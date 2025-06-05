package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"slices"
	"time"
	"github.com/alexeyco/simpletable"
)

type item struct {
	Task string
	Done bool
	CreatedAt time.Time
	CompletedAt time.Time
}
type Todo []item

func (t *Todo) Add(task string) {
	todo := item{
		Task: task,
		Done: false,
		CreatedAt: time.Now(),
		CompletedAt: time.Time{},
	}

	*t = append(*t, todo)
}

func (t *Todo) Complete(index int) error {
	if index <=0 || index > len(*t){
		return errors.New("invalid index")
	}
	ls := *t
	ls[index-1].Done = true
	ls[index-1].CompletedAt = time.Now()

	return nil
}

func (t *Todo) Delete(index int) error {
	ls := *t
	if index <=0 || index > len(ls){
		return errors.New("invalid index")
	}
	*t = slices.Delete(ls, index-1, index)

	return nil
}

func (t *Todo) Load(filename string) error{
	file, err := os.ReadFile(filename)
	if err != nil{
		if errors.Is(err, os.ErrNotExist) {
			return nil // File does not exist, nothing to load
		}
		return err
	}

	if len(file) == 0 {
		return nil // File is empty, nothing to load
	}
	err = json.Unmarshal(file, t)
	if err != nil {
		return err
	}
	return nil
}

func (t *Todo) Store(filename string) error {

	data, err := json.Marshal(t)
	if err != nil {
		return err
	}

	return os.WriteFile(filename, data, 0644)
}

func (t *Todo) Print(){
	table := simpletable.New()

	table.Header = &simpletable.Header{
		Cells: []*simpletable.Cell{
			{Align: simpletable.AlignCenter, Text: "#"},
			{Align: simpletable.AlignCenter, Text: "Task"},
			{Align: simpletable.AlignCenter, Text: "Done?"},
			{Align: simpletable.AlignRight, Text: "CreatedAt"},
			{Align: simpletable.AlignRight, Text: "CompletedAt"},
		},
	}
	var cells [][]*simpletable.Cell

	for i, item := range *t {
		i++
		task := blue(item.Task)
		done := gray("No")
		completedAt := ""
		if item.Done {
			task = green(fmt.Sprintf("\u2705 %s", item.Task))
			done = green("Yes")
		}
		if !item.CompletedAt.IsZero() {
			completedAt = item.CompletedAt.Format(time.RFC822)
		} else{
			completedAt = gray("Not completed yet")
		}
		cells = append(cells, []*simpletable.Cell{
			{Text: fmt.Sprintf("%d", i)},
			{Text: task},
			{Text: done},
			{Text: item.CreatedAt.Format(time.RFC822)},
			{Text: completedAt},
		})
	}
	table.Body = &simpletable.Body{Cells: cells}
	table.Footer = &simpletable.Footer{Cells: []*simpletable.Cell{
		{Align: simpletable.AlignCenter, Span: 5, Text: func() string {
								if t.CountPending() == 0 {
									return green("Hurray! No tasks pending.")
								}
								return red(fmt.Sprintf("Total pending tasks: %d", t.CountPending()))
							}(),
		},
	}}
	table.SetStyle(simpletable.StyleUnicode)

	table.Println()
}

func (t *Todo) CountPending() int{
	count := 0
	for _, item := range *t {
		if !item.Done {
			count++
		}
	}
	return count
}