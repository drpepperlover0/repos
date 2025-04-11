package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"

	"github.com/drpepperlover0/internal/app/service"
	"github.com/drpepperlover0/internal/app/repository"
	"github.com/drpepperlover0/internal/app/router"
	"github.com/drpepperlover0/lib"
)

const (
	addr = ":8080"
)

func main() {
	loger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	db, err := lib.ConnectDB()
	if err != nil {
		loger.Error((fmt.Sprintf("DB connect error: %v", err)))
		return
	}

	userRepo := repository.NewUserRepo(db)
	service := service.New(userRepo)
	server := &http.Server{
		Addr:    addr,
		Handler: router.InitRoutes(service),
	}

	loger.Info("Server is starting on port 8080")
	if err := server.ListenAndServe(); err != nil {
		loger.Error((fmt.Sprintf("Starting server error: %v", err)))
		return
	}
}
