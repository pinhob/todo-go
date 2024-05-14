package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/alexeyco/simpletable"
	"github.com/pinhob/todo-go"
)

func main() {
	add := flag.Bool("add", false, "Add a new task to your list")
	list := flag.Bool("list", false, "List all all tasks from your list")
	complete := flag.Int("complete", 0, "Mark one task as completed")
	del := flag.Int("del", 0, "Delete specified task from your list")

	flag.Parse()

	ls := &todo.List{}
	if err := ls.Load(".todo.json"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	table := simpletable.New()
	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignCenter, Text: "CreatedAt"},
			{Align: simpletable.AlignCenter, Text: "CompletedAt"},
		},
	}

	table.Footer = &simpletable.Footer{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Span: 5, Text: red(fmt.Sprintf("You have %d pending todos", ls.CountPendingTodos()))},
		},
	}

	for k, task := range *ls {
		taskDescription := blue(task.Task)

		if task.Done {
			taskTextWithCheckMark := fmt.Sprintf("\u2705 %s", task.Task)
			taskDescription = green(taskTextWithCheckMark)
		}

		r := []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: fmt.Sprintf("%d", k+1)},
			{Align: simpletable.AlignLeft, Text: taskDescription},
			{Align: simpletable.AlignCenter, Text: strconv.FormatBool(task.Done)},
			{Align: simpletable.AlignCenter, Text: task.CreatedAt.Format(time.RFC822)},
			{Align: simpletable.AlignCenter, Text: task.CompletedAt.Format(time.RFC822)},
		}

		table.Body.Cells = append(table.Body.Cells, r)
	}

	table.SetStyle(simpletable.StyleUnicode)

	switch {
	case *list:
		fmt.Println(table.String())

	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintf(os.Stderr, "error adding new task, %v\n", err)
			os.Exit(1)
		}

		ls.Add(task)
		if err := ls.Save(".todo.json"); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving list, %v\n", err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := ls.Complete(*complete); err != nil {
			fmt.Fprintf(os.Stderr, "Error completing the task, %v\n", err)
		}

		if err := ls.Save(".todo.json"); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving list, %v\n", err)
			os.Exit(1)
		}
	case *del > 0:
		if err := ls.Delete(*del); err != nil {
			fmt.Fprintf(os.Stderr, "Error deleting the task, %v\n", err)
		}

		if err := ls.Save(".todo.json"); err != nil {
			fmt.Fprintf(os.Stderr, "Error saving list, %v\n", err)
			os.Exit(1)
		}
	}
}

func getTask(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()

	if err := scanner.Err(); err != nil {
		return "", nil
	}

	task := scanner.Text()
	if len(task) == 0 {
		return "", errors.New("task must be described to be added")
	}

	return task, nil
}
