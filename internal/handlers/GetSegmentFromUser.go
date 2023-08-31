package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Youngporcher/avito/internal/database"
	"net/http"
)

type GetSegmentFromUserStruct struct {
	User_id string
}
type SegmentsUsers struct {
	Segments []string
}

func GetSegmentFromUser(w http.ResponseWriter, r *http.Request) {
	var request GetSegmentFromUserStruct
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&request)
	if err != nil {
		HandlerResponse(w, "Error http server", 400)
		return
	}

	var result SegmentsUsers
	result.Segments = database.GetSegmentFromUser(request.User_id)
	response, err := json.Marshal(result)
	if err != nil {
		fmt.Println(err)
		return

	}
	HandlerResponse(w, string(response), 200)
}
