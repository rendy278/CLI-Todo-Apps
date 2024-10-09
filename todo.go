package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type Todo struct {
	Title       string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type Todos []Todo

// Menambahkan todo baru
func (todos *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}

	*todos = append(*todos, todo)
}

// Memvalidasi indeks
func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// Menghapus todo berdasarkan indeks
func (todos *Todos) delete(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)

	return nil
}

// Mengubah status completed todo
func (todos *Todos) toggle(index int) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	todo := &(*todos)[index]

	if !todo.Completed {
		completionTime := time.Now()
		todo.CompletedAt = &completionTime
	}

	todo.Completed = !todo.Completed

	return nil
}

// Mengedit todo
func (todos *Todos) edit(index int, title string) error {
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	(*todos)[index].Title = title
	return nil
}

// Mencetak todos dalam bentuk tabel
func (todos *Todos) print() {
	tbl := table.New(os.Stdout)
	tbl.SetRowLines(false)
	tbl.SetHeaders("#", "Title", "Completed", "Created At", "Completed At")

	for index, todo := range *todos {
		completed := "❌"
		completedAt := ""
		if todo.Completed {
			completed = "✅"
			if todo.CompletedAt != nil {
				completedAt = todo.CompletedAt.Format(time.RFC1123)
			}
		}
		tbl.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format(time.RFC1123), completedAt)
	}

	tbl.Render()
}
