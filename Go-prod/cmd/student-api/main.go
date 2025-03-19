package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/anmol1377/student-api/internal/config"
	"github.com/anmol1377/student-api/internal/http/handlers/student"
	"github.com/anmol1377/student-api/internal/storage/sqlite"
)

func main() {

	cfg := config.MustLoad()
	// fmt.Println(cfg)
	// fmt.Println("cjscbjhsb")

	// db below

	storage, err := sqlite.New(cfg)

	if err != nil {
		log.Fatal(err)
	}

	slog.Info("Storage init", slog.String("ENV", cfg.ENV), slog.String("ver", "1.0.0"))

	router := http.NewServeMux()
	router.HandleFunc("POST /api/students", student.New(storage))
	router.HandleFunc("GET /api/students/{id}", student.GetID(storage))

	server := http.Server{
		Addr:    cfg.Address,
		Handler: router,
	}

	slog.Info("started")
	// fmt.Println("started")

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		err := server.ListenAndServe()
		if err != nil {
			log.Fatal("failed to start server: ", err)
		}
	}()

	<-done

	slog.Info("shut down server")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = server.Shutdown(ctx)

	if err != nil {
		slog.Error("failed shutdown", slog.String("error", err.Error()))
	}
	slog.Info("server shutdown sucessful")
}
