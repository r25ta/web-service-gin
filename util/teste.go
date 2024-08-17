package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Os números pares são:")

	for i := 0; i < 21; i++ {

		if isPar(i) {
			fmt.Printf("\n%d", i)
		}
	}
}

func isPar(n1 int) bool {
	if math.Mod(float64(n1), 2) != 0 {
		return false
	}

	return true

}
