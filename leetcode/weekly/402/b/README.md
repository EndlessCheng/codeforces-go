本题和 [1010. 总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/) 几乎一样，把那道题的 $60$ 改成 $24$ 即可（注意本题返回值是 64 位整数）。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1T1421k7Hi/) 和 [文字题解](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/solution/liang-shu-zhi-he-de-ben-zhi-shi-shi-yao-bd0r1/)，欢迎点赞关注！

```py [sol-Python3]
class Solution:
    def countCompleteDayPairs(self, hours: List[int]) -> int:
        ans = 0
        cnt = [0] * 24
        for t in hours:
            # 先查询 cnt，再更新 cnt，因为题目要求 i<j
            # 如果先更新，再查询，就把 i=j 的情况也考虑进去了
            ans += cnt[(24 - t % 24) % 24]
            cnt[t % 24] += 1
        return ans
```

```java [sol-Java]
class Solution {
    public long countCompleteDayPairs(int[] hours) {
        long ans = 0;
        int[] cnt = new int[24];
        for (int t : hours) {
            // 先查询 cnt，再更新 cnt，因为题目要求 i<j
            // 如果先更新，再查询，就把 i=j 的情况也考虑进去了
            ans += cnt[(24 - t % 24) % 24];
            cnt[t % 24]++;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long countCompleteDayPairs(vector<int> &hours) {
        long long ans = 0;
        int cnt[24]{};
        for (int t : hours) {
            // 先查询 cnt，再更新 cnt，因为题目要求 i<j
            // 如果先更新，再查询，就把 i=j 的情况也考虑进去了
            ans += cnt[(24 - t % 24) % 24];
            cnt[t % 24]++;
        }
        return ans;
    }
};
```

```go [sol-Go]
func countCompleteDayPairs(hours []int) (ans int64) {
	cnt := [24]int{}
	for _, t := range hours {
		// 先查询 cnt，再更新 cnt，因为题目要求 i<j
		// 如果先更新，再查询，就把 i=j 的情况也考虑进去了
		ans += int64(cnt[(24-t%24)%24])
		cnt[t%24]++
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n+M)$，其中 $n$ 为 $\textit{hours}$ 的长度，$M=24$。
- 空间复杂度：$\mathcal{O}(M)$。

## 一句话总结

对于有两个变量的题目，通常可以枚举其中一个变量，把它视作常量，从而转化成只有一个变量的问题。

附：代码中用到的取模技巧见 [模运算的世界：当加减乘除遇上取模](https://leetcode.cn/circle/discuss/mDfnkW/)。

## 相似题目

- [1. 两数之和](https://leetcode.cn/problems/two-sum/)
- [219. 存在重复元素 II](https://leetcode.cn/problems/contains-duplicate-ii/)
- [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)
- [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/) 1161 经典题
- [2815. 数组中的最大数对和](https://leetcode.cn/problems/max-pair-sum-in-an-array/) 1295
- [2748. 美丽下标对的数目](https://leetcode.cn/problems/number-of-beautiful-pairs/) 1301
- [2342. 数位和相等数对的最大和](https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits/) 1309
- [1679. K 和数对的最大数目](https://leetcode.cn/problems/max-number-of-k-sum-pairs/) 1346
- [1010. 总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/) 1377
- [2971. 找到最大周长的多边形](https://leetcode.cn/problems/find-polygon-with-the-largest-perimeter/) 1521
- [2874. 有序三元组中的最大值 II](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii/) 1583
- [1014. 最佳观光组合](https://leetcode.cn/problems/best-sightseeing-pair/) 1730
- [1814. 统计一个数组中好对子的数目](https://leetcode.cn/problems/count-nice-pairs-in-an-array/) 1738
- [454. 四数相加 II](https://leetcode.cn/problems/4sum-ii/)
- [1214. 查找两棵二叉搜索树之和](https://leetcode.cn/problems/two-sum-bsts/)（会员题）
- [2613. 美数对](https://leetcode.cn/problems/beautiful-pairs/)（会员题）
- [2964. 可被整除的三元组数量](https://leetcode.cn/problems/number-of-divisible-triplet-sums/)（会员题）

## 分类题单

以下题单没有特定的顺序，可以按照个人喜好刷题。

1. [滑动窗口（定长/不定长/多指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/贪心/脑筋急转弯）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
