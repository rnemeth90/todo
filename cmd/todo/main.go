package main

import (
	"errors"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/rnemeth90/todo"
)

var (
	add      string
	complete int
	delete   int
	list     bool
)

func init() {
	flag.StringVar(&add, "a", "", "add a todo item to the list")
	flag.IntVar(&complete, "c", 0, "mark an item in the list as complete")
	flag.IntVar(&delete, "d", 0, "delete an item from the list")
	flag.BoolVar(&list, "l", false, "list todo items")

	flag.Usage = usage
}

func usage() {
	fmt.Println(os.Args[0])
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

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
	if err := li.Get(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	// handle input
	switch {
	case add != "":
		li.Add(add + " " + strings.Join(flag.Args(), " "))

		if err := li.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case list:
		fmt.Println(li)
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
