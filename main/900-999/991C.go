package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf991C(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	ans := 1 + sort.Search(n/2, func(k int) bool {
		k++
		m := n
		s := 0
		for m > k {
			s += k
			m -= k
			m -= m / 10
		}
		return (s+m)*2 >= n
	})
	Fprint(out, ans)
}

//func main() { cf991C(os.Stdin, os.Stdout) }
