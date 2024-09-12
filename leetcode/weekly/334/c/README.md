## 方法一：二分答案

### 提示 1

如果 $2\cdot\textit{nums}[i]\le \textit{nums}[j]$，则称 $\textit{nums}[i]$ 与 $\textit{nums}[j]$ **匹配**。

如果可以匹配 $k$ 对，那么也可以匹配小于 $k$ 对，去掉一些数对即可做到。

如果无法匹配 $k$ 对，那么也无法匹配大于 $k$ 对（反证法）。

所以 $k$ 越大，越无法选出 $k$ 个能匹配的数对。有单调性，就可以**二分答案**。二分算法的理论讲解见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

### 提示 2

现在问题变成：

- 能否从 $\textit{nums}$ 中选出 $k$ 个能匹配的数对？

要让哪些数匹配呢？

**结论：从小到大排序后，如果存在 $k$ 对匹配，那么一定可以让最小的 $k$ 个数与最大的 $k$ 个数匹配。**

证明：假设不是最小的 $k$ 个数与最大的 $k$ 个数匹配，那么我们总是可以把 $\textit{nums}[i]$ 替换成比它小的且不在匹配中的数，这仍然是匹配的；同理，把 $\textit{nums}[j]$ 替换成比它大的且不在匹配中的数，这仍然是匹配的。所以如果存在 $k$ 对匹配，那么一定可以让最小的 $k$ 个数和最大的 $k$ 个数匹配。

反过来说，如果最小的 $k$ 个数无法和最大的 $k$ 个数匹配，则任意 $k$ 对都无法匹配。（也可以用反证法证明）

从小到大排序后，$\textit{nums}[0]$ 要与 $\textit{nums}[n-k]$ 匹配。如果不这样做，$\textit{nums}[0]$ 与在 $\textit{nums}[n-k]$ 右侧的数匹配，相当于占了一个位置，那么后续要选个更大的 $\textit{nums}[i]$ 与 $\textit{nums}[n-k]$ 匹配，这不一定能匹配得上。

一般地，$\textit{nums}[i]$ 要与 $\textit{nums}[n-k+i]$ 匹配。

如果对于所有的 $0\le i < k$，都满足 $2\cdot\textit{nums}[i]\le\textit{nums}[n-k+i]$，那么就可以从 $\textit{nums}$ 中选出 $k$ 个能匹配的数对。

### 细节

下面代码采用开区间二分，这仅仅是二分的一种写法，使用闭区间或者半闭半开区间都是可以的。

- 开区间左端点初始值：$0$。无论 $\textit{nums}$ 是什么样，一定能选出 $0$ 个匹配。
- 开区间右端点初始值：$\left\lfloor\dfrac{n}{2}\right\rfloor+1$。最多能选 $\left\lfloor\dfrac{n}{2}\right\rfloor$ 个匹配，再多一个就不行了。

对于开区间写法，简单来说 `check(mid) == true` 时更新的是谁，最后就返回谁。相比其他二分写法，开区间写法不需要思考加一减一等细节，更简单。推荐使用开区间写二分。

```py [sol-Python3]
class Solution:
    def maxNumOfMarkedIndices(self, nums: List[int]) -> int:
        nums.sort()
        left, right = 0, len(nums) // 2 + 1  # 开区间
        while left + 1 < right:
            k = (left + right) // 2
            if all(nums[i] * 2 <= nums[i - k] for i in range(k)):
                left = k
            else:
                right = k
        return left * 2  # 最多匹配 left 对，有 left * 2 个数
```

```py [sol-Python3 库函数]
# 二分找最小的不满足要求的 k+1，二分结束后，k 就是最大的满足要求的数对个数
class Solution:
    def maxNumOfMarkedIndices(self, nums: List[int]) -> int:
        nums.sort()
        check = lambda k: any(nums[i] * 2 > nums[i - k - 1] for i in range(k + 1))
        return bisect_left(range(len(nums) // 2), True, key=check) * 2
```

```java [sol-Java]
class Solution {
    public int maxNumOfMarkedIndices(int[] nums) {
        Arrays.sort(nums);
        int left = 0;
        int right = nums.length / 2 + 1; // 开区间
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
            if (check(nums, mid)) {
                left = mid;
            } else {
                right = mid;
            }
        }
        return left * 2; // 最多匹配 left 对，有 left * 2 个数
    }

    private boolean check(int[] nums, int k) {
        for (int i = 0; i < k; i++) {
            if (nums[i] * 2 > nums[nums.length - k + i]) {
                return false;
            }
        }
        return true;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxNumOfMarkedIndices(vector<int>& nums) {
        ranges::sort(nums);

        auto check = [&](int k) -> bool {
            for (int i = 0; i < k; i++) {
                if (nums[i] * 2 > nums[nums.size() - k + i]) {
                    return false;
                }
            }
            return true;
        };

        int left = 0, right = nums.size() / 2 + 1; // 开区间
        while (left + 1 < right) {
            int mid = (left + right) / 2;
            (check(mid) ? left : right) = mid;
        }
        return left * 2; // 最多匹配 left 对，有 left * 2 个数
    }
};
```

