package todo

import "time"

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

func (l *List) Complete(taskNumber int) {
	ls := *l
	ls[taskNumber-1].Done = true
}
