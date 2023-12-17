package hanlers

import (
	"context"
	"encoding/json"
	"net/http"
	"realtime_chat_app"
	"realtime_chat_app/internal/transform"
)

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.GetMainPage"

	http.ServeFile(w, r, "index.html")
	//w.Write([]byte("/api"))
}

func (h *Handler) GetUserChats(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.GetUserChats"
	var user realtime_chat_app.User

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		return
	}

	h.client.Auth(context.Background(), transform.UserTransform(user))

	w.Write([]byte("/api/chats"))
}

func GetChatId(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.GetChatId"
	w.Write([]byte("/api/chat/id"))
}

func DeleteChat(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.DeleteChat"
	w.Write([]byte("/api/chat/id"))
}

func CreateChat(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.CreateChat"
	w.Write([]byte("/api/chat/id/1"))
}
