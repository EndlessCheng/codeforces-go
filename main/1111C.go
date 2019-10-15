package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
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
		li := sort.Search(k, func(i int) bool { return arr[i] >= l })
		ri := sort.Search(k, func(i int) bool { return arr[i] >= r+1 })
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

func main() {
	Sol1111C(os.Stdin, os.Stdout)
}
