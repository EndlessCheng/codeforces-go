大的思路是，枚举 $j$，同时维护左边能与 $\textit{nums}[j]$ 构成影子对的 $\textit{nums}[i]$。

根据题意，一旦 $\textit{nums}[i]$ 右边出现比 $\textit{nums}[i]$ 小的数（作为 $\textit{nums}[k]$），那么 $\textit{nums}[i]$ 无法构成影子对。比如示例 1 的 $[3,1,\ldots]$，由于 $1$ 的出现，$3$ 永远无法构成影子对。

用一个栈保存遍历过的数。如果当前元素 $x$ 比栈顶小，那么栈顶永远无法构成影子对，弹出栈顶。反复操作，直到栈为空，或者栈顶元素 $\le x$。

如此操作后，从栈底到栈顶是递增的（允许相等）。由于我们去掉了无法构成影子对的数，所以在 $x$ 左侧的能与 $x$ 构成影子对的数都在栈中。

但栈顶元素可能等于 $x$，不能计入。此外，栈中有相同元素，我们不能暴力统计栈中 $x$ 的个数，那样时间复杂度会退化至 $\mathcal{O}(n^2)$。我们可以改为在栈中保存 $(元素,元素的出现次数)$，同时维护栈的大小 $\textit{size}$。

- 如果栈顶元素小于 $x$，那么有 $\textit{size}$ 个元素可以与 $x$ 组成影子对。
- 如果栈顶元素等于 $x$，那么有 $(\textit{size} - 栈顶元素的出现次数)$ 个元素可以与 $x$ 组成影子对。

[本题视频讲解](https://www.bilibili.com/video/BV1k7Yv6WE3i/?t=8m58s)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def shadowPairs(self, nums: list[int]) -> int:
        st = [[0, 0]]  # 栈底哨兵
        size = 0  # 栈的大小（出现次数之和）
        ans = 0
        for x in nums:
            while st[-1][0] > x:
                size -= st.pop()[1]  # 栈顶永远无法构成影子对
            ans += size
            if st[-1][0] == x:
                # 恰好等于 x 的 nums[i] 不能构成影子对，要减掉
                ans -= st[-1][1]
                st[-1][1] += 1
            else:
                st.append([x, 1])
            size += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long shadowPairs(int[] nums) {
        ArrayList<int[]> st = new ArrayList<>();
        st.add(new int[]{0, 0}); // 栈底哨兵
        int size = 0; // 栈的大小（出现次数之和）
        long ans = 0;
        for (int x : nums) {
            while (st.getLast()[0] > x) {
                size -= st.removeLast()[1]; // 栈顶永远无法构成影子对
            }

            ans += size;
            if (st.getLast()[0] == x) {
                // 恰好等于 x 的 nums[i] 不能构成影子对，要减掉
                ans -= st.getLast()[1];
                st.getLast()[1]++;
            } else {
                st.add(new int[]{x, 1});
            }
            size++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long shadowPairs(vector<int>& nums) {
        stack<pair<int, int>> st;
        st.emplace(0, 0); // 栈底哨兵
        int size = 0; // 栈的大小（出现次数之和）
        long long ans = 0;
        for (int x : nums) {
            while (st.top().first > x) {
                size -= st.top().second;
                st.pop(); // 栈顶永远无法构成影子对
            }

            ans += size;
            if (st.top().first == x) {
                // 恰好等于 x 的 nums[i] 不能构成影子对，要减掉
                ans -= st.top().second;
                st.top().second++;
            } else {
                st.emplace(x, 1);
            }
            size++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func shadowPairs(nums []int) (ans int64) {
	type pair struct{ x, cnt int }
	st := []pair{{}} // 栈底哨兵
	size := 0 // 栈的大小（cnt 之和）
	for _, x := range nums {
		for st[len(st)-1].x > x {
			size -= st[len(st)-1].cnt
			st = st[:len(st)-1] // 栈顶永远无法构成影子对
		}

		ans += int64(size)
		if st[len(st)-1].x == x {
			// 恰好等于 x 的 nums[i] 不能构成影子对，要减掉
			ans -= int64(st[len(st)-1].cnt)
			st[len(st)-1].cnt++
		} else {
			st = append(st, pair{x, 1})
		}
		size++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $\textit{nums}$ 的长度。虽然我们写了个二重循环，但站在每个元素的视角看，这个元素在二重循环中最多入栈出栈各一次，因此循环次数**之和**是 $\mathcal{O}(n)$，所以时间复杂度是 $\mathcal{O}(n)$。
- 空间复杂度：$\mathcal{O}(n)$。

## 专题训练

见下面单调栈题单的「**一、单调栈**」。

[1209. 删除字符串中的所有相邻重复项 II](https://leetcode.cn/problems/remove-all-adjacent-duplicates-in-string-ii/) 用到了类似的技巧。

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
