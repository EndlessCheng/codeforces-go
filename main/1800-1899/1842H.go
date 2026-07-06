package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1842H(in io.Reader, out io.Writer) {
	const mod = 998244353
	const inv2 = (mod + 1) / 2
	var n, m int
	Fscan(in, &n, &m)

	inv := make([]int, n+1)
	s := [2][]int{make([]int, n), make([]int, n)}
	for range m {
		var t, i, j int
		Fscan(in, &t, &i, &j)
		i--
		j--
		s[t][i] |= 1 << j
		s[t][j] |= 1 << i
	}

	f := make([]int, 1<<n)
	f[0] = 1
	for i := range f {
		for j := range n {
			if i>>j&1 > 0 {
				continue
			}
			cnt := 0
			if i&s[0][j] == s[0][j] {
				cnt++
			}
			if i&s[1][j] == s[1][j] {
				cnt++
			}
			f[i|1<<j] = (f[i|1<<j] + f[i]*inv2%mod*cnt) % mod
		}
	}

	ans := f[1<<n-1]
	inv[1] = 1
	for i := 2; i <= n; i++ {
		inv[i] = mod - mod/i*inv[mod%i]%mod
		ans = ans * inv[i] % mod
	}
	Fprint(out, ans)
}

//func main() { cf1842H(bufio.NewReader(os.Stdin), os.Stdout) }
