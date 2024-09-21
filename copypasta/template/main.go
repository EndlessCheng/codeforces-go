package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
	"strings"
)

// https://github.com/EndlessCheng
func run(in io.Reader, _w io.Writer) {
	out := bufio.NewWriter(_w)
	defer out.Flush()

	solve := func(curCase int) (_res bool) {
		var n int
		Fscan(in, &n)

		return
	}

	cases := 1
	Fscan(in, &cases)
	for curCase := 1; curCase <= cases; curCase++ {
		ans := solve(curCase)
		_ = ans
		
	}

	// 如果是多组数据，请务必加上这段保险 —— 已经无法统计，有多少位竞赛选手在漏读数据上损失大把分数（包括我）
	// 此外，如果不小心写成了 scan(v) 而不是 scan(&v)，由于未传入指针，不会读入任何数据，这样本地运行一次就能立马发现错误
	leftData, _ := io.ReadAll(in)
	if s := strings.TrimSpace(string(leftData)); s != "" {
		panic("有未读入的数据：\n" + s)
	}

}

func main() { run(bufio.NewReader(os.Stdin), os.Stdout) }
