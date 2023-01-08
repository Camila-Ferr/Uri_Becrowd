package main

import (
	"fmt"
	"strconv"
)

func startToEnd(start int, end int) {
	for start <= end {
		fmt.Print(start)
		start += 1
	}
}

func endToStart(start int, end int, end_string string) {
	i := 0

	for end >= start {
		for i < len(end_string) {
			fmt.Printf("%c", end_string[len(end_string)-1-i])
			i += 1
		}
		end -= 1
		end_string = strconv.FormatInt(int64(end), 10)
		i = 0
	}

}

func main() {
	var start, end, numberValues int
	fmt.Scanf("%d", &numberValues)

	cont := 0
	for cont < numberValues {
		fmt.Scanf("%d", &start)
		fmt.Scanf("%d", &end)
		end_string := strconv.FormatInt(int64(end), 10)
		startToEnd(start, end)
		endToStart(start, end, end_string)
		fmt.Printf("\n")

		cont = cont + 1
	}

}
