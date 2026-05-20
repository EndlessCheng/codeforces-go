package main

import (
	. "fmt"
	"io"
	"strings"
)

// https://github.com/EndlessCheng
func cf2030C(in io.Reader, out io.Writer) {
	T, s := 0, ""
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &s)
		if strings.Contains("1"+s+"1", "11") {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { cf2030C(bufio.NewReader(os.Stdin), os.Stdout) }
