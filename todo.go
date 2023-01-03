package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"time"
)

// add
// complete
// delete
// save
// list

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(taskName string) {

	t := item{
		Task:        taskName,
		Done:        false,
		CreatedAt:   time.Now(),
		CompletedAt: time.Time{},
	}

	*l = append(*l, t)
}

func (l *List) Complete(i int) error {
	ls := *l
	for k, _ := range *l {
		if i == k-1 {
			ls[i].Done = true
			ls[i].CompletedAt = time.Now()
		}
	}
	return nil
}

func (l *List) Delete(i int) error {
	ls := *l
	if i <= 0 || i > len(ls) {
		return fmt.Errorf("Item %d does not exist", i)
	}
	*l = append(ls[:i-1], ls[i:]...)
	return nil
}

func (l *List) Save(fileName string) error {
	js, err := json.Marshal(l)
	if err != nil {
		return err
	}

	return ioutil.WriteFile(fileName, js, 0644)
}

func (l *List) List(fileName string) error {
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return fmt.Errorf("file %s does not exist", fileName)
		}
		return err
	}

	return json.Unmarshal(file, l)
}
