package copypasta

import (
	. "fmt"
	"math/bits"
)

/*
标准库 "math/bits" 包含了位运算常用的函数，如二进制中 1 的个数、二进制表示的长度等
注意：bits.Len(0) 返回的是 0 而不是 1
     bits.Len(x) 相当于 int(Log2(x)+eps)+1
     或者说 2^(Len(x)-1) <= x < 2^Len(x)

TIPS: & 和 | 在区间求和上具有单调性；^ 的区间求和见 strings.go 中的 trie.maxXor
** 代码和题目见下面的 bitOpTrick 和 bitOpTrickCnt

常用等式（若改变了计算的顺序，注意优先级！）
a|b = (a^b) + (a&b)    a^b = (a|b) - (a&b)
a+b = (a|b) + (a&b)
    = (a&b)*2 + (a^b)
    = (a|b)*2 - (a^b)
相关题目
https://codeforces.com/problemset/problem/1325/D
https://atcoder.jp/contests/abc050/tasks/arc066_b

结合律：(a&b)^(a&c) = a&(b^c)    其他符号类似
相关题目 https://leetcode-cn.com/contest/weekly-contest-237/problems/find-xor-sum-of-all-pairs-bitwise-and/

运算符优先级 https://golang.org/ref/spec#Operators
Precedence    Operator
    5         *  /  %  <<  >>  &  &^
    4         +  -  |  ^
    3         ==  !=  <  <=  >  >=
    2         &&
    1         ||

一些子集的枚举算法见 search.go
S∪{i}: S|1<<i
S\{i}: S&^(1<<i)
构造 2^n-1，即 n 个 1 的另一种方法: ^(-1<<n)
检测是否只有一个 1：x&(x-1) == 0

https://oeis.org/A060142 每一段连续 0 的长度均为偶数的数：如 100110000100
Ordered set S defined by these rules: 0 is in S and if x is in S then 2x+1 and 4x are in S
0, 1, 3, 4, 7, 9, 12, 15, 16, 19, 25, 28, 31, 33, 36, 39, 48, 51, 57, 60, 63, 64, 67, 73, 76, 79, 97, 100
https://oeis.org/A086747 Baum-Sweet sequence
相关题目：蒙德里安的梦想 https://www.acwing.com/problem/content/293/

https://oeis.org/A048004 最长连续 1 为 k 的长为 n 的二进制串的个数
相关题目：https://codeforces.com/problemset/problem/1027/E

https://oeis.org/A047778 Concatenation of first n numbers in binary, converted to base 10
相关题目：https://leetcode-cn.com/contest/weekly-contest-218/problems/concatenation-of-consecutive-binary-numbers/
钱珀瑙恩数 Champernowne constant https://en.wikipedia.org/wiki/Champernowne_constant

https://oeis.org/A072339
Any number n can be written (in two ways, one with m even and one with m odd) in the form n = 2^k_1 - 2^k_2 + 2^k_3 - ... + 2^k_m
where the signs alternate and k_1 > k_2 > k_3 > ... >k_m >= 0; sequence gives minimal value of m
https://codeforces.com/problemset/problem/1617/E

异或和相关
https://oeis.org/A003987 异或矩阵
https://oeis.org/A003815 异或和 i  a(0)=0, a(4n+1)=1, a(4n+2)=4n+3, a(4n+3)=0, a(4n+4)=4n+4
    相关题目 https://codeforces.com/problemset/problem/1493/E
            https://codeforces.com/problemset/problem/460/D
https://oeis.org/A145768 异或和 i*i
https://oeis.org/A126084 异或和 质数
https://oeis.org/A018252 异或和 合数?
https://oeis.org/A072594 异或和 质因数分解 是积性函数 a(p^k)=p*(k&1)
	https://oeis.org/A072595 满足 A072594(n)=0 的数
https://oeis.org/A178910 异或和 因子
	https://oeis.org/A178911 满足 A178910(n)=n 的数 Perfex number

https://oeis.org/A038712 a(n) = n^(n-1) = 1, 3, 1, 7, 1, 3, 1, 15, 1, ...
https://oeis.org/A080277 A038712 的前缀和  =>  a(n) = n + 2*a(n/2)

二进制长度
https://oeis.org/A029837 Binary order of n: log_2(n) rounded up to next integer
https://oeis.org/A001855 A029837 的前缀和
https://oeis.org/A070939 a(0)=1, a(n)=bits.Len(n)
https://oeis.org/A083652 A070939 的前缀和

OnesCount 相当于二进制的 digsum
https://oeis.org/A000120 wt(n) = OnesCount(n)
https://oeis.org/A000788 前缀和 a(0) = 0, a(2n) = a(n)+a(n-1)+n, a(2n+1) = 2a(n)+n+1
https://oeis.org/A121853 前缀积 https://www.luogu.com.cn/problem/P4317
https://oeis.org/A092391 n+OnesCount(n)
	https://oeis.org/A010061 二进制自我数/哥伦比亚数（A092391 的补集）
https://oeis.org/A011371 n-OnesCount(n) Also highest power of 2 dividing n!
							a(n) = floor(n/2) + a(floor(n/2))
                         这同时是前 n 个数的质因子分解的 2 的幂次之和
https://oeis.org/A027868 Number of trailing zeros in n!; highest power of 5 dividing n!
                            a(n) = (n-A053824(n))/4, 其中 A053824(n) = Sum of digits of (n written in base 5)
推广至任意数：n! 的质因子分解中，p 的幂次为 (n-digsum_p(n))/(p-1)，其中 digsum_p(n) 表示 n 的 p 进制的数位和
https://oeis.org/A245788 n*OnesCount(n)
https://oeis.org/A049445 OnesCount(n)|n
	-  n/OnesCount(n)
https://oeis.org/A199238 n%OnesCount(n)
https://oeis.org/A010062 a(0)=1, a(n+1)=a(n)+OnesCount(a(n))
	https://oeis.org/A096303 从 n 出发不断执行 n+=OnesCount(n)，直到 n 在 A010062 中，所需要的迭代次数
	Number of iterations of n -> n + (number of 1's in binary representation of n) needed for the trajectory of n to join the trajectory of A010062
		https://oeis.org/A229743 Positions of records
		https://oeis.org/A229744 Values of records
	相关题目 https://www.luogu.com.cn/problem/P5891 https://class.luogu.com.cn/classroom/lgr66
https://oeis.org/A180094 Number of steps to reach 0 or 1, starting with n and applying the map k -> (number of 1's in binary expansion of k) repeatedly

https://oeis.org/A023416 Number of 0's in binary expansion of n
							a(n) = a(n/2) + 1 - n&1
https://oeis.org/A059015 A023416 的前缀和

十进制 digsum
一点点数学 https://codeforces.com/problemset/problem/817/C
https://oeis.org/A007953 digsum(n)
https://oeis.org/A062028 n+digsum(n)    质数 https://oeis.org/A047791    合数 https://oeis.org/A107743
	https://oeis.org/A003052 自我数/哥伦比亚数 Self number / Colombian number
	https://en.wikipedia.org/wiki/Self_number
	1, 3, 5, 7, 9, 20, 31, 42, 53, 64, 75, 86, 97, 108, ...
		https://oeis.org/A006378 自我质数 Self primes
https://oeis.org/A066568 n-digsum(n)
https://oeis.org/A057147 n*digsum(n)
https://oeis.org/A005349 digsum(n)|n   Niven (or Harshad) numbers
	https://oeis.org/A065877 digsum(n)∤n   Non-Niven (or non-Harshad) numbers
	https://oeis.org/A001101 Moran numbers: n such that (n / digsum(n)) is prime
https://oeis.org/A016052 a(1)=3, a(n+1)=a(n)+digsum(a(n))
https://oeis.org/A051885 Smallest number whose digsum = n
							int64(n%9+1) * int64(math.Pow10(n/9)) - 1
							相关题目 https://codeforces.com/contest/1373/problem/E
https://oeis.org/A077196 Smallest possible sum of the digits of a multiple of n https://oeis.org/A077194 https://oeis.org/A077195
							相关题目（0-1 最短路）https://atcoder.jp/contests/arc084/tasks/arc084_b
https://oeis.org/A118137 digsum(n)+digsum(n+1)
https://oeis.org/A003132 Sum of squares of digits of n
	https://oeis.org/A003621 Number of iterations until n reaches 1 or 4 under x goes to sum of squares of digits map
https://oeis.org/A055012 Sum of cubes of digits of n
https://oeis.org/A055013 Sum of 4th powers of digits of n
https://oeis.org/A055014 Sum of 5th powers of digits of n
https://oeis.org/A055015 Sum of 6th powers of digits of n
	相关题目 https://www.luogu.com.cn/problem/P1660
https://oeis.org/A031286 Additive persistence: number of summations of digits needed to obtain a single digit (the additive digital root)
https://oeis.org/A031346 Multiplicative persistence: number of iterations of "multiply digits" needed to reach a number < 10

https://oeis.org/A014837 Sum of all the digits of n in every base from 2 to n-1
https://oeis.org/A043306 Sum of all the digits of n in every base from 2 to n

回文数
https://oeis.org/A002113 十进制回文数
	https://oeis.org/A043269 digsum(A002113(n))
	https://oeis.org/A070199 Number of palindromes of length <= n
https://oeis.org/A002779 回文平方数
	https://oeis.org/A002778 Numbers whose square is a palindrome
https://oeis.org/A002781 回文立方数
	https://oeis.org/A002780 Numbers whose cube is a palindrome
https://oeis.org/A002385 回文素数
	https://en.wikipedia.org/wiki/Palindromic_prime
https://oeis.org/A006567 反素数 emirp (primes whose reversal is a different prime)
	https://en.wikipedia.org/wiki/Emirp
https://oeis.org/A003459 绝对素数/可交换素数 Absolute primes (or permutable primes): every permutation of the digits is a prime
	https://en.wikipedia.org/wiki/Permutable_prime
https://oeis.org/A007500 Primes whose reversal in base 10 is also prime
https://oeis.org/A006995 二进制回文数
https://oeis.org/A007632 既是二进制回文数又是十进制回文数

https://oeis.org/A090994 Number of meaningful differential operations of the n-th order on the space R^9
a(k+5) = a(k+4) + 4*a(k+3) - 3*a(k+2) - 3*a(k+1) + a(k)
相关题目 LC1215/双周赛10C https://leetcode-cn.com/contest/biweekly-contest-10/problems/stepping-numbers/

套路题 https://codeforces.com/problemset/problem/1415/D
按位归纳 https://codeforces.com/problemset/problem/925/C
*/

