package main

import (
    "fmt"
)

func main() {
    var firstNumber, secondNumber int
    fmt.Scanf("%d %d", &firstNumber, &secondNumber)

   if (firstNumber> secondNumber) && (firstNumber%secondNumber == 0){
        fmt.Println("Sao Multiplos")
    } else if (secondNumber > firstNumber) && (secondNumber%firstNumber == 0){
        fmt.Println("Sao Multiplos")
    } else{
        fmt.Println("Nao sao Multiplos")
    }

}
