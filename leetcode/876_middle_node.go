package main

import "fmt"

func main() {
	var head *ListNode
	head = createList(head)
	printList(head)

	head = middleNode(head)
	// printListWithDummy(head)
	// printList(head)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	fast := head
	slow := head

	for fast != nil {
		if fast.Next == nil {
			break
		}
		if fast.Next != nil && fast.Next.Next == nil {
			if slow != nil && slow.Next != nil {
				slow = slow.Next
			}
			break
		}

		fast = fast.Next.Next
		slow = slow.Next
	}

	fmt.Println("slow: ", slow.Val)
	return slow
}

func createList(head *ListNode) *ListNode {
	if head == nil {
		head = &ListNode{
			Val: 1,
		}
	}

	cur := head
	next := head

	for i := 2; i < 8; i++ {
		next = &ListNode{
			Val: i,
		}

		cur.Next = next
		cur = next
	}

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

func printListWithDummy(head *ListNode) *ListNode {
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
