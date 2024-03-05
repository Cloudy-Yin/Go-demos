package main

import "fmt"

type List struct {
	Val  int
	Next *List
}

func rmListNum(head *List) *List {
	if head == nil || head.Next == nil {
		return nil
	}

	cur := head.Next
	var pre *List = nil
	for cur != nil {
		pre = cur
		for pre.Next != nil {
			if cur.Val == pre.Next.Val {
				pre.Next = pre.Next.Next
			} else {
				pre = pre.Next
			}
		}
		cur = cur.Next
	}

	return head
}

func print(head *List) {
	cur := head
	for cur != nil {
		fmt.Println(cur.Val)
		cur = cur.Next
	}
}

func main() {

	node5 := &List{Val: 5, Next: nil}
	node4 := &List{Val: 2, Next: node5}
	node3 := &List{Val: 2, Next: node4}
	node2 := &List{Val: 2, Next: node3}
	node1 := &List{Val: 1, Next: node2}
	print(rmListNum(node1))
	//t1 := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}}
	//fmt.Println(mergeTwoList(t1[0], t1[1]))
}
