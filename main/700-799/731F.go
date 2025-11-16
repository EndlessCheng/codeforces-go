package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf731F(in io.Reader, out io.Writer) {
	var n, v, ans int
	Fscan(in, &n)
	const mx = 200001
	s := [mx]int{}
	for range n {
		Fscan(in, &v)
		s[v]++
	}
	for i := 2; i < mx; i++ {
		s[i] += s[i-1]
	}

	for x := 1; x < mx; x++ {
		if s[x] == s[x-1] {
			continue
		}
		res := 0
		for j := x; j < mx; j += x {
			res += (s[min(j+x, mx)-1] - s[j-1]) * j
		}
		ans = max(ans, res)
	}
	Fprint(out, ans)
}

//func main() { cf731F(bufio.NewReader(os.Stdin), os.Stdout) }
