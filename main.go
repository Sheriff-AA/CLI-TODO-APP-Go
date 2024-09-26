package main

func main() {
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.load(&todos)
	cmdFlags := NewCmdFlags()
	cmdFlags.execute(&todos)
	storage.save(todos)
}