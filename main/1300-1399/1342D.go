package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1342D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, k int
	Fscan(in, &n, &k)
	sz := make([]int, n)
	for i := range sz {
		Fscan(in, &sz[i])
	}
	sort.Ints(sz)
	uppCnt := make([]int, k)
	for i := range uppCnt {
		Fscan(in, &uppCnt[i])
	}
	// 注意这个二分的总复杂度是 O(n) 的，如果上面的排序换成计数排序，整体的复杂度就是 O(n) 的
	ans := sort.Search(n+1, func(caseNum int) bool {
		cnt := 0
		for i := n - 1; i >= 0; i -= caseNum {
			cnt++
			if cnt > uppCnt[sz[i]-1] {
				return false
			}
		}
		return true
	})
	Fprintln(out, ans)
	for st := 0; st < ans; st++ {
		a := []interface{}{}
		for i := st; i < n; i += ans {
			a = append(a, sz[i])
		}
		Fprint(out, len(a), " ")
		Fprintln(out, a...)
	}
}

//func main() { CF1342D(os.Stdin, os.Stdout) }
