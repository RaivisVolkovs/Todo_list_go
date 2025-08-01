package main

func main() {
	todos := Todos{}
	todos.add("Learn Go")
	todos.add("Build a Todo app")
	todos.toggle(0)
	todos.print()
}
