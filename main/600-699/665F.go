package main

import (
	. "fmt"
	"io"
	"math"
)

// https://github.com/EndlessCheng
func cf665F(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	m := int(math.Sqrt(float64(n)))
	pi := make([]int, m+1)
	pi2 := make([]int, m+1)
	for i := 1; i <= m; i++ {
		pi[i] = i - 1
		pi2[i] = n/i - 1
	}

	for i := 2; i <= m; i++ {
		prePi := pi[i-1]
		if pi[i] > prePi {
			for j := 1; j <= min(m, n/(i*i)); j++ {
				if i*j <= m {
					pi2[j] -= pi2[i*j] - prePi
				} else {
					pi2[j] -= pi[n/(i*j)] - prePi
				}
			}
			for j := m; j >= i*i; j-- {
				pi[j] -= pi[j/i] - prePi
			}
		}
	}

	ans := pi[int(math.Cbrt(float64(n)))]
	for i := 2; i <= m; i++ {
		if pi[i] > pi[i-1] {
			ans += pi2[i] - pi[i]
		}
	}
	Fprint(out, ans)
}

//func main() { cf665F(os.Stdin, os.Stdout) }
