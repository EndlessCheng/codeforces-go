package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF883I(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	Fprint(out, sort.Search(a[n-1]-a[0], func(mx int) bool {
		f := make([]bool, n+1)
		f[0] = true
		j0 := 0
		for i := k - 1; i < n; i++ {
			for a[i]-a[j0] > mx {
				j0++
			}
			for ; j0 <= i-k+1; j0++ {
				if f[j0] {
					f[i+1] = true
					break
				}
			}
		}
		return f[n]
	}))
}

//func main() { CF883I(os.Stdin, os.Stdout) }
