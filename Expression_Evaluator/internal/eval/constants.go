package eval

//Expression errors
const (
	ERROR_DIVISION_BY_ZERO    = "division by zero"
	ERROR_INVALID_OPERATOR    = "invalid operator"
	ERROR_INVALID_CHARACTER   = "invalid character"
	ERROR_NOT_ENOUGH_NUMBERS  = "not enough numbers"
	ERROP_INVALID_PARENTHESIS = "invalid parenthesis"
)

//Stack errors
const (
	EMPTY_STACK               = "empty stack: "
	ERROR_EMPTY_STACK_PEEKING = EMPTY_STACK + "peeking"
	ERROR_EMPTY_STACK_POPPING = EMPTY_STACK + "popping"
)
