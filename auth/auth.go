package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"git.sgu.ru/ultramarine/logserver/conf"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-ldap/ldap/v3"
	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type user struct {
	Username string
}

func parseToken(tokenString string) (user, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(conf.Conf.App.JWTKey), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return user{Username: fmt.Sprintf("%v", claims["username"])}, nil
	}

	return user{}, err
}

// LDAPAuthFunc is used by a middleware to authenticate requests.
func LDAPAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		log.Infof("Failed to get token from metadata: %s", err)
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		log.Infof("Failed to parse token: %s", err)
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	grpc_ctxtags.Extract(ctx).Set("auth.sub", tokenInfo.Username)

	type ctxKey string
	k := ctxKey("tokenInfo")
	newCtx := context.WithValue(ctx, k, tokenInfo)

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

	if err := authenticate(loginInfo.Username, loginInfo.Password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		log.Infof("Failed authenticate user %s: %s", loginInfo.Username, err)
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": loginInfo.Username})
	tokenString, err := token.SignedString([]byte(conf.Conf.App.JWTKey))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Infof("Failed create token for user %s: %s", loginInfo.Username, err)
		return
	}

	w.Write([]byte(tokenString))

}

func authenticate(login, password string) error {
	if password == "" {
		return errors.New("empty password")
	}

	l, err := ldap.Dial("tcp", conf.Conf.LDAP.Host)
	if err != nil {
		return err
	}
	defer l.Close()

	if err = l.Bind(conf.Conf.LDAP.BindDN, conf.Conf.LDAP.BindPass); err != nil {
		return fmt.Errorf("error authenticating admin user in LDAP: %s", err)
	}

	searchRequest := ldap.NewSearchRequest(
		conf.Conf.LDAP.BaseDN,
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
