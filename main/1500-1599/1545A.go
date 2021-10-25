package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1545A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n int
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		b := []int{}
		for i := range a {
			Fscan(in, &a[i])
			if i&1 == 0 {
				b = append(b, a[i])
			}
		}
		sort.Ints(a)
		sort.Ints(b)
		for i, v := range b {
			if v != a[i*2] {
				Fprintln(out, "NO")
				continue o
			}
		}
		Fprintln(out, "YES")
	}
}

//func main() { CF1545A(os.Stdin, os.Stdout) }
