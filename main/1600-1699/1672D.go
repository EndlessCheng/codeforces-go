package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1672D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, n)
		for i := range b {
			Fscan(in, &b[i])
		}
		c := make([]int, n+1)
		for i, j := n-1, n-1; i >= 0; i-- {
			if i < n-1 && b[i] == b[i+1] {
				c[b[i]]++ // a 中 <=j 的任意位置可以塞个 b[i]
				continue
			}
			for ; a[j] != b[i]; j-- {
				if c[a[j]] == 0 {
					Fprintln(out, "NO")
					continue o
				}
				c[a[j]]-- // 塞入
			}
			j--
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1672D(os.Stdin, os.Stdout) }
