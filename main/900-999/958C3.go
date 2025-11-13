package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf958C3(in io.Reader, out io.Writer) {
	var n, k, p int
	Fscan(in, &n, &k, &p)
	s := make([]int, n+1)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] += s[i-1]
		f[i] = 1e9
	}

	for range k {
		mn, mnJ := f[0], 0
		for j := 1; j <= n; j++ {
			tmp := f[j]
			f[j] = mn + (s[j]-s[mnJ])%p
			if tmp < mn {
				mn, mnJ = tmp, j
			}
		}
		f[0] = 1e9
	}
	Fprint(out, f[n])
}

//func main() { cf958C3(bufio.NewReader(os.Stdin), os.Stdout) }
