package main

import (
	"fmt"
)

func main() {
	var testes, soma int
	fmt.Scanf("%d", &testes)
	soma = 0
	for i := 0; i < testes; i++ {
		var t, v int
		fmt.Scanf("%d %d", &t, &v)
		soma += t * v
	}
	fmt.Printf("%d\n", soma);
}
