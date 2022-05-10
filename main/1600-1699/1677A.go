package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1677A(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		l := make([][]int, n)
		for i := range a {
			Fscan(in, &a[i])
			l[i] = make([]int, n+1)
			for j, v := range a[:i] {
				l[i][j+1] = l[i][j]
				if v < a[i] {
					l[i][j+1]++
				}
			}
		}
		ans := int64(0)
		for i := 1; i < n-2; i++ {
			for j, c := n-2, 0; j > i; j-- {
				if a[i] > a[j+1] {
					c++
				}
				ans += int64(l[j][i] * c)
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1677A(os.Stdin, os.Stdout) }
