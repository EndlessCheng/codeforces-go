package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1076F(in io.Reader, out io.Writer) {
	var n, k, y, leftX, leftY int
	Fscan(in, &n, &k)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}
	for _, x := range a {
		Fscan(in, &y)
		leftX = max(leftX+x-y*k, 0)
		leftY = max(leftY+y-x*k, 0)
		if leftX > k || leftY > k {
			Fprint(out, "NO")
			return
		}
	}
	Fprint(out, "YES")
}

//func main() { cf1076F(bufio.NewReader(os.Stdin), os.Stdout) }
