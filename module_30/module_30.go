package module_30

import (
	"encoding/json"
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
	//serverRun()

	r := mux.NewRouter()
	r.HandleFunc("/persons", getAll).Methods("GET")
	r.HandleFunc("/friends/user_id", getFriends).Methods("GET")
	//r.HandleFunc("/books/{id}", getBook).Methods("GET")
	r.HandleFunc("/create", createPerson).Methods("POST")
	//r.HandleFunc("/books/{id}", updateBook).Methods("PUT")
	r.HandleFunc("/user", deletePerson).Methods("DELETE")
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
		for _, person := range persons {
			if index, isOk := person.FriendContains(val.ID); isOk {
				name = person.FriendRemove(index)
			}
		}
		delete(persons, t.TargetId)
	}
	json.NewEncoder(w).Encode(persons)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(name))
	return
}

/*4. Сделайте обработчик, который возвращает всех друзей пользователя.*/
func getFriends(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for key, val := range persons {
		if key == params["user_id"] {
			err := json.NewEncoder(w).Encode(val.Friends)
			if err != nil {
				return
			}
			return
		}
	}
	json.NewEncoder(w).Encode(&Friends{})
}

/*
5. Сделайте обработчик, который обновляет возраст пользователя.
*/

func ageUpdate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	if val, ok := persons[params["id"]]; ok {
		val.AgeUpdate(3)
	}
}
