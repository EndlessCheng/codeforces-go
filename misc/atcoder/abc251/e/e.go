package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n int
	Fscan(in, &n)
	a := make([]int, n)
	for i := range a {
		Fscan(in, &a[i])
	}

	notPay, pay := int(1e18), a[0]
	for _, v := range a[1:] {
		notPay, pay = pay, min(notPay, pay)+v
	}
	ans := notPay

	notPay, pay = a[n-1], a[n-1]+a[0]
	for _, v := range a[1 : n-1] {
		notPay, pay = pay, min(notPay, pay)+v
	}
	Fprintln(out, min(ans, min(notPay, pay)))
}

func main() { run(os.Stdin, os.Stdout) }
func min(a, b int) int { if a > b { return b }; return a }
