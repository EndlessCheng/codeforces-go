package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol220B(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, q int
	Fscan(in, &n, &q)
	a := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
	}
	qs := make([][4]int, q)
	blockSize := int(math.Round(math.Sqrt(float64(n))))
	for i := range qs {
		var ql, qr int
		Fscan(in, &ql, &qr)
		// 分块，对于每一块，左端点分配在一个较小的范围，并且按照右端点从小到大排序，
		// 这样对于每一块，指针移动的次数为 O(√n*√n+n) = O(n)
		// 此外，记录的是 [l,r)，这样能稍微简化后续代码
		qs[i] = [4]int{ql / blockSize, qr + 1, ql, i}
	}
	sort.Slice(qs, func(i, j int) bool {
		for k := range qs[i] {
			if qs[i][k] != qs[j][k] {
				return qs[i][k] < qs[j][k]
			}
		}
		return true
	})

	cnt := make([]int, n+1)
	nowCnt := 0
	update := func(val, delta int) {
		if val > n {
			return
		}
		if cnt[val] == val {
			nowCnt--
		}
		cnt[val] += delta
		if cnt[val] == val {
			nowCnt++
		}
	}
	ans := make([]int, q)
	l, r := 1, 1
	for _, tuple := range qs {
		ql, qr, i := tuple[2], tuple[1], tuple[3]
		for ; l < ql; l++ {
			update(a[l], -1)
		}
		for ; r < qr; r++ {
			update(a[r], 1)
		}
		for l > ql {
			l--
			update(a[l], 1)
		}
		for r > qr {
			r--
			update(a[r], -1)
		}
		ans[i] = nowCnt
	}
	for _, v := range ans {
		Fprintln(out, v)
	}
}

//func main() {
//	Sol220B(os.Stdin, os.Stdout)
//}
