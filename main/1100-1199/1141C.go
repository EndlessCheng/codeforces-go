package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1141C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, minS, maxS int
	Fscan(in, &n)
	s := make([]int, n)
	seen := map[int]bool{0: true}
	for i := 1; i < n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
		v := s[i]
		if seen[v] {
			Fprint(out, -1)
			return
		}
		seen[v] = true
		minS = min(minS, v)
		maxS = max(maxS, v)
	}
	if maxS-minS != n-1 {
		Fprint(out, -1)
		return
	}
	for _, v := range s {
		Fprint(out, v+1-minS, " ")
	}
}

//func main() { cf1141C(bufio.NewReader(os.Stdin), os.Stdout) }
