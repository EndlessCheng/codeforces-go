package main

import (
	. "fmt"
	"io"
)

func cf1996E(in io.Reader, out io.Writer) {
	const mod = 1_000_000_007
	T, S := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &S)
		ans, s := 0, 0
		cnt := map[int]int{0: 1}
		for i, b := range S {
			s += int(b&1)*2 - 1
			ans = (ans + cnt[s]*(len(S)-i)) % mod
			cnt[s] += i + 2
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1996E(bufio.NewReader(os.Stdin), os.Stdout) }
