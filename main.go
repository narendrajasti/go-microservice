package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/narendrajasti/go-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-microservice", log.LstdFlags)

	// greet := handlers.NewGreet(l)
	// healthCheck := handlers.NewHealthCheck(l)
	products := handlers.NewProduct(l)
	healthCheck := handlers.NewHealthCheck(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", products.GetProducts)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", products.AddProduct)
	postRouter.Use(products.MiddlewareValidationProduct)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", products.UpdateProducts)
	putRouter.Use(products.MiddlewareValidationProduct)

	healthRouter := sm.Methods("GET").Subrouter()
	healthRouter.HandleFunc("/healthCheck", healthCheck.HealthCheck)

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
