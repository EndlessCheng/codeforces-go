package main

import (
	"bufio"
	. "fmt"
	"io"
	"math"
	"os"
	"strconv"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mod = 998244353
	var n, ans, sum int
	var s string
	Fscan(in, &n)
	for i := 0; i < n; i++ {
		Fscan(in, &s)
		v, _ := strconv.Atoi(s)
		ans = (ans + int(math.Pow10(len(s)))%mod*sum + v*i) % mod
		sum = (sum + v) % mod
	}
	Fprint(out, ans)
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
