package copypasta

import "errors"

// TODO: 待整理
// https://oi-wiki.org/ds/splay/

type Int int

func (p Int) Equals(key Key) bool {
	return int(p) == int(key.(Int))
}

func (p Int) Less(key Key) bool {
	return int(p) < int(key.(Int))
}

type Key interface {
	Equals(key Key) bool
	Less(key Key) bool
}

type splayNode struct {
	key   Key
	left  *splayNode
	right *splayNode
}

func newNode(key Key) *splayNode {
	return &splayNode{key, nil, nil}
}

type SplayTree struct {
	root *splayNode
	tmp  *splayNode
	len  int
}

func NewSplayTree() *SplayTree {
	return &SplayTree{nil, newNode(nil), 0}
}

func rotateLeft(x, p *splayNode) {
	p.right = x.left
	x.left = p
}

func rotateRight(x, p *splayNode) {
	p.left = x.right
	x.right = p
}

func (p *SplayTree) splay(x *splayNode, key Key) *splayNode {
	if x == nil {
		return nil
	}

	left := p.tmp
	right := p.tmp

	for {
		if key.Less(x.key) {
			y := x.left
			if y == nil {
				break
			}
			if key.Less(y.key) { // zig-zig
				rotateRight(y, x)
				x = y
				if x.left == nil {
					break
				}
			}
			// link right
			right.left = x
			right = x
			// move left
			x = x.left
		} else if x.key.Less(key) {
			y := x.right
			if y == nil {
				break
			}
			if y.key.Less(key) { // zig-zig
				rotateLeft(y, x)
				x = y
				if x.right == nil {
					break
				}
			}
			// link left
			left.right = x
			left = x
			// move right
			x = x.right
		} else {
			break
		}
	}

	left.right = x.left
	right.left = x.right
	x.left = p.tmp.right
	x.right = p.tmp.left

	return x
}

// left <= key < right
func (p *SplayTree) split(key Key) (left, right *splayNode) {
	p.root = p.splay(p.root, key)
	if p.root.key.Equals(key) || p.root.key.Less(key) {
		right := p.root.right
		p.root.right = nil
		return p.root, right
	} else {
		left := p.root.left
		p.root.left = nil
		return left, p.root
	}
}

// keys from left tree must be less then keys from right tree
func (p *SplayTree) join(left, right *splayNode) *splayNode {
	if left == nil {
		return right
	} else if right == nil {
		return left
	}
	left = p.splay(left, right.key)
	left.right = right
	return left
}

type Set interface {
	Len() int
	Insert(key Key) error
	Find(key Key) bool
	Remove(key Key) error
}

func NewSet() Set {
	return Set(NewSplayTree())
}

func (p *SplayTree) Len() int {
	return p.len
}

func (p *SplayTree) Insert(key Key) error {
	if p.root == nil {
		p.root = newNode(key)
		p.len++
	} else {
		p.root = p.splay(p.root, key)
		if p.root.key.Equals(key) {
			return errors.New("such key already exists")
		} else {
			left, right := p.split(key)
			p.root = newNode(key)
			p.root.left = left
			p.root.right = right
			p.len++
		}
	}
	return nil
}

func (p *SplayTree) Find(key Key) bool {
	if p.root == nil {
		return false
	}
	p.root = p.splay(p.root, key)
	return p.root.key.Equals(key)
}

func (p *SplayTree) Remove(key Key) error {
	p.root = p.splay(p.root, key)
	if p.root == nil || !p.root.key.Equals(key) {
		return errors.New("such key doesn't exist")
	}
	p.root = p.join(p.split(key))
	return nil
}
