package main

import (
	"fmt"
	"log"
	"strings"
)

func lex(script string, filename string) {
	var heap = make(map[string][]string)
	for num, ln := range strings.Split(script, "\n") {
		ln = strings.Replace(ln,"\n","",-1)
		var nonint bool = false
		for _, token := range strings.Split(ln, "") {
			token = strings.ToLower(token)
			if !checkInt(token) {
				nonint = true
			}
		}
		sp := strings.Split(ln," ")[0]
		oop := strings.Replace(ln,sp + " ","",-1)
		newoop := strings.Split(oop," ")

		sp = strings.ToLower(sp)
		heap[sp] = newoop
		if !nonint {
			log.Fatalf("Invalid call on ln %v", num)
		}
	}
	parseout := parse(heap)
	if parseout != "" {
		log.Fatalf("Error occured on execution of script: (%v)\n", parseout)
	}
	fmt.Printf("Ran script %v with no errors", filename)
}
