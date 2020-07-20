package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

func TestCF1385F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
1
7 2
3 1
4 5
3 6
7 4
1 2
1 4
outputCopy
3
inputCopy
1 
9 2 
2 1
3 2
4 2
5 3
6 1
7 2
8 4
9 4
outputCopy
2
inputCopy
4
8 3
1 2
1 5
7 6
6 8
3 1
6 4
6 1
10 3
1 2
1 10
2 3
1 5
1 6
2 4
7 10
10 9
8 10
7 2
3 1
4 5
3 6
7 4
1 2
1 4
5 1
1 2
2 3
4 3
5 3
outputCopy
2
3
3
4`
	testutil.AssertEqualCase(t, rawText, 0, CF1385F)
}

// 无尽对拍
func Test2(t *testing.T) {
	//rand.Seed(time.Now().UnixNano())
	inputGenerator := func() string {
		rg := testutil.NewRandGenerator()
		rg.Int(1, 1) // t
		rg.NewLine()
		n := rg.Int(10,10)
		rg.Int(2, 2) // k
		rg.NewLine()
		rg.TreeEdges(n, 1)
		return rg.String()
	}

	var _permute func([]int, int, func())
	_permute = func(a []int, i int, do func()) {
		if i == len(a) {
			do()
			return
		}
		_permute(a, i+1, do)
		for j := i + 1; j < len(a); j++ {
			a[i], a[j] = a[j], a[i]
			_permute(a, i+1, do)
			a[i], a[j] = a[j], a[i]
		}
	}
	permuteAll := func(a []int, do func()) { _permute(a, 0, do) }

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		//return
		var t int
		for Fscan(in, &t); t > 0; t-- {
			var n, k int
			Fscan(in, &n, &k)
			g := make([][]int, n)
			for i := 1; i < n; i++ {
				var v, w int
				Fscan(in, &v, &w)
				v--
				w--
				g[v] = append(g[v], w)
				g[w] = append(g[w], v)
			}
			if k == 1 {
				Fprintln(out, n-1)
				continue
			}

			ans := 0
			a := make([]int, n)
			for i := range a {
				a[i] = i
			}
			permuteAll(a, func() {
				cnt := 0
				vis := make([]bool, n)
			o:
				for i := 0; i < n; i += k {
					// 是叶子且连接到同一个顶点
					same := -1
					for j := i; j-i < k; j++ {
						if j == n {
							break o
						}
						fa := 0
						for _, w := range g[a[j]] {
							if !vis[w] {
								if fa > 0 {
									break o
								}
								fa++
								if same == -1 {
									same = w
								} else if w != same {
									break o
								}
							}
						}
					}
					for j := i; j-i < k; j++ {
						vis[a[j]] = true
					}
					cnt++
				}
				if cnt > ans {
					ans = cnt
				}
			})
			Fprintln(out, ans)
		}
	}

	testutil.AssertEqualRunResultsInf(t, inputGenerator, runBF, CF1385F)
}
