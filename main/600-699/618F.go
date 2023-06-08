package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF618F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	//    sa[r]-sa[l]=sb[r]-sb[l]
	// => sb[l]-sa[l]=sb[r]-sa[r]
	var n int
	var v int64
	Fscan(in, &n)
	sa := make([]int64, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sa[i+1] = sa[i] + v
	}
	sb := make([]int64, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sb[i+1] = sb[i] + v
	}
	flip := sa[n] > sb[n]
	if flip {
		sa, sb = sb, sa
	}
	print := func(l, r int) {
		Fprintln(out, r-l)
		for i := l + 1; i <= r; i++ {
			Fprint(out, i, " ")
		}
		Fprintln(out)
	}

	pre := make([]int, n)
	for i := range pre {
		pre[i] = -1
	}
	bj := make([]int, n+1)
	j := 0
	for i, s := range sa {
		for sb[j] < s {
			j++
		}
		d := sb[j] - s
		if pre[d] != -1 {
			if flip {
				print(bj[pre[d]], j)
				print(pre[d], i)
			} else {
				print(pre[d], i)
				print(bj[pre[d]], j)
			}
			return
		}
		pre[d] = i
		bj[i] = j
	}
}

//func main() { CF618F(os.Stdin, os.Stdout) }
