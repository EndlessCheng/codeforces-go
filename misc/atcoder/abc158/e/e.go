package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	var n, p, ans int
	var s string
	Fscan(in, &n, &p, &s)
	if p == 2 || p == 5 {
		for i, b := range s {
			if int(b-'0')%p == 0 {
				ans += i + 1
			}
		}
	} else {
		cnt := make([]int, p) // map[int]int{}
		cnt[0] = 1
		v, pow10 := 0, 1
		for i := n - 1; i >= 0; i-- {
			v = (v + int(s[i]-'0')*pow10) % p
			pow10 = pow10 * 10 % p
			ans += cnt[v]
			cnt[v]++
		}
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
