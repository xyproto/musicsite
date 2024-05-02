package main

import (
	"syscall/js"
)

func main() {
	console := js.Global().Get("console")
	console.Call("log", "WebAssembly module loaded and running!")
	js.Global().Set("add", js.FuncOf(add))
	<-make(chan bool) // Keep the wasm module running
}

func add(this js.Value, args []js.Value) interface{} {
	sum := args[0].Int() + args[1].Int()
	console := js.Global().Get("console")
	console.Call("log", "Adding numbers:", args[0], "+", args[1], "=", sum)
	return js.ValueOf(sum)
}
