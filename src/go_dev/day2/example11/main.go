package main

import "fmt"

var a = "g"
var A = "W"
func main() {
	n()
	m()
	n()
}
func n() {
	fmt.Println(a)
}
func m() {
	a = "o"
	fmt.Println(a)
}

