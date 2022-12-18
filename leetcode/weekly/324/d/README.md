[视频讲解](https://www.bilibili.com/video/BV1LW4y1T7if/) 已出炉，欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

挨个回答每个询问。

环可以看成是从 $a$ 出发往上走，在某个位置「拐弯」，往下走到 $b$。

这个拐弯的地方就是 $a$ 和 $b$ 的**最近公共祖先**。

设 $\textit{LCA}$ 为 $a$ 和 $b$ 的最近公共祖先，那么环长等于 $\textit{LCA}$ 到 $a$ 的距离加 $\textit{LCA}$ 到 $b$ 的距离加一。

如何找 $\textit{LCA}$？

不断循环，每次循环比较 $a$ 和 $b$ 的大小：

- 如果 $a>b$，则 $a$ 的深度大于等于 $b$ 的深度，那么把 $a$ 移动到其父节点，即 $a=a/2$；
- 如果 $a<b$，则 $a$ 的深度小于等于 $b$ 的深度，那么把 $b$ 移动到其父节点，即 $b=b/2$；
- 如果 $a=b$，则找到了 $\textit{LCA}$，退出循环。

循环次数加一即为环长。

```py [sol1-Python3]
class Solution:
    def cycleLengthQueries(self, n: int, queries: List[List[int]]) -> List[int]:
        for i, (a, b) in enumerate(queries):
            res = 1
            while a != b:
                if a > b: a //= 2
                else: b //= 2
                res += 1
            queries[i] = res
        return queries
```

```java [sol1-Java]
class Solution {
    public int[] cycleLengthQueries(int n, int[][] queries) {
        var m = queries.length;
        var ans = new int[m];
        for (var i = 0; i < m; ++i) {
            int res = 1, a = queries[i][0], b = queries[i][1];
            while (a != b) {
                if (a > b) a /= 2;
                else b /= 2;
                ++res;
            }
            ans[i] = res;
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> cycleLengthQueries(int n, vector<vector<int>> &queries) {
        int m = queries.size();
        vector<int> ans(m);
        for (int i = 0; i < m; ++i) {
            int res = 1, a = queries[i][0], b = queries[i][1];
            while (a != b) {
                a > b ? a /= 2 : b /= 2;
                ++res;
            }
            ans[i] = res;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func cycleLengthQueries(_ int, queries [][]int) []int {
	ans := make([]int, len(queries))
	for i, q := range queries {
		res := 1
		for a, b := q[0], q[1]; a != b; res++ {
			if a > b {
				a /= 2
			} else {
				b /= 2
			}
		}
		ans[i] = res
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$O(nm)$，其中 $m$ 为 $\textit{queries}$ 的长度。回答一个询问的时间复杂度为 $O(n)$。
- 空间复杂度：$O(1)$，仅用到若干额外变量。
