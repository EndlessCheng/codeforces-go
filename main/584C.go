package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF584C(_r io.Reader, out io.Writer) {
	var n, needSame int
	var s, t []byte
	Fscan(bufio.NewReader(_r), &n, &needSame, &s, &t)
	needSame = n - needSame
	same := 0
	for i, b := range s {
		if b == t[i] {
			same++
		}
	}
	ans := make([]byte, n)
	if needSame <= same {
		for i, b := range s {
			c := t[i]
			if needSame > 0 && b == c {
				ans[i] = b
				needSame--
				continue
			}
			ans[i] = 'a'
			if ans[i] == b || ans[i] == c {
				ans[i]++
			}
			if ans[i] == b || ans[i] == c {
				ans[i]++
			}
		}
	} else if (needSame-same)*2 > n-same {
		Fprint(out, -1)
		return
	} else {
		needSame = (needSame - same) * 2
		cnt := 0
		for i, b := range s {
			c := t[i]
			if b == c {
				ans[i] = b
				continue
			}
			if cnt*2 < needSame {
				cnt++
				ans[i] = b
			} else if cnt < needSame {
				cnt++
				ans[i] = c
			} else {
				ans[i] = 'a'
				if ans[i] == b || ans[i] == c {
					ans[i]++
				}
				if ans[i] == b || ans[i] == c {
					ans[i]++
				}
			}
		}
	}
	Fprint(out, string(ans))
}

//func main() { CF584C(os.Stdin, os.Stdout) }
