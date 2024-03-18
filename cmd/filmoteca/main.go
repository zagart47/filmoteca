package main

import (
	"filmoteca/internal/config"
	"filmoteca/internal/repository"
	"filmoteca/internal/repository/postgresql"
	"filmoteca/internal/service"
	"filmoteca/pkg/logger"
	"filmoteca/transport/http"
	"filmoteca/transport/http/handler"
)

func main() {
	db, err := postgresql.New(config.All.PostgreSQL.DSN)
	if err != nil {
		logger.Log.Error("cannot connect to db: ", err.Error())
	}
	logger.Log.Info("connected to db", "host", config.All.PostgreSQL.DSN)
	repos := repository.NewRepositories(db)
	logger.Log.Info("repositories added")
	services := service.NewServices(repos)
	logger.Log.Info("services added")
	h := handler.NewHandler(services)
	mux := h.Init()
	s := http.NewServer(config.All.HTTP.Host)
	logger.Log.Info("starting http server", "host", config.All.HTTP.Host)
	if err := s.Run(mux); err != nil {
		logger.Log.Error("cannot start server: ", err.Error())
	}
}
