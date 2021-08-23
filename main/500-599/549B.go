package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF549B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	g := make([][]byte, n)
	for i := range g {
		Fscan(in, &g[i])
	}
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	ans := []int{}
o:
	for {
		for i, v := range a {
			if v == 0 {
				ans = append(ans, i+1)
				for j, b := range g[i] {
					a[j] -= int(b & 1)
				}
				continue o
			}
		}
		break
	}
	sort.Ints(ans)
	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
}

//func main() { CF549B(os.Stdin, os.Stdout) }
