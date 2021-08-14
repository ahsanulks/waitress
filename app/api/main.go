package main

import (
	"net/http"
	"os"
	"time"

	"github.com/ahsanulks/waitress/app/handler"
	"github.com/ahsanulks/waitress/config"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/go-rel/rel"
	"github.com/go-rel/rel/adapter/postgres"
	"github.com/subosito/gotenv"
	"go.uber.org/zap"
)

func main() {
	gotenv.Load()
	if os.Getenv("ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// load all dependencies
	zapLog, _ := zap.NewProduction(zap.AddCallerSkip(1))
	logger := config.NewLog(zapLog)
	repo, dbConnection := initDatabase(logger)
	defer dbConnection.Close()

	r := gin.New()
	r.Use(ginzap.Ginzap(zapLog, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zapLog, true))

	// load all handler
	healthzHandler := handler.NewHealthz()
	healthzHandler.AddCheck("postgresql", repo)
	healthzHandler.Mount(r.Group("/healthz"))

	server := http.Server{
		Addr:    ":" + os.Getenv("PORT"),
		Handler: r,
	}

	logger.Store(nil, "server starting on port: "+os.Getenv("PORT"), map[string]interface{}{
		"tags": []string{"server_start"},
	})

	if err := server.ListenAndServe(); err != http.ErrServerClosed {
		panic("can't start the server")
	}
}

func initDatabase(log config.Log) (rel.Repository, *postgres.Adapter) {
	conf := config.NewPostgresConf()
	return config.NewPostgresConn(conf, log)
}
