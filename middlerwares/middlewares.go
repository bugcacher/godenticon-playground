package middlerwares

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bugcacher/godenticon-playground/logger"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-limiter/memorystore"
)

var (
	ip_headers = []string{"RemoteAddr", "X-Forwarded-For", "X-Real-IP"}
)

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userIP, _, _ := net.SplitHostPort(r.RemoteAddr)
		for _, ipHeader := range ip_headers {
			if userIP == "" {
				userIP = r.Header.Get(ipHeader)
			}
		}
		logger.DefaultLogger.Info("request received", "ip", userIP, "request", r.URL.Query().Encode())
		next.ServeHTTP(w, r)
	})
}

func RateLimiter(next http.Handler) http.Handler {
	// Initialize in memory store for rate limiter
	store, err := memorystore.New(&memorystore.Config{
		Tokens:   10,
		Interval: 1 * time.Minute,
	})
	if err != nil {
		log.Fatal("failed to initialize memory store for rate limiter")
	}
	// Initiliaze rate limiter middleware with in-memory store
	middlerware, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc(ip_headers...))
	if err != nil {
		log.Fatal("failed to initialize rate limiter middleware")
	}

	return middlerware.Handle(next)
}
