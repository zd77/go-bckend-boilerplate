package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

type GetTasksResponse struct {
	Message []string `json:"message"`
}

func HandleGetTasks(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	path := r.URL.Path
	parts := strings.Split(path, "/")
	fmt.Printf("parts: %+v\n", parts)

	voughtHeroes := []string{"Homelander", "A-Train", "The Deep"}
	response := GetTasksResponse{
		Message: voughtHeroes,
	}
	jsonResp, err := json.Marshal(response)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResp)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

}
