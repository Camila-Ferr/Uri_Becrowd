package main

import (
	"fmt"
	"strconv"
)

func verifica(s string) bool {
	i := 0
	for i < len(s) {
		if s[i] == 55 {
			return true
		}
		i++
	}
	return false
}

func main() {
	var pessoas, pessoa, vezes, aparecimentos, numeros int
	pessoa = 0
	for true {
		fmt.Scanf("%d", &pessoas)
		fmt.Scanf("%d", &pessoa)
		fmt.Scanf("%d", &vezes)
		numeros = pessoa
		possivel := true
		i := 0

		if (pessoa == 0) && (pessoas == 0) && (vezes == 0) {
			break
		} else if pessoa > pessoas {
			possivel = false
			fmt.Println("-1")
		}

		aparecimentos = 0
		for aparecimentos != vezes && possivel {

			if (numeros%7 == 0) || verifica(strconv.Itoa(numeros)) {
				aparecimentos += 1
			}
			if aparecimentos != vezes {
				if (pessoa == pessoas) || (pessoa == 1) {
					numeros += 2*pessoas - 2
				} else if i%2 == 0 {
					numeros += 2 * (pessoas - pessoa)
				} else {
					numeros += -2 + 2*pessoa
				}
				i += 1
			}

		}
		if possivel {
			fmt.Println(numeros)
		}
	}
}
