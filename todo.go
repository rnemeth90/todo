package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

// List represents a list of todo items
type List []item

// Add will add a todo item to the list
func (l *List) Add(taskName string) {

	t := item{
		Task:        taskName,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

// Complete marks an item in the list as complete
func (l *List) Complete(i int) error {
	ls := *l
	for k := range *l {
		if i-1 == k {
			ls[i-1].Done = true
			ls[i-1].CompletedAt = time.Now()
		}
	}
	return nil
}

// Delete is used to delete an item from the todo file
func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

// Save is used to save the list of items to the todo file
func (l *List) Save(fileName string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, js, 0644)
}

// Get returns a list of todo items from the todo file
func (l *List) Get(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s does not exist", fileName)
		}
		return err
	}

	if len(file) > 0 {
		return json.Unmarshal(file, l)
	}
	return errors.New("File is empty")
}

func (l *List) String() string {
	formatted := ""

	for k, t := range *l {
		prefix := "[ ] "
		if t.Done {
			prefix = "[x] "
		}

		formatted += fmt.Sprintf("%s%d: %s\n", prefix, k+1, t.Task)
	}
	return formatted
}
