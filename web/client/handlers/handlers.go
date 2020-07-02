package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"git.sgu.ru/ultramarine/logserver"
	"git.sgu.ru/ultramarine/logserver/web/client/server"

	"github.com/go-ldap/ldap"
	log "github.com/sirupsen/logrus"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLogin(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	http.ServeFile(w, r, "public/index.html")
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	session, _ := server.Core.Store.Get(r, "logviewer_session")

	if r.Method == "GET" {
		http.ServeFile(w, r, "public/html/login.html")
	} else if r.Method == "POST" {
		r.ParseForm()

		if alreadyLogin(r) {
			http.Redirect(w, r, "/", http.StatusFound)
			return
		}

		userName, err := auth(r.FormValue("uname"), r.FormValue("psw"))
		if err != nil {
			http.Redirect(w, r, "/login", http.StatusFound)
			return
		}

		session.Values["userName"] = userName
		session.Save(r, w)
		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func GetDHCPLogsHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLogin(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	var req logserver.DHCPLogsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errMsg := fmt.Sprintf("failed to decode request for DHCP logs: %v", err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var ctx = context.Background()

	logs, err := server.Core.Svc.GetDHCPLogs(ctx, req.MAC, req.From, req.To)
	if err != nil {
		errMsg := fmt.Sprintf("failed to get DHCP logs of %s mac: %v", req.MAC, err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	logsJSON, err := json.Marshal(logs)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal DHCP logs of %s to JSON: %v", req.MAC, err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(logsJSON)
}

func GetSwitchLogsHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLogin(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	var req logserver.SwitchLogsRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errMsg := fmt.Sprintf("failed to decode request for Switch logs: %v", err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var ctx = context.Background()

	logs, err := server.Core.Svc.GetSwitchLogs(ctx, req.Name, req.From, req.To)
	if err != nil {
		errMsg := fmt.Sprintf("failed to get Switch logs of %s: %v", req.Name, err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	logsJSON, err := json.Marshal(logs)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal Switch logs of %s to JSON: %v", req.Name, err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(logsJSON)
}

func GetSimilarSwitchesHandler(w http.ResponseWriter, r *http.Request) {
	if !alreadyLogin(r) {
		http.Redirect(w, r, "/login", http.StatusUnauthorized)
		return
	}

	var req logserver.SimilarSwitchesRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		errMsg := "failed to decode request for similar switches"

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusBadRequest)
		return
	}

	var ctx = context.Background()

	switches, err := server.Core.Svc.GetSimilarSwitches(ctx, req.Name)
	if err != nil {
		errMsg := fmt.Sprintf("failed to get similar to %s switches: %v", req.Name, err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	switchesJSON, err := json.Marshal(switches)
	if err != nil {
		errMsg := fmt.Sprintf("failed to marshal similar to %s switches to JSON: %v", req.Name, err)

		log.Warnf(errMsg)

		http.Error(w, errMsg, http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(switchesJSON)
}

func alreadyLogin(r *http.Request) bool {
	//session, _ := server.Core.Store.Get(r, "logviewer_session")
	//return session.Values["userName"] != nil
	return true
}

func auth(login, password string) (string, error) {
	if password == "" {
		return "", errors.New("Empty password")
	}

	username := ""

	l, err := ldap.Dial("tcp", server.Conf.LDAP.Host)
	if err != nil {
		return username, err
	}
	defer l.Close()

	if l.Bind(server.Conf.LDAP.BindDN, server.Conf.LDAP.BindPass); err != nil {
		return username, err
	}

	searchRequest := ldap.NewSearchRequest(
		server.Conf.LDAP.BaseDN,
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		"(&(sAMAccountName="+login+"))",
		[]string{"cn"},
		nil,
	)

	sr, err := l.Search(searchRequest)
	if err != nil || len(sr.Entries) != 1 {
		return username, errors.New("User not found")
	}

	username = sr.Entries[0].GetAttributeValue("cn")

	if err = l.Bind(username, password); err != nil {
		return "", err
	}

	return username, err
}
