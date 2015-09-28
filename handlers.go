package main

import (
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

func EnableCORS(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PUT, PATCH, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, api_key, Authorization")
}

func Index(w http.ResponseWriter, r *http.Request) {
    EnableCORS(w)
    fmt.Fprintln(w, "Welcome!")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
    todos := Todos{
        Todo{Name: "Write presentation"},
        Todo{Name: "Host meetup"},
    }

    EnableCORS(w)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    if err := json.NewEncoder(w).Encode(todos); err != nil {
        panic(err)
    }
}

func TodoShow(w http.ResponseWriter, r *http.Request) {
    EnableCORS(w)
    vars := mux.Vars(r)
    todoId := vars["todoId"]
    fmt.Fprintln(w, "Todo show:", todoId)
}

func Swagger(w http.ResponseWriter, r *http.Request) {
    EnableCORS(w)
    w.Header().Set("Content-Type", "application/json; charset=UTF-8")
    w.WriteHeader(http.StatusOK)
    s, err := ioutil.ReadFile("swagger.json")
    if err != nil {
        panic(err)
    }
    w.Write(s)
}
