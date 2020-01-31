package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1295C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var t int
	for Fscan(in, &t); t > 0; t-- {
		var s, t []byte
		Fscan(in, &s, &t)
		pos := [26][]int{}
		for i := range s {
			s[i] -= 'a'
			pos[s[i]] = append(pos[s[i]], i)
		}
		for i := range pos {
			pos[i] = append(pos[i], -1)
		}
		next := make([][]int, len(s)+1)
		next[0] = make([]int, 26)
		for i := range next[0] {
			next[0][i] = pos[i][0]
		}
		cur := [26]int{}
		for i, b := range s {
			cur[b]++
			next[i+1] = make([]int, 26)
			copy(next[i+1], next[i])
			next[i+1][b] = pos[b][cur[b]]
		}
		ans := 0
		for i := 0; i < len(t); {
			si := next[0][t[i]-'a']
			if si == -1 {
				ans = -1
				break
			}
			for i++; i < len(t); i++ {
				si = next[si+1][t[i]-'a']
				if si == -1 {
					break
				}
			}
			ans++
		}
		Fprintln(out, ans)
	}
}

//func main() {
//	CF1295C(os.Stdin, os.Stdout)
//}
