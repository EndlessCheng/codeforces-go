## 方法一：中心扩展法

**前置题目**：[5. 最长回文子串](https://leetcode.cn/problems/longest-palindromic-substring/)，[我的题解](https://leetcode.cn/problems/longest-palindromic-substring/solutions/2958179/mo-ban-on-manacher-suan-fa-pythonjavacgo-t6cx/)。

用中心扩展法，枚举最终答案的回文中心，向外扩展。

如果发现 $s[l]\ne s[r]$，那么必须删除字母，才能继续扩展：

- 选择删除 $s[l]$，从 $l-1$ 和 $r$ 开始，继续向左向右扩展。
- 选择删除 $s[r]$，从 $l$ 和 $r+1$ 开始，继续向左向右扩展。
- 两种情况取最大值。

### 答疑

**问**：如果在首次遇到 $s[l]\ne s[r]$ 之前，就执行删除操作，是否会得到更优的结果？

**答**：不会。如果提前删除的字母在回文中心的左边，我们可以对比删除 $s[l]$ 的情况，在最好情况下，都是删除了从 $l$ 到回文中心之间的一个字母，对于 $l$ 左侧的字母匹配，和原来是完全一样的，所以第二次中心扩展不会变得更长。如果删除的字母在回文中心的右边，我们可以对比删除 $s[r]$ 的情况，同理，第二次中心扩展不会变得更长。

[本题视频讲解](https://www.bilibili.com/video/BV1VgZ4BCETj/?t=4m38s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def almostPalindromic(self, s: str) -> int:
        n = len(s)
        ans = 0

        def expand(l: int, r: int) -> None:
            while l >= 0 and r < n and s[l] == s[r]:
                l -= 1
                r += 1
            nonlocal ans
            ans = max(ans, r - l - 1)  # [l+1, r-1] 是回文串

        for i in range(2 * n - 1):
            l, r = i // 2, (i + 1) // 2
            while l >= 0 and r < n and s[l] == s[r]:
                l -= 1
                r += 1
            expand(l - 1, r)  # 删除 s[l]，继续扩展
            expand(l, r + 1)  # 删除 s[r]，继续扩展
            if ans >= n:  # 优化：提前返回答案
                return n
        return ans
```

```java [sol-Java]
class Solution {
    public int almostPalindromic(String S) {
        char[] s = S.toCharArray();
        int n = s.length;
        int ans = 0;

        for (int i = 0; i < 2 * n - 1 && ans < n; i++) {
            int l = i / 2;
            int r = (i + 1) / 2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            ans = Math.max(ans, expand(s, l - 1, r)); // 删除 s[l]，继续扩展
            ans = Math.max(ans, expand(s, l, r + 1)); // 删除 s[r]，继续扩展
        }

        return Math.min(ans, n);
    }

    private int expand(char[] s, int l, int r) {
        while (l >= 0 && r < s.length && s[l] == s[r]) {
            l--;
            r++;
        }
        return r - l - 1; // [l+1, r-1] 是回文串
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int almostPalindromic(string s) {
        int n = s.size();
        int ans = 0;

        auto expand = [&](int l, int r) -> void {
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            ans = max(ans, r - l - 1); // [l+1, r-1] 是回文串
        };

        for (int i = 0; i < 2 * n - 1 && ans < n; i++) {
            int l = i / 2, r = (i + 1) / 2;
            while (l >= 0 && r < n && s[l] == s[r]) {
                l--;
                r++;
            }
            expand(l - 1, r); // 删除 s[l]，继续扩展
            expand(l, r + 1); // 删除 s[r]，继续扩展
        }
        return min(ans, n);
    }
};
```

```go [sol-Go]
func almostPalindromic(s string) (ans int) {
	n := len(s)
	expand := func(l, r int) {
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		ans = max(ans, r-l-1) // [l+1, r-1] 是回文串
	}

	for i := range 2*n - 1 {
		l, r := i/2, (i+1)/2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
		}
		expand(l-1, r) // 删除 s[l]，继续扩展
		expand(l, r+1) // 删除 s[r]，继续扩展
		if ans >= n { // 优化：提前返回答案
			return n
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n^2)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 方法二：Manacher 算法 + 后缀数组

用 Manacher 算法加速第一次中心扩展的过程。

第二次中心扩展的过程，等价于计算 $s$ 的某个后缀与 $\text{reverse}(s)$ 的某个后缀的最长公共前缀（LCP）。可以用后缀数组 $\mathcal{O}(n\log n)$ 预处理 $s + \texttt{#} + \text{reverse}(s)$ 后，$\mathcal{O}(1)$ 回答任意两个后缀的 LCP。

**注**：由于这个做法常数比较大，实际运行时间不如方法一。

```go
type sparseTable[T any] struct {
	st [][]T
	op func(T, T) T
}

// 时间复杂度 O(n * log n)
func newSparseTable[T any](a []T, op func(T, T) T) sparseTable[T] {
	n := len(a)
	w := bits.Len(uint(n))
	st := make([][]T, w)
	for i := range st {
		st[i] = make([]T, n)
	}
	copy(st[0], a)
	for i := 1; i < w; i++ {
		for j := range n - 1<<i + 1 {
			st[i][j] = op(st[i-1][j], st[i-1][j+1<<(i-1)])
		}
	}
	return sparseTable[T]{st, op}
}

// [l,r) 左闭右开，下标从 0 开始
// 返回 op(nums[l:r])
// 时间复杂度 O(1)
func (s sparseTable[T]) query(l, r int) T {
	k := bits.Len(uint(r-l)) - 1
	return s.op(s.st[k][l], s.st[k][r-1<<k])
}

func suffixArrayLCP(s string) func(int, int) int {
	// 后缀数组 sa（后缀序）
	// sa[i] 表示后缀字典序中的第 i 个字符串（的首字母）在 s 中的位置
	// 特别地，后缀 s[sa[0]:] 字典序最小，后缀 s[sa[n-1]:] 字典序最大
	type _tp struct {
		_  []byte
		sa []int32
	}
	sa := (*_tp)(unsafe.Pointer(suffixarray.New([]byte(s)))).sa

	// 计算后缀名次数组
	// 后缀 s[i:] 位于后缀字典序中的第 rank[i] 个
	// 特别地，rank[0] 即 s 在后缀字典序中的排名，rank[n-1] 即 s[n-1:] 在字典序中的排名
	// 相当于 sa 的反函数，即 rank[sa[i]] = i
	rank := make([]int, len(sa))
	for i, p := range sa {
		rank[p] = i
	}

	// 计算高度数组（也叫 LCP 数组）
	// height[0] = 0（哨兵）
	// height[i] = LCP(s[sa[i]:], s[sa[i-1]:])  (i > 0)
	// 获取 s[i] 所在位置的高度：height[rank[i]]
	height := make([]int, len(sa))
	h := 0
	// 计算 s 与 s[sa[rank[0]-1]:] 的 LCP（记作 LCP0）
	// 计算 s[1:] 与 s[sa[rank[1]-1]:] 的 LCP（记作 LCP1）
	// 计算 s[2:] 与 s[sa[rank[2]-1]:] 的 LCP
	// ...
	// 计算 s[n-1:] 与 s[sa[rank[n-1]-1]:] 的 LCP
	// 从 LCP0 到 LCP1，我们只去掉了 s[0] 和 s[sa[rank[0]-1]] 这两个字符
	// 所以 LCP1 >= LCP0 - 1
	// 这样就能加快 LCP 的计算了（类似滑动窗口）
	// 注：实际只计算了 n-1 对 LCP，因为我们跳过了 rank[i] = 0 的情况
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

	st := newSparseTable(height, func(a int, b int) int { return min(a, b) })

	// 返回 LCP(s[i:], s[j:])，即两个后缀的最长公共前缀
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
		// rj+1 是因为 query 是左闭右开
		return st.query(ri+1, rj+1)
	}
	return lcp
}

func almostPalindromic(s string) (ans int) {
	n := len(s)
	revS := []byte(s)
	slices.Reverse(revS)
	lcp := suffixArrayLCP(s + "#" + string(revS))

	// 将 s 改造为 t，这样就不需要分 len(s) 的奇偶来讨论了，因为新串 t 的每个回文子串都是奇回文串（都有回文中心）
	// s 和 t 的下标转换关系：
	// (si+1)*2 = ti
	// ti/2-1 = si
	// ti 为偶数（2,4,6,...）对应 s 中的奇回文串
	// ti 为奇数（3,5,7,...）对应 s 中的偶回文串
	t := append(make([]byte, 0, n*2+3), '^')
	for _, c := range s {
		t = append(t, '#', byte(c))
	}
	t = append(t, '#', '$')

	// 定义一个奇回文串的回文半径=(长度+1)/2，即保留回文中心，去掉一侧后的剩余字符串的长度
	// halfLen[i] 表示在 t 上的以 t[i] 为回文中心的最长回文子串的回文半径
	// 具体地，闭区间 [i-halfLen[i]+1, i+halfLen[i]-1] 是 t 上的一个回文子串
	// 由于 t 中回文子串的首尾字母一定是 #，根据下标转换关系，
	// 可以得到其在 s 中对应的回文子串的区间为 [(i-halfLen[i])/2, (i+halfLen[i])/2-2]
	halfLen := make([]int, len(t)-2)
	halfLen[1] = 1
	// boxR 表示当前右边界下标最大的回文子串的右边界下标+1（初始化成任意 <= 0 的数都可以）
	// boxM 为该最大回文子串的中心位置，二者的关系为 boxR = boxM + halfLen[boxM]
	boxM, boxR := 0, 0
	for i := 2; i < len(halfLen); i++ { // 循环的起止位置对应着原串的首尾字符
		hl := 1 // 注：如果题目比较的是抽象意义的值，单个值可能不满足要求，此时应初始化 hl = 0
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

		// 闭区间 [(i-halfLen[i])/2, (i+halfLen[i])/2-2] 是 s 上的一个回文子串
		l, r := (i-halfLen[i])/2, (i+halfLen[i])/2-2

		// s 本身是回文串，或者删除两端一个字母后是回文串
		if r-l+1 >= n-1 {
			return n // 如果 s 本身是回文串，删除回文中心后，仍然是回文串
		}

		// 删除 s[l-1]，继续扩展
		extra := 1 // 删除 [l,r] 外侧的一个字母
		if l-2 >= 0 && r+1 < n {
			extra += lcp(r+1, n*2-l+2) * 2
		}
		ans = max(ans, r-l+1+extra)

		// 删除 s[r+1]，继续扩展
		extra = 1 // 删除 [l,r] 外侧的一个字母
		if l-1 >= 0 && r+2 < n {
			extra += lcp(r+2, n*2-l+1) * 2
		}
		ans = max(ans, r-l+1+extra)
	}

	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log n)$。

## 专题训练

见下面字符串题单的「**三、Manacher 算法**」和「**八、后缀数组**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
