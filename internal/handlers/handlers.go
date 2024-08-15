package handlers

import (
	"fmt"
	"net/http"
	"time"
)

func ScheduledTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Scheduled task executed at %s", time.Now().Format(time.RFC3339))
}

func Webhook1Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Webhook 1 received")
}

func Webhook2Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Webhook 2 received")
}

func DevHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the development endpoint!")
}