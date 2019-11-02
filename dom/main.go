package main

import (
	"syscall/js"
)

func hello(this js.Value, args []js.Value) interface{} {
	doc := js.Global().Get("document")

	// title = document.createElement("h1")
	// title.innerText = "Hello World"
	title := doc.Call("createElement", "h1")
	title.Set("innerText", "Hello World")

	// document.body.append(title)
	doc.Get("body").Call("append", title)
	return nil
}

func main() {
	js.Global().Set("hello", js.FuncOf(hello))
	select {} // block the main thread forever
}
