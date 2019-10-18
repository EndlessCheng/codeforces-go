package main

import (
	"bufio"
	. "fmt"
	"io"
)

type keyType int
type valueType int

var zeroValue valueType = 0

type Tree struct {
	Root       *Node
	size       int
	comparator func(a, b keyType) int
}

type color bool

const black, red color = true, false

type Node struct {
	Left   *Node
	Right  *Node
	Parent *Node
	Key    keyType
	Value  valueType
	color  color
}

func newRBTree() *Tree {
	return &Tree{comparator: func(a, b keyType) int {
		if a < b {
			return -1
		}
		if a > b {
			return 1
		}
		return 0
	}}
}

func (t *Tree) Insert(key keyType, value valueType) {
	var insertedNode *Node
	if t.Root == nil {
		t.Root = &Node{Key: key, Value: value, color: red}
		insertedNode = t.Root
	} else {
		node := t.Root
	loop:
		for {
			compare := t.comparator(key, node.Key)
			switch {
			case compare < 0:
				if node.Left == nil {
					node.Left = &Node{Key: key, Value: value, color: red}
					insertedNode = node.Left
					break loop
				} else {
					node = node.Left
				}
			case compare > 0:
				if node.Right == nil {
					node.Right = &Node{Key: key, Value: value, color: red}
					insertedNode = node.Right
					break loop
				} else {
					node = node.Right
				}
			default:
				node.Key = key
				node.Value = value
				return
			}
		}
		insertedNode.Parent = node
	}
	t.insertCase1(insertedNode)
	t.size++
}

func (t *Tree) Lookup(key keyType) *Node {
	for o := t.Root; o != nil; {
		compare := t.comparator(key, o.Key)
		switch {
		case compare < 0:
			o = o.Left
		case compare > 0:
			o = o.Right
		default:
			return o
		}
	}
	return nil
}

func (t *Tree) Erase(key keyType) {
	node := t.Lookup(key)
	if node == nil {
		return
	}
	if node.Left != nil && node.Right != nil {
		pred := node.Left.maximumNode()
		node.Key = pred.Key
		node.Value = pred.Value
		node = pred
	}
	if node.Left == nil || node.Right == nil {
		var child *Node
		if node.Right == nil {
			child = node.Left
		} else {
			child = node.Right
		}
		if node.color == black {
			node.color = nodeColor(child)
			t.deleteCase1(node)
		}
		t.replaceNode(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
	t.size--
}

func (t *Tree) MultiInsert(key keyType) {
	if o := t.Lookup(key); o != nil {
		o.Value++
	} else {
		t.Insert(key, 1)
	}
}

func (t *Tree) MultiErase(key keyType) {
	if o := t.Lookup(key); o != nil {
		o.Value--
		if o.Value == 0 {
			t.Erase(key)
		}
	}
}

// Floor Finds floor node of the input key, return the floor node or nil if no floor is found.
//
// Floor node is defined as the largest node that is smaller than or equal to the given node.
// A floor node may not be found, either because the tree is empty, or because
// all nodes in the tree are larger than the given node.
func (t *Tree) Floor(key keyType) (floor *Node) {
	for o := t.Root; o != nil; {
		compare := t.comparator(key, o.Key)
		switch {
		case compare < 0:
			o = o.Left
		case compare > 0:
			floor = o
			o = o.Right
		default:
			return o
		}
	}
	return
}

func (node *Node) grandparent() *Node {
	if node != nil && node.Parent != nil {
		return node.Parent.Parent
	}
	return nil
}

func (node *Node) uncle() *Node {
	if node == nil || node.Parent == nil || node.Parent.Parent == nil {
		return nil
	}
	return node.Parent.sibling()
}

func (node *Node) sibling() *Node {
	if node == nil || node.Parent == nil {
		return nil
	}
	if node == node.Parent.Left {
		return node.Parent.Right
	}
	return node.Parent.Left
}

func (t *Tree) rotateLeft(node *Node) {
	right := node.Right
	t.replaceNode(node, right)
	node.Right = right.Left
	if right.Left != nil {
		right.Left.Parent = node
	}
	right.Left = node
	node.Parent = right
}

func (t *Tree) rotateRight(node *Node) {
	left := node.Left
	t.replaceNode(node, left)
	node.Left = left.Right
	if left.Right != nil {
		left.Right.Parent = node
	}
	left.Right = node
	node.Parent = left
}

func (t *Tree) replaceNode(old *Node, new *Node) {
	if old.Parent == nil {
		t.Root = new
	} else {
		if old == old.Parent.Left {
			old.Parent.Left = new
		} else {
			old.Parent.Right = new
		}
	}
	if new != nil {
		new.Parent = old.Parent
	}
}

func (t *Tree) insertCase1(node *Node) {
	if node.Parent == nil {
		node.color = black
	} else {
		t.insertCase2(node)
	}
}

func (t *Tree) insertCase2(node *Node) {
	if nodeColor(node.Parent) == black {
		return
	}
	t.insertCase3(node)
}

func (t *Tree) insertCase3(node *Node) {
	uncle := node.uncle()
	if nodeColor(uncle) == red {
		node.Parent.color = black
		uncle.color = black
		node.grandparent().color = red
		t.insertCase1(node.grandparent())
	} else {
		t.insertCase4(node)
	}
}

func (t *Tree) insertCase4(node *Node) {
	grandparent := node.grandparent()
	if node == node.Parent.Right && node.Parent == grandparent.Left {
		t.rotateLeft(node.Parent)
		node = node.Left
	} else if node == node.Parent.Left && node.Parent == grandparent.Right {
		t.rotateRight(node.Parent)
		node = node.Right
	}
	t.insertCase5(node)
}

func (t *Tree) insertCase5(node *Node) {
	node.Parent.color = black
	grandparent := node.grandparent()
	grandparent.color = red
	if node == node.Parent.Left && node.Parent == grandparent.Left {
		t.rotateRight(grandparent)
	} else if node == node.Parent.Right && node.Parent == grandparent.Right {
		t.rotateLeft(grandparent)
	}
}

func (node *Node) maximumNode() *Node {
	if node == nil {
		return nil
	}
	for node.Right != nil {
		node = node.Right
	}
	return node
}

func (t *Tree) deleteCase1(node *Node) {
	if node.Parent == nil {
		return
	}
	t.deleteCase2(node)
}

func (t *Tree) deleteCase2(node *Node) {
	sibling := node.sibling()
	if nodeColor(sibling) == red {
		node.Parent.color = red
		sibling.color = black
		if node == node.Parent.Left {
			t.rotateLeft(node.Parent)
		} else {
			t.rotateRight(node.Parent)
		}
	}
	t.deleteCase3(node)
}

func (t *Tree) deleteCase3(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == black &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		t.deleteCase1(node.Parent)
	} else {
		t.deleteCase4(node)
	}
}

func (t *Tree) deleteCase4(node *Node) {
	sibling := node.sibling()
	if nodeColor(node.Parent) == red &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == black &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		node.Parent.color = black
	} else {
		t.deleteCase5(node)
	}
}

