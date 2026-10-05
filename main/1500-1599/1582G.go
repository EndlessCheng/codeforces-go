package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1582G(in io.Reader, out io.Writer) {
	var n int
	var s string
	Fscan(in, &n)
	a := make([]int, n+1)
	mx := 0
	for i := 1; i <= n; i++ {
		Fscan(in, &a[i])
		mx = max(mx, a[i])
	}
	Fscan(in, &s)
	s = " " + s

	lpf := make([]int, mx+1)
	primes := make([]int, mx+1)
	pn := 0

	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			pn++
			primes[pn] = i
			lpf[i] = i
		}
		for j := 1; j <= pn && i*primes[j] <= mx; j++ {
			lpf[i*primes[j]] = primes[j]
			if i%primes[j] == 0 {
				break
			}
		}
	}

	pr := make([]int, n+1)
	st := make([]int, n+1)
	w := make([]int, n+1)
	ve := make([][]int, mx+1)

	ans := 0
	for i := 1; i <= n; i++ {
		pr[i] = i
		for j := a[i]; j > 1; j /= lpf[j] {
			p := lpf[j]
			if s[i] == '*' {
				ve[p] = append(ve[p], i)
			} else if len(ve[p]) == 0 {
				pr[i] = 0
			} else {
				pr[i] = min(pr[i], ve[p][len(ve[p])-1])
				ve[p] = ve[p][:len(ve[p])-1]
			}
		}
	}

	t := 0
	for i := n; i >= 1; i-- {
		v := 1
		for t > 0 && pr[i] <= pr[st[t]] {
			v += w[t]
			t--
		}
		t++
		st[t] = i
		w[t] = v
		if pr[i] == i {
			ans += v
		}
	}
	Fprint(out, ans)
}

//func main() { cf1582G(bufio.NewReader(os.Stdin), os.Stdout) }
