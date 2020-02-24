package main

import "tgin/datastructure/code"

func main()  {
	t := code.List{}
	t.PushNode(7)
	t.ShowList()
	t.PushNode(1,2,43,3,5)
	t.ShowList()
	t.AddNode(9)
	t.ShowList()
	t.Insert(4,10)
	t.Remove(3)
	t.ShowList()
	t.ReverserList()
	t.ShowList()

}
