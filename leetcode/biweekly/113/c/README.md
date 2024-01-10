[视频讲解](https://www.bilibili.com/video/BV1PV411N76R/) 第三题。

突破口：$0\le k\le 100$。

由于异或和都是非负数，所以等式 $(x_1\oplus x_2) + (y_1\oplus y_2) = k$ 暗含着：

- $0\le x_1\oplus x_2 \le k$
- $0\le y_1\oplus y_2 \le k$

如果 $x_1\oplus x_2 = i$，我们可以得到 $y_1\oplus y_2 = k-i$。

异或运算类似加法，我们可以把 $x_2$ 和 $y_2$ 移到等式右侧。 

枚举 $(x_2,y_2)$，则 $x_1 = x_2\oplus i$ 且 $y_1 =y_2\oplus (k-i)$。

用哈希表统计遍历过的点的个数，把哈希表中的 $(x_2\oplus i, y_2\oplus (k-i))$ 的个数加入答案。

如果不方便在哈希表中插入 `pair`，可以把 $(x,y)$ 压缩成一个整数，例如 $2000000x + y$。

```py [sol-Python3]
class Solution:
    def countPairs(self, coordinates: List[List[int]], k: int) -> int:
        ans = 0
        cnt = Counter()
        for x, y in coordinates:
            for i in range(k + 1):
                ans += cnt[x ^ i, y ^ (k - i)]  # tuple 的括号可以省略
            cnt[x, y] += 1  # tuple 的括号可以省略
        return ans
```

```java [sol-Java]
class Solution {
    public int countPairs(List<List<Integer>> coordinates, int k) {
        int ans = 0;
        var cnt = new HashMap<Long, Integer>();
        for (var p : coordinates) {
            int x = p.get(0), y = p.get(1);
            for (int i = 0; i <= k; i++) {
                ans += cnt.getOrDefault((x ^ i) * 2000000L + (y ^ (k - i)), 0);
            }
            cnt.merge(x * 2000000L + y, 1, Integer::sum);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPairs(vector<vector<int>> &coordinates, int k) {
        int ans = 0;
        unordered_map<long long, int> cnt;
        for (auto &p: coordinates) {
            for (int i = 0; i <= k; i++) {
                // 直接 ans += cnt[...] 会插入不存在的点
                auto it = cnt.find((p[0] ^ i) * 2000000LL + (p[1] ^ (k - i)));
                if (it != cnt.end()) {
                    ans += it->second;
                }
            }
            cnt[p[0] * 2000000LL + p[1]]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countPairs(coordinates [][]int, k int) (ans int) {
	type pair struct{ x, y int }
	cnt := map[pair]int{}
	for _, p := range coordinates {
		x, y := p[0], p[1]
		for i := 0; i <= k; i++ {
			ans += cnt[pair{x ^ i, y ^ (k - i)}]
		}
		cnt[pair{x, y}]++
	}
	return
}
```

```js [sol-JavaScript]
var countPairs = function (coordinates, k) {
    let ans = 0;
    const cnt = new Map();
    for (const [x, y] of coordinates) {
        for (let i = 0; i <= k; i++) {
            ans += cnt.get((x ^ i) * 2000000 + (y ^ (k - i))) ?? 0;
        }
        const key = x * 2000000 + y;
        cnt.set(key, (cnt.get(key) ?? 0) + 1);
    }
    return ans;
};
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(nk)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。
