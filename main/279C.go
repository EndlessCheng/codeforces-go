package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol279C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n, q, l, r int
	Fscan(in, &n, &q)
	a := make([]int, n+2)
	sameL := make([]int, n+1)
	sameLI := 0
	prev := make([]int, n+1)
	isAsc := true
	changeI := 0
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		if a[i] != a[i-1] {
			sameLI = i
		}
		sameL[i] = sameLI
		if a[i-1] < a[i] && !isAsc {
			changeI = i - 1
			isAsc = true
		} else if a[i-1] > a[i] && isAsc {
			changeI = i - 1
			isAsc = false
		}
		prev[i] = changeI
	}

	next := make([]int, n+1)
	isAsc = false
	changeI = n + 1
	for i := n; i >= 1; i-- {
		if a[i] < a[i+1] && !isAsc {
			changeI = i + 1
			isAsc = true
		} else if a[i] > a[i+1] && isAsc {
			changeI = i + 1
			isAsc = false
		}
		next[i] = changeI
	}

	for ; q > 0; q-- {
		Fscan(in, &l, &r)
		if next[l] >= r || prev[r] <= l || sameL[next[l]] == sameL[prev[r]] && a[next[l]] >= a[l] {
			Fprintln(out, "Yes")
		} else {
			Fprintln(out, "No")
		}
	}
}

//func main() {
//	Sol279C(os.Stdin, os.Stdout)
//}
