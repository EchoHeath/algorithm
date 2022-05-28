package doubleLinkedList

import (
	"testing"
)

//双向链表Test
func TestCreateList(t *testing.T) {
	capacity := 10
	l := NewDoubleLinkedListService(capacity)
	nodes := []*Node{}
	for i := 1; i <= capacity; i++ {
		node := &Node{
			Key:   i,
			Value: i,
		}
		nodes = append(nodes, node)
	}

	l.Append(nodes[0])
	l.Print()
	l.Append(nodes[1])
	l.Print()
	l.Pop()
	l.Print()
	l.Append(nodes[2])
	l.Print()
	l.AppendFront(nodes[3])
	l.Print()
	l.Append(nodes[4])
	l.Print()
	l.Remove(nodes[2])
	l.Print()
	l.Remove(nil)
	l.Print()
}
