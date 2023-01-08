package main

import (
	"fmt"
	"math/big"
)

func count(bin string) {
	result, i := 0, 0

	for i < len(bin) {

		if bin[i] == 49 {
			result += 1
		}
		i += 1
	}

	fmt.Println(result)
}

func main() {
	var cases, actual int
  var cardinal big.Int
	fmt.Scanf("%d", &cases)
	actual = 0

	for actual < cases {

		fmt.Scanf("%d", &cardinal)
    bin:= cardinal.Text(2)
		count(bin)
		actual += 1
	}

}
