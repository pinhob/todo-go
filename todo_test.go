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

func TestDelete(t *testing.T) {
	ls := todo.List{}
	task := "task 2"

	ls.Add(task)
	ls.Delete(1)

	if len(ls) > 0 {
		t.Errorf("list should be empty, but got %v", len(ls))
		t.Errorf("%v", ls[0])
	}
}
