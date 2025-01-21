package copypasta

import (
	"index/suffixarray"
	"math/bits"
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"strings"
	"time"
	"unsafe"
)

/*
库函数应用题
https://codeforces.com/problemset/problem/600/A 1600
https://atcoder.jp/contests/abc381/tasks/abc381_c split

字符串问题的特殊性：
不同子串之间会共享一些局部信息，巧妙地利用这些局部信息可以设计出更加高效的算法。

todo NOI 一轮复习 II：字符串 https://www.luogu.com.cn/blog/ix-35/noi-yi-lun-fu-xi-ii-zi-fu-chuan
金策 字符串算法选讲 https://www.bilibili.com/video/BV11K4y1p7a5 https://www.bilibili.com/video/BV19541177KU
    PDF 见 https://github.com/EndlessCheng/cp-pdf

TIPS: 若处理原串比较困难，不妨考虑下反转后的串 https://codeforces.com/contest/873/problem/F

本质不同回文串
https://codeforces.com/problemset/problem/1823/D https://codeforces.com/blog/entry/115465

斐波那契字符串：s(1) = "a", s(2) = "b", s(n) = s(n-1) + s(n-2), n>=3

https://en.wikipedia.org/wiki/Bitap_algorithm shift-or / shift-and / Baeza-Yates–Gonnet algorithm

给你两个字符串 s 和 t，问 s 有多少个子串和 t 是相似的
相似：如果至多交换一次字符串 S 中的两个相邻字母，使得 S = T，那么 S 和 T 就是相似的。

// Optimal trade-offs for pattern matching with k mismatches
// https://arxiv.org/abs/1704.01311

*/

