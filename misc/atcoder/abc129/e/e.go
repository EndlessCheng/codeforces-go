package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
const mod = 1_000_000_007

func run(in io.Reader, out io.Writer) {
	const inv3 = (mod + 1) / 3
	var s string
	Fscan(in, &s)
	ans := 0
	pow2 := 1
	pow3 := pow(3, len(s))
	for _, b := range s {
		pow3 = pow3 * inv3 % mod
		if b == '1' {
			ans = (ans + pow2*pow3) % mod
			pow2 = pow2 * 2 % mod
		}
	}
	Fprint(out, (ans+pow2)%mod)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }

func pow(x, n int) int {
	res := 1
	for ; n > 0; n /= 2 {
		if n%2 > 0 {
			res = res * x % mod
		}
		x = x * x % mod
	}
	return res
}
