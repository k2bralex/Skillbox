package module_30

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "skillbox/module_30/models/target"
	. "skillbox/module_30/models/user"
	"strconv"
)

var Id = 1010

var users = map[string]*User{
	"1001": {"1001", "User1", 20, Friends{}},
	"1002": {"1002", "User2", 21, Friends{}},
	"1003": {"1003", "User3", 22, Friends{}},
	"1004": {"1004", "User4", 34, Friends{}},
	"1005": {"1005", "User5", 32, Friends{}},
	"1006": {"1006", "User6", 31, Friends{}},
	"1007": {"1007", "User7", 13, Friends{}},
	"1008": {"1008", "User8", 12, Friends{}},
	"1009": {"1009", "User9", 53, Friends{}},
	"1010": {"1010", "User10", 41, Friends{}},
}

func Run() {
	r := mux.NewRouter()
	r.HandleFunc("/users", getAll).Methods("GET")
	r.HandleFunc("/users/{id}/friends", getFriendList).Methods("GET")
	r.HandleFunc("/users/{id}/friends", addFriend).Methods("PUT")
	r.HandleFunc("/users", createUser).Methods("POST")
	r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/users/{source_id}/friends/{target_id}",
		deleteFriend).Methods("DELETE")
	r.HandleFunc("/users/{id}", ageUpdate).Methods("PATCH")
	log.Fatal(http.ListenAndServe(":8080", r))
}

// getAll return all users in storage
func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// createUser add new user into storage
func createUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	Id++
	user.ID = strconv.Itoa(Id)
	users[user.ID] = &user
	json.NewEncoder(w).Encode(user)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(user.ID))
	return
}

// deleteUser from storage
func deleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
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
	return
}

// addFriend add new friend to friend list
func addFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t Target
	_ = json.NewDecoder(r.Body).Decode(&t)
	params := mux.Vars(r)
	source := users[params["id"]]
	target := users[t.ID]
	source.Friends = append(source.Friends, target)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s and %s friends now", source.Name, target.Name)))
	return
}

// deleteFriend from friend list
func deleteFriend(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	source := users[params["source_id"]]
	source.DeleteFriend(source.GetFriendIndex(params["target_id"]))

	//json.NewEncoder(w).Encode(users)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("user removed from friend list")))
	return
}

// getFriendList return all users in friend list
func getFriendList(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var friends = map[string]*User{}
	if val, ok := users[params["id"]]; ok {
		for _, friend := range val.Friends {
			friends[friend.ID] = friend
		}
	}
	json.NewEncoder(w).Encode(friends)
}

// ageUpdate change user age
func ageUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t Target
	_ = json.NewDecoder(r.Body).Decode(&t)
	params := mux.Vars(r)
	if val, ok := users[params["id"]]; ok {
		val.AgeUpdate(t.Age)
	}

	//json.NewEncoder(w).Encode(users)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("user age updated")))
	return
}
