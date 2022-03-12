package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1650D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := make([]interface{}, n)
		for i := n; i > 0; i-- {
			for j, v := range a {
				if v == i {
					ans[i-1] = (j + 1) % i
					a = append(a[j+1:], a[:j]...)
					break
				}
			}
		}
		Fprintln(out, ans...)
	}
}

//func main() { CF1650D(os.Stdin, os.Stdout) }
