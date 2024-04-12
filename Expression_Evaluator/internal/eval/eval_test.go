package eval

import (
	"fmt"
	"testing"
)

func TestEvaluate(t *testing.T) {
	var expressions map[string]float64 = map[string]float64{
		"2 * 4":                                8,
		"2 - 5":                                -3,
		"2 + 5":                                7,
		"2.35 - 7.178":                         -4.828,
		"2 / (16 - 2 * 6)":                     0.5,
		"3 * (-1)":                             -3,
		"((4.5 + 3.5) * (8.2  - 2.2) + 4) / 2": 26,
		"2 * 8) - 3":                           0,
	}

	for input, answer := range expressions {
		result, err := Evaluate(input)

		fmt.Printf("input: %v | result: %f | error: %v \n", input, result, err)

		if answer != result {
			t.Errorf("Want %v , got %v", answer, result)
		}
	}
}
