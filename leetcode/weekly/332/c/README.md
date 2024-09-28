把等式 $\textit{val} \oplus \textit{first} = \textit{second}$ 两边同时异或 $\textit{first}$，得到

$$
\textit{val} \oplus \textit{first} \oplus \textit{first} = \textit{second} \oplus \textit{first}
$$

由于 $\textit{first} \oplus \textit{first} = 0$，因此上式化简为

$$
\textit{val} = \textit{second}\oplus \textit{first}
$$

所以问题等价于在 $s$ 中找到值为 $\textit{second}\oplus \textit{first}$ 的数。

由于 $10^9<2^{30}$，可以直接预处理 $s$ 中长度不超过 $30$ 的子串（转成数字）及其对应的 $\textit{left}$ 和 $\textit{right}$，记到一个哈希表中，然后 $\mathcal{O}(1)$ 地回答询问。

特别地，如果存在 $s[i]=0$，我们可以单独记录 $0$ 对应的 $\textit{left}=i,\ \textit{right}=i$，记录到哈希表中，这样预处理 $s$ 子串的时候，就只需要考虑从 $1$ 开始的子串，效率更高。

本题 [视频讲解](https://www.bilibili.com/video/BV1GY411i7RP/)。

```py [sol-Python3]
class Solution:
    def substringXorQueries(self, s: str, queries: List[List[int]]) -> List[List[int]]:
        n, m = len(s), {}
        if (i := s.find('0')) >= 0:
            m[0] = (i, i)

        for l, c in enumerate(s):
            if c == '0':
                continue
            x = 0
            for r in range(l, min(l + 30, n)):
                x = (x << 1) | (ord(s[r]) & 1)
                if x not in m:
                    m[x] = (l, r)

        NOT_FOUND = (-1, -1)
        return [m.get(x ^ y, NOT_FOUND) for x, y in queries]
```

```java [sol-Java]
class Solution {
    private static final int[] NOT_FOUND = new int[]{-1, -1};

    public int[][] substringXorQueries(String S, int[][] queries) {
        Map<Integer, int[]> m = new HashMap<>();
        int i = S.indexOf('0');
        if (i >= 0) {
            m.put(0, new int[]{i, i}); // 这样下面就可以直接跳过 '0' 了，效率更高
        }

        char[] s = S.toCharArray();
        int n = s.length;
        for (int l = 0; l < n; l++) {
            if (s[l] == '0') {
                continue;
            }
            for (int r = l, x = 0; r < Math.min(l + 30, n); r++) {
                x = (x << 1) | (s[r] & 1);
                m.putIfAbsent(x, new int[]{l, r});
            }
        }

        int[][] ans = new int[queries.length][];
        for (i = 0; i < queries.length; i++) {
            ans[i] = m.getOrDefault(queries[i][0] ^ queries[i][1], NOT_FOUND);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<vector<int>> substringXorQueries(string s, vector<vector<int>>& queries) {
        unordered_map<int, pair<int, int>> m;
        if (auto i = s.find('0'); i != string::npos) {
            m[0] = {i, i};
        }
        for (int l = 0, n = s.length(); l < n; l++) {
            if (s[l] == '0') {
                continue;
            }
            for (int r = l, x = 0; r < min(l + 30, n); r++) {
                x = (x << 1) | (s[r] & 1);
                if (!m.contains(x)) {
                    m[x] = {l, r};
                }
            }
        }

        vector<vector<int>> ans;
        for (auto& q : queries) {
            auto it = m.find(q[0] ^ q[1]);
            if (it == m.end()) {
                ans.push_back({-1, -1});
            } else {
                ans.push_back({it->second.first, it->second.second});
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func substringXorQueries(s string, queries [][]int) [][]int {
	type pair struct{ l, r int }
	m := map[int]pair{}
	if i := strings.IndexByte(s, '0'); i >= 0 {
		m[0] = pair{i, i}
	}

	for l, c := range s {
		if c == '0' {
			continue
		}
		for r, x := l, 0; r < l+30 && r < len(s); r++ {
			x = x<<1 | int(s[r]&1)
			if _, ok := m[x]; !ok {
				m[x] = pair{l, r}
			}
		}
	}

	ans := make([][]int, len(queries))
	notFound := []int{-1, -1} // 避免重复创建
	for i, q := range queries {
		p, ok := m[q[0]^q[1]]
		if !ok {
			ans[i] = notFound
		} else {
			ans[i] = []int{p.l, p.r}
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log U + q)$，其中 $n$ 为 $s$ 的长度，$U=max(\textit{queries})$，$q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n\log U)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
