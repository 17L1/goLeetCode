/*
 * @Description:
 * @Author: lilongguang
 * @Date: 2022-04-08 10:34:13
 * @LastEditors: lilongguang
 * @LastEditTime: 2022-04-08 11:34:32
 */

package main

import "container/list"

type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := []*Node{root}
	for queue != nil {
		level := []int{}
		tmp := queue
		queue = nil
		for _, node := range tmp {
			level = append(level, node.Val)
			queue = append(queue, node.Children...)
		}
		ans = append(ans, level)
	}
	return ans

}

func levelOrder2(root *Node) [][]int {
	var ans [][]int
	if root == nil {
		return ans
	}
	queue := list.New()
	queue.PushBack(root)
	for queue.Len() > 0 {
		size := queue.Len()
		tmp := []int{}

		for i := 0; i < size; i++ {
			node := queue.Remove(queue.Front()).(*Node)
			tmp = append(tmp, node.Val)
			for j := 0; j < len(node.Children); j++ {
				queue.PushBack(node.Children[j])
			}

		}
		ans = append(ans, tmp)
	}
	return ans

}
