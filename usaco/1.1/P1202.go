package main

import (
	"bufio"
	. "fmt"
	"io"
	"time"
)

// github.com/EndlessCheng/codeforces-go
func SolP1202(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	cnt := [7]int{}
	var n int
	Fscan(in, &n)
	t := time.Date(1900, 1, 13, 0, 0, 0, 0, time.UTC)
	for n *= 12; n > 0; n-- {
		cnt[t.Weekday()]++
		t = t.AddDate(0, 1, 0)
	}
	ans := Sprint(append([]int{cnt[6]}, cnt[:6]...))
	Fprintln(out, ans[1:len(ans)-1])
}

//func main() {
//	SolP1202(os.Stdin, os.Stdout)
//}
