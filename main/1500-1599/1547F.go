package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1547F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	gcd := func(a, b int) int {
		for a != 0 {
			a, b = b%a, a
		}
		return b
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			Fscan(in, &a[i])
		}
		ans := 0
		for i, x := range a {
			y := a[(i+1)%n]
			j := i
			for ; x != y; j++ {
				x = gcd(x, a[(j+1)%n])
				y = gcd(y, a[(j+2)%n])
			}
			if j-i > ans {
				ans = j - i
			}
		}
		Fprintln(out, ans)
	}
}

//func main() { CF1547F(os.Stdin, os.Stdout) }
