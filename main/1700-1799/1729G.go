package main

import (
	"bufio"
	. "fmt"
	"index/suffixarray"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1729G(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var T int
	var s, t []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &t)
		a := sort.IntSlice(suffixarray.New(s).Lookup(t, -1))
		sort.Ints(a)
		n, m := len(a), len(t)
		minMove := 0
		for i := 0; i < n; i = a.Search(a[a.Search(a[i]+m)-1] + m) {
			minMove++
		}
		f := make([]int, n+1)
		f[n] = 1
		for i := 0; i < minMove; i++ {
			for j, p := range a {
				f[j] = 0
				for _, q := range a[j:a.Search(p+m)] {
					f[j] = (f[j] + f[a.Search(q+m)]) % 1000000007
				}
			}
		}
		Fprintln(out, minMove, f[0])
	}
}

//func main() { CF1729G(os.Stdin, os.Stdout) }
