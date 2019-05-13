.package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%d santimeters is %d meters", a, a/100)
}
