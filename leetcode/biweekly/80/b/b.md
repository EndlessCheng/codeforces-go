对于正整数，$xy\ge\textit{success}$ 等价于 $y\ge\left\lceil\dfrac{\textit{success}}{x}\right\rceil$。

为了方便二分，可以利用如下等式： 

$$
\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a+b-1}{b}\right\rfloor = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
$$

讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。

根据上式，我们有

$$
y\ge\left\lceil\dfrac{\textit{success}}{x}\right\rceil = \left\lfloor\dfrac{\textit{success}-1}{x}\right\rfloor + 1
$$ 

这等价于 

$$
y>\left\lfloor\dfrac{\textit{success}-1}{x}\right\rfloor
$$

对 $\textit{potions}$ 排序后，就可以二分查找了：设 $x=\textit{spells}[i]$，$j$ 是最小的满足 $\textit{potions}[j]>\left\lfloor\dfrac{\textit{success}-1}{x}\right\rfloor$ 的下标，由于数组已经排序，那么下标大于 $j$ 的也同样满足该式，这一共有 $m-j$ 个，其中 $m$ 是 $\textit{potions}$ 的长度。

为什么不等式一定要这样变形？好处是每次二分只需要做一次除法，避免多次在二分循环内做乘法，效率更高。另外的好处是部分语言可以直接调用库函数二分。

有关二分查找的原理，请看[【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

```Python [sol-Python3]
class Solution:
    def successfulPairs(self, spells: List[int], potions: List[int], success: int) -> List[int]:
        potions.sort()
        m = len(potions)
        success -= 1
        return [m - bisect_right(potions, success // x) for x in spells]
```

```java [sol-Java]
class Solution {
    public int[] successfulPairs(int[] spells, int[] potions, long success) {
        Arrays.sort(potions);
        for (int i = 0; i < spells.length; i++) {
            long target = (success - 1) / spells[i];
            if (target < potions[potions.length - 1]) { // 防止 long 转成 int 截断
                spells[i] = potions.length - upperBound(potions, (int) target);
            } else {
                spells[i] = 0;
            }
        }
        return spells;
    }

    // 直接二分 long 是 28ms，改成 int 是 26ms
    private int upperBound(int[] nums, int target) {
        int left = -1, right = nums.length; // 开区间 (left, right)
        while (left + 1 < right) { // 区间不为空
            // 循环不变量：
            // nums[left] <= target
            // nums[right] > target
            int mid = (left + right) >>> 1;
            if (nums[mid] > target) {
                right = mid; // 二分范围缩小到 (left, mid)
            } else {
                left = mid; // 二分范围缩小到 (mid, right)
            }
        }
        return right;
    }
}
```

```C++ [sol-C++]
class Solution {
public:
    vector<int> successfulPairs(vector<int> &spells, vector<int> &potions, long long success) {
        ranges::sort(potions);
        for (int &x : spells) {
            long long target = (success - 1) / x;
            if (target < potions.back()) {
                // 这样写每次二分就只用 int 比较，避免把 potions 中的元素转成 long long 比较
                x = potions.end() - ranges::upper_bound(potions, (int) target);
            } else {
                x = 0;
            }
        }
        return spells;
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

// 返回 nums 中的第一个大于 target 的元素下标
int upperBound(int* nums, int numsSize, int target) {
    int left = -1, right = numsSize; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        int mid = left + (right - left) / 2;
        if (nums[mid] > target) {
            right = mid; // 二分范围缩小到 (left, mid)
        } else {
            left = mid; // 二分范围缩小到 (mid, right)
        }
    }
    return right;
}

int* successfulPairs(int* spells, int spellsSize, int* potions, int potionsSize, long long success, int* returnSize) {
    qsort(potions, potionsSize, sizeof(int), cmp);
    for (int i = 0; i < spellsSize; i++) {
        long long target = (success - 1) / spells[i];
        if (target < potions[potionsSize - 1]) {
            // 这样写每次二分就只用 int 比较，避免把 potions 中的元素转成 long long 比较
            spells[i] = potionsSize - upperBound(potions, potionsSize, target);
        } else {
            spells[i] = 0;
        }
    }
    *returnSize = spellsSize;
    return spells;
}
```

```go [sol-Go]
func successfulPairs(spells, potions []int, success int64) []int {
	slices.Sort(potions)
	for i, x := range spells {
		spells[i] = len(potions) - sort.SearchInts(potions, (int(success)-1)/x+1)
	}
	return spells
}
```

```js [sol-JS]
var successfulPairs = function(spells, potions, success) {
    potions.sort((a, b) => a - b);
    for (let i = 0; i < spells.length; i++) {
        const target = Math.ceil(success / spells[i]);
        spells[i] = potions.length - lowerBound(potions, target);
    }
    return spells;
};

var lowerBound = function(nums, target) {
    let left = -1, right = nums.length; // 开区间 (left, right)
    while (left + 1 < right) { // 区间不为空
        // 循环不变量：
        // nums[left] < target
        // nums[right] >= target
        const mid = left + ((right - left) >> 1);
        if (nums[mid] >= target) {
            right = mid; // 范围缩小到 (left, mid)
        } else {
            left = mid; // 范围缩小到 (mid, right)
        }
    }
    return right;
}
```

```js [sol-JS lodash]
var successfulPairs = function(spells, potions, success) {
    potions.sort((a, b) => a - b);
    for (let i = 0; i < spells.length; i++) {
        const target = Math.ceil(success / spells[i]);
        spells[i] = potions.length - _.sortedIndex(potions, target);
    }
    return spells;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn successful_pairs(mut spells: Vec<i32>, mut potions: Vec<i32>, success: i64) -> Vec<i32> {
        potions.sort_unstable();
        let last = potions[potions.len() - 1] as i64;
        for x in spells.iter_mut() {
            let target = (success - 1) / *x as i64;
            if target < last { // 防止 i64 转成 i32 截断（这样不需要把 potions 转成 i64 比较）
                let j = potions.partition_point(|&x| x <= target as i32);
                *x = (potions.len() - j) as i32;
            } else {
                *x = 0;
            }
        }
        spells
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}((n+m)\log m)$，其中 $n$ 为 $\textit{spells}$ 的长度，$m$ 为 $\textit{potions}$ 的长度。排序 $\mathcal{O}(m\log m)$。二分 $n$ 次，每次 $\mathcal{O}(\log m)$。
- 空间复杂度：$\mathcal{O}(1)$。返回值不计入。忽略排序的栈开销，仅用到若干额外变量。

## 思考题

把乘法改成**异或**要怎么做？

这题是 [1803. 统计异或值在范围内的数对有多少](https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/)，做法见 [我的题解](https://leetcode.cn/problems/count-pairs-with-xor-in-a-range/solution/bu-hui-zi-dian-shu-zhi-yong-ha-xi-biao-y-p2pu/)。

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
