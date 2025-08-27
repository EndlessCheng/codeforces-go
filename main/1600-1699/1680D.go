package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1680D(in io.Reader, out io.Writer) {
	var n, k, ans int
	Fscan(in, &n, &k)
	S := make([]int, n+1)
	Z := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &S[i])
		if S[i] == 0 {
			Z[i] = k
		}
		S[i] += S[i-1]
		Z[i] += Z[i-1]
	}
	if S[n] > Z[n] || S[n] < -Z[n] {
		Fprint(out, -1)
		return
	}
	for i := range n {
		for j := i + 1; j <= n; j++ {
			s := S[j] - S[i]
			z := Z[j] - Z[i]
			cz := Z[n] - z
			ans = max(ans, min(z, cz-S[n])+s, min(z, cz+S[n])-s)
		}
	}
	Fprint(out, ans+1)
}

//func main() { cf1680D(bufio.NewReader(os.Stdin), os.Stdout) }
