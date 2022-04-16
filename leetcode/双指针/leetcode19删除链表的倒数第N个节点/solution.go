/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-16 14:03:46
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-16 14:16:35
 */
package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tummy := &ListNode{
		Val:  0,
		Next: nil,
	}
	tummy.Next = head
	first := tummy
	second := tummy
	//快指针先往前走N步
	for second.Next != nil && n > 0 {
		second = second.Next
		n--
	}
	if second == nil {
		return tummy.Next
	}
	for second.Next != nil {
		first = first.Next
		second = second.Next
	}
	first.Next = first.Next.Next
	return tummy.Next
}
