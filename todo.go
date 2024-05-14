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

type List []item

func (l *List) Add(task string) {
	newTask := item{
		Task:        task,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, newTask)
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
	list, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return os.WriteFile(fileName, list, 0666)
}

func (l *List) Load(fileName string) error {
	file, err := os.ReadFile(fileName)

	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return nil
	}

	if len(file) == 0 {
		return nil
	}

	return json.Unmarshal(file, l)
}

func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "  "
		if t.Done {
			prefix = "X "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}

	return formatted
}

func (l *List) CountPendingTodos() int {
	var total int
	for _, t := range *l {
		if !t.Done {
			total++
		}
	}

	return total
}