// Bitset
// 参考 C++ 的标准库源码 https://gcc.gnu.org/onlinedocs/libstdc++/libstdc++-html-USERS-3.4/bitset-source.html
// 若要求方法内不修改 b 而是返回一个修改后的拷贝，可以在方法开头加上 b = append(Bitset(nil), b...) 并返回 b
// 应用：https://codeforces.com/problemset/problem/33/D（也可以用 LCA）
const _w = bits.UintSize

func NewBitset(n int) Bitset { return make(Bitset, n/_w+1) } // (n+_w-1)/_w

type Bitset []uint

func (b Bitset) Has(p int) bool { return b[p/_w]&(1<<(p%_w)) != 0 } // get
func (b Bitset) Flip(p int)     { b[p/_w] ^= 1 << (p % _w) }
func (b Bitset) Set(p int)      { b[p/_w] |= 1 << (p % _w) }  // 置 1
func (b Bitset) Reset(p int)    { b[p/_w] &^= 1 << (p % _w) } // 置 0

// 左移 k 位
// 应用 https://leetcode-cn.com/problems/minimize-the-difference-between-target-and-chosen-elements/submissions/
func (b Bitset) Lsh(k int) {
	if k == 0 {
		return
	}
	shift, offset := k/_w, k%_w
	if shift >= len(b) {
		for i := range b {
			b[i] = 0
		}
		return
	}
	if offset == 0 {
		// Fast path
		copy(b[shift:], b)
	} else {
		for i := len(b) - 1; i > shift; i-- {
			b[i] = b[i-shift]<<offset | b[i-shift-1]>>(_w-offset)
		}
		b[shift] = b[0] << offset
	}
	for i := 0; i < shift; i++ {
		b[i] = 0
	}
}

