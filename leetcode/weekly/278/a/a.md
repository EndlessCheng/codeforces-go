## 题意

返回不在 $\textit{nums}$ 中的最小的 $\textit{original}\cdot 2^k$，其中 $k$ 是非负整数。

## 方法一：枚举

枚举 $k=0,1,2,\ldots$ 判断 $\textit{original}\cdot 2^k$ 是否在 $\textit{nums}$ 中。

用哈希集合记录 $\textit{nums}$ 的每个元素可以加快判断。

```py [sol-Python3]
class Solution:
    def findFinalValue(self, nums: List[int], original: int) -> int:
        st = set(nums)
        while original in st:
            original *= 2
        return original
```

```java [sol-Java]
class Solution {
    public int findFinalValue(int[] nums, int original) {
        Set<Integer> st = new HashSet<>();
        for (int x : nums) {
            st.add(x);
        }

        while (st.contains(original)) {
            original *= 2;
        }
        return original;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findFinalValue(vector<int>& nums, int original) {
        unordered_set<int> st(nums.begin(), nums.end());
        while (st.contains(original)) {
            original *= 2;
        }
        return original;
    }
};
```

```go [sol-Go]
func findFinalValue(nums []int, original int) int {
	has := map[int]bool{}
	for _, x := range nums {
		has[x] = true
	}

	for has[original] {
		original *= 2
	}
	return original
}
```

```js [sol-JavaScript]
var findFinalValue = function(nums, original) {
    const st = new Set(nums);
    while (st.has(original)) {
        original *= 2;
    }
    return original;
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn find_final_value(nums: Vec<i32>, mut original: i32) -> i32 {
        let st = nums.into_iter().collect::<HashSet<_>>();
        while st.contains(&original) {
            original *= 2;
        }
        original
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(n+\log\dfrac{U}{\textit{original}}\right)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。循环次数为 $\mathcal{O}\left(\log\dfrac{U}{\textit{original}}\right)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 方法二：只记录所有可能值

哈希集合记录的元素可以更少，只需要记录符合 $\textit{original}\cdot 2^k$ 的元素。

设 $x = \textit{nums}[i]$，如果 $x = \textit{original}\cdot 2^k$，那么：

- $x$ 是 $\textit{original}$ 的倍数。
- $\dfrac{x}{\textit{original}}$ 是 [2 的幂](https://leetcode.cn/problems/power-of-two/)，见 [我的题解](https://leetcode.cn/problems/power-of-two/solutions/2973442/yan-ge-zheng-ming-yi-xing-xie-fa-pythonj-h04o/)。

```py [sol-Python3]
class Solution:
    def findFinalValue(self, nums: List[int], original: int) -> int:
        st = set()
        for x in nums:
            k, r = divmod(x, original)
            if r == 0 and k & (k - 1) == 0:
                st.add(x)

        while original in st:
            original *= 2
        return original
