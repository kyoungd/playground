package main

import (
	"fmt"
)

type Node interface {
	SetValue(v int)
	GetValue() int
}

type SLLNode struct {
	next         *SLLNode
	value        int
	sNodeMessage string
}

func (sNode *SLLNode) SetValue(v int) {
	sNode.value = v
}

func (sNode *SLLNode) GetValue() int {
	return sNode.value
}

func NewSLLNode() *SLLNode {
	return &SLLNode{sNodeMessage: "This is a message from the normal Node"}
}

type PowerNode struct {
	next         *SLLNode
	value        int
	pNodeMessage string
}

func (sNode *PowerNode) SetValue(v int) {
	sNode.value = v * 10
}

func (sNode *PowerNode) GetValue() int {
	return sNode.value
}

func NewPowerNode() *PowerNode {
	return &PowerNode{pNodeMessage: "This is a message from the power Node"}
}

func main() {
	n := createNode(5)
	switch concreten := n.(type) {
	case *SLLNode:
		fmt.Println("Type is SSLNode, message", concreten.sNodeMessage)
	case *PowerNode:
		fmt.Println("type in PowerNode, message: ", concreten.pNodeMessage)
	}
}

func createNode(v int) Node {
	sn := NewSLLNode()
	sn.SetValue(v)
	return sn
}

// type SingleLinkedList struct {
// 	head *SLLNode
// 	tail *SLLNode
// }

// func newSingleLinkedList() *SingleLinkedList {
// 	return new(SingleLinkedList)
// }

// func (list *SingleLinkedList) Add(v int) {
// 	newNode := SLLNode{value: v}
// 	if list.head == nil {
// 		list.head = &newNode
// 	} else if list.tail == list.head {
// 		list.head.next = &newNode
// 	} else if list.tail != nil {
// 		list.tail.next = &newNode
// 	}
// 	list.tail = &newNode
// }

// func (list *SingleLinkedList) String() string {
// 	s := ""
// 	for n := list.head; n != nil; n = n.next {
// 		s += fmt.Sprintf(" {%d}", n.GetValue())
// 	}
// 	return s
// }

// func main() {
// 	list := newSingleLinkedList()
// 	list.Add(3)
// 	list.Add(4)
// 	list.Add(5)
// 	list.Add(6)
// 	fmt.Println("Hello, playground ", list)
// }
