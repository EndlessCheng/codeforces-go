package main

import . "fmt"

// http://en.wikipedia.org/wiki/Red%E2%80%93black_tree

type keyType int   // *custom* 图方便可以全局替换
type valueType int // *custom* 图方便可以全局替换
type color bool

const red, black color = true, false

type Node struct {
	Left, Right, Parent *Node

	N     int // 以该节点为根的子树中的节点总数
	MN    int
	Key   keyType
	Value valueType
	color color
}

type Tree struct {
	Root       *Node
	comparator func(a, b keyType) int
}

func NewRBTree() *Tree {
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

/***************************************************************************
 *  Node helper methods.
 ***************************************************************************/

func (node *Node) isRed() bool {
	if node != nil {
		return node.color == red
	}
	return false
}

func (node *Node) size() int {
	if node != nil {
		return node.N
	}
	return 0
}

func (node *Node) mSize() int {
	if node != nil {
		return node.MN
	}
	return 0
}

func (t *Tree) Size() int   { return t.Root.size() }
func (t *Tree) MSize() int  { return t.Root.mSize() }
func (t *Tree) Empty() bool { return t.Size() == 0 }

func (t *Tree) Put(key keyType, value valueType) *Node {
	var insertedNode *Node
	if t.Root == nil {
		// 标准的插入操作，和父节点用红链接相连
		t.Root = &Node{Key: key, Value: value, color: red}
		insertedNode = t.Root
		insertedNode.N = 1
		insertedNode.MN = 1
	} else {
		o := t.Root
		stack := []*Node{o}
	loop:
		for {
			switch cmp := t.comparator(key, o.Key); {
			case cmp < 0:
				if o.Left == nil {
					// 标准的插入操作，和父节点用红链接相连
					o.Left = &Node{Key: key, Value: value, color: red}
					insertedNode = o.Left
					break loop
				}
				o = o.Left
			case cmp > 0:
				if o.Right == nil {
					// 标准的插入操作，和父节点用红链接相连
					o.Right = &Node{Key: key, Value: value, color: red}
					insertedNode = o.Right
					break loop
				}
				o = o.Right
			default:
				// just change value
				o.Value = value
				return o
			}
			stack = append(stack, o)
		}
		insertedNode.Parent = o
		insertedNode.N = 1
		insertedNode.MN = 1
		for len(stack) > 0 {
			stack, o = stack[:len(stack)-1], stack[len(stack)-1]
			o.N = 1 + o.Left.size() + o.Right.size()
			o.MN = int(o.Value) + o.Left.mSize() + o.Right.mSize()
		}
	}
	t.insertCase1(insertedNode)
	return insertedNode
}

// Get
func (t *Tree) Lookup(key keyType) *Node {
	for o := t.Root; o != nil; {
		switch cmp := t.comparator(key, o.Key); {
		case cmp < 0:
			o = o.Left
		case cmp > 0:
			o = o.Right
		default:
			return o
		}
	}
	return nil
}

func (t *Tree) LookupStack(key keyType) (stack []*Node, found bool) {
	for o := t.Root; o != nil; {
		stack = append(stack, o)
		switch cmp := t.comparator(key, o.Key); {
		case cmp < 0:
			o = o.Left
		case cmp > 0:
			o = o.Right
		default:
			found = true
			return
		}
	}
	return
}

func (t *Tree) Contains(key keyType) bool { return t.Lookup(key) != nil }

func (t *Tree) Delete(key keyType) {
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
			node.color = colorOf(child)
			t.deleteCase1(node)
		}
		t.replaceParent(node, child)
		if node.Parent == nil && child != nil {
			child.color = black
		}
	}
}

// valueType must be int
func (t *Tree) MultiPut(key keyType) *Node {
	if o := t.Lookup(key); o != nil {
		o.Value++
		return o
	}
	return t.Put(key, 1)
}

// valueType must be int
func (t *Tree) MultiDelete(key keyType) {
	if o := t.Lookup(key); o != nil {
		o.Value--
		if o.Value == 0 {
			t.Delete(key)
		}
	}
}

// valueType must be int
func (t *Tree) MultiPutMN(key keyType) *Node {
	if stack, found := t.LookupStack(key); found {
		var leaf, o *Node
		stack, leaf = stack[:len(stack)-1], stack[len(stack)-1]
		leaf.Value++
		leaf.MN++
		for len(stack) > 0 {
			stack, o = stack[:len(stack)-1], stack[len(stack)-1]
			o.MN = int(o.Value) + o.Left.mSize() + o.Right.mSize()
		}
		return leaf
	}
	return t.Put(key, 1)
}

