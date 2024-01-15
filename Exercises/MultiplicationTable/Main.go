package main

import (
	"fmt"
)

func main() {

	const max = 5

	fmt.Println()
	// Print the top header numbers in the x axis
	fmt.Println(`    |
 X  |    0    1    2    3    4    5 `)
	fmt.Println(`____|_______________________________`)
	fmt.Println(`    |`)
	// Loop through each y axis number
	for y := 0; y <= max; y++ {
		fmt.Printf("%2d  |", y)
		// Loop through each x axis number
		for x := 0; x <= max; x++ {
			fmt.Printf("%5d", y*x)
		}
		fmt.Println()
	}
	fmt.Println()

}
