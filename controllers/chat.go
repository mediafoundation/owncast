package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/owncast/owncast/core"
	"github.com/owncast/owncast/core/user"
	"github.com/owncast/owncast/router/middleware"
	log "github.com/sirupsen/logrus"
)

// GetChatMessages gets all of the chat messages.
func GetChatMessages(w http.ResponseWriter, r *http.Request) {
	middleware.EnableCors(&w)
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case http.MethodGet:
		messages := core.GetAllChatMessages()

		err := json.NewEncoder(w).Encode(messages)
		if err != nil {
			log.Errorln(err)
		}
	default:
		w.WriteHeader(http.StatusNotImplemented)
		if err := json.NewEncoder(w).Encode(j{"error": "method not implemented (PRs are accepted)"}); err != nil {
			InternalErrorHandler(w, err)
		}
	}
}

// RegisterAnonymousChatUser will register a new user.
func RegisterAnonymousChatUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	if r.Method != POST {
		WriteSimpleResponse(w, false, r.Method+" not supported")
		return
	}

	type registerAnonymousUserResponse struct {
		AccessToken string `json:"accessToken"`
		DisplayName string `json:"displayName"`
	}

	err, newUser := user.CreateAnonymousUser()
	if err != nil {
		WriteSimpleResponse(w, false, err.Error())
		return
	}

	response := registerAnonymousUserResponse{
		AccessToken: newUser.AccessToken,
		DisplayName: newUser.DisplayName,
	}

	WriteResponse(w, response)
}
