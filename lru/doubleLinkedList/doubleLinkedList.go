package doubleLinkedList

import (
	"fmt"
)

type Node struct {
	Key   interface{} //缓存key
	Value interface{} //缓存val
	Prev  *Node       //上一个节点
	Next  *Node       //下一个节点
}

type DoubleLinkedList struct {
	Head     *Node //头结点
	Tail     *Node //尾结点
	Capacity int   //总容量
	Size     int   //当前使用的容量
}

type ListInterface interface {
	AppendFront(node *Node) *Node //从头部添加
	Append(node *Node) *Node      //从尾部添加
	Remove(node *Node) *Node      //删除节点
	Pop() *Node                   //删除头结点
	RemoveTail() *Node            //删除尾结点
	Print()                       //打印
}

//NewDoubleLinkedListService get a doubleLinkedList.
func NewDoubleLinkedListService(capacity int) *DoubleLinkedList {
	//获取双向链表
	return &DoubleLinkedList{
		Capacity: capacity,
	}
}

//AppendFront 从头部添加
func (d *DoubleLinkedList) AppendFront(node *Node) *Node {
	if d.Head == nil {
		d.Head = node
		d.Tail = node
	} else {
		node.Next = d.Head
		d.Head.Prev = node
		d.Head = node
		node.Prev = nil
	}
	d.Size++
	return node
}

//Append 从尾部添加
func (d *DoubleLinkedList) Append(node *Node) *Node {
	if d.Tail == nil {
		d.Head, d.Tail = node, node
	} else {
		d.Tail.Next = node
		node.Prev = d.Tail
		d.Tail = node
	}
	d.Size++
	return node
}

//Remove 删除节点
func (d *DoubleLinkedList) Remove(node *Node) *Node {
	if node == nil || node == d.Tail {
		return d.RemoveTail()
	} else if node == d.Head {
		return d.Pop()
	} else {
		node.Prev.Next = node.Next
		node.Next.Prev = node.Prev
		d.Size--
	}
	return node
}

//Pop 删除头结点
func (d *DoubleLinkedList) Pop() *Node {
	if d.Head == nil {
		return nil
	}
	index := d.Head
	if index.Next != nil {
		d.Head = index.Next
		d.Head.Prev = nil
		d.Size--
	} else {
		d.Head, d.Tail = nil, nil
		d.Size = 0
	}
	return index
}

//RemoveTail 删除尾结点
func (d *DoubleLinkedList) RemoveTail() *Node {
	if d.Tail == nil {
		return nil
	}

	index := d.Tail
	if index.Prev != nil {
		d.Tail = index.Prev
		d.Tail.Next = nil
		d.Size--
	} else {
		d.Head, d.Tail = nil, nil
		d.Size = 0
	}

	return index
}

//Print 打印链表
func (d *DoubleLinkedList) Print() {
	index := d.Head
	line := ""
	for index != nil {
		line += fmt.Sprintf("{%v, %v}==>", index.Key, index.Value)
		index = index.Next
	}
	fmt.Println(line)
}
