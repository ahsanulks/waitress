package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahsanulks/waitress/app/handler"
	"github.com/gin-gonic/gin"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// load all handler and dependencies
	healthzHandler := handler.NewHealthz()
	healthzHandler.Mount(r.Group("/healthz"))

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	log.Println("server starting on port: " + os.Getenv("PORT"))
	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic("can't start the server")
	}
}
