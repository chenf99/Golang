package service

import (
	"net/http"
	"fmt"
)

const (
	StatusNotImplemented = 501
)

func Error(w http.ResponseWriter, error string, code int) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.Header().Set("X-Content-Type-Options", "nosniff")
	w.WriteHeader(code)
	fmt.Fprintln(w, error)
}

func NotImplemented(w http.ResponseWriter, r *http.Request) { Error(w, "501 page not Implemented", StatusNotImplemented) }

func NotImplementedHandler() http.Handler { return http.HandlerFunc(NotImplemented)}