package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf68D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	a := map[int]int{}
	subSum := map[int]int{}
	var dfs func(int, int) float64
	dfs = func(o, maxFaS int) float64 {
		if subSum[o] <= maxFaS { // 最优性剪枝
			return float64(maxFaS)
		}
		lv := dfs(o*2, max(maxFaS, a[o]+subSum[o*2+1])) // 当前节点 + 右子树
		rv := dfs(o*2+1, max(maxFaS, a[o]+subSum[o*2])) // 当前节点 + 左子树
		return (lv + rv) / 2
	}

	var q, v, e int
	var op string
	Fscan(in, &q, &q)
	for range q {
		Fscan(in, &op)
		if op[0] == 'd' {
			Fprintf(out, "%.4f\n", dfs(1, 0))
			continue
		}
		Fscan(in, &v, &e)
		a[v] += e
		for ; v > 0; v /= 2 {
			subSum[v] += e
		}
	}
}

//func main() { cf68D(bufio.NewReader(os.Stdin), os.Stdout) }
