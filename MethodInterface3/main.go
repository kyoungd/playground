package main

import (
	"errors"
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
	SetValue(v int) error
	GetValue() int
}

type SLLNode struct {
	next  *SLLNode
	value int
}

func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return errors.New("Node is not instantiated.  It is nil.")
	}
	sNode.value = v
	return nil
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

func (sNode *PowerNode) SetValue(v int) error {
	sNode.value = v * 10
	return nil
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
	var sllnode *SLLNode
	fmt.Println(" --> ", sllnode.SetValue(4))

	n := createNode(5)
	switch cn := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SSLNode. Node is of Value: ", cn.GetValue())
	case *PowerNode:
		fmt.Println("Type is PowerNode. Node is of Value: ", cn.GetValue())
	}
	// var node Node
	// node = NewSLLNode()
	// node.SetValue(4)
	// fmt.Println("Node is of Value: ", node.GetValue())

	// node = NewPowerNode()
	// node.SetValue(5)
	// fmt.Println("Node is of Value: ", node.GetValue())

	// if n, ok := node.(*PowerNode); ok {
	// 	fmt.Println("This is a power node of value ", n.value)
	// }

	// fmt.Println("Powernode: ", node)
}

func createNode(v int) Node {
	sn := NewSLLNode()
	sn.SetValue(v)
	return sn
}
