package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf963C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var n, w, h, c, g int
	Fscan(in, &n)
	cntW := map[int]int{}
	setW := map[int]bool{}
	for range n {
		Fscan(in, &w, &h, &c)
		g = gcd(g, c)
		cntW[h]++
		setW[w] = true
	}

	if g == 1 {
		Fprint(out, 0)
		return
	}

	for _, c := range cntW {
		if c != len(setW) {
			Fprint(out, 0)
			return
		}
	}

	ans := 0
	for i := 1; i*i <= g; i++ {
		if g%i == 0 {
			ans++
			if i*i != g {
				ans++
			}
		}
	}
	Fprintln(out, ans)
}

//func main() { cf963C(bufio.NewReader(os.Stdin), os.Stdout) }
