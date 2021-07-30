package DecisionTreeClassifier

import (
	"strconv"
)

func ClassCounts(rows [][]interface{}) map[string]int {
	counts := make(map[string]int)
	for _, row := range rows {
		label := row[len(row)-1].(string)
		_, exists := counts[label]
		if !exists {
			counts[label] = 0
		}
		counts[label] += 1
	}
	return counts
}

func UniqueValues(slice []interface{}) []interface{} {
	values := make(map[string]bool)
	uniqueValues := make([]interface{}, 0)
	for i, entry := range slice {
		if _, exists := values[strconv.Itoa(i)]; !exists {
			values[strconv.Itoa(i)] = true
			uniqueValues = append(uniqueValues, entry)
		}
	}
	return uniqueValues
}

func GetColumn(slice [][]interface{}, index int) []interface{} {
	col := make([]interface{}, 0)
	for _, row := range slice {
		col = append(col, row[index])
	}
	return col
}

func Partition(rows [][]interface{}, decision Decision) ([][]interface{}, [][]interface{}) {
	trueRows := make([][]interface{}, 0)
	falseRows := make([][]interface{}, 0)
	for i := range rows {
		if decision.Match(rows[i]) {
			trueRows = append(trueRows, rows[i])
		} else {
			falseRows = append(falseRows, rows[i])
		}
	}
	return trueRows, falseRows
}
