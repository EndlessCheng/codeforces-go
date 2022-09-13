package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func CF1494C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	max := func(a, b int) int {
		if b > a {
			return b
		}
		return a
	}

	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := make([]int, m)
		for i := range b {
			Fscan(in, &b[i])
		}
		f := func(a, b []int) int {
			same, i, n := 0, 0, len(a)
			for _, v := range b {
				for i < n && a[i] < v {
					i++
				}
				if i < n && a[i] == v {
					same++
					i++
				}
			}
			res := same
			i, left := 0, 0
			for right, v := range b {
				for i < n && a[i] < v {
					i++
				}
				if i < n && a[i] == v {
					same--
					i++
				}
				for left <= right && v-b[left]+1 > i { // b[left] 到 b[right] 的特殊位置无法被 i 个连续箱子覆盖
					left++
				}
				res = max(res, right-left+1+same)
			}
			return res
		}
		x, y := sort.SearchInts(a, 0), sort.SearchInts(b, 0)
		ans := f(a[x:], b[y:])
		// 负数变正数且递增，方便复用 f
		for i := 0; i < (x+1)/2; i++ {
			a[i], a[x-1-i] = -a[x-1-i], -a[i]
		}
		for i := 0; i < (y+1)/2; i++ {
			b[i], b[y-1-i] = -b[y-1-i], -b[i]
		}
		ans += f(a[:x], b[:y])
		Fprintln(out, ans)
	}
}

//func main() { CF1494C(os.Stdin, os.Stdout) }
