package main

import (
	"bytes"
	"fmt"
	"log/slog"
	"net/http"
	"runtime/debug"
	"time"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	var (
		method = r.Method
		uri    = r.URL.RequestURI()
		trace  = string(debug.Stack())
	)
	app.logger.Error(err.Error(), slog.Any("method", method), slog.Any("url", uri), slog.Any("trace", trace))
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)

}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, page string, status int, data TemplateData) {

	ts, ok := app.templateCache[page]
	if !ok {
		err := fmt.Errorf("the page %s not exits", page)
		app.serverError(w, r, err)
		return
	}

	buf := new(bytes.Buffer)
	if err := ts.ExecuteTemplate(buf, "base", data); err != nil {
		app.serverError(w, r, err)
	}
	w.WriteHeader(status)
	buf.WriteTo(w)

}

func (app *application) newTemplateData(r *http.Request) TemplateData {
	return TemplateData{
		CurrentYear: time.Now().Year(),
	}
}
