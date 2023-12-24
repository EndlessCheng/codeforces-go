[本题视频讲解](https://www.bilibili.com/video/BV1rG411k72D/)

建图，从 $\textit{original}[i]$ 向 $\textit{changed}[i]$ 连边，边权为 $\textit{cost}[i]$。

然后用 Floyd 算法求图中任意两点最短路，得到 $\textit{dis}$ 矩阵，原理请看 [带你发明 Floyd 算法！](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/solution/dai-ni-fa-ming-floyd-suan-fa-cong-ji-yi-m8s51/)包含为什么循环顺序是 $kij$ 的讲解。

这里得到的 $\textit{dis}[i][j]$ 表示字母 $i$ 通过若干次替换操作变成字母 $j$ 的最小成本。

最后累加所有 $\textit{dis}[\textit{original}[i]][\textit{changed}[i]]$，即为答案。如果答案为无穷大，返回 $-1$。

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
                for j in range(26):
                    dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j])

        ans = sum(dis[ord(x) - ord('a')][ord(y) - ord('a')] for x, y in zip(source, target))
        return ans if ans < inf else -1
```

```java [sol-Java]
class Solution {
    public long minimumCost(String source, String target, char[] original, char[] changed, int[] cost) {
        int[][] dis = new int[26][26];
        for (int i = 0; i < 26; i++) {
            Arrays.fill(dis[i], Integer.MAX_VALUE / 2);
            dis[i][i] = 0;
        }
        for (int i = 0; i < cost.length; i++) {
            int x = original[i] - 'a';
            int y = changed[i] - 'a';
            dis[x][y] = Math.min(dis[x][y], cost[i]);
        }
        for (int k = 0; k < 26; k++) {
            for (int i = 0; i < 26; i++) {
                for (int j = 0; j < 26; j++) {
                    dis[i][j] = Math.min(dis[i][j], dis[i][k] + dis[k][j]);
                }
            }
        }

        long ans = 0;
        for (int i = 0; i < source.length(); i++) {
            int d = dis[source.charAt(i) - 'a'][target.charAt(i) - 'a'];
            if (d == Integer.MAX_VALUE / 2) {
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
    long long minimumCost(string source, string target, vector<char> &original, vector<char> &changed, vector<int> &cost) {
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
                for (int j = 0; j < 26; j++) {
                    dis[i][j] = min(dis[i][j], dis[i][k] + dis[k][j]);
                }
            }
        }

        long long ans = 0;
        for (int i = 0; i < source.length(); i++) {
            int d = dis[source[i] - 'a'][target[i] - 'a'];
            if (d == 0x3f3f3f3f) {
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
	dis := [26][26]int{}
	for i := range dis {
		for j := range dis[i] {
			if j != i {
				dis[i][j] = 1e13
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
			for j := range dis {
				dis[i][j] = min(dis[i][j], dis[i][k]+dis[k][j])
			}
		}
	}

	for i, b := range source {
		ans += int64(dis[b-'a'][target[i]-'a'])
	}
	if ans >= 1e13 {
		return -1
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m+|\Sigma|^3)$，其中 $n$ 为 $\textit{source}$ 的长度，$m$ 为 $\textit{cost}$ 的长度，$|\Sigma|$ 为字符集合的大小，本题中字符均为小写字母，所以 $|\Sigma|=26$。
- 空间复杂度：$\mathcal{O}(|\Sigma|^2)$。

#### 相似题目

- [2642. 设计可以求最短路径的图类](https://leetcode.cn/problems/design-graph-with-shortest-path-calculator/) 1811
- [1334. 阈值距离内邻居最少的城市](https://leetcode.cn/problems/find-the-city-with-the-smallest-number-of-neighbors-at-a-threshold-distance/) 1855
- [2101. 引爆最多的炸弹](https://leetcode.cn/problems/detonate-the-maximum-bombs/) 1880
