package main

import (
	. "fmt"
	"io"
)

func cf633F(in io.Reader, out io.Writer) {
	var n, ans int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	g := make([][]int, n)
	for i := 1; i < n; i++ {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	subAns := make([]int, n)
	nodes := make([]struct{ fi, se, th, fiW, seW int }, n)
	var dfs func(int, int) int
	dfs = func(v, fa int) int {
		val := a[v]
		subAns[v] = val
		maxS := val
		maxSubAnsW := 0
		p := &nodes[v]
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			s := dfs(w, v)
			ans = max(ans, maxSubAnsW+subAns[w])    // 两个子树 w 中的最大路径和之和
			maxSubAnsW = max(maxSubAnsW, subAns[w]) // 子树 w 中的最大路径和
			subAns[v] = max(subAns[v], maxS+s)      // 在 v 拐弯的最大路径和
			maxS = max(maxS, s+val)                 // 子树 v 中的最大链和
			if s > p.fi {
				p.th = p.se
				p.se = p.fi
				p.seW = p.fiW
				p.fi = s
				p.fiW = w
			} else if s > p.se {
				p.th = p.se
				p.se = s
				p.seW = w
			} else if s > p.th {
				p.th = s
			}
		}
		subAns[v] = max(subAns[v], maxSubAnsW)
		return maxS
	}
	dfs(0, -1)

	var reroot func(int, int, int)
	reroot = func(v, fa, mxFa int) {
		val := a[v]
		p := nodes[v]
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			if w == p.fiW {
				// 子树 w 中的最大路径和 + val + 在 v 拐弯的两条最大链和
				ans = max(ans, subAns[w]+val+p.se+max(p.th, mxFa))
				reroot(w, v, val+max(p.se, mxFa))
			} else {
				s := p.se
				if w == p.seW {
					s = p.th
				}
				ans = max(ans, subAns[w]+val+p.fi+max(s, mxFa))
				reroot(w, v, val+max(p.fi, mxFa))
			}
		}
	}
	reroot(0, -1, 0)
	Fprint(out, ans)
}

//func main() { cf633F(bufio.NewReader(os.Stdin), os.Stdout) }
