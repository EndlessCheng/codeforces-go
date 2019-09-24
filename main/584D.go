package main

import (
	"bufio"
	. "fmt"
	"io"
)

func Sol584D(reader io.Reader, writer io.Writer) {
	isPrime := func(n int) bool {
		for i := 2; i*i <= n; i++ {
			if n%i == 0 {
				return false
			}
		}
		return true
	}

	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	if isPrime(n) {
		Fprintln(out, 1)
		Fprintln(out, n)
		return
	}
	n -= 3
	for i := 3; ; i += 2 {
		if isPrime(i) && isPrime(n-i) {
			Fprintln(out, 3)
			Fprintln(out, 3, i, n-i)
			return
		}
	}
}

//func main() {
//	Sol584D(os.Stdin, os.Stdout)
//}
