对于数字 $x$ 和 $y$，设 $d = |x-y|$，那么：

- 指针不经过 $0\text{-}9$ 的旋转次数为 $d$。
- 指针经过 $0\text{-}9$ 的旋转次数为 $10-d$。

二者取最小值，得到指针从数字 $x$ 旋转到数字 $y$ 的最少旋转次数为

$$
\min(d, 10-d)
$$

[本题视频讲解](https://www.bilibili.com/video/BV1FiHj6uESs/)，欢迎点赞关注~

```py [sol-Python3]
class Solution:
    def minRotations(self, s: str) -> int:
        ans = 0
        pre = ord('0')
        for x in map(ord, s):
            d = abs(x - pre)
            ans += min(d, 10 - d)
            pre = x
        return ans
```

```java [sol-Java]
class Solution {
    public int minRotations(String s) {
        int ans = 0;
        char pre = '0';
        for (char ch : s.toCharArray()) {
            int d = Math.abs(ch - pre);
            ans += Math.min(d, 10 - d);
            pre = ch;
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    int minRotations(string s) {
        int ans = 0;
        char pre = '0';
        for (char ch : s) {
            int d = abs(ch - pre);
            ans += min(d, 10 - d);
            pre = ch;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minRotations(s string) (ans int) {
	pre := '0'
	for _, ch := range s {
		d := abs(int(ch) - int(pre))
		ans += min(d, 10-d)
		pre = ch
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 是 $s$ 的长度。
- 空间复杂度：$\mathcal{O}(1)$。

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
