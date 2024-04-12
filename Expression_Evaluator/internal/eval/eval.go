package eval

type Evaluator interface {
	calculateResult() (float64, error)
}

func Evaluate(input string) (float64, error) {

	var evaluator Evaluator = &algorithm{exprStr: input}

	result, err := evaluator.calculateResult()

	if err != nil {
		return 0, err
	}

	result = roundFloat64(result, 3)

	return result, err
}
