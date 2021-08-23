package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
)

// github.com/EndlessCheng/codeforces-go
func CF940C(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	var n, k int
	var s []byte
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

	if k > n { // t 比 s 长，故直接在 s 后面添加 k-n 个字母集中最小的字母
		Fprintf(out, "%s", s)
		Fprintf(out, "%s\n", bytes.Repeat([]byte{set[0]}, k-n))
		return
	}

	// 寻找可以变更字母的位置，即最靠右的小于字母集最大字母的位置
	for i := k - 1; ; i-- { // 题目保证答案存在
		if s[i] < set[len(set)-1] {
			j := bytes.IndexByte(set, s[i])
			s[i] = set[j+1] // 将 s[i] 变更为字母集中比 s[i] 大的下一个字母
			Fprintf(out, "%s", s[:i+1])
			Fprintf(out, "%s\n", bytes.Repeat([]byte{set[0]}, k-i-1)) // s[i] 右边的用字母集中最小的字母
			return
		}
	}
}

//func main() { CF940C(os.Stdin, os.Stdout) }
