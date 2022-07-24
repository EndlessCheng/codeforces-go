本题 [视频讲解](https://www.bilibili.com/video/BV14a411U7QZ?t=11m) 已出炉，额外讲解了**如何用集合论来思考二进制**。欢迎点赞三连，在评论区分享你对这场周赛的看法~

---

#### 提示 1

对于 $x|y$ 和 $x\&y$ 中的 $1$，在同一个比特位上，如果都有 $1$，那这个 $1$ 会被统计两次；如果一个为 $1$ 另一个为 $0$，那这个 $1$ 会被统计一次。

#### 提示 2

例如 $x=110$，$y=011$，只统计一次的部分为 $x'=100$，$y'=001$，统计了两次的部分为 $x\&y=010$。我们可以直接把 $010$ 重新分配到 $x'$ 和 $y'$ 上，这样又得到了 $x$ 和 $y$。

记 $c(x)$ 为 $x$ 的二进制表示中的 $1$ 的个数，则有如下等式：

$$
c(x|y)+c(x\&y)=c(x)+c(y)
$$

#### 提示 3

遍历去重后的 $\textit{nums}$，统计 $c(\textit{nums}[i])$ 的个数，记录在 $\textit{cnt}$ 中，然后写一个二重循环遍历 $\textit{cnt}$，对于所有的 $c(x)+c(y)\ge k$，累加 $\textit{cnt}[c(x)]\cdot\textit{cnt}[c(y)]$，表示从这两组中各选一个 $x$ 和 $y$ 组成优质数对的个数。

#### 复杂度分析

- 时间复杂度：$O(n+U^2)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U$ 为不同的 $c(\textit{nums}[i])$ 的个数，不超过 $30$。
- 空间复杂度：$O(n+U)$。

```py [sol1-Python3]
class Solution:
    def countExcellentPairs(self, nums: List[int], k: int) -> int:
        cnt = Counter(x.bit_count() for x in set(nums))
        ans = 0
        for cx, ccx in cnt.items():
            for cy, ccy in cnt.items():
                if cx + cy >= k:  # (x,y) 是优质数对
                    ans += ccx * ccy
        return ans
```

```java [sol1-Java]
class Solution {
    public long countExcellentPairs(int[] nums, int k) {
        var ans = 0L;
        var vis = new HashSet<Integer>();
        var cnt = new HashMap<Integer, Integer>();
        for (var x : nums)
            if (!vis.contains(x)) {
                vis.add(x);
                var c = Integer.bitCount(x);
                cnt.put(c, cnt.getOrDefault(c, 0) + 1);
            }
        for (var x : cnt.entrySet())
            for (var y : cnt.entrySet())
                if (x.getKey() + y.getKey() >= k)
                    ans += (long) x.getValue() * y.getValue();
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long countExcellentPairs(vector<int> &nums, int k) {
        unordered_map<int, int> cnt;
        for (int x : unordered_set<int>(nums.begin(), nums.end())) // 去重
            ++cnt[__builtin_popcount(x)];
        long ans = 0L;
        for (auto &[cx, ccx] : cnt)
            for (auto &[cy, ccy] : cnt)
                if (cx + cy >= k) // (x,y) 是优质数对
                    ans += (long) ccx * ccy;
        return ans;
    }
};
```

```go [sol1-Go]
func countExcellentPairs(nums []int, k int) (ans int64) {
	vis := map[int]bool{}
	cnt := map[int]int{}
	for _, x := range nums {
		if !vis[x] {
			vis[x] = true
			cnt[bits.OnesCount(uint(x))]++
		}
	}
	for cx, ccx := range cnt {
		for cy, ccy := range cnt {
			if cx+cy >= k { // (x,y) 是优质数对
				ans += int64(ccx) * int64(ccy)
			}
		}
	}
	return
}
```

进一步地，二重循环可以用前缀和（或者后缀和）来优化。

我们可以从小到大遍历 $c(x)$，根据 $c(y)\ge k-c(x)$，对应的 $c(y)$ 也会从大到小减小，我们可以用后缀和维护这些 $cnt[c(y)]$ 的和。

#### 复杂度分析

- 时间复杂度：$O(n+U)$，其中 $n$ 为 $\textit{nums}$ 的长度，$U=30$。
- 空间复杂度：$O(n+U)$。

```py [sol2-Python3]
class Solution:
    def countExcellentPairs(self, nums: List[int], k: int) -> int:
        cnt = [0] * 30
        for x in set(nums):
            cnt[x.bit_count()] += 1
        ans = 0
        s = sum(cnt[k:])
        for cx, ccx in enumerate(cnt):
            ans += ccx * s
            if 0 <= k - 1 - cx < 30:
                s += cnt[k - 1 - cx]
        return ans
```

```java [sol2-Java]
class Solution {
    static final int U = 30;

    public long countExcellentPairs(int[] nums, int k) {
        var vis = new HashSet<Integer>();
        var cnt = new int[U];
        for (var x : nums)
            if (!vis.contains(x)) {
                vis.add(x);
                ++cnt[Integer.bitCount(x)];
            }
        var ans = 0L;
        var s = 0;
        for (var i = k; i < U; i++)
            s += cnt[i];
        for (int cx = 0; cx < U; cx++) {
            ans += (long) cnt[cx] * s;
            var cy = k - 1 - cx;
            if (0 <= cy && cy < U) s += cnt[cy];
        }
        return ans;
    }
}
```

```cpp [sol2-C++]
class Solution {
    static constexpr int U = 30;
public:
    long long countExcellentPairs(vector<int> &nums, int k) {
        int cnt[U] = {};
        for (int x : unordered_set<int>(nums.begin(), nums.end())) // 去重
            ++cnt[__builtin_popcount(x)];
        long ans = 0L;
        int s = 0;
        for (int i = k; i < U; i++)
            s += cnt[i];
        for (int cx = 0; cx < U; cx++) {
            ans += (long) cnt[cx] * s;
            int cy = k - 1 - cx;
            if (0 <= cy && cy < U) s += cnt[cy];
        }
        return ans;
    }
};
```

```go [sol2-Go]
func countExcellentPairs(nums []int, k int) (ans int64) {
	vis := map[int]bool{}
	cnt := [30]int{}
	for _, x := range nums {
		if !vis[x] {
			vis[x] = true
			cnt[bits.OnesCount(uint(x))]++
		}
	}
	s := 0
	for i := k; i < 30; i++ {
		s += cnt[i]
	}
	for cx, ccx := range cnt {
		ans += int64(ccx) * int64(s)
		cy := k - 1 - cx
		if 0 <= cy && cy < 30 {
			s += cnt[cy]
		}
	}
	return
}
```
