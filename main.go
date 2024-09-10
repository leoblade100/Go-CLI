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
	fmt.Println("%+v", todos)
}
