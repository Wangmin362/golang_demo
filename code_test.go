package main

import (
	"fmt"
	"testing"
)

func sortedListToBST(head *ListNode) *TreeNode {
	var list []int
	for head != nil {
		list = append(list, head.Val)
		head = head.Next
	}

	var build func(arr []int, begin, end int) *TreeNode

	build = func(arr []int, begin, end int) *TreeNode {
		if begin > end {
			return nil
		}
		mid := begin + (end-begin)>>1

		root := &TreeNode{Val: arr[mid]}
		root.Left = build(arr, begin, mid-1)
		root.Right = build(arr, mid+1, end)

		return root
	}
	return build(list, 0, len(list)-1)
}

func TestCode(t *testing.T) {
	head := &ListNode{Val: -10, Next: &ListNode{Val: -3, Next: &ListNode{Val: 0, Next: &ListNode{Val: 5, Next: &ListNode{Val: 9}}}}}
	root := sortedListToBST(head)
	fmt.Println(root)

}

type ListNode struct {
	Val  int
	Next *ListNode
}

// TreeNode 二叉树定义
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Node N叉树定义
type Node struct {
	Val      int
	Children []*Node
}
