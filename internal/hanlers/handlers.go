package hanlers

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"os"
)

type Handler struct {
}

func MustHandler() *Handler {
	return &Handler{}
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
		fm.HandleFunc("/chats", GetUserChats).Methods("GET")

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
