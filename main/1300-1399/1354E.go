package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1354E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m, n1, n2, n3, v, w int
	Fscan(in, &n, &m, &n1, &n2, &n3)
	g := make([][]int, n)
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		v--
		w--
		g[v] = append(g[v], w)
		g[w] = append(g[w], v)
	}

	color := make([]int8, n)
	cnt := [2]int{}
	vs := []int{}
	var f func(int, int8) bool
	f = func(v int, c int8) bool {
		color[v] = c
		cnt[c-1]++
		vs = append(vs, v)
		for _, w := range g[v] {
			if color[w] == c || color[w] == 0 && !f(w, 3^c) {
				return false
			}
		}
		return true
	}
	dp := make([]bool, n2+1) // dp[i][j] 表示能否从前 i 个连通块（二分图）中标记 j 个 2（第一维可以滚动掉）
	dp[0] = true
	from := make([][]int, n)
	for i := range from {
		from[i] = make([]int, n2+1)
	}
	cc := [][]int{}
	for i, c := range color {
		if c != 0 {
			continue
		}
		cnt = [2]int{}
		vs = []int{}
		if !f(i, 1) {
			Fprint(out, "NO")
			return
		}
	o:
		for j := n2; j >= 0; j-- { // 分组背包，每组有两个元素，对应当前连通块（二分图）的两部的大小
			for k, c := range cnt {
				if c <= j && dp[j-c] {
					dp[j] = true
					from[len(cc)][j] = (j-c)<<1 | k
					continue o
				}
			}
			dp[j] = false // 由于我们是滚动数组的写法，dp[i][j] 无法满足时要标记成 false
		}
		cc = append(cc, vs)
	}

	if !dp[n2] {
		Fprint(out, "NO")
		return
	}
	Fprintln(out, "YES")
	ans := bytes.Repeat([]byte{'3'}, n)
	for i, j := len(cc)-1, n2; i >= 0; i-- {
		j = from[i][j]
		tar := int8(j&1) + 1
		j >>= 1
		for _, v := range cc[i] {
			if color[v] == tar {
				n2--
				ans[v] = '2'
			} else if n1 > 0 {
				n1--
				ans[v] = '1'
			}
		}
	}
	Fprintf(out, "%s", ans)
}

//func main() { CF1354E(os.Stdin, os.Stdout) }
