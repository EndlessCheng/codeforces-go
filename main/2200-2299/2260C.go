package main

import (
	. "fmt"
	"io"
	"math/bits"
)

// https://github.com/EndlessCheng
func cf2260C(in io.Reader, out io.Writer) {
	var T, x, y int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &x, &y)
		s := x + y
		w := bits.Len(uint(x &^ s))
		mask := 1<<w - 1
		Fprintln(out, s, x&mask-s&mask)
	}
}

//func main() { cf2260C(bufio.NewReader(os.Stdin), os.Stdout) }
