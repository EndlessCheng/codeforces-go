把小于 $a$ 的数视作 $0$，在 $[a,b]$ 中的数视作 $1$，大于 $b$ 的数视作 $2$。

问题等价于：

- 把这个只包含 $0,1,2$ 的数组排成升序，只能交换相邻元素，求最小交换次数。

**定理**：在只能交换相邻元素的前提下，把数组 $a$ 排成升序，最小交换次数等于数组 $a$ 的逆序对的个数。其中「逆序对」指 $i<j$ 且 $a[i]>a[j]$ 的下标对 $(i,j)$。

**证明**：设数组 $a$ 的逆序对的个数为 $I$。

- 由于一次相邻交换只改变一对相邻元素的逆序对，所以一次相邻交换至多把 $I$ 减少 $1$，所以至少要交换 $I$ 次。
- 这个下界是可以达到的。如果 $a$ 不是升序的，那么存在一对相邻元素 $(x,y)$ 满足 $x>y$。我们随便选一对逆序相邻元素，交换，可以让 $I$ 恰好减少 $1$。重复该过程，交换 $I$ 次，即可把 $a$ 排成升序。

如何计算逆序对的个数？一般需要值域树状数组，但本题只有三种元素，可以用两个变量 $\textit{cnt}_1$ 和 $\textit{cnt}_2$ 分别统计遍历过的 $1$ 和 $2$ 的个数。

- 如果当前元素是 $0$，那么新增 $\textit{cnt}_1+\textit{cnt}_2$ 个逆序对。
- 如果当前元素是 $1$，那么新增 $\textit{cnt}_2$ 个逆序对。
- 如果当前元素是 $2$，那么新增 $0$ 个逆序对。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def minAdjacentSwaps(self, nums: list[int], a: int, b: int) -> int:
        ans = cnt1 = cnt2 = 0
        for x in nums:
            if x < a:  # x 视作 0
                ans += cnt1 + cnt2
            elif x <= b:  # x 视作 1
                ans += cnt2
                cnt1 += 1
            else:  # x 视作 2
                cnt2 += 1
        return ans % 1_000_000_007
```

```java [sol-Java]
class Solution {
    public int minAdjacentSwaps(int[] nums, int a, int b) {
        long ans = 0;
        int cnt1 = 0;
        int cnt2 = 0;
        for (int x : nums) {
            if (x < a) { // x 视作 0
                ans += cnt1 + cnt2;
            } else if (x <= b) { // x 视作 1
                ans += cnt2;
                cnt1++;
            } else { // x 视作 2
                cnt2++;
            }
        }
        return (int) (ans % 1_000_000_007);
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minAdjacentSwaps(vector<int>& nums, int a, int b) {
        long long ans = 0;
        int cnt1 = 0, cnt2 = 0;
        for (int x : nums) {
            if (x < a) { // x 视作 0
                ans += cnt1 + cnt2;
            } else if (x <= b) { // x 视作 1
                ans += cnt2;
                cnt1++;
            } else { // x 视作 2
                cnt2++;
            }
        }
        return ans % 1'000'000'007;
    }
};
```

```go [sol-Go]
func minAdjacentSwaps(nums []int, a, b int) (ans int) {
	const mod = 1_000_000_007
	cnt1, cnt2 := 0, 0
	for _, x := range nums {
		if x < a { // x 视作 0
			ans += cnt1 + cnt2
		} else if x <= b { // x 视作 1
			ans += cnt2
			cnt1++
		} else { // x 视作 2
			cnt2++
		}
	}
	return ans % mod
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数据结构题单的「**§8.2 逆序对**」。

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
