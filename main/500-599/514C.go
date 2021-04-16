package main

import (
	"bufio"
	. "fmt"
	"io"
)

// 复杂度 O(|S|√|S|)，|S| 为字符串长度之和

// github.com/EndlessCheng/codeforces-go
type node14 struct {
	son [3]*node14
	end bool
}

func (o *node14) find(s string, miss bool) bool {
	if s == "" {
		return miss && o.end
	}
	for i, c := range o.son {
		if c != nil && (byte(i) == s[0]-'a' && c.find(s[1:], miss) || !miss && byte(i) != s[0]-'a' && c.find(s[1:], true)) {
			return true
		}
	}
	return false
}

type trie14 struct{ root *node14 }

func (t *trie14) put(s string) *node14 {
	o := t.root
	for i := range s {
		b := s[i] - 'a'
		if o.son[b] == nil {
			o.son[b] = &node14{}
		}
		o = o.son[b]
	}
	o.end = true
	return o
}

func CF514C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	t := &trie14{&node14{}}
	var n, m int
	var s string
	for Fscan(in, &n, &m); n > 0; n-- {
		Fscan(in, &s)
		t.put(s)
	}
	for ; m > 0; m-- {
		if Fscan(in, &s); t.root.find(s, false) {
			Fprintln(out, "YES")
		} else {
			Fprintln(out, "NO")
		}
	}
}

//func main() { CF514C(os.Stdin, os.Stdout) }
