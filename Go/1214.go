package main

import (
	"fmt"
	"math"
)

func main() {
	var notes,cases int 
	fmt.Scanf("%d", &notes)

	for notes != 0 {
		fmt.Scanf("%d", &cases)
		for cases != 0{

		}
		if numero < 3 || (numero%2 != 0 && verificaPrimo(numero)) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
		notes -= 1
	}
}
