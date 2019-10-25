package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol1249C2(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var q int
	for Fscan(in, &q); q > 0; q-- {
		var n int64
		Fscan(in, &n)
		coef := []int64{}
		for ; n > 0; n /= 3 {
			coef = append(coef, n%3)
		}
		coef = append(coef, 0)
		for i := len(coef) - 1; i >= 0; i-- {
			if coef[i] == 2 {
				for j := range coef[:i] {
					coef[j] = 0
				}
				for j := i; j < len(coef); j++ {
					if coef[j] == 2 {
						coef[j] = 0
						coef[j+1]++
					}
				}
			}
		}
		ans := int64(0)
		pow3 := int64(1)
		for _, ci := range coef {
			ans += pow3 * ci
			pow3 *= 3
		}
		Fprintln(out, ans)
	}
}

//func main() {
//	Sol1249C2(os.Stdin, os.Stdout)
//}
