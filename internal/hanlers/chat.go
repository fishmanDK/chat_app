package hanlers

import "net/http"

func GetMainPage(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.GetMainPage"

	http.ServeFile(w, r, "index.html")
	//w.Write([]byte("/api"))
}

func GetUserChats(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.chat.GetUserChats"
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
