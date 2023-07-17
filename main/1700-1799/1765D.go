package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1765D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	ans := int64(n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		ans += int64(a[i])
	}
	sort.Ints(a)
	l, r := 0, n-1
	for l < r {
		if a[l]+a[r] <= m {
			l++
			if l < r {
				ans -= 2
			} else {
				ans--
			}
		}
		r--
	}
	Fprint(out, ans)
}

//func main() { CF1765D(os.Stdin, os.Stdout) }
