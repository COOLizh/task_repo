package main

import (
	"context"
	"io"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	config "github.com/COOLizh/task_repo/configs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"

	"github.com/COOLizh/task_repo/pkg/db"
	"github.com/COOLizh/task_repo/pkg/routers"
)

// init function
func init() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env found, looking for OS environment")
	}
}

func main() {
	conf := config.New()
	conn, pool, err := db.Connect(conf.DatabaseURL)
	if err != nil {
		log.Println(err)
	}

	defer pool.Close()
	defer conn.Release()

	router := routers.SetupRouter()
	log.Info("Service starting on port " + conf.APIPort)
	f, _ := os.Create(conf.FileLogName)
	gin.DefaultWriter = io.MultiWriter(f)

	srv := &http.Server{
		Addr:    conf.APIPort,
		Handler: router,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall.SIGKILL but can't be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}
}
