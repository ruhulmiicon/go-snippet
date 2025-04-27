package main

import (
	"flag"
	"log/slog"
	"net/http"
	"os"
)

func main() {
	addr := flag.String("addr", ":4000", "Http Network Address")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
		AddSource: true,
	}))
	mux := http.NewServeMux()
	fileServe := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServe))
	mux.HandleFunc("GET /{$}", home)
	mux.HandleFunc("GET /snippet/view/{$}", snippetView)
	mux.HandleFunc("GET /snippet/create", snippetCreate)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("Starting Server", "addr", *addr)
	err := http.ListenAndServe(*addr, mux)
	logger.Error(err.Error())
	os.Exit(1)
}
