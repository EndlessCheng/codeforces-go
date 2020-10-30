package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

// 注意确认 Val 的类型是否和题目一致（一般都是 int）
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// https://support.leetcode.com/hc/en-us/articles/360011883654-What-does-1-null-2-3-mean-in-binary-tree-representation
func buildTreeNode(rawArray string) (root *TreeNode, err error) {
	if len(rawArray) < 2 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, fmt.Errorf("invalid test data %s", rawArray)
	}

	rawArray = rawArray[1 : len(rawArray)-1]
	if rawArray == "" {
		return
	}
	splits := strings.Split(rawArray, ",")

	nodes := make([]*TreeNode, len(splits))
	for i, s := range splits {
		s = strings.TrimSpace(s)
		if s != "null" {
			nodes[i] = &TreeNode{}
			var er error
			nodes[i].Val, er = strconv.Atoi(s)
			if er != nil {
				return nil, er
			}
			//nodes[i].Val = s[1 : len(s)-1]
		}
	}
	kids := nodes

	root, kids = kids[0], kids[1:]
	for _, node := range nodes {
		if node != nil {
			if len(kids) > 0 {
				node.Left, kids = kids[0], kids[1:]
			}
			if len(kids) > 0 {
				node.Right, kids = kids[0], kids[1:]
			}
		}
	}
	return
}

func (o *TreeNode) toRawString() string {
	nodes := []*TreeNode{}
	queue := []*TreeNode{o}
	for len(queue) > 0 {
		o, queue = queue[0], queue[1:]
		nodes = append(nodes, o)
		if o != nil {
			queue = append(queue, o.Left, o.Right)
		}
	}

	for len(nodes) > 0 && nodes[len(nodes)-1] == nil {
		nodes = nodes[:len(nodes)-1]
	}

	s := "["
	for _, node := range nodes {
		if len(s) > 1 {
			s += ","
		}
		if node != nil {
			s += strconv.Itoa(node.Val)
			//s += `"` + node.Val + `"`
		} else {
			s += "null"
		}
	}
	s += "]"
	return s
}

func (o *TreeNode) CountNodes() int {
	cnto := 0
	var _f func(o *TreeNode)
	_f = func(o *TreeNode) {
		if o == nil {
			return
		}
		cnto++
		_f(o.Left)
		_f(o.Right)
	}
	_f(o)
	return cnto
}

// 加权图见下面的 ToWeightedGraph
func (o *TreeNode) ToGraph() {
	n := o.CountNodes()
	g := make([][]int, n)
	cnt := 0
	var build func(o *TreeNode)
	build = func(o *TreeNode) {
		v := cnt
		if o.Left == nil && o.Right == nil {
			// do leaf ...

		}
		if o.Left != nil {
			cnt++
			g[v] = append(g[v], cnt)
			g[cnt] = append(g[cnt], v)
			build(o.Left)
		}
		if o.Right != nil {
			cnt++
			g[v] = append(g[v], cnt)
			g[cnt] = append(g[cnt], v)
			build(o.Right)
		}
	}
	build(o)
}

func (o *TreeNode) ToWeightedGraph() {
	type edge struct{ to, weight int }
	n := o.CountNodes()
	g := make([][]edge, n)
	cnt := 0
	var build func(o *TreeNode)
	build = func(o *TreeNode) {
		v := cnt
		if o.Left == nil && o.Right == nil {
			// do leaf ...

		}
		if o.Left != nil {
			cnt++
			g[v] = append(g[v], edge{cnt, o.Left.Val})
			g[cnt] = append(g[cnt], edge{v, o.Left.Val})
			build(o.Left)
		}
		if o.Right != nil {
			cnt++
			g[v] = append(g[v], edge{cnt, o.Right.Val})
			g[cnt] = append(g[cnt], edge{v, o.Right.Val})
			build(o.Right)
		}
	}
	build(o)
}

//

// 注意确认 Val 的类型是否和题目一致（一般都是 int）
type ListNode struct {
	Val  int
	Next *ListNode
}

func buildListNode(rawArray string) (head *ListNode, err error) {
	if len(rawArray) < 2 || rawArray[0] != '[' || rawArray[len(rawArray)-1] != ']' {
		return nil, fmt.Errorf("invalid test data %s", rawArray)
	}

	rawArray = rawArray[1 : len(rawArray)-1]
	if rawArray == "" {
		return
	}
	splits := strings.Split(rawArray, ",")

	nodes := make([]*ListNode, len(splits))
	for i, s := range splits {
		s = strings.TrimSpace(s)
		nodes[i] = &ListNode{}
		var er error
		nodes[i].Val, er = strconv.Atoi(s)
		if er != nil {
			return nil, er
		}
		//nodes[i].Val = s[1 : len(s)-1]
	}

	head = nodes[0]
	for i := 0; i < len(nodes)-1; i++ {
		nodes[i].Next = nodes[i+1]
	}
	return
}

func (o *ListNode) toRawString() string {
	s := "["
	for ; o != nil; o = o.Next {
		if len(s) > 1 {
			s += ","
		}
		s += strconv.Itoa(o.Val)
		//s += `"` + o.Val + `"`
	}
	s += "]"
	return s
}

func (o *ListNode) Values() []int {
	vals := []int{}
	for ; o != nil; o = o.Next {
		vals = append(vals, o.Val)
	}
	return vals
}

func (o *ListNode) Nodes() []*ListNode {
	nodes := []*ListNode{}
	for ; o != nil; o = o.Next {
		nodes = append(nodes, o)
	}
	return nodes
}

//

func MustBuildTreeNode(rawArray string) *TreeNode {
	root, err := buildTreeNode(rawArray)
	if err != nil {
		panic(err)
	}
	return root
}

func MustBuildListNode(rawArray string) *ListNode {
	head, err := buildListNode(rawArray)
	if err != nil {
		panic(err)
	}
	return head
}

//

// https://leetcode-cn.com/problems/populating-next-right-pointers-in-each-node-ii/
type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}
