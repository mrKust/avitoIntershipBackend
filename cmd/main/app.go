package main

import (
	"avitoIntershipBackend/internal/config"
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

	logger.Info("register user handler")
	handler := user.NewHandler(*logger)
	handler.Register(router)

	postgresSQLClient, err := postgresql.NewClient(context.Background(), 3, cfg.Storage)
	if err != nil {
		logger.Fatal(fmt.Errorf("can't connect to database due to error: %v", err))
	}

	userRepository := userDB.NewRepository(postgresSQLClient, logger)

	err = userRepository.Delete(context.Background(), "12")
	if err != nil {
		return
	}

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
