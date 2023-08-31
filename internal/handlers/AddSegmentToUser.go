package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Youngporcher/avito/internal/database"
	"net/http"
)

type AddSegmentToUserStruct struct {
	AddSegments    []string
	DeleteSegments []string
	User_id        string
}

func AddSegmentToUser(w http.ResponseWriter, r *http.Request) {
	var request AddSegmentToUserStruct
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&request)
	if err != nil {
		HandlerResponse(w, "Error http server", 400)
		fmt.Println(err)
		return
	}
	database.AddSegmentToUser(request.AddSegments, request.DeleteSegments, request.User_id)
	HandlerResponse(w, "segments deleted and created", 200)
}
