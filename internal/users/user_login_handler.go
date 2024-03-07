package users

import (
	"encoding/json"
	"flc_backend/internal/myUtils"
	"net/http"
)

var UsersSlice = []User{}

func UserLoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		myUtils.RespondWithError(w, http.StatusBadRequest, "please enter credentials")
		return
	}

	var userInput User

	err := json.NewDecoder(r.Body).Decode(&userInput)
	if err != nil {
		myUtils.RespondWithError(w, http.StatusBadRequest, "error parsing body")
		return
	}

	exists := userExists(userInput.Email)
	if !exists {
		myUtils.RespondWithError(w, http.StatusBadRequest, "no user with that email")
		return
	}

	credentialsMatch := userCredentialsMatch(userInput.Email, userInput.Password)
	if !credentialsMatch {
		myUtils.RespondWithError(w, http.StatusBadRequest, "credentials doesn't match")
		return
	}

	userIndex := getIndexOfUserFromEmail(userInput.Email)
	if userIndex < 0 {
		myUtils.RespondWithError(w, http.StatusInternalServerError, "couldn't find user with the credentials")
		return
	}

	myUtils.RespondWithJson(w, http.StatusOK, UsersSlice[userIndex])
}
