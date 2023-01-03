package main

import (
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

	fileName := os.Getenv("TODO_FILENAME")
	if fileName == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			fmt.Fprintf(os.Stderr, "")
		}
		fileName = home + string(os.PathSeparator) + ".todo.json"
	}

	li := &todo.List{}
	if err := li.List(fileName); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	switch {
	case task != "":
		li.Add(task)

		fmt.Printf("Adding %s\n", task)

		if err := li.Save(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
	case list:
		if err := li.List(fileName); err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(1)
		}
		for _, item := range *li {
			if !item.Done {
				fmt.Println(item.Task)
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
	}

}
