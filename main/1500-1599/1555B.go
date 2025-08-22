package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1555B(in io.Reader, out io.Writer) {
	var T, W, H, x1, y1, x2, y2, w, h int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &W, &H, &x1, &y1, &x2, &y2, &w, &h)
		ans := int(1e9)
		if w+x2-x1 <= W {
			ans = max(w-max(x1, W-x2), 0)
		}
		if h+y2-y1 <= H {
			ans = min(ans, max(h-max(y1, H-y2), 0))
		}
		if ans == 1e9 {
			ans = -1
		}
		Fprintln(out, ans)
	}
}

//func main() { cf1555B(bufio.NewReader(os.Stdin), os.Stdout) }
