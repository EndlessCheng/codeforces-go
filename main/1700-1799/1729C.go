package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF1729C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		pos := [26][]int{}
		for i, b := range s {
			b -= 'a'
			pos[b] = append(pos[b], i+1)
		}
		d := 1
		if s[0] > s[n-1] {
			d = -1
		}
		ans := []int{}
		for i := int(s[0] - 'a'); ; i += d {
			ans = append(ans, pos[i]...)
			if i == int(s[n-1]-'a') {
				break
			}
		}
		Fprintln(out, abs(int(s[0])-int(s[n-1])), len(ans))
		for _, v := range ans {
			Fprint(out, v, " ")
		}
		Fprintln(out)
	}
}

//func main() { CF1729C(os.Stdin, os.Stdout) }
