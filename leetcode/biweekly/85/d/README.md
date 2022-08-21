下午 2 点在 B 站直播讲周赛和双周赛的题目，感兴趣的小伙伴可以来 [关注](https://space.bilibili.com/206214/dynamic) 一波哦~

---

#### 提示 1

删除不好做，添加比较好做。不妨倒着思考，删除变成了添加。

#### 提示 2

添加时可能会合并两个子段。

我们需要考虑如何动态维护每个子段的元素和，并高效地合并两个子段。

#### 提示 3

用并查集，添加 $x=\textit{removeQueries}[i]$ 时，用并查集合并 $x$ 和 $x+1$，并把 $\textit{nums}[x]$ 加到子段和中。

$\textit{ans}[i]$ 要么取上一个 $\textit{ans}[i+1]$ 的最大子段和，要么取合并后的子段和，这两者取最大值。

```py [sol1-Python3]
class Solution:
    def maximumSegmentSum(self, nums: List[int], removeQueries: List[int]) -> List[int]:
        n = len(nums)
        fa = list(range(n + 1))
        sum = [0] * (n + 1)
        def find(x: int) -> int:
            if fa[x] != x:
                fa[x] = find(fa[x])
            return fa[x]
        ans = [0] * n
        for i in range(n - 1, 0, -1):
            x = removeQueries[i]
            to = find(x + 1)
            fa[x] = to  # 合并 x 和 x+1
            sum[to] += sum[x] + nums[x]
            ans[i - 1] = max(ans[i], sum[to])
        return ans
```

```java [sol1-Java]
class Solution {
    int[] fa;

    public long[] maximumSegmentSum(int[] nums, int[] removeQueries) {
        var n = nums.length;
        fa = new int[n + 1];
        for (var i = 0; i <= n; i++) fa[i] = i;
        var sum = new long[n + 1];

        var ans = new long[n];
        for (var i = n - 1; i > 0; --i) {
            var x = removeQueries[i];
            var to = find(x + 1);
            fa[x] = to; // 合并 x 和 x+1
            sum[to] += sum[x] + nums[x];
            ans[i - 1] = Math.max(ans[i], sum[to]);
        }
        return ans;
    }

    // 非递归并查集
    int find(int x) {
        var f = x;
        while (fa[f] != f) f = find(fa[f]);
        var tmp = x;
        while (fa[tmp] != f) {
            var t = tmp;
            tmp = fa[tmp];
            fa[t] = f;
        }
        return f;
    }
}
```

```cpp [sol1-C++]
class Solution {
public:
    vector<long long> maximumSegmentSum(vector<int> &nums, vector<int> &removeQueries) {
        int n = nums.size();
        int fa[n + 1];
        iota(fa, fa + n + 1, 0);
        long long sum[n + 1];
        memset(sum, 0, sizeof(sum));
        function<int(int)> find = [&](int x) -> int { return fa[x] == x ? x : fa[x] = find(fa[x]); };

        vector<long long> ans(n);
        for (int i = n - 1; i > 0; --i) {
            int x = removeQueries[i];
            int to = find(x + 1);
            fa[x] = to; // 合并 x 和 x+1
            sum[to] += sum[x] + nums[x];
            ans[i - 1] = max(ans[i], sum[to]);
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximumSegmentSum(nums []int, removeQueries []int) (ans []int64) {
	n := len(nums)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}
	sum := make([]int64, n+1)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans = make([]int64, n)
	for i := n - 1; i > 0; i-- {
		x := removeQueries[i]
		to := find(x + 1)
		fa[x] = to // 合并 x 和 x+1
		sum[to] += sum[x] + int64(nums[x])
		ans[i-1] = max(ans[i], sum[to])
	}
	return
}

func max(a, b int64) int64 { if b > a { return b }; return a }
```

#### 思考题

$\textit{nums}$ 有负数要怎么做？
