package main

import (
	"fmt"
	"strings"
	"time"
)

func PrintMainScreen() {
	fmt.Println("Welcome to our pizza terminal!")
	fmt.Println("Choose the size of pizza:")
	fmt.Println("1 - Small")
	fmt.Println("2 - Medium")
	fmt.Println("3 - Big")
	fmt.Println("4 - For family")
}

func PrintCheque(operator map[int]string, id int, number int, amount float32) {
	currentTime := time.Now()
	fmt.Println("------------------------------")
	fmt.Printf("Operator : %s \nNumber: %s \nAmount %v \nDate: %s \nOperationStatus: Successful \n", operator[id], number, amount, currentTime.Format("2006.01.02 15:04"))
	fmt.Println("------------------------------")
}

// Need to modify
func CheckNumber(operatorID int, number string) bool {
	var checkPrefix bool
	operatorPreffix := map[int]string{
		1: "90,44,55,88",
		2: "98,918,94",
		3: "93,50,92,77,70",
		4: "919,917,915,91",
	}

	if len(number) != 9 {
		return false
	}

	prefix := operatorPreffix[operatorID]
	arrPrefix := strings.Split(prefix, ",")

	for _, _prefixValue := range arrPrefix {
		checkPrefix = strings.HasPrefix(number, _prefixValue)

		if checkPrefix {
			return true
		}
	}

	return checkPrefix
}
