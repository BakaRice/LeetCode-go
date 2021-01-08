package main

/*
@Time : 2020/10/15 20:55
@Author : DELL ricemarch@foxmail.com
@tips:https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node/
*/

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	//每次循环从该层的最左侧节点开始
	for leftmost := root; leftmost.Left != nil; leftmost = leftmost.Left {
		//通过Next层序遍历这一层节点，为下一层的节点更新Next指针
		for node := leftmost; node != nil; node = node.Next {
			//左节点指向右节点
			node.Left.Next = node.Right
			//右节点指向下一个左节点
			if node.Next != nil {
				node.Right.Next = node.Next.Left
			}
		}
	}
	//返回根节点
	return root
}
