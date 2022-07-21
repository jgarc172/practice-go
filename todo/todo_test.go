package todo

import (
	"os"
	"testing"
)

func TestTodo(t *testing.T) {
	tasks := List{}
	taskName := "New Task"

	t.Run("Add", func(t *testing.T) {
		tasks.Add(taskName)
		got := tasks[0].Task
		if got != taskName {
			t.Errorf("task is '%v', expected '%v'", got, taskName)
		}
	})
	t.Run("Save and Get", func(t *testing.T) {
		tempFile, err := os.CreateTemp("", "")
		if err != nil {
			t.Errorf("error creating tempFile '%v'", err)
		}
		defer os.Remove(tempFile.Name())

		err = tasks.Save(tempFile.Name())
		if err != nil {
			t.Fatalf("error saving tasks to file '%v'", err)
		}

		tasks = List{}
		err = tasks.Get(tempFile.Name())
		if err != nil {
			t.Fatalf("error getting tasks from file '%'", err)
		}
		got := tasks[0].Task
		if got != taskName {
			t.Errorf("task is '%v', expected '%v'", got, taskName)
		}
	})
	t.Run("Complete", func(t *testing.T) {
		done := tasks[0].Done
		if done {
			t.Errorf("done is '%v', expected '%v'", done, false)
		}
		tasks.Complete(1)
		complete := tasks[0].Done
		if !complete {
			t.Errorf("completed is '%v', expected '%v'", complete, false)
		}
	})
	t.Run("Delete", func(t *testing.T) {
		tasks.Delete(1)
		length := len(tasks)
		if length != 0 {
			t.Errorf("length is '%v', expected '%v'", length, 0)
		}
	})
}

func TestGetNil(t *testing.T) {
	tempFile, err := os.CreateTemp("", "")
	if err != nil {
		t.Fatalf("error creating tempFile, '%v'", err)
	}
	defer os.Remove(tempFile.Name())

	for _, tc := range casesGetNil {
		t.Run(tc.name, func(t *testing.T) {
			if tc.name == "exists" {
				tc.fileName = tempFile.Name()
			}
			err := tc.tasks.Get(tc.fileName)
			if err != nil {
				t.Errorf("error when opening file, '%v'", err)
			}
		})
	}
}

var casesGetNil = []struct {
	name     string
	tasks    List
	fileName string
	err      error
}{
	{"no name", List{}, "", nil},
	{"not exist", List{}, "notexist.txt", nil},
	{"exists", List{}, "temp", nil},
}
