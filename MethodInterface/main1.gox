package main

import (
	"fmt"
)

// type level int

// func main() {
// 	sl := new(level)
// 	sl.raiseShieldLevel(4)
// 	fmt.Println(*sl)
// 	sl.raiseShieldLevel(5)
// 	fmt.Println(*sl)
// }

// func (lv *level) raiseShieldLevel(i int) {
// 	if *lv == 0 {
// 		*lv = level(1)
// 	}
// 	*lv = (*lv) * level(i)
// }

type Node interface {
	SetValue(v int)
	GetValue() int
}

type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return new(SLLNode)
}

type PowerNode struct {
	next  *PowerNode
	value int
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func (sNode *PowerNode) String() string {
	return "POWER NODE"
}

func NewPowerNode() *PowerNode {
	return new(PowerNode)
}

func main() {
	var node Node
	node = NewSLLNode()
	node.SetValue(4)
	fmt.Println("Node is of Value: ", node.GetValue())

	node = NewPowerNode()
	node.SetValue(5)
	fmt.Println("Node is of Value: ", node.GetValue())

	if n, ok := node.(*PowerNode); ok {
		fmt.Println("This is a power node of value ", n.value)
	}

	fmt.Println("Powernode: ", node)
}
