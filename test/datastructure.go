package main

import (
	"tgin/datastructure"
)

func main()  {
	t := datastructure.List{}
	t.PushNode(1)
	t.PushNode(2)
	t.PushNode(3)
	t.PushNode(4)
	t.PushNode(5)
	t.AddNode(9)
	t.ShowList()
	t.Insert(4,10)
	t.Remove(3)
	t.ShowList()
	t.ReverserList()
	t.ShowList()

}
