本题 [视频讲解](https://www.bilibili.com/video/BV1Kd4y1Z7Fv) 已出炉，欢迎点赞三连，在评论区分享你对这场双周赛的看法~

---

#### 提示 1

删除不好做，添加比较好做。不妨倒着思考，删除变成了添加。

#### 提示 2

添加时可能会合并两个子段。

我们需要考虑如何动态维护每个子段的元素和，并高效地合并两个子段。

#### 提示 3

用并查集（[视频讲解](https://www.bilibili.com/video/BV1Kd4y1Z7Fv) 中讲了原理），添加下标 $x=\textit{removeQueries}[i]$ 时，用并查集合并 $x$ 和 $x+1$，并把 $\textit{nums}[x]$ 加到子段和中。

以 $\textit{removeQueries}=[3,1,2,0]$ 为例说明。倒序遍历，我们会先合并下标 $0$ 和 $1$，这样相当于创建了一个下标子段 $[0]$；然后合并 $2$ 和 $3$，创建了下标子段 $[2]$；然后合并 $1$ 和 $2$，由于 $0$ 和 $1$ 已经合并了，这一操作会把 $0 1 2$ 都合并起来，最终形成下标子段 $[0,1,2]$。注意，这一合并过程中会形成若干条「链」，每一条链去掉最右边的元素就等价于实际的下标子段。

另外一种理解方式是，把链看成是一列火车，这列火车有一节「幽灵火车头」，不算在实际的子段中。如果要合并两条链，就需要把左边这节幽灵火车头作为一节实际的车厢加到右边这列火车中，因此我们只需要合并 $x$ 和 $x+1$，不需要合并 $x$ 和 $x-1$。

最后，对于 $\textit{ans}[i]$，要么取上一个 $\textit{ans}[i+1]$ 的最大子段和，要么取合并后的子段和，这两者取最大值。

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

    int find(int x) {
        if (fa[x] != x) fa[x] = find(fa[x]);
        return fa[x];
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

#### 相似题目

- [2334. 元素值大于变化阈值的子数组](https://leetcode.cn/problems/subarray-with-elements-greater-than-varying-threshold/)
- [1562. 查找大小为 M 的最新分组](https://leetcode.cn/problems/find-latest-group-of-size-m/)

#### 思考题

$\textit{nums}$ 有负数要怎么做？

如果询问的是 $\textit{nums}$ 某个子区间的最大子段和呢？

更加一般的情况，见 [洛谷 P4513 小白逛公园](https://www.luogu.com.cn/problem/P4513) 和一道相似的问题 [2213. 由单个字符重复的最长子字符串](https://leetcode.cn/problems/longest-substring-of-one-repeating-character/)。
