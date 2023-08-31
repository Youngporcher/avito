package handlers

import (
	"encoding/json"
	"github.com/Youngporcher/avito/internal/database"
	"net/http"
)

type DeleteSegmentStruct struct {
	Name string
}

func DeleteSegment(w http.ResponseWriter, r *http.Request) {
	var request DeleteSegmentStruct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		HandlerResponse(w, "Error http server", 400)
		return
	}
	database.DeleteSegment(request.Name)
	HandlerResponse(w, "segment "+request.Name+" deleted", 200)
}
