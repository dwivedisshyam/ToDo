package middleware

import (
	"net/http"
	"strings"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/handler"
	"github.com/dwivedisshyam/todo/service"
	"github.com/gin-gonic/gin"
)

func Auth(userSrvc service.User) func(*gin.Context) {
	return func(ctx *gin.Context) {
		if exemptPath(ctx.Request) {
			ctx.Next()
		}

		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			handler.Respond(ctx, nil, errors.Unauthorized("Unauthoried user"))

			return
		}

		bearer := strings.Split(auth, " ")
		if len(bearer) < 2 || bearer[1] == "" {
			handler.Respond(ctx, nil, errors.Unauthorized("Unauthoried user"))
			return
		}

		tokenString := bearer[1]

		payload, err := userSrvc.ValidateToken(ctx, tokenString)
		if err != nil {
			handler.Respond(ctx, nil, err)
			return
		}

		ctx.Set("loggedin-user", payload)
		ctx.Next()
	}
}

type endpoint struct {
	method string
	url    string
}

func exemptPath(r *http.Request) bool {
	path := []endpoint{
		{http.MethodPost, "/users"},
	}

	for _, p := range path {
		if r.Method == p.method && strings.HasPrefix(r.URL.Path, p.url) {
			return true
		}
	}

	return false
}
