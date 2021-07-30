package DecisionTreeClassifier

type DecisionNode struct {
	decision    Decision
	trueBranch  interface{}
	falseBranch interface{}
}

func NewDecisionNode(decision Decision, trueBranch interface{}, falseBranch interface{}) DecisionNode {
	dn := new(DecisionNode)
	dn.decision = decision
	dn.trueBranch = trueBranch
	dn.falseBranch = falseBranch
	return *dn
}
