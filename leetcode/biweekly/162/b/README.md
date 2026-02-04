由于我们只关心剩余元素的最小值和最大值，不关心元素的顺序，所以可以先从小到大排序，方便后续计算。

排序后，枚举最大值 $\textit{mx} = \textit{nums}[i]$，那么最小值 $\textit{mn} = \textit{nums}[\textit{left}]$ 必须满足

$$
\textit{mn}\cdot k \ge \textit{mx}
$$

在 $[\textit{mn},\textit{mx}]$ 中的元素保留，其余元素去掉。由于排序了，所以这些元素在数组中是**连续**的，问题转化成一个标准的**滑动窗口**模型。如果不满足上式，就把 $\textit{left}$ 加一，直到满足上式。关于滑动窗口的原理，见 [滑动窗口【基础算法精讲 03】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。

内层循环结束后，用窗口长度 $i-\textit{left}+1$ 更新保留元素个数的最大值 $\textit{maxSave}$。

最终答案为 $n - \textit{maxSave}$。

```py [sol-Python3]
class Solution:
    def minRemoval(self, nums: List[int], k: int) -> int:
        nums.sort()
        max_save = left = 0
        for i, mx in enumerate(nums):
            while nums[left] * k < mx:
                left += 1
            max_save = max(max_save, i - left + 1)
        return len(nums) - max_save
```

```java [sol-Java]
class Solution {
    public int minRemoval(int[] nums, int k) {
        Arrays.sort(nums);
        int maxSave = 0;
        int left = 0;
        for (int i = 0; i < nums.length; i++) {
            while ((long) nums[left] * k < nums[i]) {
                left++;
            }
            maxSave = Math.max(maxSave, i - left + 1);
        }
        return nums.length - maxSave;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRemoval(vector<int>& nums, int k) {
        ranges::sort(nums);
        int max_save = 0, left = 0;
        for (int i = 0; i < nums.size(); i++) {
            while (1LL * nums[left] * k < nums[i]) {
                left++;
            }
            max_save = max(max_save, i - left + 1);
        }
        return nums.size() - max_save;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int minRemoval(int* nums, int numsSize, int k) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int max_save = 0;
    int left = 0;
    for (int i = 0; i < numsSize; i++) {
        while (1LL * nums[left] * k < nums[i]) {
            left++;
        }
        max_save = MAX(max_save, i - left + 1);
    }
    return numsSize - max_save;
}
```

```go [sol-Go]
func minRemoval(nums []int, k int) int {
	slices.Sort(nums)
	maxSave, left := 0, 0
	for i, mx := range nums {
		for nums[left]*k < mx {
			left++
		}
		maxSave = max(maxSave, i-left+1)
	}
	return len(nums) - maxSave
}
```

```js [sol-JavaScript]
var minRemoval = function(nums, k) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    let maxSave = 0;
    let left = 0;
    for (let i = 0; i < n; i++) {
        while (nums[left] * k < nums[i]) {
            left++;
        }
        maxSave = Math.max(maxSave, i - left + 1);
    }
    return n - maxSave;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn min_removal(mut nums: Vec<i32>, k: i32) -> i32 {
        nums.sort_unstable();
        let mut max_save = 0;
        let mut left = 0;
        for (i, &x) in nums.iter().enumerate() {
            while (nums[left] as i64) * (k as i64) < x as i64 {
                left += 1;
            }
            max_save = max_save.max(i - left + 1);
        }
        (nums.len() - max_save) as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 专题训练

见下面滑动窗口题单的「**§2.1 越短越合法/求最长/最大**」。

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

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
