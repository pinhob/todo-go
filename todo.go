package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []*item

func (l *List) Add(task string) {
	newTask := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, &newTask)
}

func (l *List) Complete(taskNumber int) error {
	ls := *l

	if taskNumber < 1 || taskNumber > (len(ls)) {
		return errors.New("invalid task number")
	}

	ls[taskNumber-1].Done = true

	return nil
}

func (l *List) Delete(taskNumber int) error {
	ls := *l

	if taskNumber < 1 || taskNumber > (len(ls)) {
		return errors.New("invalid task number")
	}

	*l = append(ls[:taskNumber-1], ls[taskNumber:]...)

	return nil
}

func (l *List) Save(fileName string) error {
	/*
		save list to a file with bufio
		return error if needed
	*/
	list, err := json.Marshal(*l)

	if err != nil {
		return err
	}

	fileErr := os.WriteFile(".todo.json", list, 0666)

	if fileErr != nil {
		fmt.Println("error in write file")
		return fileErr
	}

	return nil
}

func (l *List) Load(fileName string) (List, error) {
	list := List{}
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(file, &list); err != nil {
		return nil, errors.New("error unmarshal json")
	}

	return list, nil
}
