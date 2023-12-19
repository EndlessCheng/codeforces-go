package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1913D(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	const mod = 998244353

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		s := make([]int, n+1)
		type data struct{ i, v, f int }
		st := []data{{-1, 0, 0}}
		stSum := 0
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			for st[len(st)-1].v > v {
				stSum -= st[len(st)-1].f
				st = st[:len(st)-1]
			}
			f := (stSum + s[i] - s[st[len(st)-1].i+1]) % mod
			if len(st) == 1 {
				f++ // 把前面全删了
			}
			stSum = (stSum + f) % mod
			s[i+1] = s[i] + f
			st = append(st, data{i, v, f})
		}
		Fprintln(out, (stSum+mod)%mod)
	}
}

//func main() { cf1913D(os.Stdin, os.Stdout) }
