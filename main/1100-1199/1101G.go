package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1101G(in io.Reader, out io.Writer) {
	var n, v, xor, num int
	b := [30]int{}
	Fscan(in, &n)
	for range n {
		Fscan(in, &v)
		xor ^= v
		//v = xor // 注意前缀和也可以可以由基表出的，所以可以不写
		for i := len(b) - 1; i >= 0; i-- {
			if v>>i == 0 {
				continue
			}
			if b[i] == 0 {
				b[i] = v
				num++
				break
			}
			v ^= b[i]
		}
	}
	if xor == 0 {
		num = -1
	}
	Fprint(out, num)
}

//func main() { cf1101G(bufio.NewReader(os.Stdin), os.Stdout) }
