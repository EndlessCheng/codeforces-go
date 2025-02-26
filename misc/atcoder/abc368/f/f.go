package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://github.com/EndlessCheng
func run(in io.Reader, out io.Writer) {
	const mx int = 1e5
	Omega := [mx + 1]int8{}
	primes := []int{}
	for i := 2; i <= mx; i++ {
		if Omega[i] == 0 {
			Omega[i] = 1
			primes = append(primes, i)
		}
		for _, p := range primes {
			if p*i > mx {
				break
			}
			Omega[p*i] = Omega[i] + 1
		}
	}
	var n, v, xor int
	for Fscan(in, &n); n > 0; n-- {
		Fscan(in, &v)
		xor ^= int(Omega[v])
	}
	if xor > 0 {
		Fprint(out, "Anna")
	} else {
		Fprint(out, "Bruno")
	}
}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
