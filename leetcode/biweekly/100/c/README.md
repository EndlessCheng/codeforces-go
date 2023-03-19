下午两点【biIibiIi@灵茶山艾府】直播讲题，记得关注哦~

---

技巧题，把 $\textit{nums}[i]$ 及其下标绑定后，按照元素值从小到大排序，元素值相同的按照下标排序。

然后按照题目模拟，用一个 $\textit{vis}$ 数组来实现标记。

也可以生成一个下标数组，对下标排序。具体见 Java 和 C++ 的实现。

```py [sol1-Python3]
class Solution:
    def findScore(self, nums: List[int]) -> int:
        ans = 0
        vis = [False] * (len(nums) + 2)  # 保证下标不越界
        for i, x in sorted(enumerate(nums, 1), key=lambda p: p[1]):
            if not vis[i]:
                vis[i - 1] = vis[i + 1] = True  # 标记相邻的两个元素
                ans += x
        return ans
```

```java [sol1-Java]
class Solution {
    public long findScore(int[] nums) {
        int n = nums.length;
        var ids = new Integer[n];
        for (int i = 0; i < n; ++i) ids[i] = i;
        Arrays.sort(ids, (i, j) -> nums[i] - nums[j]);

        long ans = 0;
        var vis = new boolean[n + 2]; // 保证下标不越界
        for (int i : ids)
            if (!vis[i + 1]) { // 避免 -1，偏移一位
                vis[i] = vis[i + 2] = true;
                ans += nums[i];
            }
        return ans;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    long long findScore(vector<int> &nums) {
        int n = nums.size(), ids[n];
        iota(ids, ids + n, 0);
        stable_sort(ids, ids + n, [&](int i, int j) {
            return nums[i] < nums[j];
        });

        long long ans = 0;
        bool vis[n + 2];  // 保证下标不越界
        memset(vis, 0, sizeof(vis));
        for (int i : ids)
            if (!vis[i + 1]) { // 避免 -1，偏移一位
                vis[i] = vis[i + 2] = true;
                ans += nums[i];
            }
        return ans;
    }
};
```

```go [sol1-Go]
func findScore(nums []int) (ans int64) {
	type pair struct{ v, i int }
	a := make([]pair, len(nums))
	for i, x := range nums {
		a[i] = pair{x, i + 1} // +1 保证下面 for 循环下标不越界
	}
	sort.Slice(a, func(i, j int) bool {
		a, b := a[i], a[j]
		return a.v < b.v || a.v == b.v && a.i < b.i
	})
	vis := make([]bool, len(nums)+2) // 保证下标不越界
	for _, p := range a {
		if !vis[p.i] {
			vis[p.i-1] = true
			vis[p.i+1] = true // 标记相邻的两个元素
			ans += int64(p.v)
		}
	}
	return
}
```

### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$O(n)$。
