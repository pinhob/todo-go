package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pinhob/todo-go"
)

func main() {
	add := flag.Bool("add", false, "Add a new task to your list")
	list := flag.Bool("list", false, "List all all tasks from your list")

	flag.Parse()

	ls := &todo.List{}
	if err := ls.Load(".todo.json"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *list:
		fmt.Print(ls)

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
