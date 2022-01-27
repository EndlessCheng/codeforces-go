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

	solve := func(curCase int) {
		var n int
		Fscan(in, &n)

	}

	cases := 1
	Fscan(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		solve(curCase)
	}

	// 如果是多组数据，请务必加上这段保险 —— 已经无法统计，有多少位竞赛选手在漏读数据上损失大把分数（包括我）
	// 此外，如果不小心写成了 scan(v) 而不是 scan(&v)，由于未传入指针，不会读入任何数据，这样本地运行一次就能立马发现错误
	_leftData, _ := ioutil.ReadAll(in)
	if _s := strings.TrimSpace(string(_leftData)); _s != "" {
		panic("有未读入的数据：\n" + _s)
	}

}

func main() { run(os.Stdin, os.Stdout) }
