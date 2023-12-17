package main

import (
	"net/http"
	"realtime_chat_app/internal/hanlers"
)

func main() {
	handler := hanlers.MustHandler()
	router, _ := handler.InitRoutes()

	//http.Handle("/", router)
	http.ListenAndServe(":8000", router)
}
