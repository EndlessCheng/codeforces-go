package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1971E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k, q, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &q)
		a := make([]int, k+1)
		for i := 1; i <= k; i++ {
			Fscan(in, &a[i])
		}
		b := make([]int, k+1)
		for i := 1; i <= k; i++ {
			Fscan(in, &b[i])
		}
		for ; q > 0; q-- {
			Fscan(in, &x)
			i := sort.SearchInts(a, x)
			if a[i] == x {
				Fprint(out, b[i], " ")
			} else {
				Fprint(out, b[i-1]+(x-a[i-1])*(b[i]-b[i-1])/(a[i]-a[i-1]), " ")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1971E(bufio.NewReader(os.Stdin), os.Stdout) }
