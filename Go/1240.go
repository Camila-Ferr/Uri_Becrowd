package main

import (
	"fmt"
	"math"
	"strconv"
)

func compare(a string, b string) bool {
	i := 0
	for i < len(a) {
		if a[i] != b[i] {
			return false
		}
		i++
	}
	return true
}

func main() {
	var a, b, casos int
	fmt.Scanf("%d", &casos)
	i := 0
	for i < casos {
		fmt.Scanf("%d", &a)
		fmt.Scanf("%d", &b)

		lenB := len(strconv.Itoa(b))
		novoA := a % (int(math.Pow10(lenB)))
		if compare(strconv.Itoa(novoA), strconv.Itoa(b)) {
			fmt.Println("encaixa")
		} else {
			fmt.Println("nao encaixa")
		}
		i += 1
	}
}