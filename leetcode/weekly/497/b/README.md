设三条边为 $a,b,c$，满足 $1\le a\le b\le c$。

由于 $a+c>b$ 和 $b+c>a$ 一定成立，所以只需判断 $a+b>c$ 是否成立。若不成立，返回空数组。

设 $a,b,c$ 的对角分别为 $A,B,C$。根据余弦定理：

$$
\begin{aligned}
\cos A &= \dfrac{b^2+c^2-a^2}{2bc}     \\
\cos B &= \dfrac{a^2+c^2-b^2}{2ac}     \\
\cos C &= \dfrac{a^2+b^2-c^2}{2ab}     \\
\end{aligned}
$$

把弧度转成角度：

$$
\begin{aligned}
A &= \arccos\left(\dfrac{b^2+c^2-a^2}{2bc}\right)\cdot\dfrac{180}{\pi}   \\
B &= \arccos\left(\dfrac{a^2+c^2-b^2}{2ac}\right)\cdot\dfrac{180}{\pi}   \\
C &= 180 - A - B \\
\end{aligned}
$$

由于「小边对小角」，所以 $A\le B\le C$。

```py [sol-Python3]
class Solution:
    def internalAngles(self, sides: List[int]) -> List[float]:
        sides.sort()
        a, b, c = sides
        if a + b <= c:
            return []

        A = degrees(acos((b * b + c * c - a * a) / (b * c * 2)))
        B = degrees(acos((a * a + c * c - b * b) / (a * c * 2)))
        return [A, B, 180 - A - B]  # 小边对小角
```

```java [sol-Java]
class Solution {
    public double[] internalAngles(int[] sides) {
        Arrays.sort(sides);
        int a = sides[0], b = sides[1], c = sides[2];
        if (a + b <= c) {
            return new double[0];
        }

        final double rad = 180 / Math.PI;
        double A = Math.acos((double) (b * b + c * c - a * a) / (double) (b * c * 2)) * rad;
        double B = Math.acos((double) (a * a + c * c - b * b) / (double) (a * c * 2)) * rad;
        return new double[]{A, B, 180 - A - B}; // 小边对小角
    }
}
```

```cpp [sol-C++]
class Solution {
public:
    vector<double> internalAngles(vector<int>& sides) {
        ranges::sort(sides);
        int a = sides[0], b = sides[1], c = sides[2];
        if (a + b <= c) {
            return {};
        }

        const double rad = 180 / acos(-1);
        double A = acos(1.0 * (b * b + c * c - a * a) / (b * c * 2)) * rad;
        double B = acos(1.0 * (a * a + c * c - b * b) / (a * c * 2)) * rad;
        return {A, B, 180 - A - B}; // 小边对小角
    }
};
```

```go [sol-Go]
func internalAngles(sides []int) []float64 {
	slices.Sort(sides)
	a, b, c := sides[0], sides[1], sides[2]
	if a+b <= c {
		return nil
	}

	const rad = 180 / math.Pi
	A := math.Acos(float64(b*b+c*c-a*a)/float64(b*c*2)) * rad
	B := math.Acos(float64(a*a+c*c-b*b)/float64(a*c*2)) * rad
	return []float64{A, B, 180 - A - B} // 小边对小角
}
```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(1)$。

## 分类题单

[如何科学刷题？](https://leetcode.cn/circle/discuss/RvFUtj/)

1. [滑动窗口与双指针（定长/不定长/单序列/双序列/三指针/分组循环）](https://leetcode.cn/circle/discuss/0viNMK/)
2. [二分算法（二分答案/最小化最大值/最大化最小值/第K小）](https://leetcode.cn/circle/discuss/SqopEo/)
3. [单调栈（基础/矩形面积/贡献法/最小字典序）](https://leetcode.cn/circle/discuss/9oZFK9/)
4. [网格图（DFS/BFS/综合应用）](https://leetcode.cn/circle/discuss/YiXPXW/)
5. [位运算（基础/性质/拆位/试填/恒等式/思维）](https://leetcode.cn/circle/discuss/dHn9Vk/)
6. [图论算法（DFS/BFS/拓扑排序/基环树/最短路/最小生成树/网络流）](https://leetcode.cn/circle/discuss/01LUak/)
7. [动态规划（入门/背包/划分/状态机/区间/状压/数位/数据结构优化/树形/博弈/概率期望）](https://leetcode.cn/circle/discuss/tXLS3i/)
8. [常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树）](https://leetcode.cn/circle/discuss/mOr1u6/)
9. [数学算法（数论/组合/概率期望/博弈/计算几何/随机算法）](https://leetcode.cn/circle/discuss/IYT3ss/)
10. [贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造）](https://leetcode.cn/circle/discuss/g6KTKL/)
11. [链表、树与回溯（前后指针/快慢指针/DFS/BFS/直径/LCA）](https://leetcode.cn/circle/discuss/K0n2gO/)
12. [字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机）](https://leetcode.cn/circle/discuss/SJFwQI/)
