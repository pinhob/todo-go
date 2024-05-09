package todo_test

import (
	"testing"

	"github.com/pinhob/todo-go"
)

const fileName = ".todo.json"

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
	err := ls.Complete(1)

	if err != nil {
		t.Errorf("Got error '%v' when using Complete funcion", err)
	}

	got := ls[0].Done

	if !got {
		t.Errorf("got %t but expected true", got)
	}
}

func TestDelete(t *testing.T) {
	ls := todo.List{}
	task := "task 2"

	ls.Add(task)
	err := ls.Delete(1)

	if err != nil {
		t.Errorf("Got error '%v' when using Complete funcion", err)
	}

	if len(ls) > 0 {
		t.Errorf("list should be empty, but got %v", len(ls))
		t.Errorf("%v", ls[0])
	}
}

func TestSave(t *testing.T) {
	ls := todo.List{}
	task := "task 1"
	ls.Add(task)

	saveErr := ls.Save(fileName)
	if saveErr != nil {
		t.Errorf("Got error '%s' when saving file", saveErr)
	}

	list, loadErr := ls.Load(fileName)

	if loadErr != nil {
		t.Errorf("Got error '%s' when loading file", loadErr)
	}

	got := list[0].Task

	if got != task {
		t.Errorf("want task %s, got %s", task, got)
	}
}
