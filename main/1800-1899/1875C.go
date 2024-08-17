package main

import (
	. "fmt"
	"io"
)

func cf1875C(in io.Reader, out io.Writer) {
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}
	var T, n, m int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &m)
		n %= m
		g := gcd(n, m)
		n /= g
		m /= g
		if m&(m-1) > 0 {
			Fprintln(out, -1)
			continue
		}
		ans := 0
		for n > 0 {
			ans += n
			n = n * 2 % m
		}
		Fprintln(out, ans*g)
	}
}

//func main() { cf1875C(bufio.NewReader(os.Stdin), os.Stdout) }
