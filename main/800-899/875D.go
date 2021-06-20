package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF875D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	ans := int64(0)
	var n, v int
	Fscan(in, &n)
	type vi struct{ v, i int }
	stk := []vi{{2e9, -1}}
	type lr struct{ v, l, r int }
	set := []lr{}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		for stk[len(stk)-1].v <= v {
			stk = stk[:len(stk)-1]
		}
		stk = append(stk, vi{v, i})
		for j := range set {
			set[j].v |= v
		}
		set = append(set, lr{v, i, i + 1})
		k := 0
		for _, q := range set[1:] {
			if set[k].v != q.v {
				k++
				set[k] = q
			} else {
				set[k].r = q.r
			}
		}
		set = set[:k+1]
		for _, p := range set {
			j := sort.Search(len(stk), func(i int) bool { return stk[i].v < p.v }) - 1
			j = stk[j].i + 1
			if j < p.l {
				j = p.l
			}
			ans += int64(p.r - j)
		}
	}
	Fprint(out, ans)
}

//func main() { CF875D(os.Stdin, os.Stdout) }
