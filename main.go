package main

import "fmt"

func main() {
	romMap := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}

	var romVal string
	fmt.Print("Please input a Roman Numeral: ")
	fmt.Scanln(&romVal)
	var total int

	// loops through the length of the user input starting from the last letter
	for i := len(romVal) - 1; i > -1; i-- {
		var currLetter string = string(romVal[i])
		var currVal int = romMap[currLetter]

		// checks the next letter if you are able to subtract
		// 0 is so there is no out of bounds error
		if i != 0 {
			var nextVal int = romMap[string(romVal[i-1])]
			if currVal > nextVal {
				total += (currVal - nextVal)
				i--
			} else {
				total += currVal
			}
		} else {
			total += currVal
		}
	}

	fmt.Println(total)
}
