package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"math/rand"
	"strconv"
	"testing"
)

// https://codeforces.com/problemset/problem/1320/E
// https://codeforces.com/problemset/status/1320/problem/E
func TestCF1320E(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
7
1 2
1 3
2 4
2 5
3 6
3 7
3
2 2
4 1
7 1
1 3
2 2
4 3
7 1
1 3
3 3
1 1
4 100
7 100
1 2 3
outputCopy
1 2
1 1
1 1 1
inputCopy
2
2 1
2
2 2
1 1
2 1
2 1
2 2
1 1
2 1
2 1
outputCopy
2 1
2 1
inputCopy
3
1 2
1 3
2
3 2
3 1
1 1
2 1
1 2
1 1
3 1
2
outputCopy
2 3
1`
	testutil.AssertEqualCase(t, rawText, 0, CF1320E)
}

func TestCompareCF1320E(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		n := rg.Int(1, 5)
		rg.NewLine()
		rg.TreeEdges(n, 1)
		q := rg.Int(2, 5) // todo
		rg.NewLine()
		for ; q > 0; q-- {
			k := rg.Int(1, n)
			m := rg.Int(1, n)
			rg.NewLine()
			p := rand.Perm(n)[:k]
			for i := range p {
				rg.Bytes(strconv.Itoa(p[i]+1) + " 1\n")
			}
			rg.UniqueSlice(m, 1, n)
		}
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		var n, q int
		Fscan(in, &n)
		g := make([][]int, n)
		for i := 1; i < n; i++ {
			var v, w int
			Fscan(in, &v, &w)
			v--
			w--
			g[v] = append(g[v], w)
			g[w] = append(g[w], v)
		}

		for Fscan(in, &q); q > 0; q-- {
			var k, m int
			Fscan(in, &k, &m)
			ans := make([]int, n)
			minT := make([]int, n)
			for i := range ans {
				ans[i] = -1
				minT[i] = 1e9
			}
			for i := 1; i <= k; i++ {
				var st, s int
				Fscan(in, &st, &s)
				st--
				vis := make([]bool, len(g))
				vis[st] = true
				type pair struct{ v, d int }
				q := []pair{{st, 0}}
				for len(q) > 0 {
					p := q[0]
					q = q[1:]
					v, d := p.v, p.d
					t := (d + s - 1) / s
					if t < minT[v] {
						ans[v] = i
						minT[v] = t
					}
					for _, w := range g[v] {
						if !vis[w] {
							vis[w] = true
							q = append(q, pair{w, d + 1})
						}
					}
				}
			}
			for ; m > 0; m-- {
				var v int
				Fscan(in, &v)
				v--
				Fprint(out, ans[v], " ")
			}
			Fprintln(out)
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1320E)
}
