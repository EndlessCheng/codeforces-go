## 构造

由于每个字符串的长度都是 $2$，如果一个字母在公共超序列中出现 $2$ 次，那么把这个字母加到最左边和最右边，即可满足所有字符串。例如把字母 $\texttt{a}$ 加到最左边和最右边，那么所有 $\texttt{a*}$ 和 $\texttt{*a}$ 都能满足。这个性质对于多个出现 $2$ 次的字母也成立。

如果字母的出现次数超过 $2$，那么去掉中间的多余字母也能满足要求。所以**每个字母在公共超序列中的出现次数要么是 $1$，要么是 $2$**。

## 枚举子集

枚举出现 $2$ 次的字母（枚举子集），那么其余字母只出现 $1$ 次。

这里用二进制和位运算实现。用二进制压缩字母的出现次数，二进制从低到高第 $i$ 位是 $0$ 表示该字母（如果有）在公共超序列中出现 $1$ 次，是 $1$ 表示该字母在公共超序列中出现 $2$ 次。具体请看 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/circle/discuss/CaOJ45/)。

## 建图

下面只考虑两个字母都只出现 $1$ 次的字符串。

对于字符串 $s$，相当于一条约束：$s[0]$ 必须在 $s[1]$ 的左边。

如果有 $\texttt{ab},\texttt{bc},\texttt{ca}$ 这三个字符串，把「$\texttt{a}$ 必须在 $\texttt{b}$ 的左边」和「$\texttt{b}$ 必须在 $\texttt{c}$ 的左边」合并，得到「$\texttt{a}$ 必须在 $\texttt{c}$ 的左边」，但同时 $\texttt{ca}$ 意味着「$\texttt{c}$ 必须在 $\texttt{a}$ 的左边」。如果每个字母都只在公共超序列中出现 $1$ 次，这是无法做到的。换句话说，如果这些约束形成了一个**环状结构**，就不符合要求，因为我们规定这些字母只能出现 $1$ 次。

一般地，把「$s[0]$ 必须在 $s[1]$ 的左边」这个规则抽象成有向图，也就是连一条从 $s[0]$ 到 $s[1]$ 的有向边，那么问题变成：

- 这个有向图是否有环？

