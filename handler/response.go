package handler

import (
	"net/http"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/gin-gonic/gin"
)

func WriteJSON(ctx *gin.Context, data interface{}, err error) {
	resp := &response{
		Data:  data,
		Error: err,
	}

	if err != nil {
		er, ok := err.(*errors.Error)
		if ok {
			resp.StatusCode = er.StatusCode
		} else {
			resp.StatusCode = http.StatusInternalServerError
		}

	} else {
		resp.StatusCode = http.StatusOK
	}

	ctx.AbortWithStatusJSON(resp.StatusCode, resp)
}

type response struct {
	Data       interface{} `json:"data,omitempty"`
	Error      error       `json:"error,omitempty"`
	StatusCode int         `json:"-"`
}
