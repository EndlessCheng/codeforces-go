package main

import (
	. "fmt"
	"io"
)

// https://space.bilibili.com/206214
func cf1066B(in io.Reader, out io.Writer) {
	var n, r, v, ans, done int
	Fscan(in, &n, &r)
	nxt := -1
	for i := 0; i < n; i++ {
		if i >= done+r {
			if nxt < 0 {
				Fprint(out, -1)
				return
			}
			ans++
			done = nxt + r
			nxt = -1
		}
		Fscan(in, &v)
		if v > 0 {
			nxt = i
		}
	}
	if done < n {
		if nxt < 0 || nxt+r < n {
			Fprint(out, -1)
			return
		}
		ans++
	}
	Fprint(out, ans)
}

//func main() { cf1066B(bufio.NewReader(os.Stdin), os.Stdout) }
