package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, l, r, pre, cnt int
	Fscan(in, &n)
	a := make([]int, 0, n*2)
	for i := 0; i < n; i++ {
		Fscan(in, &l, &r)
		a = append(a, l<<1|1, (r+1)<<1) // 记录区间端点位置，用最低位表示左右
	}
	sort.Ints(a)

	ans := make([]int, n+1)
	for i := 0; i < n*2; { // 扫描所有区间端点
		v := a[i] >> 1
		ans[cnt] += v - pre
		for ; i < n*2 && a[i]>>1 == v; i++ {
			cnt += a[i]&1<<1 - 1 // 左+1，右-1
		}
		pre = v
	}
	for _, c := range ans[1:] {
		Fprint(out, c, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
