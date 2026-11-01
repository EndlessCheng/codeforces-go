package main

import (
	. "fmt"
	"io"
)

func cf196D(in io.Reader, out io.Writer) {
	const mod = 127
	var m int
	var str string
	Fscan(in, &m, &str)
	n := len(str)
	p := make([]uint, n+1)
	h1 := make([]uint, n+1)
	h2 := make([]uint, n+1)
	p[0] = 1
	for i := 1; i <= n; i++ {
		p[i] = p[i-1] * mod
	}
	s := []byte(" " + str)
	ck := func(l, r int) bool {
		return (h1[r]-h1[l-1]*p[r-l+1])*p[l-1] == h2[r]-h2[l-1]
	}
	check := func(i int) bool {
		h1[i] = h1[i-1]*mod + uint(s[i])
		h2[i] = h2[i-1] + uint(s[i])*p[i-1]
		return i >= m && ck(i-m+1, i) || i >= m+1 && ck(i-m, i)
	}

	i := 1
	for ; i <= n; i++ {
		if check(i) {
			break
		}
	}

	ok := false
	for i = min(n, i); i > 0 && !ok; i-- {
		for s[i] = s[i] + 1; s[i] <= 'z' && !ok; s[i]++ {
			if !check(i) {
				ok = true
			}
		}
	}

	s[i+1]--
	if !ok {
		Fprintln(out, "Impossible")
		return
	}
	for i += 2; i <= n; i++ {
		for s[i] = 'a'; s[i] <= 'z'; s[i]++ {
			if !check(i) {
				break
			}
		}
	}
	Fprintf(out, "%s\n", s[1:n+1])
}

//func main() { cf196D(bufio.NewReader(os.Stdin), os.Stdout) }
