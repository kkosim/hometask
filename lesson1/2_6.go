package main

import "fmt"

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Printf("%d hours %d minutes %d seconds", a/3600, (a%3600)/60, a%60)
}
