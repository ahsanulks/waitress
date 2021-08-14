package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ahsanulks/waitress/app/handler"
	"github.com/ahsanulks/waitress/config"
	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	"github.com/subosito/gotenv"
)

func main() {
	gotenv.Load()
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// load all handler and dependencies
	repo, dbConnection := initDatabase()
	defer dbConnection.Close()
	healthzHandler := handler.NewHealthz()
	healthzHandler.AddCheck("postgresql", repo)
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

func initDatabase() (rel.Repository, *postgres.Adapter) {
	conf := config.NewPostgresConf()
	return config.NewPostgresConn(conf)
}
