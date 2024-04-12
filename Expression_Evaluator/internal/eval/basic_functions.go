package eval

import (
	"errors"
	"math"
)

// performs arithmetic operation on two numbers
func simpleArthm(second, first float64, operator byte) (result float64, err error) {

	switch operator {
	case '+':
		result = first + second
	case '-':
		result = first - second
	case '*':
		result = first * second
	case '/':
		if first != 0 {
			result = first / second
		} else {
			err = errors.New(ERROR_DIVISION_BY_ZERO)
		}
	default:
		err = errors.New(ERROR_INVALID_OPERATOR + " " + string(operator))
	}

	return
}

func roundFloat64(input float64, prec int) (rounded float64) {
	shift := math.Pow(10, float64(prec))
	input *= shift

	rounded = math.Round(input)
	rounded /= shift

	return
}
