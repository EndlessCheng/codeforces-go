package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1156D(reader io.Reader, writer io.Writer) {
	var find func([]int, int) int
	find = func(fa []int, i int) int {
		if fa[i] != i {
			fa[i] = find(fa, fa[i])
		}
		return fa[i]
	}
	merge := func(fa []int, size []int, from, to int) {
		from, to = find(fa, from), find(fa, to)
		if from != to {
			fa[from] = to
			size[to] += size[from]
		}
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	fa := [2][]int{}
	size := [2][]int{}
	for i := range fa {
		fa[i] = make([]int, n)
		for j := range fa[i] {
			fa[i][j] = j
		}
		size[i] = make([]int, n)
		for j := range size[i] {
			size[i][j] = 1
		}
	}
	for m := n - 1; m > 0; m-- {
		var v, w, c int
		Fscan(in, &v, &w, &c)
		merge(fa[c], size[c], v-1, w-1)
	}

	ans := int64(0)
	for i := range fa[0] {
		// 考察顶点 v，假设通过 0-边与其相连的点一共有 x 个（包括 v），通过 1-边与其相连的有 y 个（包括 v），
		// 那么以 x 个点为起点，以 y 个点为终点的路径条数 x*y-1 即为顶点 v 对答案的贡献（-1 是因为 v 到 v 不算）。
		ans += int64(size[0][find(fa[0], i)])*int64(size[1][find(fa[1], i)]) - 1
	}
	Fprint(out, ans)
}

//func main() {
//	Sol1156D(os.Stdin, os.Stdout)
//}
