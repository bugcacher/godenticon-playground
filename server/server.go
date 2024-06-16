package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bugcacher/godenticon-playground/api"
	"github.com/bugcacher/godenticon-playground/middlerwares"
)

const (
	HTTP_PORT = "8000"
)

func Serve() {
	router := http.NewServeMux()

	router.Handle("/", http.HandlerFunc(api.HandleGenerateIdenticon))
	router.Handle("/health", http.HandlerFunc(api.HandleHealth))

	routerWithMiddlewares := middlerwares.ApplyMiddlewares(router, middlerwares.LogRequest, middlerwares.RateLimiter, middlerwares.EnableCORS)

	addr := fmt.Sprintf("0.0.0.0:%s", HTTP_PORT)
	err := http.ListenAndServe(addr, routerWithMiddlewares)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
