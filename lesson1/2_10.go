package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	n1 := n / 10
	n2 := n % 10
	n3 := n1 + n2
	n4 := n1 * n2
	fmt.Printf("tens = %d \nunits = %d \nsumm = %d \nproduct = %d", n1, n2, n3, n4)
}
