package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf318B(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	ans, cnt := 0, 0
	for i := 5; i <= len(s); i++ {
		if s[i-5:i] == "heavy" {
			cnt++
		} else if s[i-5:i] == "metal" {
			ans += cnt
		}
	}
	Fprint(out, ans)
}

//func main() { cf318B(bufio.NewReader(os.Stdin), os.Stdout) }
