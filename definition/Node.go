package definition

type DoublyLinkedNode struct {
	Val   int
	Prev  *DoublyLinkedNode
	Next  *DoublyLinkedNode
	Child *DoublyLinkedNode
}
