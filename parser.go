package main

import (
	"fmt"
	"strconv"
	"strings"
)

var tokens = make(map[string]func(val string,etc ...string) string)
var globals = make(map[string]func(loc int) string)

var intstack []string
var stringstack []string
var funcstack []func(loc int) string

func parse(opcodes map[string][]string) string {
	//--Opcodes--\\
	Opcode("pushstring",pushstring)

	Opcode("pushing",pushint)

	Opcode("getglobal",getglobal)

	Opcode("deferints",defInt)

	Opcode("deferstrings",defString)

	Opcode("call",call)

	//--Globals--\\
	Global("prints",printS)

	Global("printi",printI)

	//--Actual code--\\

	for op, val := range opcodes {
		if tokens[op] == nil {
			return "Invalid operation passed"
		}
		var ret string
		if len(val) > 1{
			ret = tokens[op](val[0],val[1])
		}else{
			ret = tokens[op](val[0],"")
		}
		if ret != "" {
			return ret
		}
	}
	return ""
}

//--Utils--\\
func checkInt(a string) bool {
	var alpha string = "abcdefghijklmnopqrstuvwxyz"
	b := strings.Split(alpha, "")
	for _, let := range b {
		if strings.Contains(a, let) {
			return false
		}
	}
	return true
}

func Global(name string,f func(loc int) string){
	globals[name] = f
}

func Opcode(name string,f func(val string,etc ...string) string){
	tokens[name] = f
}

func Purify(val string) string{
	alpha := "abcdefghijklmnopqrstuvwxyz"
	var out string = ""
	for _,v := range strings.Split(val,""){
		if strings.Contains(alpha,v){
			out = out + v
		}
	}
	return out
}

func Fill(s []string) string{
	var out string = ""

	for _,v := range s{
		out = out + " " + v
	}

	return out
}

//--Token funcs --\\
func pushstring(val string,etc ...string) string {
	if checkInt(val) {
		return "Int pushed as string"
	}

	if etc[0] != ""{
		val = val + Fill(etc)
	}
	stringstack = append(stringstack, val)
	return ""
}

func pushint(val string,etc ...string) string {
	if !checkInt(val) {
		return "String pushed as int"
	}
	intstack = append(intstack, val)
	return ""
}

func getglobal(val string,etc ...string) string {
	val = Purify(val)
	if globals[val] == nil {
		return "Invalid global fetched"
	}
	funcstack = append(funcstack, globals[val])
	return ""
}

func defInt(val string,etc ...string) string {
	var temp []string
	intstack = temp
	return ""
}

func defString(val string,etc ...string) string {
	var temp []string
	stringstack = temp
	return ""
}

func call(val string,etc ...string) string {
	if !checkInt(val) {
		return "Invalid global call"
	}
	loc, err := strconv.Atoi(val)
	if err != nil {
		return "Invalid global call"
	}
	loc3, err2 := strconv.Atoi(etc[0])
	if err2 != nil {
		return "Invalid global call"
	}

	if len(funcstack) < 1 || len(funcstack) < loc {
		return "Invalid global called."
	}

	funcstack[loc](loc3)
	return ""
}

//--Globals--\\

func printS(loc int) string {
	if len(stringstack) < 1 || len(stringstack) < loc {
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
