package main

import (
	"bufio"
	. "bytes"
	. "fmt"
	"io"
	"os"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, k int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &k, &s)
		has := [26]bool{}
		for _, b := range s {
			has[b-'a'] = true
		}
		set := []byte{} // 字母集
		for i, b := range has {
			if b {
				set = append(set, 'a'+byte(i))
			}
		}

		if k > n { // t 比 s 长，故直接在 s 后面添加字母集中最小的字母
			Fprintf(out, "%s", s)
			Fprintf(out, "%s\n", Repeat([]byte{set[0]}, k-n))
			continue
		}

		// 寻找可以变更字母的位置，即最靠右的小于字母集最大字母的位置
		for i := k - 1; ; i-- {
			if s[i] < set[len(set)-1] {
				j := IndexByte(set, s[i]) + 1
				s[i] = set[j] // 将 s[i] 变更为字母集的下一个字母
				Fprintf(out, "%s", s[:i+1])
				Fprintf(out, "%s\n", Repeat([]byte{set[0]}, k-i-1)) // s[i] 右边的用字母集中最小的字母
				break
			}
		}
	}
}

func main() { run(os.Stdin, os.Stdout) }
