package main

import (
	"fmt"
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
	fmt.Println("---------------------------------------------")
	fmt.Printf("Size : %s \nSpecies: %s \nNumber %d \nDate: %s \nTotal Sum: %d \n",pizzaTypesIngredients[speciesID] , , ,currentTime.Format("2006.01.02 15:04"), )
	fmt.Println("---------------------------------------------")
}
