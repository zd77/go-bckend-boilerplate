package api

import (
	"fmt"
	"net/http"
)

func RouterApi() http.Handler {
	fmt.Println("Router api")
	router := http.NewServeMux()
	router.Handle("/api/", http.StripPrefix("/api", apiRouter()))
	return router
}

func apiRouter() http.Handler {
	router := http.NewServeMux()
	router.HandleFunc("/tasks/", HandleGetTasks)
	return router
}
