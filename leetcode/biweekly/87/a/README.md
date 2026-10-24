把字符串日期转换成在这一年的第几天，问题相当于：

- 给你两个闭区间，求这两个闭区间的交集中的整数个数。

例如 $[10,30]$ 与 $[20,40]$ 的交集为 $[20,30]$，区间内的整数个数为 $30-20+1=11$。

如何计算区间交集的左右端点？

- 左端点 $\textit{left}$ 等于两个区间左端点的最大值。
- 右端点 $\textit{right}$ 等于两个区间右端点的最小值。

交集中的整数个数为 $\textit{right}-\textit{left}+1$。

如果交集为空，即 $\textit{right}-\textit{left}+1\le 0$，那么答案为 $0$。

```py [sol-Python3]
# 每月天数的前缀和（注：最后一个数可以省略）
DAYS_PRE_SUM = list(accumulate((31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31), initial=0))

def calc_days(date: str) -> int:
    return DAYS_PRE_SUM[int(date[:2]) - 1] + int(date[3:])

class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        left = calc_days(max(arriveAlice, arriveBob))  # 直接比较字符串再计算
        right = calc_days(min(leaveAlice, leaveBob))
        return max(right - left + 1, 0)  # 答案不能为负数
```

```py [sol-Python3 库函数]
def calc_dt(date: str) -> datetime.datetime:
    return datetime.datetime.strptime(date, '%m-%d')  # 默认是 1900 年（平年）

class Solution:
    def countDaysTogether(self, arriveAlice: str, leaveAlice: str, arriveBob: str, leaveBob: str) -> int:
        left = calc_dt(max(arriveAlice, arriveBob))  # 直接比较字符串再计算
        right = calc_dt(min(leaveAlice, leaveBob))
        return max((right - left).days + 1, 0)  # 答案不能为负数
```

```java [sol-Java]
class Solution {
    // 每月天数的前缀和（注：最后一个数可以省略）
    private static final int[] daysPreSum = new int[]{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365};

    private int calcDays(String S) {
        char[] s = S.toCharArray();
        int month = (s[0] - '0') * 10 + s[1] - '0';
        int day = (s[3] - '0') * 10 + s[4] - '0';
        return daysPreSum[month - 1] + day;
    }

    public int countDaysTogether(String arriveAlice, String leaveAlice, String arriveBob, String leaveBob) {
        int left = calcDays(arriveAlice.compareTo(arriveBob) > 0 ? arriveAlice : arriveBob); // 直接比较字符串再计算
        int right = calcDays(leaveAlice.compareTo(leaveBob) < 0 ? leaveAlice : leaveBob);
        return Math.max(right - left + 1, 0); // 答案不能为负数
    }
}
```

```cpp [sol-C++]
class Solution {
    // 每月天数的前缀和（注：最后一个数可以省略）
    static constexpr int days_pre_sum[] = {0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365};

    int calc_days(string s) {
        int month = (s[0] - '0') * 10 + s[1] - '0';
        int day = (s[3] - '0') * 10 + s[4] - '0';
        return days_pre_sum[month - 1] + day;
    }

public:
    int countDaysTogether(string arriveAlice, string leaveAlice, string arriveBob, string leaveBob) {
        int left = calc_days(max(arriveAlice, arriveBob)); // 直接比较字符串再计算
        int right = calc_days(min(leaveAlice, leaveBob));
        return max(right - left + 1, 0); // 答案不能为负数
    }
};
```

```go [sol-Go]
// 每月天数的前缀和（注：最后一个数可以省略）
var daysPreSum = []int{0, 31, 59, 90, 120, 151, 181, 212, 243, 273, 304, 334, 365}

func calcDays(s string) int {
	month := (s[0]-'0')*10 + (s[1] - '0')
	day := (s[3]-'0')*10 + (s[4] - '0')
	return daysPreSum[month-1] + int(day)
}

func countDaysTogether(arriveAlice, leaveAlice, arriveBob, leaveBob string) int {
	left := calcDays(max(arriveAlice, arriveBob)) // 直接比较字符串再计算
	right := calcDays(min(leaveAlice, leaveBob))
	return max(right-left+1, 0) // 答案不能为负数
}
```

### 复杂度分析

- 时间复杂度：$\mathcal{O}(1)$。
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
