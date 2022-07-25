package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF1467C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, min [3]int
	s := [3]int64{}
	Fscan(in, &n[0], &n[1], &n[2])
	var v int
	for i, n := range n {
		min[i] = 1e9
		for j := 0; j < n; j++ {
			Fscan(in, &v)
			if v < min[i] { min[i] = v }
			s[i] += int64(v)
		}
	}
	sort.Slice(s[:], func(i, j int) bool { return s[i] < s[j] })
	ans := s[1] + s[2] - s[0]
	sort.Ints(min[:])
	s2 := s[0] + s[1] + s[2] - int64(min[0]+min[1])*2
	if s2 > ans { ans = s2 }
	Fprint(out, ans)
}

//func main() { CF1467C(os.Stdin, os.Stdout) }
