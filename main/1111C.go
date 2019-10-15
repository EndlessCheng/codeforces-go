package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func Sol1111C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, k int
	var a, b int64
	Fscan(in, &n, &k, &a, &b)
	arr := make([]int, k)
	for i := range arr {
		Fscan(in, &arr[i])
	}
	sort.Ints(arr)

	var f func(l, r int) int64
	f = func(l, r int) int64 {
		// 改成手写二分能快一倍！
		// 非手写二分 514ms https://codeforces.com/contest/1111/submission/62602788
		// 　手写二分 249ms https://codeforces.com/contest/1111/submission/62603480
		li := sort.Search(k, func(i int) bool { return arr[i] >= l })
		ri := sort.Search(k, func(i int) bool { return arr[i] >= r+1 })
		// 也可以写成下面这样，可能会快一点点(<10%)
		//ri := sort.Search(k-li, func(i int) bool { return arr[i+li] >= r+1 }) + li
		if li == ri {
			return a
		}
		ans := b * int64(ri-li) * int64(r-l+1)
		if l < r {
			mid := (l + r) >> 1
			if newAns := f(l, mid) + f(mid+1, r); newAns < ans {
				ans = newAns
			}
		}
		return ans
	}
	Fprint(out, f(1, 1<<uint(n)))
}

//func main() {
//	Sol1111C(os.Stdin, os.Stdout)
//}
