package __99

import (
	"bufio"
	. "fmt"
	"io"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func CF95A(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	mp := map[string]bool{}
	var n int
	var s string
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &s)
		mp[strings.ToLower(s)] = true
	}
	var l []byte
	Fscan(in, &s, &l)
	t := strings.ToLower(s)
	ans := []byte(s)
	c := l[0]

	for i := range t {
		for j := i + 1; j <= len(t); j++ {
			if mp[t[i:j]] {
				for k := i; k < j; k++ {
					tar := byte('a')
					if t[k] != c {
						tar = c
					} else if t[k] == 'a' {
						tar = 'b'
					}
					ans[k] = tar
					if s[k] < 'a' {
						ans[k] -= 32
					}
				}
			}
		}
	}
	Fprint(out, string(ans))
}

//func main() { CF95A(os.Stdin, os.Stdout) }
