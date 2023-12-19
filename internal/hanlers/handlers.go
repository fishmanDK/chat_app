package hanlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"log/slog"
	"net/http"
	"os"
	"realtime_chat_app/internal/service"

	authgrpc "realtime_chat_app/internal/clients/auth/grpc"
)

type Handler struct {
	client  *authgrpc.Client
	service *service.Service
	logger  *slog.Logger
}

func MustHandler(client *authgrpc.Client, service *service.Service, logger *slog.Logger) *Handler {
	return &Handler{
		client:  client,
		service: service,
		logger:  logger,
	}
}

func (h *Handler) InitRoutes() (*mux.Router, error) {
	router := mux.NewRouter()
	const dir = "internal/static/html"
	if _, err := os.Stat(dir); err != nil {
		fmt.Println(err.Error())
	}

	{
		fm := router.PathPrefix("/fm").Subrouter()
		fm.PathPrefix("/files/").Handler(http.StripPrefix("/fm/files/", http.FileServer(http.Dir(dir))))

		fm.HandleFunc("/chats", h.GetUserChats).Methods("GET")

		id := fm.PathPrefix("/chats/{id}").Subrouter()
		{
			id.HandleFunc("", GetChatId).Methods("GET")
			id.HandleFunc("", DeleteChat).Methods("DELETE")
			//id.HandleFunc("", DeleteChat).Methods("PUTCH")
			id.HandleFunc("", CreateChat).Methods("POST")
		}

		info := fm.PathPrefix("/info").Subrouter()
		{
			info.HandleFunc("/{userID}", GetInfoByUserID)
		}
	}

	return router, nil
}
