package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf842D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type node struct {
		son  [2]*node
		size int
	}
	root := &node{}

	var n, m, v, xor int
	Fscan(in, &n, &m)
	vis := [3e5 + 1]bool{}
	for range n {
		Fscan(in, &v)
		if vis[v] {
			continue
		}
		vis[v] = true
		o := root
		for i := 18; i >= 0; i-- {
			b := v >> i & 1
			if o.son[b] == nil {
				o.son[b] = &node{}
			}
			o = o.son[b]
			o.size++
		}
	}

	for range m {
		Fscan(in, &v)
		xor ^= v
		o := root
		mex := 0
		for i := 18; i >= 0 && o != nil; i-- {
			b := xor >> i & 1
			if o.son[b] != nil && o.son[b].size == 1<<i {
				mex |= 1 << i
				b ^= 1
			}
			o = o.son[b]
		}
		Fprintln(out, mex)
	}
}

//func main() { cf842D(bufio.NewReader(os.Stdin), os.Stdout) }
