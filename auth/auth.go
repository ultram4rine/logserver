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
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func parseToken(token string) (struct{}, error) {
	return struct{}{}, nil
}

func userClaimFromToken(struct{}) string {
	return "foobar"
}

// LDAPAuthFunc is used by a middleware to authenticate requests.
func LDAPAuthFunc(ctx context.Context) (context.Context, error) {
	token, err := grpc_auth.AuthFromMD(ctx, "bearer")
	if err != nil {
		return nil, err
	}

	tokenInfo, err := parseToken(token)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid auth token: %v", err)
	}

	grpc_ctxtags.Extract(ctx).Set("auth.sub", userClaimFromToken(tokenInfo))

	// WARNING: in production define your own type to avoid context collisions
	newCtx := context.WithValue(ctx, "tokenInfo", tokenInfo)

	return newCtx, nil
}

// Handler handles auth endpoint.
func Handler(w http.ResponseWriter, r *http.Request) {
	var user struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := authenticate(user.Username, user.Password); err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	} else {
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{"username": user.Username})
		tokenString, err := token.SignedString(conf.Conf.App.JWTKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write([]byte(tokenString))
	}
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