// valueType must be int
func (t *Tree) MultiDeleteMN(key keyType) {
	if stack, found := t.LookupStack(key); found {
		var leaf, o *Node
		stack, leaf = stack[:len(stack)-1], stack[len(stack)-1]
		leaf.Value--
		leaf.MN--
		for len(stack) > 0 {
			stack, o = stack[:len(stack)-1], stack[len(stack)-1]
			o.MN = int(o.Value) + o.Left.mSize() + o.Right.mSize()
		}
		if leaf.Value == 0 {
			t.Delete(key)
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
		switch cmp := t.comparator(key, o.Key); {
		case cmp < 0:
			o = o.Left
		case cmp > 0:
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
		switch cmp := t.comparator(key, o.Key); {
		case cmp < 0:
			ceiling = o
			o = o.Left
		case cmp > 0:
			o = o.Right
		default:
			return o
		}
	}
	return
}

// 排名为 k 的节点（k 从 0 开始）
// 即小于节点的键的数量为 k
func (t *Tree) Select(k int) *Node {
	if k < 0 {
		return nil
	}
	for o := t.Root; o != nil; {
		switch ls := o.Left.size(); {
		case k < ls:
			o = o.Left
		case k > ls:
			k -= 1 + ls
			o = o.Right
		default:
			return o
		}
	}
	return nil
}

// 小于 key 的键的数量
func (t *Tree) Rank(key keyType) (cnt int) {
	for o := t.Root; o != nil; {
		switch cmp := t.comparator(key, o.Key); {
		case cmp < 0:
			o = o.Left
		case cmp > 0:
			cnt += 1 + o.Left.size()
			o = o.Right
		default:
			cnt += o.Left.size()
			return
		}
	}
	return
}

// 排名为 k 的节点（k 从 0 开始）
// 即小于节点的键的数量为 k
func (t *Tree) SelectMN(k int) *Node {
	if k < 0 {
		return nil
	}
	for o := t.Root; o != nil; {
		switch ls := o.Left.mSize(); {
		case k < ls:
			o = o.Left
		case k > ls:
			k -= int(o.Value) + ls
			if k < 0 {
				return o
			}
			o = o.Right
		default:
			return o
		}
	}
	return nil
}

// 小于 key 的键的数量
func (t *Tree) RankMN(key keyType) (cnt int) {
	for o := t.Root; o != nil; {
		switch cmp := t.comparator(key, o.Key); {
		case cmp < 0:
			o = o.Left
		case cmp > 0:
			cnt += int(o.Value) + o.Left.mSize()
			o = o.Right
		default:
			cnt += o.Left.mSize()
			return
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

func (t *Tree) rotateLeft(o *Node) {
	r := o.Right
	t.replaceParent(o, r)
	o.Right = r.Left
	if r.Left != nil {
		r.Left.Parent = o
	}
	r.Left = o
	o.Parent = r
	r.N = o.N
	r.MN = o.MN
	o.N = 1 + o.Left.size() + o.Right.size()
	o.MN = int(o.Value) + o.Left.mSize() + o.Right.mSize()
}

func (t *Tree) rotateRight(o *Node) {
	l := o.Left
	t.replaceParent(o, l)
	o.Left = l.Right
	if l.Right != nil {
		l.Right.Parent = o
	}
	l.Right = o
	o.Parent = l
	l.N = o.N
	l.MN = o.MN
	o.N = 1 + o.Left.size() + o.Right.size()
	o.MN = int(o.Value) + o.Left.mSize() + o.Right.mSize()
}

func (t *Tree) replaceParent(old, new *Node) {
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

func (t *Tree) insertCase1(o *Node) {
	if o.Parent == nil {
		o.color = black
	} else {
		t.insertCase2(o)
	}
}

func (t *Tree) insertCase2(o *Node) {
	if colorOf(o.Parent) == black {
		return
	}
	t.insertCase3(o)
}

func (t *Tree) insertCase3(o *Node) {
	uncle := o.uncle()
	if colorOf(uncle) == red {
		o.Parent.color = black
		uncle.color = black
		o.grandparent().color = red
		t.insertCase1(o.grandparent())
	} else {
		t.insertCase4(o)
	}
}

func (t *Tree) insertCase4(o *Node) {
	grandparent := o.grandparent()
	if o == o.Parent.Right && o.Parent == grandparent.Left {
		t.rotateLeft(o.Parent)
		o = o.Left
	} else if o == o.Parent.Left && o.Parent == grandparent.Right {
		t.rotateRight(o.Parent)
		o = o.Right
	}
	t.insertCase5(o)
}

func (t *Tree) insertCase5(o *Node) {
	o.Parent.color = black
	grandparent := o.grandparent()
	grandparent.color = red
	if o == o.Parent.Left && o.Parent == grandparent.Left {
		t.rotateRight(grandparent)
	} else if o == o.Parent.Right && o.Parent == grandparent.Right {
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

func (t *Tree) deleteCase1(o *Node) {
	if o.Parent == nil {
		return
	}
	t.deleteCase2(o)
}

func (t *Tree) deleteCase2(o *Node) {
	sibling := o.sibling()
	if colorOf(sibling) == red {
		o.Parent.color = red
		sibling.color = black
		if o == o.Parent.Left {
			t.rotateLeft(o.Parent)
		} else {
			t.rotateRight(o.Parent)
		}
	}
	t.deleteCase3(o)
}

func (t *Tree) deleteCase3(o *Node) {
	sibling := o.sibling()
	if colorOf(o.Parent) == black &&
		colorOf(sibling) == black &&
		colorOf(sibling.Left) == black &&
		colorOf(sibling.Right) == black {
		sibling.color = red
		t.deleteCase1(o.Parent)
	} else {
		t.deleteCase4(o)
	}
}

func (t *Tree) deleteCase4(o *Node) {
	sibling := o.sibling()
	if colorOf(o.Parent) == red &&
		colorOf(sibling) == black &&
		colorOf(sibling.Left) == black &&
		colorOf(sibling.Right) == black {
		sibling.color = red
		o.Parent.color = black
	} else {
		t.deleteCase5(o)
	}
}

func (t *Tree) deleteCase5(o *Node) {
	sibling := o.sibling()
	if o == o.Parent.Left &&
		colorOf(sibling) == black &&
		colorOf(sibling.Left) == red &&
		colorOf(sibling.Right) == black {
		sibling.color = red
		sibling.Left.color = black
		t.rotateRight(sibling)
	} else if o == o.Parent.Right &&
		colorOf(sibling) == black &&
		colorOf(sibling.Right) == red &&
		colorOf(sibling.Left) == black {
		sibling.color = red
		sibling.Right.color = black
		t.rotateLeft(sibling)
	}
	t.deleteCase6(o)
}

func (t *Tree) deleteCase6(o *Node) {
	sibling := o.sibling()
	sibling.color = colorOf(o.Parent)
	o.Parent.color = black
	if o == o.Parent.Left && colorOf(sibling.Right) == red {
		sibling.Right.color = black
		t.rotateLeft(o.Parent)
	} else if colorOf(sibling.Left) == red {
		sibling.Left.color = black
		t.rotateRight(o.Parent)
	}
}

func colorOf(o *Node) color {
	if o == nil {
		return black
	}
	return o.color
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
func (it *Iterator) End() bool { return it.position == end }

// RBegin moves the iterator to the last element and returns true if there was a last element in the container.
// Modifies the state of the iterator.
func (t *Tree) RBegin() *Iterator {
	it := &Iterator{tree: t, position: end}
	return it.Prev()
}
func (it *Iterator) REnd() bool { return it.position == begin }

// Keys returns all keys in-order
func (t *Tree) Keys() []keyType {
	keys := make([]keyType, 0, t.Size())
	for it := t.Begin(); !it.End(); it.Next() {
		keys = append(keys, it.node.Key)
	}
	return keys
}

// valueType must be int
func (t *Tree) MultiKeys() []keyType {
	keys := make([]keyType, 0, t.Size())
	for it := t.Begin(); !it.End(); it.Next() {
		k, v := it.node.Key, int(it.node.Value) // it.node.MN
		for i := 0; i < v; i++ {
			keys = append(keys, k)
		}
	}
	return keys
}

// Values returns all values in-order based on the key.
func (t *Tree) Values() []valueType {
	values := make([]valueType, 0, t.Size())
	for it := t.Begin(); !it.End(); it.Next() {
		values = append(values, it.node.Value)
	}
	return values
}

//

//func (node *Node) String() string {return Sprint(node.Key)}

func (node *Node) String() string {
	if node.Value == 1 {
		return Sprint(node.Key)
	}
	return Sprintf("%d(%d)", node.Key, node.Value)
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
	if !t.Empty() {
		t.Root.draw("", true, &str)
	}
	return str
}
