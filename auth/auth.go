package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/go-ldap/ldap/v3"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type claims struct {
	jwt.StandardClaims
	Username string `json:"username"`
}

func createToken(username, password string) (string, error) {
	if err := authenticate(username, password); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 30)),
			IssuedAt:  jwt.At(time.Now()),
		},
		Username: username,
	})

	return token.SignedString([]byte(viper.GetString("jwt_key")))
}

func parseToken(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(viper.GetString("jwt_key")), nil
	})

	if customClaims, ok := token.Claims.(*claims); ok && token.Valid {
		return customClaims.Username, nil
	}

	return "", err
}

// LDAPAuthFunc is used by a middleware to authenticate requests.
func LDAPAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		log.Infof("Failed to get token from metadata: %s", err)
		return nil, err
	}

	username, err := parseToken(token)
	if err != nil {
		log.Infof("Failed to parse token: %s", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	grpc_ctxtags.Extract(ctx).Set("auth.sub", username)

	type ctxKey string
	k := ctxKey("tokenInfo")
	newCtx := context.WithValue(ctx, k, username)

	return newCtx, nil
}

// Handler handles auth endpoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	var loginInfo struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	if err := json.NewDecoder(r.Body).Decode(&loginInfo); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Infof("Failed to decode auth request: %s", err)
		return
	}

	token, err := createToken(loginInfo.Username, loginInfo.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Infof("Failed to create token for %s: %s", loginInfo.Username, err)
		return
	}

	tokenParts := strings.Split(token, ".")

	http.SetCookie(w, &http.Cookie{
		Name:   "info",
		Value:  fmt.Sprintf("%s.%s", tokenParts[0], tokenParts[1]),
		Secure: false,
	})
	http.SetCookie(w, &http.Cookie{
		Name:     "sig",
		Value:    tokenParts[2],
		Secure:   false,
		HttpOnly: true,
	})

	w.Write([]byte(fmt.Sprintf("%s.%s", tokenParts[0], tokenParts[1])))
}

// TwoCookieAuthMiddleware used for SPA auth.
func TwoCookieAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Authorization") != "" {
			next.ServeHTTP(w, r)
		} else {
			infoPart, _ := r.Cookie("info")
			sigPart, _ := r.Cookie("sig")

			r.Header.Set("Authorization", fmt.Sprintf("Bearer %s.%s", infoPart.Value, sigPart.Value))

			next.ServeHTTP(w, r)
		}
	})
}

func authenticate(login, password string) error {
	if password == "" {
		return errors.New("empty password")
	}

	l, err := ldap.Dial("tcp", viper.GetString("ldap_host"))
	if err != nil {
		return err
	}
	defer l.Close()

	if err = l.Bind(viper.GetString("ldap_bind_dn"), viper.GetString("ldap_bind_pass")); err != nil {
		return fmt.Errorf("error authenticating admin user in LDAP: %s", err)
	}

	searchRequest := ldap.NewSearchRequest(
		viper.GetString("ldap_base_dn"),
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(sAMAccountName="+login+"))",
		[]string{"cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil {
		return fmt.Errorf("error searching user in LDAP: %s", err)
	}
	if len(sr.Entries) != 1 {
		return errors.New("user not found in LDAP")
	}

	username := sr.Entries[0].GetAttributeValue("cn")

	if err = l.Bind(username, password); err != nil {
		return fmt.Errorf("error authenticating user in LDAP: %s", err)
	}

	return nil
}
