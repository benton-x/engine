package engine

import (
	"context"
	"fmt"

	"src/common"
)

// Condition rule tree node
type Condition struct {
	ID       int           `json:"id"`
	CondKey  string        `json:"cond_key"`
	Op       common.OpType `json:"op"`
	CondVal  interface{}   `json:"cond_val"`
	Children []*Condition  `json:"children"`
	IsLeaf   bool          `json:"is_leaf"`
}

// NonRecursiveMatch non recursive rule match
func (c *Condition) NonRecursiveMatch(ctx context.Context, feature map[string]interface{}) bool {
	stack := common.NewStack()
	stack.Push(c)
	for !stack.Empty() && stack.Len() > 0 {
		top := stack.Pop()
		cur, ok := top.(*Condition)
		if !ok {
			fmt.Println("type assert failed")
			continue
		}
		if cur.Exec(ctx, feature) {
			fmt.Println("exec rule match failed")
			continue
		}
		if cur.IsLeaf {
			return true
		}
		for _, child := range cur.Children {
			stack.Push(child)
		}
	}
	return false
}

// Exec exec rule match
func (c *Condition) Exec(ctx context.Context, feature map[string]interface{}) bool {
	if val, ok := feature[c.CondKey]; ok {
		return common.Calc(ctx, c.Op, c.CondVal, val)
	}
	return false
}
