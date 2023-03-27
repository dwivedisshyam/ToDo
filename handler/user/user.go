package user

import (
	"strconv"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/handler"
	"github.com/dwivedisshyam/todo/model"
	"github.com/dwivedisshyam/todo/service"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	srvc service.User
}

func New(us service.User) *userHandler {
	return &userHandler{srvc: us}
}

func (uh userHandler) Create(ctx *gin.Context) {
	var u model.User

	err := ctx.BindJSON(&u)
	if err != nil {
		handler.WriteJSON(ctx, nil, errors.BadRequest(err.Error()))
		return
	}

	user, err := uh.srvc.Create(&u)
	if err != nil {
		handler.WriteJSON(ctx, nil, err)
		return
	}

	handler.WriteJSON(ctx, user, nil)
}

func (uh userHandler) Update(ctx *gin.Context) {
	var u model.User

	id, ok := ctx.Params.Get("id")
	if !ok {
		handler.WriteJSON(ctx, nil, errors.Validation("id missing"))
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.WriteJSON(ctx, nil, errors.Unexpected(err.Error()))
		return
	}

	err = ctx.BindJSON(&u)
	if err != nil {
		handler.WriteJSON(ctx, nil, errors.BadRequest(err.Error()))
		return
	}

	u.ID = userID

	errSrvc := uh.srvc.Update(&u)
	if errSrvc != nil {
		handler.WriteJSON(ctx, nil, errSrvc)
		return
	}

	handler.WriteJSON(ctx, nil, nil)
}

func (uh userHandler) List(ctx *gin.Context) {
	users, err := uh.srvc.List()
	if err != nil {
		handler.WriteJSON(ctx, nil, err)
		return
	}

	handler.WriteJSON(ctx, users, err)
}

func (uh userHandler) Get(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		handler.WriteJSON(ctx, nil, errors.Validation("id missing"))
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.WriteJSON(ctx, nil, errors.Unexpected(err.Error()))
		return
	}

	user, srvcErr := uh.srvc.Get(userID)
	if srvcErr != nil {
		handler.WriteJSON(ctx, nil, srvcErr)
		return
	}

	handler.WriteJSON(ctx, user, nil)
}

func (uh userHandler) Delete(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		handler.WriteJSON(ctx, nil, errors.Validation("id missing"))
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.WriteJSON(ctx, nil, errors.Unexpected(err.Error()))
		return
	}

	srvcErr := uh.srvc.Delete(userID)
	if srvcErr != nil {
		handler.WriteJSON(ctx, nil, srvcErr)
		return
	}

	handler.WriteJSON(ctx, nil, nil)
}

func (uh userHandler) Login(ctx *gin.Context) {
	var l model.Login

	err := ctx.BindJSON(&l)
	if err != nil {
		handler.WriteJSON(ctx, nil, errors.BadRequest(err.Error()))
		return
	}

	token, srvcErr := uh.srvc.Login(&l)
	if srvcErr != nil {
		handler.WriteJSON(ctx, nil, srvcErr)
		return
	}

	handler.WriteJSON(ctx, token, nil)
}
