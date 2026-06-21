统计每个位置被多少个灯泡照亮，记作 $\textit{cnt}$。那么 $\textit{cnt}[i] = 0$ 的那些位置，就需要额外用灯泡照亮了。

设一个正常工作的灯泡，照亮 $[L,R]$ 中的每个位置。我们把 $\textit{cnt}$ 中的子数组 $[L,R]$ 都加一。如何快速实现区间加一？做法见 [差分数组原理讲解](https://leetcode.cn/problems/car-pooling/solution/suan-fa-xiao-ke-tang-chai-fen-shu-zu-fu-9d4ra/)。计算差分数组的前缀和，就得到了 $\textit{cnt}$ 数组。

我们可以在计算差分数组的前缀和的同时，计算额外安装的灯泡数量。

如果遍历到位置 $i$ 时，发现位置 $i$ 没被照亮，那么贪心地，在 $i+1$ 处安装灯泡，照亮 $[i,i+2]$，更新差分数组。（如果 $i=n-1$，则在 $n-1$ 处安装灯泡）

> **注**：位置 $i+1$ 一定没有灯泡，否则位置 $i$ 一定已被照亮，矛盾。 

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minLights(self, lights: list[int]) -> int:
        n = len(lights)
        diff = [0] * (n + 1)
        for i, v in enumerate(lights):
            if v > 0:
                # 照亮 [max(i-v, 0), min(i+v, n-1)]
                diff[max(i - v, 0)] += 1
                diff[min(i + v + 1, n)] -= 1

        ans = sum_d = 0
        for i in range(n):
            sum_d += diff[i]
            if sum_d == 0:
                # 在 i+1 装一个灯泡，照亮 [i, i+2]
                ans += 1
                sum_d += 1  # diff[i] += 1 直接更新到 sum_d 中
                diff[min(i + 3, n)] -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int minLights(int[] lights) {
        int n = lights.length;
        int[] diff = new int[n + 1];
        for (int i = 0; i < n; i++) {
            int v = lights[i];
            if (v > 0) {
                // 照亮 [max(i-v, 0), min(i+v, n-1)]
                diff[Math.max(i - v, 0)] += 1;
                diff[Math.min(i + v + 1, n)] -= 1;
            }
        }

        int ans = 0;
        int sumD = 0;
        for (int i = 0; i < n; i++) {
            sumD += diff[i];
            if (sumD == 0) {
                // 在 i+1 装一个灯泡，照亮 [i, i+2]
                ans += 1;
                sumD += 1; // diff[i]++ 直接更新到 sumD 中
                diff[Math.min(i + 3, n)] -= 1;
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minLights(vector<int>& lights) {
        int n = lights.size();
        vector<int> diff(n + 1);
        for (int i = 0; i < n; i++) {
            int v = lights[i];
            if (v > 0) {
                // 照亮 [max(i-v, 0), min(i+v, n-1)]
                diff[max(i - v, 0)]++;
                diff[min(i + v + 1, n)]--;
            }
        }

        int ans = 0, sum_d = 0;
        for (int i = 0; i < n; i++) {
            sum_d += diff[i];
            if (sum_d == 0) {
                // 在 i+1 装一个灯泡，照亮 [i, i+2]
                ans++;
                sum_d++; // diff[i]++ 直接更新到 sum_d 中
                diff[min(i + 3, n)]--;
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func minLights(lights []int) (ans int) {
	n := len(lights)
	diff := make([]int, n+1)
	for i, v := range lights {
		if v > 0 {
			// 照亮 [max(i-v, 0), min(i+v, n-1)]
			diff[max(i-v, 0)]++
			diff[min(i+v+1, n)]--
		}
	}

	sumD := 0
	for i, d := range diff[:n] {
		sumD += d
		if sumD == 0 {
			// 在 i+1 装一个灯泡，照亮 [i, i+2]
			ans++
			sumD++ // diff[i]++ 直接更新到 sumD 中
			diff[min(i+3, n)]--
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{lights}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面数据结构题单的「**§2.1 一维差分**」。

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
