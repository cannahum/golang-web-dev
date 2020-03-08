package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

type SessionHandler func(http.ResponseWriter, *http.Request, string)

type SessionMiddleware struct {
	handler SessionHandler
}

func (sh *SessionMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c, err := r.Cookie("session")
	if err != nil {
		log.Println("Error getting cookie", err.Error)

		sessionID, _ := uuid.NewV4()
		c = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
	}

	sh.handler(w, r, c.Value)
}

func NewEnsureSession(handlerToWrap SessionHandler) *SessionMiddleware {
	return &SessionMiddleware{handlerToWrap}
}
