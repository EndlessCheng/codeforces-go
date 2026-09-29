根据题意，在 $\textit{nums}_1[i]$ 上 $+1$，等价于在 $\textit{nums}_2[i]$ 上 $-1$，反之亦然。

定义 $a[i]=|\textit{nums}_1[i]-\textit{nums}_2[i]|$，$k=k_1+k_2$。原问题等价于：

- 对数组 $a$ 执行至多 $k$ 次 $-1$ 操作后，$\sum\limits_{i} a[i]^2$ 的最小值。

对于 $a$ 中的两个数，先把大的 $-1$ 会更优。

把 $a$ 从大到小排序，然后从左到右遍历 $a$，同时更新剩余操作次数 $k$。遍历至 $a[i]$ 时，$a[0]$ 到 $a[i-1]$ 均已减小至 $a[i]$，我们需要判断能否用 $k$ 次操作把 $a[0]$ 到 $a[i]$ 全部减小至 $a[i+1]$，即比较 $k$ 与所需操作次数 $c = (i + 1)(a[i] - a[i+1])$ 的大小：

- 如果 $c<k$，那么从 $a[0]$ 到 $a[i]$ 均可以减小至 $a[i+1]$，更新 $k=k-c$。
- 如果 $c\ge k$，那么从 $a[0]$ 到 $a[i]$ 中：
    - 有 $k \bmod (i+1)$ 个元素可以额外减小 $\left\lfloor\dfrac{k}{i+1}\right\rfloor+1$；
    - 有 $i+1-k \bmod (i+1)$ 个元素可以额外减小 $\left\lfloor\dfrac{k}{i+1}\right\rfloor$。
    - 后续无法继续减小，应退出循环。

代码实现时，可以在 $a$ 末尾加一个哨兵 $0$，这样无需判断下标 $i+1$ 是否越界。

```py [sol-Python3]
class Solution:
    def minSumSquareDiff(self, a: List[int], nums2: List[int], k1: int, k2: int) -> int:
        ans, k = 0, k1 + k2
        for i in range(len(a)):
            a[i] = abs(a[i] - nums2[i])
            ans += a[i] * a[i]
        if sum(a) <= k:
            return 0  # 所有 a[i] 均可为 0

        a.sort(reverse=True)
        a.append(0)  # 哨兵
        for i, v in enumerate(a):
            ans -= v * v  # 撤销上面的 ans += a[i] * a[i]
            j = i + 1
            c = j * (v - a[j])
            if c < k:
                k -= c
                continue
            v -= k // j
            return ans + k % j * (v - 1) * (v - 1) + (j - k % j) * v * v
```

```java [sol-Java]
class Solution {
    public long minSumSquareDiff(int[] a, int[] nums2, int k1, int k2) {
        int n = a.length;
        int k = k1 + k2;
        long ans = 0;
        long sum = 0;
        for (int i = 0; i < n; i++) {
            a[i] = Math.abs(a[i] - nums2[i]);
            sum += a[i];
            ans += (long) a[i] * a[i];
        }
        if (sum <= k) {
            return 0; // 所有 a[i] 均可为 0
        }

        Arrays.sort(a);
        for (int i = n - 1; ; i--) {
            int m = n - i;
            long v = a[i];
            long c = m * (v - (i > 0 ? a[i - 1] : 0));
            ans -= v * v; // 撤销上面的 ans += a[i] * a[i]
            if (c < k) {
                k -= c;
                continue;
            }
            v -= k / m;
            return ans + k % m * (v - 1) * (v - 1) + (m - k % m) * v * v;
        }
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minSumSquareDiff(vector<int>& a, vector<int>& nums2, int k1, int k2) {
        int n = a.size(), k = k1 + k2;
        long long ans = 0, sum = 0;
        for (int i = 0; i < n; i++) {
            a[i] = abs(a[i] - nums2[i]);
            sum += a[i];
            ans += 1LL * a[i] * a[i];
        }
        if (sum <= k) { // 所有 a[i] 均可为 0
            return 0;
        }

        ranges::sort(a, greater());
        a.push_back(0); // 哨兵
        for (int i = 0; ; i++) {
            long long v = a[i];
            ans -= v * v; // 撤销上面的 ans += a[i] * a[i]
            int j = i + 1;
            long long c = j * (v - a[j]);
            if (c < k) {
                k -= c;
                continue;
            }
            v -= k / j;
            return ans + k % j * (v - 1) * (v - 1) + (j - k % j) * v * v;
        }
    }
};
```

```go [sol-Go]
func minSumSquareDiff(a, nums2 []int, k1, k2 int) int64 {
	ans, sum := 0, 0
	for i, v := range a {
		a[i] = abs(v - nums2[i])
		sum += a[i]
		ans += a[i] * a[i]
	}
	k := k1 + k2
	if sum <= k {
		return 0 // 所有 a[i] 均可为 0
	}

	slices.SortFunc(a, func(a, b int) int { return b - a })
	a = append(a, 0) // 哨兵
	for i, v := range a {
		i++
		ans -= v * v // 撤销上面的 ans += a[i] * a[i]
		if c := i * (v - a[i]); c < k {
			k -= c
			continue
		}
		v -= k / i
		ans += k%i*(v-1)*(v-1) + (i-k%i)*v*v
		break
	}
	return int64(ans)
}

func abs(x int) int { if x < 0 { return -x }; return x }
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}_i$ 的长度。瓶颈在排序上。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销和哨兵的开销。

## 分类题单

[如何科学刷题？](https://leetcode.cn/discuss/post/3141566/ru-he-ke-xue-shua-ti-by-endlesscheng-q3yd/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/discuss/post/3578981/ti-dan-hua-dong-chuang-kou-ding-chang-bu-rzz7/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/discuss/post/3579164/ti-dan-er-fen-suan-fa-er-fen-da-an-zui-x-3rqn/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/discuss/post/3579480/ti-dan-dan-diao-zhan-ju-xing-xi-lie-zi-d-u4hk/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/discuss/post/3580195/fen-xiang-gun-ti-dan-wang-ge-tu-dfsbfszo-l3pa/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/discuss/post/3580371/fen-xiang-gun-ti-dan-wei-yun-suan-ji-chu-nth4/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/discuss/post/3581143/fen-xiang-gun-ti-dan-tu-lun-suan-fa-dfsb-qyux/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/discuss/post/3581838/fen-xiang-gun-ti-dan-dong-tai-gui-hua-ru-007o/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/discuss/post/3583665/fen-xiang-gun-ti-dan-chang-yong-shu-ju-j-bvmv/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/discuss/post/3584388/fen-xiang-gun-ti-dan-shu-xue-suan-fa-shu-gcai/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/discuss/post/3091107/fen-xiang-gun-ti-dan-tan-xin-ji-ben-tan-k58yb/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/discuss/post/3142882/fen-xiang-gun-ti-dan-lian-biao-er-cha-sh-6srp/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/discuss/post/3144832/fen-xiang-gun-ti-dan-zi-fu-chuan-kmpzhan-ugt4/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)

欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
