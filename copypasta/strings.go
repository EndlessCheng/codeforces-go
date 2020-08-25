package copypasta

import (
	"index/suffixarray"
	"math/bits"
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

// TIPS: 对于环形的字符串匹配，可以将文本串复制一倍后再匹配
// TIPS: 若处理原串比较困难，不妨考虑下反转后的串 https://codeforces.com/contest/873/problem/F

// 斐波那契字符串：s(1) = "a", s(2) = "b", s(n) = s(n-1) + s(n-2), n>=3

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

	// 返回 s 中 t 的所有位置
	indexAll := func(s, t []byte) []int {
		pos := suffixarray.New(s).Lookup(t, -1)
		sort.Ints(pos)
		return pos
	}

	// 字符串哈希
	// https://oi-wiki.org/string/hash/
	// 利用 set 可以求出固定长度的不同子串个数
	// 模板题 https://www.luogu.com.cn/problem/P3370
	// 最长重复子串（二分哈希）LC1044 https://leetcode-cn.com/problems/longest-duplicate-substring/
	// 题目推荐 https://cp-algorithms.com/string/string-hashing.html#toc-tgt-7
	// TODO 二维 hash
	// TODO anti-hash: 最好不要自然溢出 https://codeforces.com/blog/entry/4898 https://codeforces.com/blog/entry/60442
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

	// https://oi-wiki.org/string/kmp/ todo 统计每个前缀的出现次数
	// TODO https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/prefix-function.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KMP.java.html
	// 下面的代码来自我在知乎上的回答 https://www.zhihu.com/question/21923021/answer/37475572
	// 模板题 https://www.luogu.com.cn/problem/P3375
	// https://codeforces.com/problemset/problem/432/D
	// https://codeforces.com/problemset/problem/1003/F
	// http://acm.hdu.edu.cn/showproblem.php?pid=2087
	calcMaxMatchLengths := func(s []byte) []int {
		n := len(s)
		maxMatchLengths := make([]int, n)
		c := 0
		for i := 1; i < n; i++ {
			b := s[i]
			for c > 0 && s[c] != b {
				c = maxMatchLengths[c-1]
			}
			if s[c] == b {
				c++
			}
			maxMatchLengths[i] = c
		}
		return maxMatchLengths
	}
	// search pattern from text, return all start positions
	kmpSearch := func(text, pattern []byte) (pos []int) {
		maxMatchLengths := calcMaxMatchLengths(pattern)
		lenP := len(pattern)
		c := 0
		for i, b := range text {
			for c > 0 && pattern[c] != b {
				c = maxMatchLengths[c-1]
			}
			if pattern[c] == b {
				c++
			}
			if c == lenP {
				pos = append(pos, i-lenP+1)
				c = maxMatchLengths[c-1] // 不允许重叠时 c = 0
			}
		}
		return
	}
	// EXTRA: 最小循环节
	// http://poj.org/problem?id=2406
	calcMinPeriod := func(s []byte) int {
		n := len(s)
		maxMatchLengths := calcMaxMatchLengths(s)
		if val := maxMatchLengths[n-1]; val > 0 && n%(n-val) == 0 {
			return n / (n - val)
		}
		return 1 // 无小于 n 的循环节
	}

	// Z-function（扩展 KMP）
	// z[i] = LCP(s, s[i:])   串与串后缀的最长公共前缀
	// 参考 Competitive Programmer’s Handbook Ch.26
	// https://oi-wiki.org/string/z-func/
	// https://cp-algorithms.com/string/z-function.html
	// https://www.geeksforgeeks.org/z-algorithm-linear-time-pattern-searching-algorithm/
	// 模板题 https://codeforces.com/edu/course/2/lesson/3/3/practice/contest/272263/problem/A https://www.luogu.com.cn/problem/P5410
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
	//
	calcZArray := func(s []byte) []int {
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
		z := calcZArray(s)
		for i, l := range z[len(pattern)+1:] {
			if l == len(pattern) {
				pos = append(pos, i)
			}
		}
		return
	}
	// todo 反向：z[i] = LCS(s, s[:i])  串与串前缀的最长公共后缀

	// 最小表示法 - 求串的循环同构串中字典序最小的串
	// 找到位置 i，从这个位置输出即得到字典序最小的串
	// https://oi-wiki.org/string/minimal-string/
	// 模板题 https://www.luogu.com.cn/problem/P1368
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
		s := make([]byte, m+1)
		s[0] = '^'
		for i, c := range origin {
			s[2*i+1] = '#'
			s[2*i+2] = c
		}
		s[2*n+1] = '#'
		s[2*n+2] = '$'
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

	https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/SuffixArray.java.html
	https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/SuffixArrayX.java.html

	讲解+例题+套题 https://oi-wiki.org/string/sa/
	CF 上的课程 https://codeforces.com/edu/course/2
	题目推荐 https://cp-algorithms.com/string/suffix-array.html#toc-tgt-11
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=string+suffix+structures

	题目总结：（部分参考《后缀数组——处理字符串的有力工具》）
	单个字符串
		模板题 https://www.luogu.com.cn/problem/P3809
		可重叠最长重复子串 LC1044 https://leetcode-cn.com/problems/longest-duplicate-substring/
			相当于求 max(height)，实现见下面的 longestDupSubstring
		不可重叠最长重复子串 http://poj.org/problem?id=1743
			可参考《算法与实现》p.223 以及 https://oi-wiki.org/string/sa/#_14
			重要技巧：按照 height 分组，每组中根据 sa 来处理组内后缀的位置
		可重叠的至少出现 k 次的最长重复子串 https://www.luogu.com.cn/problem/P2852 http://poj.org/problem?id=3261
			二分答案，对 height 分组，判定组内元素个数不小于 k
		不同子串个数 https://www.luogu.com.cn/problem/P2408 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/A
			枚举每个后缀，计算前缀总数，再减掉重复，即 height[i]
			所以个数为 n*(n+1)/2-sum{height[i]} https://oi-wiki.org/string/sa/#_13
		不同子串长度之和 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/H
			思路同上，即 n*(n+1)*(n+2)/6-sum{height[i]*(height[i]+1)/2}
		重复次数最多的连续重复子串 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/F http://poj.org/problem?id=3693 (数据弱)
			核心思想是枚举长度然后计算 LCP(i,i+l)，然后看是否还能再重复一次，具体细节见 main/edu/...
		所有子串的所有公共前后缀个数 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
			单调栈
			思想类似的题目
				https://codeforces.com/problemset/problem/123/D 本质上和上面求的是同一个
				https://codeforces.com/problemset/problem/802/I 稍作改动
		从字符串首尾取字符最小化字典序 https://oi-wiki.org/string/sa/#_10
			todo
	两个字符串
		最长公共子串 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/B http://poj.org/problem?id=2774
			用 '#' 拼接两字符串，遍历 height[1:] 若 sa[i]<len(s1) != (sa[i-1]<len(s1)) 则更新 maxLen
		长度不小于 k 的公共子串的个数 http://poj.org/problem?id=3415
			单调栈
		最短公共唯一子串 https://codeforces.com/contest/427/problem/D
			唯一性可以用 height[i] 与前后相邻值的大小来判定
	多个字符串
		不小于 k 个字符串中的最长子串 http://poj.org/problem?id=3294
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数不小于 k
		在每个字符串中至少出现两次且不重叠的最长子串 https://www.luogu.com.cn/problem/SP220
			拼接，二分答案，对 height 分组，判定组内元素在每个字符串中至少出现两次且 sa 的最大最小之差不小于二分值（用于判定是否重叠）
		出现或反转后出现在每个字符串中的最长子串 http://poj.org/problem?id=1226
			拼接反转后的串 s[i]+="#"+reverse(s)，拼接所有串，二分答案，对 height 分组，判定组内元素在每个字符串或其反转串中出现
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

		// height[i] = LCP(s[sa[i]:], s[sa[i-1]:])
		// 由于 height 数组的性质，可以和二分/单调栈/单调队列结合
		// 见 https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/D
		// 	  https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/E
		//    https://codeforces.com/problemset/problem/873/F
		height := make([]int, n)
		h := 0
		for i, ri := range rank {
			if h > 0 {
				h--
			}
			if ri > 0 {
				for j := int(sa[ri-1]); i+h < n && j+h < n && s[i+h] == s[j+h]; h++ {
				}
			}
			height[ri] = h
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

		// EXTRA: 比较两个子串 s[l1:r1] s[l2:r2]
		// https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/C
		compareSub := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l2, r2-l2
			if l := lcp(l1, l2); l >= len1 || l >= len2 {
				return len1 < len2
			}
			return rank[l1] < rank[l2] // 或者 s[l1+l] < s[l2+l]
		}

		// EXTRA: 可重叠最长重复子串
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

		_ = []interface{}{compareSub, longestDupSubstring, findAllSubstring}
	}

	_ = []interface{}{
		unsafeGetBytes, unsafeToString,
		indexAll,
		initPowP, calcHash,
		kmpSearch, calcMinPeriod,
		zSearch,
		smallestRepresentation,
		manacher, isP, midP, leftP,
		suffixArray,
	}
}

