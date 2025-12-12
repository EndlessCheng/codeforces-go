package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1413D(in io.Reader, out io.Writer) {
	var n, v, now int
	var op string
	Fscan(in, &n)
	ans := make([]any, n)

	type pair struct{ low, i int }
	st := []pair{}
	for range n * 2 {
		Fscan(in, &op)
		if op == "+" {
			st = append(st, pair{0, now})
			now++
			continue
		}
		Fscan(in, &v)
		top := len(st) - 1
		if top < 0 || v < st[top].low {
			Fprint(out, "NO")
			return
		}
		ans[st[top].i] = v
		st = st[:top]
		if top > 0 {
			st[top-1].low = max(st[top-1].low, v)
		}
	}

	Fprintln(out, "YES")
	Fprintln(out, ans...)
}

//func main() { cf1413D(bufio.NewReader(os.Stdin), os.Stdout) }
