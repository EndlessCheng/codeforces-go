package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF557C(_r io.Reader, out io.Writer) {
	const mx int = 1e5
	in := bufio.NewReader(_r)
	var n, e int
	Fscan(in, &n)
	el := [201][]int{}
	cnt := [mx + 1][2]int{}
	l := make([]int, n)
	for i := range l {
		Fscan(in, &l[i])
	}
	for _, v := range l {
		Fscan(in, &e)
		el[e] = append(el[e], v)
		cnt[v][0]++
		cnt[v][1] += e
	}
	for _, ls := range el {
		sort.Ints(ls)
	}

	ans := int(1e9)
	for i, cost := mx, 0; i > 0; i-- {
		cl, s, c := cnt[i][0], cnt[i][1], cost
		if cl == 0 {
			continue
		}
		for j, cut := 1, n-cl*2+1; cut > 0; j++ {
			k := sort.SearchInts(el[j], i)
			if cut < k {
				c += cut * j
				break
			}
			c += k * j
			cut -= k
		}
		if c < ans {
			ans = c
		}
		n -= cl
		cost += s
	}
	Fprint(out, ans)
}

//func main() { CF557C(os.Stdin, os.Stdout) }
