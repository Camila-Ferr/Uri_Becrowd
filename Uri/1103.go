package main

import "fmt"

func main() {
	var h1, m1, h2, m2, sono int

	for {
		fmt.Scanln(&h1, &m1, &h2, &m2)
		if h1 == 0 && m1 == 0 && h2 == 0 && m2 == 0 {
			break
		}
		sono = (h2-h1)*60 + (m2-m1)
		if sono < 0 {
			sono += 24 * 60
		}
		fmt.Println(sono)
	}
}