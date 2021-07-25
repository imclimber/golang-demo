package main

import (
	"fmt"
)

func main() {
	heada, headb := createIntersectionList()
	printList(heada)
	printList(headb)
	res := getIntersectionNode_V2(headb, heada)
	fmt.Println("res: ", res)
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB
	lena := 0
	lenb := 0

	for pa != nil && pa.Next != nil {
		pa = pa.Next
		lena++
	}

	for pb != nil && pb.Next != nil {
		pb = pb.Next
		lenb++
	}

	pa = headA
	pb = headB
	if lena > lenb {
		pa, pb = pb, pa
		lena, lenb = lenb, lena
	}
	distance := lenb - lena
	for i := 0; i < distance; i++ {
		pb = pb.Next
	}

	// nil == nil 为 true，依然可以跳出循环
	for pa != pb {
		pa = pa.Next
		pb = pb.Next
	}
	// fmt.Println(pa == nil)
	// fmt.Println("same point: ", pa.Val)

	return pa
}

func getIntersectionNode_V2(headA, headB *ListNode) *ListNode {
	pa := headA
	pb := headB

	for pa != pb {
		if pa != nil {
			pa = pa.Next
		} else {
			pa = headB
		}

		if pb != nil {
			pb = pb.Next
		} else {
			pb = headA
		}
	}

	return pa
}

func createIntersectionList() (heada *ListNode, headb *ListNode) {
	if heada == nil {
		heada = &ListNode{
			Val: 1,
		}
	}
	if headb == nil {
		headb = &ListNode{
			Val: 10,
		}
	}

	cur := heada
	next := heada
	psame := heada
	for i := 2; i < 7; i++ {
		next = &ListNode{
			Val: i,
		}

		if i == 3 {
			psame = next
		}

		cur.Next = next
		cur = next
	}

	// headb
	cur = headb
	next = headb

	for i := 11; i < 17; i++ {
		next = &ListNode{
			Val: i,
		}

		if i == 14 {
			cur.Next = psame
			break
		}

		cur.Next = next
		cur = next
	}

	return heada, headb
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
