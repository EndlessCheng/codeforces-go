package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2050C(in io.Reader, out io.Writer) {
	T, s := 0, ""
o:
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		cnt := [10]int{}
		sum := 0
		for _, b := range s {
			cnt[b-'0']++
			sum += int(b - '0')
		}
		for i := range min(cnt[2], 9) + 1 {
			for j := range min(cnt[3], 3) + 1 {
				if (sum+i*2+j*6)%9 == 0 {
					Fprintln(out, "YES")
					continue o
				}
			}
		}
		Fprintln(out, "NO")
	}
}

//func main() { cf2050C(bufio.NewReader(os.Stdin), os.Stdout) }
