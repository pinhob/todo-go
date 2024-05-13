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

	flag.Parse()

	list := &todo.List{}
	if err := list.Load(".todo.json"); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case *add:
		task, err := getTask(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Errorf("Error adding new task, %v", err)
		}

		list.Add(task)
		if err := list.Save(".todo.json"); err != nil {
			fmt.Errorf("Error saving list, %v", err)
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
