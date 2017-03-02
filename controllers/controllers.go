package controllers

import (
	"encoding/json"
	"fmt"
	"models"
	"net/http"

	"github.com/gorilla/mux"
)

type MafiaAPI struct{}

func (m MafiaAPI) index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome!")
}

func (m MafiaAPI) todoIndex(w http.ResponseWriter, r *http.Request) {
	todos := models.Todos{
		models.Todo{Name: "Write presentation"},
		models.Todo{Name: "Host meetup"},
	}

	if err := json.NewEncoder(w).Encode(todos); err != nil {
		panic(err)
	}
}

func (m MafiaAPI) todoShow(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	todoId := vars["todoId"]
	fmt.Fprintln(w, "Todo show:", todoId)
}
