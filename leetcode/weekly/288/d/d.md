### 提示 1

枚举种花后，**完善的**花园的数目。

### 提示 2

为了留下更多的花去填充**不完善**的花园，我们应当选择那些花朵更多的花园去变成完善的。

### 提示 3

将 $\textit{flowers}$ 从小到大排序，从而方便枚举哪些是完善的花园。

那么剩下要解决的问题就是，怎么最大化其余花园的花的最小数目。

### 提示 4

枚举 $\textit{flowers}$ 的后缀，让这些花园的花增加至 $\textit{target}$，同时我们需要求出 $\textit{flowers}$ 的最长前缀（设其长为 $x$），满足前缀中的花园的花**都**能填充至**至少** $\textit{flowers}[x-1]$ 朵。（这可以用二分或双指针来实现，下面代码用的双指针）

设在填充后缀之后，剩余 $\textit{leftFlowers}$ 朵花可以种植，且长为 $x$ 的前缀一共有 $\textit{sumFlowers}$ 朵花（这里的 $x$ 需要满足上面填充的要求）。那么在填充后，这 $x$ 个花园一共有 $\textit{leftFlowers}+\textit{sumFlowers}$ 朵花。由于**最小值不会超过平均值**，在均匀分配的情况下，最小值的最大值可以为平均值的下取整，即

$$
\left\lfloor\dfrac{\textit{leftFlowers}+\textit{sumFlowers}}{x}\right\rfloor
$$

注意这个值不能超过 $\textit{target}-1$，否则不满足题目「不完善」的要求。

按照上述方法，枚举后缀的同时计算出对应的最长前缀，及其最小值的最大值，进而计算出对应的总美丽值。所有总美丽值的最大值即为答案。

```py [sol1-Python3]
class Solution:
    def maximumBeauty(self, flowers: List[int], newFlowers: int, target: int, full: int, partial: int) -> int:
        flowers.sort()
        n = len(flowers)
        if flowers[0] >= target:  # 剪枝，此时所有花园都是完善的
            return n * full

        leftFlowers = newFlowers - target * n  # 填充后缀后，剩余可以种植的花
        for i in range(n):
            flowers[i] = min(flowers[i], target)  # 去掉多余的花
            leftFlowers += flowers[i]  # 补上已有的花

        ans, x, sumFlowers = 0, 0, 0
        for i in range(n + 1):  # 枚举后缀长度 n-i
            if leftFlowers >= 0:
                # 计算最长前缀的长度
                while x < i and flowers[x] * x - sumFlowers <= leftFlowers:
                    sumFlowers += flowers[x]
                    x += 1  # 注意 x 只增不减，二重循环的时间复杂度为 O(n)
                beauty = (n - i) * full  # 计算总美丽值
                if x:
                    beauty += min((leftFlowers + sumFlowers) // x, target - 1) * partial
                ans = max(ans, beauty)
            if i < n:
                leftFlowers += target - flowers[i]
        return ans
```

```java [sol1-Java]
class Solution {
    public long maximumBeauty(int[] flowers, long newFlowers, int target, int full, int partial) {
        Arrays.sort(flowers);
        long n = flowers.length;
        if (flowers[0] >= target) return n * full; // 剪枝，此时所有花园都是完善的

        var leftFlowers = newFlowers - target * n; // 填充后缀后，剩余可以种植的花
        for (var i = 0; i < n; ++i) {
            flowers[i] = Math.min(flowers[i], target); // 去掉多余的花
            leftFlowers += flowers[i]; // 补上已有的花
        }

        var ans = 0L;
        var sumFlowers = 0L;
        for (int i = 0, x = 0; i <= n; ++i) { // 枚举后缀长度 n-i
            if (leftFlowers >= 0) {
                // 计算最长前缀的长度
                while (x < i && (long) flowers[x] * x - sumFlowers <= leftFlowers)
                    sumFlowers += flowers[x++]; // 注意 x 只增不减，二重循环的时间复杂度为 O(n)
                var beauty = (n - i) * full; // 计算总美丽值
                if (x > 0) beauty += Math.min((leftFlowers + sumFlowers) / x, (long) target - 1) * partial;
                ans = Math.max(ans, beauty);
            }
            if (i < n) leftFlowers += target - flowers[i];
        }
        return ans;
    }
}
```

```C++ [sol1-C++]
class Solution {
public:
    long long maximumBeauty(vector<int> &flowers, long long newFlowers, int target, int full, int partial) {
        sort(flowers.begin(), flowers.end());
        long n = flowers.size();
        if (flowers[0] >= target) return n * full; // 剪枝，此时所有花园都是完善的

        long leftFlowers = newFlowers - target * n; // 填充后缀后，剩余可以种植的花
        for (int i = 0; i < n; ++i) {
            flowers[i] = min(flowers[i], target); // 去掉多余的花
            leftFlowers += flowers[i]; // 补上已有的花
        }

        long ans = 0L, sumFlowers = 0L;
        for (int i = 0, x = 0; i <= n; ++i) { // 枚举后缀长度 n-i
            if (leftFlowers >= 0) {
                // 计算最长前缀的长度
                while (x < i && (long) flowers[x] * x - sumFlowers <= leftFlowers)
                    sumFlowers += flowers[x++]; // 注意 x 只增不减，二重循环的时间复杂度为 O(n)
                long beauty = (n - i) * full; // 计算总美丽值
                if (x) beauty += min((leftFlowers + sumFlowers) / x, (long) target - 1) * partial;
                ans = max(ans, beauty);
            }
            if (i < n) leftFlowers += target - flowers[i];
        }
        return ans;
    }
};
```

```go [sol1-Go]
func maximumBeauty(flowers []int, newFlowers int64, target, full, partial int) int64 {
	sort.Ints(flowers)
	n := len(flowers)
	if flowers[0] >= target { // 剪枝，此时所有花园都是完善的
		return int64(n * full)
	}

	leftFlowers := int(newFlowers) - target*n // 填充后缀后，剩余可以种植的花
	for i, f := range flowers {
		flowers[i] = min(f, target) // 去掉多余的花
		leftFlowers += flowers[i] // 补上已有的花
	}

	ans := 0
	for i, x, sumFlowers := 0, 0, 0; i <= n; i++ { // 枚举后缀长度 n-i
		if leftFlowers >= 0 {
			// 计算最长前缀的长度
			for ; x < i && flowers[x]*x-sumFlowers <= leftFlowers; x++ {
				sumFlowers += flowers[x] // 注意 x 只增不减，二重循环的时间复杂度为 O(n)
			}
			beauty := (n - i) * full // 计算总美丽值
			if x > 0 {
				beauty += min((leftFlowers+sumFlowers)/x, target-1) * partial
			}
			ans = max(ans, beauty)
		}
		if i < n {
			leftFlowers += target - flowers[i]
		}
	}
	return int64(ans)
}
```

#### 复杂度分析

- 时间复杂度：$O(n\log n)$，其中 $n$ 为 $\textit{flowers}$ 的长度。瓶颈在排序上。
- 空间复杂度：$O(1)$，仅用到若干变量。如果考虑快排时的栈开销，则空间复杂度为 $O(\log n)$。

## 分类题单

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（矩形系列/字典序最小/贡献法）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)

更多题单，点我个人主页 - 讨论发布。

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)

[往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
