package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"
)

// === Equivalent of a Java class ===
type Server struct {
	delayMs int
	sem     chan struct{} // acts like Semaphore
}

// === Equivalent of controller method ===
func (s *Server) handler(w http.ResponseWriter, r *http.Request) {

	// tryAcquire()
	select {
	case s.sem <- struct{}{}:
		// acquired
	default:
		http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
		return
	}

	// finally { release() }
	defer func() { <-s.sem }()

	// simulate IO delay
	time.Sleep(time.Duration(s.delayMs) * time.Millisecond)

	// response
	w.WriteHeader(http.StatusOK)
}

func main() {

	// === defaults ===
	delay := 5000
	concurrency := 10000

	// === ENV overrides ===
	if v := os.Getenv("STUB_DELAY_MS"); v != "" {
		if d, err := strconv.Atoi(v); err == nil {
			delay = d
		}
	}

	if v := os.Getenv("STUB_CONCURRENCY"); v != "" {
		if c, err := strconv.Atoi(v); err == nil {
			concurrency = c
		}
	}

	// === "constructor" ===
	server := &Server{
		delayMs: delay,
		sem:     make(chan struct{}, concurrency), // Semaphore(concurrency)
	}

	http.HandleFunc("/delay", server.handler)

	fmt.Printf("Starting stub: delay=%dms, concurrency=%d\n", delay, concurrency)
	
	err := http.ListenAndServe(":8081", nil)
    if err != nil {
	   panic(err)
    }
}
