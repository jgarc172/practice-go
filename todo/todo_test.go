package todo

import (
	"fmt"
	"testing"
	"time"
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

type item struct {
	Task        string
	Done        bool
	CreatedAt   time.Time
	CompletedAt time.Time
}

type List []item

func (l *List) Add(task string) {
	t := item{task, false, time.Now(), time.Time{}}
	*l = append(*l, t)
}

func (l *List) Complete(item int) (err error) {
	list := *l
	index := item - 1
	if index < 0 || index > (len(list)-1) {
		err = fmt.Errorf("item '%v' does not exist", item)
		return
	}
	list[index].Done = true
	list[index].CompletedAt = time.Now()
	return
}

func (l *List) Delete(item int) (err error) {
	list := *l
	index := item - 1
	if index < 0 || index > (len(list)-1) {
		err = fmt.Errorf("item '%v' does not exist", item)
		return
	}
	*l = append(list[:index], list[index+1:]...)
	return
}
