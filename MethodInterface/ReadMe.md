
if n, ok := node.(*PowerNode); ok {
    fmt.Println("This is a power node of value ", n.value)

node.(*PowerNode) returns a type variable (PowerNode) of the interface defined by node
this is why n.value is avaiable.

interface Stringer defineds String()
func (sNode *PowerNode) String() string {
	return "POWER NODE"
}
...
	fmt.Println("Powernode: ", node)
