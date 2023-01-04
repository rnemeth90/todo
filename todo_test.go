package todo_test

import (
	"testing"

	"github.com/rnemeth90/todo"
)

func TestAdd(t *testing.T) {
	li := todo.List{}
	taskName := "Test Task 1"
	li.Add("Test Task 1")

	if li[0].Task != taskName {
		t.Errorf("expected %s, got %s", taskName, li[0].Task)
	}
}

func TestDelete(t *testing.T) {
	li := todo.List{}

	tasks := []string{
		"Test Task 1",
		"Test Task 2",
		"Test Task 3",
	}

	for _, v := range tasks {
		li.Add(v)
	}

	if li[0].Task != tasks[0] {
		t.Errorf("expected %s, got %s", tasks[0], li[0].Task)
	}

	li.Delete(2)
	if len(li) != 2 {
		t.Errorf("expected list to contain %d tasks", len(li))
	}
}

func TestComplete(t *testing.T) {
	li := todo.List{}

	item := "Test Task 1"
	li.Add(item)

	if li[0].Task != item {
		t.Errorf("expected %s, got %s", item, li[0].Task)
	}

	li.Complete(1)
	if li[0].Done != true {
		t.Errorf("expected %s to be complete, but it was not. Index %d, %v", li[0].Task, li[0].Index, li[0].Done)
	}
}

func TestListAndSave(t *testing.T) {

}
