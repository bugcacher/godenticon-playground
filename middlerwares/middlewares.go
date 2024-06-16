package middlerwares

import (
	"log"
	"net"
	"net/http"
	"os"
	"slices"
	"strings"
	"time"

	"github.com/bugcacher/godenticon-playground/logger"
	"github.com/sethvargo/go-limiter/httplimit"
	"github.com/sethvargo/go-limiter/memorystore"
)

const (
	allowedHosts = "ALLOWED_HOSTS"
)

var (
	ipHeaders = []string{"X-Real-IP", "X-Forwarded-For"}
)

type Middleware func(http.Handler) http.Handler

func ApplyMiddlewares(h http.Handler, middlewares ...Middleware) http.Handler {
	for _, middleware := range middlewares {
		h = middleware(h)
	}
	return h
}

func LogRequest(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var userIP string
		for _, ipHeader := range ipHeaders {
			if userIP == "" {
				userIP = r.Header.Get(ipHeader)
			}
		}
		if userIP == "" {
			userIP, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
		logger.DefaultLogger.Info("request received", "ip", userIP, "query_params", r.URL.Query().Encode(), "headers", r.Header)
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
	middlerware, err := httplimit.NewMiddleware(store, httplimit.IPKeyFunc(ipHeaders...))
	if err != nil {
		log.Fatal("failed to initialize rate limiter middleware")
	}

	return middlerware.Handle(next)
}

func EnableCORS(next http.Handler) http.Handler {
	// List of hosts enabled for cors
	allowedHostsForCors := os.Getenv(allowedHosts)
	allowedHostsForCorsList := strings.Split(allowedHostsForCors, ",")
	for idx, host := range allowedHostsForCorsList {
		allowedHostsForCorsList[idx] = strings.TrimSpace(host)
	}

	logger.DefaultLogger.Info("Allowed hosts", "hosts", allowedHostsForCorsList)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		origin := r.Header.Get("Origin")
		if slices.Contains(allowedHostsForCorsList, origin) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		// Handle preflight OPTIONS request
		if r.Method == "OPTIONS" {
			return
		}
		next.ServeHTTP(w, r)
	})
}
