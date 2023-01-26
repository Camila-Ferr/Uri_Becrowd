package main

import (
	"fmt"
	"math"
)

func main() {
	var alunos, cafeL, consumoMl float64
	fmt.Scanf("%f %f %f", &alunos, &cafeL, &consumoMl)
	fmt.Println(math.Ceil((alunos*consumoMl)/(cafeL*1000)) * cafeL)
}
