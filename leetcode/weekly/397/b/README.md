## k = 1

**从特殊到一般**，先想一想，$k=1$ 怎么做？

此时只能一步一步地向右走。无论起点在哪，终点都是 $n-1$。

如果选择 $i$ 为起点，我们计算的是子数组 $[i,n-1]$ 的元素和，即**后缀和**。

后缀和怎么算？我们可以倒着遍历 $\textit{energy}$，同时累加元素和，即为后缀和。

答案等于**所有后缀和的最大值**。

## k = 2

再想一想，$k=2$ 怎么做？

此时我们有两个终点：$n-2$ 和 $n-1$。

对于终点 $n-1$：

- 如果选择 $n-3$ 为起点，那么我们累加的是下标为 $n-3,n-1$ 的元素和。
- 如果选择 $n-5$ 为起点，那么我们累加的是下标为 $n-5,n-3,n-1$ 的元素和。
- 如果选择 $n-7$ 为起点，那么我们累加的是下标为 $n-7,n-5,n-3,n-1$ 的元素和。
- 一般地，从 $n-1$ 开始倒着遍历，步长为 $-k=-2$，累加元素和，计算元素和的最大值。

对于终点 $n-2$：

- 如果选择 $n-4$ 为起点，那么我们累加的是下标为 $n-4,n-2$ 的元素和。
- 如果选择 $n-6$ 为起点，那么我们累加的是下标为 $n-6,n-4,n-2$ 的元素和。
- 如果选择 $n-8$ 为起点，那么我们累加的是下标为 $n-8,n-6,n-4,n-2$ 的元素和。
- 一般地，从 $n-2$ 开始倒着遍历，步长为 $-k=-2$，累加元素和，计算元素和的最大值。

是否可以从 $n-3$ 开始倒着遍历？

不行，因为 $n-3$ 还可以向右跳到 $n-1$，所以 $n-3$ 不是终点，不能作为倒着遍历的起点。

## 一般情况

枚举终点 $n-k,n-k+1,\dots,n-1$，倒着遍历，步长为 $-k$。

遍历的同时累加元素和 $\textit{sufSum}$，计算 $\textit{sufSum}$ 的最大值，即为答案。

## 写法一

```py [sol-Python3]
class Solution:
    def maximumEnergy(self, energy: List[int], k: int) -> int:
        n = len(energy)
        ans = -inf
        for i in range(n - k, n):  # 枚举终点 i
            suf_sum = accumulate(energy[j] for j in range(i, -1, -k))  # 计算后缀和
            ans = max(ans, max(suf_sum))
        return ans
```

```java [sol-Java]
class Solution {
    public int maximumEnergy(int[] energy, int k) {
        int n = energy.length;
        int ans = Integer.MIN_VALUE;
        for (int i = n - k; i < n; i++) { // 枚举终点 i
            int sufSum = 0;
            for (int j = i; j >= 0; j -= k) {
                sufSum += energy[j]; // 计算后缀和
                ans = Math.max(ans, sufSum);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumEnergy(vector<int>& energy, int k) {
        int n = energy.size();
        int ans = INT_MIN;
        for (int i = n - k; i < n; i++) { // 枚举终点 i
            int suf_sum = 0;
            for (int j = i; j >= 0; j -= k) {
                suf_sum += energy[j]; // 计算后缀和
                ans = max(ans, suf_sum);
            }
        }
        return ans;
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maximumEnergy(int* energy, int energySize, int k){
    int ans = INT_MIN;
    for (int i = energySize - k; i < energySize; i++) { // 枚举终点 i
        int suf_sum = 0;
        for (int j = i; j >= 0; j -= k) {
            suf_sum += energy[j]; // 计算后缀和
            ans = MAX(ans, suf_sum);
        }
    }
    return ans;
}
```

```go [sol-Go]
func maximumEnergy(energy []int, k int) int {
	n := len(energy)
	ans := math.MinInt
	for i := n - k; i < n; i++ { // 枚举终点 i
		sufSum := 0
		for j := i; j >= 0; j -= k {
			sufSum += energy[j] // 计算后缀和
			ans = max(ans, sufSum)
		}
	}
	return ans
}
```

```js [sol-JavaScript]
var maximumEnergy = function(energy, k) {
    const n = energy.length;
    let ans = -Infinity;
    for (let i = n - k; i < n; i++) { // 枚举终点 i
        let sufSum = 0;
        for (let j = i; j >= 0; j -= k) {
            sufSum += energy[j]; // 计算后缀和
            ans = Math.max(ans, sufSum);
        }
    }
    return ans;
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_energy(energy: Vec<i32>, k: i32) -> i32 {
        let n = energy.len();
        let k = k as usize;
        let mut ans = i32::MIN;
        for i in n - k..n { // 枚举终点 i
            let mut suf_sum = 0;
            for j in (0..=i).rev().step_by(k) {
                suf_sum += energy[j]; // 计算后缀和
                ans = ans.max(suf_sum);
            }
        }
        ans
    }
}
```

## 写法二

原地计算后缀和，把后缀和保存到 $\textit{energy}$ 中。

最后返回 $\textit{energy}$ 的最大值，即为所有后缀和的最大值。

```py [sol-Python3]
class Solution:
    def maximumEnergy(self, energy: List[int], k: int) -> int:
        for i in range(len(energy) - k - 1, -1, -1):
            energy[i] += energy[i + k]
        return max(energy)
```

```java [sol-Java]
class Solution {
    public int maximumEnergy(int[] energy, int k) {
        for (int i = energy.length - k - 1; i >= 0; i--) {
            energy[i] += energy[i + k];
        }
        return Arrays.stream(energy).max().getAsInt();
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int maximumEnergy(vector<int>& energy, int k) {
        int n = energy.size();
        for (int i = n - k - 1; i >= 0; i--) {
            energy[i] += energy[i + k];
        }
        return ranges::max(energy);
    }
};
```

```c [sol-C]
#define MAX(a, b) ((b) > (a) ? (b) : (a))

int maximumEnergy(int* energy, int energySize, int k) {
    int ans = INT_MIN;
    for (int i = energySize - 1; i >= 0; i--) {
        if (i + k < energySize) {
            energy[i] += energy[i + k];
        }
        ans = MAX(ans, energy[i]);
    }
    return ans;
}
```

```go [sol-Go]
func maximumEnergy(energy []int, k int) int {
	for i := len(energy) - k - 1; i >= 0; i-- {
		energy[i] += energy[i+k]
	}
	return slices.Max(energy)
}
```

```js [sol-JavaScript]
var maximumEnergy = function(energy, k) {
    for (let i = energy.length - k - 1; i >= 0; i--) {
        energy[i] += energy[i + k];
    }
    return Math.max(...energy);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn maximum_energy(mut energy: Vec<i32>, k: i32) -> i32 {
        let k = k as usize;
        for i in (0..energy.len() - k).rev() {
            energy[i] += energy[i + k];
        }
        *energy.iter().max().unwrap()
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{energy}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面贪心与思维题单的「**§5.3 逆向思维**」。

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
