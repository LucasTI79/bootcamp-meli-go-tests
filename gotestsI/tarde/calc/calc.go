package calc

import "errors"

// Função que recebe dois inteiros e retorna a soma resultante
func Sum(num1, num2 int) int {
	return num1 + num2
}

// Função que recebe dois inteiros e retorna a diferença
// ou diferença resultante
func Sub(num1, num2 int) int {
	return num1 - num2
}

// Função que recebe dois inteiros (numerador e denominador) e retorna a divisão resultante
func Divide(num, den int) (int, error) {
	if den == 0 {
		return 0, errors.New("denominador não pode ser zero")
	}

	return num / den, nil
}
