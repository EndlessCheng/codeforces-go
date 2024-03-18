请看 [视频讲解](https://www.bilibili.com/video/BV11x421r7q5/) 第二题。

由于要按照元素从小到大标记，但又不能直接对数组排序（因为有对特定 $\textit{index}$ 的标记），我们可以创建一个 $\textit{ids}$ 数组，其中 $\textit{ids}[i]=i$，然后对该数组按照 $\textit{nums}[\textit{ids}[i]]$ 从小到大排序。注意要使用**稳定排序**，因为相同元素值的下标需要按照下标从小到大排。也可以使用不稳定排序（如快速排序），但要对于相同元素值按照下标从小到大排序。

设 $\textit{nums}$ 的元素和为 $s$。对于每个询问，我们先将 $s$ 减少 $\textit{nums}[\textit{index}]$，然后将 $\textit{nums}[\textit{index}]$ 置为 $0$，就相当于标记了这个数（因为题目保证数组元素都是正数）。然后依照 $\textit{ids}$ 找 $k$ 个最小的没有被标记的数，将其标记，标记的同时维护 $s$。

```py [sol-Python3]
class Solution:
    def unmarkedSumArray(self, nums: List[int], queries: List[List[int]]) -> List[int]:
        n = len(nums)
        s = sum(nums)
        ids = sorted(range(n), key=lambda i: nums[i])  # 稳定排序
        ans = []
        j = 0
        for i, k in queries:
            s -= nums[i]
            nums[i] = 0  # 标记
            while j < n and k:
                i = ids[j]
                if nums[i]:  # 没有被标记
                    s -= nums[i]
                    nums[i] = 0
                    k -= 1
                j += 1
            ans.append(s)
        return ans
```

```java [sol-Java]
class Solution {
    public long[] unmarkedSumArray(int[] nums, int[][] queries) {
        int n = nums.length;
        long s = 0;
        Integer[] ids = new Integer[n];
        for (int i = 0; i < n; i++) {
            s += nums[i];
            ids[i] = i;
        }
        Arrays.sort(ids, (i, j) -> nums[i] - nums[j]); // 稳定排序

        long[] ans = new long[queries.length];
        int j = 0;
        for (int qi = 0; qi < queries.length; qi++) {
            int[] q = queries[qi];
            int i = q[0];
            int k = q[1];
            s -= nums[i];
            nums[i] = 0; // 标记
            for (; j < n && k > 0; j++) {
                i = ids[j];
                if (nums[i] > 0) { // 没有被标记
                    s -= nums[i];
                    nums[i] = 0;
                    k--;
                }
            }
            ans[qi] = s;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<long long> unmarkedSumArray(vector<int> &nums, vector<vector<int>> &queries) {
        int n = nums.size();
        long long s = accumulate(nums.begin(), nums.end(), 0LL);
        vector<int> ids(n);
        iota(ids.begin(), ids.end(), 0);
        ranges::stable_sort(ids, [&](int i, int j) { return nums[i] < nums[j]; });

        vector<long long> ans;
        int j = 0;
        for (auto &q : queries) {
            int i = q[0], k = q[1];
            s -= nums[i];
            nums[i] = 0; // 标记
            for (; j < n && k; j++) {
                i = ids[j];
                if (nums[i] > 0) { // 没有被标记
                    s -= nums[i];
                    nums[i] = 0;
                    k--;
                }
            }
            ans.push_back(s);
        }
        return ans;
    }
};
```

```go [sol-Go]
func unmarkedSumArray(nums []int, queries [][]int) []int64 {
	s, n := 0, len(nums)
	ids := make([]int, n)
	for i, x := range nums {
		s += x
		ids[i] = i
	}
	slices.SortStableFunc(ids, func(i, j int) int { return nums[i] - nums[j] })

	ans := make([]int64, len(queries))
	j := 0
	for qi, p := range queries {
		i, k := p[0], p[1]
		s -= nums[i]
		nums[i] = 0 // 标记
		for ; j < n && k > 0; j++ {
			i := ids[j]
			if nums[i] > 0 { // 没有被标记
				s -= nums[i]
				nums[i] = 0
				k--
			}
		}
		ans[qi] = int64(s)
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(n)$。忽略返回值的空间。

## 题单

- [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
- [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
- [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
- [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
- [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
- [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)

更多题单，请点我个人主页 - 讨论发布。
