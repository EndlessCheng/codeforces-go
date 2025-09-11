package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1101E(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var q, x, y, maxX, maxY int
	var op string
	Fscan(in, &q)
	for range q {
		Fscan(in, &op, &x, &y)
		if x > y {
			x, y = y, x
		}
		if op == "+" {
			maxX = max(maxX, x)
			maxY = max(maxY, y)
		} else if maxX <= x && maxY <= y {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf1101E(bufio.NewReader(os.Stdin), os.Stdout) }
