package DecisionTreeClassifier

import (
	"fmt"
	"math"
)

type DecisionTreeClassifier struct {
	root interface{}
}

type DecisionTreeClassifierer interface {
	BuildTree()
	FindBestSplit()
	InformationGain()
	GiniImpurity()
	Predict()
	PredictionMap()
	predict()
}

func NewDecisionTreeClassifier(rows [][]interface{}, maxDepth int) DecisionTreeClassifier {
	tree := new(DecisionTreeClassifier)
	tree.root = tree.BuildTree(rows, maxDepth)
	return *tree
}

func (dtc DecisionTreeClassifier) PredictionMap(row []interface{}) map[string]int {
	predictions := dtc.predict(row, dtc.root)
	return predictions
}

func PredictFromMap(predictionMap map[string]int) string {
	var max int = 0
	var bestLabel string
	for key, val := range predictionMap {
		if val >= max {
			max = val
			bestLabel = key
		}
	}
	return bestLabel
}

func (dtc DecisionTreeClassifier) Predict(row []interface{}) string {
	predictionByCounts := dtc.predict(row, dtc.root)
	return PredictFromMap(predictionByCounts)
}

func (dtc DecisionTreeClassifier) predict(row []interface{}, node interface{}) map[string]int {
	node__Leaf, isLeaf := node.(Leaf)

	if isLeaf {
		return node__Leaf.predictions
	}

	node__DecisionNode := node.(DecisionNode)

	if node__DecisionNode.decision.Match(row) {
		return dtc.predict(row, node__DecisionNode.trueBranch)
	} else {
		return dtc.predict(row, node__DecisionNode.falseBranch)
	}

}

func (dtc DecisionTreeClassifier) BuildTree(rows [][]interface{}, maxDepth int) interface{} {
	gain, decision := dtc.FindBestSplit(rows)

	if gain == 0 || maxDepth == 0 {
		return NewLeaf(rows)
	}

	trueRows, falseRows := Partition(rows, decision)

	trueBranch := dtc.BuildTree(trueRows, maxDepth-1)
	falseBranch := dtc.BuildTree(falseRows, maxDepth-1)

	return NewDecisionNode(decision, trueBranch, falseBranch)
}

func (dtc DecisionTreeClassifier) FindBestSplit(rows [][]interface{}) (float64, Decision) {
	var bestGain float64 = 0
	var bestDecision Decision
	currentUncertainty := dtc.GiniImpurity(rows)
	numFeatures := len(rows[0]) - 1

	for col := 0; col < numFeatures; col++ {
		values := UniqueValues(GetColumn(rows, col))

		for _, value := range values {
			dec := NewDecision(col, value)

			trueRows, falseRows := Partition(rows, *dec)

			if len(trueRows) == 0 || len(falseRows) == 0 {
				continue
			}

			gain := dtc.InformationGain(trueRows, falseRows, currentUncertainty)

			if gain >= bestGain {
				bestGain, bestDecision = gain, *dec
			}
		}
	}
	return bestGain, bestDecision
}

func (dtc DecisionTreeClassifier) GiniImpurity(rows [][]interface{}) float64 {
	counts := ClassCounts(rows)
	impurity := 1.0
	for _, occurrances := range counts {
		var probOfOccurrance float64 = float64(occurrances) / float64(len(rows))
		impurity -= math.Pow(probOfOccurrance, 2)
	}
	return impurity
}

func (dtc DecisionTreeClassifier) InformationGain(left [][]interface{}, right [][]interface{}, currentUncertainty float64) float64 {
	p := float64(len(left)) / float64(len(left)+len(right))
	return currentUncertainty - p*dtc.GiniImpurity(left) - (1-p)*dtc.GiniImpurity(right)
}

func (dtc DecisionTreeClassifier) PrintTree() {
	dtc.printTree(dtc.root, "")
}

func (dtc DecisionTreeClassifier) printTree(node interface{}, spacing string) {
	node__Leaf, isLeaf := node.(Leaf)

	if isLeaf {
		fmt.Println("Predict", node__Leaf.predictions)
		return
	}

	node.(DecisionNode).decision.PrintFormattedDecision()

	fmt.Println("--> True:")
	dtc.printTree(node.(DecisionNode).trueBranch, "  ")

	fmt.Println("--> False:")
	dtc.printTree(node.(DecisionNode).falseBranch, "  ")
}
