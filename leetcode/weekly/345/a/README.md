下午两点[【biIibiIi@灵茶山艾府】](https://space.bilibili.com/206214)直播讲题，记得关注哦~

---

按题意模拟。

为方便计算，循环中的下标可以从 $0$ 开始，在返回时再加一。

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
        int m = n;
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

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。每个朋友至多遍历一次。
- 空间复杂度：$\mathcal{O}(n)$。
