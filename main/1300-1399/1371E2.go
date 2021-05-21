package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1371E2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}

	var n, p int
	Fscan(in, &n, &p)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	sort.Ints(a)
	mi, mx := 0, int(2e9)
	for i, v := range a {
		// 下界：x+i>=a[i]，即打败一个怪前必须有 x+i 个糖果
		mi = max(mi, v-i)
		if i-p+1 >= 0 {
			// 上界：x+i-p+1<a[i]，即对于每个 i>=p-1 的 a[i]，其对应的位置 i-p+1 的糖果个数不能 >=a[i]，否则此刻在 i-p+1 位置上能选到 p 个怪物
			mx = min(mx, v-(i-p+1))
		}
	}
	Println(mi, mx)
	Fprintln(out, max(mx-mi, 0))
	for i := mi; i < mx; i++ {
		Fprint(out, i, " ")
	}
}

//func main() { CF1371E2(os.Stdin, os.Stdout) }
