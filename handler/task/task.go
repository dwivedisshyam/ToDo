package task

import (
	"strconv"

	"github.com/dwivedisshyam/go-lib/pkg/errors"
	"github.com/dwivedisshyam/todo/handler"
	"github.com/dwivedisshyam/todo/model"
	"github.com/dwivedisshyam/todo/service"
	"github.com/gin-gonic/gin"
)

type taskHandler struct {
	srvc service.Task
}

func New(us service.Task) *taskHandler {
	return &taskHandler{srvc: us}
}

func (uh taskHandler) Create(ctx *gin.Context) {
	var t model.Task

	err := ctx.BindJSON(&t)
	if err != nil {
		handler.Respond(ctx, nil, errors.BadRequest(err.Error()))
		return
	}

	errSrvc := uh.srvc.Create(&t)
	if errSrvc != nil {
		handler.Respond(ctx, nil, errSrvc)
		return
	}

	handler.Respond(ctx, nil, nil)
}

func (uh taskHandler) Update(ctx *gin.Context) {
	var t model.Task

	id, ok := ctx.Params.Get("id")
	if !ok {
		handler.Respond(ctx, nil, errors.Validation("id missing"))
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.Respond(ctx, nil, errors.Unexpected(err.Error()))
		return
	}

	err = ctx.BindJSON(&t)
	if err != nil {
		handler.Respond(ctx, nil, errors.BadRequest(err.Error()))
		return
	}

	t.ID = userID

	errSrvc := uh.srvc.Update(&t)
	if errSrvc != nil {
		handler.Respond(ctx, nil, errSrvc)
		return
	}

	handler.Respond(ctx, nil, nil)
}

func (uh taskHandler) List(ctx *gin.Context) {
	tasks, err := uh.srvc.List()
	if err != nil {
		handler.Respond(ctx, nil, err)
		return
	}

	handler.Respond(ctx, tasks, err)
}

func (uh taskHandler) Get(ctx *gin.Context) {
	id, ok := ctx.Params.Get("id")
	if !ok {
		handler.Respond(ctx, nil, errors.Validation("id missing"))
		return
	}

	userID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		handler.Respond(ctx, nil, errors.Unexpected(err.Error()))
		return
	}

	user, srvcErr := uh.srvc.Get(userID)
	if srvcErr != nil {
		handler.Respond(ctx, nil, srvcErr)
		return
	}

	handler.Respond(ctx, user, nil)
}
