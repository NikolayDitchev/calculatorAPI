package main

import (
	"net/http"
	"strings"
)

type middleware func(http.HandlerFunc) http.HandlerFunc

var middlewareFuncs []middleware = []middleware{
	bodyLenghtCheck,
	requireContentType,
}

func middlewareChain(hf http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, req *http.Request) {
		for i := len(middlewareFuncs) - 1; i >= 0; i-- {
			hf = middlewareFuncs[i](hf)
		}

		hf(writer, req)
	}
}

//middleware functions declarations

func bodyLenghtCheck(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		req.Body = http.MaxBytesReader(w, req.Body, 256)
		next(w, req)
	}
}

func requireContentType(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		if req.Method == "POST" {

			conType := req.Header.Get("Content-Type")
			mediaType := strings.ToLower(strings.TrimSpace(conType))

			if mediaType != "application/json" {
				http.Error(w, "Content-Type header is not application/json", http.StatusUnsupportedMediaType)
				return
			}
		}

		next(w, req)
	}
}
