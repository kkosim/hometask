package main

import (
	"fmt"
	"strconv"
)

func main() {
	var exitTerminal bool
	var number string
	var sum float32
	operator := map[int]string{
		1: "Megafon",
		2: "Babilon",
		3: "Tcell",
		4: "Beeline",
	}

	for !exitTerminal {
		var operatorCheck bool
		var numberCheck bool
		var sumCheck bool
		var commandCheck bool
		var command string
		var id int
		PrintMainScreen()

		for !operatorCheck {
			fmt.Scan(&id)

			if id >= 1 && id <= 4 {
				operatorCheck = true
				for i, oper := range operator {
					if id == i {
						fmt.Printf("Please type your phone number (%s): ", oper)
					}
				}
			} else {
				fmt.Println("Please, choose right operator")
			}
		}
		for !numberCheck {
			fmt.Scan(&number)

			if CheckNumber(id, number) {
				numberCheck = true
			} else {
				fmt.Println("Please, enter right number")
			}
		}
		number, _ := strconv.Atoi(number)

		fmt.Print("Enter Amount: ")
		for !sumCheck {
			fmt.Scan(&sum)

			if sum > 0 {
				sumCheck = true
			} else {
				fmt.Println("Please, enter right sum")
			}
		}

		PrintCheque(operator, id, number, sum)
		fmt.Println("esc- to go exit, back - to go to terminal")

		for !commandCheck {
			fmt.Scan(&command)

			if command == "esc" {
				commandCheck = true
				exitTerminal = true
			} else if command == "back" {
				commandCheck = true
				continue
			}
		}
	}

	fmt.Println("application terminated")
}