func (t *Tree) deleteCase5(node *Node) {
	sibling := node.sibling()
	if node == node.Parent.Left &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Left) == red &&
		nodeColor(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		t.rotateRight(sibling)
	} else if node == node.Parent.Right &&
		nodeColor(sibling) == black &&
		nodeColor(sibling.Right) == red &&
		nodeColor(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		t.rotateLeft(sibling)
	}
	t.deleteCase6(node)
}

func (t *Tree) deleteCase6(node *Node) {
	sibling := node.sibling()
	sibling.color = nodeColor(node.Parent)
	node.Parent.color = black
	if node == node.Parent.Left && nodeColor(sibling.Right) == red {
		sibling.Right.color = black
		t.rotateLeft(node.Parent)
	} else if nodeColor(sibling.Left) == red {
		sibling.Left.color = black
		t.rotateRight(node.Parent)
	}
}

func nodeColor(node *Node) color {
	if node == nil {
		return black
	}
	return node.color
}

type position byte

const begin, between, end position = 0, 1, 2

// Iterator holding the iterator's state
type Iterator struct {
	tree     *Tree
	node     *Node
	position position
}

// Iterator returns a stateful iterator whose elements are key/value pairs.
func (t *Tree) NewIterator(node *Node) *Iterator {
	return &Iterator{tree: t, node: node, position: between}
}

// Min returns the left-most (min) node or nil if tree is empty.
func (t *Tree) Min() (min *Node) {
	for o := t.Root; o != nil; o = o.Left {
		min = o
	}
	return
}

// Max returns the right-most (max) node or nil if tree is empty.
func (t *Tree) Max() (max *Node) {
	for o := t.Root; o != nil; o = o.Right {
		max = o
	}
	return
}

// Next moves the iterator to the next element and returns itself.
// Modifies the state of the iterator.
func (it *Iterator) Next() *Iterator {
	if it.position == end {
		goto end
	}
	if it.position == begin {
		left := it.tree.Min()
		if left == nil {
			goto end
		}
		it.node = left
		goto between
	}
	if it.node.Right != nil {
		it.node = it.node.Right
		for it.node.Left != nil {
			it.node = it.node.Left
		}
		goto between
	}
	if it.node.Parent != nil {
		node := it.node
		for it.node.Parent != nil {
			it.node = it.node.Parent
			if it.tree.comparator(node.Key, it.node.Key) <= 0 {
				goto between
			}
		}
	}

end:
	it.node = nil
	it.position = end
	return it

between:
	it.position = between
	return it
}

// github.com/EndlessCheng/codeforces-go
func Sol527C(reader io.Reader, writer io.Writer) {
	cut := func(t, mt *Tree, mid keyType) {
		it := t.NewIterator(t.Floor(mid))
		l := it.node.Key
		r := it.Next().node.Key
		mt.MultiErase(r - l)
		mt.MultiInsert(r - mid)
		mt.MultiInsert(mid - l)
		t.Insert(mid, 1)
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	w, h, mw, mh := newRBTree(), newRBTree(), newRBTree(), newRBTree()
	var w0, h0 keyType
	var n int
	Fscan(in, &w0, &h0, &n)
	w.Insert(0, 1)
	w.Insert(w0, 1)
	mw.MultiInsert(w0)
	h.Insert(0, 1)
	h.Insert(h0, 1)
	mh.MultiInsert(h0)
	for ; n > 0; n-- {
		var op string
		var mid keyType
		Fscan(in, &op, &mid)
		if op[0] == 'V' {
			cut(w, mw, mid)
		} else {
			cut(h, mh, mid)
		}
		Fprintln(out, int64(mw.Max().Key)*int64(mh.Max().Key))
	}
}

//func main() {
//	Sol527C(os.Stdin, os.Stdout)
//}
