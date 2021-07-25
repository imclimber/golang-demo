package main

import "fmt"

func main() {
	var head *listNode
	head = createList(head)
	printList(head)

	head = removeNthFromEnd(head, 2)
	printListWithDummy(head)
	// printList(head)
}

type listNode struct {
	Val  int
	Next *listNode
}

func removeNthFromEnd(head *listNode, n int) *listNode {
	dummy := &listNode{}
	dummy.Next = head

	fast := dummy
	for i := 0; i < n && fast.Next != nil; i++ {
		fast = fast.Next
	}
	fmt.Println("fast: ", fast.Val)

	slow := dummy
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	fmt.Println("slow: ", slow.Val)

	del := slow.Next
	slow.Next = del.Next
	del.Next = nil

	// return dummy.Next // 使用 printList 打印，即不使用虚拟头结点
	return dummy // 使用 printListWithDummy 打印，使用虚拟头结点
}

func createList(head *listNode) *listNode {
	if head == nil {
		head = &listNode{
			Val: 1,
		}
	}

	cur := head
	next := head

	for i := 2; i < 1; i++ {
		next = &listNode{
			Val: i,
		}

		cur.Next = next
		cur = next
	}

	return head
}

func printList(head *listNode) *listNode {
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

func printListWithDummy(head *listNode) *listNode {
	if head == nil || head.Next == nil {
		fmt.Println("val: ", "nil")
		return head
	}

	cur := head.Next

	for cur != nil {
		fmt.Println("val: ", cur.Val)
		cur = cur.Next
	}

	return head
}
