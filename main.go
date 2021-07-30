package main

import (
	"fmt"
	"time"

	"github.com/danielrjohnson/DecisionTreeClassifier/DecisionTreeClassifier"
)

func main() {
	trainRows := [][]interface{}{
		{"Green", 3, "Apple"},
		{"Yellow", 3, "Apple"},
		{"Red", 1, "Grape"},
		{"Red", 1, "Grape"},
		{"Yellow", 3, "Lemon"},
		{"Green", 3, "Apple"},
		{"Yellow", 3, "Apple"},
		{"Red", 1, "Grape"},
		{"Red", 1, "Grape"},
		{"Yellow", 3, "Lemon"},
		{"Green", 3, "Apple"},
		{"Yellow", 3, "Apple"},
		{"Red", 1, "Grape"},
		{"Red", 1, "Grape"},
		{"Yellow", 3, "Lemon"},
		{"Green", 3, "Apple"},
		{"Yellow", 3, "Apple"},
		{"Red", 1, "Grape"},
		{"Red", 1, "Grape"},
		{"Yellow", 3, "Lemon"},
		{"Green", 3, "Apple"},
		{"Yellow", 3, "Apple"},
		{"Red", 1, "Grape"},
		{"Red", 1, "Grape"},
		{"Yellow", 3, "Lemon"},
	}

	start := time.Now()
	forest := DecisionTreeClassifier.NewRandomForestClassifier(1000, 3)
	forest.Fit(trainRows)
	t := time.Now()
	elapsed := t.Sub(start)
	fmt.Println("Forest took", elapsed, "to make and fit.")

	testRows := [][]interface{}{
		{"Green", 3},
	}

	fmt.Println("Prediction:", forest.Predict(testRows[0]))
}
