package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/pinhob/todo-go"
)

const todoFileName = "../todo/.todo.json"

type task struct {
	Description string `json:"task"`
}

func main() {
	http.HandleFunc("GET /todos", listTodos)
	http.HandleFunc("POST /todos", addTodo)
	http.HandleFunc("DELETE /todos/{id}", handleDeleteTodo)
	http.HandleFunc("PUT /todos/{id}", handleUpdateTodo)
	http.HandleFunc("/", handleRoot)
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Bem-vindo à API de tarefas!"))
}

func listTodos(w http.ResponseWriter, r *http.Request) {
	ls := &todo.List{}
	if err := ls.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(ls)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
	var task task
	ls := &todo.List{}
	if err := ls.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	ls.Add(task.Description)
	if err := ls.Save(todoFileName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)
}

func handleDeleteTodo(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	ls := &todo.List{}
	if err := ls.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	taskNum, err := strconv.Atoi(id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	if err := ls.Delete(taskNum); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	if err := ls.Save(todoFileName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Item successfully deleted")
}

func handleUpdateTodo(w http.ResponseWriter, r *http.Request) {
	var task task
	ls := &todo.List{}

	id, idErr := strconv.Atoi(r.PathValue("id"))
	if idErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(idErr)
	}

	if err := ls.Load(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err := json.NewDecoder(r.Body).Decode(&task)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	updatedTask, updateTaskErr := ls.Update(id, task.Description)
	if updateTaskErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(updateTaskErr)
	}

	if err := ls.Save(todoFileName); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedTask)
}
