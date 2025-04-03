package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"restapi/internal/config"
	"restapi/internal/user"
	"restapi/internal/user/db"
	"restapi/pkg/client/mongodb"
	"restapi/pkg/logging"
	"time"

	"github.com/julienschmidt/httprouter"
)

func main() {
	logger := logging.GetLogger() //Info< Debug< Trace
	logger.Info("create router")
	router := httprouter.New()

	cfg := config.GetConfig()

	cfgMongo := cfg.MongoDB
	mongoDBClient, err := mongodb.NewClient(context.Background(), cfgMongo.Host, cfgMongo.Port, cfgMongo.Username, cfgMongo.Password, cfgMongo.Database, cfgMongo.AuthDB)
	if err != nil {
		panic(err)
	}

	storage := db.NewStorage(mongoDBClient, cfgMongo.Collection, logger)

	user1 := user.User{
		ID:           "",
		Email:        "AlexExample@mail.ru",
		Username:     "Alexander",
		PasswordHash: "12345",
	}
	user1ID, err := storage.Create(context.Background(), user1)
	if err != nil {
		panic(err)
	}
	logger.Info(user1ID)

	logger.Info("register user handler")
	handler := user.NewHandler(logger) //во всех обработчиках доступен логгер
	handler.Register(router)

	start(router, cfg)
}

func start(router *httprouter.Router, cfg *config.Config) {
	logger := logging.GetLogger()
	logger.Info("start application")

	var listener net.Listener
	var listenErr error

	if cfg.Listen.Type == "sock" {
		// /path/to/binary
		// Dir() -- /path/to
		logger.Info("detect app path")
		appDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
		if err != nil {
			logger.Fatal(err)
		}
		logger.Info("create socket")
		socketPath := path.Join(appDir, "app.sock")

		logger.Info("listen unix socket")
		listener, listenErr = net.Listen("unix", socketPath)
		logger.Info("server is listening unix socket %s", socketPath)
	} else {
		logger.Info("listen tcp")
		listener, listenErr = net.Listen("tcp", fmt.Sprintf("%s:%s", cfg.Listen.BindIP, cfg.Listen.Port))
		logger.Info("server is listening port %s:%s", cfg.Listen.BindIP, cfg.Listen.Port)
	}

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		panic(err)
	}

	if listenErr != nil {
		logger.Fatal(listenErr)
	}

	server := &http.Server{
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	logger.Fatal(server.Serve(listener))
}
