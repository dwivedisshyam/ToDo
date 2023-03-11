package main

import (
	"github.com/dwivedisshyam/todo/db"
	taskH "github.com/dwivedisshyam/todo/handler/task"
	userH "github.com/dwivedisshyam/todo/handler/user"
	"github.com/dwivedisshyam/todo/middleware"
	taskSrvc "github.com/dwivedisshyam/todo/service/task"
	userSrvc "github.com/dwivedisshyam/todo/service/user"
	taskStore "github.com/dwivedisshyam/todo/store/task"
	userStore "github.com/dwivedisshyam/todo/store/user"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := db.Connect()

	userSt := userStore.New(db)
	taskSt := taskStore.New(db)

	taskSvc := taskSrvc.New(taskSt)
	userSvc := userSrvc.New(userSt)

	r.Use(middleware.Auth(userSvc))

	uh := userH.New(userSvc)
	th := taskH.New(taskSvc)

	r.POST("/login", uh.Login)

	r.POST("/users", uh.Create)
	r.GET("/users", uh.List)
	r.GET("/users/:id", uh.Get)
	r.PUT("/users/:id", uh.Update)

	r.POST("/tasks", th.Create)
	r.GET("/tasks", th.List)
	r.GET("/tasks/:id", th.Get)
	r.PUT("/tasks/:id", th.Update)

	r.Run(":8000")
}
