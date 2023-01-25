package main_test

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

var (
	binaryName   = "todo"
	todoFileName = ".testtodo.json"
)

func TestMain(m *testing.M) {

	todoFile := os.TempDir() + string(os.PathSeparator) + todoFileName

	fmt.Println("Building tool...")

	if runtime.GOOS == "windows" {
		binaryName += ".exe"
	}

	if err := os.Setenv("TODO_FILENAME", todoFile); err != nil {
		fmt.Fprintf(os.Stderr, "cannot set env var for todo file: %s", err)
		os.Exit(1)
	}
	envVar := os.Getenv("TODO_FILENAME")
	fmt.Println("using", envVar)

	cmd := exec.Command("go", "build", "-o", binaryName)

	if err := cmd.Run(); err != nil {
		fmt.Fprintf(os.Stderr, "cannot build tool: %s %s\n", binaryName, err)
		os.Exit(1)
	}

	fmt.Println("running...")
	result := m.Run()

	fmt.Println("cleaning up...")
	os.Remove(binaryName)
	os.Remove(todoFile)

	os.Exit(result)
}

func TestAddTask(t *testing.T) {
	task := "test task number 1"

	dir, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}

	cmdPath := filepath.Join(dir, binaryName)

	t.Run("AddNewTask", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-add", task)

		if err := cmd.Run(); err != nil {
			t.Fatal(err)
		}
	})

	t.Run("ListTasks", func(t *testing.T) {
		cmd := exec.Command(cmdPath, "-list")
		out, err := cmd.CombinedOutput()
		if err != nil {
			t.Fatal(err)
		}

		out = []byte(strings.TrimSuffix(string(out), " \n\n"))

		expected := fmt.Sprintf("[ ] 1: %s", task)

		if strings.Compare(string(out), expected) != 0 {
			t.Errorf("got %v, expected %v\n", out, []byte(expected))
		}
	})
}
