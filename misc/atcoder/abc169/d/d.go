package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	primeFactorization := func(n int) (factors [][2]int) {
		for i := 2; i*i <= n; i++ {
			e := 0
			for ; n%i == 0; n /= i {
				e++
			}
			if e > 0 {
				factors = append(factors, [2]int{i, e})
			}
		}
		if n > 1 {
			factors = append(factors, [2]int{n, 1})
		}
		return
	}

	var n int
	Fscan(in, &n)
	if n == 1 {
		Fprintln(out, 0)
		return
	}
	ans := 0
	for _, f := range primeFactorization(n) {
		e := f[1]
		for i := 1; i <= e; i++ {
			ans++
			e -= i
		}
	}
	Fprint(out, ans)
}

func main() { run(os.Stdin, os.Stdout) }
