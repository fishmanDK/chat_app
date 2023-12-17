package main

import (
	"context"
	"net/http"
	authgrpc "realtime_chat_app/internal/clients/auth/grpc"
	"realtime_chat_app/internal/hanlers"
)

func main() {
	client, _ := authgrpc.NewClient(context.Background(), "localhost:5001")

	handler := hanlers.MustHandler(client)
	router, _ := handler.InitRoutes()

	//http.Handle("/", router)
	http.ListenAndServe(":8000", router)
}
