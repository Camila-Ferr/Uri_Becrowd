package main

import (
	"fmt"
	"math"
)

func verificaPrimo(numero int) bool {
	i := 3
	for i <= int(math.Sqrt(float64(numero))) {
		if numero%i == 0 {
			return false
		}
		i += 2
	}
	return true
}
func main() {
	var cases, numero int
	fmt.Scanf("%d", &cases)

	for cases != 0 {
		fmt.Scanf("%d", &numero)
		if numero < 3 || (numero%2 != 0 && verificaPrimo(numero)) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not Prime")
		}
		cases -= 1
	}
}
