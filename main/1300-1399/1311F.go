package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1311F(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	type pair struct{ v, x int }

	var n int
	Fscan(in, &n)
	x := make([]int, n)
	ps := make([]pair, n)
	for i := range x {
		Fscan(in, &x[i])
	}
	for i := range ps {
		Fscan(in, &ps[i].v)
		ps[i].x = x[i]
	}
	sort.Ints(x)
	sort.Slice(ps, func(i, j int) bool { a, b := ps[i], ps[j]; return a.v < b.v || a.v == b.v && a.x < b.x })

	ans := int64(0)
	rank := map[int]int{}
	for i, v := range x {
		ans += int64(v) * int64(2*i+1-n) // 先计算 Σ|ai-aj| 的结果
		rank[v] = i
	}
	for i, p := range ps {
		ans -= int64(p.x) * int64(rank[p.x]-i) // 若 rank 发生了变动，则修改相应贡献
	}
	Fprint(out, ans)
}

//func main() { CF1311F(os.Stdin, os.Stdout) }
