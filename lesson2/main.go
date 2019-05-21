package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	operator := map[int]string{
		1: "Megafon",
		2: "Beeline",
		3: "Babilon",
		4: "Tcell",
	}
	var id string
	megafon := []string{"90", "41", "55", "88"}
	beeline := []string{"919", "915", "911", "91"}
	babilon := []string{"98", "918", "94"}
	tcell := []string{"93", "92", "50", "77", "70"}
	uOperator := map[int][]string{
		1: megafon,
		2: beeline,
		3: babilon,
		4: tcell,
	}
	fmt.Println("Welcome to our pay terminal!")
	fmt.Println("Choose your operator:")
	fmt.Println("1 - Megafon\n2 - Beeline\n3 - Babilon\n4 - Tcell")
	//Block that checkes if the operator typed correctly
	var exitTerminal bool
	for !exitTerminal {
		var operatorCheck bool
		var numberCheck bool
		var sumCheck bool
		/*var commandCheck bool
		var command string*/
		for !operatorCheck {
			fmt.Scan(&id)
			if id == "esc" {
				numberCheck = true
				sumCheck = true
				exitTerminal = true
			} else if id == "back" {
				numberCheck = true
				sumCheck = true
			}
			id, err := strconv.Atoi(id)
			if err == nil {
				if (id >= 1) || (id <= 4) {
					operatorCheck = true
					for i, oper := range operator {
						if id == i {
							fmt.Printf("Please type your phone number (%s): ", oper)
						}
					}
				} else {
					fmt.Println("Please choose right operator 1")
				}
			}
		}
		id, err := strconv.Atoi(id)
		if err == nil {
			continue
		}
		//checking if number dimension is correct
		var number string
		for !numberCheck {
			fmt.Scan(&number)
			if number == "esc" {
				sumCheck = true
				exitTerminal = true
			} else if number == "back" {
				sumCheck = true
			}
			//checking if number dimension is correct
			if (number < "410000000") || (number > "989999999") {
				fmt.Println("Please type right phone number 2")
			} else {
				counter := 0
				//if the dim is corrct, check if the prefix is correct
				for i, uOper := range uOperator {
					//finds the needed operator
					if id == i {

						//goes through all prefixes of operator
						for j := 0; j < len(uOper); j = j + 1 {
							//checks if any of them matches
							if number[:len(uOper[j])] == uOper[j] {
								counter = counter + 1
							}
						}

					}
				}
				//if there is no match ask for reiteration
				if counter == 0 {
					fmt.Print("Please type right phone number 1")
				} else {
					//else move forvard
					numberCheck = true
				}
			}
		}
		var amount string
		fmt.Print("Amount: ")
		for !sumCheck {
			fmt.Scan(&amount)
			if number == "esc" {
				//sumCheck = true
				exitTerminal = true
			} else if number == "back" {
				continue
				//sumCheck = true
			}
			amount, err := strconv.ParseFloat(amount, 64)
			if err == nil {
				if amount < 0.0 {
					fmt.Println("Please type correct amount")
				} else {
					sumCheck = true
				}
			}
		}
		currentTime := time.Now()
		fmt.Println("------------------------------")
		fmt.Printf("Operator : %s \nNumber: %d \nAmount %v \nDate: %s \nOperationStatus: Successful \n", operator[id], number, amount, currentTime.Format("2006.01.02 15:04"))
		fmt.Println("------------------------------")
	}
}
