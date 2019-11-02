package main

import (
	"syscall/js"
)

func add(this js.Value, args []js.Value) interface{} {
	a, b := args[0].Int(), args[1].Int()
	return a + b
}

func main() {
	js.Global().Set("add", js.FuncOf(add))
	select {} // block the main thread forever
}
