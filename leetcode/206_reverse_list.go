package main

import "fmt"

func main() {
	var head *listNode
	head = createList(head)
	printList(head)
	head = reverseList_recursion(head)
	printList(head)
}

//  * Definition for singly-linked list.
type listNode struct {
	Val  int
	Next *listNode
}

func createList(head *listNode) *listNode {
	if head == nil {
		head = &listNode{
			Val: 1,
		}
	}
	cur := head
	next := head

	for i := 2; i < 6; i++ {
		next = &listNode{
			Val: i,
		}

		cur.Next = next
		cur = next
	}

	return head
}

func printList(head *listNode) *listNode {
	cur := head

	for cur != nil {
		fmt.Println("val: ", cur.Val)
		cur = cur.Next
	}

	return head
}

func reverseList(head *listNode) *listNode {
	var pre *listNode
	cur := head

	for cur != nil {
		after := cur.Next

		cur.Next = pre

		pre = cur
		cur = after
	}

	return pre
}

func reverseList_recursion(head *listNode) *listNode {

	// 递归终止条件
	if head == nil || head.Next == nil {
		return head
	}

	// 核心逻辑
	next := head.Next
	head.Next = nil
	newhead := reverseList_recursion(next)
	next.Next = head

	return newhead
}
