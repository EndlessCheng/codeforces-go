package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF792C(_r io.Reader, out io.Writer) {
	var s, ans string
	Fscan(bufio.NewReader(_r), &s)
	pos := [3][]int{}
	for i, b := range s {
		pos[b%3] = append(pos[b%3], i)
	}
	sum := (len(pos[1]) + len(pos[2])*2) % 3
	if sum == 0 {
		Fprint(out, s)
		return
	}

	if sum == 2 {
		pos[1], pos[2] = pos[2], pos[1]
	}
	rm0 := func(s string) string { // 删去前导零，注意长度为 1 时会停止
		for len(s) > 1 && s[0] == '0' {
			s = s[1:]
		}
		return s
	}
	if ps := pos[1]; len(ps) > 0 {
		p := ps[len(ps)-1]
		ans = rm0(s[:p] + s[p+1:])
	}
	if ps := pos[2]; len(ps) > 1 {
		p, q := ps[len(ps)-2], ps[len(ps)-1]
		if s := rm0(s[:p] + s[p+1:q] + s[q+1:]); len(s) > len(ans) {
			ans = s
		}
	}
	if ans != "" {
		Fprint(out, ans)
	} else {
		Fprint(out, -1)
	}
}

//func main() { CF792C(os.Stdin, os.Stdout) }
