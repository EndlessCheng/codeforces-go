package copypasta

import (
	"bytes"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/* 其他无法分类的算法

Smallest number h such that n*h is a repunit (111...1), or 0 if no such h exists
https://oeis.org/A190301 111...1
https://oeis.org/A216485 222...2

Least k such that the decimal representation of k*n contains only 1's and 0's
https://oeis.org/A079339
0's and d's (2~9): A096681-A096688

a(n) is the least value of k such that k*n uses only digits 1 and 2. a(n) = -1 if no such multiple exists
https://oeis.org/A216482

a(n) is the smallest positive number such that the decimal digits of n*a(n) are all 0, 1 or 2
https://oeis.org/A181061

三维 n 皇后 http://oeis.org/A068940
Maximal number of chess queens that can be placed on a 3-dimensional chessboard of order n so that no two queens attack each other

Smallest positive integer k such that n = +-1+-2+-...+-k for some choice of +'s and -'s https://oeis.org/A140358
相关题目 https://codeforces.com/problemset/problem/1278/B

Numbers n such that n is the substring identical to the least significant bits of its base 2 representation.
http://oeis.org/A181891
http://oeis.org/A181929 前缀
http://oeis.org/A038102 子串

Maximal number of regions obtained by joining n points around a circle by straight lines.
Also number of regions in 4-space formed by n-1 hyperplanes.
a(n) = n*(n-1)*(n*n-5*n+18)/24+1 https://oeis.org/A000127
*/
func miscCollection() {
	// 预处理 log 的整数部分
	logInit := func() {
		const mx int = 1e6
		log := [mx + 1]int{}
		for i := 2; i <= mx; i++ {
			log[i] = log[i>>1] + 1
		}
	}

	// 找环
	// 1<=next[i]<=n
	getCycle := func(next []int, n, st int) (beforeCycle, cycle []int) {
		vis := make([]int8, n+1)
		for v := st; vis[v] < 2; v = next[v] {
			if vis[v] == 1 {
				cycle = append(cycle, v)
			}
			vis[v]++
		}
		for v := st; vis[v] == 1; v = next[v] {
			beforeCycle = append(beforeCycle, v)
		}
		return
	}

	ceilK := func(n, k int) int {
		if n%k == 0 {
			return n
		}
		return n + k - n%k
	}

	// 把 n 用 m 等分，得到 m-n%m 个 n/m 和 n%m 个 n/m+1
	partition := func(n, m int) (q, cntQ, cntQ1 int) {
		// m must > 0
		return n / m, m - n%m, n % m
	}

	// 从 st 出发，步长为 gap，不超过 upper 的最大值
	// st <= upper, gap > 0
	maxValueStepToUpper := func(st, upper, gap int) int {
		upper -= st
		return st + upper - upper%gap
	}

	// 从 st 跳到 [l,r]，每次跳 d 个单位长度，问首次到达的位置（或无法到达）
	moveToRange := func(st, d, l, r int) (firstPos int, ok bool) {
		switch {
		case st < l:
			if d <= 0 {
				return
			}
			return l + ((st-l)%d+d)%d, true
		case st <= r:
			return st, true
		default:
			if d >= 0 {
				return
			}
			return r + ((st-r)%d+d)%d, true
		}
	}

	hash01Mat := func(mat [][]int) int {
		hash := 0
		cnt := 0
		for _, row := range mat {
			for _, v := range row {
				hash |= v << cnt
				cnt++
			}
		}
		return hash
	}

	// 用堆求前 k 小
	smallK := func(a []int, k int) []int {
		k++
		q := hp{} // 最大堆
		for _, v := range a {
			if q.Len() < k || v < q.top() {
				q.push(v)
			}
			if q.Len() > k {
				q.pop() // 不断弹出更大的元素，留下的就是较小的
			}
		}
		return q.IntSlice // 注意返回的不是有序数组
	}

	removeLeadingZero := func(s string) string {
		for i, b := range s {
			if b > '0' {
				return s[i:]
			}
		}
		return "0"
	}

	// floatStr must contain a .
	// all decimal part must have same length
	// floatToInt("3.000100", 1e6) => 3000100
	// "3.0001" is not allowed
	floatToInt := func(floatStr string, shift10 int) int {
		splits := strings.SplitN(floatStr, ".", 2)
		i, _ := strconv.Atoi(splits[0])
		d, _ := strconv.Atoi(splits[1])
		return i*shift10 + d
	}

	// floatToRat("1.2", 1e1) => (6, 5)
	floatToRat := func(floatStr string, shift10 int) (m, n int) {
		m = floatToInt(floatStr, shift10)
		n = shift10
		var g int // g:= calcGCD(m, n)
		m /= g
		n /= g
		return
	}

	// 括号拼接
	// 代码来源 https://codeforces.ml/gym/101341/problem/A
	// 类似题目 https://atcoder.jp/contests/abc167/tasks/abc167_f
	//         https://codeforces.ml/problemset/problem/1203/F1
	concatBrackets := func(ss [][]byte) (ids []int) {
		type pair struct{ x, y, i int }

		d := 0
		var ls, rs []pair
		for i, s := range ss {
			l, r := 0, 0
			for _, b := range s {
				if b == '(' {
					l++
				} else if l > 0 {
					l--
				} else {
					r++
				}
			}
			if r < l {
				ls = append(ls, pair{r, l, i})
			} else {
				rs = append(rs, pair{l, r, i})
			}
			d += l - r
		}

		sort.Slice(ls, func(i, j int) bool { return ls[i].x < ls[j].x })
		sort.Slice(rs, func(i, j int) bool { return rs[i].x < rs[j].x })
		f := func(ps []pair) []int {
			_ids := []int{}
			s := 0
			for _, p := range ps {
				if s < p.x {
					return nil
				}
				s += p.y - p.x
				_ids = append(_ids, p.i)
			}
			return _ids
		}
		idsL := f(ls)
		idsR := f(rs)
		if d != 0 || idsL == nil || idsR == nil {
			return
		}
		for _, id := range idsL {
			ids = append(ids, id)
		}
		for i := len(idsR) - 1; i >= 0; i-- {
			ids = append(ids, idsR[i])
		}
		return
	}

	sliceToStr := func(a []int) []byte {
		b := bytes.Buffer{}
		b.WriteByte('{')
		for i, v := range a {
			if i > 0 {
				b.WriteByte(',')
			}
			b.WriteString(strconv.Itoa(v))
		}
		b.WriteString("}\n")
		return b.Bytes()
	}

	getMapRangeValues := func(m map[int]int, l, r int) (a []int) {
		for i := l; i <= r; i++ {
			v, ok := m[i]
			if !ok {
				v = -1
			}
			a = append(a, v)
		}
		return
	}

	_ = []interface{}{
		logInit,
		getCycle,
		ceilK, partition, maxValueStepToUpper, moveToRange,
		hash01Mat,
		smallK,
		removeLeadingZero,
		floatToRat,
		concatBrackets,
		sliceToStr,
		getMapRangeValues,
	}
}

// 逆序数
// LC面试题51 https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
func mergeCount(a []int) int64 {
	n := len(a)
	if n <= 1 {
		return 0
	}
	b := make([]int, n/2)
	c := make([]int, n-n/2)
	copy(b, a[:n/2])
	copy(c, a[n/2:])
	cnt := mergeCount(b) + mergeCount(c)
	ai, bi, ci := 0, 0, 0
	for ai < n {
		// 归并排序的同时计算逆序数
		if bi < len(b) && (ci == len(c) || b[bi] <= c[ci]) {
			a[ai] = b[bi]
			bi++
		} else {
			cnt += int64(n/2 - bi)
			a[ai] = c[ci]
			ci++
		}
		ai++
	}
	return cnt
}

// N 皇后 - 回溯法
// LC51 https://leetcode-cn.com/problems/n-queens/
// LC52 https://leetcode-cn.com/problems/n-queens-ii/

// 格雷码 http://oeis.org/A003188
// https://en.wikipedia.org/wiki/Gray_code
// LC89 https://leetcode-cn.com/problems/gray-code/
func grayCode(length int) []int {
	ans := make([]int, 1<<length)
	for i := range ans {
		ans[i] = i ^ i>>1
	}
	return ans
}

// 负二进制数相加
// LC1073/周赛139C https://leetcode-cn.com/problems/adding-two-negabinary-numbers/ https://leetcode-cn.com/contest/weekly-contest-139/
func addNegabinary(a1, a2 []int) []int {
	if len(a1) < len(a2) {
		a1, a2 = a2, a1
	}
	for i, j := len(a1)-1, len(a2)-1; j >= 0; {
		a1[i] += a2[j]
		i--
		j--
	}
	ans := append(make([]int, 2), a1...)
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] >= 2 {
			ans[i] -= 2
			if ans[i-1] >= 1 {
				ans[i-1]--
			} else {
				ans[i-1]++
				ans[i-2]++
			}
		}
	}
	for i, v := range ans {
		if v != 0 {
			return ans[i:]
		}
	}
	return []int{0}
}