```c [sol-C]
int cmp(const void *a, const void *b) {
    return *(int*)a - *(int*)b;
}

bool check(int* nums, int numsSize, int k) {
    for (int i = 0; i < k; i++) {
        if (nums[i] * 2 > nums[numsSize - k + i]) {
            return false;
        }
    }
    return true;
}

int maxNumOfMarkedIndices(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int left = 0, right = numsSize / 2 + 1; // 开区间
    while (left + 1 < right) {
        int mid = (left + right) / 2;
        if (check(nums, numsSize, mid)) {
            left = mid;
        } else {
            right = mid;
        }
    }
    return left * 2; // 最多匹配 left 对，有 left * 2 个数
}
```

```go [sol-Go]
func maxNumOfMarkedIndices(nums []int) int {
    slices.Sort(nums)
    n := len(nums)
    pairs := sort.Search(n/2, func(k int) bool {
        k++
        for i, x := range nums[:k] {
            if x*2 > nums[n-k+i] {
                return true
            }
        }
        return false
    })
    return pairs * 2 // 最多匹配 pairs 对，有 pairs * 2 个数
}
```

```js [sol-JavaScript]
var maxNumOfMarkedIndices = function(nums) {
    nums.sort((a, b) => a - b);

    function check(k) {
        for (let i = 0; i < k; i++) {
            if (nums[i] * 2 > nums[nums.length - k + i]) {
                return false;
            }
        }
        return true;
    }

    let left = 0, right = Math.floor(nums.length / 2) + 1; // 开区间
    while (left + 1 < right) {
        const mid = Math.floor((left + right) / 2);
        if (check(mid)) {
            left = mid;
        } else {
            right = mid;
        }
    }
    return left * 2; // 最多匹配 left 对，有 left * 2 个数
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_num_of_marked_indices(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();

        let check = |k: usize| -> bool {
            for i in 0..k {
                if nums[i] * 2 > nums[nums.len() - k + i] {
                    return false;
                }
            }
            true
        };

        let mut left = 0;
        let mut right = nums.len() / 2 + 1; // 开区间
        while left + 1 < right {
            let mid = (left + right) / 2;
            if check(mid) {
                left = mid;
            } else {
                right = mid;
            }
        }
        (left * 2) as _ // 最多匹配 left 对，有 left * 2 个数
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。

## 方法二：同向双指针

由方法一的匹配方式可知，我们需要用 $\textit{nums}$ 左半部分中的数，去匹配 $\textit{nums}$ 右半部分中的数。

在 $\textit{nums}$ 的右半部分中，找到第一个满足 $2\cdot\textit{nums}[0]\le \textit{nums}[j]$ 的 $j$，那么 $\textit{nums}[1]$ 只能匹配右半部分中的下标大于 $j$ 的数，依此类推。

这可以用同向双指针实现。

[本题视频讲解](https://www.bilibili.com/video/BV1wj411G7sH/)（第三题）

```py [sol-Python3]
class Solution:
    def maxNumOfMarkedIndices(self, nums: List[int]) -> int:
        nums.sort()
        i = 0
        for x in nums[(len(nums) + 1) // 2:]:
            if nums[i] * 2 <= x:  # 找到一个匹配
                i += 1
        return i * 2
```

```java [sol-Java]
class Solution {
    public int maxNumOfMarkedIndices(int[] nums) {
        Arrays.sort(nums);
        int n = nums.length;
        int i = 0;
        for (int j = (n + 1) / 2; j < n; j++) {
            if (nums[i] * 2 <= nums[j]) { // 找到一个匹配
                i++;
            }
        }
        return i * 2;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maxNumOfMarkedIndices(vector<int>& nums) {
        ranges::sort(nums);
        int i = 0, n = nums.size();
        for (int j = (n + 1) / 2; j < n; j++) {
            if (nums[i] * 2 <= nums[j]) { // 找到一个匹配
                i++;
            }
        }
        return i * 2;
    }
};
```

```c [sol-C]
int cmp(const void *a, const void *b) {
    return *(int*)a - *(int*)b;
}

int maxNumOfMarkedIndices(int* nums, int numsSize) {
    qsort(nums, numsSize, sizeof(int), cmp);
    int i = 0;
    for (int j = (numsSize + 1) / 2; j < numsSize; j++) {
        if (nums[i] * 2 <= nums[j]) { // 找到一个匹配
            i++;
        }
    }
    return i * 2;
}
```

```go [sol-Go]
func maxNumOfMarkedIndices(nums []int) int {
    slices.Sort(nums)
    i := 0
    for _, x := range nums[(len(nums)+1)/2:] {
        if nums[i]*2 <= x { // 找到一个匹配
            i++
        }
    }
    return i * 2
}
```

```js [sol-JavaScript]
var maxNumOfMarkedIndices = function(nums) {
    nums.sort((a, b) => a - b);
    const n = nums.length;
    let i = 0;
    for (let j = Math.floor((n + 1) / 2); j < n; j++) {
        if (nums[i] * 2 <= nums[j]) { // 找到一个匹配
            i++;
        }
    }
    return i * 2;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn max_num_of_marked_indices(mut nums: Vec<i32>) -> i32 {
        nums.sort_unstable();
        let mut i = 0;
        for &x in &nums[(nums.len() + 1) / 2..] {
            if nums[i] * 2 <= x { // 找到一个匹配
                i += 1;
            }
        }
        (i * 2) as _
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销，仅用到若干额外变量。Python 忽略切片开销。

更多相似题目，见下面的二分题单，以及贪心题单中的「**§1.2 单序列配对**」和「**§1.3 双序列配对**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
