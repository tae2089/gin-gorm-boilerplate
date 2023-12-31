package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/tae2089/gin-boilerplate/common/cache"
	"github.com/tae2089/gin-boilerplate/common/config"
	"github.com/tae2089/gin-boilerplate/common/router"
	"github.com/tae2089/gin-boilerplate/notification"
	"github.com/tae2089/gin-boilerplate/user/model"
)

func main() {
	e := gin.New()
	db := config.NewDBConfig()
	jwtKey := config.NewJwtKey()
	gin.SetMode(gin.ReleaseMode)
	db.AutoMigrate(&model.User{})
	redisStore := cache.NewRedisStore(redis.NewClient(&redis.Options{
		Network: "tcp",
		Addr:    "127.0.0.1:6379",
	}))

	go notification.Listen(notification.UseLoggerProvider)
	router.SetupRouter(e, db, jwtKey, redisStore)
	srv := &http.Server{
		Addr:    ":8080",
		Handler: e,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")

}
