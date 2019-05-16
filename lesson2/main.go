package main

import (
	"fmt"
	"time"
)

func main() {
	operator := map[string]string{
		"1": "Megafon",
		"2": "Beeline",
		"3": "Babilon",
		"4": "Tcell",
	}
	id := "5"
	megafon := []string{"90", "41", "55", "88"}
	beeline := []string{"919", "915", "911", "91"}
	babilon := []string{"98", "918", "94"}
	tcell := []string{"93", "92", "50", "77", "70"}
	uOperator := map[string][]string{
		"1": megafon,
		"2": beeline,
		"3": babilon,
		"4": tcell,
	}
	fmt.Println("Welcome to our pay terminal!")
	fmt.Println("Choose your operator:")
	fmt.Println("1 - Megafon\n2 - Beeline\n3 - Babilon\n4 - Tcell")
	//Block that checkes if the operator typed correctly
	for {
		fmt.Scan(&id)
		if (id < "1") || (id > "4") {
			fmt.Println("Please choose right operator")
		} else {
			break
		}
	}
	//type your number stage
	for i, oper := range operator {
		if id == i {
			fmt.Printf("Please type your phone number (%s): ", oper)
		}
	}
	//checking if number dimension is correct
	var number string
	for {
		fmt.Scan(&number)
		//checking if number dimension is correct
		if (number < "410000000") || (number > "989999999") {
			fmt.Println("Please type right phone number")
		} else {
			//if the dim is corrct, check if the prefix is correct
			for i, uOper := range uOperator {
				//finds the needed operator
				if id == i {
					counter := 0
					//goes through all prefixes of operator
					for j := 0; j < len(uOper); j = j + 1 {
						//checks if any of them matches
						if number[:len(uOper[j])] == uOper[j] {
							counter = counter + 1
						}
					}
					//if there is no concide ask for reiteration
					if counter == 0 {
						fmt.Println("Please type right phone number ---")
					} else {
						//else move forvard
						break
					}
				}
			}
			break
		}
	}
	var amount float32
	fmt.Print("Amount: ")
	for {
		fmt.Scan(&amount)
		if amount < 0 {
			fmt.Println("Please type correct amount")
		} else {
			break
		}
	}
	currentTime := time.Now()
	fmt.Println("------------------------------")
	fmt.Printf("Operator : %s \nNumber: %s \nAmount %v \nDate: %s \nOperationStatus: Successful \n", operator[id], number, amount, currentTime.Format("2006.01.02 15:04"))
	fmt.Println("------------------------------")

}
