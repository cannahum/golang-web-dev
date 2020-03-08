package main

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"
)

// SessionHandler is a function that handles a request and also takes a 3rd parameter of cookie value (string)
type SessionHandler func(http.ResponseWriter, *http.Request, *http.Cookie)

// SessionMiddleware calls the handler (of type SessionHandler) with cookie
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

		http.SetCookie(w, c)
	}

	sh.handler(w, r, c)
}

// NewEnsureSession returns a new sessionMiddleware
func NewEnsureSession(handlerToWrap SessionHandler) *SessionMiddleware {
	return &SessionMiddleware{handlerToWrap}
}
