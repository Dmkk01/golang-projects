package handlers

import (
	"net/http"
	"strings"

	"gopkg.in/mgo.v2/bson"
)

func UsersRouter(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimSuffix(r.URL.Path, "/")
	if path == "/users" {
		switch r.Method {
		case http.MethodGet:
			usersGetAll(w, r)
			return
		case http.MethodPost:
			// createUser(w, r)
			return
		default:
			postError(w, http.StatusMethodNotAllowed)
			return
		}
	}

	path = strings.TrimPrefix(path, "/users/")

	if !bson.IsObjectIdHex(path) {
		postError(w, http.StatusNotFound)
		return
	}

	// id := bson.ObjectIdHex(path)

	switch r.Method {
	case http.MethodGet:
		// getUser(id, w, r)
		return
	case http.MethodPut:
		// updateUser(id, w, r)
		return
	case http.MethodPatch:
		// updateUser(id, w, r)
		return
	case http.MethodDelete:
		// deleteUser(id, w, r)
		return
	default:
		postError(w, http.StatusMethodNotAllowed)
	}
}
