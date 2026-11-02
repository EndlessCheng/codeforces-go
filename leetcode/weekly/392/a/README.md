教程：[【简单题杀手】分组循环](https://leetcode.cn/problems/longest-alternating-subarray/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-r57bz/)

本题和教程中的情况一样，第一个单调序列末尾和第二个单调序列开头，有一个元素是**重叠**的，所以下面代码在外层循环末尾要把 $i$ 减一。

[本题视频讲解](https://www.bilibili.com/video/BV1ut421H7Wv/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def longestMonotonicSubarray(self, a: List[int]) -> int:
        ans = 1
        i, n = 0, len(a)
        while i < n - 1:
            if a[i + 1] == a[i]:
                i += 1  # 直接跳过
                continue
            i0 = i  # 记录这一组的开始位置
            inc = a[i + 1] > a[i]  # 定下基调：是严格递增还是严格递减
            i += 2  # i 和 i+1 已经满足要求，从 i+2 开始判断
            while i < n and a[i] != a[i - 1] and (a[i] > a[i - 1]) == inc:
                i += 1
            # 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = max(ans, i - i0)
            i -= 1
        return ans
```

```java [sol-Java]
class Solution {
    public int longestMonotonicSubarray(int[] a) {
        int ans = 1;
        int i = 0, n = a.length;
        while (i < n - 1) {
            if (a[i + 1] == a[i]) {
                i++; // 直接跳过
                continue;
            }
            int i0 = i; // 记录这一组的开始位置
            boolean inc = a[i + 1] > a[i]; // 定下基调：是严格递增还是严格递减
            i += 2; // i 和 i+1 已经满足要求，从 i+2 开始判断
            while (i < n && a[i] != a[i - 1] && (a[i] > a[i - 1]) == inc) {
                i++;
            }
            // 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = Math.max(ans, i - i0);
            i--;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int longestMonotonicSubarray(vector<int> &a) {
        int ans = 1;
        int i = 0, n = a.size();
        while (i < n - 1) {
            if (a[i + 1] == a[i]) {
                i++; // 直接跳过
                continue;
            }
            int i0 = i; // 记录这一组的开始位置
            bool inc = a[i + 1] > a[i]; // 定下基调：是严格递增还是严格递减
            i += 2; // i 和 i+1 已经满足要求，从 i+2 开始判断
            while (i < n && a[i] != a[i - 1] && (a[i] > a[i - 1]) == inc) {
                i++;
            }
            // 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
            ans = max(ans, i - i0);
            i--;
        }
        return ans;
    }
};
```

```go [sol-Go]
func longestMonotonicSubarray(a []int) int {
	ans := 1
	i, n := 0, len(a)
	for i < n-1 {
		if a[i+1] == a[i] {
			i++ // 直接跳过
			continue
		}
		i0 := i              // 记录这一组的开始位置
		inc := a[i+1] > a[i] // 定下基调：是严格递增还是严格递减
		i += 2               // i 和 i+1 已经满足要求，从 i+2 开始判断
		for i < n && a[i] != a[i-1] && a[i] > a[i-1] == inc {
			i++
		}
		// 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-i0)
		i--
	}
	return ans
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。时间复杂度乍一看是 $\mathcal{O}(n^2)$，但注意变量 $i$ 减少的次数是 $\mathcal{O}(n)$ 的，其它情况一直在增加，由于 $i$ 最大是 $n$，所以增加的次数是 $\mathcal{O}(n)$，所以二重循环总共循环 $\mathcal{O}(n)$ 次，时间复杂度是 $\mathcal{O}(n)$ 的。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面滑动窗口与双指针题单的「**六、分组循环**」。

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
