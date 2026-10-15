package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf909F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n int
	Fscan(in, &n)
	if n&1 != 0 {
		Fprintln(out, "NO")
	} else {
		Fprintln(out, "YES")
		var pr func(int)
		pr = func(n int) {
			for i := 20; i > 0; i-- {
				if n>>i&1 > 0 {
					x := 1<<(i+1) - 1 ^ n
					if x > 1 {
						pr(x - 1)
					}
					for ; n >= x; n-- {
						Fprint(out, n, " ")
					}
					return
				}
			}
		}
		pr(n)
		Fprintln(out)
	}

	if n < 6 || n&(n-1) == 0 {
		Fprintln(out, "NO")
		return
	}
	Fprintln(out, "YES")
	a := make([]int, n+2)
	a[1], a[2], a[3], a[4], a[5], a[6] = 3, 6, 2, 5, 1, 4
	for i := 7; i <= n; i++ {
		if i == i&-i {
			a[i], a[i+1] = i+1, i
			i++
		} else {
			a[i] = a[i&-i]
			a[i&-i] = i
		}
	}
	for i := 1; i <= n; i++ {
		Fprint(out, a[i], " ")
	}
}

//func main() { cf909F(bufio.NewReader(os.Stdin), os.Stdout) }
