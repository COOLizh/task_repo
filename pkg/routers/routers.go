// Package routers with routing
package routers

import (
	"fmt"
	"time"

	"github.com/COOLizh/task_repo/pkg/handlers"
	"github.com/COOLizh/task_repo/pkg/middleware"
	"github.com/gin-gonic/gin"
)

// SetupRouter - all endpoints
func SetupRouter() *gin.Engine {
	gin.ForceConsoleColor()
	router := gin.New()
	router.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
		return fmt.Sprintf("%s - [%s] \"%s %s %s %d \n",
			param.ClientIP,
			param.TimeStamp.Format(time.RFC1123),
			param.Method,
			param.Path,
			param.Request.Proto,
			param.StatusCode,
		)
	}))
	router.Use(gin.Recovery())

	v1 := router.Group("/api/v1")
	v1.GET("/", handlers.HealthCheck)
	v1.POST("/registration", handlers.Registration)
	v1.POST("/login", handlers.Login)
	v1.GET("/tasks", handlers.GetAllTasks)
	v1.Use(middleware.AuthMiddleware()).GET("/auth", handlers.AuthCheck)
	v1.Use(middleware.AuthMiddleware()).POST("/tasks/:task_id", handlers.SendSolutionHandler)
	return router
}
