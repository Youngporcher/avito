package handlers

import "net/http"

func HandlerResponse(w http.ResponseWriter, text string, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write([]byte(text))
}
