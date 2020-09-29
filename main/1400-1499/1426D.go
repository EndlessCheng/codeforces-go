package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1426D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, ans, v int
	Fscan(in, &n)
	sum := make([]int64, n+1)
	mp := map[int64]bool{0: true}
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		sum[i+1] = sum[i] + int64(v)
		if mp[sum[i+1]] {
			ans++
			mp = map[int64]bool{sum[i]: true}
		}
		mp[sum[i+1]] = true
	}
	Fprint(out, ans)
}

//func main() { CF1426D(os.Stdin, os.Stdout) }
