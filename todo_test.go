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

func TestUpdate(t *testing.T) {
	ls := todo.List{}
	description := "task to update"
	updatedDescription := "updated task"

	ls.Add(description)
	task, err := ls.Update(1, updatedDescription)

	if err != nil {
		t.Errorf("task cannot be updated, %v", err)
	}

	if task.Task != updatedDescription {
		t.Errorf("task should be equal %v, but got %v", updatedDescription, task.Task)
	}
}

func TestSaveLoad(t *testing.T) {
	ls := todo.List{}
	task := "task 1"
	ls.Add(task)

	if err := ls.Save(fileName); err != nil {
		t.Errorf("Got error '%s' when saving file", err)
	}

	err := ls.Load(fileName)

	if err != nil {
		t.Errorf("Got error '%s' when loading file", err)
	}

	got := ls[0].Task

	if got != task {
		t.Errorf("want task %s, got %s", task, got)
	}
}
