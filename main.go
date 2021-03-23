package main

import (
	"context"
	"devops/handler"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("template/html/*")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"name": "shay",
		})
	})
	api := router.Group("/api")
	{
		api.GET("/group/hub", handler.GroupDockerHandler)
		api.GET("/group/project", handler.GroupProjectHandler)
		api.POST("/devops", handler.GennerateHandler)
	}

	//优雅关停
	srv := &http.Server{
		Addr:    ":3001",
		Handler: router,
	}
	go func() {
		log.Printf("start listen at -> %s", srv.Addr)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exited")
	// router.Run(":3001")
}
