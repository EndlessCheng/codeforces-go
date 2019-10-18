package copypasta

import "fmt"

// http://en.wikipedia.org/wiki/Red%E2%80%93black_tree

type keyType int            // *custom*
type valueType int          // *custom*
var zeroValue valueType = 0 // "" for string

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

// valueType must be int
func (t *Tree) MultiInsert(key keyType) {
	if o := t.Lookup(key); o != nil {
		o.Value++
	} else {
		t.Insert(key, 1)
	}
}

// valueType must be int
func (t *Tree) MultiErase(key keyType) {
	if o := t.Lookup(key); o != nil {
		o.Value--
		if o.Value == 0 {
			t.Erase(key)
		}
	}
}

func (t *Tree) IsEmpty() bool {
	return t.size == 0
}

func (t *Tree) Size() int {
	return t.size
}

// Keys returns all keys in-order
func (t *Tree) Keys() []keyType {
	keys := make([]keyType, 0, t.size)
	for it := t.Begin(); !it.IsEnd(); it.Next() {
		keys = append(keys, it.node.Key)
	}
	return keys
}

// valueType must be int
func (t *Tree) MultiKeys() []keyType {
	keys := make([]keyType, 0, t.size)
	for it := t.Begin(); !it.IsEnd(); it.Next() {
		k, v := it.node.Key, int(it.node.Value)
		for i := 0; i < v; i++ {
			keys = append(keys, k)
		}
	}
	return keys
}

// Values returns all values in-order based on the key.
func (t *Tree) Values() []valueType {
	values := make([]valueType, 0, t.size)
	for it := t.Begin(); !it.IsEnd(); it.Next() {
		values = append(values, it.node.Value)
	}
	return values
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

// Ceiling finds ceiling node of the input key, return the ceiling node or nil if no ceiling is found.
//
// Ceiling node is defined as the smallest node that is larger than or equal to the given node.
// A ceiling node may not be found, either because the tree is empty, or because
// all nodes in the tree are smaller than the given node.
func (t *Tree) Ceiling(key keyType) (ceiling *Node) {
	for o := t.Root; o != nil; {
		compare := t.comparator(key, o.Key)
		switch {
		case compare < 0:
			ceiling = o
			o = o.Left
		case compare > 0:
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

//

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

// Prev moves the iterator to the previous element and returns itself.
// Modifies the state of the iterator.
func (it *Iterator) Prev() *Iterator {
	if it.position == begin {
		goto begin
	}
	if it.position == end {
		right := it.tree.Max()
		if right == nil {
			goto begin
		}
		it.node = right
		goto between
	}
	if it.node.Left != nil {
		it.node = it.node.Left
		for it.node.Right != nil {
			it.node = it.node.Right
		}
		goto between
	}
	if it.node.Parent != nil {
		node := it.node
		for it.node.Parent != nil {
			it.node = it.node.Parent
			if it.tree.comparator(node.Key, it.node.Key) >= 0 {
				goto between
			}
		}
	}

begin:
	it.node = nil
	it.position = begin
	return it

between:
	it.position = between
	return it
}

// Begin moves the iterator to the first element and returns true if there was a first element in the container.
// Modifies the state of the iterator
func (t *Tree) Begin() *Iterator {
	it := &Iterator{tree: t, position: begin}
	return it.Next()
}
func (it *Iterator) IsEnd() bool { return it.position == end }

// RBegin moves the iterator to the last element and returns true if there was a last element in the container.
// Modifies the state of the iterator.
func (t *Tree) RBegin() *Iterator {
	it := &Iterator{tree: t, position: end}
	return it.Prev()
}
func (it *Iterator) IsREnd() bool { return it.position == begin }

//

func (node *Node) String() string {
	return fmt.Sprint(node.Key)
}

func (node *Node) draw(prefix string, isTail bool, str *string) {
	if node.Right != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		node.Right.draw(newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += node.String() + "\n"
	if node.Left != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		node.Left.draw(newPrefix, true, str)
	}
}

func (t *Tree) String() string {
	str := "RedBlackTree\n"
	if !t.IsEmpty() {
		t.Root.draw("", true, &str)
	}
	return str
}
