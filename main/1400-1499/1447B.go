package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1447B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n*m)
		c, s := 0, 0
		for i := range a {
			Fscan(in, &a[i])
			if a[i] < 0 {
				c++
				a[i] = -a[i]
			}
			s += a[i]
		}
		sort.Ints(a)
		if a[0] > 0 && c&1 > 0 {
			s -= 2 * a[0]
		}
		Fprintln(out, s)
	}
}

//func main() { CF1447B(os.Stdin, os.Stdout) }
