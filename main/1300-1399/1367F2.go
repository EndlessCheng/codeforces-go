package main

import (
	. "fmt"
	"io"
	"slices"
	"sort"
)

// https://github.com/EndlessCheng
func cf1367F2(in io.Reader, out io.Writer) {
	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		b := slices.Clone(a)
		slices.Sort(b)
		b = slices.Compact(b)

		m := len(b)
		tot := make([]int, m)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v)
			tot[a[i]]++
		}

		// f[v] 表示当前以值 v 结尾的最长子序列的长度
		f := make([]int, m)
		// f2[v] 记录以一整块以值 v 结尾的最长子序列的长度
		// 「一整块」指的是子序列中包含了数组中所有的 v
		// 只有当所有 v 都遍历到，这个状态才有意义，表示一个「已完成」的、可以被 v+1 安全转移的状态
		full := make([]int, m)
		cnt := make([]int, m)
		for _, v := range a {
			if v > 0 {
				if cnt[v-1] == tot[v-1] {
					f[v] = max(f[v], full[v-1])
				} else {
					f[v] = max(f[v], cnt[v-1])
				}
			}
			f[v]++ // 虽然在 cnt[v-1] 的基础上 +1 不一定对，但这不是最优的
			if cnt[v] > 0 {
				full[v]++
			} else {
				full[v] = f[v]
			}
			cnt[v]++
		}
		Fprintln(out, n-slices.Max(f))
	}
}

//func main() { cf1367F2(bufio.NewReader(os.Stdin), os.Stdout) }
