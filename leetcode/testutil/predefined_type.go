package testutil

import (
	"fmt"
	"strconv"
	"strings"
)

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

	sb := &strings.Builder{}
	sb.WriteByte('[')
	for _, node := range nodes {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		if node != nil {
			sb.WriteString(strconv.Itoa(node.Val))
		} else {
			sb.WriteString("null")
		}
	}
	sb.WriteByte(']')
	return sb.String()
}

func GetParents(root *TreeNode) map[*TreeNode]*TreeNode {
	parents := map[*TreeNode]*TreeNode{}
	var dfs func(o, pa *TreeNode)
	dfs = func(o, pa *TreeNode) {
		if o == nil {
			return
		}
		parents[o] = pa
		dfs(o.Left, o)
		dfs(o.Right, o)
	}
	dfs(root, nil)
	return parents
}

func CountNodes(root *TreeNode) (cnt int) {
	var dfs func(*TreeNode)
	dfs = func(o *TreeNode) {
		if o == nil {
			return
		}
		cnt++
		dfs(o.Left)
		dfs(o.Right)
	}
	dfs(root)
	return
}

// 加权图见下面的 ToWeightedGraph
func ToGraph(root *TreeNode) [][]int {
	n := CountNodes(root)
	g := make([][]int, n)
	id := 0
	var build func(*TreeNode)
	build = func(o *TreeNode) {
		v := id
		if o.Left == nil && o.Right == nil {
			// do leaf ...

		}
		if o.Left != nil {
			id++
			g[v] = append(g[v], id)
			g[id] = append(g[id], v)
			build(o.Left)
		}
		if o.Right != nil {
			id++
			g[v] = append(g[v], id)
			g[id] = append(g[id], v)
			build(o.Right)
		}
	}
	build(root)
	return g
}

type DirEdge struct {
	to  int
	dir byte
}

func ToGraphWithDir(root *TreeNode) [][]DirEdge {
	const (
		left  = 'L'
		right = 'R'
		up    = 'U'
	)
	n := CountNodes(root)
	g := make([][]DirEdge, n)
	id := 0
	var build func(o *TreeNode)
	build = func(o *TreeNode) {
		v := id
		if o.Left == nil && o.Right == nil {
			// do leaf ...

		}
		if o.Left != nil {
			id++
			g[v] = append(g[v], DirEdge{id, left})
			g[id] = append(g[id], DirEdge{v, up})
			build(o.Left)
		}
		if o.Right != nil {
			id++
			g[v] = append(g[v], DirEdge{id, right})
			g[id] = append(g[id], DirEdge{v, up})
			build(o.Right)
		}
	}
	build(root)
	return g
}

type Edge struct{ to, wt int }

func ToWeightedGraph(root *TreeNode) [][]Edge {
	n := CountNodes(root)
	g := make([][]Edge, n)
	id := 0
	var build func(o *TreeNode)
	build = func(o *TreeNode) {
		v := id
		if o.Left == nil && o.Right == nil {
			// do leaf ...

		}
		if o.Left != nil {
			id++
			g[v] = append(g[v], Edge{id, o.Left.Val})
			g[id] = append(g[id], Edge{v, o.Left.Val})
			build(o.Left)
		}
		if o.Right != nil {
			id++
			g[v] = append(g[v], Edge{id, o.Right.Val})
			g[id] = append(g[id], Edge{v, o.Right.Val})
			build(o.Right)
		}
	}
	build(root)
	return g
}

//

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
	sb := &strings.Builder{}
	sb.WriteByte('[')
	for ; o != nil; o = o.Next {
		if sb.Len() > 1 {
			sb.WriteByte(',')
		}
		sb.WriteString(strconv.Itoa(o.Val))
	}
	sb.WriteByte(']')
	return sb.String()
}

func Nodes(head *ListNode) []*ListNode {
	nodes := []*ListNode{}
	for o := head; o != nil; o = o.Next {
		nodes = append(nodes, o)
	}
	return nodes
}

func Values(head *ListNode) (values []int) {
	for o := head; o != nil; o = o.Next {
		values = append(values, o.Val)
	}
	return values
}

func BuildListNodeFromInts(a []int) *ListNode {
	dummy := ListNode{}
	cur := &dummy
	for _, v := range a {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return dummy.Next
}

func ModifyNodes(head *ListNode, f func([]int) []int) *ListNode {
	a := Values(head)
	res := f(a)
	return BuildListNodeFromInts(res)
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
