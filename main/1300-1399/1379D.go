package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1379D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	type event struct{ e, i int }

	var n, m, k, x int
	Fscan(in, &n, &m, &m, &k)
	half := m / 2
	es := make([]event, 0, 2*n)
	has := map[int]bool{}
	for i := 1; i <= n; i++ {
		Fscan(in, &x, &x)
		if x >= half {
			x -= half
		}
		es = append(es, event{x<<1 | 1, i})
		if x+k < half {
			es = append(es, event{(x + k) << 1, i})
		} else if x+k > half {
			es = append(es, event{(x + k - half) << 1, i})
			has[i] = true
		}
	}

	mi := len(has)
	tmp := map[int]bool{}
	for i := range has {
		tmp[i] = true
	}

	sort.Slice(es, func(i, j int) bool { return es[i].e < es[j].e })
	for _, p := range es {
		if p.e&1 > 0 {
			has[p.i] = true
		} else {
			delete(has, p.i)
		}
		if len(has) < mi {
			mi = len(has)
		}
	}
	if mi == len(tmp) {
		Fprintln(out, mi, 0)
		for i := range tmp {
			Fprint(out, i, " ")
		}
		return
	}

	// 说明 mi < len(tmp)
	has = tmp
	for _, p := range es {
		if p.e&1 > 0 {
			has[p.i] = true
		} else if delete(has, p.i); len(has) == mi {
			Fprintln(out, mi, p.e>>1)
			for i := range has {
				Fprint(out, i, " ")
			}
			return
		}
	}
}

//func main() { CF1379D(os.Stdin, os.Stdout) }
