package eval

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

type algorithm struct {
	opPrecedence map[byte]int
	operators    stack[byte]
	numbers      stack[float64]
	mainError    error
	exprStr      string
	index        int
	numberBuffer strings.Builder
}

func (A *algorithm) insertNumber() {

	if A.numberBuffer.Len() == 0 {
		return
	}

	resultNumber, err := strconv.ParseFloat(A.numberBuffer.String(), 64)

	if err != nil {
		A.mainError = err
	} else {
		A.numbers.push(resultNumber)
	}

	A.numberBuffer.Reset()
	A.numberBuffer.Grow(len(A.exprStr))
}

func (A *algorithm) doArthmOps(condition func(byte) bool) {

	if A.operators.len() == 0 || condition(A.operators.peek()) {
		return
	}

	if A.numbers.len() < 2 {
		A.mainError = errors.New(ERROR_NOT_ENOUGH_NUMBERS)
		return
	}

	result, err := simpleArthm(A.numbers.pop(), A.numbers.pop(), A.operators.pop())

	if err != nil {
		A.mainError = err
	} else {
		A.numbers.push(result)
		A.doArthmOps(condition)
	}
}

func (A *algorithm) readExpression() {
	if A.index >= len(A.exprStr) || A.mainError != nil {
		return
	}

	switch {

	case unicode.IsDigit(rune(A.exprStr[A.index])) || A.exprStr[A.index] == '.':

		A.numberBuffer.WriteByte(A.exprStr[A.index])

		if A.index+1 >= len(A.exprStr) ||
			!(unicode.IsDigit(rune(A.exprStr[A.index+1])) || A.exprStr[A.index+1] == '.') {
			A.insertNumber()
		}

	case A.exprStr[A.index] == '(':

		if A.exprStr[A.index+1] == '-' {
			A.numberBuffer.WriteByte('-')
			A.index++
		}
		A.operators.push('(')

	case A.opPrecedence[A.exprStr[A.index]] != 0:

		A.doArthmOps(func(lastOp byte) bool {
			return lastOp == '(' || A.opPrecedence[lastOp] < A.opPrecedence[A.exprStr[A.index]]
		})

		A.operators.push(A.exprStr[A.index])

	case A.exprStr[A.index] == ')':

		A.doArthmOps(func(lastOp byte) bool {
			return lastOp == '('
		})

		if A.operators.len() != 0 {
			A.operators.pop()
		} else {
			A.mainError = errors.New(ERROP_INVALID_PARENTHESIS)
		}

	case A.exprStr[A.index] == ' ':

	default:
		A.mainError = errors.New(ERROR_INVALID_CHARACTER + string(A.exprStr[A.index]))
	}

	A.index++
	A.readExpression()
}

func (A *algorithm) calculateResult() (float64, error) {

	A.opPrecedence = map[byte]int{'+': 1, '-': 1, '*': 2, '/': 2}
	A.numberBuffer.Grow(len(A.exprStr))

	A.readExpression()

	if A.mainError != nil {
		return 0, A.mainError
	}

	A.doArthmOps(func(lastOp byte) bool {
		return false
	})

	if A.numbers.len() != 1 || A.mainError != nil {
		return 0, A.mainError
	}

	return A.numbers.pop(), A.mainError
}
