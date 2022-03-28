package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1658C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		c1 := 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] == 1 {
				c1++
			}
			if i > 0 && a[i]-a[i-1] > 1 {
				c1 = 2
			}
		}
		if c1 != 1 || a[0]-a[n-1] > 1 {
			Fprintln(out, "NO")
		} else {
			Fprintln(out, "YES")
		}
	}
}

//func main() { CF1658C(os.Stdin, os.Stdout) }
