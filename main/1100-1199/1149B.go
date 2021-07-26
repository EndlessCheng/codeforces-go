package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF1149B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}

	var n, q, id int
	var text, op, ch []byte
	Fscan(in, &n, &q, &text)
	nxt := make([][26]int, n+3)
	for i := 0; i < 26; i++ {
		nxt[n+1][i] = n + 2
		nxt[n+2][i] = n + 2
	}
	for i := n; i > 0; i-- {
		for j := 0; j < 26; j++ {
			nxt[i][j] = nxt[i+1][j]
		}
		nxt[i][text[i-1]-'a'] = i
	}

	s := [3][]byte{{0}, {0}, {0}}
	dp := [251][251][251]int{}
	for ; q > 0; q-- {
		Fscan(in, &op, &id)
		id--
		if op[0] == '+' {
			Fscan(in, &ch)
			s[id] = append(s[id], ch[0]-'a')
			st := [3]int{}
			st[id] = len(s[id]) - 1
			for i := st[0]; i < len(s[0]); i++ {
				for j := st[1]; j < len(s[1]); j++ {
					for k := st[2]; k < len(s[2]); k++ {
						dp[i][j][k] = n + 1
						if i > 0 {
							dp[i][j][k] = min(dp[i][j][k], nxt[dp[i-1][j][k]+1][s[0][i]])
						}
						if j > 0 {
							dp[i][j][k] = min(dp[i][j][k], nxt[dp[i][j-1][k]+1][s[1][j]])
						}
						if k > 0 {
							dp[i][j][k] = min(dp[i][j][k], nxt[dp[i][j][k-1]+1][s[2][k]])
						}
					}
				}
			}
		} else {
			s[id] = s[id][:len(s[id])-1]
		}
		if dp[len(s[0])-1][len(s[1])-1][len(s[2])-1] <= n {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF1149B(os.Stdin, os.Stdout) }
