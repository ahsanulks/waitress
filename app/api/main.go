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

	cartRepository "github.com/ahsanulks/waitress/carts/repository"
	cartUsecase "github.com/ahsanulks/waitress/carts/usecase"
	productRepository "github.com/ahsanulks/waitress/products/repository"
	productUsecase "github.com/ahsanulks/waitress/products/usecase"
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

	// load relate healthz
	healthzHandler := handler.NewHealthz()
	healthzHandler.AddCheck("postgresql", repo)

	// load relate product domain
	productRepo := productRepository.NewProductRepository(repo)
	productUsecase := productUsecase.NewProductUsecase(productRepo)
	productHandler := handler.NewProductHandler(productUsecase)

	// load relate cart domain
	cartRepo := cartRepository.NewCartRepository(repo)
	cartUsecase := cartUsecase.NewCartUsecase(cartRepo)
	cartHandler := handler.NewCartHandler(cartUsecase)

	r := gin.New()
	r.Use(ginzap.Ginzap(zapLog, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(ginzap.RecoveryWithZap(zapLog, true))

	// mount all endpoint from handler
	healthzHandler.Mount(r.Group("/healthz"))
	productHandler.Mount(r.Group("/products"))
	cartHandler.Mount(r.Group("/carts"))

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
