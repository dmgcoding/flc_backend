package users

import (
	"encoding/json"
	"flc_backend/internal/myUtils"
	"net/http"
)

const SampleJwtToken = "32|9cSNrbjAeq6lwD4Uefi5k47nRHksQb7XhsciId1Uf65cfa89"
const SampleTokenHeader = "Bearer 32|9cSNrbjAeq6lwD4Uefi5k47nRHksQb7XhsciId1Uf65cfa89"

func CreateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Body == nil {
		myUtils.RespondWithError(w, http.StatusBadRequest, "body can't be empty")
		return
	}

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		myUtils.RespondWithError(w, http.StatusBadRequest, "error parsing body")
		return
	}

	//check for existing user with that email
	exists := userExists(newUser.Email)
	if exists {
		myUtils.RespondWithError(w, http.StatusBadRequest, "user with that email already exists. please login")
		return
	}

	newUser.AccessToken = SampleJwtToken

	UsersSlice = append(UsersSlice, newUser)

	type successMsg struct {
		Msg string `json:"msg"`
	}

	myUtils.RespondWithJson(w, http.StatusCreated, successMsg{Msg: "User created successfully!"})
}
