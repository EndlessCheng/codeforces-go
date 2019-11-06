package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol546D(reader io.Reader, writer io.Writer) {
	in := bufio.NewScanner(reader)
	in.Split(bufio.ScanWords)
	out := bufio.NewWriter(writer)
	defer out.Flush()
	read := func() (x int) {
		in.Scan()
		for _, b := range in.Bytes() {
			x = x*10 + int(b-'0')
		}
		return
	}

	n := 5000000
	cnt := make([]int, n+1)
	primes := make([]int, 0, 348513)
	for i := 2; i <= n; i++ {
		if cnt[i] == 0 {
			primes = append(primes, i)
			cnt[i] = 1
		}
		for _, p := range primes {
			if j := i * p; j <= n {
				cnt[j] = cnt[i] + 1
			} else {
				break
			}
		}
	}
	for i := 3; i <= n; i++ {
		cnt[i] += cnt[i-1]
	}
	for n := read(); n > 0; n-- {
		Fprintln(out, cnt[read()]-cnt[read()])
	}
}

//func main() {
//	Sol546D(os.Stdin, os.Stdout)
//}
