package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

var _x = uint(1)

func fastRand() uint {
	_x ^= _x << 13
	_x ^= _x >> 17
	_x ^= _x << 5
	return _x
}

type node struct {
	lr       [2]*node
	priority uint
	key      int
	value    int
}

func (o *node) rotate(d int) *node {
	x := o.lr[d^1]
	o.lr[d^1] = x.lr[d]
	x.lr[d] = o
	return x
}

type treap struct {
	root       *node
	comparator func(a, b int) int
}

func newTreap() *treap {
	return &treap{comparator: func(a, b int) int {
		if a < b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}}
}

func (t *treap) _put(o *node, key int) *node {
	if o == nil {
		return &node{priority: fastRand(), key: key, value: 1}
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._put(o.lr[cmp], key)
		if o.lr[cmp].priority > o.priority {
			o = o.rotate(cmp ^ 1)
		}
	} else {
		o.value++
	}
	return o
}

func (t *treap) put(key int) { t.root = t._put(t.root, key) }

func (t *treap) _delete(o *node, key int) *node {
	if o == nil {
		return nil
	}
	if cmp := t.comparator(key, o.key); cmp >= 0 {
		o.lr[cmp] = t._delete(o.lr[cmp], key)
	} else {
		if o.value > 1 {
			o.value--
		} else {
			if o.lr[1] == nil {
				return o.lr[0]
			}
			if o.lr[0] == nil {
				return o.lr[1]
			}
			cmp2 := 0
			if o.lr[0].priority > o.lr[1].priority {
				cmp2 = 1
			}
			o = o.rotate(cmp2)
			o.lr[cmp2] = t._delete(o.lr[cmp2], key)
		}
	}
	return o
}

func (t *treap) delete(key int) { t.root = t._delete(t.root, key) }

func (t *treap) max() int {
	var max *node
	for o := t.root; o != nil; o = o.lr[1] {
		max = o
	}
	return max.key
}

func (t *treap) floor(key int) int {
	var floor *node
	for o := t.root; o != nil; {
		switch cmp := t.comparator(key, o.key); {
		case cmp == 0:
			o = o.lr[0]
		case cmp > 0:
			floor = o
			o = o.lr[1]
		default:
			return o.key
		}
	}
	if floor == nil {
		return 0
	}
	return floor.key
}

func (t *treap) ceiling(key int) (ceiling *node) {
	for o := t.root; o != nil; {
		switch cmp := t.comparator(key, o.key); {
		case cmp == 0:
			ceiling = o
			o = o.lr[0]
		case cmp > 0:
			o = o.lr[1]
		default:
			return o
		}
	}
	return
}

//

func (o *node) String() string {
	var s string
	if o.value == 1 {
		s = Sprintf("%v", o.key)
	} else {
		s = Sprintf("%v(%v)", o.key, o.value)
	}
	return s
}

func (o *node) draw(prefix string, isTail bool, str *string) {
	if o.lr[1] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "│   "
		} else {
			newPrefix += "    "
		}
		o.lr[1].draw(newPrefix, false, str)
	}
	*str += prefix
	if isTail {
		*str += "└── "
	} else {
		*str += "┌── "
	}
	*str += o.String() + "\n"
	if o.lr[0] != nil {
		newPrefix := prefix
		if isTail {
			newPrefix += "    "
		} else {
			newPrefix += "│   "
		}
		o.lr[0].draw(newPrefix, true, str)
	}
}

func (t *treap) String() string {
	if t.root == nil {
		return "BST (empty)\n"
	}
	str := "BST\n"
	t.root.draw("", true, &str)
	return str
}

// github.com/EndlessCheng/codeforces-go
func Sol706D(reader io.Reader, writer io.Writer) {
	bitLength := func(n int) int {
		c := 1
		if n>>16 > 0 {
			c += 16
			n >>= 16
		}
		if n>>8 > 0 {
			c += 8
			n >>= 8
		}
		if n>>4 > 0 {
			c += 4
			n >>= 4
		}
		if n>>2 > 0 {
			c += 2
			n >>= 2
		}
		if n-1 > 0 {
			c++
		}
		return c
	}
	min := func(a, b int) int {
		if a <= b {
			return a
		}
		return b
	}
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	t := newTreap()
	t.put(0)

	nextFloor := func(bitPos int, floor, x int) (f int, i int) {
		for i = bitPos; i >= 0; i-- {
			ui := uint(i)
			bitFloor := floor >> ui & 1
			bitX := x >> ui & 1
			if bitFloor != bitX {
				continue
			}
			if bitFloor == 1 {
				// mask floor
				check := floor&^(1<<ui) | (1<<ui - 1)
				newFloor := t.floor(check)
				return newFloor, i
			}
		}
		return -1, -1
	}

	var q, x int
	var op string
	for Fscan(in, &q); q > 0; q-- {
		switch Fscan(in, &op, &x); op[0] {
		case '+':
			t.put(x)
		case '-':
			t.delete(x)
		default:
			floor := t.max()
			if floor == 0 {
				Fprintln(out, x)
				continue
			}
			ans := floor ^ x
			fl := bitLength(floor)
			xl := bitLength(x)
			minL := min(fl, xl)
			for i := minL - 1; i >= 0; {
				floor, i = nextFloor(i, floor, x)
				if floor == -1 {
					break
				}
				if newXor := floor ^ x; newXor > ans {
					ans = newXor
				}
			}
			Fprintln(out, ans)
		}
	}
}

func main() {
	Sol706D(os.Stdin, os.Stdout)
}
