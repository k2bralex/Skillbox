package module_30

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"skillbox/module_30/models/user"
	"skillbox/module_30/pkg/logging"
)

func Start() {
	logger := logging.GetLogger()
	log.Println("create router")
	r := mux.NewRouter()

	log.Println("register user handler")
	handler := user.NewHandler(logger)
	handler.Register(r)

	log.Fatal(http.ListenAndServe(":8080", r))
}
