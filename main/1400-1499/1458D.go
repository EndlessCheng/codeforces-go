package main

import (
	"bufio"
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf1458D(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()
	var T int
	var s string
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		cnt := make([]int, n*2+1)
		bias := n
		d := 0
		for _, b := range s {
			if b == '1' {
				cnt[bias+d]++
				d++
			} else {
				d--
				cnt[bias+d]++
			}
		}

		d = 0
		for range n {
			if cnt[bias+d-1] > 0 && (cnt[bias+d] == 0 || cnt[bias+d-1] > 1) {
				d--
				cnt[bias+d]--
				Fprint(out, "0")
			} else {
				cnt[bias+d]--
				d++
				Fprint(out, "1")
			}
		}
		Fprintln(out)
	}
}

//func main() { cf1458D(bufio.NewReader(os.Stdin), os.Stdout) }
