package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"sort"
)

// https://space.bilibili.com/206214
func cf1781C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n int
	var s []byte
	cnt := [26]int{}
	id := make([]int, 26)
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n, &s)
		for _, b := range s {
			cnt[b-'a']++
		}
		for i := range id {
			id[i] = i
		}
		sort.Slice(id, func(i, j int) bool { return cnt[id[i]] > cnt[id[j]] })

		maxSave := 0
		k := 0
		for i := 1; i <= 26; i++ {
			if n%i > 0 {
				continue
			}
			save := 0
			for _, j := range id[:i] {
				save += min(cnt[j], n/i)
			}
			if save > maxSave {
				maxSave = save
				k = i // 变成出现次数最多的 k 种字母
			}
		}

		need := []byte{}
		m := n / k // 每种字母最终有 m 个
		for _, i := range id[:k] {
			if cnt[i] > m {
				cnt[i] = m
			} else {
				// 还需要 m-cnt[i] 个字母
				need = append(need, bytes.Repeat([]byte{'a' + byte(i)}, m-cnt[i])...)
			}
		}
		for _, i := range id[k:] {
			cnt[i] = 0
		}

		j := 0
		for i, b := range s {
			b -= 'a'
			if cnt[b] > 0 {
				cnt[b]-- // 保持不变
			} else {
				s[i] = need[j] // 变
				j++
			}
		}
		Fprintln(out, n-maxSave)
		Fprintf(out, "%s\n", s)
	}
}

//func main() { cf1781C(os.Stdin, os.Stdout) }
