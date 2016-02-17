package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/robertkrimen/otto"
)

const src = `console.log("Hello, Otto!")`

func main() {
	flag.Parse()

	err := func() error {
		vm := otto.New()
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
