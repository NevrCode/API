package middleware

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func RequireAuth(c *gin.Context) {
	c.Next()
}
func SessionMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := Store.Get(c.Request, "session-name")
		print(session.ID)
		c.Set("session", session)
		c.Next()
	}
}