// 右移 k 位
func (b Bitset) Rsh(k int) {
	if k == 0 {
		return
	}
	shift, offset := k/_w, k%_w
	if shift >= len(b) {
		for i := range b {
			b[i] = 0
		}
		return
	}
	lim := len(b) - 1 - shift
	if offset == 0 {
		// Fast path
		copy(b, b[shift:])
	} else {
		for i := 0; i < lim; i++ {
			b[i] = b[i+shift]>>offset | b[i+shift+1]<<(_w-offset)
		}
		// 注意：若前后调用 lsh 和 rsh，需要注意超出 n 的范围的 1 对结果的影响
		b[lim] = b[len(b)-1] >> offset
	}
	for i := lim + 1; i < len(b); i++ {
		b[i] = 0
	}
}

// 返回 1 的个数
func (b Bitset) OnesCount() (c int) {
	for _, v := range b {
		c += bits.OnesCount(v)
	}
	return
}

// 遍历所有 1 的位置
// 如果对范围有要求，可在 f 中 return p < n
func (b Bitset) Foreach(f func(p int) (Break bool)) {
	for i, v := range b {
		for ; v > 0; v &= v - 1 {
			j := i*_w | bits.TrailingZeros(v)
			if f(j) {
				return
			}
		}
	}
}

// 返回第一个 0 的下标，不存在时会返回一个不小于 n 的位置
func (b Bitset) Index0() int {
	for i, v := range b {
		if ^v != 0 {
			return i*_w | bits.TrailingZeros(^v)
		}
	}
	return len(b) * _w
}

