package main_test

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binName  = "todo"
	fileName = ".todo.json"
)

func TestMain(m *testing.M) {
	fmt.Println("Building tool...")

	if os.Getenv("TODO_FILENAME") != "" {
		fileName = os.Getenv("TODO_FILENAME")
	}

	if runtime.GOOS == "windows" {
		binName += ".exe"
	}

	build := exec.Command("go", "build", "-o", binName)

	if err := build.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "Error building tool %s: %s", binName, err)
		os.Exit(1)
	}

	fmt.Println("Running tests...")
	result := m.Run()

	fmt.Println("Cleaning up...")
	os.Remove(binName)
	os.Remove(fileName)

	os.Exit(result)
}

func TestTodoCLI(t *testing.T) {
	task := "task number 1"
	secondTask := "task number 2"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("error to find working directory: %s", err)
	}

	cmdPath := filepath.Join(dir, binName)

	t.Run("add new task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Fatalf("error adding new task: %s", err)
		}
	})

	t.Run("add new task from stdin", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add")

		cmdStdIn, err := cmd.StdinPipe()
		if err != nil {
			t.Fatalf("error calling StdinPipe: %s", err)
		}

		io.WriteString(cmdStdIn, secondTask)
		cmdStdIn.Close()

		if err := cmd.Run(); err != nil {
			t.Fatalf("error add new task from stdin: %s", err)
		}
	})

	t.Run("list all tasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error running list command: %s", err)
		}

		outHasTaskOne := strings.Contains(string(out), task)
		outHasTaskTwo := strings.Contains(string(out), secondTask)

		if !outHasTaskOne {
			t.Errorf("task '%s' should be in output, but we got %v", task, outHasTaskOne)
		}

		if !outHasTaskTwo {
			t.Errorf("task '%s' should be in output, but we got %v", secondTask, outHasTaskTwo)
		}
	})

	t.Run("update first task", func(t *testing.T) {
		updatedTask := "updated task"

		cmd := exec.Command(cmdPath, "update", "-id", "1", "-task", updatedTask)

		if err := cmd.Run(); err != nil {
			t.Fatalf("error adding new task: %s", err)
		}

		cmdList := exec.Command(cmdPath, "-list")
		out, err := cmdList.CombinedOutput()
		if err != nil {
			t.Fatalf("error updating task: %s", err)
		}

		outHasUpdatedTask := strings.Contains(string(out), updatedTask)

		if !outHasUpdatedTask {
			t.Errorf("task '%s' should be in output, but we got %v", updatedTask, outHasUpdatedTask)
		}
	})

	t.Run("complete first task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-complete", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error completing first task: %s", err)
		}

		if len(out) != 0 {
			t.Fatalf("Expected no return, got %s", out)
		}
	})

	t.Run("delete first task", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-del", "1")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatalf("error deleting first task: %s", err)
		}

		if len(out) != 0 {
			t.Fatalf("Expected no return, got %s", out)
		}

		cmdList := exec.Command(cmdPath, "-list")
		outList, err := cmdList.CombinedOutput()
		if err != nil {
			t.Fatalf("error listing tasks: %s", err)
		}

		if len(outList) == 1 {
			t.Errorf("expected list with only second item, got: %s", outList)
		}
	})
}
