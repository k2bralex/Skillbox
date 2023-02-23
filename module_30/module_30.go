package module_30

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	. "skillbox/module_30/models/person"
	. "skillbox/module_30/models/target"
	"strconv"
)

var Id = 1010

var persons = map[string]*Person{
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
	r.HandleFunc("/persons", getAll).Methods("GET")
	r.HandleFunc("/friends/{user_id}", getFriends).Methods("GET")
	r.HandleFunc("/make_friends", makeFriends).Methods("POST")
	r.HandleFunc("/create", createPerson).Methods("POST")
	r.HandleFunc("/user", deletePerson).Methods("DELETE")
	r.HandleFunc("/{user_id}", ageUpdate).Methods("PUT")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func getAll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(persons)
}

/*
 1. Сделайте обработчик создания пользователя. У пользователя должны быть следующие поля:
    имя, возраст и массив друзей. Пользователя необходимо сохранять в мапу.
*/
func createPerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var person Person
	_ = json.NewDecoder(r.Body).Decode(&person)
	Id++
	person.ID = strconv.Itoa(Id)
	persons[person.ID] = &person
	json.NewEncoder(w).Encode(person)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(person.ID))
	return
}

/*
2. Сделайте обработчик, который делает друзей из двух пользователей. Например, если мы создали
двух пользователей и нам вернулись их ID, то в запросе мы можем указать ID пользователя,
который инициировал запрос на дружбу, и ID пользователя, который примет инициатора в друзья.
*/

func makeFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t Target
	_ = json.NewDecoder(r.Body).Decode(&t)
	source := persons[t.SourceId]
	target := persons[t.TargetId]
	source.Friends = append(source.Friends, target)

	//json.NewEncoder(w).Encode(persons)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s and %s friends now", source.Name, target.Name)))
	return
}

/*
 3. Сделайте обработчик, который удаляет пользователя. Данный обработчик принимает ID пользователя
и удаляет его из хранилища, а также стирает его из массива friends у всех его друзей.
*/

func deletePerson(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var (
		name string
		t    Target
	)
	_ = json.NewDecoder(r.Body).Decode(&t)
	if val, ok := persons[t.TargetId]; ok {
		name = val.Name
		for _, person := range persons {
			if index, isOk := person.FriendContains(val.ID); isOk {
				person.FriendRemove(index)
			}
		}
		delete(persons, t.TargetId)
	}

	//json.NewEncoder(w).Encode(persons)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name))
	return
}

/*4. Сделайте обработчик, который возвращает всех друзей пользователя.*/
func getFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	var friends = map[string]*Person{}
	if val, ok := persons[params["user_id"]]; ok {
		for _, friend := range val.Friends {
			friends[friend.ID] = friend
		}
	}
	json.NewEncoder(w).Encode(friends)
}

/*
5. Сделайте обработчик, который обновляет возраст пользователя.
*/

func ageUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var t Target
	_ = json.NewDecoder(r.Body).Decode(&t)
	newAge, _ := strconv.Atoi(t.NewAge)
	params := mux.Vars(r)
	if val, ok := persons[params["user_id"]]; ok {
		val.AgeUpdate(newAge)
	}

	//json.NewEncoder(w).Encode(persons)

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(fmt.Sprintf("person age updated")))
	return

}
