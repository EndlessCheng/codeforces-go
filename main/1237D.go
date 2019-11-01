package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1237D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	a := make([]int, 3*n)
	for i := range a[:n] {
		Fscan(in, &a[i])
	}
	copy(a[n:2*n], a[:n])
	copy(a[2*n:], a[:n])

	type pair struct{ a, i int }
	q := []pair{{a[0], 0}}
	ql, j := 0, 1
	for i := range a[:n] {
		// 确保 q 是单调递减的，这样最值就是 q[ql].a
		for ; j < 3*n && (ql == len(q) || 2*a[j] >= q[ql].a); j++ {
			for ql < len(q) && q[len(q)-1].a < a[j] {
				q = q[:len(q)-1]
			}
			q = append(q, pair{a[j], j})
		}
		ans := j - i
		if ans > 2*n {
			ans = -1
		}
		Fprint(out, ans, " ")
		if ql < len(q) && q[ql].i == i {
			ql++
		}
	}
}

//func main() {
//	Sol1237D(os.Stdin, os.Stdout)
//}
