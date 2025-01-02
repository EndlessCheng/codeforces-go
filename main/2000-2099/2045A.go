package main

import (
	. "fmt"
	"io"
)

// https://github.com/EndlessCheng
func cf2045A(in io.Reader, out io.Writer) {
	var s string
	Fscan(in, &s)
	n := len(s)
	cnt := ['Z' + 1]int{}
	for _, b := range s {
		cnt[b]++
	}

	a := cnt['A'] + cnt['E'] + cnt['I'] + cnt['O'] + cnt['U']
	y := cnt['Y']
	b := n - a - y
	ng := min(cnt['N'], cnt['G'])

	if (a+y)*2 <= ng { // 元音少，甚至比 NG 还少
		Fprint(out, (a+y)*5) // 只用 NG 辅音
	} else if (a+y)*2 <= b-ng { // 元音少，即使把 NG 合并（减少辅音个数）仍然少
		Fprint(out, (a+y)*3+ng) // 所有 NG 全部用上
	} else if (b+y)/2 <= a { // 辅音少
		res := (b + y) / 2 * 3
		if (b+y)%2 > 0 && ng > 0 {
			res++ // 多出的一个辅音可以是 N 或者 G，合并到 NG 中
		}
		Fprint(out, res)
	} else {
		// 如果没有 NG，那么答案一定是 3 的倍数
		// 如果只有一个 NG，那么当 n 是 3k+2 时，一定会多出一个字母，例如 AYNGG
		// 其余情况可以用 NG 和 Y 灵活调整，答案是 n
		Fprint(out, n-max(n%3-ng, 0))
	}
}

//func main() { cf2045A(bufio.NewReader(os.Stdin), os.Stdout) }
