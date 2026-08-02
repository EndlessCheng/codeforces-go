从时刻 $0$ 开始计时，设当前累计时间为 $t$。如何快速计算 $t$ 时间内可以完成多少个任务？

例如 $\textit{tasks} = [2,3,4]$，其前缀和数组为 $s = [2,5,9]$。

由于 $\textit{tasks}$ 中的数都是非负数，所以 $s$ 是递增的。我们可以在有序数组 $s$ 中**二分查找**最后一个 $\le t$ 的数的下标 $k$，那么已完成的任务下标为 $[0,k]$，未完成的任务下标为 $[k+1,n-1]$，这有 $n-k-1$ 个。关于二分查找的原理，请看 [二分查找 红蓝染色法【基础算法精讲 04】](https://www.bilibili.com/video/BV1AP41137w7/)。

代码实现时，可以改成在 $s$ 中二分查找第一个 $> t$ 的数的下标 $k' = k+1$，那么未完成的任务数为 $n-k'$ 个。

如果 $k' = n$，则重置累计时间 $t = 0$。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def countTasks(self, tasks: list[int], shifts: list[int]) -> list[int]:
        n = len(tasks)
        # 原地计算 tasks 的前缀和
        for i in range(1, n):
            tasks[i] += tasks[i - 1]

        t = 0
        for i, shift in enumerate(shifts):
            t += shift
            if t >= tasks[-1]:  # 完成所有任务
                t = 0
                shifts[i] = 0
            else:
                shifts[i] = n - bisect_right(tasks, t)
        return shifts
```

```java [sol-Java]
class Solution {
    public int[] countTasks(int[] tasks, int[] shifts) {
        int n = tasks.length;
        long[] s = new long[n];
        s[0] = tasks[0];
        for (int i = 1; i < n; i++) {
            s[i] = s[i - 1] + tasks[i];
        }

        long t = 0;
        for (int i = 0; i < shifts.length; i++) {
            t += shifts[i];
            if (t >= s[n - 1]) { // 完成所有任务
                t = 0;
                shifts[i] = 0;
            } else {
                // s 无重复元素，可以用库函数二分
                int j = Arrays.binarySearch(s, t + 1);
                if (j < 0) j = ~j; // 见 Arrays.binarySearch 源码
                shifts[i] = n - j;
            }
        }
        return shifts;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<int> countTasks(vector<int>& tasks, vector<int>& shifts) {
        int n = tasks.size();
        vector<long long> s(n);
        s[0] = tasks[0];
        for (int i = 1; i < n; i++) {
            s[i] = s[i - 1] + tasks[i];
        }

        long long t = 0;
        for (int& shift : shifts) {
            t += shift;
            if (t >= s[n - 1]) { // 完成所有任务
                t = 0;
                shift = 0;
            } else {
                shift = s.end() - ranges::upper_bound(s, t);
            }
        }
        return shifts;
    }
};
```

```go [sol-Go]
func countTasks(tasks, shifts []int) []int {
	n := len(tasks)
	// 原地计算 tasks 的前缀和
	for i := 1; i < n; i++ {
		tasks[i] += tasks[i-1]
	}

	t := 0
	for i, shift := range shifts {
		t += shift
		if t >= tasks[n-1] { // 完成所有任务
			t = 0
			shifts[i] = 0
		} else {
			shifts[i] = n - sort.SearchInts(tasks, t+1)
		}
	}
	return shifts
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n + m\log n)$，其中 $n$ 是 $\textit{tasks}$ 的长度，$m$ 是 $\textit{shifts}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$ 或 $\mathcal{O}(1)$。

## 附：二分查找常用转化表

| **需求**  | **写法**  |  **如果不存在** | 
|---|---|---|
| $\ge x$ 的第一个元素的下标  | $\texttt{lowerBound}(\textit{nums},x)$  | 结果为 $n$ |
| $> x$ 的第一个元素的下标 | $\texttt{lowerBound}(\textit{nums},x+1)$  | 结果为 $n$ |
| $< x$ 的最后一个元素的下标  | $\texttt{lowerBound}(\textit{nums},x)-1$  | 结果为 $-1$ |
| $\le x$ 的最后一个元素的下标  | $\texttt{lowerBound}(\textit{nums},x+1)-1$  | 结果为 $-1$ |

| **需求**  | **写法**  |
|---|---|
| $< x$ 的元素个数  | $\texttt{lowerBound}(\textit{nums},x)$  | 
| $\le x$ 的元素个数 | $\texttt{lowerBound}(\textit{nums},x+1)$  | 
| $\ge x$ 的元素个数  | $n - \texttt{lowerBound}(\textit{nums},x)$  | 
| $> x$ 的元素个数  | $n - \texttt{lowerBound}(\textit{nums},x+1)$  | 

## 专题训练

见下面二分题单的「**一、二分查找**」。

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
