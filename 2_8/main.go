package main

import "fmt"

func main() {
	var k, n int
	fmt.Scan(&k, &n)
	fmt.Println((k + n - 1) % 7)
}
