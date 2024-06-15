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

	rateLimitedHandler := middlerwares.RateLimiter(http.HandlerFunc(api.HandleGenerateIdenticon))
	loggedHanlder := middlerwares.LogRequest(rateLimitedHandler)

	router.Handle("/", loggedHanlder)
	addr := fmt.Sprintf("0.0.0.0:%s", HTTP_PORT)
	err := http.ListenAndServe(addr, router)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
