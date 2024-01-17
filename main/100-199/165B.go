package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf165B(in io.Reader, out io.Writer) {
	var n, k int
	Fscan(in, &n, &k)
	ans := sort.Search(n, func(v int) bool {
		s := v
		for ; v >= k; v /= k {
			s += v / k
		}
		return s >= n
	})
	Fprint(out, ans)
}

//func main() { cf165B(os.Stdin, os.Stdout) }
