package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf2167E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k, x int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &x)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		slices.Sort(a)
		low := sort.Search(x, func(low int) bool {
			low++
			s := max(a[0]-low+1, 0)
			for i := 1; i < n; i++ {
				s += max(a[i]-a[i-1]-low*2+1, 0)
			}
			s += max(x-a[n-1]-low+1, 0)
			return s < k
		})
		pr := func(l, r int) {
			for i := l; i <= r && k > 0; i++ {
				k--
				Fprint(out, i, " ")
			}
		}
		if low == 0 {
			pr(0, k-1)
		} else {
			pr(0, a[0]-low)
			for i := 1; i < n; i++ {
				pr(a[i-1]+low, a[i]-low)
			}
			pr(a[n-1]+low, x)
		}
		Fprintln(out)
	}
}

//func main() { cf2167E(bufio.NewReader(os.Stdin), os.Stdout) }
