
myType := n.(type) myType is the type struct variable.

sNode := &SSLNode{value:15}  // Make sure you reference the type with &.

The receiver of a method is allowed to be nil.  We can test for nil or uninstantiated condition inside the method.  Neat.

func (sNode *SLLNode) SetValue(v int) error {
	if sNode == nil {
		return errors.New("Node is not instantiated.  It is nil.")
	}
	sNode.value = v
	return nil
}

func main() {
	var sllnode *SLLNode
	fmt.Println(sllnode.SetValue(4))
}

This works.
