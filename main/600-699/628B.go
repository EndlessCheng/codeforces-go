package main

import (
	. "fmt"
	"io"
)

func cf628B(in io.Reader, out io.Writer) {
	var s []byte
	Fscan(in, &s)
	ans := 0
	for i, c := range s {
		if c%4 == 0 {
			ans++
		}
		if i > 0 && (s[i-1]*2+c)%4 == 0 {
			ans += i
		}
	}
	Fprint(out, ans)
}

//func main() { cf628B(bufio.NewReader(os.Stdin), os.Stdout) }
