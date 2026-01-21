package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1646D(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	if n == 2 {
		Fprint(out, "2 2\n1 1")
		return
	}

	g := make([][]int, n)
	for range n - 1 {
		var v, w int
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	f := make([][4]int, n)
	var dfs func(int, int) (int, int, int, int)
	dfs = func(v, fa int) (notChosen, chosen, notChosenS, chosenS int) {
		chosen = 1
		notChosenS = 1
		chosenS = len(g[v])
		for _, w := range g[v] {
			if w == fa {
				continue
			}
			nc, c, ncs, cs := dfs(w, v)
			if nc > c {
				notChosen += nc
				notChosenS += ncs
			} else if nc < c {
				notChosen += c
				notChosenS += cs
			} else {
				notChosen += c
				notChosenS += min(ncs, cs)
			}
			chosen += nc
			chosenS += ncs
		}
		f[v] = [4]int{notChosen, chosen, notChosenS, chosenS}
		return
	}
	nc, c, ncs, cs := dfs(0, -1)

	ans := make([]any, n)
	var makeAns func(int, int, bool)
	makeAns = func(v, fa int, chosen bool) {
		if chosen {
			ans[v] = len(g[v])
		} else {
			ans[v] = 1
		}
		for _, w := range g[v] {
			if w != fa {
				makeAns(w, v, !chosen && (f[w][0] < f[w][1] || f[w][0] == f[w][1] && f[w][2] > f[w][3]))
			}
		}
		return
	}
	choose := nc < c || nc == c && ncs > cs
	makeAns(0, -1, choose)

	tot := ncs
	if choose {
		tot = cs
	}
	Fprintln(out, max(nc, c), tot)
	Fprintln(out, ans...)
}

//func main() { cf1646D(bufio.NewReader(os.Stdin), os.Stdout) }
