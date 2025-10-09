下文记 $m=\textit{value}$。

由于同一个数可以加减任意倍的 $m$，我们可以先把每个 $\textit{nums}[i]$ 变成与 $\textit{nums}[i]$ 关于模 $m$ **同余**的最小非负整数，以备后用。关于同余的介绍，请看 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

本题有负数，根据这篇文章中的公式，我们可以把每个 $\textit{nums}[i]$ 变成

$$
(\textit{nums}[i]\bmod m + m)\bmod m
$$

从而保证取模结果在 $[0,m)$ 中。

例如 $\textit{nums}=[1,-6,-4,3,5]$，$m=3$，取模后变成 $[1,0,2,0,2]$。

然后枚举答案：

- 有没有与 $0$ 关于模 $m$ 同余的数？有，我们消耗掉一个 $0$。
- 有没有与 $1$ 关于模 $m$ 同余的数？有，我们消耗掉一个 $1$。
- 有没有与 $2$ 关于模 $m$ 同余的数？有，我们消耗掉一个 $2$。
- 有没有与 $3$ 关于模 $m$ 同余的数？有，我们消耗掉一个 $0$。这个取模后等于 $0$ 的数，可以继续操作，变成 $3$。
- 有没有与 $4$ 关于模 $m$ 同余的数？也就是看是否还有 $1$，没有，那么答案等于 $4$。

怎么知道还有没有剩余元素？用一个哈希表 $\textit{cnt}$ 统计 $(\textit{nums}[i]\bmod m + m) \bmod m$ 的个数。

