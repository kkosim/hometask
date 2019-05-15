package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%d meters is %d kilometers'", a, a/1000)
}
