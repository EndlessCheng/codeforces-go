package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1260D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var m, n, k, t int
	Fscan(in, &m, &n, &k, &t)
	t -= n + 1
	a := make([]int, m)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	traps := make([]struct{ l, r, d int }, k)
	for i := range traps {
		Fscan(in, &traps[i].l, &traps[i].r, &traps[i].d)
	}
	sort.Slice(traps, func(i, j int) bool { return traps[i].l < traps[j].l })

	ans := sort.Search(m, func(mx int) bool {
		low := a[mx]
		i := 0
		for ; i < k && traps[i].d <= low; i++ {
		}
		if i == k {
			return false
		}
		s := 0
		l0 := traps[i].l
		maxR := traps[i].r
		for _, p := range traps[i+1:] {
			if p.d <= low {
				continue
			}
			l, r := p.l, p.r
			if l > maxR {
				s += maxR - l0 + 1
				l0 = l
			}
			maxR = max(maxR, r)
		}
		s += maxR - l0 + 1
		return s*2 > t
	})
	Fprint(out, ans)
}

//func main() { cf1260D(os.Stdin, os.Stdout) }
