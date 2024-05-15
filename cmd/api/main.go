package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/pinhob/todo-go"
)

func main() {
	http.HandleFunc("/", handleRoot)
	http.HandleFunc("/todos", handleTodos)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo à API de tarefas!"))
}

func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	default:
		listTodos(w, r)
	}
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	ls := &todo.List{}
	if err := ls.Load("../todo/.todo.json"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ls)
}
