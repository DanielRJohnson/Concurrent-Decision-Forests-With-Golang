package DecisionTreeClassifier

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"time"
)

type RandomForestClassifier struct {
	forest        []DecisionTreeClassifier
	numEstimators int
	maxDepth      int
}

type RandomForestClassifierer interface {
	Fit()
	Predict()
}

func NewRandomForestClassifier(numEstimators int, maxDepth int) RandomForestClassifier {
	rfc := new(RandomForestClassifier)
	rfc.forest = make([]DecisionTreeClassifier, numEstimators)
	rfc.numEstimators = numEstimators
	rfc.maxDepth = maxDepth
	return *rfc
}

func (rfc RandomForestClassifier) Fit(rows [][]interface{}) {
	var wg sync.WaitGroup
	for i := range rfc.forest {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()

			// Random subset of the rows with length = sqrt(len(rows))
			subset := make([][]interface{}, int(math.Sqrt(float64(len(rows)))))
			rand.Seed(time.Now().UnixNano())
			for j := 0; j < len(subset); j++ {
				subset[j] = rows[rand.Intn(len(rows))]
			}

			rfc.forest[i] = NewDecisionTreeClassifier(subset, rfc.maxDepth)
		}(i)
	}
	wg.Wait()
}

func (rfc RandomForestClassifier) Predict(row []interface{}) string {
	predictions := make(map[string]int)
	for _, tree := range rfc.forest {
		prd := tree.PredictionMap(row)
		for key := range prd {
			predictions[key] += prd[key]
		}
	}
	fmt.Println(predictions)
	return PredictFromMap(predictions)
}
