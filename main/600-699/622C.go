package main

import (
	"bufio"
	. "fmt"
	"io"
)

func cf622C(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, m, l, r, v int
	Fscan(in, &n, &m)
	a := make([]struct{ v, l int }, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i].v)
		if a[i].v != a[i-1].v {
			a[i].l = i - 1
		} else {
			a[i].l = a[i-1].l
		}
	}
	for ; m > 0; m-- {
		Fscan(in, &l, &r, &v)
		if v != a[r].v {
			Fprintln(out, r)
		} else if l <= a[r].l {
			Fprintln(out, a[r].l)
		} else {
			Fprintln(out, -1)
		}
	}
}

//func main() { cf622C(bufio.NewReader(os.Stdin), os.Stdout) }
