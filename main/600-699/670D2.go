package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF670D2(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	b := make([]int, n)
	for i := range b {
		Fscan(in, &b[i])
	}
	ans := sort.Search(k+1e9, func(mx int) bool {
		mx++
		k := k
		for i, v := range a {
			k -= max(mx*v-b[i], 0)
			if k < 0 {
				return true
			}
		}
		return false
	})
	Fprint(out, ans)
}

//func main() { CF670D2(os.Stdin, os.Stdout) }
