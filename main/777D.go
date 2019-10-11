package main

import (
	"bufio"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func Sol777D(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	Fscan(in, &n)
	tags := make([]string, n)
	for i := range tags {
		Fscan(in, &tags[i])
	}

	for i := len(tags) - 2; i >= 0; i-- {
		t0, t1 := tags[i], tags[i+1]
		l := len(t0)
		if len(t1) < l {
			l = len(t1)
		}
		if t0[:l] == t1[:l] {
			tags[i] = t0[:l]
			continue
		}
		j := 1
		for ; j < l; j++ {
			if t0[j] < t1[j] {
				break
			}
			if t0[j] > t1[j] {
				tags[i] = tags[i][:j]
				break
			}
		}
	}
	for _, tag := range tags {
		Fprintln(out, tag)
	}
}

//func main() {
//	Sol777D(os.Stdin, os.Stdout)
//}
