package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

// github.com/EndlessCheng/codeforces-go
func CF748D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var k, n, v, mx, ans int
	var s string
	np := map[string][]int{}
	p := map[string][]int{}
o:
	for Fscan(in, &k, &n); k > 0; k-- {
		Fscan(in, &s, &v)
		for i := 0; i < n/2; i++ {
			if s[i] != s[n-1-i] {
				np[s] = append(np[s], -v)
				continue o
			}
		}
		p[s] = append(p[s], v)
	}
	for s, a := range np {
		t := []byte(s)
		for i := 0; i < n/2; i++ {
			t[i], t[n-1-i] = t[n-1-i], t[i]
		}
		rev := string(t)
		b := np[rev]
		if b == nil {
			continue
		}
		sort.Ints(a)
		sort.Ints(b)
		for i := 0; i < len(a) && i < len(b) && a[i]+b[i] < 0; i++ {
			ans -= a[i] + b[i]
		}
		delete(np, rev)
	}
	for _, a := range p {
		sort.Ints(a)
		i := len(a) - 1
		for ; i > 0 && a[i]+a[i-1] > 0; i -= 2 {
			ans += a[i] + a[i-1]
			if -a[i-1] > mx {
				mx = -a[i-1]
			}
		}
		if i >= 0 && a[i] > mx {
			mx = a[i]
		}
	}
	Fprint(out, ans+mx)
}

//func main() { CF748D(os.Stdin, os.Stdout) }
