package main

import (
	"fmt"
)

func main() {
	var head *ListNode
	head = createList(head)
	// printList(head)

	res := hasCycle_V2(head)
	fmt.Println("res: ", res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	fast := head
	slow := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if fast == slow {
			fmt.Println("slow, fast: ", slow.Val, fast.Val)
			fmt.Printf("slow point[%p], fast point[%p]\n ", slow, fast)
			return true
		}
	}

	return false
}

func hasCycle_V2(head *ListNode) bool {
	nodeMap := make(map[string]bool)

	cur := head
	for cur != nil && cur.Next != nil {
		p := fmt.Sprintln("%p", cur)
		fmt.Println("p:", p)
		_, ok := nodeMap[p]
		if ok {
			return true
		} else {
			nodeMap[p] = true
		}

		cur = cur.Next
	}

	return false
}

func createList(head *ListNode) *ListNode {
	if head == nil {
		head = &ListNode{
			Val: 1,
		}
	}

	cur := head
	next := head

	for i := 2; i < 7; i++ {
		next = &ListNode{
			Val: i,
		}

		cur.Next = next
		cur = next
	}

	cur.Next = head

	return head
}

func printList(head *ListNode) *ListNode {
	if head == nil {
		fmt.Println("val: ", "nil")
		return head
	}

	cur := head

	for cur != nil {
		fmt.Println("val: ", cur.Val)
		cur = cur.Next
	}

	return head
}
