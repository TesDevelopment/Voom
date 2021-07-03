package main

import (
	"fmt"
	"strconv"
	"strings"
)

var tokens = make(map[string]func(val string, secondary string) string)
var globals = make(map[string]func(loc int) string)

var intstack []string
var stringstack []string
var funcstack []func(loc int) string

func parse(opcodes map[string]string) string {
	//--Opcodes--\\
	tokens["pushstring"] = pushstring

	tokens["pushing"] = pushint

	tokens["getglobal"] = getglobal

	tokens["deferints"] = defInt

	tokens["deferstrings"] = defString

	tokens["call"] = call

	//--Globals--\\
	globals["prints"] = printS

	globals["printi"] = printI

	//--Actual code--\\

	for op, val := range opcodes {
		if tokens[op] == nil {
			return "Invalid operation passed"
		}
		val2 := strings.Split(val, "|/|")[1]
		var ret string = tokens[op](val, val2)
		if ret != "" {
			return ret
		}
	}
	return ""
}

//--Utils--\\
func checkInt(a string) bool {
	var alpha string = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := strings.Split(alpha, "")
	for _, let := range b {
		if strings.Contains(a, let) {
			return false
		}
	}
	return true
}

//--Token funcs --\\
func pushstring(val string, s string) string {
	if checkInt(val) {
		return "Int pushed as string"
	}
	stringstack = append(stringstack, val)
	return ""
}

func pushint(val string, s string) string {
	if !checkInt(val) {
		return "String pushed as int"
	}
	intstack = append(intstack, val)
	return ""
}

func getglobal(val string, s string) string {
	if globals[val] == nil {
		return "Invalid global"
	}
	funcstack = append(funcstack, globals[val])
	return ""
}

func defInt(val string, s string) string {
	var temp []string
	intstack = temp
	return ""
}

func defString(val string, s string) string {
	var temp []string
	stringstack = temp
	return ""
}

func call(val string, loc2 string) string {
	if !checkInt(val) {
		return "Invalid global call"
	}
	loc, err := strconv.Atoi(val)
	if err != nil {
		return "Invalid global call"
	}
	loc3, err2 := strconv.Atoi(loc2)
	if err2 != nil {
		return "Invalid global call"
	}
	funcstack[loc](loc3)
	return ""
}

//--Globals--\\

func printS(loc int) string {
	if len(stringstack) < loc {
		return "Attempted to print a nonexistant string"
	}
	fmt.Println(stringstack[loc])
	return ""
}
func printI(loc int) string {
	if len(intstack) < loc {
		return "Attempted to print a nonexistant int"
	}
	fmt.Println(intstack[loc])
	return ""
}
