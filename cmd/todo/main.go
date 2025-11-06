package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/ukiran03/todo"
)

var todoFileName = ".todo.json"

func main() {
	flag.Usage = func() {
		fmt.Fprintln(flag.CommandLine.Output(), "Usage information:")
		flag.PrintDefaults()
	}

	file := flag.String("file", "", "Write the todo items to File")
	add := flag.Bool("add", false, "Add task to the ToDo list via Args or STDIN\nmutiple lines for multiple tasks")
	list := flag.Bool("list", false, "List all tasks")
	complete := flag.Int("complete", 0, "Item to be completed")
	delete := flag.Int("delete", 0, "Item to be deleted")
	verbose := flag.Bool("verbose", false, "Show verbose List")
	pending := flag.Bool("pending", false, "Show Incomplete tasks")

	flag.Parse()

	// check if the user defined the ENV VAR for a custom filename
	if os.Getenv("TODO_FILENAME") != "" {
		todoFileName = os.Getenv("TODO_FILENAME")
	}
	if *file != "" {
		todoFileName = *file
		fmt.Println(todoFileName)
	}

	l := &todo.List{}
	// initial read to do items from file
	if err := l.Get(todoFileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case len(os.Args) == 1 || *list:
		fmt.Print(l)
	case *delete > 0:
		if err := l.Delete(*delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *complete > 0:
		if err := l.Complete(*complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *add:
		tasks, err := getTasks(os.Stdin, flag.Args()...)
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, t := range tasks {
			l.Add(t)
		}
		if err := l.Save(todoFileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case *pending:
		var formated string
		for k, t := range *l {
			if !t.Done {
				formated += t.FormatTask(k, true, true)
			}
		}
		fmt.Print(formated)
	case *verbose:
		var formated string
		for k, t := range *l {
			formated += t.FormatTask(k, true, true)
		}
		fmt.Print(formated)
	default:
		fmt.Fprintln(os.Stderr, "Invalid option")
		os.Exit(1)
	}
}

// getTask function decides where to get the description for a new
// task from: Args or STDIN
func getTasks(r io.Reader, args ...string) ([]string, error) {
	var tasks []string
	// via ARGs
	if len(args) > 0 {
		task := strings.Join(args, " ")
		tasks = append(tasks, task)
		return tasks, nil
	}
	// via STDIN
	s := bufio.NewScanner(r)
	for s.Scan() {
		if len(s.Text()) == 0 {
			return nil, fmt.Errorf("Task cannot be blank")
		}
		tasks = append(tasks, s.Text())
	}

	if err := s.Err(); err != nil {
		return nil, err
	}
	return tasks, nil
}
