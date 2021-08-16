package main

import (
	"bufio"
	"bytes"
	. "fmt"
	"io"
	"sort"
	"strings"
)

/* 思考记录
- 全一样，直接输出
- 如果有只出现一次的字母，把只出现一次且最小的放到首位，其余排序放后面（这样可以构造答案为 0 的）
- 此时答案必然为 1，尝试拼凑字典序最小的：（假设 a 为最小，b 为次小，c 为第三小）
  - 如果最小字母的个数-2 不超过其余字母个数：aa+其余字母排序，其余 a 间隔插入排序中的字母
  - 如果只有两种字母：abb..bbaa..aa
  - 如果有三种或以上的字母：abaa..aac+其余字母排序
*/

// github.com/EndlessCheng/codeforces-go
func CF1530E(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T int
	var s []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s)
		n := len(s)
		cnt := [26]int{}
		for _, b := range s {
			cnt[b-'a']++
		}
		set := []int{}
		pc1 := -1
		for i, c := range cnt {
			if c > 0 {
				set = append(set, i)
				if c == 1 && pc1 < 0 {
					pc1 = i
				}
			}
		}
		if len(set) == 1 {
			Fprintf(out, "%s\n", s)
			continue
		}
		if pc1 >= 0 {
			sort.Slice(s, func(i, j int) bool { return s[i] < s[j] })
			b := 'a' + byte(pc1)
			i := bytes.IndexByte(s, b)
			Fprintf(out, "%c%s%s\n", b, s[:i], s[i+1:])
			continue
		}
		cnt0 := cnt[set[0]]
		b0 := 'a' + byte(set[0])
		if cnt0-2 <= n-cnt0 {
			Fprintf(out, "%c%c", b0, b0)
			cnt0 -= 2
			for _, i := range set[1:] {
				for c := cnt[i]; c > 0; c-- {
					Fprintf(out, "%c", 'a'+byte(i))
					if cnt0 > 0 {
						Fprintf(out, "%c", b0)
						cnt0--
					}
				}
			}
			Fprintln(out)
		} else if len(set) == 2 {
			Fprintf(out, "%c%s%s\n", b0, strings.Repeat(string('a'+byte(set[1])), cnt[set[1]]), strings.Repeat(string('a'+byte(set[0])), cnt[set[0]]-1))
		} else {
			Fprintf(out, "%c%c%s%c", b0, 'a'+byte(set[1]), strings.Repeat(string('a'+byte(set[0])), cnt[set[0]]-1), 'a'+byte(set[2]))
			cnt[set[1]]--
			cnt[set[2]]--
			for _, i := range set[1:] {
				Fprintf(out, "%s", strings.Repeat(string('a'+byte(i)), cnt[i]))
			}
			Fprintln(out)
		}
	}
}

//func main() { CF1530E(os.Stdin, os.Stdout) }
