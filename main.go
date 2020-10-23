package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/narendrajasti/go-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-microservice", log.LstdFlags)

	greet := handlers.NewGreet(l)
	healthCheck := handlers.NewHealthCheck(l)
	products := handlers.NewProduct(l)

	sm := http.NewServeMux()
	sm.Handle("/", greet)
	sm.Handle("/healthCheck", healthCheck)
	sm.Handle("/products", products)
	s := &http.Server{
		Addr:         ":9090",
		Handler:      sm,
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		err := s.ListenAndServe()

		if err != nil {
			l.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	l.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
