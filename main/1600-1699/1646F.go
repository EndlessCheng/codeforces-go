package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1646F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var n, tot int
	Fscan(in, &n)
	a := make([][]int, n+1)
	for i := range a {
		a[i] = make([]int, n+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var x int
			Fscan(in, &x)
			a[i][x]++
		}
	}

	op := make([][]int, n*n)
	for i := range op {
		op[i] = make([]int, n+1)
	}

	calc := func(x, y int) {
		tot++
		for {
			for i := 1; i <= n; i++ {
				if a[x][i] > 1 {
					op[tot-1][x] = i
					a[x][i]--
					a[x%n+1][i]++
					break
				}
			}
			x = x%n + 1
			if x == y {
				break
			}
		}
	}

	var f func()
	f = func() {
		for i := 1; i <= n; i++ {
			for j := 1; j <= n; j++ {
				if a[i][j] > 1 {
					calc(i, i)
					f()
					return
				}
			}
		}
	}
	f()

	Fprintln(out, tot+(n*n-n)/2)
	for i := range tot {
		for j := 1; j <= n; j++ {
			Fprint(out, op[i][j], " ")
		}
		Fprintln(out)
	}
	for i := 2; i <= n; i++ {
		for j := i; j >= 2; j-- {
			for k := 1; k <= n; k++ {
				Fprint(out, (j+k-2)%n+1, " ")
			}
			Fprintln(out)
		}
	}
}

//func main() { cf1646F(bufio.NewReader(os.Stdin), os.Stdout) }
