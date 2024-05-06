package todo_test

import (
	"testing"

	"github.com/pinhob/todo-go"
)

func TestAdd(t *testing.T) {
	ls := todo.List{}
	task := "task 1"

	ls.Add(task)
	got := ls[0].Task

	if got != task {
		t.Errorf("got %s want %s", got, task)
	}
}

func TestComplete(t *testing.T) {
	ls := todo.List{}
	task := "task 1"

	ls.Add(task)
	ls.Complete(1)

	got := ls[0].Done

	if !got {
		t.Errorf("got %t but expected true", got)
	}
}
