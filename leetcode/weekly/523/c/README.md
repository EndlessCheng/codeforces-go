维护每个座位是否被预订，

维护每个座位预订时，所消耗的秒数。释放座位时，减去记录的秒数。

对于 $\texttt{getMinTime}$，这其实是个脑筋急转弯：**由于总是可以先预定靠窗的，再预定中间的，所以返回当前已预定座位数即可**。

代码实现时，可以用一个二进制数记录一行中的哪些座位被预订。原理见 [从集合论到位运算，常见位运算技巧分类总结](https://leetcode.cn/discuss/post/3571304/cong-ji-he-lun-dao-wei-yun-suan-chang-ji-enve/)。

下午两点 [B站@灵茶山艾府](https://space.bilibili.com/206214) 直播讲题，欢迎关注~

```py [sol-Python3]
class BusBooking:
    def __init__(self, n: int):
        self.states = [0] * (n + 1)
        self.times = [[0] * 4 for _ in range(n + 1)]
        self.total_time = 0
        self.booked_num = 0

    def toggle(self, seat: str) -> None:
        row = int(seat[1:])
        col = ord(seat[0]) - ord('A')
        mask = self.states[row]

        if mask >> col & 1 == 0:
            # 预订
            if col == 0 and mask >> 1 & 1 or col == 3 and mask >> 2 & 1:
                time = 3  # 靠窗，且相邻座位已被预订
            else:
                time = 1
            self.total_time += time
            self.times[row][col] = time  # 记录预订该座位所需的秒数，将来释放时用
            self.booked_num += 1
        else:
            # 释放
            self.total_time -= self.times[row][col]
            self.times[row][col] = 0
            self.booked_num -= 1

        self.states[row] ^= 1 << col

    def getTotalTime(self) -> int:
        return self.total_time

    def getMinTime(self) -> int:
        return self.booked_num
```

```java [sol-Java]
class BusBooking {
    private final int[] states;
    private final int[][] times;
    private int totalTime;
    private int bookedNum;

    public BusBooking(int n) {
        states = new int[n + 1];
        times = new int[n + 1][4];
    }

    public void toggle(String seat) {
        int row = Integer.parseInt(seat.substring(1));
        int col = seat.charAt(0) - 'A';
        int mask = states[row];

        if ((mask >> col & 1) == 0) {
            // 预订
            int time = col == 0 && (mask >> 1 & 1) > 0 || col == 3 && (mask >> 2 & 1) > 0 ? 3 : 1;
            totalTime += time;
            times[row][col] = time; // 记录预订该座位所需的秒数，将来释放时用
            bookedNum++;
        } else {
            // 释放
            totalTime -= times[row][col];
            times[row][col] = 0;
            bookedNum--;
        }

        states[row] ^= 1 << col;
    }

    public int getTotalTime() {
        return totalTime;
    }

    public int getMinTime() {
        return bookedNum;
    }
}
```

```cpp [sol-C++]
class BusBooking {
    vector<int> states;
    vector<array<int, 4>> times;
    int total_time = 0;
    int booked_num = 0;

public:
    BusBooking(int n) : states(n + 1), times(n + 1) {}

    void toggle(string seat) {
        int row = stoi(seat.substr(1));
        char col = seat[0] - 'A';
        int mask = states[row];

        if ((mask >> col & 1) == 0) {
            // 预订
            int time = col == 0 && (mask >> 1 & 1) || col == 3 && (mask >> 2 & 1) ? 3 : 1;
            total_time += time;
            times[row][col] = time; // 记录预订该座位所需的秒数，将来释放时用
            booked_num++;
        } else {
            // 释放
            total_time -= times[row][col];
            times[row][col] = 0;
            booked_num--;
        }

        states[row] ^= 1 << col;
    }

    int getTotalTime() {
        return total_time;
    }

    int getMinTime() {
        return booked_num;
    }
};
```

```go [sol-Go]
type BusBooking struct {
	states    []int
	times     [][4]int
	totalTime int
	bookedNum int
}

func Constructor(n int) BusBooking {
	return BusBooking{
		states: make([]int, n+1),
		times:  make([][4]int, n+1),
	}
}

func (b *BusBooking) Toggle(seat string) {
	row, _ := strconv.Atoi(seat[1:])
	mask := b.states[row]
	col := seat[0] - 'A'

	if mask>>col&1 == 0 {
		// 预订
		time := 1
		if col == 0 && mask>>1&1 > 0 || col == 3 && mask>>2&1 > 0 {
			time = 3 // 靠窗，且相邻座位已被预订
		}
		b.totalTime += time
		b.times[row][col] = time // 记录预订该座位所需的秒数，将来释放时用
		b.bookedNum++
	} else {
		// 释放
		b.totalTime -= b.times[row][col]
		b.times[row][col] = 0
		b.bookedNum--
	}

	b.states[row] ^= 1 << col
}

func (b *BusBooking) GetTotalTime() int {
	return b.totalTime
}

func (b *BusBooking) GetMinTime() int {
	return b.bookedNum
}
```

#### 复杂度分析

- 时间复杂度：
    - 初始化：$\mathcal{O}(n)$。
    - $\texttt{toggle}$：$\mathcal{O}(\textit{seat})$。把字符串转成整数需要 $\mathcal{O}(\textit{seat}) = \mathcal{O}(\log n)$ 的时间。
    - $\texttt{getTotalTime}$：$\mathcal{O}(1)$。
    - $\texttt{getMinTime}$：$\mathcal{O}(1)$。
- 空间复杂度：$\mathcal{O}(n)$。

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
