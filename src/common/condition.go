package common

// Condition rule tree node
type Condition struct {
	ID       int          `json:"id"`
	CondKey  string       `json:"cond_key"`
	Op       OpType       `json:"op"`
	CondVal  interface{}  `json:"cond_val"`
	Children []*Condition `json:"children"`
	IsLeaf   bool         `json:"is_leaf"`
}
