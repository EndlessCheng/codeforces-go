package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1234D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var s []byte
	Fscan(in, &s)
	for i := range s {
		s[i] -= 'a'
	}

	n := len(s)
	trees := make([][]int, 26)
	for i := range trees {
		trees[i] = make([]int, n+1)
	}
	add := func(tree []int, idx int, val int) {
		for ; idx <= n; idx += idx & -idx {
			tree[idx] += val
		}
	}
	sum := func(tree []int, idx int) (res int) {
		for ; idx > 0; idx &= idx - 1 {
			res += tree[idx]
		}
		return
	}
	query := func(tree []int, l, r int) int {
		return sum(tree, r) - sum(tree, l-1)
	}

	for i, c := range s {
		add(trees[c], i+1, 1)
	}
	var q, op int
	for Fscan(in, &q); q > 0; q-- {
		if Fscan(in, &op); op == 1 {
			var idx int
			var raw string
			Fscan(in, &idx, &raw)
			ch := raw[0] - 'a'
			add(trees[s[idx-1]], idx, -1)
			add(trees[ch], idx, 1)
			s[idx-1] = ch
		} else {
			var l, r int
			Fscan(in, &l, &r)
			res := 0
			for _, tree := range trees {
				if query(tree, l, r) > 0 {
					res++
				}
			}
			Fprintln(out, res)
		}
	}
}

//func main() {
//	Sol1234D(os.Stdin, os.Stdout)
//}
