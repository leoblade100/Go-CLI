package main

import (
	"fmt"
)

func main() {
	todos := Todos{}
	todos.add("buy milk")
	todos.add("buy products")
	fmt.Println("%+v\n\n", todos)
	todos.delete(2)
	todos.togle(1)
	fmt.Println("%+v\n\n", todos)
	todos.togle(1)
	fmt.Println("%+v\n\n", todos)
	todos.edit(0, "buy game")
	fmt.Println("%+v\n\n", todos)
	todos.print()
}
