package copypasta

import (
	"bytes"
	"math"
	"regexp"
	"slices"
	"sort"
	"strconv"
	"strings"
)

/* 其他无法分类的算法 ad-hoc

- [764. 最大加号标志](https://leetcode.cn/problems/largest-plus-sign/) 1753
最大子矩形 https://www.luogu.com.cn/problem/P1578

倒数平方根 https://www.bilibili.com/video/BV17N41167dR/

小奥
https://codeforces.com/problemset/problem/700/A

消消乐
- [954. 二倍数对数组](https://leetcode.cn/problems/array-of-doubled-pairs/) 1548
- [2007. 从双倍数组中还原原数组](https://leetcode.cn/problems/find-original-array-from-doubled-array/) 1557

原地哈希
LC442 https://leetcode.cn/problems/find-all-duplicates-in-an-array/

三维 n 皇后 https://oeis.org/A068940
Maximal number of chess queens that can be placed on a 3-dimensional chessboard of order n so that no two queens attack each other

Smallest positive integer k such that n = +-1+-2+-...+-k for some choice of +'s and -'s https://oeis.org/A140358
https://codeforces.com/problemset/problem/11/B
https://codeforces.com/problemset/problem/1278/B
LC754 https://leetcode.cn/problems/reach-a-number/ 1880

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

长度为 n 的所有二进制串，最多能划分出的 11 的个数之和 https://oeis.org/A045883
相关题目 https://codeforces.com/contest/1511/problem/E

https://oeis.org/A007302 Optimal cost function between two processors at distance n
bits.OnesCount(3*n ^ n)
LC2571 https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0/
解释 https://leetcode.cn/problems/minimum-operations-to-reduce-an-integer-to-0/solution/ji-yi-hua-sou-suo-by-endlesscheng-cm6l/

4 汉诺塔 https://oeis.org/A007664
Reve's puzzle: number of moves needed to solve the Towers of Hanoi puzzle with 4 pegs and n disks, according to the Frame-Stewart algorithm
https://www.acwing.com/problem/content/description/98/

麻将
2021·昆明 https://ac.nowcoder.com/acm/contest/12548/K

五子棋
https://codeforces.com/contest/825/problem/B

调度场算法 shunting-yard algorithm
中缀转后缀
https://en.wikipedia.org/wiki/Shunting-yard_algorithm
*/
func miscCollection() {
	// debug 用
	toArray := func(a []int) (res [100]int) {
		for i, v := range a {
			res[i] = v
		}
		return
	}

	// 预处理 log 的整数部分
	logInit := func() {
		const mx int = 1e6
		log := [mx + 1]int{} // log[0] 未定义，请勿访问
		for i := 2; i <= mx; i++ {
			log[i] = log[i>>1] + 1
		}
	}

	// 找环
	// 1<=next[i]<=n
	// 相关题目 https://atcoder.jp/contests/abc167/tasks/abc167_d
	// EXTRA: 周期追逐 https://codeforces.com/problemset/problem/547/A
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

	// Floyd 判圈算法
	// https://zh.wikipedia.org/wiki/Floyd%E5%88%A4%E5%9C%88%E7%AE%97%E6%B3%95
	// https://en.wikipedia.org/wiki/Cycle_detection
	// https://codeforces.com/problemset/problem/1137/D
	// 设环长为 c，链长为 t，则快慢指针相遇时，慢指针在环上走过的距离为 c-t%c（具体证明见 CF1137D 这题的题解）

	// max record pos
	// 相关题目（这也是一道好题）https://codeforces.com/problemset/problem/1381/B
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

	// 小结论：把 n 用 m 等分，得到 m-n%m 个 n/m 和 n%m 个 n/m+1
	// 相关题目 https://codeforces.com/problemset/problem/663/A
	partition := func(n, m int) (q, cntQ, cntQ1 int) {
		// m must > 0
		return n / m, m - n%m, n % m
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

	// 适用于需要频繁读取 a 中元素到一个 map 中的情况
	// 调用 quickHashMapRead(a) 后
	// 原来的类似 cnt[a[i]]++ 的操作，可以让 cnt 由 map[int]int 改为 make([]int, len(rk))
	// 若需要访问 a[i] 原有元素，可以访问 origin[a[i]]
	// 这样后续操作就与 map 无关了
	quickHashMapRead := func(a []int) ([]int, int) {
		origin := make([]int, len(a))
		rk := map[int]int{}
		for i, v := range a {
			if _, has := rk[v]; !has {
				rk[v] = len(rk)
				origin[rk[v]] = v
			}
			a[i] = rk[v]
		}
		return origin, len(rk)
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

	// 01 矩阵，每个 1 位置向四个方向延伸连续 1 的最远距离
	// Kick Start 2021 Round A L Shaped Plots https://codingcompetitions.withgoogle.com/kickstart/round/0000000000436140/000000000068c509
	max1dir4 := func(a [][]int) (ls, rs, us, ds [][]int) {
		n, m := len(a), len(a[0])
		ls, rs = make([][]int, n), make([][]int, n)
		for i, row := range a {
			ls[i] = make([]int, m)
			for j, v := range row {
				if v == 0 {
					continue
				}
				if j == 0 || row[j-1] == 0 {
					ls[i][j] = j
				} else {
					ls[i][j] = ls[i][j-1]
				}
			}
			rs[i] = make([]int, m)
			for j := m - 1; j >= 0; j-- {
				if row[j] == 0 {
					continue
				}
				if j == m-1 || row[j+1] == 0 {
					rs[i][j] = j
				} else {
					rs[i][j] = rs[i][j+1]
				}
			}
		}

		us, ds = make([][]int, n), make([][]int, n)
		for i := range us {
			us[i] = make([]int, m)
			ds[i] = make([]int, m)
		}
		for j := 0; j < m; j++ {
			for i, row := range a {
				if row[j] == 0 {
					continue
				}
				if i == 0 || a[i-1][j] == 0 {
					us[i][j] = i
				} else {
					us[i][j] = us[i-1][j]
				}
			}
			for i := n - 1; i >= 0; i-- {
				if a[i][j] == 0 {
					continue
				}
				if i == n-1 || a[i+1][j] == 0 {
					ds[i][j] = i
				} else {
					ds[i][j] = ds[i+1][j]
				}
			}
		}

		return
	}

	// a 是环形，若相邻元素 (v,w) 符合某种条件，则合并，删除 w
	// 在 a 上不断循环合并直至没有可以合并的相邻元素，返回删除的元素
	// 相关题目 https://codeforces.com/problemset/problem/1483/B
	loopMergeOnRing := func(a []int, canMerge func(v, w int) bool) (deletedElements []int) {
		n := len(a)
		r := make([]int, n)
		for i := 0; i < n-1; i++ {
			r[i] = i + 1
		}

		q := []int{}
		for i, v := range a {
			if canMerge(v, a[r[i]]) {
				q = append(q, i)
			}
		}
		del := make([]bool, n)
		for len(q) > 0 {
			i := q[0]
			q = q[1:]
			if del[i] {
				continue
			}
			if !del[r[i]] {
				deletedElements = append(deletedElements, r[i]) // +1
				del[r[i]] = true
			}
			r[i] = r[r[i]]
			if canMerge(a[i], a[r[i]]) {
				q = append(q, i)
			}
		}
		return
	}

	// 最小栈，支持动态 push pop，查询栈中最小元素
	// 思路是用另一个栈，同步 push pop，处理 push 时压入 min(当前元素,栈顶元素)，注意栈为空的时候直接压入元素
	// https://ac.nowcoder.com/acm/contest/1055/A
	// https://blog.nowcoder.net/n/ceb3214b89594af481ef9b794c75a929

	_ = []interface{}{
		toArray,
		logInit,
		getCycle,
		recordPos,
		partition,
		smallK,
		floatToRat,
		isInt,
		quickHashMapRead,
		concatBrackets,
		sliceToStr,
		getMapRangeValues,
		max1dir4,
		loopMergeOnRing,
	}
}

// b 是 a 的一个排列（允许有重复元素）
// 返回 b 中各个元素在 a 中的下标（重复的元素顺序保持一致）
// 可用于求从 a 变到 b 需要的相邻位元素交换的最小次数，即返回结果的逆序对个数
// LC1850 https://leetcode.cn/problems/minimum-adjacent-swaps-to-reach-the-kth-smallest-number/
func mapPos(a, b []int) []int {
	pos := map[int][]int{}
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}
	ids := make([]int, len(b))
	for i, v := range b {
		ids[i] = pos[v][0]
		pos[v] = pos[v][1:]
	}
	return ids
}

/* 归并排序与逆序对
一张关于归并排序的好图 https://www.cnblogs.com/chengxiao/p/6194356.html
- 归并排序：[912. 排序数组](https://leetcode.cn/problems/sort-an-array/)
- [LCR 170. 交易逆序对的总数](https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/)
- [327. 区间和的个数](https://leetcode.cn/problems/count-of-smaller-numbers-after-self/)
- [327. 区间和的个数](https://leetcode.cn/problems/count-of-range-sum/)
- [493. 翻转对](https://leetcode.cn/problems/reverse-pairs/)
https://atcoder.jp/contests/arc075/tasks/arc075_c
*/
func mergeCount(a []int) int {
	n := len(a)
	if n <= 1 {
		return 0
	}
	left := slices.Clone(a[:n/2])
	right := slices.Clone(a[n/2:])
	cnt := mergeCount(left) + mergeCount(right)
	l, r := 0, 0
	for i := range a {
		// 归并排序的同时计算逆序对
		if l < len(left) && (r == len(right) || left[l] <= right[r]) {
			a[i] = left[l]
			l++
		} else {
			cnt += n/2 - l
			a[i] = right[r]
			r++
		}
	}
	return cnt
}

// 状压 N 皇后
// https://oeis.org/A000170 Number of ways of placing n non-attacking queens on an n X n board
// Strong conjecture: there is a constant c around 2.54 such that a(n) is asymptotic to n!/c^n
// Weak conjecture: lim_{n -> infinity} (1/n) * log(n!/a(n)) = constant = 0.90....
// https://arxiv.org/pdf/2107.13460.pdf
// LC51 https://leetcode.cn/problems/n-queens/
// LC52 https://leetcode.cn/problems/n-queens-ii/
func totalNQueens(n int) (ans int) {
	var dfs func(int, int, int, int)
	dfs = func(row, columns, diagonals1, diagonals2 int) {
		if row == 1 {
			ans++
			return
		}
		availablePositions := (1<<n - 1) &^ (columns | diagonals1 | diagonals2)
		for availablePositions > 0 {
			position := availablePositions & -availablePositions
			dfs(row+1, columns|position, (diagonals1|position)<<1, (diagonals2|position)>>1)
			availablePositions &^= position // 移除该比特位
		}
	}
	dfs(0, 0, 0, 0)
	return
}

// 格雷码 https://oeis.org/A003188 https://oeis.org/A014550
// https://en.wikipedia.org/wiki/Gray_code
// LC89 https://leetcode.cn/problems/gray-code/
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
// LC1073 https://leetcode.cn/problems/adding-two-negabinary-numbers/
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
// https://atcoder.jp/contests/abc105/tasks/abc105_c
// LC1017 https://leetcode.cn/problems/convert-to-base-2/
func toNegabinary(n int) (ans string) {
	if n == 0 {
		return "0"
	}
	for ; n != 0; n = -(n >> 1) {
		ans = string(byte('0'+n&1)) + ans
	}
	return
}

// 分数转小数
// https://en.wikipedia.org/wiki/Repeating_decimal
// Number of digits in decimal expansion of 1/n before the periodic part begins https://oeis.org/A051628
// - 设 n=2^c2*5^c5*...，那么 A051628(n) = max(c2,c5)
// Period of decimal representation of 1/n, or 0 if 1/n terminates https://oeis.org/A051626
// - 如果 n 的质因子只有 2 和 5，那么不存在循环节
// - 否则，先把 n 移除所有质因子 2 和 5，得到 m，那么 A051626(n) = n_order(10, m) 请看 math.go 中的「阶」
// - 参考 https://zhuanlan.zhihu.com/p/346536813 https://www.zhihu.com/question/462266812
// The periodic part of the decimal expansion of 1/n https://oeis.org/A036275
// 例如 (2, -3) => ("-0.", "6")
// b must not be zero
//
// https://oeis.org/A007732 Period of decimal representation of 1/n
// https://oeis.org/A084680 Order of 10 modulo n [i.e., least m such that 10^m = 1 (mod n)] or 0 when no such number exists
// https://oeis.org/A002329 Periods of reciprocals of integers prime to 10
//
// LC166 https://leetcode.cn/problems/fraction-to-recurring-decimal/
// WF1990 https://www.luogu.com.cn/problem/UVA202
// 1e12 加强版 https://ac.nowcoder.com/acm/contest/62622/E
// Python 代码 https://ac.nowcoder.com/acm/contest/view-submission?submissionId=63288994
// b 进制下是否为有限小数 https://codeforces.com/problemset/problem/983/A 1700
func fractionToDecimal(a, b int) (beforeCycle, cycle []byte) {
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
	res = append(res, strconv.Itoa(a/b)...)

	r := a % b
	if r == 0 {
		return res, nil
	}
	res = append(res, '.')

	posMap := map[int]int{}
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
func decimalToFraction(decimal string) (a, b int) {
	r := regexp.MustCompile(`(?P<integerPart>\d+)\.?(?P<nonRepeatingPart>\d*)\(?(?P<repeatingPart>\d*)\)?`)
	match := r.FindStringSubmatch(decimal)
	integerPart, nonRepeatingPart, repeatingPart := match[1], match[2], match[3]
	intPartNum, _ := strconv.Atoi(integerPart)
	if repeatingPart == "" {
		repeatingPart = "0"
	}
	b, _ = strconv.Atoi(strings.Repeat("9", len(repeatingPart)) + strings.Repeat("0", len(nonRepeatingPart)))
	a, _ = strconv.Atoi(nonRepeatingPart + repeatingPart)
	if nonRepeatingPart != "" {
		v, _ := strconv.Atoi(nonRepeatingPart)
		a -= v
	}
	a += intPartNum * b
	// 后续需要用 gcd 化简
	// 或者用 return big.NewRat(a, b)
	return
}

// 表达式计算（无括号）
// LC227 https://leetcode.cn/problems/basic-calculator-ii/
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

// 倒序思想
// 来源自被删除的 C 题 https://ac.nowcoder.com/acm/contest/view-submission?submissionId=45798625
// 有一个大小为 n*m 的网格图，和一个长为 n*m 的目标位置列表，每个位置表示网格图中的一个格点且互不相同
// 网格图初始为空
// 按照此列表的顺序，一个士兵从网格图边缘的任意位置进入并到达目标位置，到达后该士兵停留在此格点，下一个士兵开始进入网格图
// 每个士兵会尽可能地避免经过有士兵的格点
// 输出每个士兵必须经过的士兵数之和
// n, m <= 500
// 思路：
// 将问题转化成从最后一个士兵开始倒着退出网格
// 对于一个填满的网格图，每个士兵到边缘的最短路径就是离他最近的边缘的距离
// 当一个士兵退出网格后，BFS 地更新他周围的士兵到边缘的最短路径（空格点为 0，有人的格点为 1）
// 复杂度 O((n+m)*min(n,m)^2)
func minMustPassSum(n, m int, targetCells [][2]int) int {
	dis := make([][]int, n)
	filled := make([][]int, n) // 格子是否有人
	inQ := make([][]bool, n)
	for i := range dis {
		dis[i] = make([]int, m)
		filled[i] = make([]int, m)
		for j := range dis[i] {
			dis[i][j] = min(i, n-1-i, j, m-1-j)
			filled[i][j] = 1
		}
		inQ[i] = make([]bool, m)
	}

	ans := 0
	type pair struct{ x, y int }
	dir4 := []pair{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := len(targetCells) - 1; i >= 0; i-- {
		p := targetCells[i]
		x, y := p[0], p[1]
		//x--
		//y--
		ans += dis[x][y]
		filled[x][y] = 0
		q := []pair{{x, y}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, y := p.x, p.y
			inQ[x][y] = false
			for _, d := range dir4 {
				if xx, yy := x+d.x, y+d.y; 0 <= xx && xx < n && 0 <= yy && yy < m && dis[x][y]+filled[x][y] < dis[xx][yy] {
					dis[xx][yy] = dis[x][y] + filled[x][y]
					if !inQ[xx][yy] {
						inQ[xx][yy] = true
						q = append(q, pair{xx, yy})
					}
				}
			}
		}
	}
	return ans
}

// 马走日从 (0,0) 到 (x,y) 所需最小步数
// 无边界 LC1197 https://leetcode.cn/problems/minimum-knight-moves/
// http://poj.org/problem?id=1915
// 有边界+打印方案 https://www.acwing.com/problem/content/3527/
/*
# https://leetcode.cn/problems/maximum-number-of-moves-to-kill-all-pawns/
def minKnightMoves(self, x: int, y: int) -> int:
   x, y = abs(x), abs(y)
   x, y = min(x, y), max(x, y)  # swap
   if x == 0 and y == 1:
	   return 3
   if x == y == 2:
	   return 4
   if y > 2 * x:
	   x += (y - 2 * x + 1) // 4 * 2
   return (x + y) // 3 + (x + y) % 3
*/
func minKnightMoves(x, y int, abs func(int) int) int {
	// TODO 算法来源？
	x, y = abs(x), abs(y)
	if x == 2 && y == 2 {
		return 4
	}
	if x+y == 1 {
		return 3
	}
	ans := max((x+1)/2, (y+1)/2, (x+y+2)/3)
	ans += (ans ^ x ^ y) & 1
	return ans
}

// 网格图
// 从 (sx,sy) 出发，向右下走，遇到边界反弹，返回到达 (tx,ty) 的最小步数
// 若无法到达，返回 -1
func minDiagonalMove(n, m, sx, sy, tx, ty int) (step int) {
	vis := make([][][2][2]bool, n)
	for i := range vis {
		vis[i] = make([][2][2]bool, m)
	}
	dx, dy := 1, 1
	vis[sx][sy][dx][dy] = true
	x, y := sx, sy
	for x != tx || y != ty {
		step++
		xx, yy := x+dx, y+dy
		if xx < 0 || xx >= n || yy < 0 || yy >= m {
			if xx < 0 || xx >= n {
				dx = -dx
			}
			if yy < 0 || yy >= m {
				dy = -dy
			}
			//xx, yy = x+dx, y+dy // 有的题目直接反弹
			xx, yy = x, y // 有的题目停一步
		}
		x, y = xx, yy
		if vis[x][y][(dx+1)/2][(dy+1)/2] {
			return -1
		}
		vis[x][y][(dx+1)/2][(dy+1)/2] = true
	}
	return
}

// 判断 6 个矩形是否为长方体的 6 个面
// NEERC04 https://www.luogu.com.cn/problem/UVA1587
func isCuboid(rect [][2]int) bool {
	for i, r := range rect {
		if r[0] > r[1] {
			rect[i] = [2]int{r[1], r[0]}
		}
	}
	sort.Slice(rect, func(i, j int) bool { a, b := rect[i], rect[j]; return a[0] < b[0] || a[0] == b[0] && a[1] < b[1] })
	for i := 0; i < 6; i += 2 {
		if rect[i] != rect[i+1] { // NOTE: [2]
			return false
		}
	}
	y0, y2 := rect[0][1], rect[2][1]
	return rect[2][0] == rect[0][0] && (rect[4] == [2]int{y0, y2} || rect[4] == [2]int{y2, y0})
}

// 约瑟夫问题
// 思路：用递推公式，自底向上计算
// https://zh.wikipedia.org/wiki/%E7%BA%A6%E7%91%9F%E5%A4%AB%E6%96%AF%E9%97%AE%E9%A2%98
// https://oi-wiki.org/misc/josephus/ 注意当 k 较小时，存在 O(klogn) 的做法
// https://www.scirp.org/pdf/OJDM_2019101516120841.pdf Generalizations of the Feline and Texas Chainsaw Josephus Problems
//
// LCR187 https://leetcode.cn/problems/yuan-quan-zhong-zui-hou-sheng-xia-de-shu-zi-lcof/
// LC1823 https://leetcode.cn/problems/find-the-winner-of-the-circular-game/
// https://codeforces.com/gym/101955/problem/K
func josephusProblem(n, k int) int {
	cur := 0
	for i := 2; i <= n; i++ {
		cur = (cur + k) % i
	}
	return cur + 1 // 1-index
}

// 均分纸牌 https://www.luogu.com.cn/problem/P1031
// 环形 https://www.luogu.com.cn/problem/P2512 https://www.luogu.com.cn/problem/P3051 https://www.luogu.com.cn/problem/P4016 UVa11300 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=25&page=show_problem&problem=2275
// 环形+打印方案 https://www.luogu.com.cn/problem/P2125
// 二维环形 https://www.acwing.com/problem/content/107/
func minMoveToAllSameInCircle(a []int, abs func(int) int) (ans int) {
	n := len(a)
	avg := 0
	for _, v := range a {
		avg += v
	}
	if avg%n != 0 {
		return -1
	}
	avg /= n
	sum := make([]int, n)
	sum[0] = a[0] - avg
	for i := 1; i < n; i++ {
		sum[i] = sum[i-1] + a[i] - avg
	}
	slices.Sort(sum) // 也可以用快速选择求中位数
	mid := sum[n/2]
	for _, v := range sum {
		ans += abs(v - mid)
	}
	return
}

// 表达式转表达式树
// https://leetcode.cn/submissions/detail/186220993/
func parseExpression(s string) {
	s = strings.TrimSpace(s)
	type node struct {
		lo, ro *node
		res    int
		op     byte
	}

	n := len(s)
	left := make([]int, n)
	stk := []int{}
	for i := n - 1; i >= 0; i-- {
		if s[i] == ')' {
			stk = append(stk, i)
		} else if s[i] == '(' {
			left[stk[len(stk)-1]] = i
			stk = stk[:len(stk)-1]
		}
	}
	var parse func(l, r int) *node
	parse = func(l, r int) *node {
		o := &node{}
		// 因为表达式是左结合的，我们需要从右向左构造这棵表达式树
		if s[r] == ')' {
			ll := left[r]
			ro := parse(ll+1, r-1)
			if ll == l {
				return ro
			}
			o.ro = ro
			o.op = s[ll-1]
			o.lo = parse(l, ll-2)
		} else {
			ro := &node{res: int(s[r] & 15)} // 单个数字
			if l == r {
				return ro
			}
			o.ro = ro
			o.op = s[r-1]
			o.lo = parse(l, r-2)
		}
		//calc(o) // 计算表达式
		return o
	}
	root := parse(0, n-1)
	_ = root
}

// 钱珀瑙恩数 Champernowne constant
// https://en.wikipedia.org/wiki/Champernowne_constant
// https://oeis.org/A033307
// 返回第 k 位数字
// https://leetcode.cn/contest/espressif-2021/problems/fSghVj/
func champernowneConstant(k int) int {
	for i, p10 := 1, 10; ; i++ {
		if i*p10 > k {
			return int(strconv.Itoa(k / i)[k%i] & 15)
		}
		k += p10
		p10 *= 10
	}
}

// 摩尔投票法求绝对众数（absolute mode, majority）
// Boyer–Moore majority vote algorithm
// https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
// LC169 https://leetcode.cn/problems/majority-element/
// LC229 变形 https://leetcode.cn/problems/majority-element-ii/
// LC2780 https://leetcode.cn/problems/minimum-index-of-a-valid-split/
func majorityVote(a []int) (mode int) {
	cnt := 0
	for _, v := range a {
		if cnt == 0 {
			mode = v
		}
		if v == mode {
			cnt++
		} else {
			cnt--
		}
	}
	return
}

// 给出二维平面上的 n 个坐标点，以 (x,y) 为中心的十字最多能覆盖多少个点？
// https://atcoder.jp/contests/abc176/tasks/abc176_e
// 进阶：每个点有不同的点权 https://atcoder.jp/contests/abc298/tasks/abc298_f
func maxCover(a [][3]int) (ans int) {
	type pair struct{ r, c int }
	grid := make(map[pair]int, len(a))
	rowSum := map[int]int{}
	colSum := map[int]int{}
	for _, p := range a {
		r, c, v := p[0], p[1], p[2]
		grid[pair{r, c}] = v
		rowSum[r] += v
		colSum[c] += v
	}

	type cs struct{ c, s int }
	colList := make([]cs, 0, len(colSum))
	for x, s := range colSum {
		colList = append(colList, cs{x, s})
	}
	sort.Slice(colList, func(i, j int) bool { return colList[i].s > colList[j].s })

	for r, s := range rowSum {
		for _, cs := range colList {
			v, ok := grid[pair{r, cs.c}] // 每个点至多访问一次
			if !ok {
				ans = max(ans, s+cs.s)
				break // 保证时间复杂度是 O(n) 的关键：colList 从大到小排序，后面的只会更小，无需遍历
			}
			ans = max(ans, s+cs.s-v)
		}
	}
	return
}

// 至多 k 次操作，每次操作把一个数减一，目标是最小化数组最大值
// 不能把 a[i] 改成负数
// https://codeforces.com/problemset/problem/960/B
func minMaxArray(a []int, k int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	a = append(a, 0) // 哨兵

	sum := 0
	for i := 0; a[i] > 0; {
		v := a[i]
		for ; a[i] == v; i++ {
			sum += v
		}
		if sum-k >= a[i]*i {
			sum -= k
			h, ex := sum/i, sum%i
			// ex 个 h+1，i-ex 个 h，其余不变
			for j := 0; j < ex; j++ {
				a[j] = h + 1
			}
			for j := ex; j < i; j++ {
				a[j] = h
			}
			return
		}
	}
	// 全为 0
	for i := range a {
		a[i] = 0
	}
}

// 将 "aa...abb...b" (a 个 'a' 和 b 个 'b') 中的 b 向左移动 k 次，得到的字典序最大的字符串是什么？
func moveAB(a, b, k int) string {
	if k > a*b { // 非法
		return ""
	}
	const A, B = "a", "b"
	lb, ex := k/a, k%a
	if ex == 0 {
		return strings.Repeat(B, lb) + strings.Repeat(A, a) + strings.Repeat(B, b-lb)
	}
	return strings.Repeat(B, lb) + strings.Repeat(A, a-ex) + B + strings.Repeat(A, ex) + strings.Repeat(B, b-1-lb)
}

// https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/C
func twoApplesADay(a []int) int {
	n := len(a)
	if n == 1 {
		return 1
	}

	const low = 1
	slices.Sort(a)
o:
	for _, sum := range []int{a[0] + a[n-2], a[0] + a[n-1], a[1] + a[n-1]} {
		ans := low - 1
		l, r := 0, n-1
		for l < r {
			if a[l]+a[r] == sum {
				l++
				r--
				continue
			}
			if ans >= low {
				continue o
			}
			if a[l]+a[r] < sum {
				ans = sum - a[l]
				l++
			} else {
				ans = sum - a[r]
				r--
			}
			if ans < low {
				continue o
			}
		}
		if ans < low {
			ans = sum - a[l]
		}
		if ans >= low {
			return ans
		}
	}
	return -1
}

// 只能交换相邻的相差为 1 的数字（也可以交换 0 和 9）
// 输出最小字典序
//
// 设 a=s[0]，答案的首位只可能是 a-1 或 a（如果 a=9，那么可能和 0 互换；如果是 0，那么不需要互换，特判即可）
// 能不能是 a-1 呢？判断第一个 a-1 左边是否只有 a 和 a-2
// 然后讨论 a=s[1]，依此类推
// 用 10 个队列存对应数字出现的下标，每次取队首判断即可
// 时间复杂度 O(10n)：一次判断过程复杂度是 O(10)，一共有 n 次判断
//
// https://ac.nowcoder.com/acm/contest/65259/C
// 不能交换 0 和 9 的版本见后
func makeLexicographicallySmallestStringBySwappingAdjacentElements(s string, abs func(int) int) string {
	n := len(s)
	pos := [10][]int{}
	for i, b := range s {
		b -= '0'
		pos[b] = append(pos[b], i)
	}

	check := func(b byte) bool {
		if len(pos[b]) == 0 {
			return false
		}
		for i := byte(0); i <= 9; i++ {
			if i == b {
				continue
			}
			d := abs(int(i) - int(b))
			if d != 1 && d != 9 && len(pos[i]) > 0 && pos[i][0] < pos[b][0] {
				return false
			}
		}
		return true
	}

	ans := make([]byte, 0, n)
	used := make([]bool, n+1)
	for i := 0; i < n; {
		b := s[i] - '0'
		x := (b + 9) % 10 // b-1
		y := (b + 1) % 10
		if x > y {
			x, y = y, x
		}
		if x < b && check(x) {
			b = x
		} else if y < b && check(y) {
			b = y
		}

		ans = append(ans, '0'+b)
		used[pos[b][0]] = true
		pos[b] = pos[b][1:]

		for used[i] {
			i++
		}
	}
	return string(ans)
}

// 不能交换 0 和 9 的版本
func makeLexicographicallySmallestStringBySwappingAdjacentElements2(a []int) []int {
	const mx = 9
	pos := [mx + 1][]int{}
	for i, v := range a {
		pos[v] = append(pos[v], i)
	}

	check := func(v int) bool {
		if len(pos[v]) == 0 {
			return false
		}
		for i := 0; i <= mx; i++ {
			if (i < v-1 || i > v+1) && len(pos[i]) > 0 && pos[i][0] < pos[v][0] {
				return false
			}
		}
		return true
	}

	n := len(a)
	ans := make([]int, 0, n)
	used := make([]bool, n+1)
	for i := 0; i < n; {
		v := a[i]
		if v == 0 {
			continue
		}
		if check(v - 1) { // 把 v-1 换过来
			v--
		}

		ans = append(ans, v)
		used[pos[v][0]] = true
		pos[v] = pos[v][1:]

		for used[i] {
			i++
		}
	}
	return ans
}

// 解析时分
// - [1736. 替换隐藏数字得到的最晚时间](https://leetcode.cn/problems/latest-time-by-replacing-hidden-digits/) 1264
// - [3114. 替换字符可以得到的最晚时间](https://leetcode.cn/problems/latest-time-you-can-obtain-after-replacing-characters/) 1291
// - [2224. 转化时间需要的最少操作数](https://leetcode.cn/problems/minimum-number-of-operations-to-convert-time/) 1296
// - [2933. 高访问员工](https://leetcode.cn/problems/high-access-employees/) 1537
// 解析月日：LC2409 https://leetcode.cn/problems/count-days-spent-together/
func parseTime(s string) (hour, minute, total int) {
	hour = int(s[0]&15)*10 + int(s[1]&15)
	minute = int(s[3]&15)*10 + int(s[4]&15)
	total = hour*60 + minute
	return
}

// 区间覆盖
// 给定一些区间，从中选择尽量少的区间，覆盖一条指定线段 [s,t]。
// 变形说法：从 i 可以跳到 [i,i+a[i]] 中的任意整点
// 返回从 0 跳到 n-1 的最小跳跃次数
// 如果无法到达 n-1，返回 -1
// 注：对于复杂变形题，采用分组循环不易写错
// - [45. 跳跃游戏 II](https://leetcode.cn/problems/jump-game-ii/)
// - [1024. 视频拼接](https://leetcode.cn/problems/video-stitching/) 1746
// - [1326. 灌溉花园的最少水龙头数目](https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/) 1885
// 【图解】https://leetcode.cn/problems/minimum-number-of-taps-to-open-to-water-a-garden/solution/yi-zhang-tu-miao-dong-pythonjavacgo-by-e-wqry/
// https://codeforces.com/problemset/problem/1630/C 2200 变形 
// https://codeforces.com/problemset/problem/1066/B 1500 这题其实不算，但如果每个加热器的 r 不同，就是跳跃游戏 II 了
func minJumpNumbers(a []int) (ans int) {
	curR := 0 // 已建造的桥的右端点
	nxtR := 0 // 下一座桥的右端点的最大值
	// 这里没有遍历到 n-1，因为它已经是终点了
	for i, d := range a[:len(a)-1] {
		r := i + d
		nxtR = max(nxtR, r)
		if i == curR { // 到达已建造的桥的右端点
			if i == nxtR { // 无论怎么造桥，都无法从 i 到 i+1
				return -1
			}
			curR = nxtR // 建造下一座桥
			ans++
		}
	}
	return
}

// 区间合并
// 合并 a 中所有重叠的闭区间（哪怕只有一个端点重叠，也算重叠）
// 注意 [1,1] 和 [2,2] 不能合并成 [1,2]
// 注：这种做法在变形题中容易写错，更加稳定的做法是差分数组
// - [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
// - [55. 跳跃游戏](https://leetcode.cn/problems/jump-game/)
// - [763. 划分字母区间](https://leetcode.cn/problems/partition-labels/) 1443
// - [3169. 无需开会的工作日](https://leetcode.cn/problems/count-days-without-meetings/) ~1500
// - [2580. 统计将重叠区间合并成组的方案数](https://leetcode.cn/problems/count-ways-to-group-overlapping-ranges/) 1632
// - [2963. 统计好分割方案的数目](https://leetcode.cn/problems/count-the-number-of-good-partitions/) 1985
// - [2584. 分割数组使乘积互质](https://leetcode.cn/problems/split-the-array-to-make-coprime-products/) 2159
// - [2655. 寻找最大长度的未覆盖区间](https://leetcode.cn/problems/find-maximal-uncovered-ranges/)（会员题）
// 另见 common.go 中的「区间贪心」
// https://codeforces.com/problemset/problem/1626/C 1700
// - 倒序合并代码 https://codeforces.com/contest/1626/submission/211306494
// https://codeforces.com/problemset/problem/1859/D 1800
// https://codeforces.com/problemset/problem/1260/D 1900
func mergeIntervals(a [][]int) [][]int {
	slices.SortFunc(a, func(a, b []int) int { return a[0] - b[0] }) // 按区间左端点排序
	merged := [][]int{}
	l0 := a[0][0]
	maxR := a[0][1]
	for _, p := range a[1:] { // 从第二个区间开始
		l, r := p[0], p[1]
		// 如果要合并 [1,1] 和 [2,2]，下面改成 if l-1 > maxR
		if l > maxR { // 发现一个新区间
			merged = append(merged, []int{l0, maxR}) // 先把旧的加入答案
			l0 = l                                   // 记录新区间左端点
		}
		maxR = max(maxR, r)
	}
	merged = append(merged, []int{l0, maxR}) // 最后发现的新区间加入答案

	{
		// 包含 x 的闭区间
		var x int
		i := sort.Search(len(merged), func(i int) bool { return merged[i][1] >= x })
		if i < len(merged) && merged[i][0] <= x {
			// ans[i]...
		}
	}

	return merged
}
