下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class Solution:
    def parse(self, t: str) -> int:
        return int(t[:2]) * 3600 + int(t[3:5]) * 60 + int(t[6:])

    def secondsBetweenTimes(self, startTime: str, endTime: str) -> int:
        return self.parse(endTime) - self.parse(startTime)
```

```java [sol-Java]
class Solution {
    public int secondsBetweenTimes(String startTime, String endTime) {
        return parse(endTime) - parse(startTime);
    }

    private int parse(String time) {
        char[] t = time.toCharArray();
        int hour = (t[0] - '0') * 10 + (t[1] - '0');
        int minute = (t[3] - '0') * 10 + (t[4] - '0');
        int second = (t[6] - '0') * 10 + (t[7] - '0');
        return hour * 3600 + minute * 60 + second;
    }
}
```

```cpp [sol-C++]
class Solution {
    int parse(const string& t) {
        int hour = (t[0] - '0') * 10 + (t[1] - '0');
        int minute = (t[3] - '0') * 10 + (t[4] - '0');
        int second = (t[6] - '0') * 10 + (t[7] - '0');
        return hour * 3600 + minute * 60 + second;
    }

public:
    int secondsBetweenTimes(string startTime, string endTime) {
        return parse(endTime) - parse(startTime);
    }
};
```

```go [sol-Go]
func parse(t string) int {
	hour := int(t[0]-'0')*10 + int(t[1]-'0')
	minute := int(t[3]-'0')*10 + int(t[4]-'0')
	second := int(t[6]-'0')*10 + int(t[7]-'0')
	return hour*3600 + minute*60 + second
}

func secondsBetweenTimes(startTime, endTime string) int {
	return parse(endTime) - parse(startTime)
}
```

#### 复杂度分析

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
