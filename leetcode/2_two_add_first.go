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

// 看了题解后，发现想得有点复杂了，但是效率确实很高
// 第一次运行结果
// 执行用时： 12 ms , 在所有 Go 提交中击败了 70.96% 的用户
// 内存消耗： 4.6 MB , 在所有 Go 提交中击败了 96.73% 的用户

// 某一次运行结果
// 执行用时：8 ms, 在所有 Go 提交中击败了 92.24% 的用户
// 内存消耗：4.7 MB, 在所有 Go 提交中击败了94.72% 的用户

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil || l2 == nil {
		return nil
	}

	head := l1

	lastCount := 0
	oneVal := 0
	twoVal := 0
	for l1 != nil || l2 != nil {
		oneVal = 0
		twoVal = 0
		if l1 != nil {
			oneVal = l1.Val
		}
		if l2 != nil {
			twoVal = l2.Val
		}

		// l2 存在下一位，l1 才能扩展下一位
		if (l1 != nil && l1.Next == nil) && (l2 != nil && l2.Next != nil) {
			l1.Next = &ListNode{}
		}
		l1.Val = lastCount + (oneVal+twoVal)%10
		lastCount = (oneVal + twoVal) / 10

		// 修正上一次进 1，本次和为 9
		if l1.Val == 10 {
			l1.Val = 0
			lastCount = 1
		}

		// 修正：最后一位超出 10，必须补充一位
		if (l1 != nil && l1.Next == nil) && (l2 == nil || l2.Next == nil) && lastCount == 1 {
			l1.Next = &ListNode{Val: 1}
			return head
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
