小于 $\textit{nums}$ 第 $k$ 大的数都是满足要求的数。

可以排序+二分查找求小于第 $k$ 大的元素个数；也可以先快速选择求出第 $k$ 大，再遍历 $\textit{nums}$ 统计。

特判 $k=0$，所有元素都满足要求，返回 $n$。

[本题视频讲解](https://www.bilibili.com/video/BV1D4SiB5Ee3/?t=50m32s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def countElements(self, nums: List[int], k: int) -> int:
        n = len(nums)
        if k == 0:
            return n
        nums.sort()
        return bisect_left(nums, nums[-k])  # 小于第 k 大的元素个数
```

```java [sol-Java]
class Solution {
    public int countElements(int[] nums, int k) {
        int n = nums.length;
        if (k == 0) {
            return n;
        }
        Arrays.sort(nums);
        return lowerBound(nums, nums[n - k]); // 小于第 k 大的元素个数
    }

    // 见 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int target) {
        int left = -1;
        int right = nums.length;
        while (left + 1 < right) {
            int mid = left + (right - left) / 2;
            if (nums[mid] >= target) {
                right = mid;
            } else {
                left = mid;
            }
        }
        return right;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countElements(vector<int>& nums, int k) {
        int n = nums.size();
        if (k == 0) {
            return n;
        }
        ranges::sort(nums);
        return ranges::lower_bound(nums, nums[n - k]) - nums.begin(); // 小于第 k 大的元素个数
    }
};
```

```cpp [sol-C++ 快速选择]
class Solution {
public:
    int countElements(vector<int>& nums, int k) {
        int n = nums.size();
        if (k == 0) {
            return n;
        }
        ranges::nth_element(nums, nums.end() - k);
        int kth = nums[n - k], ans = 0;
        for (int x : nums) {
            ans += x < kth;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countElements(nums []int, k int) int {
	n := len(nums)
	if k == 0 {
		return n
	}
	slices.Sort(nums)
	return sort.SearchInts(nums, nums[n-k]) // 小于第 k 大的元素个数
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$ 或 $\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。快速选择可以做到 $\mathcal{O}(n)$，见 C++ 第二份代码。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 附表

| **需求**  | **写法**  |
|---|---|
| $< x$ 的元素个数  | $\texttt{lowerBound}(\textit{nums},x)$  | 
| $\le x$ 的元素个数 | $\texttt{lowerBound}(\textit{nums},x+1)$  | 
| $\ge x$ 的元素个数  | $n - \texttt{lowerBound}(\textit{nums},x)$  | 
| $> x$ 的元素个数  | $n - \texttt{lowerBound}(\textit{nums},x+1)$  | 

注意 $< x$ 和 $\ge x$ 互为补集，元素个数之和为 $n$。$\le x$ 和 $> x$ 同理。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
