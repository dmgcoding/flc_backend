package routes

import (
	"flc_backend/internal/myUtils"
	"net/http"
)

func healthzHandler(w http.ResponseWriter, r *http.Request) {
	type successMsg struct {
		Msg string `json:"msg"`
	}
	myUtils.RespondWithJson(w, http.StatusOK, successMsg{Msg: "Success"})
}
