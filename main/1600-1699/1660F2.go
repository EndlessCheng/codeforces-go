package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1660F2(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		sum := make([]int, n+1)
		for i, b := range s {
			sum[i+1] = sum[i] + int(b&3) - 2
		}
		ans := int64(0)
		t := [3][]int{}
		for i := range t {
			t[i] = make([]int, n*2+2)
		}
		for i := n; i >= 0; i-- {
			j, k := (sum[i]%3+3)%3, sum[i]+n+1
			for k := k; k > 0; k &= k - 1 {
				ans += int64(t[j][k])
			}
			for ; k < len(t[j]); k += k & -k {
				t[j][k]++
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1660F2(os.Stdin, os.Stdout) }
