package main

import (
	"database/sql"
	"flag"
	"log/slog"
	"net/http"
	"os"
	"ruhulaminjr/internal/models"

	_ "github.com/go-sql-driver/mysql"
)

type application struct {
	logger  *slog.Logger
	snippet *models.SnippetModel
}

func main() {
	addr := flag.String("addr", ":4000", "Http Network Address")
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "My sql database connection string")
	flag.Parse()
	logger := slog.New(slog.NewTextHandler(os.Stdin, &slog.HandlerOptions{
		AddSource: true,
	}))
	db, err := openDB(*dsn)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
	defer db.Close()
	app := application{
		logger: logger,
		snippet: &models.SnippetModel{
			DB: db,
		},
	}

	logger.Info("Starting Server", "addr", *addr)
	err = http.ListenAndServe(*addr, app.routes())
	logger.Error(err.Error())
	os.Exit(1)
}

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		defer db.Close()
		return nil, err
	}
	return db, nil
}
