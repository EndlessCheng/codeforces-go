package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf701C(in io.Reader, out io.Writer) {
	var ans, all, l int
	var s []byte
	Fscan(in, &ans, &s)
	for _, b := range s {
		all |= 1 << (b - 'A')
	}
	less := bits.OnesCount(uint(all))
	cnt := [256]int{}
	for r, b := range s {
		if cnt[b] == 0 {
			less--
		}
		cnt[b]++
		for less == 0 {
			ans = min(ans, r-l+1)
			cnt[s[l]]--
			if cnt[s[l]] == 0 {
				less++
			}
			l++
		}
	}
	Fprint(out, ans)
}

//func main() { cf701C(bufio.NewReader(os.Stdin), os.Stdout) }