// 前缀树/字典树/单词查找树
// 另类解读：如果将字符串长度视作定值的话，trie 树是一种 O(n) 排序，O(1) 查询的数据结构
//          这点上和哈希表很像，但是 trie 树可以在路径上保存信息，从而能做到一些哈希表做不到的前缀操作
// https://oi-wiki.org/string/trie/
// https://www.quora.com/q/threadsiiithyderabad/Tutorial-on-Trie-and-example-problems
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieST.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TrieSET.java.html
// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/TST.java.html
// 另见 strings_index_trie.go
// NOTE: 为保证连续性，分隔符可取 'Z'+1 或 'z'+1
// 模板题 LC208 https://leetcode-cn.com/problems/implement-trie-prefix-tree/
// 好题：前缀和后缀搜索 周赛62D/LC745 https://leetcode-cn.com/problems/prefix-and-suffix-search/
// 回文对（配合 Manacher 可以做到线性复杂度）LC336 https://leetcode-cn.com/problems/palindrome-pairs/
// LC 套题（推荐困难难度的题） https://leetcode-cn.com/tag/trie/
// todo https://codeforces.com/contest/455/problem/B
type trieNode struct {
	son    [26]*trieNode
	dupCnt int
	val    int // val 也可以是个 []int，此时 dupCnt == len(val)

	// AC 自动机: 当 o.son[i] 不能匹配 text 中的某个字符时，o.fail 即为下一个应该查找的结点
	fail *trieNode
}

