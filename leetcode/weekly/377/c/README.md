建图，从 $\textit{original}[i]$ 向 $\textit{changed}[i]$ 连边，边权为 $\textit{cost}[i]$。

然后用 Floyd 算法求图中任意两点最短路，得到 $\textit{dis}$ 矩阵，原理请看 [带你发明 Floyd 算法！](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/)包含为什么循环顺序是 $kij$ 的讲解。

这里得到的 $\textit{dis}[i][j]$ 表示字母 $i$ 通过若干次替换操作变成字母 $j$ 的最小成本。

最后，累加所有 $\textit{dis}[\textit{source}[i]][\textit{target}[i]]$，即为答案。如果答案为无穷大，返回 $-1$。

[本题视频讲解](https://www.bilibili.com/video/BV1rG411k72D/)

```py [sol-Python3]
class Solution:
    def minimumCost(self, source: str, target: str, original: List[str], changed: List[str], cost: List[int]) -> int:
        dis = [[inf] * 26 for _ in range(26)]
        for i in range(26):
            dis[i][i] = 0

        for x, y, c in zip(original, changed, cost):
            x = ord(x) - ord('a')
            y = ord(y) - ord('a')
            dis[x][y] = min(dis[x][y], c)

        for k in range(26):
            for i in range(26):
                if dis[i][k] == inf:
                    continue  # 巨大优化！
                for j in range(26):
                    dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j])

        ans = sum(dis[ord(x) - ord('a')][ord(y) - ord('a')] for x, y in zip(source, target))
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public long minimumCost(String source, String target, char[] original, char[] changed, int[] cost) {
        final int INF = Integer.MAX_VALUE / 2;
        int[][] dis = new int[26][26];
        for (int i = 0; i < 26; i++) {
            Arrays.fill(dis[i], INF);
            dis[i][i] = 0;
        }
        for (int i = 0; i < cost.length; i++) {
            int x = original[i] - 'a';
            int y = changed[i] - 'a';
            dis[x][y] = Math.min(dis[x][y], cost[i]);
        }
        for (int k = 0; k < 26; k++) {
            for (int i = 0; i < 26; i++) {
                if (dis[i][k] == INF) {
                    continue; // 巨大优化！
                }
                for (int j = 0; j < 26; j++) {
                    dis[i][j] = Math.min(dis[i][j], dis[i][k] + dis[k][j]);
                }
            }
        }

        long ans = 0;
        for (int i = 0; i < source.length(); i++) {
            int d = dis[source.charAt(i) - 'a'][target.charAt(i) - 'a'];
            if (d == INF) {
                return -1;
            }
            ans += d;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minimumCost(string source, string target, vector<char>& original, vector<char>& changed, vector<int>& cost) {
        const int INF = 0x3f3f3f3f;
        int dis[26][26];
        memset(dis, 0x3f, sizeof(dis));
        for (int i = 0; i < 26; i++) {
            dis[i][i] = 0;
        }
        for (int i = 0; i < cost.size(); i++) {
            int x = original[i] - 'a';
            int y = changed[i] - 'a';
            dis[x][y] = min(dis[x][y], cost[i]);
        }
        for (int k = 0; k < 26; k++) {
            for (int i = 0; i < 26; i++) {
                if (dis[i][k] == INF) {
                    continue; // 巨大优化！
                }
                for (int j = 0; j < 26; j++) {
                    dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j]);
                }
            }
        }

        long long ans = 0;
        for (int i = 0; i < source.length(); i++) {
            int d = dis[source[i] - 'a'][target[i] - 'a'];
            if (d == INF) {
                return -1;
            }
            ans += d;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minimumCost(source, target string, original, changed []byte, cost []int) (ans int64) {
	const inf = math.MaxInt / 2
	dis := [26][26]int{}
	for i := range dis {
		for j := range dis[i] {
			if j != i {
				dis[i][j] = inf
			}
		}
	}
	for i, c := range cost {
		x := original[i] - 'a'
		y := changed[i] - 'a'
		dis[x][y] = min(dis[x][y], c)
	}
	for k := range dis {
		for i := range dis {
			if dis[i][k] == inf {
				continue // 巨大优化！
			}
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	for i, b := range source {
		d := dis[b-'a'][target[i]-'a']
		if d == inf {
			return -1
		}
		ans += int64(d)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m+|\Sigma|^3)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 为 $\textit{cost}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
