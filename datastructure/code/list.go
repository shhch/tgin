package code

import (
	"fmt"
)

type Node struct {
	Data interface{}
	Next *Node
}

type List struct {
	Head *Node
}

func (list *List) IsEmpty() bool {
	if list.Head != nil {
		return false
	}
	return true
}

func (list *List) GetLen() int {
	if list.IsEmpty() {
		return 0
	}
	cur := list.Head
	count := 1
	for cur.Next != nil {
		count++
		cur = cur.Next
	}
	return count
}

func (list *List) AddNode(obj interface{}) {
	n := &Node{Data: obj}
	if list.IsEmpty() {
		list.Head = n
	} else {
		n.Next = list.Head
		list.Head = n
	}
}

func (list *List) PushNode(obj ...interface{}) {
	var cur *Node
	for _, v := range obj {
		if list.IsEmpty() {
			fmt.Println("???")
			list.Head = &Node{Data: v}
			continue
		}

		cur = list.Head
		for cur.Next != nil {
			cur = cur.Next
		}
		cur.Next = &Node{Data: v}
	}

}

func (list *List) ShowList() {
	cur := list.Head
	for cur != nil {
		fmt.Print(cur.Data, " ")
		cur = cur.Next
	}
	fmt.Print("\n")
}

func (list *List) Insert(index int, obj interface{}) {
	if index <= 0 {
		list.AddNode(obj)
	} else if index >= list.GetLen() {
		list.PushNode(obj)
	} else {
		cur := list.Head
		for i := 1; i < index; i++ {
			cur = cur.Next
		}
		n := &Node{Data: obj}
		n.Next = cur.Next
		cur.Next = n
	}
}

func (list *List) Remove(index int) interface{} {
	if index <= 0 {
		return list.RemoveHead()
	} else if index >= list.GetLen() {
		return list.Pop()
	} else {
		cur := list.Head
		for i := 1; i < index-1; i++ {
			cur = cur.Next
		}
		data := cur.Next.Data
		cur.Next = cur.Next.Next
		return data
	}
}

func (list *List) Pop() interface{} {
	if list.IsEmpty() {
		return nil
	}
	length := list.GetLen()
	if length == 1 {
		data := list.Head.Data
		list.Head = nil
		return data
	} else {
		pre := list.Head
		cur := pre.Next
		for cur.Next != nil {
			pre = pre.Next
			cur = pre.Next
		}
		data := pre.Next.Data
		pre.Next = nil
		return data
	}

}

func (list *List) RemoveHead() interface{} {
	if !list.IsEmpty() {
		list.Head = list.Head.Next
		return list.Head.Data
	}
	return nil
}

func (list *List) ReverserList() {
	if list.IsEmpty() || list.Head.Next == nil {
		return
	}
	cur := list.Head
	var pre *Node = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	list.Head = pre
}
