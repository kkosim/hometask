package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	fmt.Println((k + 1) % 12)
}
