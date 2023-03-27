package main

import (
	"net/http"

	"github.com/dwivedisshyam/todo/db"
	"github.com/dwivedisshyam/todo/handler"
	"github.com/dwivedisshyam/todo/middleware"
	"github.com/dwivedisshyam/todo/service"
	"github.com/dwivedisshyam/todo/store"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db := db.Connect()

	userSt := store.NewUser(db)
	taskSt := store.NewTask(db)

	taskSvc := service.NewTask(taskSt)
	userSvc := service.NewUser(userSt)

	r.Use(middleware.Auth(userSvc))

	uh := handler.NewUser(userSvc)
	th := handler.NewTask(taskSvc)

	r.POST("/login", uh.Login)

	r.POST("/users", uh.Create)
	r.GET("/users", uh.List)
	r.GET("/users/:id", uh.Get)
	r.PUT("/users/:id", uh.Update)
	r.DELETE("/users/:id", uh.Delete)

	r.POST("users/:id/tasks", th.Create)
	r.GET("users/:id/tasks", th.List)
	r.GET("users/:id/tasks/:id", th.Get)
	r.PUT("users/:id/tasks/:id", th.Update)

	// fs := http.FileServer(http.Dir("./swagger"))

	r.StaticFile("/.well-known/openapi.json", "./doc/openapi.json")
	r.StaticFS("/swagger", http.Dir("./swagger"))

	r.Run(":8000")
}
