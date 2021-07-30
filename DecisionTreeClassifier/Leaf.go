package DecisionTreeClassifier

type Leaf struct {
	predictions map[string]int
}

func NewLeaf(rows [][]interface{}) Leaf {
	l := new(Leaf)
	l.predictions = ClassCounts(rows)
	return *l
}