// 返回第一个 1 的下标，不存在时会返回一个不小于 n 的位置（同 C++ 中的 _Find_first）
func (b Bitset) Index1() int {
	for i, v := range b {
		if v != 0 {
			return i*_w | bits.TrailingZeros(v)
		}
	}
	return len(b) * _w
}

// 返回下标严格大于 p 的第一个 1 的下标，不存在时会返回一个不小于 n 的位置（同 C++ 中的 _Find_next）
func (b Bitset) Next1(p int) int {
	p++ // make bound inclusive
	if i := p / _w; i < len(b) {
		v := b[i] & (^uint(0) << (p % _w)) // mask off bits below bound
		if v != 0 {
			return i*_w | bits.TrailingZeros(v)
		}
		for i++; i < len(b); i++ {
			if b[i] != 0 {
				return i*_w | bits.TrailingZeros(b[i])
			}
		}
	}
	return len(b) * _w
}

// 判断 [l,r] 范围内的数是否全为 0
// https://codeforces.com/contest/1107/problem/D（标准做法是二维前缀和）
func (b Bitset) All0(l, r int) bool {
	i := l / _w
	if i == r/_w {
		return b[i]>>(l%_w)&(1<<(r-l+1)-1) == 0
	}
	if b[i]>>(l%_w) != 0 {
		return false
	}
	for i++; i < r/_w; i++ {
		if b[i] != 0 {
			return false
		}
	}
	return b[r/_w]&(1<<(r%_w+1)-1) == 0
}

