package main

import (
	"fmt"
)

func main() {
  var entrada int 
  fmt.Scanf("%d", &entrada)

	fmt.Printf("%d:%d:%d\n", entrada/3600, (entrada%3600)/60, entrada%60)
}
