package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf757B(in io.Reader, out io.Writer) {
	const mx int = 1e5
	cnt := [mx + 1]int{}
	var n, v int
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		cnt[v]++
	}
	ans := 1
	for i := 2; i <= mx; i++ {
		s := 0
		for j := i; j <= mx; j += i {
			s += cnt[j]
		}
		ans = max(ans, s)
	}
	Fprint(out, ans)
}

//func main() { cf757B(bufio.NewReader(os.Stdin), os.Stdout) }
