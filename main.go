package main

import (
	"flag"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
	"github.com/robertkrimen/otto"
)

const src = `1+1`

func main() {
	flag.Parse()

	vm := otto.New()
	result, err := vm.Run(src)
	if err != nil {
		switch err := err.(type) {
		case *otto.Error:
			fmt.Print(err.String())
		default:
			fmt.Println(err)
		}
		return
	}
	ans := js.Global.Get("document").Call("getElementById", "answer")
	ans.Set("textContent", result)
}
