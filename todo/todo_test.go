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
		os.TempDir()
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
