package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1175D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	Fscan(in, &n, &k)
	sum := make([]int64, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &sum[i])
		sum[i] += sum[i-1]
	}
	ans := int64(k) * sum[n]
	sum = sum[1 : len(sum)-1]
	sort.Slice(sum, func(i, j int) bool { return sum[i] < sum[j] })
	for _, s := range sum[:k-1] {
		ans -= s
	}
	Fprint(out, ans)
}

//func main() { CF1175D(os.Stdin, os.Stdout) }
