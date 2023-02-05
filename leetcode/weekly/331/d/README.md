首先，把两个数组中都有的数去掉，那么每个剩余数字的出现次数必须为偶数。这可以用哈希表来统计。

设处理后的剩余数组分别 $a$ 和 $b$。

贪心地想，如果要交换 $a$ 中最小的数，那么找一个 $b$ 中最大的数是最合适的；对于 $b$ 中最小的数也同理。

那么把 $a$ **从小到大**排序，$b$ **从大到小**排序，两两匹配。

但是，还有一种方案。

把 $\textit{basket}_1$ 和 $\textit{basket}_2$ 中的最小值 $\textit{mn}$ 当作「工具人」，对于 $a[i]$ 和 $b[i]$ 的交换，可以分别和 $\textit{mn}$ 交换一次，就相当于 $a[i]$ 和 $b[i]$ 交换了。

因此每次交换的代价为

$$
\min(a[i], b[i], 2\cdot\textit{mn})
$$

累加代价，即为答案。

上式也表明，如果工具人也在需要交换的数字中，那么它的最小代价必然是和其他数交换，不会发生工具人和工具人交换的情况。

设 $m$ 为 $a$ 的长度。代码实现时，由于 $\min(a[i], b[i])$ 的数字都在 $a$ 的某个前缀与 $b$ 某个后缀中，而剩下没有选的数（$a$ 的后缀和 $b$ 的前缀）不比这 $m$ 个数小，所以取出的数一定是这 $2m$ 个数中最小的 m 个数。

> 更详细的证明：设选了 $a[0],\cdots,a[i]$ 和 $b[i+1],\cdots,b[m-1]$，由于 $a[i+1]\ge b[i+1]$ 且 $a[i+1] \ge a[i]$，所以 $a[i+1]$ 大于等于任意已选数字，进而推出 $a[i+2],\cdots,a[m-1]$ 都是大于等于任意已选数字的；对于 $b[i]$ 和 $b[i-1],\cdots,b[0]$ 同理。所以剩下没有选的数字都比已选数字大，进而说明已选数字是这 $2m$ 个数中最小的 $m$ 个数。

那么可以直接把 $a$ 和 $b$ 拼起来，从小到大排序后，遍历前一半的数即可（排序可以用快速选择代替，见 C++）。

附：[视频讲解](https://www.bilibili.com/video/BV1sG4y1T7oc/)

```py [sol1-Python3]
class Solution:
    def minCost(self, basket1: List[int], basket2: List[int]) -> int:
        cnt = Counter()
        for x, y in zip(basket1, basket2):
            cnt[x] += 1
            cnt[y] -= 1
        mn = min(cnt)
        a = []
        for x, c in cnt.items():
            if c % 2: return -1
            a.extend([x] * (abs(c) // 2))
        a.sort()  # 也可以用快速选择
        return sum(min(x, mn * 2) for x in a[:len(a) // 2])
```

```java [sol1-Java]
class Solution {
    public long minCost(int[] basket1, int[] basket2) {
        var cnt = new HashMap<Integer, Integer>();
        for (int i = 0; i < basket1.length; ++i) {
            cnt.merge(basket1[i], 1, Integer::sum);
            cnt.merge(basket2[i], -1, Integer::sum);
        }

        int mn = Integer.MAX_VALUE;
        var a = new ArrayList<Integer>();
        for (var e : cnt.entrySet()) {
            int x = e.getKey(), c = e.getValue();
            if (c % 2 != 0) return -1;
            mn = Math.min(mn, x);
            for (c = Math.abs(c) / 2; c > 0; --c)
                a.add(x);
        }

        long ans = 0;
        a.sort((x, y) -> x - y); // 也可以用快速选择
        for (int i = 0; i < a.size() / 2; ++i)
            ans += Math.min(a.get(i), mn * 2);
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long minCost(vector<int> &basket1, vector<int> &basket2) {
        unordered_map<int, int> cnt;
        for (int i = 0; i < basket1.size(); ++i) {
            ++cnt[basket1[i]];
            --cnt[basket2[i]];
        }

        int mn = INT_MAX;
        vector<int> a;
        for (auto [x, c] : cnt) {
            if (c % 2) return -1;
            mn = min(mn, x);
            for (c = abs(c) / 2; c > 0; --c)
                a.push_back(x);
        }

        long ans = 0;
        nth_element(a.begin(), a.begin() + a.size() / 2, a.end()); // 快速选择
        for (int i = 0; i < a.size() / 2; ++i)
            ans += min(a[i], mn * 2);
        return ans;
    }
};
```

```go [sol1-Go]
func minCost(basket1, basket2 []int) (ans int64) {
	cnt := map[int]int{}
	for i, x := range basket1 {
		cnt[x]++
		cnt[basket2[i]]--
	}

	mn, a := math.MaxInt, []int{}
	for x, c := range cnt {
		if c%2 != 0 {
			return -1
		}
		mn = min(mn, x)
		for c = abs(c) / 2; c > 0; c-- {
			a = append(a, x)
		}
	}

	sort.Ints(a) // 也可以用快速选择
	for _, x := range a[:len(a)/2] {
		ans += int64(min(x, mn*2))
	}
	return
}

func abs(x int) int { if x < 0 { return -x }; return x }
func min(a, b int) int { if b < a { return b }; return a }
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$ 或 $O(n)$，其中 $n$ 为 $\textit{basket}_1$ 的长度。用快速选择可以做到 $O(n)$（见 C++）。
- 空间复杂度：$O(n)$。
