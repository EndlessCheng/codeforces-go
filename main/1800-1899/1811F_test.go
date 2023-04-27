package main

import (
	. "fmt"
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"io"
	"testing"
)

// https://codeforces.com/problemset/problem/1811/F
// https://codeforces.com/problemset/status/1811/problem/F
func TestCF1811F(t *testing.T) {
	// just copy from website
	rawText := `
inputCopy
5

9 12
1 2
3 1
2 3
1 6
4 1
6 4
3 8
3 5
5 8
9 7
2 9
7 2

8 12
1 2
3 1
2 3
1 6
4 1
6 4
3 8
3 5
5 8
8 7
2 8
7 2

4 3
1 2
4 2
3 1

6 8
6 3
6 4
5 3
5 2
3 2
3 1
2 1
2 4

5 7
2 4
2 5
3 4
3 5
4 1
4 5
1 5
outputCopy
YES
NO
NO
NO
NO
inputCopy
4

2 1
1 2

8 9
1 2
8 4
8 2
6 4
6 5
4 7
3 2
3 7
2 5

9 12
2 9
2 8
6 9
6 8
6 5
6 1
9 8
9 3
9 1
8 3
8 7
5 7

3 3
1 2
1 3
2 3
outputCopy
NO
NO
NO
NO
inputCopy
1
16 20
10 2
3 10
3 9
2 9
2 12
12 7
15 7
15 4
4 12
12 11
11 5
5 14
16 14
11 16
11 1
8 1
6 8
13 6
1 13
1 2
outputCopy
YES`
	testutil.AssertEqualCase(t, rawText, -1, CF1811F)
}

func TestCompareCF1811F(_t *testing.T) {
	//return
	testutil.DebugTLE = 0

	inputGenerator := func() string {
		//return ``
		rg := testutil.NewRandGenerator()
		rg.One() // 若不是多测则 remove
		n := rg.Int(25,25)
		m := rg.Int(30,30)
		rg.NewLine()
		rg.GraphEdges(n, m, 1, false)
		return rg.String()
	}

	// 暴力算法
	runBF := func(in io.Reader, out io.Writer) {
		T := 1
		Fscan(in, &T)
		for Case := 1; Case <= T; Case++ {
			Fprintln(out, solveCF1811F(in))
		}
	}

	testutil.AssertEqualRunResultsInf(_t, inputGenerator, runBF, CF1811F)
}

func solveCF1811F(in io.Reader) string {
	var n, m int
	Fscan(in, &n, &m)

	// 建图，并记录每个节点的边数
	g := make([][]int, n)
	var u, v int
	edges := make([]int, n)
	for i := 0; i < m; i++ {
		Fscan(in, &u, &v)
		g[u-1] = append(g[u-1], v-1)
		g[v-1] = append(g[v-1], u-1)
		edges[u-1]++
		edges[v-1]++
	}
	// println(edges)

	k := 1
	for k*k < n {
		k++
	}
	// k至少要为3
	if k < 3 {
		return "NO"
	}
	// 节点个数需要为平方数
	if k*k != n {
		return "NO"
	}
	// 边的个数需要为节点个数 + k
	if m != k*(k+1) {
		return "NO"
	}

	// 找出边为4的节点，并且检查所有节点的边，节点的边只能是2或4
	deg4 := []int{}
	for i, x := range edges {
		if x == 4 {
			deg4 = append(deg4, i)
		} else if x != 2 {
			return "NO"
		}
	}

	// 检查所有边为4的节点是否能构成一个环
	vis := []int{}
	v = deg4[0]
	for {
		nxt := -1
		for _, u := range g[v] {
			// 只遍历度为4的节点
			if edges[u] == 2 {
				continue
			}
			// 已经访问过的节点
			if len(vis) > 0 && vis[len(vis)-1] == u {
				continue
			}
			nxt = u
			break
		}
		// 没找到下一个环节点
		if nxt == -1 {
			return "NO"
		}
		vis = append(vis, v)
		if nxt == deg4[0] {
			break
		}
		v = nxt
	}

	// 度为4的节点构成的环的长度要为k
	if len(vis) != k {
		return "NO"
	}

	// 从度为4的节点开始检查与度为2的节点是否能构成环，且长度为k
	for _, s := range deg4 {
		vis = []int{}
		u = s
		for {
			nxt := -1
			for _, v := range g[u] {
				// 如果是初始节点，需要跳过度为4的节点，否则因为初始节点的度也为4
				// 其它节点的话遍历到的下一个度为4的节点意味着回到了初始节点，不需要跳过
				if u == s && edges[v] == 4 {
					continue
				}
				// 如果是前一个节点，也需要跳过
				if len(vis) > 0 && v == vis[len(vis)-1] {
					continue
				}
				nxt = v
				break
			}
			// 没找到下一个节点，意味着不构成环
			if nxt == -1 {
				return "NO"
			}
			vis = append(vis, u)
			u = nxt
			// 找到一个度为4的节点则退出
			if edges[u] == 4 {
				break
			}
		}
		// 如果没回到初始节点
		if u != s {
			return "NO"
		}
		// 如果环的节点个数不为k
		if len(vis) != k {
			return "NO"
		}
	}

	return "YES"

}
