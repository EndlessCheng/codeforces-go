视频讲解：[【周赛 345】](https://www.bilibili.com/video/BV1ka4y137ua/)

按题意模拟，用 $\textit{vis}$ 数组标记接到过球的人。

为方便取模运算，循环中的下标可以从 $0$ 开始，在返回时再加一。

```py [sol1-Python3]
class Solution:
    def circularGameLosers(self, n: int, k: int) -> List[int]:
        vis = [False] * n
        i, d = 0, k
        while not vis[i]:
            vis[i] = True
            i = (i + d) % n
            d += k
        return [i for i, b in enumerate(vis, 1) if not b]
```

```java [sol1-Java]
class Solution {
    public int[] circularGameLosers(int n, int k) {
        var vis = new boolean[n];
        int m = n; // 答案长度
        for (int i = 0, d = k; !vis[i]; d += k, m--) {
            vis[i] = true;
            i = (i + d) % n;
        }
        var ans = new int[m];
        for (int i = 0, j = 0; i < n; i++)
            if (!vis[i])
                ans[j++] = i + 1;
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<int> circularGameLosers(int n, int k) {
        vector<int> vis(n);
        for (int i = 0, d = k; !vis[i]; d += k) {
            vis[i] = true;
            i = (i + d) % n;
        }
        vector<int> ans;
        for (int i = 0; i < n; i++)
            if (!vis[i])
                ans.push_back(i + 1);
        return ans;
    }
};
```

```go [sol1-Go]
func circularGameLosers(n int, k int) (ans []int) {
	vis := make([]bool, n)
	for i, d := 0, k; !vis[i]; d += k {
		vis[i] = true
		i = (i + d) % n
	}
	for i, b := range vis {
		if !b {
			ans = append(ans, i+1)
		}
	}
	return
}
```

```js [sol1-JavaScript]
var circularGameLosers = function (n, k) {
    let vis = Array(n).fill(false);
    for (let i = 0, d = k; !vis[i]; d += k) {
        vis[i] = true;
        i = (i + d) % n;
    }
    let ans = [];
    for (let i = 0; i < n; i++)
        if (!vis[i])
            ans.push(i + 1);
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。每个人至多用 $\textit{vis}$ 标记一次。
- 空间复杂度：$\mathcal{O}(n)$。
