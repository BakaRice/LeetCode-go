package main

/*
@Time : 2020/10/15 21:03
@Author : DELL ricemarch@foxmail.com
@tips:https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/solution/
*/
func connect_2(root *Node) *Node {
	start := root
	for start != nil {
		var nextStart, last *Node
		handle := func(cur *Node) {
			if cur == nil {
				return
			}
			if nextStart == nil {
				nextStart = cur
			}
			if last != nil {
				last.Next = cur
			}
			last = cur
		}
		for p := start; p != nil; p = p.Next {
			handle(p.Left)
			handle(p.Right)
		}
		start = nextStart
	}
	return root
}
