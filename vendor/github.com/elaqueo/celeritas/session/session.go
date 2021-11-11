package session

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/alexedwards/scs/v2"
)

type Session struct {
	CookieLifetime string
	CookiePersist  string
	CookieSecure   string
	CookieName     string
	CookieDomain   string
	SessionType    string
}

func (c *Session) InitSession() *scs.SessionManager {
	var persist, secure bool

	// how long session should last
	minutes, err := strconv.Atoi(c.CookieLifetime)
	if err != nil {
		minutes = 60
	}

	// should cookie persist
	persist = strings.ToLower(c.CookiePersist) == "true"

	// must cookies be secure
	secure = strings.ToLower(c.CookieSecure) == "true"

	// create session
	session := scs.New()
	session.Lifetime = time.Duration(minutes) * time.Minute
	session.Cookie.Persist = persist
	session.Cookie.Secure = secure
	session.Cookie.Name = c.CookieName
	session.Cookie.Domain = c.CookieDomain
	session.Cookie.SameSite = http.SameSiteLaxMode

	// which session store
	switch strings.ToLower(c.SessionType) {
	case "redis":
	case "mysql", "mariadb":
	case "postgres", "postgresql":
	default:
		// cookie
	}

	return session
}
