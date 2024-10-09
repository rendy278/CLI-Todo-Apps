package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	todos.add("I need  to buy milk")
	// todos.add("I like  to eat pizza")
	todos.toggle(1)
	todos.print()
	storage.Save(todos)
	// fmt.Println("%+v\n\n", todos)
	todos.delete(7)
	// fmt.Println("%+v", todos)
}
