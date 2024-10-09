package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type cmdFlags struct {
	Add    string
	Delete int
	Edit   string
	Toogle int
	List   bool
}

func NewCmdFlags() *cmdFlags {
	cf := cmdFlags{}
	flag.StringVar(&cf.Add, "add", "", "Add a new todo specify title")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a todo by index & specify a new title, id:new_title")
	flag.IntVar(&cf.Delete, "delete", -1, "Specify a todo by index to delete")
	flag.IntVar(&cf.Toogle, "toogle", -1, "Specify a todo by index to toogle")
	flag.BoolVar(&cf.List, "list", false, "List all todos")
	flag.Parse()
	return &cf
}

func (cf *cmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			fmt.Println("Error invalid format for edit. Please use id:new_title")
		}
		index, err := strconv.Atoi(parts[0])

		if err != nil {
			fmt.Println("Error: invalid index for edit")
			os.Exit(1)
		}
		todos.edit(index, parts[1])
	case cf.Toogle != -1:
		todos.toggle(cf.Toogle)

	case cf.Delete != -1:
		todos.delete(cf.Delete)

	default:
		fmt.Println("Invalid command")
	}
}