// 判断 [l,r] 范围内的数是否全为 1
func (b Bitset) All1(l, r int) bool {
	i := l / _w
	if i == r/_w {
		mask := uint(1)<<(r-l+1) - 1
		return b[i]>>(l%_w)&mask == mask
	}
	if ^(b[i] | (1<<(l%_w) - 1)) != 0 {
		return false
	}
	for i++; i < r/_w; i++ {
		if ^b[i] != 0 {
			return false
		}
	}
	mask := uint(1)<<(r%_w+1) - 1
	return b[r/_w]&mask == mask
}

// 下面几个方法均需保证长度相同
func (b Bitset) Equals(c Bitset) bool {
	for i, v := range b {
		if v != c[i] {
			return false
		}
	}
	return true
}

func (b Bitset) HasSubset(c Bitset) bool {
	for i, v := range b {
		if v|c[i] != v {
			return false
		}
	}
	return true
}

// 将 c 的元素合并进 b
func (b Bitset) Merge(c Bitset) {
	for i, v := range c {
		b[i] |= v
	}
}

// 注：有关子集枚举的位运算技巧，见 search.go
func _() {
	// 利用 -v = ^v+1
	lowbit := func(v int64) int64 { return v & -v }

	// 1,2,4,8,...
	isPow2 := func(v int64) bool { return v > 0 && v&(v-1) == 0 }

	// 是否有两个相邻的 1    有 https://oeis.org/A004780 没有 https://oeis.org/A003714
	hasAdjacentOnes := func(v uint) bool { return v>>1&v > 0 }

	// 是否有两个相邻的 0（不考虑前导零）    有 https://oeis.org/A004753 没有 http://oeis.org/A003754
	hasAdjacentZeros := func(v uint) bool {
		v |= v >> 1 // 若没有相邻的 0，则 v 会变成全 1 的数
		return v&(v+1) > 0
	}

	bits31 := func(v int) []byte {
		bits := make([]byte, 31)
		for i := range bits {
			bits[i] = byte(v >> (30 - i) & 1)
		}
		return bits
	}
	_bits31 := func(v int) string { return Sprintf("%031b", v) }
	_bits32 := func(v uint) string { return Sprintf("%032b", v) }

	// 返回最小的非负 x，其满足 n^x >= m
	// https://codeforces.com/problemset/problem/1554/C
	leastXor := func(n, m int) (res int) {
		for i := 29; i >= 0; i-- { // 29 for 1e9
			bn, bm := n>>i&1, m>>i&1
			if bn == 1 && bm == 0 { // 后面都填 0
				break
			}
			if bn == 0 && bm == 1 { // 必须填 1
				res |= 1 << i
			}
			// bn = bm 的情况填 0
		}
		return
	}

	// 对于数组 a 的所有区间，返回 op(区间元素) 的全部运算结果    logTrick
	// 利用操作的单调性求解
	// 复杂度 O(f * n * logU)，f 为 op(x,y) 的时间复杂度，n 为 a 的长度，U 为 a 中元素最大值
	// |: LC898/周赛100C https://leetcode-cn.com/contest/weekly-contest-100/problems/bitwise-ors-of-subarrays/
	// &: LC1521/周赛198D https://leetcode-cn.com/contest/weekly-contest-198/problems/find-a-value-of-a-mysterious-function-closest-to-target/
	// GCD: https://codeforces.com/edu/course/2/lesson/9/2/practice/contest/307093/problem/G
	//      https://codeforces.com/problemset/problem/475/D (见下面的 bitOpTrickCnt)
	//      https://codeforces.com/problemset/problem/1632/D (见下面的 bitOpTrickCnt)
	bitOpTrick := func(a []int, op func(x, y int) int) map[int]bool {
		ans := map[int]bool{} // 统计 op(一段区间) 的不同结果
		set := []int{}
		for _, v := range a {
			for i, w := range set {
				set[i] = op(w, v)
			}
			set = append(set, v)
			// 去重
			k := 0
			for _, w := range set[1:] {
				if set[k] != w {
					k++
					set[k] = w
				}
			}
			set = set[:k+1]
			for _, w := range set {
				// do w...
				ans[w] = true
			}
		}
		return ans
	}

	// 进阶：对于数组 a 的所有区间，返回 op(区间元素) 的全部运算结果及其出现次数
	// https://codeforces.com/problemset/problem/475/D
	// https://codeforces.com/problemset/problem/1632/D
	// 与单调栈结合 https://codeforces.com/problemset/problem/875/D
	// CERC13，紫书例题 10-29，UVa 1642 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4517
	bitOpTrickCnt := func(a []int, op func(x, y int) int) map[int]int64 {
		cnt := map[int]int64{}
		type result struct{ v, l, r int } // [l,r)
		set := []result{}
		for i, v := range a {
			for j, p := range set {
				set[j].v = op(p.v, v)
			}
			set = append(set, result{v, i, i + 1})
			// 去重
			k := 0
			for _, q := range set[1:] {
				if set[k].v != q.v {
					k++
					set[k] = q
				} else {
					set[k].r = q.r
				}
			}
			set = set[:k+1]
			// 此时我们将区间 [0,i] 划分成了 len(set) 个左闭右开区间
			// 对 ∀p∈set，∀j∈[p.l,p.r)，op(区间[j,i]) 的计算结果均为 p.v
			for _, p := range set {
				// do p...     [l,r)
				cnt[p.v] += int64(p.r - p.l)
			}
		}
		return cnt
	}

	//（接上）考虑乘法
	// 问题：给一数组 a，元素均为正整数，求区间和等于区间积的区间个数
	// 我们来考虑对每个区间右端点，有多少个合法的区间左端点
	// 核心思路是，对于每个满足题目要求的区间，其区间积不会超过 sum(a)
	// 由于乘积至少要乘 2 才会变化，所以对于一个固定的区间右端点，不同的区间积至多有 O(log(sum(a))) 个
	// 同时由于元素均为正数，所以对一个固定的区间右端点，区间左端点也至多有 O(log(sum(a))) 个
	// 据此我们只需要在加入一个新的数后，去重并去掉区间积超过 sum(a) 的区间，就可以暴力做出此题
	// 注：根据以上推导过程，我们还可以得出总的答案个数至多为 O(nlog(sum(a)))
	// https://www.dotcpp.com/oj/problem2622.html
	countSumEqMul := func(a []int) (ans int) {
		tot := 0
		for _, v := range a {
			tot += v
		}
		// 每个前缀和互不相同
		posS := map[int]int{0: 0} // int64
		sum := 0
		type result struct{ v, l, r int }
		muls := []result{}
		for i, v := range a {
			sum += v
			for j := range muls {
				muls[j].v *= v
			}
			muls = append(muls, result{v, i, i + 1})
			// 去重
			k := 0
			for _, q := range muls[1:] {
				if muls[k].v != q.v {
					k++
					muls[k] = q
				} else {
					muls[k].r = q.r
				}
			}
			muls = muls[:k+1]
			// 去掉超过 tot 的，从而保证 muls 中至多有 O(log(tot)) 个元素
			for muls[0].v > tot {
				muls = muls[1:]
			}
			// 此时我们将区间 [0,i] 划分成了 len(muls) 个（左闭右开）区间，对 ∀j∈[muls[k].l,muls[k].r)，[j,i] 的区间积均为 muls[k].v
			for _, p := range muls {
				// 判断左端点前缀和对应下标是否在范围内
				if pos, has := posS[sum-p.v]; has && p.l <= pos && pos < p.r {
					ans++
				}
			}
			posS[sum] = i + 1
		}
		return
	}

	// 找三个不同的在 [l,r] 范围内的数，其异或和为 0
	// 考虑尽可能地小化最大减最小的值，构造 (x, y, z) = (b*2-1, b*3-1, b*3), b=2^k
	// 相关题目 https://codeforces.com/problemset/problem/460/D
	zeroXorSum3 := func(l, r int64) []int64 {
		for b := int64(1); b*3 <= r; b <<= 1 {
			if x, y, z := b*2-1, b*3-1, b*3; l <= x && z <= r {
				return []int64{x, y, z}
			}
		}
		return nil
	}

	// 在 [low,high] 区间内找两个数字 A B，使其异或值最大且不超过 limit
	// 返回值保证 A <= B
	// 复杂度 O(log(high))
	maxXorWithLimit := func(low, high, limit int, min func(int, int) int) (int, int) {
		n := bits.Len(uint(high ^ low))
		maxXor := 1<<n - 1
		mid := high&^maxXor | 1<<(n-1)
		if limit >= maxXor { // 无约束，相关题目 https://codeforces.com/problemset/problem/276/D
			return mid - 1, mid
		}
		if limit >= 1<<(n-1) { // A 和 B 能否在第 n-1 位不同的情况下，构造出一个满足要求的解？
			a, b := mid&(mid-1), mid
			for i := n - 2; i >= 0; i-- {
				bt := 1 << i
				if limit&bt > 0 { // a 取 1，b 取 0 总是优于 a 取 0，b 取 1
					a |= bt
				} else if high&(bt<<1-1) > ^low&(bt<<1-1) { // high 侧大，都取 1
					if high&bt == 0 { // b 没法取 1
						goto next
					}
					a |= bt
					b |= bt
				} else {            // low 侧大，都取 0
					if low&bt > 0 { // a 没法取 0
						goto next
					}
				}
				if (a^low)&bt > 0 { // a 不受 low 的约束
					a |= limit & (bt - 1)
					break
				}
				if (b^high)&bt > 0 { // b 不受 high 的约束
					a |= bt - 1
					b |= ^limit & (bt - 1)
					break
				}
			}
			return a, b
		}
		// A 和 B 在第 n-1 位上必须相同
	next:
		f := func(high int) (int, int) {
			n := bits.Len(uint(high ^ mid))
			maxXor := min(1<<n-1, limit)
			// 只有当 maxXor 为 0 时，返回值才必须相等
			if maxXor == 0 {
				return mid, mid
			}
			// maxXor 的最高位置于 B，其余置于 A
			mb := 1 << (bits.Len(uint(maxXor)) - 1)
			return mid | maxXor&^mb, mid | mb
		}
		if high-mid > mid-1-low { // 选区间长的一侧
			return f(high)
		}
		a, b := f(2*mid - 1 - low) // 对称到 high
		return 2*mid - 1 - b, 2*mid - 1 - a
	}

	_ = []interface{}{lowbit, isPow2, hasAdjacentOnes, hasAdjacentZeros, bits31, _bits31, _bits32, leastXor, bitOpTrick, bitOpTrickCnt, countSumEqMul, zeroXorSum3, maxXorWithLimit}
}

// https://halfrost.com/go_s2_de_bruijn/

// LC137 https://leetcode-cn.com/problems/single-number-ii/
// 除了某个元素只出现一次以外，其余每个元素均出现了三次。返回只出现了一次的元素
// 		定义两个集合 ones 和 twos，初始为空
// 		第一次出现就放在 ones 中
//		第二次出现就在 ones 中删除并放在 twos
//		第三次出现就从 twos 中删除
//		这样最终 ones 中就留下了最后的结果
func singleNumber(a []int) int {
	ones, twos := 0, 0
	for _, v := range a {
		ones = (ones ^ v) &^ twos
		twos = (twos ^ v) &^ ones
	}
	return ones
}
