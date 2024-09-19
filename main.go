package main

import (
	"firstSteps/calc"
	"fmt"
)

func main() {
	var num1 int = 10
	var num2 int = 0
	var result, error = calc.Div(num1, num2)
	if error != nil {
		fmt.Println(error)
	} else {
		fmt.Println(result)
	}
}
