package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1482B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		if n > 2 {
			for i, d := 2, a[1]-a[0]; i < n; i++ {
				if dd := a[i] - a[i-1]; dd != d {
					m := abs(dd - d)
					if d < 0 {
						d += m
					}
					for i, v := range a {
						if v >= m || i > 0 && (a[i-1]+d)%m != v {
							Fprintln(out, -1)
							continue o
						}
					}
					Fprintln(out, m, d)
					continue o
				}
			}
		}
		Fprintln(out, 0)
	}
}

//func main() { CF1482B(os.Stdin, os.Stdout) }
