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
	fmt.Println(A)
}
func m() {
	A := "o"
	fmt.Println(A)
}

