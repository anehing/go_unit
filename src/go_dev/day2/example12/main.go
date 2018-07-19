package main


import "fmt"

var a string

func main() {
	a = "G"
	print(a)
	f1()
}
func f1() {
	a := "1"
	fmt.Println(a)
	f2()
}
func f2() {
	fmt.Println(a)
}
