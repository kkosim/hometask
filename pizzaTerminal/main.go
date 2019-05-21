package main

import (
	"fmt"
	"os"
)

func main() {
	//var exitTerminal bool
	var size string
	var SizeCheck bool
	var speciesID int
	/*var number string
	var sum float32
	pizzaSize := map[string]float32{
		"Small":      0.75,
		"Medium":     1,
		"Big":        1.5,
		"For family": 2,
	}
	pizzaSizeId := map[int]string{

		1: "Small",
		2: "Medium",
		3: "Big",
		4: "For family",
	}*/
	PrintMainScreen()
	pizzaTypesIngredients := map[int]string{
		1:  "Margherita",
		2:  "Marinara",
		3:  "Assorti",
		4:  "Carbonara",
		5:  "Frutti di Mare",
		6:  "Quattro Formaggi",
		7:  "Crudo",
		8:  "Napoletana",
		9:  "Pugliese",
		10: "Montanara",
		11: "Emiliana",
		12: "Romana",
		13: "Fattoria",
		14: "Schiacciata",
		15: "prosciutto",
	}
	//pizzaPrice := []float32{31, 33, 32, 35, 36, 31, 33, 25, 36, 21, 39, 32, 30, 28, 32}

	for !SizeCheck {
		fmt.Scan(&size)

		if size == "esc" {
			os.Exit(1)
		}
		if size >= "1" && size <= "4" {
			SizeCheck = true
		} else {
			fmt.Println("Please choose right size")
		}
	}
	for i, spec := range pizzaTypesIngredients {
		fmt.Println("Please choose the pizza you prefer")
		fmt.Printf("%d -- %s", i, spec)
		fmt.Scan(&speciesID)
		for {
			if speciesID > 15 || speciesID < 1 {
				fmt.Println("Please select one of the following ")
			} else {
				break
			}
		}
	}

}
