## 前置知识：相向双指针

请看[【基础算法精讲】](https://www.bilibili.com/video/BV1bP411c7oJ/)。

## 思路

题目相当于从数组中选两个数，**我们只关心这两个数的和是否小于** $\textit{target}$，由于 $a+b=b+a$，无论如何排列数组元素，都不会影响加法的结果，所以**排序不影响答案**。

排序后：

1. 初始化左右指针 $\textit{left}=0,\ \textit{right}=n-1$。
2. 如果 $\textit{nums}[\textit{left}]+\textit{nums}[\textit{right}] < \textit{target}$，由于数组是有序的，$\textit{nums}[\textit{left}]$ 与下标 $i$ 在区间 $[\textit{left}+1,\textit{right}]$ 中的任何 $\textit{nums}[i]$ 相加，都是 $<\textit{target}$ 的，因此直接找到了 $\textit{right}-\textit{left}$ 个合法数对，加到答案中，然后将 $\textit{left}$ 加一。
3. 如果 $\textit{nums}[\textit{left}]+\textit{nums}[\textit{right}] \ge \textit{target}$，由于数组是有序的，$\textit{nums}[\textit{right}]$ 与下标 $i$ 在区间 $[\textit{left},\textit{right}-1]$ 中的任何 $\textit{nums}[i]$ 相加，都是 $\ge\textit{target}$ 的，因此后面无需考虑 $\textit{nums}[\textit{right}]$，将 $\textit{right}$ 减一。
4. 重复上述过程直到 $\textit{left}\ge \textit{right}$ 为止。

```py [sol-Python3]
class Solution:
    def countPairs(self, nums: List[int], target: int) -> int:
        nums.sort()
        ans = left = 0
        right = len(nums) - 1
        while left < right:
            if nums[left] + nums[right] < target:
                ans += right - left
                left += 1
            else:
                right -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int countPairs(List<Integer> nums, int target) {
        Collections.sort(nums);
        int ans = 0;
        int left = 0;
        int right = nums.size() - 1;
        while (left < right) {
            if (nums.get(left) + nums.get(right) < target) {
                ans += right - left;
                left++;
            } else {
                right--;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int countPairs(vector<int>& nums, int target) {
        ranges::sort(nums);
        int ans = 0, left = 0, right = nums.size() - 1;
        while (left < right) {
            if (nums[left] + nums[right] < target) {
                ans += right - left;
                left++;
            } else {
                right--;
            }
        }
        return ans;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int countPairs(int* nums, int numsSize, int target) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int ans = 0, left = 0, right = numsSize - 1;
    while (left < right) {
        if (nums[left] + nums[right] < target) {
            ans += right - left;
            left++;
        } else {
            right--;
        }
    }
    return ans;
}
```

```go [sol-Go]
func countPairs(nums []int, target int) (ans int) {
	slices.Sort(nums)
	left, right := 0, len(nums)-1
	for left < right {
		if nums[left]+nums[right] < target {
			ans += right - left
			left++
		} else {
			right--
		}
	}
	return
}
```

```js [sol-JavaScript]
var countPairs = function(nums, target) {
    nums.sort((a, b) => a - b);
    let ans = 0, left = 0, right = nums.length - 1;
    while (left < right) {
        if (nums[left] + nums[right] < target) {
            ans += right - left;
            left++;
        } else {
            right--;
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_pairs(mut nums: Vec<i32>, target: i32) -> i32 {
        nums.sort_unstable();
        let mut ans = 0;
        let mut left = 0;
        let mut right = nums.len() - 1;
        while left < right {
            if nums[left] + nums[right] < target {
                ans += right - left;
                left += 1;
            } else {
                right -= 1;
            }
        }
        ans as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。不计入排序的栈开销，仅用到若干额外变量。

注：由于本题数据范围小，评测机对运行时间的影响更大（运行时间存在一定波动）。如果你发现击败百分比不高，可以无视。也可以多提交几次试试。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心算法（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
