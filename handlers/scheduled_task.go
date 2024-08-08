package handlers

import (
	"fmt"
	"net/http"
)

func ScheduledTaskHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Scheduled task executed")
}
