package main

import (
	. "fmt"
	"io"
)

func cf196D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var m int
	var s []byte
	Fscan(in, &m, &s)
	n := len(s)

	w := n - 1
	for w >= 0 && s[w] == 'z' {
		s[w] = 'a'
		w--
	}
	if w < 0 {
		Fprintln(out, "Impossible")
		return
	}
	s[w]++

	pw := make([]int, n+1)
	pw[0] = 1
	for i := 1; i <= n; i++ {
		pw[i] = pw[i-1] * mod
	}

	a := make([]int, n+1)
	b := make([]int, n+1)
	check := func(x, y int) bool {
		x++
		return x < y || (a[x]-a[x-y]*pw[y])*pw[x-y] != b[x]-b[x-y]
	}

	t := make([]byte, n+1)
	ok := false

	var dfs func(int, bool)
	dfs = func(x int, y bool) {
		if x == n {
			Fprintf(out, "%s", t[:n])
			ok = true
			return
		}
		st := byte('a')
		if y {
			st = s[x]
		}
		for t[x] = st; t[x] <= 'z'; t[x]++ {
			a[x+1] = a[x]*mod + int(t[x])
			b[x+1] = b[x] + int(t[x])*pw[x]
			if check(x, m) && check(x, m+1) {
				dfs(x+1, y && t[x] == s[x])
			}
			if ok {
				return
			}
		}
	}

	dfs(0, true)
	if !ok {
		Fprint(out, "Impossible")
	}
}

//func main() { cf196D(bufio.NewReader(os.Stdin), os.Stdout) }
