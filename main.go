package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/narendrajasti/go-microservice/data"
	"github.com/narendrajasti/go-microservice/handlers"
)

func main() {

	l := log.New(os.Stdout, "go-microservice", log.LstdFlags)
	v := data.NewValidation()

	// greet := handlers.NewGreet(l)
	// healthCheck := handlers.NewHealthCheck(l)
	products := handlers.NewProducts(l, v)
	healthCheck := handlers.NewHealthCheck(l)

	sm := mux.NewRouter()
	getRouter := sm.Methods("GET").Subrouter()
	getRouter.HandleFunc("/products", products.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", products.ListSingle)

	postRouter := sm.Methods("POST").Subrouter()
	postRouter.HandleFunc("/products", products.AddProduct)
	postRouter.Use(products.MiddlewareValidateProduct)

	putRouter := sm.Methods("PUT").Subrouter()
	putRouter.HandleFunc("/products/{id:[0-9]+}", products.UpdateProducts)
	putRouter.Use(products.MiddlewareValidateProduct)

	healthRouter := sm.Methods("GET").Subrouter()
	healthRouter.HandleFunc("/healthCheck", healthCheck.HealthCheck)

	optns := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(optns, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer((http.Dir("./"))))

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
