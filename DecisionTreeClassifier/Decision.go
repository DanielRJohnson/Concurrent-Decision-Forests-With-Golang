package DecisionTreeClassifier

import (
	"fmt"
)

type Decision struct {
	column int
	value  interface{}
}

func NewDecision(column int, value interface{}) *Decision {
	d := new(Decision)
	d.column = column
	d.value = value
	return d
}

type Decisioner interface {
	Match()
	printFormattedDecision()
}

func (dec Decision) Match(exampleRow []interface{}) bool {
	val := exampleRow[dec.column]
	_, isString := val.(string)
	if !isString {
		_, isInt := val.(int)
		if isInt {
			return val.(int) >= dec.value.(int)
		}
		return val.(float64) >= dec.value.(float64)
	} else {
		return val == dec.value
	}
}

func (dec Decision) PrintFormattedDecision() {
	_, isFloat := dec.value.(float64)
	_, isInt := dec.value.(int)
	condition := "=="
	if isFloat || isInt {
		condition = ">="
	}
	fmt.Println("Is (label)", condition, dec.value)
}
