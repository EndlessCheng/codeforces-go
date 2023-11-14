package main

import (
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF388A(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	for n > 0 {
		ans++
		c := 0
		for i, v := range a {
			if v >= c {
				c++
				n--
				a[i] = -1
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF388A(os.Stdin, os.Stdout) }
