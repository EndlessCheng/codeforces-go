package main

import (
	"bufio"
	. "fmt"
	"io"
	"sort"
)

func Sol520C(reader io.Reader, writer io.Writer) {
	in := bufio.NewReader(reader)
	out := bufio.NewWriter(writer)
	defer out.Flush()

	var n int
	var s string
	Fscan(in, &n, &s)

	// 找规律，观察发现对于 AAG 这样只有一种字母最多的情况，t 只有一种情况（AAA）；
	// 对于 AAGGT 这样有 m 种字母一样最多的情况，由于滚动循环的特性，每个匹配位置上出现 A 和出现 G 的情况是一样多的（都为两次），
	// 所以 t 的每个字符上是 A 还是 G 对计算结果没有影响，这样答案是 pow(m,n)。
	cnts := make([]int, 26)
	for _, c := range s {
		cnts[c-'A']++
	}
	sort.Ints(cnts)
	m := int64(0)
	for _, c := range cnts {
		if c == cnts[25] {
			m++
		}
	}
	const mod = int64(1e9 + 7)
	ans := int64(1)
	for ; n > 0; n >>= 1 {
		if n&1 == 1 {
			ans = ans * m % mod
		}
		m = m * m % mod
	}
	Fprintln(out, ans)
}

//func main() {
//	Sol520C(os.Stdin, os.Stdout)
//}
