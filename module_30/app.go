package module_30

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"skillbox/module_30/models/user"
)

func Start() {
	r := mux.NewRouter()

	handler := user.NewHandler()
	handler.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
