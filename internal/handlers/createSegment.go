package handlers

import (
	"encoding/json"
	"github.com/Youngporcher/avito/internal/database"
	"net/http"
)

type CreateSegmentStruct struct {
	Name string
}

func CreateSegment(w http.ResponseWriter, r *http.Request) {
	var request CreateSegmentStruct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		HandlerResponse(w, "Error http server", 400)
		return
	}
	database.CreateSegment(request.Name)
	HandlerResponse(w, "segment "+request.Name+" created", 200)

}
