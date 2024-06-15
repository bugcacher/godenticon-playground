package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bugcacher/godenticon-playground/api"
	"github.com/bugcacher/godenticon-playground/middlerwares"
)

const (
	HTTP_PORT = 8000
)

func Serve() {
	router := http.NewServeMux()

	rateLimitedHandler := middlerwares.RateLimiter(http.HandlerFunc(api.HandleGenerateIdenticon))
	loggedHanlder := middlerwares.LogRequest(rateLimitedHandler)

	router.Handle("/", loggedHanlder)

	err := http.ListenAndServe(fmt.Sprintf(":%d", HTTP_PORT), router)
	if err != nil {
		log.Fatal("failed to start server")
	}
}
