package main

import (
	. "fmt"
	"io"
)

func cf437C(in io.Reader, out io.Writer) {
	var n, m, v, w, ans int
	Fscan(in, &n, &m)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for ; m > 0; m-- {
		Fscan(in, &v, &w)
		ans += min(a[v-1], a[w-1])
	}
	Fprint(out, ans)
}

//func main() { cf437C(bufio.NewReader(os.Stdin), os.Stdout) }
