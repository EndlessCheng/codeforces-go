package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf1622F(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n int
	Fscan(in, &n)

	pr := func(x, y int) {
		ans := n - 1
		if x != y {
			ans--
		}
		Fprintln(out, ans)
		for i := 1; i <= n; i++ {
			if i != x && i != y {
				Fprint(out, i, " ")
			}
		}
		Fprintln(out)
	}

	if n == 1 {
		Fprintln(out, 1)
		Fprintln(out, 1)
		return
	}

	t := n / 2
	if n%2 > 0 {
		m := n * (t - 1)
		k := int(math.Sqrt(float64(m)))
		if k*k == m {
			pr(t-2, n-2)
			return
		}
		n--
	}
	if t%2 > 0 {
		x := 0
		k := (t + 1) / 2
		s := int(math.Sqrt(float64(k)))
		if s*s == k {
			x = t + 1
		} else if t == 9 {
			x = 7
		}
		if x > 0 {
			pr(x, x)
		} else {
			pr(2, t)
		}
	} else {
		pr(t, t)
	}
}

//func main() { cf1622F(bufio.NewReader(os.Stdin), os.Stdout) }
