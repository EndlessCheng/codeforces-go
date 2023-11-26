[本题视频讲解](https://www.bilibili.com/video/BV19N411j7Dj/)

视频中通过一个例子，讲了如何具体排序。排序方案都有了，也就可以说明如下结论的正确性：

把 $\textit{nums}[i]$ 及其下标 $i$ 绑在一起排序（也可以单独排序下标），然后把 $\textit{nums}$ 分成若干段，每一段都是递增的且相邻元素之差不超过 $\textit{limit}$，那么这一段可以随意排序。

下面代码用到了分组循环的技巧，详见视频。

```py [sol-Python3]
class Solution:
    def lexicographicallySmallestArray(self, nums: List[int], limit: int) -> List[int]:
        n = len(nums)
        a = sorted(zip(nums, range(n)))
        ans = [0] * n
        i = 0
        while i < n:
            st = i
            i += 1
            while i < n and a[i][0] - a[i - 1][0] <= limit:
                i += 1
            sub = a[st:i]
            sub_idx = sorted(i for _, i in sub)
            for j, (x, _) in zip(sub_idx, sub):
                ans[j] = x
        return ans
```

```java [sol-Java]
class Solution {
    public int[] lexicographicallySmallestArray(int[] nums, int limit) {
        int n = nums.length;
        Integer[] ids = new Integer[n];
        for (int i = 0; i < n; i++) {
            ids[i] = i;
        }
        Arrays.sort(ids, (i, j) -> nums[i] - nums[j]);

        int[] ans = new int[n];
        for (int i = 0; i < n; ) {
            int st = i;
            for (i++; i < n && nums[ids[i]] - nums[ids[i - 1]] <= limit; i++) ;
            Integer[] subIds = Arrays.copyOfRange(ids, st, i);
            Arrays.sort(subIds);
            for (int j = 0; j < subIds.length; j++) {
                ans[subIds[j]] = nums[ids[st + j]];
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> lexicographicallySmallestArray(vector<int> &nums, int limit) {
        int n = nums.size();
        vector<int> ids(n);
        iota(ids.begin(), ids.end(), 0);
        sort(ids.begin(), ids.end(), [&](int i, int j) { return nums[i] < nums[j]; });

        vector<int> ans(n);
        for (int i = 0; i < n;) {
            int st = i;
            for (i++; i < n && nums[ids[i]] - nums[ids[i - 1]] <= limit; i++);
            vector<int> subIds(ids.begin() + st, ids.begin() + i);
            sort(subIds.begin(), subIds.end());
            for (int j = 0; j < subIds.size(); j++) {
                ans[subIds[j]] = nums[ids[st + j]];
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func lexicographicallySmallestArray(nums []int, limit int) []int {
	n := len(nums)
	ids := make([]int, n)
	for i := range ids {
		ids[i] = i
	}
	slices.SortFunc(ids, func(i, j int) int { return nums[i] - nums[j] })

	ans := make([]int, n)
	for i := 0; i < n; {
		st := i
		for i++; i < n && nums[ids[i]]-nums[ids[i-1]] <= limit; i++ {}
		subIds := slices.Clone(ids[st:i])
		slices.Sort(subIds)
		for j, idx := range subIds {
			ans[idx] = nums[ids[st+j]]
		}
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 相似题目

欢迎在评论区补充更多类似题目。

- [1202. 交换字符串中的元素](https://leetcode.cn/problems/smallest-string-with-swaps/)
