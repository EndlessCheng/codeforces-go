package main

import (
	. "fmt"
	"io"
	"runtime/debug"
)

// https://github.com/EndlessCheng
func init() { debug.SetGCPercent(-1) }

type node57 struct {
	son [2]*node57
	cnt int
	sum int
}

type trie57 struct{ root *node57 }

func (t *trie57) put(s []byte, isPal []bool) {
	tot := 0
	for _, b := range isPal {
		if b {
			tot++
		}
	}
	o := t.root
	for i, b := range s {
		b -= 'a'
		if o.son[b] == nil {
			o.son[b] = &node57{}
		}
		o = o.son[b]
		o.sum += tot
		if isPal[i] {
			o.cnt++
			tot--
		}
	}
}

func (t *trie57) kth(k int) (s []byte) {
	o := t.root
	for {
		for i, son := range o.son {
			if son == nil {
				continue
			}
			if k > son.sum {
				k -= son.sum
				continue
			}
			s = append(s, 'a'+byte(i))
			o = son
			k -= o.cnt
			if k <= 0 {
				return
			}
			break
		}
	}
}

func cf557E(in io.Reader, out io.Writer) {
	var s []byte
	var k int
	Fscan(in, &s, &k)
	n := len(s)
	isPal := make([][]bool, n)
	for i := range isPal {
		isPal[i] = make([]bool, n)
	}
	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			isPal[l][r] = true
			l -= 2
			r += 2
		}
		l, r = i/2-1, (i+1)/2+1
		for l >= 0 && r < n && s[l] == s[r] {
			isPal[l][r] = true
			l -= 2
			r += 2
		}
	}

	t := &trie57{&node57{}}
	for i, row := range isPal {
		t.put(s[i:], row[i:])
	}
	Fprintf(out, "%s", t.kth(k))
}

//func main() { cf557E(os.Stdin, os.Stdout) }
