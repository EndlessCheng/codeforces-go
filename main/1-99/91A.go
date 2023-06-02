package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func CF91A(in io.Reader, out io.Writer) {
	var s, t string
	Fscan(bufio.NewReader(in), &s, &t)
	n := len(s)
	pos := [26]int{}
	for i := range pos {
		pos[i] = n
	}
	nxt := make([][26]int, n)
	for k := 2; k > 0; k-- {
		for i := n - 1; i >= 0; i-- {
			nxt[i] = pos
			pos[s[i]-'a'] = i
		}
	}

	ans := 1
	i, j := 0, 0
	if t[0] == s[0] {
		j = 1
	}
	for ; j < len(t); j++ {
		i2 := nxt[i][t[j]-'a']
		if i2 == n {
			Fprint(out, -1)
			return
		}
		if i2 <= i {
			ans++
		}
		i = i2
	}
	Fprint(out, ans)
}

//func main() { CF91A(os.Stdin, os.Stdout) }
