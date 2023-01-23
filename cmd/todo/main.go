package main

import (
	"errors"
	"flag"
	"fmt"
	"os"

	"github.com/rnemeth90/todo"
)

var (
	task     string
	complete int
	delete   int
	list     bool
)

func init() {
	flag.StringVar(&task, "task", "", "add a todo item")
	flag.IntVar(&complete, "complete", 0, "mark an item as complete")
	flag.IntVar(&delete, "delete", 0, "delete an item")
	flag.BoolVar(&list, "list", false, "list todo items")

	flag.Usage = usage
}

func usage() {
	fmt.Println(os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	// args := flag.Args()
	// task = task + " " + strings.Join(args, " ")

	fileName := os.Getenv("TODO_FILENAME")
	if fileName == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		fileName = home + string(os.PathSeparator) + ".todo.json"
	}

	// create the file if it doesn't exist
	if _, err := os.Stat(fileName); err != nil {
		if errors.Is(err, os.ErrNotExist) {
			if _, err := os.Create(fileName); err != nil {
				fmt.Fprintln(os.Stderr, err)
			}
		}
	}

	// create a list struct and load it from the todo file
	li := &todo.List{}
	li.List(fileName)

	// handle input
	switch {
	case task != "":
		li.Add(task)

		if err := li.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}

	case list:
		if err := li.List(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for index, item := range *li {
			if !item.Done {
				fmt.Printf("%d [ ] %s\n", index+1, item.Task)
			} else {
				fmt.Printf("%d [x] %s\n", index+1, item.Task)

			}
		}
	case complete > 0:
		if err := li.Complete(complete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := li.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case delete > 0:
		if err := li.Delete(delete); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		if err := li.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	default:
		usage()
	}
}