```

```java [sol-Java]
class Solution {
    public int findFinalValue(int[] nums, int original) {
        Set<Integer> st = new HashSet<>();
        for (int x : nums) {
            int k = x / original;
            if (x % original == 0 && (k & (k - 1)) == 0) {
                st.add(x);
            }
        }

        while (st.contains(original)) {
            original *= 2;
        }
        return original;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findFinalValue(vector<int>& nums, int original) {
        unordered_set<int> st;
        for (int x : nums) {
            int k = x / original;
            if (x % original == 0 && (k & (k - 1)) == 0) {
                st.insert(x);
            }
        }

        while (st.contains(original)) {
            original *= 2;
        }
        return original;
    }
};
```

```go [sol-Go]
func findFinalValue(nums []int, original int) int {
	has := map[int]bool{}
	for _, x := range nums {
		k := x / original
		if x%original == 0 && k&(k-1) == 0 {
			has[x] = true
		}
	}
	for has[original] {
		original *= 2
	}
	return original
}
```

```js [sol-JavaScript]
var findFinalValue = function(nums, original) {
    const st = new Set();
    for (const x of nums) {
        const k = x / original;
        if (x % original === 0 && (k & (k - 1)) === 0) {
            st.add(x);
        }
    }

    while (st.has(original)) {
        original *= 2;
    }
    return original;
};
```

```rust [sol-Rust]
use std::collections::HashSet;

impl Solution {
    pub fn find_final_value(nums: Vec<i32>, mut original: i32) -> i32 {
        let st = nums.into_iter()
            .filter(|x| x % original == 0 && ((x / original) & (x / original - 1)) == 0)
            .collect::<HashSet<_>>();
        while st.contains(&original) {
            original *= 2;
        }
        original
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}\left(n+\log\dfrac{U}{\textit{original}}\right)$，其中 $n$ 是 $\textit{nums}$ 的长度，$U=\max(\textit{nums})$。
- 空间复杂度：$\mathcal{O}\left(\log\dfrac{U}{\textit{original}}\right)$。

## 方法三：位运算

改成记录 $\textit{original}\cdot 2^k$ 中的 $k$。

由于 $k$ 很小，我们可以把 $k$ 保存到一个二进制数 $\textit{mask}$ 中，具体请看 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)

答案中的 $k$ 就是 $\textit{mask}$ 的最低 $0$ 的位置。

这可以通过位运算 $\mathcal{O}(1)$ 地计算出来：把 $\textit{mask}$ 取反，找最低位的 $1$，其对应的二进制数 $\textit{lowbit}$ 即为答案中的 $2^k$。再乘以 $\textit{original}$，得到最终答案。

```py [sol-Python3]
class Solution:
    def findFinalValue(self, nums: List[int], original: int) -> int:
        mask = 0
        for x in nums:
            k, r = divmod(x, original)
            if r == 0 and k & (k - 1) == 0:
                mask |= k

        # 找最低的 0，等价于取反后，找最低的 1（lowbit）
        mask = ~mask
        return original * (mask & -mask)
```

```java [sol-Java]
class Solution {
    public int findFinalValue(int[] nums, int original) {
        int mask = 0;
        for (int x : nums) {
            int k = x / original;
            if (x % original == 0 && (k & (k - 1)) == 0) {
                mask |= k;
            }
        }

        // 找最低的 0，等价于取反后，找最低的 1（lowbit）
        mask = ~mask;
        return original * (mask & -mask);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int findFinalValue(vector<int>& nums, int original) {
        int mask = 0;
        for (int x : nums) {
            int k = x / original;
            if (x % original == 0 && (k & (k - 1)) == 0) {
                mask |= k;
            }
        }

        // 找最低的 0，等价于取反后，找最低的 1（lowbit）
        mask = ~mask;
        return original * (mask & -mask);
    }
};
```

```go [sol-Go]
func findFinalValue(nums []int, original int) int {
	mask := 0
	for _, x := range nums {
		k := x / original
		if x%original == 0 && k&(k-1) == 0 {
			mask |= k
		}
	}

	// 找最低的 0，等价于取反后，找最低的 1（lowbit）
	mask = ^mask
	return original * (mask & -mask)
}
```

```js [sol-JavaScript]
var findFinalValue = function(nums, original) {
    let mask = 0;
    for (const x of nums) {
        const k = x / original;
        if (x % original === 0 && (k & (k - 1)) === 0) {
            mask |= k;
        }
    }

    // 找最低的 0，等价于取反后，找最低的 1（lowbit）
    mask = ~mask;
    return original * (mask & -mask);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn find_final_value(nums: Vec<i32>, original: i32) -> i32 {
        let mut mask = 0;
        for x in nums {
            let k = x / original;
            if x % original == 0 && (k & (k - 1)) == 0 {
                mask |= k;
            }
        }

        // 找最低的 0，等价于取反后，找最低的 1（lowbit）
        mask = !mask;
        original * (mask & -mask)
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 相似题目

[2568. 最小无法得到的或值](https://leetcode.cn/problems/minimum-impossible-or/)

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
