package main

import (
	. "fmt"
	"io"
	"slices"
)

// https://github.com/EndlessCheng
func cf28D(in io.Reader, out io.Writer) {
	var n, mx, mxI int
	Fscan(in, &n)
	f := map[[2]int][2]int{}
	from := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var v, c, l, r int
		Fscan(in, &v, &c, &l, &r)
		fv, ok := f[[2]int{l, r + c}]
		if l > 0 && !ok {
			continue
		}
		from[i] = fv[1]
		nf := v + fv[0]
		p := [2]int{l + c, r}
		if nf > f[p][0] {
			f[p] = [2]int{nf, i}
		}
		if r == 0 && nf > mx {
			mx, mxI = nf, i
		}
	}

	ans := []any{}
	for i := mxI; i > 0; i = from[i] {
		ans = append(ans, i)
	}
	slices.Reverse(ans)
	Fprintln(out, len(ans))
	Fprintln(out, ans...)
}

//func main() { cf28D(bufio.NewReader(os.Stdin), os.Stdout) }
