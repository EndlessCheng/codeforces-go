package main

import (
	"bufio"
	. "fmt"
	"io"
	"math/rand/v2"
	"slices"
)

// https://github.com/EndlessCheng
func cf995E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var u, v, p int
	Fscan(in, &u, &v, &p)

	e := []int{}
	find := func(z int) bool {
		e = e[:0]
		x := rand.IntN(p-1) + 1
		a := z * x % p
		b := x
		cnt := 0
		for a != 0 {
			cnt++
			if a < b {
				e = append(e, 3)
				a, b = b, a
			} else {
				e = append(e, 2)
				a -= b
			}
			if cnt > 100 {
				return false
			}
		}
		return true
	}

	for !find(u) {
	}
	ans := slices.Clone(e)

	for !find(v) {
	}
	for i := len(e) - 1; i >= 0; i-- {
		if e[i] == 3 {
			ans = append(ans, 3)
		} else {
			ans = append(ans, 1)
		}
	}

	Fprintln(out, len(ans))
	for _, v := range ans {
		Fprint(out, v, " ")
	}
	Fprintln(out)
}

//func main() { cf995E(bufio.NewReader(os.Stdin), os.Stdout) }
