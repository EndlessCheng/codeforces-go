package copypasta

import (
	"bytes"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

/* 其他无法分类的算法

三维 n 皇后 https://oeis.org/A068940
Maximal number of chess queens that can be placed on a 3-dimensional chessboard of order n so that no two queens attack each other

Smallest positive integer k such that n = +-1+-2+-...+-k for some choice of +'s and -'s https://oeis.org/A140358
相关题目 https://codeforces.com/problemset/problem/1278/B

Numbers n such that n is the substring identical to the least significant bits of its base 2 representation.
https://oeis.org/A181891
https://oeis.org/A181929 前缀
https://oeis.org/A038102 子串

Maximal number of regions obtained by joining n points around a circle by straight lines.
Also number of regions in 4-space formed by n-1 hyperplanes.
a(n) = n*(n-1)*(n*n-5*n+18)/24+1 https://oeis.org/A000127

https://oeis.org/A001069 Log2*(n) (version 2): take log_2 of n this many times to get a number < 2
https://oeis.org/A211667 Number of iterations sqrt(sqrt(sqrt(...(n)...))) such that the result is < 2
    a(n) = 1, 2, 3, 4, 5, ... for n = 2^1, 2^2, 2^4, 2^8, 2^16, ..., i.e., n = 2, 4, 16, 256, 65536, ... = https://oeis.org/A001146

https://oeis.org/A002024 n appears n times; a(n) = floor(sqrt(2n) + 1/2) https://www.zhihu.com/question/25045244

找规律 https://codeforces.com/problemset/problem/1034/B
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

	// max record pos
	// 相关题目 https://codeforces.com/problemset/problem/1381/B
	recordPos := func(a []int) []int {
		pos := []int{0}
		for i, v := range a {
			if v > a[pos[len(pos)-1]] {
				pos = append(pos, i)
			}
		}
		//pos = append(pos, len(a))
		return pos
	}

	ceilK := func(n, k int) int {
		if n%k == 0 {
			return n
		}
		return n + k - n%k
	}

	shift1Mod := func(v, mod int) int {
		return (v-1)%mod + 1
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
			if q.Len() < k || v < q.IntSlice[0] {
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
	// https://www.luogu.com.cn/problem/UVA10555
	floatToRat := func(floatStr string, shift10 int) (m, n int) {
		m = floatToInt(floatStr, shift10)
		n = shift10
		var g int // g := gcd(m, n)
		m /= g
		n /= g
		return
	}

	isInt := func(x float64) bool {
		const eps = 1e-8
		return math.Abs(x-math.Round(x)) < eps
	}

	// 括号拼接
	// 代码来源 https://codeforces.com/gym/101341/problem/A
	// 类似题目 https://atcoder.jp/contests/abc167/tasks/abc167_f
	//         https://codeforces.com/problemset/problem/1203/F1
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

	// 螺旋矩阵 Spiral Matrix
	// https://ac.nowcoder.com/acm/contest/6489/C
	// 另：只考虑枚举顺序 LC54 https://leetcode-cn.com/problems/spiral-matrix/
	genSpiralMatrix := func(n, m int) [][]int {
		mat := make([][]int, n)
		for i := range mat {
			mat[i] = make([]int, m)
			for j := range mat[i] { // 如果从 1 开始这里可以不要，下面的 != -1 改成 > 0
				mat[i][j] = -1
			}
		}
		type pair struct{ x, y int }
		dir4 := [4]pair{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} // 右下左上
		x, y, di := 0, 0, 0
		//pos := make([]pair, n*m+1)
		for i := 0; i < n*m; i++ { // 从 0 到 n*m-1
			mat[x][y] = i
			//pos[i] = pair{x, y}
			d := dir4[di&3]
			if xx, yy := x+d.x, y+d.y; xx < 0 || xx >= n || yy < 0 || yy >= m || mat[xx][yy] != -1 {
				di++
			}
			d = dir4[di&3]
			x += d.x
			y += d.y
		}
		return mat
	}

	// 最小栈，支持动态 push pop，查询栈中最小元素
	// 思路是用另一个栈，同步 push pop，处理 push 时压入 min(当前元素,栈顶元素)，注意栈为空的时候直接压入元素
	// https://ac.nowcoder.com/acm/contest/1055/A
	// https://blog.nowcoder.net/n/ceb3214b89594af481ef9b794c75a929

	_ = []interface{}{
		logInit,
		getCycle,
		recordPos,
		ceilK, shift1Mod, partition, maxValueStepToUpper, moveToRange,
		hash01Mat,
		smallK,
		removeLeadingZero,
		floatToRat,
		isInt,
		concatBrackets,
		sliceToStr,
		getMapRangeValues,
		genSpiralMatrix,
	}
}

// 逆序对
// LC 面试题 51 https://leetcode-cn.com/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// EXTRA: https://leetcode-cn.com/problems/count-of-range-sum/
// 一张关于归并排序的好图 https://www.cnblogs.com/chengxiao/p/6194356.html
func mergeCount(a []int) int64 {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := append([]int(nil), a[:n/2]...)
	right := append([]int(nil), a[n/2:]...)
	cnt := mergeCount(left) + mergeCount(right)
	l, r := 0, 0
	for i := range a {
		// 归并排序的同时计算逆序对
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			a[i] = left[l]
			l++
		} else {
			cnt += int64(n/2 - l)
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

// 状压 N 皇后
// LC51 https://leetcode-cn.com/problems/n-queens/
// LC52 https://leetcode-cn.com/problems/n-queens-ii/
func totalNQueens(n int) (ans int) {
	var f func(row, columns, diagonals1, diagonals2 int)
	f = func(row, columns, diagonals1, diagonals2 int) {
		if row == 1 {
			ans++
			return
		}
		availablePositions := (1<<n - 1) &^ (columns | diagonals1 | diagonals2)
		for availablePositions > 0 {
			position := availablePositions & -availablePositions
			f(row+1, columns|position, (diagonals1|position)<<1, (diagonals2|position)>>1)
			availablePositions &^= position // 移除该比特位
		}
	}
	f(0, 0, 0, 0)
	return
}

// 格雷码 https://oeis.org/A003188 https://oeis.org/A014550
// https://en.wikipedia.org/wiki/Gray_code
// LC89 https://leetcode-cn.com/problems/gray-code/
// 转换 https://codeforces.com/problemset/problem/1419/E
func grayCode(length int) []int {
	ans := make([]int, 1<<length)
	for i := range ans {
		ans[i] = i ^ i>>1
	}
	return ans
}

// 输入两个无重复元素的序列，返回通过交换相邻元素，从 a 到 b 所需的最小交换次数
// 保证 a b 包含相同的元素
func countSwap(a, b []int) int {
	// 可能要事先 copy 一份 a
	// 暴力法
	ans := 0
	for _, tar := range b {
		for i, v := range a {
			if v == tar {
				ans += i
				a = append(a[:i], a[i+1:]...)
				break
			}
		}
	}
	return ans
}

// 输入一个仅包含 () 的括号串，返回右括号个数不少于左括号个数的非空子串个数
func countValidSubstring(s string) (ans int) {
	cnt := map[int]int{0: 1}
	leSum := 1 // less equal than v
	v := 0
	for _, b := range s {
		if b == '(' {
			leSum -= cnt[v]
			v--
		} else {
			v++
			leSum += cnt[v]
		}
		ans += leSum
		cnt[v]++
		leSum++
	}
	return
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
		res = string(byte('0'+n&1)) + res
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
