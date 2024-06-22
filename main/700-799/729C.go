package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://space.bilibili.com/206214
func cf729C(in io.Reader, out io.Writer) {
	var n, k, s, t int
	Fscan(in, &n, &k, &s, &t)
	a := make([]struct{ price, cap int }, n)
	for i := range a {
		Fscan(in, &a[i].price, &a[i].cap)
	}
	sort.Slice(a, func(i, j int) bool { return a[i].cap < a[j].cap })
	p := make([]int, k, k+2)
	for i := range p {
		Fscan(in, &p[i])
	}
	p = append(p, 0, s)
	slices.Sort(p)
	gap := make([]int, k+1)
	for i := range gap {
		gap[i] = p[i+1] - p[i]
	}
	slices.Sort(gap)
	mx := gap[k]

	time := s * 2
	i := 0
	for j, c := range a {
		if c.cap < mx {
			continue
		}
		for ; i <= k && gap[i]*2 <= c.cap; i++ {
			time -= gap[i]
			s -= gap[i]
		}
		if time-c.cap*(k+1-i)+s <= t {
			ans := int(2e9)
			for _, c := range a[j:] {
				ans = min(ans, c.price)
			}
			Fprint(out, ans)
			return
		}
	}
	Fprint(out, -1)
}

//func main() { cf729C(bufio.NewReader(os.Stdin), os.Stdout) }
