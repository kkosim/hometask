package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%d kilograms is %d tons'", a, a/1000)
}
