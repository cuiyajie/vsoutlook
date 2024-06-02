package svcinfra

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
	"vsoutlook.com/vsoutlook/infra/config"
)

var SessionKey = "yfak"
var SessionUIDKey = "u"
var sessionStore *sessions.CookieStore

func NewCookieStore(keyPairs ...[]byte) *sessions.CookieStore {
	maxAge := 86400 * 90
	opts := sessions.Options{
		Path:   "/",
		MaxAge: maxAge,

		// the following(including SameSite=None) are needed only
		//     when we call api with cors in local frontend development mode
		// not that needed for production
		// https://stackoverflow.com/a/67001424
		HttpOnly: true,
	}

	cs := &sessions.CookieStore{
		Codecs:  securecookie.CodecsFromPairs(keyPairs...),
		Options: &opts,
	}
	cs.MaxAge(cs.Options.MaxAge)
	return cs
}

func GetSession(r *http.Request) *sessions.Session {
	if sessionStore == nil {
		sessionStore = NewCookieStore(
			[]byte(config.Get("SESSION_KEY")),
			[]byte(config.Get("SESSION_SECRET")))
	}

	session, _ := sessionStore.Get(r, SessionKey)
	return session
}

func sessGetString(s *sessions.Session, key string) string {
	if s == nil {
		return ""
	}
	v, exist := s.Values[key]
	if !exist {
		return ""
	}
	return v.(string)
}

func sessSetString(c *Context, s *sessions.Session, key, value string) error {
	s.Values[key] = value
	if err := s.Save(c.Request, c.Writer); err != nil {
		return fmt.Errorf("sess set: %s, %w", key, err)
	}
	return nil
}

func SetSessionUser(c *Context, uid string) bool {
	if err := sessSetString(c, c.Session, SessionUIDKey, uid); err != nil {
		log.Printf("set sess err: %s", err)
		c.GeneralError("")
		return false
	}
	return true
}
