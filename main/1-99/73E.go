package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf73E(in io.Reader, out io.Writer) {
	var n, x, v, ans int
	Fscan(in, &n, &x)
	if x == 2 {
		Fprint(out, 0)
		return
	}

	const mx int = 2e6
	has := [mx]bool{}
	for range n {
		Fscan(in, &v)
		if v == 1 {
			Fprint(out, 1)
			return
		}
		if v < mx {
			has[v] = true
		}
	}

	var np = [mx]bool{}
	for i := 2; i < x; i++ {
		if np[i] {
			continue
		}
		if !has[i] {
			Fprint(out, -1)
			return
		}
		ans++
		for j := i * i; j < mx; j += i {
			np[j] = true
		}
	}
	Fprint(out, ans)
}

//func main() { cf73E(bufio.NewReader(os.Stdin), os.Stdout) }
