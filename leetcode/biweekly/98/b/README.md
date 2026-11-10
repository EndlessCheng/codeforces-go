为方便计算，先把 $\textit{nums}$ 从小到大排序。

如果不修改，那么最大得分为 $\textit{nums}[n-1] - \textit{nums}[0]$。

只改一个数呢？如果最小值和最大值都不改，那么上式不变。所以要改最小值或者最大值。

推广这个想法，在可以改两个数的情况下，只需考虑改前二小和前二大中的数：

- 改最大的两个数为 $\textit{nums}[n-3]$，最大得分为 $\textit{nums}[n-3]-\textit{nums}[0]$。
- 改最小的为 $\textit{nums}[1]$，最大的为 $\textit{nums}[n-2]$，最大得分为 $\textit{nums}[n-2]-\textit{nums}[1]$。
- 改最小的两个数为 $\textit{nums}[2]$，最大得分为 $\textit{nums}[n-1]-\textit{nums}[2]$。

取这三种情况的最小值，即最大得分的最小值。

由于这三种情况都会产生相同的数，所以最小得分为 $0$。

综上所述，答案为

$$
\min(\textit{nums}[n-3]-\textit{nums}[0], \textit{nums}[n-2]-\textit{nums}[1],\textit{nums}[n-1]-\textit{nums}[2])
$$

[视频讲解](https://www.bilibili.com/video/BV15D4y1G7ms/) 第二题。

```py [sol-Python3]
class Solution:
    def minimizeSum(self, a: List[int]) -> int:
        a.sort()
        return min(a[-3] - a[0], a[-2] - a[1], a[-1] - a[2])
```

```java [sol-Java]
class Solution {
    public int minimizeSum(int[] a) {
        Arrays.sort(a);
        int n = a.length;
        return Math.min(Math.min(a[n - 3] - a[0], a[n - 2] - a[1]), a[n - 1] - a[2]);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minimizeSum(vector<int>& a) {
        ranges::sort(a);
        int n = a.size();
        return min({a[n - 3] - a[0], a[n - 2] - a[1], a[n - 1] - a[2]});
    }
};
```

```c [sol-C]
int cmp(const void* a, const void* b) {
    return *(int*)a - *(int*)b;
}

int minimizeSum(int* a, int n) {
    qsort(a, n, sizeof(int), cmp);
    return MIN(MIN(a[n - 3] - a[0], a[n - 2] - a[1]), a[n - 1] - a[2]);
}
```

```go [sol-Go]
func minimizeSum(a []int) int {
	slices.Sort(a)
	n := len(a)
	return min(a[n-3]-a[0], a[n-2]-a[1], a[n-1]-a[2])
}
```

```js [sol-JavaScript]
var minimizeSum = function(a) {
    a.sort((x, y) => x - y);
    const n = a.length;
    return Math.min(Math.min(a[n - 3] - a[0], a[n - 2] - a[1]), a[n - 1] - a[2]);
};
```

```rust [sol-Rust]
impl Solution {
    pub fn minimize_sum(mut a: Vec<i32>) -> i32 {
        a.sort_unstable();
        let n = a.len();
        (a[n - 3] - a[0]).min(a[n - 2] - a[1]).min(a[n - 1] - a[2])
    }
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{nums}$ 的长度。**注**：手动维护（或者快速选择算法）可以做到 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(1)$。忽略排序时栈的开销。

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
