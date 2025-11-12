package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf958C2(in io.Reader, out io.Writer) {
	var n, k, p int
	Fscan(in, &n, &k, &p)
	s := make([]int, n+1)
	f := make([]int, n+1)
	for i := 1; i <= n; i++ {
		Fscan(in, &s[i])
		s[i] = (s[i] + s[i-1]) % p
		f[i] = -1e9
	}

	mx := make([]int, p)
	for range k {
		for i := range mx {
			mx[i] = -1e9
		}
		mx[0] = f[0]
		for j := 1; j <= n; j++ {
			tmp := f[j]
			f[j] = -1e9
			for rem, fv := range mx {
				f[j] = max(f[j], fv+(s[j]-rem+p)%p)
			}
			mx[s[j]] = max(mx[s[j]], tmp)
		}
		f[0] = -1e9
	}
	Fprint(out, f[n])
}

//func main() { cf958C2(bufio.NewReader(os.Stdin), os.Stdout) }