func (o *trieNode) empty() bool {
	for _, son := range o.son {
		if son != nil {
			return false
		}
	}
	return true
}

type trie struct{ root *trieNode }

func newTrie() *trie {
	// init with a root (empty string)
	return &trie{&trieNode{}}
}

func (*trie) ord(c byte) byte { return c - 'a' }
func (*trie) chr(v byte) byte { return v + 'a' }

func (t *trie) put(s []byte, val int) *trieNode {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		if o.son[b] == nil {
			o.son[b] = &trieNode{}
		}
		o = o.son[b]
		//o.dupCnt++ // 经过节点 o 的字符串个数（EXTRA: 统计前缀个数）
		//o.val = val // 更新 s 的所有前缀的值
	}
	o.dupCnt++
	o.val = val
	return o
}

func (t *trie) find(s []byte) *trieNode {
	o := t.root
	for _, b := range s {
		o = o.son[t.ord(b)]
		if o == nil {
			return nil
		}
	}
	if o.dupCnt == 0 {
		return nil
	} // s 只是某个字符串的前缀
	return o
}

// s 必须在 trie 中存在
func (t *trie) del(s []byte) {
	fa := make([]*trieNode, len(s)+1)
	o := t.root
	for i, b := range s {
		fa[i] = o
		o = o.son[t.ord(b)]
		//o.dupCnt-- // 对应 put 的写法
	}
	o.dupCnt--
	if o.dupCnt == 0 {
		for i := len(s) - 1; i >= 0; i-- {
			f := fa[i]
			f.son[t.ord(s[i])] = nil
			if !f.empty() {
				break
			}
		}
	}
}

// 在 trie 中寻找字典序最小的以 p 为前缀的字符串
func (t *trie) minPrefix(p string) (s []byte, node *trieNode) {
	o := t.root
	for i := range p {
		o = o.son[t.ord(p[i])]
		if o == nil {
			return
		}
	}
	// trie 中存在字符串 s，使得 p 是 s 的前缀
	s = []byte(p)
	for o.dupCnt == 0 {
		for i, son := range o.son {
			if son != nil {
				s = append(s, t.chr(byte(i)))
				o = son
				break
			}
		}
	}
	return s, o
}

// rank 和 kth 分别求小于 s 的字符串个数和第 k 小
// 此时 o.val 保存子树字符串个数
func (t *trie) rank(s []byte) (k int) {
	o := t.root
	for _, b := range s {
		b = t.ord(b)
		for _, son := range o.son[:b] {
			if son != nil {
				k += son.val
			}
		}
		o = o.son[b]
		if o == nil {
			return
		}
	}
	//k += o.val // 等于 s 的也算上
	return
}

func (t *trie) kth(k int) (s []byte) {
	o := t.root
outer:
	for {
		for i, son := range o.son {
			if son != nil {
				if k >= son.val {
					k -= son.val
				} else {
					o = son
					s = append(s, t.chr(byte(i)))
					continue outer
				}
			}
		}
		return
	}
}

