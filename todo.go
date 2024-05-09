package todo

import (
	"errors"
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
