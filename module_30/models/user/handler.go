package user

import (
	"github.com/gorilla/mux"
	"net/http"
	. "skillbox/module_30/hendlers"
)

var _ Handler = &handler{}

const (
	usersURL = "/users"
	userURL  = "/users/{id}"
)

type handler struct {
}

func NewHandler() Handler {
	return &handler{}
}

func (h *handler) Register(r *mux.Router) {
	r.HandleFunc(usersURL, GetList).Methods("GET")
	r.HandleFunc(usersURL, CreateUser).Methods("POST")
	r.HandleFunc(userURL, AgeUpdate).Methods("PATCH")
	r.HandleFunc(userURL, DeleteUser).Methods("DELETE")
}

func GetList(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("this is list of users"))
	/*w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)*/
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	Id++
	user.ID = strconv.Itoa(Id)
	users[user.ID] = &user
	json.NewEncoder(w).Encode(user)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(user.ID))
	return*/
	w.Write([]byte("this is create users"))
}

func AgeUpdate(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	var t Target
	_ = json.NewDecoder(r.Body).Decode(&t)
	params := mux.Vars(r)
	if val, ok := users[params["id"]]; ok {
		val.AgeUpdate(t.Age)
	}

	//json.NewEncoder(w).Encode(users)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user age updated")))
	return*/
	w.Write([]byte("this is update users"))
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	var (
		name string
		t    Target
	)
	_ = json.NewDecoder(r.Body).Decode(&t)
	if val, ok := users[t.ID]; ok {
		name = val.Name
		for _, user := range users {
			if index, isOk := user.FriendContains(val.ID); isOk {
				user.DeleteFriend(index)
			}
		}
		delete(users, t.ID)
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name))
	return*/

	w.Write([]byte("this is delete users"))
}
