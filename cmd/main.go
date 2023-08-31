package main

import (
	"github.com/Youngporcher/avito/internal/handlers"
	"github.com/Youngporcher/avito/internal/server"
	"log"
	"net/http"
)

func main() {
	http.Handle("/CreateSegment", http.HandlerFunc(handlers.CreateSegment))
	http.Handle("/DeleteSegment", http.HandlerFunc(handlers.DeleteSegment))
	http.Handle("/AddSegmentToUser", http.HandlerFunc(handlers.AddSegmentToUser))
	http.Handle("/GetSegmentFromUser", http.HandlerFunc(handlers.GetSegmentFromUser))
	// Запуск сервера
	srv := new(server.Server)
	err := srv.Run("8080")
	if err != nil {
		log.Fatalf("error running http server: %s", err.Error())
	}
}
