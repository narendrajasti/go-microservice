package handlers

import (
	"fmt"
	"log"
	"net/http"
)

// HealthCheck is a simple handler
type HealthCheck struct {
	l *log.Logger
}

// NewHealthCheck creates a new healthCheck handler with the given logger
func NewHealthCheck(l *log.Logger) *HealthCheck {
	return &HealthCheck{l}
}

// HealthCheck returns UP
func (hc *HealthCheck) HealthCheck(rw http.ResponseWriter, _ *http.Request) {
	hc.l.Println("UP!!!")
	fmt.Fprintf(rw, "UP! \n")
}
