package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1311B(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T, n, m, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		p := make([]bool, n)
		for ; m > 0; m-- {
			Fscan(in, &v)
			p[v-1] = true
		}
		for i := 0; i < n; i++ {
			st := i
			for ; p[i]; i++ {
			}
			sort.Ints(a[st : i+1])
		}
		if sort.IntsAreSorted(a) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1311B(os.Stdin, os.Stdout) }
