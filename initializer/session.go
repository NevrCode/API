package initializer

import (
	"os"

	"github.com/gorilla/sessions"
)

const SESSION_ID = "SDHLFSJASLAF"

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))
