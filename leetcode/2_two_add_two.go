package main

import (
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)

	l1.Val = 9
	l2.Val = 9
	// l1 := l1.Next
	// log.Println("%l1", l1.Next)
	// log.Println("%l1", l1)
	// if l1 == nil {
	// 	l1 = &ListNode{Val: 9}  // bug: 原来 l1 指向 l1.Next，现在 l1 指向新的节点了
	// }
	// log.Println("%l1", l1.Next)
	// log.Println("%l1", l1)
	// for l1 != nil {
	// 	log.Println("l1: ", l1.Val)
	// 	l1 = l1.Next
	// }
	l1.Next = &ListNode{Val: 9}

	res := addTwoNumbers(l1, l2)
	for res != nil {
		log.Println("res: ", res.Val)
		res = res.Next
	}
}

// 写法上简化了，但是效率却不是很高
// 执行用时： 16 ms , 在所有 Go 提交中击败了 38.01% 的用户
// 内存消耗： 4.7 MB , 在所有 Go 提交中击败了 95.98% 的用户

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := l1

	lastCount := 0
	oneVal := 0
	twoVal := 0
	for l1 != nil || l2 != nil || lastCount != 0 {
		oneVal = 0
		twoVal = 0
		if l1 != nil {
			oneVal = l1.Val
		}
		if l2 != nil {
			twoVal = l2.Val
		}

		l1.Val = (lastCount + oneVal + twoVal) % 10
		lastCount = (lastCount + oneVal + twoVal) / 10

		if (l1.Next == nil && lastCount != 0) || (l1.Next == nil && l2 != nil && l2.Next != nil) {
			l1.Next = &ListNode{}
		}

		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}

	return head
}
