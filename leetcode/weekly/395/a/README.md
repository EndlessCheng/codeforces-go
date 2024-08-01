考虑数组中的**最小元素**，有

$$
\min(\textit{nums}_1) + x = \min(\textit{nums}_2)
$$

即

$$
x = \min(\textit{nums}_2) - \min(\textit{nums}_1)
$$

```py [sol-Python3]
class Solution:
    def addedInteger(self, nums1: List[int], nums2: List[int]) -> int:
        return min(nums2) - min(nums1)
```

```java [sol-Java]
class Solution {
    public int addedInteger(int[] nums1, int[] nums2) {
        return min(nums2) - min(nums1);
    }

    private int min(int[] nums) {
        int res = Integer.MAX_VALUE;
        for (int x : nums) {
            res = Math.min(res, x);
        }
        return res;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int addedInteger(vector<int>& nums1, vector<int>& nums2) {
        return ranges::min(nums2) - ranges::min(nums1);
    }
};
```

```c [sol-C]
int min(int* nums, int size) {
    int res = nums[0];
    for (int i = 1; i < size; i++) {
        if (nums[i] < res) {
            res = nums[i];
        }
    }
    return res;
}

int addedInteger(int* nums1, int nums1Size, int* nums2, int nums2Size) {
    return min(nums2, nums2Size) - min(nums1, nums1Size);
}
```

```go [sol-Go]
func addedInteger(nums1, nums2 []int) int {
	return slices.Min(nums2) - slices.Min(nums1)
}
```

```js [sol-JavaScript]
var addedInteger = function(nums1, nums2) {
    return Math.min(...nums2) - Math.min(...nums1);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn added_integer(nums1: Vec<i32>, nums2: Vec<i32>) -> i32 {
        nums2.iter().min().unwrap() - nums1.iter().min().unwrap()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}_1$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
