package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	addr := flag.String("addr", ":4000", "Http Network Address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
		AddSource: true,
	}))
	app := application{
		logger: logger,
	}

	logger.Info("Starting Server", "addr", *addr)
	err := http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}
