package hanlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"

	authgrpc "realtime_chat_app/internal/clients/auth/grpc"
)

type Handler struct {
	client *authgrpc.Client
}

func MustHandler(client *authgrpc.Client) *Handler {
	return &Handler{
		client: client,
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
