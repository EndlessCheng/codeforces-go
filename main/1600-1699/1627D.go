package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1627D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var n, v, ans int
	Fscan(in, &n)
	vis := [1e6 + 1]bool{}
	for ; n > 0; n-- {
		Fscan(in, &v)
		vis[v] = true
	}
	for i := 1; i <= 333333; i++ {
		if !vis[i] {
			g := 0
			for j := i * 2; j <= 1e6; j += i {
				if vis[j] {
					g = gcd(g, j)
				}
			}
			if g == i {
				ans++
			}
		}
	}
	Fprint(out, ans)
}

//func main() { CF1627D(os.Stdin, os.Stdout) }
