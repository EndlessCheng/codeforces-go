package main

import (
	. "fmt"
	"io"
)

func cf628B(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	var ans, x int
	for i, c := range s {
		if c%4 == 0 {
			ans++
		}
		x = x*10 + int(c-'0')
		if x%4 == 0 {
			ans += i
		}
		if i > 0 {
			x -= int(s[i-1]-'0') * 10
		}
	}
	Fprint(out, ans)
}

//func main() { cf628B(bufio.NewReader(os.Stdin), os.Stdout) }
