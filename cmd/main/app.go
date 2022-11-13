package main

import (
	"avitoIntershipBackend/internal/config"
	masterBalDB "avitoIntershipBackend/internal/masterBalance/db"
	serviceDB "avitoIntershipBackend/internal/service/db"
	transactionDB "avitoIntershipBackend/internal/transaction/db"
	"avitoIntershipBackend/internal/user"
	userDB "avitoIntershipBackend/internal/user/db"
	"avitoIntershipBackend/pkg/client/postgresql"
	"avitoIntershipBackend/pkg/logging"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"time"
)

func main() {

	logger := logging.GetLogger()
	logger.Info("create router")
	router := gin.Default()

	cfg := config.GetConfig()

	postgresSQLClient, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err != nil {
		logger.Fatal(fmt.Errorf("can't connect to database due to error: %v", err))
	}

	userRepository := userDB.NewRepository(postgresSQLClient, logger)
	transactionRepository := transactionDB.NewRepository(postgresSQLClient, logger)
	serviceRepository := serviceDB.NewRepository(postgresSQLClient, logger)
	masterBalRepository := masterBalDB.NewRepository(postgresSQLClient, logger)

	logger.Info("create service and repositories")
	serv := user.NewService(userRepository, masterBalRepository, transactionRepository, serviceRepository, logger)

	logger.Info("register user handler")
	handler := user.NewHandler(*logger, serv)
	handler.Register(router)

	start(router, cfg)

}

func start(router *gin.Engine, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("Start application")

	var listener net.Listener
	var listenerError error

	if cfg.Listen.Type == "port" {
		logger.Info("listen tcp")
		listener, listenerError = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Infof("server is listening on port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	} else {
		panic("Wrong config")
	}

	if listenerError != nil {
		logger.Fatal(listenerError)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
