package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1550C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int64, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := n*2 - 1
		for i := 0; i < n-2; i++ {
			if (a[i]-a[i+1])*(a[i+1]-a[i+2]) < 0 {
				ans++
			}
		}
		for i := 0; i < n-3; i++ {
			if (a[i+1]-a[i+3])*(a[i+3]-a[i+2]) > 0 && (a[i+1]-a[i])*(a[i]-a[i+2]) > 0 {
				ans++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1550C(os.Stdin, os.Stdout) }
