package notifications

import (
	"flc_backend/internal/myUtils"
	"flc_backend/internal/users"
	"net/http"
)

func GetNotificationsHandler(w http.ResponseWriter, r *http.Request) {
	authHeaders := r.Header["Authorization"]
	if len(authHeaders) == 0 {
		myUtils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	authorizationHeader := authHeaders[0]

	tokenValid := users.IsTokenMatch(authorizationHeader)
	if !tokenValid {
		myUtils.RespondWithError(w, http.StatusUnauthorized, "Invalid token")
		return
	}

	myUtils.RespondWithJson(w, http.StatusOK, notificationsSlice)
}
