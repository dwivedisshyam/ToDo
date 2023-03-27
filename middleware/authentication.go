package middleware

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/handler"
	"github.com/dwivedisshyam/todo/service"
	"github.com/gin-gonic/gin"
)

func Auth(userSrvc service.User) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if exemptPath(ctx.Request) {
			ctx.Next()
			return
		}

		auth := ctx.Request.Header.Get("Authorization")
		if auth == "" {
			handler.WriteJSON(ctx, nil, errors.Unauthorized("Unauthoried user"))

			return
		}

		bearer := strings.Split(auth, " ")
		if len(bearer) < 2 || bearer[1] == "" {
			handler.WriteJSON(ctx, nil, errors.Unauthorized("Unauthoried user"))
			return
		}

		tokenString := bearer[1]

		payload, err := userSrvc.ValidateToken(ctx, tokenString)
		if err != nil {
			handler.WriteJSON(ctx, nil, err)
			return
		}

		if strconv.Itoa(int(payload.User.ID)) != ctx.Param("id") {
			handler.WriteJSON(ctx, nil, errors.Unauthorized("Forbidden"))
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
		{http.MethodPost, "/login"},
		// ToDo: Anyone can create a user
		{http.MethodPost, "/users"},
		{http.MethodGet, "/swagger"},
		{http.MethodGet, "/.well-known/openapi.json"},
	}

	for _, p := range path {
		if r.Method == p.method && strings.HasPrefix(r.URL.Path, p.url) {
			return true
		}
	}

	return false
}
