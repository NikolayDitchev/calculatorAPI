package main

import "net/http"

type APIserver struct {
	addr string
}

func NewAPIServer(addr string) *APIserver {
	return &APIserver{
		addr: addr,
	}
}

func (s *APIserver) Run() error {

	router := http.NewServeMux()
	router.HandleFunc("POST /evaluate", middlewareChain(evaluateExpression))
	router.HandleFunc("GET /evaluate", getHtmlFile)
	
	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}

	return server.ListenAndServe()
}
