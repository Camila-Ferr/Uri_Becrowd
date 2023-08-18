package main

import (
	"fmt"
)

func bubbleSort(original []float64, modificado []float64) {
	n := len(original)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if modificado[j] > modificado[j+1] {
				modificado[j], modificado[j+1] = modificado[j+1], modificado[j]
			}
		}
	}
	printArray(original, modificado, n)
}

func printArray(original []float64, modificado []float64, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", int(modificado[i]))
	}
	fmt.Println()
	for i := 0; i < n; i++ {
		fmt.Printf("%d\n", int(original[i]))
	}
}

func main() {
	numeros := make([]float64, 3)
	modificado := make([]float64, len(numeros))
  
	for i := 0; i < len(numeros); i++ {
		fmt.Scan(&numeros[i])
	}

	copy(modificado, numeros)
	bubbleSort(numeros, modificado)
}
