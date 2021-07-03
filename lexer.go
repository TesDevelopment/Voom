package main

import (
	"log"
	"strings"
)

func lex(script string, filename string) {
	var heap = make(map[string]string)
	for num, ln := range strings.Split(script, "\n") {
		var out string = ""
		var nonint bool = false
		for _, token := range strings.Split(ln, "") {
			token = strings.ToLower(token)
			if checkInt(token) {
				out = out + "|/|" + token
			} else {
				nonint = true
				out = out + token
			}
		}
		if !nonint {
			log.Fatalf("Invalid call on ln %v", num)
		}
	}
	parseout := parse(heap)
	if parseout != "" {
		log.Fatalf("Error occured on execution of script %v", parseout)
	}
	log.Printf("Ran script %v with no errors", filename)
}