[本题视频讲解](https://www.bilibili.com/video/BV1EL411C7YU/?t=42m20s)，欢迎点赞关注~

## 写法一

```py [sol-Python3]
class Solution:
    def findSmallestInteger(self, nums: List[int], m: int) -> int:
        cnt = Counter(x % m for x in nums)
        mex = 0
        while cnt[mex % m]:
            cnt[mex % m] -= 1
            mex += 1
        return mex
```

```java [sol-Java]
class Solution {
    public int findSmallestInteger(int[] nums, int m) {
        int[] cnt = new int[m];
        for (int x : nums) {
            cnt[(x % m + m) % m]++; // 保证取模结果在 [0, m) 中
        }

        int mex = 0;
        while (cnt[mex % m]-- > 0) {
            mex++;
        }
        return mex;
    }
}
```

```java [sol-Java 哈希表]
class Solution {
    public int findSmallestInteger(int[] nums, int m) {
        Map<Integer, Integer> cnt = new HashMap<>();
        for (int x : nums) {
            cnt.merge((x % m + m) % m, 1, Integer::sum);
        }

        int mex = 0;
        while (cnt.merge(mex % m, -1, Integer::sum) >= 0) {
            mex++;
        }
        return mex;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findSmallestInteger(vector<int>& nums, int m) {
        unordered_map<int, int> cnt;
        for (int x : nums) {
            cnt[(x % m + m) % m]++; // 保证取模结果在 [0, m) 中
        }

        int mex = 0;
        while (cnt[mex % m]-- > 0) {
            mex++;
        }
        return mex;
    }
};
```

```c [sol-C]
int findSmallestInteger(int* nums, int numsSize, int m) {
    int* cnt = calloc(m, sizeof(int));
    for (int i = 0; i < numsSize; i++) {
        cnt[(nums[i] % m + m) % m]++; // 保证取模结果在 [0, m) 中
    }

    int mex = 0;
    while (cnt[mex % m]-- > 0) {
        mex++;
    }

    free(cnt);
    return mex;
}
```

```go [sol-Go]
func findSmallestInteger(nums []int, m int) (mex int) {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[(x%m+m)%m]++ // 保证取模结果在 [0, m) 中
	}

	for cnt[mex%m] > 0 {
		cnt[mex%m]--
		mex++
	}
	return
}
```

```js [sol-JavaScript]
var findSmallestInteger = function(nums, m) {
    const cnt = new Map();
    for (const x of nums) {
        const v = (x % m + m) % m; // 保证取模结果在 [0, m) 中
        cnt.set(v, (cnt.get(v) ?? 0) + 1);
    }

    let mex = 0;
    while ((cnt.get(mex % m) ?? 0) > 0) {
        cnt.set(mex % m, cnt.get(mex % m) - 1);
        mex++;
    }
    return mex;
};
```

```rust [sol-Rust]
use std::collections::HashMap;

impl Solution {
    pub fn find_smallest_integer(nums: Vec<i32>, m: i32) -> i32 {
        let mut cnt = HashMap::new();
        for x in nums {
            // 保证取模结果在 [0, m) 中
            *cnt.entry((x % m + m) % m).or_insert(0) += 1;
        }

        for mex in 0.. {
            if let Some(c) = cnt.get_mut(&(mex % m)) {
                if *c > 0 {
                    *c -= 1;
                    continue;
                }
            }
            return mex;
        }
        unreachable!()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。由于加多少个数，就只能减多少个数，所以第二个循环至多循环 $\mathcal{O}(n)$ 次。
- 空间复杂度：$\mathcal{O}(\min(n,m))$。哈希表中至多有 $\mathcal{O}(\min(n,m))$ 个元素。

## 写法二

![lc2598-c.png](https://pic.leetcode.cn/1759992283-huZjMi-lc2598-c.png){:width=500px}

此外，把哈希表换成数组更快。

```py [sol-Python3]
class Solution:
    def findSmallestInteger(self, nums: List[int], m: int) -> int:
        cnt = [0] * m
        for x in nums:
            cnt[x % m] += 1

        i = cnt.index(min(cnt))
        return m * cnt[i] + i
```

```java [sol-Java]
class Solution {
    public int findSmallestInteger(int[] nums, int m) {
        int[] cnt = new int[m];
        for (int x : nums) {
            cnt[(x % m + m) % m]++;
        }

        int i = 0;
        for (int j = 1; j < m; j++) {
            if (cnt[j] < cnt[i]) {
                i = j;
            }
        }

        return m * cnt[i] + i;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findSmallestInteger(vector<int>& nums, int m) {
        vector<int> cnt(m);
        for (int x : nums) {
            cnt[(x % m + m) % m]++;
        }

        int i = ranges::min_element(cnt) - cnt.begin();
        return m * cnt[i] + i;
    }
};
```

```c [sol-C]
int findSmallestInteger(int* nums, int numsSize, int m) {
    int* cnt = calloc(m, sizeof(int));
    for (int i = 0; i < numsSize; i++) {
        cnt[(nums[i] % m + m) % m]++;
    }

    int i = 0;
    for (int j = 1; j < m; j++) {
        if (cnt[j] < cnt[i]) {
            i = j;
        }
    }

    int ans = m * cnt[i] + i;
    free(cnt);
    return ans;
}
```

```go [sol-Go]
func findSmallestInteger(nums []int, m int) int {
	cnt := make([]int, m)
	for _, x := range nums {
		cnt[(x%m+m)%m]++
	}

	i := 0
	for j := 1; j < m; j++ {
		if cnt[j] < cnt[i] {
			i = j
		}
	}

	return m*cnt[i] + i
}
```

```js [sol-JavaScript]
var findSmallestInteger = function(nums, m) {
    const cnt = Array(m).fill(0);
    for (const x of nums) {
        cnt[(x % m + m) % m]++;
    }

    let i = 0;
    for (let j = 1; j < m; j++) {
        if (cnt[j] < cnt[i]) {
            i = j;
        }
    }

    return m * cnt[i] + i;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_smallest_integer(nums: Vec<i32>, m: i32) -> i32 {
        let mut cnt = vec![0; m as usize];
        for x in nums {
            cnt[((x % m + m) % m) as usize] += 1;
        }

        let mut i = 0;
        for j in 1..m as usize {
            if cnt[j] < cnt[i] {
                i = j;
            }
        }

        m * cnt[i] + i as i32
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+m)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(m)$。

## 相似题目

见下面数学题单的「**§1.9 同余**」。

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
11. [链表、二叉树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA/一般树）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
