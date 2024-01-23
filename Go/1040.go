package main

import (
	"fmt"
)

func calculaMedia(notas []float64, vs bool) float64 {
	if !vs {
		return ((notas[0]*2 + notas[1]*3 + notas[2]*4 + notas[3]) / 10)
	}
	return ((notas[4] + notas[5]) / 2)
}

func main() {
	var notas = make([]float64, 4)
	var exame float64

	fmt.Scan(&notas[0], &notas[1], &notas[2], &notas[3])

	media := calculaMedia(notas, false)
	fmt.Printf("Media: %.1f\n", media)

	if media >= 7.0 {
		fmt.Println("Aluno aprovado.")
	} else if media < 5 {
		fmt.Println("Aluno reprovado.")
	} else {
		fmt.Println("Aluno em exame.")
		fmt.Scanf("%f", &exame)
		notas = append(notas, media, exame)
		fmt.Printf("Nota do exame: %.1f\n", exame)
		if exame >= 5.0 {
			fmt.Println("Aluno aprovado.")
		} else {
			fmt.Println("Aluno reprovado.")
		}
		fmt.Printf("Media final: %.1f\n", calculaMedia(notas, true))
	}

}
