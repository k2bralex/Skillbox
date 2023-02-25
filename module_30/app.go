package module_30

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"skillbox/module_30/models/user"
	. "skillbox/module_30/models/user"
)

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

func Start() {
	fmt.Println("create router")
	r := mux.NewRouter()

	storage := user.NewWorkStorage()

	for key, val := range users {
		storage.Add(val, key)
	}

	service := NewService(storage)

	fmt.Println("register user handler")
	handler := user.NewHandler(service)
	handler.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
