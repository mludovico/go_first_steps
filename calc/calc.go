package calc

import "errors"

func Add(a int, b int) (int, error) {
	return a + b, nil
}

func Sub(a int, b int) (int, error) {
	return a - b, nil
}

func Mul(a int, b int) (int, error) {
	return a * b, nil
}

func Div(a int, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Division by zero")
	}
	return a / b, nil
}
