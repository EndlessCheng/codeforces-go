package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF519D(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	a := [26]int64{}
	for i := range a {
		Fscan(in, &a[i])
	}
	var s []byte
	Fscan(in, &s)
	var sum, ans int64
	cnt := map[int64][26]int{}
	for i, v := range s[:len(s)-1] {
		v -= 'a'
		w := s[i+1] - 'a'
		if v == w {
			ans++
		}
		sum += a[v]
		c := cnt[sum]
		ans += int64(c[w])
		c[v]++
		cnt[sum] = c
	}
	Fprint(out, ans)
}

//func main() { CF519D(os.Stdin, os.Stdout) }
