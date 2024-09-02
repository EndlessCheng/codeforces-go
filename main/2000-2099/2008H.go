package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

func cf2008H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, q, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &q)
		s := make([]int, n+2)
		for i := 0; i < n; i++ {
			Fscan(in, &x)
			s[x+1]++
		}
		for i := 2; i <= n; i++ {
			s[i+1] += s[i]
		}
		ans := make([]int, n+1)
		for x = 1; x <= n; x++ {
			ans[x] = sort.Search(x-1, func(m int) bool {
				c := 0
				for i := 0; i <= n; i += x {
					c += s[min(i+m, n)+1] - s[i]
				}
				return c > n/2
			})
		}
		for ; q > 0; q-- {
			Fscan(in, &x)
			Fprint(out, ans[x], " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2008H(bufio.NewReader(os.Stdin), os.Stdout) }
