设子数组的和为 $S$。删除元素 $x$ 后，剩余元素的和 $S-x$ 必须是 $k$ 的倍数。

根据 [同余](https://leetcode.cn/discuss/post/3584387/fen-xiang-gun-mo-yun-suan-de-shi-jie-dan-7xgu/) 的定义，有

$$
x\equiv S \pmod k
$$

或者说

$$
x\bmod k = S\bmod k
$$

这意味着，**子数组的每个数模 $k$ 后都相等**。

设子数组的长度为 $\textit{cnt}$。去掉一个数后，子数组的和模 $k$ 必须为 $0$，即

$$
((x\bmod k)\cdot(\textit{cnt} - 1)) \bmod k = 0
$$

我们可以在遍历 $\textit{nums}$ 的过程中，维护连续同余段的长度 $\textit{cnt}$。如果上式成立，用 $\textit{cnt}$ 更新答案的最大值。

[本题视频讲解](https://www.bilibili.com/video/BV1rHpj6kE8u/?t=3m15s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def resilientSubarray(self, nums: list[int], k: int) -> int:
        ans = cnt = 0
        for i, x in enumerate(nums):
            r = x % k
            if i > 0 and r == nums[i - 1] % k:
                cnt += 1
            else:  # 新的模 k 同余段
                cnt = 1

            # 去掉一个数后，剩余元素之和是 k 的倍数
            if r * (cnt - 1) % k == 0:
                ans = max(ans, cnt)
        return ans
```

```java [sol-Java]
class Solution {
    public int resilientSubarray(int[] nums, int k) {
        int ans = 0;
        int cnt = 0;
        for (int i = 0; i < nums.length; i++) {
            int r = nums[i] % k;
            if (i > 0 && r == nums[i - 1] % k) {
                cnt++;
            } else { // 新的模 k 同余段
                cnt = 1;
            }

            // 去掉一个数后，剩余元素之和是 k 的倍数
            if (r * (cnt - 1) % k == 0) {
                ans = Math.max(ans, cnt);
            }
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int resilientSubarray(vector<int>& nums, int k) {
        int ans = 0, cnt = 0;
        for (int i = 0; i < nums.size(); i++) {
            int r = nums[i] % k;
            if (i > 0 && r == nums[i - 1] % k) {
                cnt++;
            } else { // 新的模 k 同余段
                cnt = 1;
            }

            // 去掉一个数后，剩余元素之和是 k 的倍数
            if (r * (cnt - 1) % k == 0) {
                ans = max(ans, cnt);
            }
        }
        return ans;
    }
};
```

```go [sol-Go]
func resilientSubarray(nums []int, k int) (ans int) {
	cnt := 0
	for i, x := range nums {
		r := x % k
		if i > 0 && r == nums[i-1]%k {
			cnt++
		} else { // 新的模 k 同余段
			cnt = 1
		}

		// 去掉一个数后，剩余元素之和是 k 的倍数
		if r*(cnt-1)%k == 0 {
			ans = max(ans, cnt)
		}
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

## 专题训练

见下面数学题单的「**§1.9 同余**」。

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
