首先，一直攻击同一个敌人，相比来回攻击多个敌人（雨露均沾）更好，因为这样我们被敌人攻击的次数更少。

**从特殊到一般**，想一想，如果只有**两个**敌人 A 和 B，我们应该先攻击谁？

消灭 A 需要攻击的次数为

$$
k_A = \left\lceil\dfrac{\textit{health}_A}{\textit{power}}\right\rceil = \left\lfloor\dfrac{\textit{health}_A-1}{\textit{power}}\right\rfloor + 1
$$

> 讨论 $\textit{health}_A$ 被 $\textit{power}$ 整除，和不被 $\textit{power}$ 整除两种情况，可以证明上式的正确性。

同理可得消灭 B 需要的攻击次数，记作 $k_B$。

如果先消灭 A，再消灭 B，那么受到的伤害总和为

$$
k_A\cdot \textit{damage}_A + (k_A+k_B)\cdot \textit{damage}_B
$$

如果先消灭 B，再消灭 A，那么受到的伤害总和为

$$
k_B\cdot \textit{damage}_B + (k_A+k_B)\cdot \textit{damage}_A
$$

如果先消灭 A，再消灭 B 更好，那么有

$$
k_A\cdot \textit{damage}_A + (k_A+k_B)\cdot \textit{damage}_B < k_B\cdot \textit{damage}_B + (k_A+k_B)\cdot \textit{damage}_A
$$

化简得

$$
k_A\cdot \textit{damage}_B < k_B\cdot \textit{damage}_A
$$

> 也就是优先消灭 $k/\textit{damage}$ 更小的敌人。结合你玩过的某些游戏，应当优先消灭那个**又脆皮伤害又高**的敌人。

推广到更多的敌人，可以按照上式对 $\textit{damage}$ 和 $\textit{health}$ 排序，理由如下。

先假定按照输入的顺序消灭敌人。如果发现相邻两个敌人不满足上面的不等式，就交换这两个敌人的位置，这可以让受到的总伤害变小。

不断交换敌人，直到所有相邻敌人都满足上面的不等式。

本质上来说，这个不断交换相邻敌人的过程，和**冒泡排序**是一样的。那么按照不等式对数组排序即可。

排序后，按照顺序消灭敌人。用一个变量 $s$ 维护从一开始到击败当前敌人，所经过的秒数。把 $s\cdot \textit{damage}[i]$ 加入答案。

具体请看 [视频讲解](https://www.bilibili.com/video/BV1ajHYeoEG5/) 第四题，欢迎点赞关注~

## 答疑

**问**：本题可以用 $k/\textit{damage}$ 比大小吗？除法的结果需要用浮点数表示，精度误差是否会影响结果？

**答**：本题 $10^5$ 的值域范围是可以的，但如果范围是 $10^9$ 就可能会有问题，见 [2280. 表示一个折线图的最少线段数](https://leetcode.cn/problems/minimum-lines-to-represent-a-line-chart/)，可以用浮点数写试试。

```py [sol-Py]
class Solution:
    def minDamage(self, power: int, damage: List[int], health: List[int]) -> int:
        a = [((h - 1) // power + 1, d) for h, d in zip(health, damage)]
        a.sort(key=lambda p: p[0] / p[1])

        ans = s = 0
        for k, d in a:
            s += k
            ans += s * d
        return ans
```

```py [sol-Py 写法二]
class Solution:
    def minDamage(self, power: int, damage: List[int], health: List[int]) -> int:
        a = [((h - 1) // power + 1, d) for h, d in zip(health, damage)]
        a.sort(key=cmp_to_key(lambda p, q: p[0] * q[1] - q[0] * p[1]))

        ans = s = 0
        for k, d in a:
            s += k
            ans += s * d
        return ans
```

```py [sol-Py 写法三]
class Solution:
    def minDamage(self, power: int, damage: List[int], health: List[int]) -> int:
        n = len(damage)
        for i in range(n):
            health[i] = (health[i] - 1) // power + 1
        ans = s = 0
        for i in sorted(range(n), key=lambda i: health[i] / damage[i]):
            s += health[i]
            ans += s * damage[i]
        return ans
```

```java [sol-Java]
class Solution {
    public long minDamage(int power, int[] damage, int[] health) {
        int n = health.length;
        int[][] a = new int[n][2];
        for (int i = 0; i < n; i++) {
            a[i][0] = (health[i] - 1) / power + 1;
            a[i][1] = damage[i];
        }
        Arrays.sort(a, (p, q) -> p[0] * q[1] - q[0] * p[1]);

        long ans = 0;
        long s = 0;
        for (int[] p : a) {
            s += p[0];
            ans += s * p[1];
        }
        return ans;
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    long long minDamage(int power, vector<int>& damage, vector<int>& health) {
        int n = health.size();
        vector<pair<int, int>> a(n);
        for (int i = 0; i < n; i++) {
            a[i] = {(health[i] - 1) / power + 1, damage[i]};
        }
        ranges::sort(a, [](const auto& p, const auto& q) {
            return p.first * q.second < q.first * p.second;
        });

        long long ans = 0, s = 0;
        for (auto& [k, d] : a) {
            s += k;
            ans += s * d;
        }
        return ans;
    }
};
```

```go [sol-Go]
func minDamage(power int, damage, health []int) (ans int64) {
	type pair struct{ k, d int }
	a := make([]pair, len(health))
	for i, h := range health {
		a[i] = pair{(h-1)/power + 1, damage[i]}
	}
	slices.SortFunc(a, func(p, q pair) int { return p.k*q.d - q.k*p.d })

	s := 0
	for _, p := range a {
		s += p.k
		ans += int64(s) * int64(p.d)
	}
	return
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 是 $\textit{damage}$ 的长度。瓶颈在排序上。
- 时间复杂度：$\mathcal{O}(n)$。

更多相似题目，见下面贪心题单中的「**§1.7 交换论证法**」。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)

[我的题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
