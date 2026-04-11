package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://github.com/EndlessCheng
func cf2185H(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, k int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k)
		a := make([]int, n+1)
		s := make([]int, n+1)
		p := []int{}
		for i := 1; i <= n; i++ {
			Fscan(in, &a[i])
			s[i] = s[i-1] + a[i]
			if a[i] > s[i-1] {
				p = append(p, i)
			}
		}

		for i := 1; i <= n; i++ {
			h := []int{}
			for _, x := range p {
				if x == i {
					continue
				}
				ex := 0
				if x < i {
					ex = a[i]
				}
				if a[x] > s[x-1]+ex {
					h = append(h, x)
				}
			}

			if len(h) < k {
				Fprint(out, n, " ")
				continue
			}
			ans := 0
			if k > 0 {
				x := h[len(h)-k]
				ans = n - x
				if x >= i {
					ans++
				}
			}
			if len(h) == k {
				le := 0
				if s[i-1] < a[i] {
					le = 1
				}
				v := (le + 1) * a[i]
				j := sort.Search(n, func(j int) bool { return s[j+1] >= v }) + 1
				ans += j - le
			}
			Fprint(out, ans, " ")
		}
		Fprintln(out)
	}
}

//func main() { cf2185H(bufio.NewReader(os.Stdin), os.Stdout) }
