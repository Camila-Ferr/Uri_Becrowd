package main

import (
	"fmt"
)

func bubbleSort(notas []float64) {
	n := len(notas)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if notas[j] < notas[j+1] {
				notas[j], notas[j+1] = notas[j+1], notas[j]
			}
		}
	}
}

func main() {
	for {
		var cidadaos, quest,consultas int
		_, err := fmt.Scanf("%d %d", &cidadaos, &quest)
		
		if(err!= nil){
			break
		}
		notas := make([]float64, cidadaos)
		
		for i := 0; i < len(notas); i++ {
			fmt.Scan(&notas[i])
		}
		bubbleSort(notas)
		for i := 0; i < quest; i++ {
			fmt.Scan(&consultas)
			fmt.Printf("%d\n", int(notas[consultas-1]))
		}
  	}
}
