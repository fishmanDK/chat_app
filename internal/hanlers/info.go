package hanlers

import "net/http"

func GetInfoByUserID(w http.ResponseWriter, r *http.Request) {
	const op = "internal.handlers.info.GetInfoByUserID"
	w.Write([]byte("/api"))
}
