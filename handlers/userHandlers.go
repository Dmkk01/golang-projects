package handlers

import (
	"log"
	"net/http"
	"simple-api/user"
)

func usersGetAll(w http.ResponseWriter, r *http.Request) {
	users, err := user.All()

	log.Println(users)

	if err != nil {
		postError(w, http.StatusInternalServerError)
		return
	}

	postBodyResponse(w, http.StatusOK, jsonResponse{"users": users})
}
