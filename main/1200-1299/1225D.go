package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1225D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx int = 1e5
	lpf := [mx + 1]int{1: 1}
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}

	var n, k, v int
	ans := int64(0)
	cnt := map[string]int{} // 或者用 [6]int 来代替 string，下面改成 p<<8|e 和 p<<8|(k-e)
	for Fscan(in, &n, &k); n > 0; n-- {
		s := []byte{}
		t := []byte{}
		Fscan(in, &v)
		for v > 1 {
			p, e := lpf[v], 1
			for v /= p; lpf[v] == p; v /= p {
				e++
			}
			if e %= k; e > 0 {
				s = append(s, byte(p>>16), byte(p>>8), byte(p), byte(e))
				t = append(t, byte(p>>16), byte(p>>8), byte(p), byte(k-e))
			}
		}
		ans += int64(cnt[string(t)])
		cnt[string(s)]++
	}
	Fprint(out, ans)
}

//func main() { CF1225D(os.Stdin, os.Stdout) }
