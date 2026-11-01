package main

import (
	. "fmt"
	"io"
)

func cf196D(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	var m int
	var str string
	Fscan(in, &m, &str)
	n := len(str)
	s := make([]byte, n+1)
	copy(s[1:], str)
	i := n
	for i > 0 && s[i] == 'z' {
		s[i] = 'a'
		i--
	}
	if i == 0 {
		Fprint(out, "Impossible")
		return
	}
	s[i]++

	pw := make([]int, n+1)
	pw[0] = 1
	for i = 1; i <= n; i++ {
		pw[i] = pw[i-1] * mod
	}

	a := make([]int, n+1)
	b := make([]int, n+1)

	check := func(x, y int) bool {
		if x < 0 {
			return true
		}
		return (a[y]-a[x]*pw[y-x])*pw[x] != b[y]-b[x]
	}

	ans := make([]byte, n+1)

	var dfs func(int, bool) bool
	dfs = func(x int, w bool) bool {
		if x > n {
			Fprintf(out, "%s", ans[1:])
			return true
		}
		start := byte('a')
		if w {
			start = s[x]
		}
		for ans[x] = start; ans[x] <= 'z'; ans[x]++ {
			a[x] = a[x-1]*mod + int(ans[x])
			b[x] = b[x-1] + pw[x-1]*int(ans[x])
			if check(x-m, x) && check(x-m-1, x) && dfs(x+1, w && ans[x] == s[x]) {
				return true
			}
		}
		return false
	}

	if !dfs(1, true) {
		Fprint(out, "Impossible")
	}
}

//func main() { cf196D(bufio.NewReader(os.Stdin), os.Stdout) }
