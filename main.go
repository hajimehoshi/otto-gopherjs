package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/gopherjs/gopherjs/js"
	"github.com/robertkrimen/otto"
)

const src = `
var div = document.createElement('div');
div.textContent = 'Hello, Otto!';
document.body.appendChild(div);`

func main() {
	flag.Parse()

	err := func() error {
		vm := otto.New()
		vm.Set("document", js.Global.Get("document"))
		_, err := vm.Run(src)
		return err
	}()
	if err != nil {
		switch err := err.(type) {
		case *otto.Error:
			fmt.Print(err.String())
		default:
			fmt.Println(err)
		}
		os.Exit(64)
	}
}
