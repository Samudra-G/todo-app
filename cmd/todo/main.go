package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"strings"
	"bufio"
	"github.com/Samudra-G/todo-app/internal/todo"
)

const (
	todoFile = ".todo.json"
)

func main() {
	if err := run(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run() error {
	add := flag.Bool("add", false, "Add a new task")
	complete := flag.Int("complete", 0, "Mark a task as complete by index")
	del := flag.Int("del", 0, "Delete a task by index")
	list := flag.Bool("list", false, "List all tasks")
	flag.Parse()

	todos := &todo.Todo{}
	if err := todos.Load(todoFile); err != nil {
		return err
	}

	switch {
	case *add:
		return handleAdd(todos, flag.Args())
	case *complete > 0:
		return handleComplete(todos, *complete)
	case *del > 0:
		return handleDelete(todos, *del)
	case *list:
		todos.Print()
		return nil
	default:
		return fmt.Errorf("invalid command")
	}
}

func handleAdd(todos *todo.Todo, args []string) error {
	task, err := getInput(os.Stdin, args...)
	if err != nil {
		return err
	}
	todos.Add(task)
	return todos.Store(todoFile)
}

func handleComplete(todos *todo.Todo, index int) error {
	if err := todos.Complete(index); err != nil {
		return err
	}
	return todos.Store(todoFile)
}

func handleDelete(todos *todo.Todo, index int) error {
	if err := todos.Delete(index); err != nil {
		return err
	}
	return todos.Store(todoFile)
}

func getInput(r io.Reader, args ...string) (string, error) {
	if len(args) > 0 {
		return strings.Join(args, " "), nil
	}

	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}

	text := scanner.Text()
	if len(text) == 0 {
			return "", fmt.Errorf("empty todo not allowed")
		}

	return text, nil
}