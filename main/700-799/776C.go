package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF776C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	var k, v, ans int64
	Fscan(in, &n, &k)
	s := make([]int64, n+1)
	for i := 0; i < n; i++ {
		Fscan(in, &v)
		s[i+1] = s[i] + v
	}
	a := []int64{1}
	if k == -1 {
		a = append(a, -1)
	} else if k != 1 {
		for p := k; -1e14 <= p && p <= 1e14; p *= k {
			a = append(a, p)
		}
	}
	for _, d := range a {
		cnt := map[int64]int{}
		for _, v := range s {
			ans += int64(cnt[v-d])
			cnt[v]++
		}
	}
	Fprint(out, ans)
}

//func main() { CF776C(os.Stdin, os.Stdout) }
