package main

import (
	"syscall/js"
)

func add(this js.Value, inputs []js.Value) interface{} {
	a, b := inputs[0].Int(), inputs[1].Int()
	return a + b
}

func main() {
	js.Global().Set("add", js.FuncOf(add))
	select {} // block the main thread forever
}
