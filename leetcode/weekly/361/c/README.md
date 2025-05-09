## 前置题目

1. [303. 区域和检索 - 数组不可变](https://leetcode.cn/problems/range-sum-query-immutable/)
2. [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/)

## 分析

首先处理 $\textit{cnt}$，把满足 $\textit{nums}[i]\bmod \textit{modulo} = k$ 的 $\textit{nums}[i]$ 视作 $1$，不满足的视作 $0$。

示例 1 的 $\textit{nums}=[3,2,4],\textit{modulo}=2,k=1$，只有 $3\bmod 2 = 1$，所以转化后得到数组 $a=[1,0,0]$。

示例 2 的 $\textit{nums}=[3,1,9,6],\textit{modulo}=3,k=0$，转化后得到数组 $a=[1,0,1,1]$。

现在问题变成统计 $a$ 的子数组中的 $1$ 的个数，这等于子数组的元素和，于是问题转化成：

- 计算 $a$ 中有多少个非空子数组 $b$，满足 $b$ 的元素和模 $\textit{modulo}$ 等于 $k$。

子数组和问题，通常用 [前缀和](https://leetcode.cn/problems/range-sum-query-immutable/solution/qian-zhui-he-ji-qi-kuo-zhan-fu-ti-dan-py-vaar/) 处理。

求出 $a$ 的前缀和数组 $s$，问题转化成：

- 计算有多少个下标对 $(l,r)$，满足 $0\le l<r\le n$ 且 $(s[r]-s[l])\bmod \textit{modulo} = k$。

对于初次接触模运算的同学，可以先从简单的情况开始思考：想一想，如果 $k=0$，要怎么做？这是 [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/) 的取模加强版。

本题保证 $0\le k < \textit{modulo}$，所以 $(s[r]-s[l])\bmod \textit{modulo} = k$ 等价于

$$
(s[r]-s[l])\bmod \textit{modulo} = k \bmod \textit{modulo}
$$

所以 $s[r]-s[l]$ 与 $k$ 关于模 $\textit{modulo}$ **同余**。由于模运算加减法封闭，可以移项，得

$$
(s[r]-k)\bmod \textit{modulo} = s[l]\bmod \textit{modulo}
$$

## 思路

**枚举右，维护左**。根据上式，我们可以一边枚举 $r$，一边统计答案。比如 $(s[r]-k)\bmod \textit{modulo}=6$，我们需要知道在 $r$ 左侧有多少个 $s[l]\bmod \textit{modulo}$ 也等于 $6$，这可以在遍历 $s$ 的过程中，用哈希表（其实只需要数组）统计 $s[l]\bmod \textit{modulo}$ 的出现次数。

## 细节

由于 $a$ 中只有 $0$ 和 $1$，所以 $s[l]\le n$。

由于 $s[l]\bmod \textit{modulo}\le \min(n, \textit{modulo}-1)$，不需要哈希表，用长为 $\min(n+1, \textit{modulo})$ 的数组代替，效率更高。

> 此外，如果 $k>n$，那么 $s[l]\bmod \textit{modulo}\le s[l]\le n < k$，所有子数组都不满足要求，直接返回 $0$。不过这个优化不影响运行时间，可以不写。

因为 $k < \textit{modulo}$，所以 $s[r]-k + \textit{modulo} > s[r]$。又因为 $s$ 是非降数组，所以当我们遍历到 $s[r] < k$ 时，$\textit{cnt}$ 中必然没有 $s[r]-k + \textit{modulo}$，所以此时无需更新 $\textit{ans}$。

## 写法一：前缀和数组

```py [sol-Python3]
class Solution:
    def countInterestingSubarrays(self, nums: List[int], modulo: int, k: int) -> int:
        pre_sum = list(accumulate((x % modulo == k for x in nums), initial=0))
        cnt = [0] * min(len(nums) + 1, modulo)
        ans = 0
        for s in pre_sum:
            if s >= k:
                ans += cnt[(s - k) % modulo]
            cnt[s % modulo] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countInterestingSubarrays(List<Integer> nums, int modulo, int k) {
        int n = nums.size();
        int[] sum = new int[n + 1];
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums.get(i) % modulo == k ? 1 : 0);
        }

        int[] cnt = new int[Math.min(n + 1, modulo)];
        long ans = 0;
        for (int s : sum) {
            if (s >= k) {
                ans += cnt[(s - k) % modulo];
            }
            cnt[s % modulo]++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countInterestingSubarrays(vector<int>& nums, int modulo, int k) {
        int n = nums.size();
        vector<int> sum(n + 1);
        for (int i = 0; i < n; i++) {
            sum[i + 1] = sum[i] + (nums[i] % modulo == k);
        }

        vector<int> cnt(min(n + 1, modulo));
        long long ans = 0;
        for (int s : sum) {
            if (s >= k) {
                ans += cnt[(s - k) % modulo];
            }
            cnt[s % modulo]++;
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

long long countInterestingSubarrays(int* nums, int numsSize, int modulo, int k) {
    int* sum = malloc((numsSize + 1) * sizeof(int));
    sum[0] = 0;
    for (int i = 0; i < numsSize; i++) {
        sum[i + 1] = sum[i] + (nums[i] % modulo == k);
    }

    int* cnt = calloc(MIN(numsSize + 1, modulo), sizeof(int));
    long long ans = 0;
    for (int i = 0; i <= numsSize; i++) {
        int s = sum[i];
        if (s >= k) {
            ans += cnt[(s - k) % modulo];
        }
        cnt[s % modulo]++;
    }

    free(sum);
    free(cnt);
    return ans;
}
```

```go [sol-Go]
func countInterestingSubarrays(nums []int, modulo, k int) (ans int64) {
    sum := make([]int, len(nums)+1)
    for i, x := range nums {
        sum[i+1] = sum[i]
        if x%modulo == k {
            sum[i+1]++
        }
    }

    cnt := make([]int, min(len(nums)+1, modulo))
    for _, s := range sum {
        if s >= k {
            ans += int64(cnt[(s-k)%modulo])
        }
        cnt[s%modulo]++
    }
    return
}
```

```js [sol-JavaScript]
var countInterestingSubarrays = function(nums, modulo, k) {
    const n = nums.length;
    const sum = Array(n + 1);
    sum[0] = 0;
    for (let i = 0; i < n; i++) {
        sum[i + 1] = sum[i] + (nums[i] % modulo === k ? 1 : 0);
    }

    const cnt = Array(Math.min(n + 1, modulo)).fill(0);
    let ans = 0;
    for (const s of sum) {
        if (s >= k) {
            ans += cnt[(s - k) % modulo];
        }
        cnt[s % modulo]++;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_interesting_subarrays(nums: Vec<i32>, modulo: i32, k: i32) -> i64 {
        let n = nums.len();
        let mut sum = vec![0; n + 1];
        for (i, x) in nums.into_iter().enumerate() {
            sum[i + 1] = sum[i];
            if x % modulo == k {
                sum[i + 1] += 1;
            }
        }

        let mut ans = 0;
        let mut cnt = vec![0; (n + 1).min(modulo as usize)];
        for s in sum {
            if s >= k {
                ans += cnt[((s - k) % modulo) as usize] as i64;
            }
            cnt[(s % modulo) as usize] += 1;
        }
        ans
    }
}
```

## 写法二：前缀和变量

```py [sol-Python3]
class Solution:
    def countInterestingSubarrays(self, nums: List[int], modulo: int, k: int) -> int:
        cnt = [0] * min(len(nums) + 1, modulo)
        cnt[0] = 1  # 单独统计 s[0]=0
        ans = s = 0
        for x in nums:
            if x % modulo == k:
                s += 1
            if s >= k:
                ans += cnt[(s - k) % modulo]
            cnt[s % modulo] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countInterestingSubarrays(List<Integer> nums, int modulo, int k) {
        int[] cnt = new int[Math.min(nums.size() + 1, modulo)];
        cnt[0] = 1; // 单独统计 s[0]=0
        long ans = 0;
        int s = 0;
        for (int x : nums) {
            if (x % modulo == k) {
                s++;
            }
            if (s >= k) {
                ans += cnt[(s - k) % modulo];
            }
            cnt[s % modulo]++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countInterestingSubarrays(vector<int>& nums, int modulo, int k) {
        vector<int> cnt(min((int) nums.size() + 1, modulo));
        cnt[0] = 1; // 单独统计 s[0]=0
        long long ans = 0;
        int s = 0;
        for (int x : nums) {
            s += x % modulo == k;
            if (s >= k) {
                ans += cnt[(s - k) % modulo];
            }
            cnt[s % modulo]++;
        }
        return ans;
    }
};
```

```c [sol-C]
#define MIN(a, b) ((b) < (a) ? (b) : (a))

long long countInterestingSubarrays(int* nums, int numsSize, int modulo, int k) {
    int* cnt = calloc(MIN(numsSize + 1, modulo), sizeof(int));
    cnt[0] = 1; // 单独统计 s[0]=0
    long long ans = 0;
    int s = 0;
    for (int i = 0; i < numsSize; i++) {
        s += nums[i] % modulo == k;
        if (s >= k) {
            ans += cnt[(s - k) % modulo];
        }
        cnt[s % modulo]++;
    }
    free(cnt);
    return ans;
}
```

```go [sol-Go]
func countInterestingSubarrays(nums []int, modulo, k int) (ans int64) {
    cnt := make([]int, min(len(nums)+1, modulo))
    cnt[0] = 1 // 单独统计 s[0]=0
    s := 0
    for _, x := range nums {
        if x%modulo == k {
            s++
        }
        if s >= k {
            ans += int64(cnt[(s-k)%modulo])
        }
        cnt[s%modulo]++
    }
    return
}
```

```js [sol-JavaScript]
var countInterestingSubarrays = function(nums, modulo, k) {
    const cnt = Array(Math.min(nums.length + 1, modulo)).fill(0);
    cnt[0] = 1; // 单独统计 s[0]=0
    let ans = 0, s = 0;
    for (const x of nums) {
        if (x % modulo === k) {
            s++;
        }
        if (s >= k) {
            ans += cnt[(s - k) % modulo];
        }
        cnt[s % modulo]++;
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn count_interesting_subarrays(nums: Vec<i32>, modulo: i32, k: i32) -> i64 {
        let mut cnt = vec![0; (nums.len() + 1).min(modulo as usize)];
        cnt[0] = 1; // 单独统计 s[0]=0
        let mut ans = 0;
        let mut s = 0;
        for x in nums {
            if x % modulo == k {
                s += 1;
            }
            if s >= k {
                ans += cnt[((s - k) % modulo) as usize] as i64;
            }
            cnt[(s % modulo) as usize] += 1;
        }
        ans
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(\min(n,\textit{modulo}))$。

更多相似题目，见下面数据结构题单的「**§1.2 前缀和与哈希表**」。

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
