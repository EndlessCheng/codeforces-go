package main

import (
	"bufio"
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func runE(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		u1 := (m + 1) / 2 * slices.Min(a)
		u2 := m / n * slices.Max(a)
		ans := sort.Search(min(u1, u2), func(low int) bool {
			// 二分最小的不满足要求的 low+1，即可得到最大的满足要求的 low
			low++
			left := m
			pre := 0
			for i, p := range a {
				k := (low-1)/p + 1 - pre // 还需要操作的次数
				if i == n-1 && k <= 0 {  // 最后一个数已经满足要求
					break
				}
				k = max(k, 1)   // 至少要走 1 步
				left -= k*2 - 1 // 左右横跳
				if left < 0 {
					return true
				}
				pre = k - 1 // 右边那个数已经操作 k-1 次
			}
			return false
		})
		Fprintln(out, ans)
	}
}

//func main() { runE(bufio.NewReader(os.Stdin), os.Stdout) }
