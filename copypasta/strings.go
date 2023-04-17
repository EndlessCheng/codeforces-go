package copypasta

import (
	"index/suffixarray"
	"math/bits"
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

/*
字符串问题的特殊性：
不同子串之间会共享一些局部信息，巧妙地利用这些局部信息可以设计出更加高效的算法。

todo NOI 一轮复习 II：字符串 https://www.luogu.com.cn/blog/ix-35/noi-yi-lun-fu-xi-ii-zi-fu-chuan
金策 字符串算法选讲 https://www.bilibili.com/video/BV11K4y1p7a5 https://www.bilibili.com/video/BV19541177KU
    PDF 见 https://github.com/EndlessCheng/cp-pdf

TIPS: 若处理原串比较困难，不妨考虑下反转后的串 https://codeforces.com/contest/873/problem/F

斐波那契字符串：s(1) = "a", s(2) = "b", s(n) = s(n-1) + s(n-2), n>=3

https://en.wikipedia.org/wiki/Bitap_algorithm shift-or / shift-and / Baeza-Yates–Gonnet algorithm
*/

func _(min, max func(int, int) int) {
	// 注：如果 s 是常量的话，由于其在编译期分配到只读段，对应的地址是无法写入的
	unsafeToBytes := func(s string) []byte { return *(*[]byte)(unsafe.Pointer(&s)) }
	unsafeToString := func(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

	// 返回 t 在 s 中的所有位置（允许重叠）
	indexAll := func(s, t []byte) []int {
		pos := suffixarray.New(s).Lookup(t, -1)
		sort.Ints(pos)
		return pos
	}

	// 字符串哈希 rolling hash, Rabin–Karp algorithm
	// https://en.wikipedia.org/wiki/Hash_function
	// https://en.wikipedia.org/wiki/Rolling_hash
	// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
	// 线性同余方法（LCG）https://en.wikipedia.org/wiki/Linear_congruential_generator
	// https://oi-wiki.org/string/hash/
	// 利用 set 可以求出固定长度的不同子串个数
	// todo 浅谈字符串 hash 的应用 https://www.luogu.com.cn/blog/Flying2018/qian-tan-zi-fu-chuan-hash
	//      anti-hash: 最好不要自然溢出 https://codeforces.com/blog/entry/4898
	//      On the mathematics behind rolling hashes and anti-hash tests https://codeforces.com/blog/entry/60442
	//      hash killer https://loj.ac/p/6758
	//  Kapun's algorithm https://codeforces.com/blog/entry/99973
	// 题目推荐 https://cp-algorithms.com/string/string-hashing.html#toc-tgt-7
	// 模板题 https://www.luogu.com.cn/problem/P3370
	// LC187 找出所有重复出现的长为 10 的子串 https://leetcode-cn.com/problems/repeated-dna-sequences/
	// LC1044 最长重复子串（二分哈希）https://leetcode-cn.com/problems/longest-duplicate-substring/
	// LC1554 只有一个不同字符的字符串 https://leetcode-cn.com/problems/strings-differ-by-one-character/
	// 倒序哈希 https://leetcode-cn.com/problems/find-substring-with-given-hash-value/solution/dao-xu-hua-dong-chuang-kou-o1-kong-jian-xpgkp/
	hash := func(s string) {
		// 注意：由于哈希很容易被卡，能用其它方法实现尽量用其它方法
		const prime uint64 = 1e8 + 7
		powP := make([]uint64, len(s)+1) // powP[i] = prime^i
		powP[0] = 1
		preHash := make([]uint64, len(s)+1) // preHash[i] = hash(s[:i]) 前缀哈希
		for i, b := range s {
			powP[i+1] = powP[i] * prime
			preHash[i+1] = preHash[i]*prime + uint64(b) // 本质是秦九韶算法
		}

		// 计算子串 s[l:r] 的哈希   0<=l<=r<=len(s)
		// 空串的哈希值为 0
		subHash := func(l, r int) uint64 { return preHash[r] - preHash[l]*powP[r-l] }
		_ = subHash
	}

	// todo 二维字符串哈希
	// https://www.luogu.com.cn/problem/solution/UVA11019
	// UVa 11019 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=22&page=show_problem&problem=1960

	// KMP (Knuth–Morris–Pratt algorithm)
	// match[i] 为 s[:i+1] 的真前缀和真后缀的最长的匹配长度
	// 特别地，match[n-1] 为 s 的真前缀和真后缀的最长的匹配长度
	// 我在知乎上对 KMP 的讲解 https://www.zhihu.com/question/21923021/answer/37475572
	// https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
	// https://oi-wiki.org/string/kmp/ todo 统计每个前缀的出现次数
	// TODO https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/prefix-function.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KMP.java.html
	//
	// 模板题 https://loj.ac/p/103 https://www.luogu.com.cn/problem/P3375
	//       LC28 https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
	//       LC1392 https://leetcode-cn.com/problems/longest-happy-prefix/
	// 最长回文前缀 LC214 https://leetcode.cn/problems/shortest-palindrome/
	// LC1316 https://leetcode.cn/problems/distinct-echo-substrings/
	// https://codeforces.com/problemset/problem/432/D
	// https://codeforces.com/problemset/problem/471/D
	// 与 LCS 结合 https://codeforces.com/problemset/problem/346/B
	// 与 DP 结合 https://codeforces.com/problemset/problem/1163/D
	// 与计数 DP 结合 https://codeforces.com/problemset/problem/494/B
	// https://codeforces.com/problemset/problem/1003/F
	// http://acm.hdu.edu.cn/showproblem.php?pid=2087
	// 最大匹配个数 https://codeforces.com/problemset/problem/615/C
	// 与贝尔数（集合划分）结合 https://codeforces.com/problemset/problem/954/I
	calcMaxMatchLengths := func(s string) []int {
		match := make([]int, len(s))
		for i, c := 1, 0; i < len(s); i++ {
			v := s[i]
			for c > 0 && s[c] != v {
				c = match[c-1]
			}
			if s[c] == v {
				c++
			}
			match[i] = c
		}
		return match
	}
	// search pattern from text, return all start positions
	kmpSearch := func(text, pattern string) (pos []int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != byte(v) {
				c = match[c-1]
			}
			if pattern[c] == byte(v) {
				c++
			}
			if c == lenP {
				pos = append(pos, i-lenP+1)
				c = match[c-1] // 不允许重叠时 c = 0
			}
		}
		return
	}
	// EXTRA: 最小循环节
	// 返回循环节以及循环次数
	// http://poj.org/problem?id=2406 https://www.luogu.com.cn/problem/UVA455
	// LC459 https://leetcode.cn/problems/repeated-substring-pattern/
	calcMinPeriod := func(s string) (string, int) {
		n := len(s)
		match := calcMaxMatchLengths(s)
		if m := match[n-1]; m > 0 && n%(n-m) == 0 {
			return s[:n-m], n / (n - m)
		}
		return s, 1 // 无小于 n 的循环节
	}

	// todo 失配树（border 树）
	//  https://www.luogu.com.cn/problem/P5829

	// Z-function（扩展 KMP，Z-array）      exkmp
	// z[i] = LCP(s, s[i:])   串与串后缀的最长公共前缀
	// 参考 Competitive Programmer’s Handbook Ch.26
	// https://oi-wiki.org/string/z-func/
	// 可视化 https://personal.utdallas.edu/~besp/demo/John2010/z-algorithm.htm
	// https://cp-algorithms.com/string/z-function.html
	// https://www.geeksforgeeks.org/z-algorithm-linear-time-pattern-searching-algorithm/
	//
	// 模板题 https://codeforces.com/edu/course/2/lesson/3/3/practice/contest/272263/problem/A
	//       https://www.luogu.com.cn/problem/P5410
	// 结论 https://codeforces.com/problemset/problem/535/D
	// 最小循环节（允许末尾截断）https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/A
	// s 和 t 是否本质相同，shift 多少次 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/B
	//		即 strings.Index(s+s, t)
	// 每个前缀的出现次数 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/C
	//		用 z[i] 来进行区间更新操作，实现时用一个差分数组即可
	//		注：字符串倒过来就是每个后缀的出现次数
	// 既是前缀又是后缀的子串个数 https://codeforces.com/problemset/problem/432/D
	//		解法之一是 a[z[i]]++ 然后求 a 的后缀和
	//		解法之二是对 z 排序二分，见我的代码
	//		其他解法有 KMP+DP 或 SA，见 https://www.luogu.com.cn/problem/solution/CF432D
	// 最长回文前缀 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/D
	//		构造 s+reverse(s)
	// 判断是否存在 i 使得 s[i:]+reverse(s[:i]) == t https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/E
	//		构造 t+s
	// 最短的包含 s 和 t 的字符串 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/F
	// 		构造 s+t 和 t+s
	// 判断一个字符串 t 是否为 prefix+reverse(s)+suffix，其中 prefix+suffix=s https://atcoder.jp/contests/abc284/tasks/abc284_f
	//      构造 t+reverse(t) 和 reverse(t)+t
	calcZ := func(s string) []int {
		n := len(s)
		z := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			if i <= r {
				z[i] = min(z[i-l], r-i+1)
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
		}
		z[0] = n
		return z
	}
	zSearch := func(text, pattern string) (pos []int) {
		s := pattern + "#" + text
		z := calcZ(s)
		for i, l := range z[len(pattern)+1:] {
			if l == len(pattern) {
				pos = append(pos, i)
			}
		}
		return
	}

	// 最小表示法 - 求串的循环同构串中字典序最小的串
	// 找到位置 i，从这个位置输出即得到字典序最小的串
	// https://oi-wiki.org/string/minimal-string/
	// 其他方法 https://codeforces.com/blog/entry/90035
	// 模板题 https://www.luogu.com.cn/problem/P1368 http://poj.org/problem?id=1509 https://codeforces.com/gym/103585/problem/K
	// https://codeforces.com/problemset/problem/496/B
	smallestRepresentation := func(s string) string {
		n := len(s)
		s += s
		// 如果要返回一个和原串不同的字符串，初始化 i=1, j=2
		i := 0
		for j := 1; j < n; {
			k := 0
			for k < n && s[i+k] == s[j+k] {
				k++
			}
			if k >= n {
				break
			}
			if s[i+k] < s[j+k] { // 改成 > 则返回字典序最大的
				// j 到 j+k 都不会是最小串的开头位置
				j += k + 1
			} else {
				// i 到 i+k 都不会是最小串的开头位置
				i, j = j, max(j, i+k)+1
			}
		}
		return s[i : i+n]
	}

	// 子序列自动机
	// LC727 https://leetcode-cn.com/problems/minimum-window-subsequence/
	// LC792 https://leetcode-cn.com/problems/number-of-matching-subsequences/
	// LC2014 https://leetcode-cn.com/problems/longest-subsequence-repeated-k-times/
	subsequenceAutomaton := func(s string) {
		// build nxt
		// nxt[i][j] 表示在 i 右侧的字符 j 的最近位置
		pos := [26]int{}
		for i := range pos {
			pos[i] = len(s)
		}
		nxt := make([][26]int, len(s))
		for i := len(s) - 1; i >= 0; i-- {
			nxt[i] = pos
			pos[s[i]-'a'] = i
		}

		// 返回是 s 的子序列的最长的 t 的前缀的长度
		match := func(t string) int {
			if t == "" || s == "" {
				return 0
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1
			}
			for ; j < len(t); j++ {
				i = nxt[i][t[j]-'a']
				if i == len(s) {
					break
				}
			}
			return j
		}
		_ = match
	}

	// 最长回文子串 Manacher（马拉车算法）
	// https://www.bilibili.com/video/BV1AX4y1F79W
	// https://www.bilibili.com/video/BV1ft4y117a4
	// https://oi-wiki.org/string/manacher/
	// https://cp-algorithms.com/string/manacher.html
	// https://codeforces.com/blog/entry/12143
	// https://leetcode-cn.com/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
	// https://www.zhihu.com/question/37289584
	// https://blog.csdn.net/synapse7/article/details/18908413
	// http://manacher-viz.s3-website-us-east-1.amazonaws.com
	//
	// 模板题 https://www.luogu.com.cn/problem/P3805
	//       LC5 https://leetcode-cn.com/problems/longest-palindromic-substring/
	// https://codeforces.com/problemset/problem/1326/D2
	// https://codeforces.com/problemset/problem/7/D https://codeforces.com/problemset/problem/835/D
	// todo 相交的回文子串对数 https://codeforces.com/problemset/problem/17/E
	//  https://codeforces.com/problemset/problem/1081/H
	//  https://www.luogu.com.cn/blog/user25308/proof-cf1081h
	//  LC1745 分割成三个非空回文子字符串 https://leetcode-cn.com/problems/palindrome-partitioning-iv/
	// LC2472 不重叠回文子字符串（长度至少为 k）的最大数目 https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/
	// - 只需要考虑长度为 k or k+1 的
	manacher := func(s string) {
		// 将 s 改造为 t，这样就不需要分 len(s) 的奇偶来讨论了，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
		// s 和 t 的下标转换关系：
		// (si+1)*2 = ti
		// ti/2-1 = si
		// ti 为偶数对应奇回文串（从 2 开始）
		// ti 为奇数对应偶回文串（从 3 开始）
		t := append(make([]byte, 0, len(s)*2+3), '^')
		for _, c := range s {
			t = append(t, '#', byte(c))
		}
		t = append(t, '#', '$')

		// 定义一个奇回文串的半长度=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
		// halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的半长度
		// 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
		halfLen := make([]int, len(t)-2)
		halfLen[1] = 1
		// r 表示当前右边界下标最大的回文子串的右边界下标+1
		// mid 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
		for i, mid, r := 2, 1, 0; i < len(halfLen); i++ { // 循环的起止位置对应着原串的首尾字符
			hl := 1
			if i < r {
				// 记 i 关于 mid 的对称位置 i'=mid*2-i
				// 若以 i' 为中心的最长回文子串范围超出了以 mid 为中心的回文串的范围（即 i+halfLen[i'] >= r）
				// 则 halfLen[i] 应先初始化为已知的半长度 r-i，然后再继续暴力匹配
				// 否则 halfLen[i] 与 halfLen[i'] 相等
				hl = min(halfLen[mid*2-i], r-i)
			}
			// 暴力扩展
			// 算法的复杂度取决于这部分执行的次数
			// 由于扩展之后 r 必然会更新（右移），且扩展的的次数就是 r 右移的次数
			// 因此算法的复杂度和 t 的长度成正比
			for t[i-hl] == t[i+hl] {
				hl++
			}
			if i+hl > r {
				mid, r = i, i+hl
			}
			halfLen[i] = hl
		}

		// t 中回文子串的长度为 hl*2-1
		// 由于其中 # 的数量总是比字符的数量多 1
		// 因此其在 s 中对应的回文子串的长度为 hl-1
		// 这一结论可用在下面的各个代码中

		// 判断 s[l..r] 是否为回文串  0<=l<=r<len(s)
		// 根据下标转换关系得到其在 t 中对应的回文中心下标为 l+r+2
		// https://codeforces.com/problemset/problem/7/D https://codeforces.com/problemset/problem/835/D
		isP := func(l, r int) bool { return halfLen[l+r+2]-1 >= r-l+1 } // 或者 halfLen[l+r+2] > r-l+1

		// 计算最长回文子串的长度，以及所有最长回文子串的首字母在 s 中的下标
		maxPL, starts := 0, []int{}
		for i := 2; i < len(halfLen); i++ {
			if hl := halfLen[i]; hl-1 > maxPL {
				// 由于 t 中回文子串的首尾字母一定是 #，再根据下标转换关系，可以得到其在 s 中对应的回文子串的区间为：
				// [(i-hl)/2, (i+hl)/2-2]
				maxPL, starts = hl-1, []int{(i - hl) / 2}
			} else if hl-1 == maxPL {
				starts = append(starts, (i-hl)/2)
			}
		}

		// odd=true:  以 s[x] 为回文中心的最长奇回文子串长度
		// odd=false: 以 s[x] 和 s[x+1] 为回文中心的最长偶回文子串长度
		// 根据下标转换关系得到其在 t 中对应的回文中心下标
		midPL := func(x int, odd bool) int {
			if odd {
				return halfLen[x*2+2] - 1
			}
			return halfLen[x*2+3] - 1
		}

		// EXTRA: 计算两个数组，其中
		// startPL[i] 表示以 s[i] 为首字母的最长回文子串的长度
		// endPL[i]   表示以 s[i] 为尾字母的最长回文子串的长度
		// [国家集训队]最长双回文串 https://www.luogu.com.cn/problem/P4555
		// LC1960 两个回文子字符串长度的最大乘积 https://leetcode-cn.com/problems/maximum-product-of-the-length-of-two-palindromic-substrings/
		// LC214 https://leetcode-cn.com/problems/shortest-palindrome/
		startPL := make([]int, len(s))
		endPL := make([]int, len(s))
		for i := 2; i < len(halfLen); i++ {
			hl := halfLen[i]
			left, right := (i-hl)/2, (i+hl)/2-2 // 见上面计算 maxPL 的注释
			startPL[left] = max(startPL[left], hl-1)
			endPL[right] = max(endPL[right], hl-1)
		}
		for i := 1; i < len(startPL); i++ {
			startPL[i] = max(startPL[i], startPL[i-1]-2) // startPL[i] 还可能是一个更长的回文串缩短后的结果，两者取最大值，同时也方便传递
		}
		for i := len(endPL) - 2; i >= 0; i-- {
			endPL[i] = max(endPL[i], endPL[i+1]-2) // 同上
		}

		// EXTRA: 计算回文子串个数
		// 易证其为 ∑(halfLen[i]/2)
		// LC647 https://leetcode-cn.com/problems/palindromic-substrings/
		totP := int64(0)
		for _, hl := range halfLen {
			totP += int64(hl / 2)
		}

		_ = []interface{}{isP, midPL}
	}

	/* 后缀数组
	SA-IS 与 DC3 的效率对比 https://riteme.site/blog/2016-6-19/sais.html#5
	注：Go1.13 开始使用 SA-IS 算法

	可视化 https://visualgo.net/zh/suffixarray

	常用分隔符 #(35) $(36) _(95) |(124)

	讲解+例题+套题 https://oi-wiki.org/string/sa/
	todo 题目推荐 https://www.luogu.com.cn/blog/luckyblock/post-bi-ji-hou-zhui-shuo-zu
	CF 上的课程 https://codeforces.com/edu/course/2
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=string+suffix+structures

	题目总结：（部分参考《后缀数组——处理字符串的有力工具》，PDF 见 https://github.com/EndlessCheng/cp-pdf）
	单个字符串
		模板题 https://www.luogu.com.cn/problem/P3809 https://loj.ac/p/111
		可重叠最长重复子串 LC1044 https://leetcode-cn.com/problems/longest-duplicate-substring/ LC1062 https://leetcode-cn.com/problems/longest-repeating-substring/
			相当于求 max(height)，实现见下面的 longestDupSubstring
		不可重叠最长重复子串 https://atcoder.jp/contests/abc141/tasks/abc141_e http://poj.org/problem?id=1743
			可参考《算法与实现》p.223 以及 https://oi-wiki.org/string/sa/#是否有某字符串在文本串中至少不重叠地出现了两次
			重要技巧：按照 height 分组，每组中根据 sa 来处理组内后缀的位置
		可重叠的至少出现 k 次的最长重复子串 https://www.luogu.com.cn/problem/P2852 http://poj.org/problem?id=3261
			二分答案，对 height 分组，判定组内元素个数不小于 k
		本质不同子串个数 https://www.luogu.com.cn/problem/P2408 https://www.luogu.com.cn/problem/SP694 https://atcoder.jp/contests/practice2/tasks/practice2_i https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/A LC1698 https://leetcode-cn.com/problems/number-of-distinct-substrings-in-a-string/
			枚举每个后缀，计算前缀总数，再减掉重复，即 height[i]
			所以个数为 n*(n+1)/2-sum{height[i]} https://oi-wiki.org/string/sa/#_13
			相似思路 LC2261 含最多 K 个可整除元素的子数组 https://leetcode-cn.com/problems/k-divisible-elements-subarrays/solution/by-freeyourmind-2m6j/
		不同子串长度之和 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/H
			思路同上，即 n*(n+1)*(n+2)/6-sum{height[i]*(height[i]+1)/2}
		带限制的不同子串个数
			https://codeforces.com/problemset/problem/271/D
			这题可以枚举每个后缀，跳过 height[i] 个字符，然后在前缀和上二分
		重复次数最多的连续重复子串 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/F http://poj.org/problem?id=3693 (数据弱)
			核心思想是枚举长度然后计算 LCP(i,i+l)，然后看是否还能再重复一次，具体代码见 main/edu/2/suffixarray/step5/f/main.go
		重复两次的最长连续重复子串 LC1316 https://leetcode.cn/problems/distinct-echo-substrings/
			题解 https://leetcode.cn/problems/distinct-echo-substrings/solution/geng-kuai-de-onlog2n-jie-fa-hou-zhui-shu-8wby/
		子串统计类题目
			用单调栈统计矩形面积 + 用单调栈跳过已经统计的
			https://codeforces.com/problemset/problem/123/D (注：这是《挑战》上推荐的题目)
			https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D 本质上就是 CF123D
			https://codeforces.com/problemset/problem/802/I 稍作改动
			todo https://www.luogu.com.cn/problem/P2178
			 https://www.luogu.com.cn/problem/P3804
			 AHOI13 差异 https://www.luogu.com.cn/problem/P4248
			 - 任意两后缀的 LCP 之和
			 对所有 i，求出 ∑j=1..n LCP(i,j) https://atcoder.jp/contests/abc213/tasks/abc213_f
		从字符串首尾取字符最小化字典序 https://oi-wiki.org/string/sa/#_10
			todo
		第 k 小子串 https://www.luogu.com.cn/problem/P3975 https://codeforces.com/problemset/problem/128/B
			todo
	两个字符串
		最长公共子串 SPOJ LCS https://www.luogu.com.cn/problem/SP1811 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/B http://poj.org/problem?id=2774 LC718 https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
			用 '#' 拼接两字符串，遍历 height[1:] 若 sa[i]<len(s1) != (sa[i-1]<len(s1)) 则更新 maxLen
		长度不小于 k 的公共子串的个数 http://poj.org/problem?id=3415
			单调栈
		最短公共唯一子串 https://codeforces.com/contest/427/problem/D
			唯一性可以用 height[i] 与前后相邻值的大小来判定
		公共回文子串 http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=2292
			todo
		所有循环串的比较计数 https://atcoder.jp/contests/abc272/tasks/abc272_f https://atcoder.jp/contests/abc272/submissions/35520643
			构造 s+s+"#"+t+t+"|"
		todo http://poj.org/problem?id=3729
	多个字符串
		多串最长公共子串 SPOJ LCS2 https://www.luogu.com.cn/problem/SP1812 https://loj.ac/p/171 LC1923 https://leetcode-cn.com/problems/longest-common-subpath/ http://poj.org/problem?id=3450
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数等于字符串个数
		不小于 k 个字符串中的最长子串 http://poj.org/problem?id=3294
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数不小于 k
		在每个字符串中至少出现两次且不重叠的最长子串 https://www.luogu.com.cn/problem/SP220
			拼接，二分答案，对 height 分组，判定组内元素在每个字符串中至少出现两次且 sa 的最大最小之差不小于二分值（用于判定是否重叠）
		出现或反转后出现在每个字符串中的最长子串 http://poj.org/problem?id=1226
			拼接反转后的串 s[i]+="#"+reverse(s)，拼接所有串，二分答案，对 height 分组，判定组内元素在每个字符串或其反转串中出现
		acSearch（https://www.luogu.com.cn/problem/P3796）的后缀数组做法
			拼接所有串（模式+文本，# 隔开），对每个模式 p 找其左右范围，满足该范围内 height[i] >= len(p)，这可以用 ST+二分或线段树二分求出，然后统计区间内的属于文本串的后缀
	todo 待整理 http://poj.org/problem?id=3581
	 https://www.luogu.com.cn/problem/P5546
	 https://www.luogu.com.cn/problem/P2048
	 https://www.luogu.com.cn/problem/P4248
	 https://www.luogu.com.cn/problem/P4341
	 https://www.luogu.com.cn/problem/P6095
	 https://www.luogu.com.cn/problem/P4070
	*/
	suffixArray := func(s string) {
		n := len(s)

		// 后缀数组 sa
		// sa[i] 表示后缀字典序中的第 i 个字符串在 s 中的位置
		// 特别地，后缀 s[sa[0]:] 字典序最小，后缀 s[sa[n-1]:] 字典序最大
		//sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").UnsafeAddr()))
		sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))

		{
			// 别的求法
			type _tp struct {
				_  []byte
				sa []int32
			}
			sa = (*_tp)(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
		}

		// 后缀名次数组 rank
		// 后缀 s[i:] 位于后缀字典序中的第 rank[i] 个
		// 特别地，rank[0] 即 s 在后缀字典序中的排名，rank[n-1] 即 s[n-1:] 在字典序中的排名
		rank := make([]int, n)
		for i, p := range sa {
			rank[p] = i
		}

		// 高度数组 height
		// height[0] = 0
		// height[i] = LCP(s[sa[i]:], s[sa[i-1]:])
		// 由于 height 数组的性质，可以和二分/单调栈/单调队列结合
		// 见 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
		// 	  https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/E
		//    https://codeforces.com/problemset/problem/873/F
		height := make([]int, n)
		h := 0
		for i, rk := range rank {
			if h > 0 {
				h--
			}
			if rk > 0 {
				for j := int(sa[rk-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
				}
			}
			height[rk] = h
		}

		// 任意两后缀的 LCP
		// 注：若允许离线可以用 Trie+Tarjan 做到线性
		const mx = 17 // 131072, 262144, 524288, 1048576
		st := make([][mx]int, n)
		for i, v := range height {
			st[i][0] = v
		}
		for j := 1; 1<<j <= n; j++ {
			for i := 0; i+1<<j <= n; i++ {
				st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
			}
		}
		_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
		lcp := func(i, j int) int {
			if i == j {
				return n - i
			}
			// 将 s[i:] 和 s[j:] 通过 rank 数组映射为 height 的下标
			ri, rj := rank[i], rank[j]
			if ri > rj {
				ri, rj = rj, ri
			}
			// ri+1 是因为 height 的定义是 sa[i] 和 sa[i-1]
			// rj+1 是因为 _q 是左闭右开
			return _q(ri+1, rj+1)
		}

		// EXTRA: 比较两个子串，返回 s[l1:r1] == s[l2:r2]，注意这里是左闭右开区间
		// https://www.acwing.com/problem/content/140/
		equalSub := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l1, r2-l2
			return len1 == len2 && lcp(l1, l2) >= len1
		}

		// EXTRA: 比较两个子串，返回 s[l1:r1] < s[l2:r2]，注意这里是左闭右开区间
		// https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/C
		lessSub := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l1, r2-l2
			if l := lcp(l1, l2); l >= len1 || l >= len2 { // 一个是另一个的前缀
				return len1 < len2
			}
			return rank[l1] < rank[l2] // 或者 s[l1+l] < s[l2+l]
		}

		// EXTRA: 比较两个子串，返回 strings.Compare(s[l1:r1], s[l2:r2])，注意这里是左闭右开区间
		// https://codeforces.com/problemset/problem/611/D
		// LC1977 https://leetcode-cn.com/problems/number-of-ways-to-separate-numbers/
		compareSub := func(l1, r1, l2, r2 int) int {
			len1, len2 := r1-l1, r2-l2
			l := lcp(l1, l2)
			if len1 == len2 && l >= len1 {
				return 0
			}
			if l >= len1 || l >= len2 { // 一个是另一个的前缀
				if len1 < len2 {
					return -1
				}
				return 1
			}
			if rank[l1] < rank[l2] { // 或者 s[l1+l] < s[l2+l]
				return -1
			}
			return 1
		}

		// EXTRA: 可重叠最长重复子串
		// LC1044 https://leetcode-cn.com/problems/longest-duplicate-substring/
		longestDupSubstring := func() string {
			maxP, maxH := 0, 0
			for i, h := range height {
				if h > maxH {
					maxP, maxH = i, h
				}
			}
			return s[sa[maxP] : int(sa[maxP])+maxH]
		}

		// EXTRA: 按后缀字典序求前缀和
		// vals[i] 表示 s[i] 的某个属性
		vals := make([]int, n)
		prefixSum := make([]int, n+1)
		for i, p := range sa {
			prefixSum[i+1] = prefixSum[i] + vals[p]
		}

		// EXTRA: 找出数组中的所有字符串，其是某个字符串的子串
		// 先拼接字符串，然后根据 height 判断前后是否有能匹配的
		// NOTE: 下面的代码展示了一种「标记 s[i] 属于原数组的哪个元素」的技巧: 在 i>0&&s[i]=='#' 时将 cnt++，其余的 s[i] 指向的 cnt 就是原数组的下标
		// LC1408 https://leetcode-cn.com/problems/string-matching-in-an-array/ 「小题大做」
		findAllSubstring := func(a []string) (ans []string) {
			s := "#" + strings.Join(a, "#")
			n := len(s)
			lens := make([]int, n) // lens[i] > 0 表示 s[i] 是原数组中的某个字符串的首字母，且 lens[i] 为该字符串长度
			cnt := 0
			for i := 1; i < n; i++ {
				if s[i-1] == '#' {
					lens[i] = len(a[cnt])
					cnt++
				}
			}
			// sa & height ...
			for i, p := range sa {
				if l := lens[p]; l > 0 {
					if height[i] >= l || i+1 < n && height[i+1] >= l {
						ans = append(ans, s[p:int(p)+l])
					}
				}
			}
			return
		}

		// debug
		for i, h := range height[:n] {
			suffix := string(s[sa[i]:])
			if h == 0 {
				println(" ", suffix)
			} else {
				println(h, suffix)
			}
		}

		_ = []interface{}{lessSub, compareSub, equalSub, longestDupSubstring, findAllSubstring}
	}

	// 若输入为 []int32，通过将每个元素拆成 4 个 byte，来满足调库条件
	// 若有负数，且需要满足有序性，可以整体减去 math.MinInt32 转成 uint32
	suffixArrayInt := func(a []int32) []int32 {
		n := len(a)
		_s := make([]byte, 0, n*4)
		for _, v := range a {
			_s = append(_s, byte(v>>24), byte(v>>16), byte(v>>8), byte(v))
		}
		_sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(_s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
		sa := make([]int32, 0, n)
		for _, p := range _sa {
			if p&3 == 0 { // 是 4 的倍数的 _sa[i] 就对应着数组 a 的 sa[i]
				sa = append(sa, p>>2)
			}
		}
		return sa
	}

	// 另一种写法，O(1) 得到 _s
	// 注意由于小端序的缘故，这里得到的 _s 和上面是不一样的，所以只有当题目与顺序无关时才可以使用
	suffixArrayInt = func(a []int32) []int32 {
		_sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))
		_sh.Len *= 4
		_sh.Cap *= 4
		_s := *(*[]byte)(unsafe.Pointer(_sh))
		_sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(_s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))
		sa := make([]int32, 0, len(a))
		for _, p := range _sa {
			if p&3 == 0 { // 是 4 的倍数的 _sa[i] 就对应着数组 a 的 sa[i]
				sa = append(sa, p>>2)
			}
		}
		return sa
	}

	// 注：这是《挑战》上的实现方案，复杂度 O(nlog^2(n))
	suffixArrayInt2 := func(a []int) []int {
		n := len(a)
		sa := make([]int, n+1)
		rank := make([]int, n+1)
		k := 0
		compareSA := func(i, j int) bool {
			if rank[i] != rank[j] {
				return rank[i] < rank[j]
			}
			ri, rj := -1, -1
			if i+k <= n {
				ri = rank[i+k]
			}
			if j+k <= n {
				rj = rank[j+k]
			}
			return ri < rj
		}
		for i := 0; i <= n; i++ {
			sa[i] = i
			rank[i] = -1
			if i < n {
				rank[i] = a[i]
			}
		}
		for k = 1; k <= n; k *= 2 {
			sort.Slice(sa, func(i, j int) bool { return compareSA(sa[i], sa[j]) })
			tmp := make([]int, n+1)
			tmp[sa[0]] = 0
			for i := 1; i <= n; i++ {
				tmp[sa[i]] = tmp[sa[i-1]]
				if compareSA(sa[i-1], sa[i]) {
					tmp[sa[i]]++
				}
			}
			rank = tmp
		}
		sa = sa[1:]
		return sa
	}

	_ = []interface{}{
		unsafeToBytes, unsafeToString,
		indexAll,
		hash,
		kmpSearch, calcMinPeriod,
		zSearch,
		smallestRepresentation,
		subsequenceAutomaton,
		manacher,
		suffixArray, suffixArrayInt, suffixArrayInt2,
	}
}

// AC 自动机见 trie.go
// 后缀自动机见 sam.go
