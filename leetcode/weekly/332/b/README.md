## 方法一：排序 + 二分查找

由于排序不影响答案，可以先（从小到大）排序，这样可以二分查找。

> $\textit{nums}$ 是 $[1,2,3]$ 还是 $[3,2,1]$，算出来的答案都是一样的，本质上就是从 $\textit{nums}$ 中选两个数。

排序后，枚举 $\textit{nums}[j]$，那么 $\textit{nums}[i]$ 需要满足 $0\le i < j$ 以及

$$
\textit{lower} - \textit{nums}[j] \le \textit{nums}[i] \le \textit{upper} - \textit{nums}[j]
$$

计算 $\le \textit{upper} - \textit{nums}[j]$ 的元素个数，减去 $< \textit{lower} - \textit{nums}[j]$ 的元素个数，即为满足上式的元素个数。（联想一下前缀和）

由于 $\textit{nums}$ 是有序的，我们可以在 $[0,j-1]$ 中**二分查找**，原理见[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)：

- 找到 $> \textit{upper} - \textit{nums}[j]$ 的第一个数，设其下标为 $r$，那么下标在 $[0,r-1]$ 中的数都是 $\le \textit{upper} - \textit{nums}[j]$ 的，这有 $r$ 个。如果 $[0,j-1]$ 中没有找到这样的数，那么二分结果为 $j$。这意味着 $[0,j-1]$ 中的数都是 $\le \textit{upper} - \textit{nums}[j]$ 的，这有 $j$ 个。
- 找到 $\ge \textit{lower} - \textit{nums}[j]$ 的第一个数，设其下标为 $l$，那么下标在 $[0,l-1]$ 中的数都是 $< \textit{lower} - \textit{nums}[j]$ 的，这有 $l$ 个。如果 $[0,j-1]$ 中没有找到这样的数，那么二分结果为 $j$。这意味着 $[0,j-1]$ 中的数都是 $< \textit{lower} - \textit{nums}[j]$ 的，这有 $j$ 个。
- 满足 $\textit{lower} - \textit{nums}[j] \le \textit{nums}[i] \le \textit{upper} - \textit{nums}[j]$ 的 $\textit{nums}[i]$ 的个数为 $r-l$，加入答案。

```py [sol-Python3]
class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        nums.sort()
        ans = 0
        for j, x in enumerate(nums):
            # 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
            r = bisect_right(nums, upper - x, 0, j)
            l = bisect_left(nums, lower - x, 0, j)
            ans += r - l  
        return ans
```

```java [sol-Java]
class Solution {
    public long countFairPairs(int[] nums, int lower, int upper) {
        Arrays.sort(nums);
        long ans = 0;
        for (int j = 0; j < nums.length; j++) {
            // 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
            int r = lowerBound(nums, j, upper - nums[j] + 1);
            int l = lowerBound(nums, j, lower - nums[j]);
            ans += r - l;
        }
        return ans;
    }

    // 原理请看 https://www.bilibili.com/video/BV1AP41137w7/
    private int lowerBound(int[] nums, int right, int target) {
        int left = -1;
        while (left + 1 < right) {
            int mid = (left + right) >>> 1;
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
    long long countFairPairs(vector<int>& nums, int lower, int upper) {
        ranges::sort(nums);
        long long ans = 0;
        for (int j = 0; j < nums.size(); j++) {
            // 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
            auto r = upper_bound(nums.begin(), nums.begin() + j, upper - nums[j]);
            auto l = lower_bound(nums.begin(), nums.begin() + j, lower - nums[j]);
            ans += r - l;
        }
        return ans;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

// 原理请看 https://www.bilibili.com/video/BV1AP41137w7/
int lowerBound(int* nums, int right, int target) {
    int left = -1;
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

long long countFairPairs(int* nums, int numsSize, int lower, int upper) {
    qsort(nums, numsSize, sizeof(int), cmp);
    long long ans = 0;
    for (int j = 0; j < numsSize; j++) {
        // 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
        int r = lowerBound(nums, j, upper - nums[j] + 1);
        int l = lowerBound(nums, j, lower - nums[j]);
        ans += r - l;
    }
    return ans;
}
```

```go [sol-Go]
func countFairPairs(nums []int, lower, upper int) (ans int64) {
    slices.Sort(nums)
    for j, x := range nums {
        // 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
        r := sort.SearchInts(nums[:j], upper-x+1)
        l := sort.SearchInts(nums[:j], lower-x)
        ans += int64(r - l)
    }
    return
}
```

