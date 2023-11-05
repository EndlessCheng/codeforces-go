[视频讲解](https://www.bilibili.com/video/BV1Fc411R7xA/) 第二题。

对每个节点，判断它是否出现在 $\textit{edges}[i][1]$ 中。

如果恰好有一个节点没有出现，说明没有可以击败它的队伍，返回这个节点的编号。否则返回 $-1$。

```py [sol-Python3]
class Solution:
    def findChampion(self, n: int, edges: List[List[int]]) -> int:
        weak = [False] * n
        for _, y in edges:
            weak[y] = True
        ans = -1
        for i, w in enumerate(weak):
            if not w:
                if ans != -1:
                    return -1
                ans = i
        return ans
```

```java [sol-Java]
class Solution {
    public int findChampion(int n, int[][] edges) {
        boolean[] weak = new boolean[n];
        for (int[] e : edges) {
            weak[e[1]] = true;
        }
        int ans = -1;
        for (int i = 0; i < n; i++) {
            if (!weak[i]) {
                if (ans != -1) {
                    return -1;
                }
                ans = i;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findChampion(int n, vector<vector<int>> &edges) {
        vector<int> weak(n);
        int ans = -1;
        for (auto &e: edges) {
            weak[e[1]] = true;
        }
        for (int i = 0; i < n; i++) {
            if (!weak[i]) {
                if (ans != -1) {
                    return -1;
                }
                ans = i;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func findChampion(n int, edges [][]int) int {
	weak := make([]bool, n)
	for _, e := range edges {
		weak[e[1]] = true
	}
	ans := -1
	for i, w := range weak {
		if !w {
			if ans != -1 {
				return -1
			}
			ans = i
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $m$ 为 $\textit{edges}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
