package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1168A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, m int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	Fprint(out, sort.Search(m, func(lim int) bool {
		pre := 0
		for _, v := range a {
			// (pre-v+m)%m 表示把 v 改成 pre，需要的操作次数
			if (pre-v+m)%m > lim { // 无法修改成 pre
				if v < pre { // 无法保证单调非降
					return false
				}
				pre = v // 只能 v 不变了
			}
		}
		return true
	}))
}

//func main() { CF1168A(os.Stdin, os.Stdout) }
