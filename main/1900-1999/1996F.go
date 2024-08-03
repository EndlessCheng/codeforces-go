package main

import (
	. "fmt"
	"io"
	"sort"
)

func cf1996F(in io.Reader, out io.Writer) {
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]struct{ a, b int }, n)
		for i := range a {
			Fscan(in, &a[i].a)
		}
		for i := range a {
			Fscan(in, &a[i].b)
		}
		left := k
		mx := sort.Search(1e9, func(mx int) bool {
			k := k
			for _, p := range a {
				if p.a > mx {
					k -= (p.a-mx-1)/p.b + 1
				}
			}
			if k < 0 {
				return false
			}
			left = min(left, k)
			return true
		})
		ans := 0
		for i, p := range a {
			if p.a > mx {
				k := (p.a-mx-1)/p.b + 1
				ans += (p.a*2 - (k-1)*p.b) * k / 2
				a[i].a -= k * p.b
			}
		}
		Fprintln(out, ans+mx*left)
	}
}

//func main() { cf1996F(bufio.NewReader(os.Stdin), os.Stdout) }