```js [sol-JavaScript]
var countFairPairs = function(nums, lower, upper) {
    nums.sort((a, b) => a - b);
    let ans = 0;
    for (let j = 0; j < nums.length; j++) {
        // 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
        const r = lowerBound(nums, j, upper - nums[j] + 1);
        const l = lowerBound(nums, j, lower - nums[j]);
        ans += r - l;
    }
    return ans;
};

// 原理请看 https://www.bilibili.com/video/BV1AP41137w7/
var lowerBound = function(nums, right, target) {
    let left = -1;
    while (left + 1 < right) {
        const mid = Math.floor((left + right) / 2);
        if (nums[mid] >= target) {
            right = mid;
        } else {
            left = mid;
        }
    }
    return right;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_fair_pairs(mut nums: Vec<i32>, lower: i32, upper: i32) -> i64 {
        nums.sort_unstable();
        let mut ans = 0;
        for j in 0..nums.len() {
            // 注意要在 [0, j-1] 中二分，因为题目要求两个下标 i < j
            let l = nums[..j].partition_point(|&x| x < lower - nums[j]);
            let r = nums[..j].partition_point(|&x| x <= upper - nums[j]);
            ans += r - l;
        }
        ans as _
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 方法二：排序 + 相向三指针

由于随着 $\textit{nums}[j]$ 的变大，$\textit{upper}-\textit{nums}[j]$ 和 $\textit{lower} - \textit{nums}[j]$ 都在变小，有单调性，可以用**相向三指针** $j,l,r$ 代替方法一中的二分查找：

1. 初始化 $l=r=n$。
2. 从左到右遍历（排序后的）$\textit{nums}$。
3. 找 $> \textit{upper} - \textit{nums}[j]$ 的第一个数：如果 $\textit{nums}[r-1] > \textit{upper}-\textit{nums}[j]$，说明 $r$ 太大了，可以继续减小。循环结束后的 $r$，与 $j$ 取最小值后，就是方法一的二分查找计算出的 $r$。
4. 找 $\ge \textit{lower} - \textit{nums}[j]$ 的第一个数：如果 $\textit{nums}[l-1] \ge \textit{lower}-\textit{nums}[j]$，说明 $l$ 太大了，可以继续减小。循环结束后的 $l$，与 $j$ 取最小值后，就是方法一的二分查找计算出的 $l$。

```py [sol-Python3]
class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        nums.sort()
        ans = 0
        l = r = len(nums)
        for j, x in enumerate(nums):
            while r and nums[r - 1] > upper - x:
                r -= 1
            while l and nums[l - 1] >= lower - x:
                l -= 1
            # 在方法一中，二分的结果必须 <= j，方法二同理
            ans += min(r, j) - min(l, j)
        return ans
```

```py [sol-Python3 手写 min]
class Solution:
    def countFairPairs(self, nums: List[int], lower: int, upper: int) -> int:
        nums.sort()
        ans = 0
        l = r = len(nums)
        for j, x in enumerate(nums):
            while r and nums[r - 1] > upper - x:
                r -= 1
            while l and nums[l - 1] >= lower - x:
                l -= 1
            if l < j:
                ans += (r if r < j else j) - l
        return ans
```

```java [sol-Java]
class Solution {
    public long countFairPairs(int[] nums, int lower, int upper) {
        Arrays.sort(nums);
        long ans = 0;
        int l = nums.length;
        int r = nums.length;
        for (int j = 0; j < nums.length; j++) {
            while (r > 0 && nums[r - 1] > upper - nums[j]) {
                r--;
            }
            while (l > 0 && nums[l - 1] >= lower - nums[j]) {
                l--;
            }
            // 在方法一中，二分的结果必须 <= j，方法二同理
            ans += Math.min(r, j) - Math.min(l, j);
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countFairPairs(vector<int>& nums, int lower, int upper) {
        ranges::sort(nums);
        long long ans = 0;
        int l = nums.size(), r = l;
        for (int j = 0; j < nums.size(); j++) {
            while (r && nums[r - 1] > upper - nums[j]) {
                r--;
            }
            while (l && nums[l - 1] >= lower - nums[j]) {
                l--;
            }
            // 在方法一中，二分的结果必须 <= j，方法二同理
            ans += min(r, j) - min(l, j);
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

long long countFairPairs(int* nums, int numsSize, int lower, int upper) {
    qsort(nums, numsSize, sizeof(int), cmp);
    long long ans = 0;
    int l = numsSize, r = numsSize;
    for (int j = 0; j < numsSize; j++) {
        while (r && nums[r - 1] > upper - nums[j]) {
            r--;
        }
        while (l && nums[l - 1] >= lower - nums[j]) {
            l--;
        }
        // 在方法一中，二分的结果必须 <= j，方法二同理
        ans += MIN(r, j) - MIN(l, j);
    }
    return ans;
}
```

```go [sol-Go]
func countFairPairs(nums []int, lower, upper int) (ans int64) {
    slices.Sort(nums)
    l, r := len(nums), len(nums)
    for j, x := range nums {
        for r > 0 && nums[r-1] > upper-x {
            r--
        }
        for l > 0 && nums[l-1] >= lower-x {
            l--
        }
        // 在方法一中，二分的结果必须 <= j，方法二同理
        ans += int64(min(r, j)-min(l, j))
    }
    return
}
```

```js [sol-JavaScript]
var countFairPairs = function(nums, lower, upper) {
    nums.sort((a, b) => a - b);
    let ans = 0, l = nums.length, r = nums.length;
    for (let j = 0; j < nums.length; j++) {
        while (r && nums[r - 1] > upper - nums[j]) {
            r--;
        }
        while (l && nums[l - 1] >= lower - nums[j]) {
            l--;
        }
        // 在方法一中，二分的结果必须 <= j，方法二同理
        ans += Math.min(r, j) - Math.min(l, j);
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_fair_pairs(mut nums: Vec<i32>, lower: i32, upper: i32) -> i64 {
        nums.sort_unstable();
        let mut ans = 0;
        let mut l = nums.len();
        let mut r = nums.len();
        for (j, &x) in nums.iter().enumerate() {
            while r > 0 && nums[r - 1] > upper - x {
                r -= 1;
            }
            while l > 0 && nums[l - 1] >= lower - x {
                l -= 1;
            }
            // 在方法一中，二分的结果必须 <= j，方法二同理
            ans += r.min(j) - l.min(j);
        }
        ans as _
    }
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
