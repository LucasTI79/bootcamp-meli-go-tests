package calc

// O pacote testing é importado
import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	// Os dados a serem usados ​​no teste são inicializados (entrada/saída)
	num1 := 3
	num2 := 5
	expectedResult := 8

	// O teste é executado
	result := Sum(num1, num2)

	assert.Equal(t, expectedResult, result, "devem ser iguais")
}

func TestDivide(t *testing.T) {
	num := 10
	den := 0
	expectedResult := 0

	result, err := Divide(num, den)

	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "denominador não pode ser zero")
	assert.Equal(t, expectedResult, result, "devem ser iguais")
}
