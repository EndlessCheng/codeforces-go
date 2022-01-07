package main

import (
	"bufio"
	"container/heap"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
type hp65 struct{ sort.IntSlice }

func (h hp65) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp65) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (hp65) Pop() (_ interface{})  { return }

func CF965E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type node struct {
		son [26]*node
		end bool
	}

	root := &node{end: true}
	hs := map[*node]*hp65{root: {}}
	var n, ans int
	var s []byte
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		ans += len(s)
		o := root
		for _, b := range s {
			b -= 'a'
			if o.son[b] == nil {
				o.son[b] = &node{}
				hs[o.son[b]] = &hp65{}
			}
			o = o.son[b]
		}
		o.end = true
		hs[o].IntSlice = []int{len(s)}
	}

	var f func(*node, int)
	f = func(o *node, d int) {
		for _, son := range o.son[:] {
			if son != nil {
				f(son, d+1)
				for _, v := range hs[son].IntSlice {
					heap.Push(hs[o], v)
				}
			}
		}
		if !o.end {
			h := hs[o]
			ans -= h.IntSlice[0] - d
			h.IntSlice[0] = d
			heap.Fix(h, 0)
		}
	}
	f(root, 0)
	Fprint(out, ans)
}

//func main() { CF965E(os.Stdin, os.Stdout) }
