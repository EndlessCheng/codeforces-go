### 本题视频讲解

见[【周赛 344】](https://www.bilibili.com/video/BV1YL41187Rx/)第三题，欢迎点赞投币！

### 思路

由于每次修改只会影响当前元素与其左右元素的关系，所以模拟即可。

用 $\textit{cnt}$ 统计相邻相同的个数。

代码实现时，可以先去掉当前元素对 $\textit{cnt}$ 的影响，修改颜色后，再加上当前元素对 $\textit{cnt}$ 的影响。

```py [sol1-Python3]
class Solution:
    def colorTheArray(self, n: int, queries: List[List[int]]) -> List[int]:
        ans = []
        a, cnt = [0] * (n + 2), 0  # 避免讨论下标出界的情况
        for i, c in queries:
            i += 1  # 下标改成从 1 开始
            if a[i]: cnt -= (a[i] == a[i - 1]) + (a[i] == a[i + 1])
            a[i] = c
            cnt += (a[i] == a[i - 1]) + (a[i] == a[i + 1])
            ans.append(cnt)
        return ans
```

```java [sol1-Java]
class Solution {
    public int[] colorTheArray(int n, int[][] queries) {
        int q = queries.length, cnt = 0;
        int[] ans = new int[q], a = new int[n + 2]; // 避免讨论下标出界的情况
        for (int qi = 0; qi < q; qi++) {
            int i = queries[qi][0] + 1, c = queries[qi][1]; // 下标改成从 1 开始
            if (a[i] > 0)
                cnt -= (a[i] == a[i - 1] ? 1 : 0) + (a[i] == a[i + 1] ? 1 : 0);
            a[i] = c;
            cnt += (a[i] == a[i - 1] ? 1 : 0) + (a[i] == a[i + 1] ? 1 : 0);
            ans[qi] = cnt;
        }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> colorTheArray(int n, vector<vector<int>> &queries) {
        int q = queries.size(), cnt = 0;
        vector<int> ans(q), a(n + 2); // 避免讨论下标出界的情况
        for (int qi = 0; qi < q; qi++) {
            int i = queries[qi][0] + 1, c = queries[qi][1]; // 下标改成从 1 开始
            if (a[i]) cnt -= (a[i] == a[i - 1]) + (a[i] == a[i + 1]);
            a[i] = c;
            cnt += (a[i] == a[i - 1]) + (a[i] == a[i + 1]);
            ans[qi] = cnt;
        }
        return ans;
    }
};
```

```go [sol1-Go]
func colorTheArray(n int, queries [][]int) []int {
	ans := make([]int, len(queries))
	a := make([]int, n+2) // 避免讨论下标出界的情况
	cnt := 0
	for qi, q := range queries {
		i, c := q[0]+1, q[1] // 下标改成从 1 开始
		if a[i] > 0 {
			if a[i] == a[i-1] {
				cnt--
			}
			if a[i] == a[i+1] {
				cnt--
			}
		}
		a[i] = c
		if a[i] == a[i-1] {
			cnt++
		}
		if a[i] == a[i+1] {
			cnt++
		}
		ans[qi] = cnt
	}
	return ans
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+q)$，其中 $q$ 为 $\textit{queries}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

### 思考题

如果求的是最长连续同色长度要怎么做？
