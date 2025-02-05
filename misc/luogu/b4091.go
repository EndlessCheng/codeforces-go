package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func b4091(in io.Reader, out io.Writer) {
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	cnt := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
		cnt[i] = 1
	}
	for {
		ok := false
		for i, v := range a {
			j := (i + 1) % n
			if v < a[j] && cnt[i] >= cnt[j] {
				cnt[j] = cnt[i] + 1
				ok = true
			} else if v > a[j] && cnt[j] >= cnt[i] {
				cnt[i] = cnt[j] + 1
				ok = true
			}
		}
		if !ok {
			break
		}
	}
	ans := 0
	for _, c := range cnt {
		ans += c
	}
	Fprint(out, ans)
}

//func main() { b4091(bufio.NewReader(os.Stdin), os.Stdout) }
