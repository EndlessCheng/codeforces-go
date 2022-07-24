下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

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

遍历去重后的 $\textit{nums}$，统计 $c(\textit{nums}[i])$ 的个数，记录在 $\textit{cnt}$ 中，然后写一个二重循环遍历 $\textit{cnt}$，对于所有的 $c(x)+c(y)\ge k$，累加 $\textit{cnt}[c(x)]\cdot\textit{cnt}[c(y)]$，表示从这两组中各选一个 $x$ 和 $y$ 组成优质数对个数。

#### 复杂度分析

- 时间复杂度：$O(n+k^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n+k)$。

```py [sol1-Python3]
class Solution:
    def countExcellentPairs(self, nums: List[int], k: int) -> int:
        cnt = Counter(x.bit_count() for x in set(nums))
        ans = 0
        for cx, ccx in cnt.items():
            for cy, ccy in cnt.items():
                if cx + cy >= k:
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
        long ans = 0L;
        map<int, int> cnt;
        unordered_set<int> set(nums.begin(), nums.end());
        for (int x : set) ++cnt[__builtin_popcount(x)];
        for (auto &[cx, ccx] : cnt)
            for (auto &[cy, ccy] : cnt)
                if (cx + cy >= k)
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