func _() {
	// 注：如果 s 是常量的话，由于其在编译期分配到只读段，对应的地址是无法写入的
	unsafeToBytes := func(s string) []byte { return *(*[]byte)(unsafe.Pointer(&s)) }
	unsafeToString := func(b []byte) string { return *(*string)(unsafe.Pointer(&b)) }

	// 返回 t 在 s 中的所有位置（允许重叠）
	indexAll := func(s, t []byte) []int {
		pos := suffixarray.New(s).Lookup(t, -1)
		slices.Sort(pos)
		return pos
	}

	// 字符串哈希 rolling hash, Rabin–Karp algorithm
	//
	// mod 和 base 一定不能是固定的，哪怕用了双模数也不行！
	// 见 Anti-hash Test Generator https://heltion.github.io/anti-hash/ https://codeforces.com/blog/entry/129538
	// - 记得点一下 reverse 按钮
	// - 回文串题目勾选 ensure that two strings are the reverse of each other
	// - https://github.com/LeetCode-Feedback/LeetCode-Feedback/issues/24862
	// 可以用随机 base 避免被 hack
	//
	// https://en.wikipedia.org/wiki/Hash_function
	// https://en.wikipedia.org/wiki/Rolling_hash
	// https://en.wikipedia.org/wiki/Rabin%E2%80%93Karp_algorithm
	// https://oi-wiki.org/string/hash/
	// 如何卡自然溢出（随机 base 也可以卡）https://prutekoi.github.io/post/qia-zi-ran-yi-chu-ha-xi/
	// https://blog.csdn.net/weixin_45750972/article/details/107457997
	// On the mathematics behind rolling hashes and anti-hash tests https://codeforces.com/blog/entry/60442
	// Hacking a weak hash https://codeforces.com/blog/entry/113484
	// Rolling hash and 8 interesting problems https://codeforces.com/blog/entry/60445
	// 选一个合适的质数 https://planetmath.org/goodhashtableprimes
	// 999727999, 1070777777, 1000000007
	// 生日问题 https://en.wikipedia.org/wiki/Birthday_problem
	// 线性同余方法（LCG）https://en.wikipedia.org/wiki/Linear_congruential_generator
	// https://rng-58.blogspot.com/2017/02/hashing-and-probability-of-collision.html
	// 【推荐】滚动哈希和卡哈希的数学原理 https://notes.sshwy.name/Math/Rolling-Hash-and-Hack
	// 浅谈字符串 hash 的应用 https://www.luogu.com.cn/blog/Flying2018/qian-tan-zi-fu-chuan-hash
	// 从 Hash Killer I、II、III 论字符串哈希 https://www.cnblogs.com/HansBug/p/4288118.html
	// anti-hash: 最好不要自然溢出 https://codeforces.com/blog/entry/4898
	// hash killer https://loj.ac/p/6758
	// Kapun's algorithm https://codeforces.com/blog/entry/99973
	// 比较：给每个元素分配一个随机哈希系数 + 滑动窗口 https://codeforces.com/problemset/problem/1418/G
	//
	// Python/Java/C++ 实现见 https://leetcode.cn/problems/construct-string-with-minimum-cost/solution/hou-zhui-shu-zu-by-endlesscheng-32h9/
	//
	// 模板题 https://www.luogu.com.cn/problem/P3370
	// 测试哈希碰撞 https://codeforces.com/problemset/problem/514/C
	//            https://codeforces.com/problemset/problem/1200/E
	//            https://leetcode.cn/problems/count-prefix-and-suffix-pairs-ii/
	// 拼接字符串 https://codeforces.com/problemset/problem/1800/D
	// LC3213 最小代价构造字符串 https://leetcode.cn/problems/construct-string-with-minimum-cost/
	// LC187 找出所有重复出现的长为 10 的子串 https://leetcode.cn/problems/repeated-dna-sequences/
	// LC1044 最长重复子串（二分哈希）https://leetcode.cn/problems/longest-duplicate-substring/
	// LC1554 只有一个不同字符的字符串 https://leetcode.cn/problems/strings-differ-by-one-character/
	// https://codeforces.com/problemset?tags=hashing,strings
	// https://cp-algorithms.com/string/string-hashing.html#toc-tgt-7
	// 倒序哈希 https://leetcode.cn/problems/find-substring-with-given-hash-value/solution/dao-xu-hua-dong-chuang-kou-o1-kong-jian-xpgkp/
	// todo https://ac.nowcoder.com/acm/contest/64384/D
	// 与线段树结合，可以做到单点修改 s[i]
	// - 合并：hash(s+t) = hash(s) * base^|t| + hash(t)
	// 与线段树结合，可以做到区间更新
	// - [l,r] 区间加一：根据 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]
	//                 从右往左看，区间增加了 base^(n-1-r) + ... + base^(n-1-l)
	//                 见 segment_tree.go 中的「区间加等比数列」
	// - 区间替换成某个数：https://codeforces.com/problemset/problem/580/E 2500
	// todo 带修回文串 https://atcoder.jp/contests/abc331/tasks/abc331_f
	//      带修回文串 https://ac.nowcoder.com/acm/contest/90074/E
	// todo https://www.luogu.com.cn/problem/P2757
	stringHashSingleMod := func(s string) {
		// 如果 OJ 的 Go 版本低于 1.20，加上这句话
		rand.Seed(time.Now().UnixNano())

		// 下面实现的是单模哈希
		// 双模哈希见后面

		// mod 和 base 随机其中一个就行，无需两个都随机
		const mod = 1_070_777_777
		base := 9e8 - rand.Intn(1e8)

		// 多项式字符串哈希（方便计算子串哈希值）
		// 哈希函数 hash(s) = s[0] * base^(n-1) + s[1] * base^(n-2) + ... + s[n-2] * base + s[n-1]   其中 n 为 s 的长度
		powBase := make([]int, len(s)+1) // powBase[i] = base^i，用它当作哈希系数是为了方便求任意子串哈希，求拼接字符串的哈希等
		preHash := make([]int, len(s)+1) // preHash[i] = hash(s[:i]) 前缀哈希
		powBase[0] = 1
		for i, b := range s {
			powBase[i+1] = powBase[i] * base % mod
			preHash[i+1] = (preHash[i]*base + int(b)) % mod // 秦九韶算法计算多项式哈希
		}

		// 计算子串 s[l:r] 的哈希值，注意这是左闭右开区间 [l,r)    0<=l<=r<=len(s)
		// 空串的哈希值为 0
		// 计算方法类似前缀和
		subHash := func(l, r int) int {
			return ((preHash[r]-preHash[l]*powBase[r-l])%mod + mod) % mod
		}

		// 计算（准备与 s 匹配的）其他字符串的哈希值
		calcHash := func(t string) (h int) {
			for _, b := range t {
				h = (h*base + int(b)) % mod
			}
			return
		}

		// 比较 s[l1:r1] 和 s[l2:r2] 的字典序大小，注意这是左闭右开区间 [l,r)
		compare := func(l1, r1, l2, r2 int) int {
			len1, len2 := r1-l1, r2-l2
			sz := min(len1, len2)
			// 二分长度求 LCP
			lcp := sort.Search(sz, func(m int) bool {
				m++
				return subHash(l1, l1+m) != subHash(l2, l2+m)
			})
			if lcp == sz {
				// 一个是另一个的前缀，或者完全相等
				return len1 - len2
			}
			// 比较 LCP 的下一个字母
			return int(s[l1+lcp]) - int(s[l2+lcp])
		}

		// 计算 s[l1:r1] + s[l2:r2] 的哈希值，注意这是左闭右开区间 [l,r)
		concatHash := func(l1, r1, l2, r2 int) int {
			h1 := preHash[r1] - preHash[l1]*powBase[r1-l1]
			h2 := preHash[r2] - preHash[l2]*powBase[r2-l2]
			return ((h1%mod*powBase[r2-l2]+h2)%mod + mod) % mod
		}

		_ = []any{subHash, calcHash, compare, concatHash}
	}

	// 双模哈希 1133ms https://leetcode.cn/problems/construct-string-with-minimum-cost/submissions/545112087/
	// 单模哈希  468ms https://leetcode.cn/problems/construct-string-with-minimum-cost/submissions/545112323/
	// 时间多一到两倍是正常的（注意内存占用变大了，缓存也会影响性能）
	stringHashDoubleMod := func(s string) {
		const mod1 = 1_070_777_777
		const mod2 = 1_000_000_007
		base1 := 9e8 - rand.Intn(1e8)
		base2 := 9e8 - rand.Intn(1e8)

		type hPair struct{ h1, h2 int }
		powBase := make([]hPair, len(s)+1)
		preHash := make([]hPair, len(s)+1)
		powBase[0] = hPair{1, 1}
		for i, b := range s {
			powBase[i+1] = hPair{powBase[i].h1 * base1 % mod1, powBase[i].h2 * base2 % mod2}
			preHash[i+1] = hPair{(preHash[i].h1*base1 + int(b)) % mod1, (preHash[i].h2*base2 + int(b)) % mod2}
		}

		// 计算子串 s[l:r] 的哈希值
		// 空串的哈希值为 0
		subHash := func(l, r int) hPair {
			h1 := ((preHash[r].h1-preHash[l].h1*powBase[r-l].h1)%mod1 + mod1) % mod1
			h2 := ((preHash[r].h2-preHash[l].h2*powBase[r-l].h2)%mod2 + mod2) % mod2
			return hPair{h1, h2}
		}

		// 计算（准备与 s 匹配的）其他字符串的哈希值
		calcHash := func(t string) (p hPair) {
			for _, b := range t {
				p.h1 = (p.h1*base1 + int(b)) % mod1
				p.h2 = (p.h2*base2 + int(b)) % mod2
			}
			return
		}

		// 计算 s[l1:r1] + s[l2:r2] 的哈希值
		concatHash := func(l1, r1, l2, r2 int) hPair {
			h1 := (((preHash[r1].h1-preHash[l1].h1*powBase[r1-l1].h1)%mod1*powBase[r2-l2].h1+preHash[r2].h1-preHash[l2].h1*powBase[r2-l2].h1)%mod1 + mod1) % mod1
			h2 := (((preHash[r1].h2-preHash[l1].h2*powBase[r1-l1].h2)%mod2*powBase[r2-l2].h2+preHash[r2].h2-preHash[l2].h2*powBase[r2-l2].h2)%mod2 + mod2) % mod2
			return hPair{h1, h2}
		}

		_ = []any{subHash, calcHash, concatHash}
	}

	// todo 二维字符串哈希
	// https://www.luogu.com.cn/problem/solution/UVA11019
	// UVa 11019 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=22&page=show_problem&problem=1960

	//

	// KMP (Knuth–Morris–Pratt algorithm)
	// pi[i] 为 s[:i+1] 的真前缀和真后缀的最长的匹配长度    pi[0] = 0
	// 特别地，pi[n-1] 为 s 的真前缀和真后缀的最长的匹配长度
	//
	// KMP 是个不完全的 DFA，而 AC 自动机是一个完全的 DFA
	// 如果 KMP 也像 AC 自动机那样构建的话，虽然可以消除掉下面代码中的内层循环，但空间需要 * |Σ|
	//
	// 我在知乎上对 KMP 的讲解 https://www.zhihu.com/question/21923021/answer/37475572
	// https://en.wikipedia.org/wiki/Knuth%E2%80%93Morris%E2%80%93Pratt_algorithm
	// https://oi-wiki.org/string/kmp/ todo 统计每个前缀的出现次数
	// https://cp-algorithms.com/string/prefix-function.html
	// https://algs4.cs.princeton.edu/code/edu/princeton/cs/algs4/KMP.java.html
	//
	// todo 题单 https://www.luogu.com.cn/training/53971
	// 模板题 https://loj.ac/p/103 https://www.luogu.com.cn/problem/P3375
	//       LC28 https://leetcode.cn/problems/find-the-index-of-the-first-occurrence-in-a-string/
	//       LC1392 https://leetcode.cn/problems/longest-happy-prefix/
	//       LC3036 https://leetcode.cn/problems/number-of-subarrays-that-match-a-pattern-ii/ 
	//       LC3037 https://leetcode.cn/problems/find-pattern-in-infinite-stream-ii/
	// LC3008 https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-ii/ 2016 *二分/双指针找最近下标
	// LC686 https://leetcode.cn/problems/repeated-string-match/
	// - a 复制 k 或 k+1 份，k=(len(b)-1)/len(a)+1
	// 最长回文前缀 LC214 https://leetcode.cn/problems/shortest-palindrome/
	// LC1316 https://leetcode.cn/problems/distinct-echo-substrings/ 1837
	// https://codeforces.com/problemset/problem/1137/B 1600 构造
	// https://codeforces.com/problemset/problem/126/B 1700
	// https://codeforces.com/problemset/problem/471/D 1800
	// https://codeforces.com/problemset/problem/1269/B ~1800 做到 O(nlogn)
	// 与 LCS 结合 https://codeforces.com/problemset/problem/346/B 2000
	// 与计数 DP 结合 https://codeforces.com/problemset/problem/494/B 2000
	// https://codeforces.com/problemset/problem/1200/E 2000
	// 最大匹配个数 https://codeforces.com/problemset/problem/615/C 2000
	// 与 DP 结合 https://codeforces.com/problemset/problem/1163/D 2100
	// https://codeforces.com/problemset/problem/526/D 2200
	// https://codeforces.com/problemset/problem/954/I 2200
	// https://codeforces.com/problemset/problem/1003/F 2200
	// 构造 t+"#"+s https://codeforces.com/problemset/problem/25/E 2200
	// - 不加 # 的话会面临 "cabc"+"abca" 这样的例子，算出的 border 是 "cabca"
	// - LC2800 https://leetcode.cn/problems/shortest-string-that-contains-three-strings/ 1856
	// todo 与 DP 结合 https://codeforces.com/problemset/problem/808/G 2300
	// LC1397 与数位 DP 结合 https://leetcode.cn/problems/find-all-good-strings/ 2667
	// http://acm.hdu.edu.cn/showproblem.php?pid=2087
	// - https://oj.socoding.cn/p/1446 
	// - https://github.com/tdzl2003/leetcode_live/blob/master/socoding/1446.md
	// 在循环同构字符串 s 中查找 t，等价于在 s+(s[:n-1]) 中查找 t 
	// - LC2851 https://leetcode.cn/problems/string-transformation/ 2858
	// https://www.lanqiao.cn/problems/5132/learning/?contest_id=144
	// todo https://www.luogu.com.cn/problem/P4391
	//  https://www.luogu.com.cn/problem/UVA10298
	//  https://www.luogu.com.cn/problem/P3435
	//  https://www.luogu.com.cn/problem/UVA11022
	//  https://www.luogu.com.cn/problem/P4824
	//  https://www.luogu.com.cn/problem/P2375
	//  https://www.luogu.com.cn/problem/P7114
	//  https://www.luogu.com.cn/problem/P3426
	//  https://www.luogu.com.cn/problem/P3193
	//  https://www.luogu.com.cn/problem/P4503
	//  https://www.luogu.com.cn/problem/P3538
	//  https://www.luogu.com.cn/problem/P4036

	// 计算前缀函数
	// pi[i] 为 s[:i+1] 的真前缀和真后缀的最长匹配长度    pi[0] = 0
	// 定义 s[:i+1] 的最大真 border 为 s[:pi[i]]    完整定义见 https://www.luogu.com.cn/problem/P5829
	// 注：「真」表示不等于整个字符串
	calcPi := func(s string) []int {
		pi := make([]int, len(s))
		match := 0
		for i := 1; i < len(pi); i++ {
			v := s[i]
			for match > 0 && s[match] != v {
				match = pi[match-1]
			}
			if s[match] == v {
				match++
			}
			pi[i] = match
		}
		return pi
	}

	// 在文本串 text 中查找模式串 pattern，返回所有成功匹配的位置（pattern[0] 在 text 中的下标）
	kmpSearch := func(text, pattern string) (pos []int) {
		pi := calcPi(pattern)
		match := 0
		for i := range text {
			v := text[i]
			for match > 0 && pattern[match] != v {
				match = pi[match-1]
			}
			if pattern[match] == v {
				match++
			}
			if match == len(pi) {
				pos = append(pos, i-len(pi)+1)
				match = pi[match-1] // 如果不允许重叠，将 cnt 置为 0
			}
		}
		return
	}

	// EXTRA: 最小循环节
	// 返回循环节以及循环次数
	// 如果没有循环节，那么返回原串和 1
	// https://codeforces.com/problemset/problem/182/D 1400
	// https://codeforces.com/problemset/problem/1690/F 1700
	// https://codeforces.com/problemset/problem/526/D 2200
	// http://poj.org/problem?id=2406 https://www.luogu.com.cn/problem/UVA455
	// LC459 https://leetcode.cn/problems/repeated-substring-pattern/
	calcMinPeriod := func(s string) (string, int) {
		n := len(s)
		pi := calcPi(s)
		if m := pi[n-1]; m > 0 && n%(n-m) == 0 {
			return s[:n-m], n / (n - m)
		}
		return s, 1 // 无小于 n 的循环节
	}

	// fail 树 / 失配树 / border 树
	// i+1 的父节点为 pi[i]    注：i 从 0 开始
	// 性质：
	// 1. s[:i] 的所有 border 长度：i 到根节点路径上的所有节点
	// 2.（1 的推论）前缀 s[:i] 有长为 x 的真 border：等价于 x 是 i 的祖先节点
	// 3. s[:i] 和 s[:j] 的最长公共 border：s[:LCA(i,j)]   
	// - 求真 border 的情况见下面的注释
	// https://ac.nowcoder.com/study/live/738/3/1
	// https://www.luogu.com.cn/problem/P5829
	failTree := func(s string) {
		pi := calcPi(s)
		g := make([][]int, len(s)+1)
		for i, p := range pi {
			g[p] = append(g[p], i+1)
		}

		// 求 s[:i] 和 s[:j] 的最长公共真 border
		// fail 树上跑 LCA
		// 直接把 graph_tree.go 中的倍增 LCA 搬过来
		// 唯一需要修改的是 getLCA 中，当 w==v 时，改成返回 pa[v][0]，表示真 border
		// https://www.luogu.com.cn/record/136415394
	}

	//

	// Z-function（扩展 KMP，Z-array）      exkmp
	// z[i] = LCP(s, s[i:])   串与串后缀的最长公共前缀
	//
	// 核心思想：
	// z-Box：上次暴力匹配的左右边界
	// z-Box 的「影子」是等长的 s 的前缀，例如 s=abababzabababab 中 s[7:13] 的影子是 s[:6]，这两段是完全一样的
	// 所以 z-Box 中的后缀，和影子中的后缀是一样的
	// 如果 i 在 z-Box 中，那么 LCP(s, s[i:]) 可以提前计算出一部分（而不是暴力），
	// 这可以通过「影子」中的对应位置的 LCP(s, s[i-boxL:]) 得出，
	// 例如 s=abababzabababab，要计算 LCP(s, s[9:])，可以看影子中的对应位置 LCP(s, s[2:]) = 4，从而直接得到 LCP(s, s[9:]) 至少是 4
	//
	// 视频讲解 https://www.bilibili.com/video/BV1it421W7D8/
	// https://oi-wiki.org/string/z-func/
	// 可视化 https://personal.utdallas.edu/~besp/demo/John2010/z-algorithm.htm
	// - 用这个 abababzabababab
	// 如果 i 在 z-Box 中，那么根据 z[i-boxL] 和 z[boxL] 的信息可以推出，有 min(z[i-boxL], boxR-i+1) 的长度是已匹配的，无需暴力匹配
	// https://cp-algorithms.com/string/z-function.html
	// https://www.geeksforgeeks.org/z-algorithm-linear-time-pattern-searching-algorithm/
	//
	// 模板题
	// https://judge.yosupo.jp/problem/zalgorithm
	// https://codeforces.com/edu/course/2/lesson/3/3/practice/contest/272263/problem/A
	// https://www.luogu.com.cn/problem/P5410
	//
	// https://codeforces.com/problemset/problem/2010/C2 1700
	// todo 结论 https://codeforces.com/problemset/problem/535/D 1900
	// https://codeforces.com/problemset/problem/1968/G2 2200
	// https://codeforces.com/problemset/problem/1051/E 2600 DP 
	// 最小循环节（允许末尾截断）https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/A
	// s 和 t 是否本质相同，shift 多少次 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/B
	//		即 strings.Index(s+s, t)
	// 每个前缀的出现次数 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/C
	//      例如 z[i]=3，那么把长为 1,2,3 的前缀的出现次数都 +1
	//		用差分数组实现，代码 https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/submission/89189534
	//		注：字符串倒过来就是每个后缀的出现次数
	// 既是前缀又是后缀的子串个数 https://codeforces.com/problemset/problem/432/D
	//		方法一
	//      求出 z 数组，以及 z 数组排序后的数组 sortedZ。
	//      设长为 L 的同时是 s 前缀和后缀的字符串为 t，那么 t 需要满足 z[n-L] == L。
	//      如果 t 出现在 s 的中间某一段，那么 t 对应的 z[i] 一定 >= L。（后缀的前缀是子串。）
	//      所以，在 sortedZ 中二分找第一个 >= L 的数的下标 j，那么 t 的出现次数就是 n-j。
	//      时间复杂度 O(nlogn)。
	//      https://codeforces.com/contest/432/submission/247583742
	//
	//      方法二
	//      统计 z[i] 的出现次数，记到 cnt 数组中。
	//      原地计算 cnt 数组的后缀和，那么 cnt[L] 就是在方法一中，原本要通过排序+二分才能算出的个数了。
	//      时间复杂度 O(n)。
	//      https://codeforces.com/contest/432/submission/247584529
	//
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
		boxL, boxR := 0, 0 // z-box 左右边界（闭区间）
		for i := 1; i < n; i++ {
			if i <= boxR {
				z[i] = min(z[i-boxL], boxR-i+1)
			}
			for i+z[i] < n && s[z[i]] == s[i+z[i]] {
				boxL, boxR = i, i+z[i]
				z[i]++
			}
		}
		z[0] = n
		return z
	}
	// 在 text 中查找 pattern 的所有（首字母）位置
	// 技巧：把 pattern 拼在 text 前面（中间插入一个不在输入中的字符），得到字符串 s，
	//      只要 LCP(s, s[i:]) == len(pattern)，就说明 i 是一个匹配的位置
	// 可以用这题测试 LC3008 https://leetcode.cn/problems/find-beautiful-indices-in-the-given-array-ii/ 2016
	zSearch := func(text, pattern string) (pos []int) {
		z := calcZ(pattern + "#" + text)
		for i, l := range z[len(pattern)+1:] {
			if l == len(pattern) {
				pos = append(pos, i)
			}
		}
		return
	}
	// zSearch 中的技巧还可以用来比较 text 的任意后缀 text[i:] 与 pattern 的大小（字典序）
	// res[i] 等同于 strings.Compare(text[i:], pattern)
	// https://codeforces.com/contest/1051/problem/E 2600
	zCompare := func(text, pattern string) []int {
		z := calcZ(pattern + "#" + text)
		compare := func(i int) int {
			lcp := z[len(pattern)+1+i]
			if lcp == len(pattern) { // 相等
				return 0
			}
			// 比较 LCP 的下一个字母
			if text[i+lcp] < pattern[lcp] {
				return -1
			}
			return 1
		}
		res := make([]int, len(text))
		for i := range res {
			res[i] = compare(i)
		}
		return res
	}

	// Main–Lorentz 算法
	// https://oi-wiki.org/string/main-lorentz/
	mainLorentz := func() {
		// todo
	}
	_ = mainLorentz

	// 最小表示法 - 求串的循环同构串中字典序最小的串
	// 找到位置 i，从这个位置输出即得到字典序最小的串
	// https://oi-wiki.org/string/minimal-string/
	// 其他方法 https://codeforces.com/blog/entry/90035
	// 模板题 https://www.luogu.com.cn/problem/P1368 http://poj.org/problem?id=1509 https://codeforces.com/gym/103585/problem/K
	// https://codeforces.com/problemset/problem/496/B
	// LC1163 非循环的情况 https://leetcode.cn/problems/last-substring-in-lexicographical-order/
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

	// 判断子序列 / 最长匹配长度
	// 返回 s 最长前缀的长度，满足该前缀是 t 的子序列
	// 见 https://leetcode.cn/circle/discuss/0viNMK/ 中的【判断子序列】
	// https://leetcode.cn/problems/is-subsequence/
	// https://leetcode.cn/problems/subsequence-with-the-minimum-score/
	// - https://leetcode.cn/problems/find-the-lexicographically-smallest-valid-sequence/
	// https://codeforces.com/problemset/problem/1194/C 1300
	// https://codeforces.com/problemset/problem/778/A 1700
	isSubseq := func(s, t string) int {
		if s == "" {
			return 0
		}
		cnt := 0
		for _, b := range t {
			if s[cnt] != byte(b) {
				continue
			}
			cnt++
			if cnt == len(s) {
				break
			}
		}
		return cnt
	}

	// 子序列自动机
	// 如果值域很大可以用哈希表/数组记录 pos 然后二分查找 https://www.luogu.com.cn/problem/P5826
	// https://leetcode.cn/problems/is-subsequence/
	// - [514. 自由之路](https://leetcode.cn/problems/freedom-trail/)
	// LC727 https://leetcode.cn/problems/minimum-window-subsequence/
	// LC792 https://leetcode.cn/problems/number-of-matching-subsequences/
	// LC2014 https://leetcode.cn/problems/longest-subsequence-repeated-k-times/
	// LC466 https://leetcode.cn/problems/count-the-repetitions/
	// - [727. 最小窗口子序列](https://leetcode.cn/problems/minimum-window-subsequence/)（会员题）
	// https://codeforces.com/problemset/problem/91/A
	// - https://www.luogu.com.cn/problem/P9572?contestId=124047
	// - 【子串】 LC686 https://leetcode.cn/problems/repeated-string-match/
	// https://codeforces.com/contest/1845/problem/C
	// - 相关 LC2350 https://leetcode.cn/problems/shortest-impossible-sequence-of-rolls/

	// 写法一
	// nxt[i][j] 表示下标 > i 的最近字符 j 的下标
	subsequenceAutomaton := func(s string) {
		const base = 'a'
		// build nxt
		pos := [26]int{}
		for i := range pos {
			pos[i] = len(s)
		}
		nxt := make([][26]int, len(s))
		for i := len(s) - 1; i >= 0; i-- {
			nxt[i] = pos
			pos[s[i]-base] = i
		}

		// 返回是 s 的子序列的最长的 t 的前缀的长度
		match := func(t string) int {
			if t == "" || s == "" {
				return 0
			}
			i, j := 0, 0
			if t[0] == s[0] {
				j = 1 // t[0] 匹配 ok
			}
			for ; j < len(t); j++ {
				i = nxt[i][t[j]-base]
				if i == len(s) {
					break
				}
			}
			return j
		}
		_ = match
	}

	// 写法二
	// nxt[i][j] 表示下标 >= i 的最近字符 j 的下标
	subsequenceAutomaton2 := func(s string) {
		const base = 'a'
		// build nxt
		pos := [26]int{}
		n := len(s)
		for i := range pos {
			pos[i] = n
		}
		nxt := make([][26]int, n+1)
		nxt[n] = pos
		for i := n - 1; i >= 0; i-- {
			pos[s[i]-base] = i
			nxt[i] = pos
		}

		// 返回是 s 的子序列的最长的 t 的前缀的长度
		match := func(t string) int {
			i := -1
			for j, b := range t {
				i = nxt[i+1][b-base]
				if i == n { // 找不到 t[j]
					return j
				}
			}
			// 此时 s[i] 匹配 t[-1]
			return len(t)
		}
		_ = match
	}

	// 最长回文子串 Manacher（马拉车算法）
	// 视频讲解：https://www.bilibili.com/video/BV1UcyYY4EnQ/
	// https://blog.csdn.net/synapse7/article/details/18908413
	// https://www.bilibili.com/video/BV1AX4y1F79W
	// https://www.bilibili.com/video/BV1ft4y117a4
	// https://oi-wiki.org/string/manacher/
	// https://cp-algorithms.com/string/manacher.html
	// https://codeforces.com/blog/entry/12143
	// https://leetcode.cn/problems/longest-palindromic-substring/solution/zui-chang-hui-wen-zi-chuan-by-leetcode-solution/
	// https://www.zhihu.com/question/37289584
	// 可视化 http://manacher-viz.s3-website-us-east-1.amazonaws.com
	//
	// https://oeis.org/A002620 全为 a 的字符串的奇回文子串个数 floor((n+1)^2/4)
	// https://oeis.org/A002620 全为 a 的字符串的偶回文子串个数 floor(n^2/4)
	// https://oeis.org/A000217 全为 a 的字符串的回文子串个数 n*(n+1)/2
	//
	// todo 题单 https://www.luogu.com.cn/training/53971
	// 模板题 https://judge.yosupo.jp/problem/enumerate_palindromes
	//       https://www.luogu.com.cn/problem/P3805
	//       LC5 https://leetcode.cn/problems/longest-palindromic-substring/
	// https://codeforces.com/problemset/problem/1326/D2 1800 去掉子串后，剩余部分是回文串
	// https://codeforces.com/problemset/problem/7/D 2200
	// - https://codeforces.com/problemset/problem/835/D 1900
	// https://codeforces.com/problemset/problem/1827/C 2600
	// https://www.luogu.com.cn/problem/P4555
	// todo 相交的回文子串对数 https://codeforces.com/problemset/problem/17/E
	//  https://codeforces.com/problemset/problem/1081/H
	//  https://www.luogu.com.cn/blog/user25308/proof-cf1081h
	//  LC1745 分割成三个非空回文子字符串 https://leetcode.cn/problems/palindrome-partitioning-iv/
	// LC2472 不重叠回文子字符串（长度至少为 k）的最大数目 https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/
	// - 只需要考虑长度为 k or k+1 的
	// todo https://www.luogu.com.cn/problem/P1659
	//  https://www.luogu.com.cn/problem/P3501
	//  https://www.luogu.com.cn/problem/UVA11475
	//  https://www.luogu.com.cn/problem/P6216
	//  https://www.luogu.com.cn/problem/P5446
	manacher := func(s string) {
		// 将 s 改造为 t，这样就不需要分 len(s) 的奇偶来讨论了，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
		// s 和 t 的下标转换关系：
		// (si+1)*2 = ti
		// ti/2-1 = si
		// ti 为偶数，对应奇回文串（从 2 开始）
		// ti 为奇数，对应偶回文串（从 3 开始）
		t := append(make([]byte, 0, len(s)*2+3), '^')
		for _, c := range s {
			t = append(t, '#', byte(c))
		}
		t = append(t, '#', '$')

		// 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
		// halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
		// 即 [i-halfLen[i]+1,i+halfLen[i]-1] 是 t 上的一个回文子串
		halfLen := make([]int, len(t)-2)
		halfLen[1] = 1
		// boxR 表示当前右边界下标最大的回文子串的右边界下标+1
		// boxM 为该回文子串的中心位置，二者的关系为 r=mid+halfLen[mid]
		boxM, boxR := 0, 0
		for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原串的首尾字符
			hl := 1
			if i < boxR {
				// 记 i 关于 boxM 的对称位置 i'=boxM*2-i
				// 若以 i' 为中心的最长回文子串范围超出了以 boxM 为中心的回文串的范围（即 i+halfLen[i'] >= boxR）
				// 则 halfLen[i] 应先初始化为已知的回文半径 boxR-i，然后再继续暴力匹配
				// 否则 halfLen[i] 与 halfLen[i'] 相等
				hl = min(halfLen[boxM*2-i], boxR-i)
			}
			// 暴力扩展
			// 算法的复杂度取决于这部分执行的次数
			// 由于扩展之后 boxR 必然会更新（右移），且扩展的的次数就是 boxR 右移的次数
			// 因此算法的复杂度 = O(len(t)) = O(len(s))
			for t[i-hl] == t[i+hl] {
				hl++
				boxM, boxR = i, i+hl
			}
			halfLen[i] = hl
		}

		// t 中回文子串的长度为 hl*2-1
		// 由于其中 # 的数量总是比字母的数量多 1
		// 因此其在 s 中对应的回文子串的长度为 hl-1
		// 这一结论可用在下面的各个代码中

		// 判断闭区间 [l,r] 是否为回文串  0<=l<=r<len(s)
		// 根据下标转换关系得到子串 s[l:r+1] 在 t 中对应的回文中心下标为 l+r+2
		// https://codeforces.com/problemset/problem/1326/D2
		// https://codeforces.com/problemset/problem/7/D 
		// - https://codeforces.com/problemset/problem/835/D
		// LC3327 https://leetcode.cn/problems/check-if-dfs-strings-are-palindromes/
		isP := func(l, r int) bool { return halfLen[l+r+2] > r-l+1 } // halfLen[l+r+2]-1 >= r-l+1

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
		// LC1960 两个回文子字符串长度的最大乘积 https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-substrings/
		// LC214 https://leetcode.cn/problems/shortest-palindrome/
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

		// todo EXTRA: 以 s[i] 为尾字母的最短回文子串的长度
		// 结合单调栈
		// 回文中心越来越大
		// https://codeforces.com/contest/1827/problem/C

		// EXTRA: 计算回文子串个数
		// 易证其为 ∑(halfLen[i]/2)
		// LC647 https://leetcode.cn/problems/palindromic-substrings/
		totP := 0
		for _, hl := range halfLen {
			totP += hl / 2
		}

		// EXTRA: 线段树区间更新可以求出前缀/后缀回文子串个数

		// todo 任意区间回文子串个数

		_ = []interface{}{isP, midPL}
	}

	// todo 只考虑奇回文串
	manacherOdd := func() {}
	_ = manacherOdd

	// 只考虑偶回文串
	// todo https://codeforces.com/contest/1827/submission/205941641
	manacherEven := func() {}
	_ = manacherEven

	/* 后缀数组
	https://riteme.site/blog/2016-6-19/sais.html（包含 SA-IS 与 DC3 的效率对比）
	https://zork.net/~st/jottings/sais.html
	https://zhuanlan.zhihu.com/p/338547483
	注：Go1.13 开始使用 SA-IS 算法

	可视化 https://visualgo.net/zh/suffixarray

	常用分隔符 #(35) $(36) _(95) |(124)

	讲解+例题+套题 https://oi-wiki.org/string/sa/
	todo 题目推荐 https://www.luogu.com.cn/blog/luckyblock/post-bi-ji-hou-zhui-shuo-zu
	CF 上的课程 https://codeforces.com/edu/course/2
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=string+suffix+structures

	题目总结：（部分参考《后缀数组——处理字符串的有力工具》，PDF 见 https://github.com/EndlessCheng/cp-pdf）
	单个字符串
		模板题
			https://www.luogu.com.cn/problem/P3809
			https://atcoder.jp/contests/abc362/tasks/abc362_g 子串计数 注意有重复的 t
			https://judge.yosupo.jp/problem/suffixarray
			https://loj.ac/p/111
		可重叠最长重复子串
			LC1044 https://leetcode.cn/problems/longest-duplicate-substring/
			LC1062 https://leetcode.cn/problems/longest-repeating-substring/
			相当于求 max(height)，实现见下面的 longestDupSubstring
		不可重叠最长重复子串
			https://atcoder.jp/contests/abc141/tasks/abc141_e
			- http://poj.org/problem?id=1743
			可参考《算法与实现》p.223 以及 https://oi-wiki.org/string/sa/#是否有某字符串在文本串中至少不重叠地出现了两次
			重要技巧：按照 height 分组，每组中根据 sa 来处理组内后缀的位置
		可重叠的至少出现 k 次的最长重复子串
			https://www.luogu.com.cn/problem/P2852
			- http://poj.org/problem?id=3261
			二分答案，对 height 分组，判定组内元素个数不小于 k
		本质不同子串个数
			LC1698 https://leetcode.cn/problems/number-of-distinct-substrings-in-a-string/
			https://www.luogu.com.cn/problem/P2408
			https://www.luogu.com.cn/problem/SP694
			https://judge.yosupo.jp/problem/number_of_substrings
			https://atcoder.jp/contests/practice2/tasks/practice2_i
			https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/A
			枚举每个后缀，计算前缀总数，再减掉重复，即 height[i]
			所以个数为 n*(n+1)/2-sum{height[i]} https://oi-wiki.org/string/sa/#_13
			相似思路 LC2261 含最多 K 个可整除元素的子数组 https://leetcode.cn/problems/k-divisible-elements-subarrays/solution/by-freeyourmind-2m6j/
		不同子串长度之和
			https://codeforces.com/edu/course/2/lesson/3/4/practice/contest/272262/problem/H
			思路同上，即 n*(n+1)*(n+2)/6-sum{height[i]*(height[i]+1)/2}
		带限制的不同子串个数
			https://codeforces.com/problemset/problem/271/D
			这题可以枚举每个后缀，跳过 height[i] 个字符，然后在前缀和上二分
		重复次数最多的连续重复子串
			https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/F
			- http://poj.org/problem?id=3693 (数据弱)
			核心思想是枚举长度然后计算 LCP(i,i+l)，然后看是否还能再重复一次，具体代码见 main/edu/2/suffixarray/step5/f/main.go
		重复两次的最长连续重复子串
			LC1316 https://leetcode.cn/problems/distinct-echo-substrings/
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
			 https://atcoder.jp/contests/abc213/tasks/abc213_f
		从字符串首尾取字符最小化字典序
			https://oi-wiki.org/string/sa/#_10
			todo
		第 k 小子串
			https://www.luogu.com.cn/problem/P3975 不同位置的相同子串算作一个/多个
			todo https://codeforces.com/problemset/problem/128/B 2100
			todo
	两个字符串
		最长公共子串 LC718 https://leetcode.cn/problems/maximum-length-of-repeated-subarray/
	               LC3135 https://leetcode.cn/problems/equalize-strings-by-adding-or-removing-characters-at-ends/
	               SPOJ LCS https://www.luogu.com.cn/problem/SP1811
	               https://codeforces.com/edu/course/2/lesson/2/5/practice/contest/269656/problem/B
	               http://poj.org/problem?id=2774
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
		多串最长公共子串 SPOJ LCS2 https://www.luogu.com.cn/problem/SP1812 https://loj.ac/p/171
	        LC1923 https://leetcode.cn/problems/longest-common-subpath/ http://poj.org/problem?id=3450
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数等于字符串个数
		不小于 k 个字符串中的最长子串 http://poj.org/problem?id=3294
			拼接，二分答案，对 height 分组，判定组内元素对应不同字符串的个数不小于 k
		在每个字符串中至少出现两次且不重叠的最长子串 https://www.luogu.com.cn/problem/SP220
			拼接，二分答案，对 height 分组，判定组内元素在每个字符串中至少出现两次且 sa 的最大最小之差不小于二分值（用于判定是否重叠）
		出现或反转后出现在每个字符串中的最长子串 http://poj.org/problem?id=1226
			拼接反转后的串 s[i]+="#"+reverse(s)，拼接所有串，二分答案，对 height 分组，判定组内元素在每个字符串或其反转串中出现
		acSearch（https://www.luogu.com.cn/problem/P3796）的后缀数组做法
			拼接所有串（模式+文本，# 隔开），对每个模式 p 找其左右范围，满足该范围内 height[i] >= len(p)，这可以用 ST+二分或线段树二分求出，然后统计区间内的属于文本串的后缀
		不在其它字符串中出现的最短字典序最小子串 https://leetcode.cn/problems/shortest-uncommon-substring-in-an-array/
	逆向
		todo 根据 sa 反推有多少个能生成 sa 的字符串 https://codeforces.com/problemset/problem/1526/E
	todo 待整理 http://poj.org/problem?id=3581
	 https://www.luogu.com.cn/problem/P5546
	 https://www.luogu.com.cn/problem/P2048
	 https://www.luogu.com.cn/problem/P4248
	 https://www.luogu.com.cn/problem/P4341
	 https://www.luogu.com.cn/problem/P6095
	 https://www.luogu.com.cn/problem/P4070
	*/
	suffixArray := func(s string) {
		// 后缀数组 sa（后缀序）
		// sa[i] 表示后缀字典序中的第 i 个字符串在 s 中的位置
		// 特别地，后缀 s[sa[0]:] 字典序最小，后缀 s[sa[n-1]:] 字典序最大
		//sa := *(*[]int)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").UnsafeAddr()))
		sa := *(*[]int32)(unsafe.Pointer(reflect.ValueOf(suffixarray.New([]byte(s))).Elem().FieldByName("sa").Field(0).UnsafeAddr()))

		{
			// 不用反射（推荐）
			type _tp struct {
				_  []byte
				sa []int32
			}
			sa = (*_tp)(unsafe.Pointer(suffixarray.New([]byte(s)))).sa
		}

		// 后缀名次数组 rank（相当于 sa 的反函数）
		// 后缀 s[i:] 位于后缀字典序中的第 rank[i] 个
		// 特别地，rank[0] 即 s 在后缀字典序中的排名，rank[n-1] 即 s[n-1:] 在字典序中的排名
		rank := make([]int, len(sa))
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
		height := make([]int, len(sa))
		h := 0
		for i, rk := range rank {
			if h > 0 {
				h--
			}
			if rk > 0 {
				for j := int(sa[rk-1]); i+h < len(s) && j+h < len(s) && s[i+h] == s[j+h]; h++ {
				}
			}
			height[rk] = h
		}

		// 任意两后缀的 LCP
		// 注：若允许离线可以用 Trie+Tarjan 做到线性
		//st := make([][17]int, n) // 131072, 262144, 524288, 1048576
		logN := bits.Len(uint(len(sa)))
		st := make([][]int, len(sa))
		for i, v := range height {
			st[i] = make([]int, logN)
			st[i][0] = v
		}
		for j := 1; 1<<j <= len(sa); j++ {
			for i := 0; i+1<<j <= len(sa); i++ {
				st[i][j] = min(st[i][j-1], st[i+1<<(j-1)][j-1])
			}
		}
		_q := func(l, r int) int { k := bits.Len(uint(r-l)) - 1; return min(st[l][k], st[r-1<<k][k]) }
		lcp := func(i, j int) int {
			if i == j {
				return len(sa) - i
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
		// LC1977 https://leetcode.cn/problems/number-of-ways-to-separate-numbers/
		compareSub := func(l1, r1, l2, r2 int) int {
			len1, len2 := r1-l1, r2-l2
			l := lcp(l1, l2)
			if l >= min(len1, len2) {
				// 一个是另一个的前缀，或者完全相等
				return len1 - len2
			}
			// 或者 int(s[l1+l]) - int(s[l2+l])
			return rank[l1] - rank[l2]
		}

		// EXTRA: 可重叠最长重复子串
		// LC1044 https://leetcode.cn/problems/longest-duplicate-substring/
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
		vals := make([]int, len(sa))
		prefixSum := make([]int, len(sa)+1)
		for i, p := range sa {
			prefixSum[i+1] = prefixSum[i] + vals[p]
		}

		// EXTRA: 找出数组中的所有字符串，其是某个字符串的子串
		// 先拼接字符串，然后根据 height 判断前后是否有能匹配的
		// NOTE: 下面的代码展示了一种「标记 s[i] 属于原数组的哪个元素」的技巧: 在 i>0&&s[i]=='#' 时将 cnt++，其余的 s[i] 指向的 cnt 就是原数组的下标
		// LC1408 https://leetcode.cn/problems/string-matching-in-an-array/ 「小题大做」
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
			// 计算 sa & height ...
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
		for i, h := range height {
			suffix := s[sa[i]:]
			if h == 0 {
				println(" ", suffix)
			} else {
				println(h, suffix)
			}
		}

		// 注：极限优化用
		// 同 sa.Lookup([]byte(t), -1) 
		// 但 sa.Lookup 会 copy sa slice，下面的代码可以避免 copy
		// ！请勿修改返回值
		lookUp := func(t string) []int32 {
			// 找后缀序中的第一个 >= t 的后缀
			i := sort.Search(len(sa), func(i int) bool { return strings.Compare(s[sa[i]:], t) >= 0 })
			// 找后缀序中的第一个「前缀不含 t」的后缀
			j := i + sort.Search(len(sa)-i, func(j int) bool { return !strings.HasPrefix(s[sa[j+i]:], t) })
			return sa[i:j]
		}

		_ = []any{lessSub, compareSub, equalSub, longestDupSubstring, findAllSubstring, lookUp}
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

	// O(n^2) 计算 LCP —— 如果你不想用后缀数组的话
	// 见字符串题单 https://leetcode.cn/circle/discuss/SJFwQI/
	// https://codeforces.com/problemset/problem/1948/D 1700 也有更简单的做法
	// https://codeforces.com/problemset/problem/2045/H 2200
	lcpArray := func(s string) {
		n := len(s)
		lcp := make([][]int, n+1)
		for i := range lcp {
			lcp[i] = make([]int, n+1)
		}
		for i := n - 1; i >= 0; i-- {
			for j := n - 1; j >= 0; j-- { // 或者 j >= i，j > i
				if s[i] == s[j] {
					lcp[i][j] = lcp[i+1][j+1] + 1
				}
			}
		}

		// 返回 strings.Compare(s[l1:r1], s[l2:r2])
		// 如果上面写 j >= 0，那么这两个子串没有左右位置要求
		compare := func(l1, r1, l2, r2 int) int {
			len1, len2 := r1-l1, r2-l2
			l := lcp[l1][l2]
			if l >= min(len1, len2) {
				if len1 == len2 { // 相等
					return 0
				}
				if len1 < len2 { // <    s[l1:r1] 是 s[l2:r2] 的前缀
					return -1
				}
				return 1 // >    s[l2:r2] 是 s[l1:r1] 的前缀
			}
			if s[l1+l] < s[l2+l] { // <
				return -1
			}
			return 1 // >
		}

		// compare 的简化版
		less := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l1, r2-l2
			l := lcp[l1][l2]
			if l >= min(len1, len2) {
				return len1 < len2
			}
			return s[l1+l] < s[l2+l]
		}

		// 判断 s[l1:r1] 是否为 s[l2:r2] 的前缀
		isPrefix := func(l1, r1, l2, r2 int) bool {
			len1, len2 := r1-l1, r2-l2
			return len1 <= len2 && lcp[l1][l2] >= len1
		}

		_ = []any{compare, less, isPrefix}
	}

	_ = []any{
		unsafeToBytes, unsafeToString,
		indexAll,
		stringHashSingleMod, stringHashDoubleMod,
		kmpSearch, calcMinPeriod, failTree, // KMP
		calcZ, zSearch, zCompare, // Z 函数
		smallestRepresentation,
		isSubseq, subsequenceAutomaton, subsequenceAutomaton2,
		manacher,
		suffixArray, suffixArrayInt, suffixArrayInt2, // 后缀数组
		lcpArray,
	}
}

// AC 自动机见 trie.go
// 后缀自动机见 sam.go
