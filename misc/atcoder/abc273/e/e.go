package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

// https://space.bilibili.com/206214
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	type node struct {
		fa  *node
		val int
	}
	root := &node{val: -1}
	root.fa = root // 小技巧
	cur := root
	nodes := map[int]*node{}

	var q, x int
	var op string
	for Fscan(in, &q); q > 0; q-- {
		Fscan(in, &op, &x)
		switch op[0] {
		case 'A':
			cur = &node{cur, x}
		case 'D':
			cur = cur.fa
		case 'S':
			nodes[x] = cur
		default:
			if o, ok := nodes[x]; ok {
				cur = o
			} else {
				cur = root
			}
		}
		Fprint(out, cur.val, " ")
	}
}

func main() { run(os.Stdin, os.Stdout) }
