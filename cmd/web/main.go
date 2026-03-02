package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	app := &application{
		logger: logger,
	}

	// addr := flag.String("addr", ":4000", "HTTP network address")
	// flag.Parse()
	port := flag.Int("port", 4000, "HTTP network address")
	flag.Parse()

	mux := app.routes()

	addr := ":" + fmt.Sprintf("%v", *port)
	logger.Info("Starting server", "port", addr)
	err := http.ListenAndServe(addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