做法见 [207. 课程表](https://leetcode.cn/problems/course-schedule/)，可以用**三色标记法**或者拓扑排序解决，详见 [我的题解](https://leetcode.cn/problems/course-schedule/solutions/2992884/san-se-biao-ji-fa-pythonjavacgojsrust-by-pll7/)。

如果无环（也就是有向无环图），则说明我们当前枚举的子集是合法的，加入答案候选项中。

答案候选项中的大小最小的那些集合，即为最终答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV15sFNewEia/?t=34m17s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def supersequences(self, words: List[str]) -> List[List[int]]:
        # 收集有哪些字母，同时建图
        all_mask = 0
        g = defaultdict(list)
        for x, y in words:
            x, y = ord(x) - ord('a'), ord(y) - ord('a')
            all_mask |= 1 << x | 1 << y
            g[x].append(y)

        # 判断是否有环
        def has_cycle(sub: int) -> bool:
            color = [0] * 26
            def dfs(x: int) -> bool:
                color[x] = 1
                for y in g[x]:
                    # 只遍历不在 sub 中的字母
                    if sub >> y & 1:
                        continue
                    if color[y] == 1 or color[y] == 0 and dfs(y):
                        return True
                color[x] = 2
                return False
            for i, c in enumerate(color):
                # 只遍历不在 sub 中的字母
                if c == 0 and (sub >> i & 1) == 0 and dfs(i):
                    return True
            return False

        st = set()
        min_size = inf
        # 枚举 all_mask 的所有子集 sub
        sub = all_mask
        while True:
            size = sub.bit_count()
            # 剪枝：如果 size > min_size 就不需要判断了
            if size <= min_size and not has_cycle(sub):
                if size < min_size:
                    min_size = size
                    st.clear()
                st.add(sub)
            sub = (sub - 1) & all_mask
            if sub == all_mask:
                break

        return [[(all_mask >> i & 1) + (sub >> i & 1) for i in range(26)]
                for sub in st]
```

```java [sol-Java]
class Solution {
    public List<List<Integer>> supersequences(String[] words) {
        // 收集有哪些字母，同时建图
        int all = 0;
        List<Integer>[] g = new ArrayList[26];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (String s : words) {
            int x = s.charAt(0) - 'a';
            int y = s.charAt(1) - 'a';
            all |= 1 << x | 1 << y;
            g[x].add(y);
        }

        Set<Integer> set = new HashSet<>();
        int minSize = Integer.MAX_VALUE;
        // 枚举 all 的所有子集 sub
        int sub = all;
        do {
            int size = Integer.bitCount(sub);
            // 剪枝：如果 size > minSize 就不需要判断了
            if (size <= minSize && !hasCycle(sub, g)) {
                if (size < minSize) {
                    minSize = size;
                    set.clear();
                }
                set.add(sub);
            }
            sub = (sub - 1) & all;
        } while (sub != all);

        List<List<Integer>> ans = new ArrayList<>(set.size()); // 预分配空间
        for (int s : set) {
            List<Integer> cnt = new ArrayList<>(26); // 预分配空间
            for (int i = 0; i < 26; i++) {
                cnt.add((all >> i & 1) + (s >> i & 1));
            }
            ans.add(cnt);
        }
        return ans;
    }

    private boolean hasCycle(int sub, List<Integer>[] g) {
        int[] color = new int[26];
        for (int i = 0; i < 26; i++) {
            // 只遍历不在 sub 中的字母
            if (color[i] == 0 && (sub >> i & 1) == 0 && dfs(i, color, g, sub)) {
                return true;
            }
        }
        return false;
    }

    private boolean dfs(int x, int[] color, List<Integer>[] g, int sub) {
        color[x] = 1;
        for (int y : g[x]) {
            // 只遍历不在 sub 中的字母
            if ((sub >> y & 1) != 0) {
                continue;
            }
            if (color[y] == 1 || color[y] == 0 && dfs(y, color, g, sub)) {
                return true;
            }
        }
        color[x] = 2;
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> supersequences(vector<string>& words) {
        // 收集有哪些字母，同时建图
        int all = 0;
        vector<int> g[26]{};
        for (auto& s : words) {
            int x = s[0] - 'a', y = s[1] - 'a';
            all |= 1 << x | 1 << y;
            g[x].push_back(y);
        }

        // 判断是否有环
        auto has_cycle = [&](int sub) -> bool {
            int color[26]{};
            auto dfs = [&](this auto&& dfs, int x) -> bool {
                color[x] = 1;
                for (int y : g[x]) {
                    // 只遍历不在 sub 中的字母
                    if (sub >> y & 1) {
                        continue;
                    }
                    if (color[y] == 1 || color[y] == 0 && dfs(y)) {
                        return true;
                    }
                }
                color[x] = 2;
                return false;
            };
            for (int i = 0; i < 26; i++) {
                // 只遍历不在 sub 中的字母
                if (color[i] == 0 && (sub >> i & 1) == 0 && dfs(i)) {
                    return true;
                }
            }
            return false;
        };

        unordered_set<int> st;
        int min_size = INT_MAX;
        // 枚举 all 的所有子集 sub
        int sub = all;
        do {
            int size = popcount((unsigned) sub);
            // 剪枝：如果 size > min_size 就不需要判断了
            if (size <= min_size && !has_cycle(sub)) {
                if (size < min_size) {
                    min_size = size;
                    st.clear();
                }
                st.insert(sub);
            }
            sub = (sub - 1) & all;
        } while (sub != all);

        vector<vector<int>> ans;
        for (int sub : st) {
            vector<int> cnt(26);
            for (int i = 0; i < 26; i++) {
                cnt[i] = (all >> i & 1) + (sub >> i & 1);
            }
            ans.push_back(cnt);
        }
        return ans;
    }
};
```

```go [sol-Go]
func supersequences(words []string) [][]int {
	// 收集有哪些字母，同时建图
	all := 0
	g := [26][]int{}
	for _, s := range words {
		x, y := int(s[0]-'a'), int(s[1]-'a')
		all |= 1<<x | 1<<y
		g[x] = append(g[x], y)
	}

	// 判断是否有环
	hasCycle := func(sub int) bool {
		color := [26]int8{}
		var dfs func(int) bool
		dfs = func(x int) bool {
			color[x] = 1
			for _, y := range g[x] {
				// 只遍历不在 sub 中的字母
				if sub>>y&1 > 0 {
					continue
				}
				if color[y] == 1 || color[y] == 0 && dfs(y) {
					return true
				}
			}
			color[x] = 2
			return false
		}
		for i, c := range color {
			// 只遍历不在 sub 中的字母
			if c == 0 && sub>>i&1 == 0 && dfs(i) {
				return true
			}
		}
		return false
	}

	set := map[int]struct{}{}
	minSize := math.MaxInt
	// 枚举 all 的所有子集 sub
	for sub, ok := all, true; ok; ok = sub != all {
		size := bits.OnesCount(uint(sub))
		// 剪枝：如果 size > minSize 就不需要判断了
		if size <= minSize && !hasCycle(sub) {
			if size < minSize {
				minSize = size
				clear(set)
			}
			set[sub] = struct{}{}
		}
		sub = (sub - 1) & all
	}

	ans := make([][]int, 0, len(set)) // 预分配空间
	for sub := range set {
		cnt := make([]int, 26)
		for i := range cnt {
			cnt[i] = all>>i&1 + sub>>i&1
		}
		ans = append(ans, cnt)
	}
	return ans
}
```

## 优化

如果 $s[0]=s[1]$，那么这个字母一定要出现两次，不需要枚举其只出现一次的情况。

为方便计算，这里改成在子集中的字母出现一次，不在子集中的字母出现两次。

```py [sol-Python3]
class Solution:
    def supersequences(self, words: List[str]) -> List[List[int]]:
        # 收集有哪些字母，同时建图
        all_mask = mask2 = 0
        g = defaultdict(list)
        for x, y in words:
            x, y = ord(x) - ord('a'), ord(y) - ord('a')
            all_mask |= 1 << x | 1 << y
            if x == y:
                mask2 |= 1 << x
            g[x].append(y)

        # 判断是否有环
        def has_cycle(sub: int) -> bool:
            color = [0] * 26

            def dfs(x: int) -> bool:
                color[x] = 1
                for y in g[x]:
                    # 只遍历在 sub 中的字母
                    if (sub >> y & 1) == 0:
                        continue
                    if color[y] == 1 or color[y] == 0 and dfs(y):
                        return True
                color[x] = 2
                return False

            for i, c in enumerate(color):
                # 只遍历在 sub 中的字母
                if c == 0 and sub >> i & 1 and dfs(i):
                    return True
            return False

        st = set()
        max_size = 0
        # 枚举 mask1 的所有子集 sub
        sub = mask1 = all_mask ^ mask2
        while True:
            size = sub.bit_count()
            # 剪枝：如果 size < max_size 就不需要判断了
            if size >= max_size and not has_cycle(sub):
                if size > max_size:
                    max_size = size
                    st.clear()
                st.add(sub)
            sub = (sub - 1) & mask1
            if sub == mask1:
                break

        return [[(all_mask >> i & 1) + ((all_mask ^ sub) >> i & 1) for i in range(26)]
                for sub in st]
```

```java [sol-Java]
class Solution {
    public List<List<Integer>> supersequences(String[] words) {
        // 收集有哪些字母，同时建图
        int all = 0;
        int mask2 = 0;
        List<Integer>[] g = new ArrayList[26];
        Arrays.setAll(g, i -> new ArrayList<>());
        for (String s : words) {
            int x = s.charAt(0) - 'a';
            int y = s.charAt(1) - 'a';
            all |= 1 << x | 1 << y;
            if (x == y) {
                mask2 |= 1 << x;
            }
            g[x].add(y);
        }

        Set<Integer> set = new HashSet<>();
        int maxSize = 0;
        // 枚举 mask1 的所有子集 sub
        int mask1 = all ^ mask2;
        int sub = mask1;
        do {
            int size = Integer.bitCount(sub);
            // 剪枝：如果 size < maxSize 就不需要判断了
            if (size >= maxSize && !hasCycle(sub, g)) {
                if (size > maxSize) {
                    maxSize = size;
                    set.clear();
                }
                set.add(sub);
            }
            sub = (sub - 1) & mask1;
        } while (sub != mask1);

        List<List<Integer>> ans = new ArrayList<>(set.size()); // 预分配空间
        for (int s : set) {
            List<Integer> cnt = new ArrayList<>(26); // 预分配空间
            for (int i = 0; i < 26; i++) {
                cnt.add((all >> i & 1) + ((all ^ s) >> i & 1));
            }
            ans.add(cnt);
        }
        return ans;
    }

    private boolean hasCycle(int sub, List<Integer>[] g) {
        int[] color = new int[26];
        for (int i = 0; i < 26; i++) {
            // 只遍历在 sub 中的字母
            if (color[i] == 0 && (sub >> i & 1) != 0 && dfs(i, color, g, sub)) {
                return true;
            }
        }
        return false;
    }

    private boolean dfs(int x, int[] color, List<Integer>[] g, int sub) {
        color[x] = 1;
        for (int y : g[x]) {
            // 只遍历在 sub 中的字母
            if ((sub >> y & 1) == 0) {
                continue;
            }
            if (color[y] == 1 || color[y] == 0 && dfs(y, color, g, sub)) {
                return true;
            }
        }
        color[x] = 2;
        return false;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> supersequences(vector<string>& words) {
        // 收集有哪些字母，同时建图
        int all = 0, mask2 = 0;
        vector<int> g[26]{};
        for (auto& s : words) {
            int x = s[0] - 'a', y = s[1] - 'a';
            all |= 1 << x | 1 << y;
            if (x == y) {
                mask2 |= 1 << x;
            }
            g[x].push_back(y);
        }

        // 判断是否有环
        auto has_cycle = [&](int sub) -> bool {
            int color[26]{};
            auto dfs = [&](this auto&& dfs, int x) -> bool {
                color[x] = 1;
                for (int y : g[x]) {
                    // 只遍历在 sub 中的字母
                    if ((sub >> y & 1) == 0) {
                        continue;
                    }
                    if (color[y] == 1 || color[y] == 0 && dfs(y)) {
                        return true;
                    }
                }
                color[x] = 2;
                return false;
            };
            for (int i = 0; i < 26; i++) {
                // 只遍历在 sub 中的字母
                if (color[i] == 0 && sub >> i & 1 && dfs(i)) {
                    return true;
                }
            }
            return false;
        };

        unordered_set<int> st;
        int max_size = 0;
        // 枚举 mask1 的所有子集 sub
        int mask1 = all ^ mask2;
        int sub = mask1;
        do {
            int size = popcount((unsigned) sub);
            // 剪枝：如果 size < min_size 就不需要判断了
            if (size >= max_size && !has_cycle(sub)) {
                if (size > max_size) {
                    max_size = size;
                    st.clear();
                }
                st.insert(sub);
            }
            sub = (sub - 1) & mask1;
        } while (sub != mask1);

        vector<vector<int>> ans;
        for (int sub : st) {
            vector<int> cnt(26);
            for (int i = 0; i < 26; i++) {
                cnt[i] = (all >> i & 1) + ((all ^ sub) >> i & 1);
            }
            ans.push_back(cnt);
        }
        return ans;
    }
};
```

```go [sol-Go]
func supersequences(words []string) [][]int {
	// 收集有哪些字母，同时建图
	all, mask2 := 0, 0
	g := [26][]int{}
	for _, s := range words {
		x, y := int(s[0]-'a'), int(s[1]-'a')
		all |= 1<<x | 1<<y
		if x == y {
			mask2 |= 1 << x
		}
		g[x] = append(g[x], y)
	}

	// 判断是否有环
	hasCycle := func(sub int) bool {
		color := [26]int8{}
		var dfs func(int) bool
		dfs = func(x int) bool {
			color[x] = 1
			for _, y := range g[x] {
				// 只遍历在 sub 中的字母
				if sub>>y&1 == 0 {
					continue
				}
				if color[y] == 1 || color[y] == 0 && dfs(y) {
					return true
				}
			}
			color[x] = 2
			return false
		}
		for i, c := range color {
			// 只遍历在 sub 中的字母
			if c == 0 && sub>>i&1 > 0 && dfs(i) {
				return true
			}
		}
		return false
	}

	set := map[int]struct{}{}
	maxSize := 0
	mask1 := all ^ mask2
	// 枚举 mask1 的所有子集 sub
	for sub, ok := mask1, true; ok; ok = sub != mask1 {
		size := bits.OnesCount(uint(sub))
		// 剪枝：如果 size < maxSize 就不需要判断了
		if size >= maxSize && !hasCycle(sub) {
			if size > maxSize {
				maxSize = size
				clear(set)
			}
			set[sub] = struct{}{}
		}
		sub = (sub - 1) & mask1
	}

	ans := make([][]int, 0, len(set)) // 预分配空间
	for sub := range set {
		cnt := make([]int, 26)
		for i := range cnt {
			cnt[i] = all>>i&1 + (all^sub)>>i&1
		}
		ans = append(ans, cnt)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(m2^k)$，其中 $m$ 是 $\textit{words}$ 的长度，$k\le 16$ 是字母种类数。枚举 $2^k$ 个子集，每个子集需要 $\mathcal{O}(m)$ 的时间跑 DFS。
- 空间复杂度：$\mathcal{O}(2^k)$。返回值不计入。由于 $m\le k^2 \ll 2^k$，所以空间复杂度以 $2^k$ 为主。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