// 01-trie：val 与树上所有数中的最大异或值
// 也可以说这是一颗（所有叶节点深度都相同的）二叉树
// 参考《算法竞赛进阶指南》0x16
// 模板题：数组中两个数的最大异或值 LC421 https://leetcode-cn.com/problems/maximum-xor-of-two-numbers-in-an-array/
// 模板题：树上最长异或路径 https://www.luogu.com.cn/problem/P4551
// todo 好题：区间异或第 k 大 https://www.luogu.com.cn/problem/P5283
func (t *trie) maxXor(val int) (ans int) {
	bits := [31]byte{}
	for i := range bits {
		bits[i] = byte(val >> (30 - i) & 1)
	}

	o := t.root
	for i, b := range bits {
		if o.son[b^1] != nil {
			ans |= 1 << (30 - i)
			b ^= 1
		}
		o = o.son[b]
	}
	return
}

// 也可以用哈希表做，效率是一样的
func findMaximumXOR(a []int) (ans int) {
	n := len(a)
	for i := 30; i >= 0; i-- {
		ans <<= 1
		prefixes := make(map[int]bool, n)
		for _, v := range a {
			prefixes[v>>i] = true
		}
		tmp := ans + 1
		for p := range prefixes {
			if prefixes[tmp^p] {
				ans = tmp
				break
			}
		}
	}
	return
}

// EXTRA: AC 自动机 Aho–Corasick algorithm / Deterministic Finite Automaton (DFA)
// https://en.wikipedia.org/wiki/Aho%E2%80%93Corasick_algorithm
// https://en.wikipedia.org/wiki/Deterministic_finite_automaton
// 基础实现 https://zhuanlan.zhihu.com/p/80325757
// 基础实现 https://www.cnblogs.com/nullzx/p/7499397.html
// 改进实现 https://oi-wiki.org/string/ac-automaton/
// 应用 https://cp-algorithms.com/string/aho_corasick.html
// 模板题
// https://leetcode-cn.com/problems/stream-of-characters/
// https://www.luogu.com.cn/problem/P3808
// https://www.luogu.com.cn/problem/P3796
// todo https://www.luogu.com.cn/problem/P5357 二次加强版
// todo https://codeforces.com/problemset/problem/963/D
func (t *trie) buildDFA() {
	q := []*trieNode{}
	for _, son := range t.root.son {
		if son != nil {
			son.fail = t.root
			q = append(q, son)
		}
	}
	for len(q) > 0 {
		o := q[0]
		q = q[1:]
		if o.fail == nil {
			o.fail = t.root
		}
		for i, son := range o.son {
			if son != nil {
				son.fail = o.fail.son[i]
				q = append(q, son)
			} else {
				o.son[i] = o.fail.son[i]
			}
		}
	}
}

// 有多少个（编号）不同的模式串在文本串 text 里出现过
func (t *trie) sumCountAllPatterns(text []byte) (cnt int) {
	o := t.root
	for _, b := range text {
		o = o.son[t.ord(b)]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != nil && f.val > 0; f = f.fail {
			cnt += f.val
			f.val = 0
		}
	}
	return
}

// 返回所有模式串 patterns 的开头在文本串 text 的所有位置（未找到时对应数组为空）
// patterns 为模式串数组（无重复元素），为方便起见，patterns 从 1 开始
func (t *trie) acSearch(text []byte, patterns [][]byte) [][]int {
	pos := make([][]int, len(patterns))
	o := t.root
	for i, b := range text {
		o = o.son[t.ord(b)]
		if o == nil {
			o = t.root
			continue
		}
		for f := o; f != nil; f = f.fail {
			if pid := f.val; pid != 0 {
				pos[pid] = append(pos[pid], i-len(patterns[pid])+1) // 也可以只记录 i，代表模式串末尾在文本的位置
			}
		}
	}
	return pos
}

//

// 可持久化 trie
// TODO https://oi-wiki.org/ds/persistent-trie/
// 模板题（最大异或和） https://www.luogu.com.cn/problem/P4735

//

// Suffix automaton (SAM)
// https://en.wikipedia.org/wiki/Suffix_automaton
//《后缀自动机》，陈立杰
//《后缀自动机在字典树上的拓展》，刘研绎
//《后缀自动机及其应用》，张天扬
// todo https://baobaobear.github.io/post/20200220-sam/
// todo https://codeforces.com/blog/entry/20861
// TODO https://oi-wiki.org/string/sam/
// TODO https://cp-algorithms.com/string/suffix-automaton.html
//      后缀树简介 https://eternalalexander.github.io/2019/10/31/%E5%90%8E%E7%BC%80%E6%A0%91%E7%AE%80%E4%BB%8B/
// 模板题 https://www.luogu.com.cn/problem/P3804

// 广义 SAM
// todo https://www.luogu.com.cn/problem/P6139

// 回文自动机 PAM
// todo https://baobaobear.github.io/post/20200416-pam/
//  https://www.luogu.com.cn/problem/P5496