// 负二进制转换
// LC1017/周赛130B https://leetcode-cn.com/problems/convert-to-base-2/ https://leetcode-cn.com/contest/weekly-contest-130/
func toNegabinary(n int) (res string) {
	if n == 0 {
		return "0"
	}
	for ; n != 0; n = -(n >> 1) {
		res = string('0'+n&1) + res
	}
	return
}

// 分数转小数
// https://en.wikipedia.org/wiki/Repeating_decimal
// Period of decimal representation of 1/n, or 0 if 1/n terminates https://oeis.org/A051626
// The periodic part of the decimal expansion of 1/n https://oeis.org/A036275
// 例如 (2, -3) => ("-0.", "6")
// b must not be zero
// LC166 https://leetcode-cn.com/problems/fraction-to-recurring-decimal/
func fractionToDecimal(a, b int64) (beforeCycle, cycle []byte) {
	if a == 0 {
		return []byte{'0'}, nil
	}
	var res []byte
	if a < 0 && b > 0 || a > 0 && b < 0 {
		res = []byte{'-'}
	}
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	res = append(res, strconv.FormatInt(a/b, 10)...)

	r := a % b
	if r == 0 {
		return res, nil
	}
	res = append(res, '.')

	posMap := map[int64]int{}
	for r != 0 {
		if pos, ok := posMap[r]; ok {
			return res[:pos], res[pos:]
		}
		posMap[r] = len(res)
		r *= 10
		res = append(res, '0'+byte(r/b))
		r %= b
	}
	return res, nil
}

