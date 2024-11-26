package main

import (
	"net/http"
	"os"
	"trigonal/backend-auth/api/database"
	"trigonal/backend-auth/api/helper"
	"trigonal/backend-auth/api/routers"

	"github.com/gin-gonic/gin"
)

func init() {
	helper.LoadEnv()
	database.ConnectDB()
	database.SyncDB()
}

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		helper.SUCCESS(ctx, http.StatusOK, "API successfully run at port "+os.Getenv("PORT"))
	})
	routers.AuthRoutes(r)
	routers.UserRoutes(r)

	r.Run()
}