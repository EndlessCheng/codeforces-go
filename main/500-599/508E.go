package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF508E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, l, r int
	Fscan(in, &n)
	type pair struct{ l, r int }
	stk := []pair{}
	ans := make([]byte, 0, n*2)
	for ; n > 0; n-- {
		Fscan(in, &l, &r)
		stk = append(stk, pair{l + len(ans), r + len(ans)})
		ans = append(ans, '(')
		for len(stk) > 0 && stk[len(stk)-1].l <= len(ans) && len(ans) <= stk[len(stk)-1].r {
			stk = stk[:len(stk)-1]
			ans = append(ans, ')')
		}
	}
	if len(stk) > 0 {
		Fprint(out, "IMPOSSIBLE")
	} else {
		Fprintf(out, "%s", ans)
	}
}

//func main() { CF508E(os.Stdin, os.Stdout) }