// 小数转分数
// decimal like "2.15(376)", which means "2.15376376376..."
// https://zh.wikipedia.org/wiki/%E5%BE%AA%E7%8E%AF%E5%B0%8F%E6%95%B0#%E5%8C%96%E7%82%BA%E5%88%86%E6%95%B8%E7%9A%84%E6%96%B9%E6%B3%95
func decimalToFraction(decimal string) (a, b int64) {
	r := regexp.MustCompile(`(?P<integerPart>\d+)\.?(?P<nonRepeatingPart>\d*)\(?(?P<repeatingPart>\d*)\)?`)
	match := r.FindStringSubmatch(decimal)
	integerPart, nonRepeatingPart, repeatingPart := match[1], match[2], match[3]
	intPartNum, _ := strconv.ParseInt(integerPart, 10, 64)
	if repeatingPart == "" {
		repeatingPart = "0"
	}
	b, _ = strconv.ParseInt(strings.Repeat("9", len(repeatingPart))+strings.Repeat("0", len(nonRepeatingPart)), 10, 64)
	a, _ = strconv.ParseInt(nonRepeatingPart+repeatingPart, 10, 64)
	if nonRepeatingPart != "" {
		v, _ := strconv.ParseInt(nonRepeatingPart, 10, 64)
		a -= v
	}
	a += intPartNum * b
	// 后续需要用 gcd 化简
	// 或者用 return big.NewRat(a, b)
	return
}

// 表达式计算（无括号）
// LC227 https://leetcode-cn.com/problems/basic-calculator-ii/
func calculate(s string) (ans int) {
	s = strings.ReplaceAll(s, " ", "")
	v, sign, stack := 0, '+', []int{}
	for i, b := range s {
		if '0' <= b && b <= '9' {
			v = v*10 + int(b-'0')
			if i+1 < len(s) {
				continue
			}
		}
		switch sign {
		case '+':
			stack = append(stack, v)
		case '-':
			stack = append(stack, -v)
		case '*':
			w := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, w*v)
		default: // '/'
			w := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			stack = append(stack, w/v)
		}
		v = 0
		sign = b
	}
	for _, v := range stack {
		ans += v
	}
	return
}
