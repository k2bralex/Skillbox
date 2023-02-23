package user

import (
	"github.com/gorilla/mux"
	"net/http"
	. "skillbox/module_30/hendlers"
	"skillbox/module_30/pkg/logging"
)

var _ Handler = &handler{}

const (
	usersURL   = "/users"
	userURL    = "/users/{id}"
	friendsURL = "/users/{id}/friends"
	friendURL  = "/users/{source_id}/friends/{target_id}"
)

type handler struct {
	logger logging.Logger
}

func NewHandler(l logging.Logger) Handler {
	return &handler{logger: l}
}

func (h *handler) Register(r *mux.Router) {
	r.HandleFunc(usersURL, GetList).Methods("GET")
	r.HandleFunc(usersURL, CreateUser).Methods("POST")
	r.HandleFunc(userURL, AgeUpdate).Methods("PATCH")
	r.HandleFunc(userURL, DeleteUser).Methods("DELETE")

	r.HandleFunc(friendsURL, GetFriendList).Methods("GET")
	r.HandleFunc(friendsURL, AddFriend).Methods("PUT")
	r.HandleFunc(friendURL, DeleteFriend).Methods("DELETE")
}

func GetList(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)*/
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is list of users"))
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
	w.WriteHeader(http.StatusCreated)
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
	w.WriteHeader(http.StatusOK)
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

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("this is delete users"))
}

func GetFriendList(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var friends = map[string]*User{}
	if val, ok := users[params["id"]]; ok {
		for _, friend := range val.Friends {
			friends[friend.ID] = friend
		}
	}
	json.NewEncoder(w).Encode(friends)*/

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("this is all friends"))
}

func AddFriend(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	var t Target
	_ = json.NewDecoder(r.Body).Decode(&t)
	params := mux.Vars(r)
	source := users[params["id"]]
	target := users[t.ID]
	source.Friends = append(source.Friends, target)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s and %s friends now", source.Name, target.Name)))
	return*/

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("added friend to user"))
}

func DeleteFriend(w http.ResponseWriter, r *http.Request) {
	/*w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	source := users[params["source_id"]]
	source.DeleteFriend(source.GetFriendIndex(params["target_id"]))

	//json.NewEncoder(w).Encode(users)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("user removed from friend list")))
	return*/

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("user friend deleted"))
}
