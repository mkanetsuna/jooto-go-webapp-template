package handlers

import (
	"fmt"
	"net/http"
)

func WebhookHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Webhook received")
}
