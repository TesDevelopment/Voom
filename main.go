package main

import (
	"io/ioutil"
	"os"
)

func main() {
	goargs := os.Args

	if len(goargs) != 2 {
		return
	}

	script, err := ioutil.ReadFile(goargs[1])

	if err != nil {
		return
	}

	lex(string(script), goargs[1])
}
