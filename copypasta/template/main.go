package main

import (
	"bufio"
	. "fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// github.com/EndlessCheng/codeforces-go
func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	// 如果是多组数据，请务必加上这段保险 —— 已经无法统计，有多少位竞赛选手在漏读数据上损失大把分数（包括我）
	defer func() {
		leftData, _ := ioutil.ReadAll(in)
		s := strings.TrimSpace(string(leftData))
		if s != "" {
			panic("有未读入的数据：\n" + s)
		}
	}()

	var n int
	Fscan(in, &n)

}

func main() { run(os.Stdin, os.Stdout) }
