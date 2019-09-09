package main

import (
	"fmt"
	"syscall/js"
)

func main() {
	js.Global().Get("document").
		Call("getElementById", "myTextBox").
		Set("value", "hello wasm")
	fmt.Println("WebAssembly is running")
}
