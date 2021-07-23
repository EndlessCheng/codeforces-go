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
todo NOI 一轮复习 II：字符串 https://www.luogu.com.cn/blog/ix-35/noi-yi-lun-fu-xi-ii-zi-fu-chuan
金策 字符串算法选讲 https://www.bilibili.com/video/BV11K4y1p7a5 https://www.bilibili.com/video/BV19541177KU
    PDF 见 misc

TIPS: 若处理原串比较困难，不妨考虑下反转后的串 https://codeforces.com/contest/873/problem/F

斐波那契字符串：s(1) = "a", s(2) = "b", s(n) = s(n-1) + s(n-2), n>=3

https://en.wikipedia.org/wiki/Bitap_algorithm shift-or / shift-and / Baeza-Yates–Gonnet algorithm
*/

func stringCollection() {
	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a >= b {
			return a
		}
		return b
	}

	// 注：如果 s 是常量的话，由于其在编译期分配到只读段，对应的地址是无法写入的
	unsafeGetBytes := func(s string) []byte {
		return *(*[]byte)(unsafe.Pointer(&s))
	}

	unsafeToString := func(b []byte) string {
		return *(*string)(unsafe.Pointer(&b))
	}

	// 返回 t 在 s 中的所有位置（允许重叠）
	indexAll := func(s, t []byte) []int {
		pos := suffixarray.New(s).Lookup(t, -1)
		sort.Ints(pos)
		return pos
	}

	// 字符串哈希
	// https://oi-wiki.org/string/hash/
	// 利用 set 可以求出固定长度的不同子串个数
	// 模板题 https://www.luogu.com.cn/problem/P3370
	// LC187 找出所有重复出现的长为 10 的子串 https://leetcode-cn.com/problems/repeated-dna-sequences/
	// LC1044 最长重复子串（二分哈希）https://leetcode-cn.com/problems/longest-duplicate-substring/
	// LC1554 只有一个不同字符的字符串 https://leetcode-cn.com/problems/strings-differ-by-one-character/
	var powP []uint64
	initPowP := func(maxLen int) {
		const prime uint64 = 1e8 + 7
		powP = make([]uint64, maxLen)
		powP[0] = 1
		for i := 1; i < maxLen; i++ {
			powP[i] = powP[i-1] * prime
		}
	}
	calcHash := func(s []byte) (val uint64) {
		for i, c := range s {
			val += uint64(c) * powP[i]
		}
		return
	}

	// KMP
	// match[i] 为 s[:i+1] 的真前缀和真后缀的最长的匹配长度
	// 特别地，match[n-1] 为 s 的真前缀和真后缀的最长的匹配长度
	// 我在知乎上对 KMP 的讲解 https://www.zhihu.com/question/21923021/answer/37475572
	// https://oi-wiki.org/string/kmp/ todo 统计每个前缀的出现次数
	// TODO https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/prefix-function.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KMP.java.html
	// 模板题 https://loj.ac/p/103 https://www.luogu.com.cn/problem/P3375
	//       LC1392 https://leetcode-cn.com/problems/longest-happy-prefix/submissions/
	// https://codeforces.com/problemset/problem/432/D
	// https://codeforces.com/problemset/problem/471/D
	// 与 LCS 结合 https://codeforces.com/problemset/problem/346/B
	// 与 DP 结合 https://codeforces.com/problemset/problem/1163/D
	// 与计数 DP 结合 https://codeforces.com/problemset/problem/494/B
	// https://codeforces.com/problemset/problem/1003/F
	// http://acm.hdu.edu.cn/showproblem.php?pid=2087
	calcMaxMatchLengths := func(s []byte) []int {
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
	kmpSearch := func(text, pattern []byte) (pos []int) {
		match := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, v := range text {
			for c > 0 && pattern[c] != v {
				c = match[c-1]
			}
			if pattern[c] == v {
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
	// http://poj.org/problem?id=2406 https://www.luogu.com.cn/problem/UVA455
	calcMinPeriod := func(s []byte) int {
		n := len(s)
		match := calcMaxMatchLengths(s)
		if val := match[n-1]; val > 0 && n%(n-val) == 0 {
			return n / (n - val)
		}
		return 1 // 无小于 n 的循环节
	}

	// Z-function（扩展 KMP）      exkmp
	// z[i] = LCP(s, s[i:])   串与串后缀的最长公共前缀
	// 参考 Competitive Programmer’s Handbook Ch.26
	// https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/z-function.html
	// https://www.geeksforgeeks.org/z-algorithm-linear-time-pattern-searching-algorithm/
	// 模板题 https://codeforces.com/edu/course/2/lesson/3/3/practice/contest/272263/problem/A
	//       https://www.luogu.com.cn/problem/P5410
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
	calcZ := func(s []byte) []int {
		n := len(s)
		z := make([]int, n)
		for i, l, r := 1, 0, 0; i < n; i++ {
			z[i] = max(0, min(z[i-l], r-i+1))
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				l, r = i, i+z[i]
				z[i]++
			}
		}
		z[0] = n
		return z
	}
	zSearch := func(text, pattern []byte) (pos []int) {
		s := append(append(pattern, '#'), text...)
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
	// 模板题 https://www.luogu.com.cn/problem/P1368 http://poj.org/problem?id=1509
	smallestRepresentation := func(s []byte) []byte {
		n := len(s)
		s = append(s, s...)
		i := 0
		for j := 1; j < n; {
			k := 0
			for ; k < n && s[i+k] == s[j+k]; k++ {
			}
			if k >= n {
				break
			}
			if s[i+k] < s[j+k] {
				// j 到 j+k 都不会是最小串的开头位置
				j += k + 1
			} else {
				// i 到 i+k 都不会是最小串的开头位置
				i, j = j, max(j, i+k)+1
			}
		}
		return s[i : i+n]
	}

	// 最长回文子串 Manacher
	// https://blog.csdn.net/synapse7/article/details/18908413
	// https://www.bilibili.com/video/BV1ft4y117a4
	// https://codeforces.com/blog/entry/12143
	// http://manacher-viz.s3-website-us-east-1.amazonaws.com
	// https://oi-wiki.org/string/manacher/#manacher
	// https://cp-algorithms.com/string/manacher.html
	// 模板题 https://www.luogu.com.cn/problem/P3805
	//       LC5 https://leetcode-cn.com/problems/longest-palindromic-substring/
	// https://codeforces.com/contest/1326/problem/D2
	// todo 类似思想 https://codeforces.com/contest/359/problem/D
	var maxLen, left []int
	manacher := func(origin []byte) int {
		min := func(a, b int) int {
			if a < b {
				return a
			}
			return b
		}
		n := len(origin)
		m := 2*n + 2
		s := make([]byte, 1, m+1)
		s[0] = '^'
		for _, c := range origin {
			s = append(s, '#', c)
		}
		s = append(s, '#', '$')
		maxLen = make([]int, m+1) // 以处理后的字符 s[i] 为中心的最长回文子串的半长度（包括 s[i]）
		ans, mid, r := 0, 0, 0
		for i := 2; i < m; i++ {
			mx := 1
			if i < r {
				// 取 min 的原因：记点 i 关于 mid 的对称点为 i'=2*mid-i，
				// 若以 i' 为中心的回文串范围超过了以 mid 为中心的回文串的范围
				//（此时有 i + len[2*mid-i] >= r，这里 len 是包括中心的半长度）
				// 则 len[i] 应取 r - i (总不能超过边界吧)
				mx = min(maxLen[2*mid-i], r-i)
			}
			for ; s[i-mx] == s[i+mx]; mx++ {
			}
			if i+mx > r {
				mid, r = i, i+mx
			}
			if mx > ans {
				ans = mx
			}
			maxLen[i] = mx
		}

		// EXTRA: 计算以每个位置为起点的最长回文子串位置
		left = make([]int, m+1)
		for i := 2; i < m; i++ {
			if left[i-maxLen[i]+1] < i+1 {
				left[i-maxLen[i]+1] = i + 1
			}
		}
		for i := 1; i <= m; i++ {
			if left[i] < left[i-1] {
				left[i] = left[i-1]
			}
		}

		// todo 以每个位置为终点的...

		return ans - 1
	}
	// 判断 [l,r] 是否为回文串  0<=l<=r<n
	isP := func(l, r int) bool { return maxLen[l+r+2]-1 >= r-l+1 }
	// odd=true: 以下标 x 为中心的最长奇回文子串长度
	// odd=false: 以下标 x,x+1 中间为中心的最长偶回文子串长度
	midP := func(x int, odd bool) int {
		if odd {
			return maxLen[2*x+2] - 1
		}
		return maxLen[2*x+3] - 1
	}
	// EXTRA: 从下标 x 开始的最长回文子串长度
	leftP := func(x int) int { return left[2*x+2] - 2*x - 2 }

	/* 后缀数组
	SA-IS 与 DC3 的效率对比 https://riteme.site/blog/2016-6-19/sais.html#5
	NOTE: Go1.13 开始使用 SA-IS 算法

	讲解+例题+套题 https://oi-wiki.org/string/sa/
	todo 题目推荐 https://www.luogu.com.cn/blog/luckyblock/post-bi-ji-hou-zhui-shuo-zu
	CF 上的课程 https://codeforces.com/edu/course/2
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=string+suffix+structures

	题目总结：（部分参考《后缀数组——处理字符串的有力工具》，PDF 在 misc 文件夹下）
	单个字符串
		模板题 https://www.luogu.com.cn/problem/P3809
		可重叠最长重复子串 LC1044 https://leetcode-cn.com/problems/longest-duplicate-substring/ LC1062 https://leetcode-cn.com/problems/longest-repeating-substring/
			相当于求 max(height)，实现见下面的 longestDupSubstring
		不可重叠最长重复子串 http://poj.org/problem?id=1743
			可参考《算法与实现》p.223 以及 https://oi-wiki.org/string/sa/#_14
			重要技巧：按照 height 分组，每组中根据 sa 来处理组内后缀的位置
		可重叠的至少出现 k 次的最长重复子串 https://www.luogu.com.cn/problem/P2852 http://poj.org/problem?id=3261
			二分答案，对 height 分组，判定组内元素个数不小于 k
		不同子串个数 https://www.luogu.com.cn/problem/P2408 https://atcoder.jp/contests/practice2/tasks/practice2_i https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/A LC1698 https://leetcode-cn.com/problems/number-of-distinct-substrings-in-a-string/
			枚举每个后缀，计算前缀总数，再减掉重复，即 height[i]
			所以个数为 n*(n+1)/2-sum{height[i]} https://oi-wiki.org/string/sa/#_13
		不同子串长度之和 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/H
			思路同上，即 n*(n+1)*(n+2)/6-sum{height[i]*(height[i]+1)/2}
		带限制的不同子串个数
			https://codeforces.com/problemset/problem/271/D
			这题可以枚举每个后缀，跳过 height[i] 个字符，然后在前缀和上二分
		重复次数最多的连续重复子串 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/F http://poj.org/problem?id=3693 (数据弱)
			核心思想是枚举长度然后计算 LCP(i,i+l)，然后看是否还能再重复一次，具体代码见 main/edu/2/suffixarray/step5/f/main.go
		子串统计类题目
			用单调栈统计矩形面积 + 用单调栈跳过已经统计的
			https://codeforces.com/problemset/problem/123/D (注：这是挑战上推荐的题目)
			https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D 本质上就是 CF123D
			https://codeforces.com/problemset/problem/802/I 稍作改动
			todo https://www.luogu.com.cn/problem/P2178
			 https://www.luogu.com.cn/problem/P3804
		从字符串首尾取字符最小化字典序 https://oi-wiki.org/string/sa/#_10
			todo
		第 k 小子串 https://www.luogu.com.cn/problem/P3975
			todo
	两个字符串
		最长公共子串 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/B http://poj.org/problem?id=2774 LC718 https://leetcode-cn.com/problems/maximum-length-of-repeated-subarray/
			用 '#' 拼接两字符串，遍历 height[1:] 若 sa[i]<len(s1) != (sa[i-1]<len(s1)) 则更新 maxLen
		长度不小于 k 的公共子串的个数 http://poj.org/problem?id=3415
			单调栈
		最短公共唯一子串 https://codeforces.com/contest/427/problem/D
			唯一性可以用 height[i] 与前后相邻值的大小来判定
		公共回文子串 http://judge.u-aizu.ac.jp/onlinejudge/description.jsp?id=2292
			todo
		todo http://poj.org/problem?id=3729
	多个字符串
	    多串最长公共子串 https://loj.ac/p/171 LC周赛248D https://leetcode-cn.com/problems/longest-common-subpath/ http://poj.org/problem?id=3450
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数等于字符串个数
		不小于 k 个字符串中的最长子串 http://poj.org/problem?id=3294
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数不小于 k
		在每个字符串中至少出现两次且不重叠的最长子串 https://www.luogu.com.cn/problem/SP220
			拼接，二分答案，对 height 分组，判定组内元素在每个字符串中至少出现两次且 sa 的最大最小之差不小于二分值（用于判定是否重叠）
		出现或反转后出现在每个字符串中的最长子串 http://poj.org/problem?id=1226
			拼接反转后的串 s[i]+="#"+reverse(s)，拼接所有串，二分答案，对 height 分组，判定组内元素在每个字符串或其反转串中出现
	todo http://poj.org/problem?id=3581
	*/
	suffixArray := func(s []byte) {
		n := len(s)
		// sa[i] 表示后缀字典序中的第 i 个字符串在 s 中的位置
		//      后缀 s[sa[0]:] 字典序最小，后缀 s[sa[n-1]:] 字典序最大
		//sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").UnsafeAddr()))
		sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New(s)).Elem().FieldByName("sa").Field(0).UnsafeAddr()))

		// 后缀 s[i:] 位于后缀字典序中的第 rank[i] 个
		//     rank[0] 即 s 在后缀字典序中的排名，rank[n-1] 即 s[n-1:] 在字典序中的排名
		rank := make([]int, n)
		for i := range rank {
			rank[sa[i]] = i
		}

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
			return _q(ri+1, rj+1)
		}

		// EXTRA: 比较两个子串 s[l1,r1) 和 s[l2,r2)
		// https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/C
		lessSub := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l2, r2-l2
			if l := lcp(l1, l2); l >= len1 || l >= len2 { // 一个是另一个的前缀
				return len1 < len2
			}
			return rank[l1] < rank[l2] // 或者 s[l1+l] < s[l2+l]
		}

		// 返回值含义同 strings.Compare
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

		// https://www.acwing.com/problem/content/140/
		equalSub := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l2, r2-l2
			return len1 == len2 && len1 == lcp(l1, l2)
		}

		// EXTRA: 可重叠最长重复子串
		// https://leetcode-cn.com/problems/longest-duplicate-substring/ https://leetcode-cn.com/problems/longest-repeating-substring/
		longestDupSubstring := func() []byte {
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
		// LC1408/周赛184A https://leetcode-cn.com/problems/string-matching-in-an-array/ 「小题大做」
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
		tmp := make([]int, n+1)
		for i := 0; i <= n; i++ {
			sa[i] = i
			rank[i] = -1
			if i < n {
				rank[i] = a[i]
			}
		}
		for k = 1; k <= n; k *= 2 {
			sort.Slice(sa, func(i, j int) bool { return compareSA(sa[i], sa[j]) })
			tmp[sa[0]] = 0
			for i := 1; i <= n; i++ {
				tmp[sa[i]] = tmp[sa[i-1]]
				if compareSA(sa[i-1], sa[i]) {
					tmp[sa[i]]++
				}
			}
			copy(rank, tmp)
		}
		sa = sa[1:]
		return sa
	}

	_ = []interface{}{
		unsafeGetBytes, unsafeToString,
		indexAll,
		initPowP, calcHash,
		kmpSearch, calcMinPeriod,
		zSearch,
		smallestRepresentation,
		manacher, isP, midP, leftP,
		suffixArray, suffixArrayInt, suffixArrayInt2,
	}
}
