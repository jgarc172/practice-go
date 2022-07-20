package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"
)

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

func (l *List) Save(fileName string) (err error) {
	bytes, err := json.Marshal(l)
	if err != nil {
		return
	}
	err = os.WriteFile(fileName, bytes, 0644)
	return
}

func (l *List) Get(fileName string) (err error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			err = nil
		}
		return
	}
	if len(bytes) == 0 {
		return
	}
	err = json.Unmarshal(bytes, l)
	return
}
