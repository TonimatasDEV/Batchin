package operators

import (
	"fmt"

	"github.com/Knetic/govaluate"
)

/*
Calculate returns the result of the expression or the same input string if the expressions is not correct.
*/
func Calculate(expression string) string {
	exp, err := govaluate.NewEvaluableExpression(expression)
	if err != nil {
		return expression
	}

	result, err := exp.Evaluate(nil)
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%v", result)
}
