package copypasta

import (
	"maps"
	"math"
	"math/bits"
	"math/rand"
	"reflect"
	"slices"
	"sort"
	"unsafe"
)

/*
如何科学刷题？ https://leetcode.cn/circle/discuss/RvFUtj/
1. 滑动窗口与双指针（定长/不定长/单序列/双序列/三指针） https://leetcode.cn/circle/discuss/0viNMK/
2. 二分算法（二分答案/最小化最大值/最大化最小值/第K小） https://leetcode.cn/circle/discuss/SqopEo/
3. 单调栈（基础/矩形面积/贡献法/最小字典序） https://leetcode.cn/circle/discuss/9oZFK9/
4. 网格图（DFS/BFS/综合应用） https://leetcode.cn/circle/discuss/YiXPXW/
5. 位运算（基础/性质/拆位/试填/恒等式/思维） https://leetcode.cn/circle/discuss/dHn9Vk/
6. 图论算法（DFS/BFS/拓扑排序/最短路/最小生成树/二分图/基环树/欧拉路径） https://leetcode.cn/circle/discuss/01LUak/
7. 动态规划（入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望） https://leetcode.cn/circle/discuss/tXLS3i/
8. 常用数据结构（前缀和/差分/栈/队列/堆/字典树/并查集/树状数组/线段树） https://leetcode.cn/circle/discuss/mOr1u6/
9. 数学算法（数论/组合/概率期望/博弈/计算几何/随机算法） https://leetcode.cn/circle/discuss/IYT3ss/
10. 贪心与思维（基本贪心策略/反悔/区间/字典序/数学/思维/脑筋急转弯/构造） https://leetcode.cn/circle/discuss/g6KTKL/
11. 链表、二叉树与一般树（前后指针/快慢指针/DFS/BFS/直径/LCA） https://leetcode.cn/circle/discuss/K0n2gO/
12. 字符串（KMP/Z函数/Manacher/字符串哈希/AC自动机/后缀数组/子序列自动机） https://leetcode.cn/circle/discuss/SJFwQI/

https://leetcode.cn/studyplan/primers-list/
https://leetcode.cn/studyplan/programming-skills/ 可选

字符串基础
https://codeforces.com/problemset/problem/1101/B
https://leetcode.cn/problems/apply-operations-to-make-string-empty/

暴力枚举
https://codeforces.com/problemset/problem/681/B 1300
- [2207. 字符串中最多数目的子序列](https://leetcode.cn/problems/maximize-number-of-subsequences-in-a-string/) 1550
另见 search.go

枚举右，维护左
- [1. 两数之和](https://leetcode.cn/problems/two-sum/)
   - https://codeforces.com/problemset/problem/702/B
- [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/) 1161 经典题
    - https://leetcode.cn/problems/sum-of-digit-differences-of-all-pairs/
    - 反向构造 https://codeforces.com/problemset/problem/1927/B 900
https://leetcode.com/discuss/interview-question/3685049/25-variations-of-Two-sum-question
https://codeforces.com/problemset/problem/1420/B 1200
https://codeforces.com/problemset/problem/318/B 1300 子串
https://codeforces.com/problemset/problem/1926/D 1300
https://codeforces.com/problemset/problem/1800/F 1900 异或

枚举右，维护左：需要维护两种值（pair）
https://codeforces.com/problemset/problem/1931/D 1300
https://leetcode.cn/problems/count-beautiful-substrings-ii/ 2445

枚举中间
https://codeforces.com/problemset/problem/1957/D 1900 前后缀分解 从高到低思考

任意下标 i 和 j
https://codeforces.com/problemset/problem/1895/C 1400

哈希表
- [2260. 必须拿起的最小连续卡牌数](https://leetcode.cn/problems/minimum-consecutive-cards-to-pick-up/) 1365
- [982. 按位与为零的三元组](https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/) 2085
- [面试题 16.21. 交换和](https://leetcode.cn/problems/sum-swap-lcci/)

前缀和
https://codeforces.com/problemset/problem/466/C

前缀和+哈希表（双变量思想）
- [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
   - 子数组中的和为 k 的子数组的个数之和 https://codeforces.com/problemset/problem/1996/E 1600
- [974. 和可被 K 整除的子数组](https://leetcode.cn/problems/subarray-sums-divisible-by-k/) 1676
   - 变形：乘积可以被 k 整除
   - a[i] = gcd(a[i], k) 之后窗口乘积是 k 的倍数就行，不会乘爆
https://atcoder.jp/contests/abc146/tasks/abc146_e 1762
https://atcoder.jp/contests/abc233/tasks/abc233_d
交错前缀和 https://codeforces.com/contest/1915/problem/E
https://codeforces.com/problemset/problem/1446/D1 2600 转换
https://www.luogu.com.cn/problem/AT_joisc2014_h 三个字母映射到一些大整数上，从而区分开

前缀和思想 LC1523 https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/
有点数形结合 https://codeforces.com/problemset/problem/1748/C

前缀和的前缀和（二重前缀和）
LC2281 https://leetcode.cn/problems/sum-of-total-strength-of-wizards/
https://atcoder.jp/contests/abc058/tasks/arc071_b

前缀和+异或
模 3 & 字符集大小为 n https://codeforces.com/problemset/problem/1418/G 2500
https://atcoder.jp/contests/abc295/tasks/abc295_d
https://ac.nowcoder.com/acm/contest/75174/E

https://leetcode.cn/problems/find-longest-subarray-lcci/
https://codeforces.com/problemset/problem/1296/C

前后缀分解
部分题目也可以用状态机 DP 解决
- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)（[视频讲解](https://www.bilibili.com/video/BV1Qg411q7ia/?t=3m05s)）
  注：带修改的接雨水 https://codeforces.com/gym/104821/problem/M
  - https://www.zhihu.com/question/627281278/answer/3280684055
  - 全排列接雨水 https://atcoder.jp/contests/tenka1-2015-final/tasks/tenka1_2015_final_e
- [926. 将字符串翻转到单调递增](https://leetcode.cn/problems/flip-string-to-monotone-increasing/)
  - https://codeforces.com/problemset/problem/180/C 1400
  - https://codeforces.com/problemset/problem/846/A 1500
https://codeforces.com/problemset/problem/1178/B 1300
https://codeforces.com/problemset/problem/1443/B 1300
https://codeforces.com/problemset/problem/2026/B 1300 做到 O(n)
https://codeforces.com/problemset/problem/1706/C 1400
https://codeforces.com/problemset/problem/1980/D 1400 pairwise 的前后缀分解
https://codeforces.com/problemset/problem/2008/E 1500
https://codeforces.com/problemset/problem/1029/C 1600
https://codeforces.com/problemset/problem/2028/C 1600
https://codeforces.com/problemset/problem/2031/D 1700
https://codeforces.com/problemset/problem/1957/D 1900
https://codeforces.com/problemset/problem/1969/D 1900
https://codeforces.com/problemset/problem/1837/F 2400
https://codeforces.com/problemset/problem/2005/D 2400 GCD logTrick
昆明 2024：至多修改一个子数组 [L,R] ：把元素都加上 k，最大化整个数组的 GCD
- 预处理前后缀 GCD，由于前缀 GCD 只有 O(logU) 个不同的值，可以只枚举 O(logU) 个 L 和 O(n) 个 R，
- 枚举 R 的同时计算修改后的子数组 GCD，然后和前后缀 GCD 求 GCD

定长滑动窗口
https://codeforces.com/problemset/problem/716/B 1300
https://codeforces.com/problemset/problem/1955/D 1400
https://codeforces.com/problemset/problem/608/B 1500
https://codeforces.com/problemset/problem/1687/A 1600
https://codeforces.com/problemset/problem/69/E 1800
https://codeforces.com/problemset/problem/371/E 2000

不定长滑动窗口（求最长/最大）
- [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
   - 翻转至多一个任意子串后的无重复字符的最长子串 https://codeforces.com/contest/1234/problem/F
- [1658. 将 x 减到 0 的最小操作数](https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/) 1817
    - https://codeforces.com/problemset/problem/1692/E 1200
与单调队列结合 https://www.luogu.com.cn/problem/P3594
https://codeforces.com/problemset/problem/1873/F 1300
https://codeforces.com/problemset/problem/1794/C 1300 式子变形

不定长滑动窗口（求最短/最小）
https://codeforces.com/problemset/problem/1354/B 1200
https://codeforces.com/problemset/problem/224/B 1500 和最小
https://codeforces.com/problemset/problem/701/C 1500
https://codeforces.com/problemset/problem/1777/C 1700

不定长滑动窗口（求子数组个数）
https://atcoder.jp/contests/abc130/tasks/abc130_d 和至少为 k 的子数组个数
https://codeforces.com/problemset/problem/550/B 变形：改成子数组
其它题目见【前缀和】

滑窗的同时维护数据
https://codeforces.com/problemset/problem/898/D 1600

多指针
https://codeforces.com/problemset/problem/1971/F 1600

LC2234 https://leetcode.cn/problems/maximum-total-beauty-of-the-gardens/ 2562
类似 [795. 区间子数组个数](https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/) 1817
入门题 https://codeforces.com/problemset/problem/602/B
入门题 https://codeforces.com/problemset/problem/279/B
https://atcoder.jp/contests/abc229/tasks/abc229_d
LC2271 毯子覆盖的最多白色砖块数 需要多思考一点点 https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/
- https://atcoder.jp/contests/abc098/tasks/arc098_b
较为复杂 https://atcoder.jp/contests/abc294/tasks/abc294_e
      - https://ac.nowcoder.com/acm/contest/62033/D
https://codeforces.com/problemset/problem/1208/B
https://codeforces.com/problemset/problem/1765/D
多指针 https://codeforces.com/problemset/problem/895/B
https://codeforces.com/contest/1833/problem/F
计算有多少子数组，其中有至少 k 个相同的数 https://codeforces.com/problemset/problem/190/D
https://codeforces.com/problemset/problem/165/C
- [1099. 小于 K 的两数之和](https://leetcode.cn/problems/two-sum-less-than-k/)（会员题）

单序列双指针
- [2972. 统计移除递增子数组的数目 II](https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii/) 2153
     - https://codeforces.com/problemset/problem/1167/E 2100
https://codeforces.com/contest/2032/problem/C 1400

双序列双指针
LC88 https://leetcode.cn/problems/merge-sorted-array/
LC360（背向双指针）https://leetcode.cn/problems/sort-transformed-array/
- [986. 区间列表的交集](https://leetcode.cn/problems/interval-list-intersections/) 1542
- [1537. 最大得分](https://leetcode.cn/problems/get-the-maximum-score/) 1961
https://codeforces.com/contest/489/problem/B 1200

MEX
https://codeforces.com/problemset/problem/1793/D 1800
https://atcoder.jp/contests/abc194/tasks/abc194_e

相向双指针
题单 https://leetcode.cn/leetbook/read/sliding-window-and-two-pointers/odt2yh/
LC2824 https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
LC923 https://leetcode.cn/problems/3sum-with-multiplicity/
https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/C

同时用到同向双指针和相向双指针的题
https://atcoder.jp/contests/abc155/tasks/abc155_d
- 相似题目 https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/

a[i] + b[j] = target 的方案数
a[i] + b[j] < target 的方案数    相向双指针 https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
                                         https://codeforces.com/problemset/problem/1538/C 1300
                               - [259. 较小的三数之和](https://leetcode.cn/problems/3sum-smaller/)（会员题）
a[i] + b[j] > target 的方案数    同上
a[i] - b[j] = target 的方案数
a[i] - b[j] < target 的方案数    滑窗
a[i] - b[j] > target 的方案数    同上 >= https://atcoder.jp/contests/abc353/tasks/abc353_c
子数组元素和 = < > target 的方案数：用前缀和，转换成上面 a[i] - b[j] 的形式
子序列元素和 = < > target 的方案数：0-1 背包恰好/至多/至少，见 https://www.bilibili.com/video/BV16Y411v7Y6/ 末尾的总结

分组循环
https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/
**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。
**核心思想**：
- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
- 内层循环负责遍历组，找出这一组最远在哪结束。
这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。
见题单第六章 https://leetcode.cn/circle/discuss/0viNMK/
LC1180（会员）https://leetcode.cn/problems/count-substrings-with-only-one-distinct-letter/
LC2257 https://leetcode.cn/problems/count-unguarded-cells-in-the-grid/
- https://atcoder.jp/contests/abc317/tasks/abc317_e
LC2495（会员）逆向思维 https://leetcode.cn/problems/number-of-subarrays-having-even-product/
https://codeforces.com/problemset/problem/1272/C 1200
https://codeforces.com/problemset/problem/1343/C 1200
https://codeforces.com/problemset/problem/1821/C 1300 枚举答案
https://codeforces.com/problemset/problem/1873/F 1300
https://codeforces.com/problemset/problem/363/C 1400 分组循环的分组循环
https://codeforces.com/problemset/problem/1380/C 1400
https://codeforces.com/problemset/problem/620/C 1500
https://codeforces.com/problemset/problem/525/C 1600
https://codeforces.com/problemset/problem/1748/C 1600
https://codeforces.com/problemset/problem/1849/D 1700
https://codeforces.com/problemset/problem/2031/D 1700

哨兵
- [1465. 切割后面积最大的蛋糕](https://leetcode.cn/problems/maximum-area-of-a-piece-of-cake-after-horizontal-and-vertical-cuts/) 1445
- [2975. 移除栅栏得到的正方形田地的最大面积](https://leetcode.cn/problems/maximum-square-area-by-removing-fences-from-a-field/) 1873
不是哨兵，但图像类似 [2943. 最大化网格图中正方形空洞的面积](https://leetcode.cn/problems/maximize-area-of-square-hole-in-grid/) 1677

巧妙枚举
LC939 https://leetcode.cn/problems/minimum-area-rectangle/
- [1577. 数的平方等于两数乘积的方法数](https://leetcode.cn/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/) 1594
LC3234 https://leetcode.cn/problems/count-the-number-of-substrings-with-dominant-ones/
- 这个代码的时间复杂度怎么证明？https://leetcode.cn/circle/discuss/GNUiDD/view/wPi6zR/
https://codeforces.com/problemset/problem/846/C 1800
https://codeforces.com/problemset/problem/1181/C 1900
https://codeforces.com/problemset/problem/1626/D 2100
https://codeforces.com/problemset/problem/339/E 2700

贪心及其证明
- [881. 救生艇](https://leetcode.cn/problems/boats-to-save-people/) 1530
    - https://codeforces.com/problemset/problem/1690/E
    - https://www.lanqiao.cn/problems/4174/learning/?contest_id=135
    - https://codeforces.com/problemset/problem/1765/D
- [2350. 不可能得到的最短骰子序列](https://leetcode.cn/problems/shortest-impossible-sequence-of-rolls/) 1961
    - https://codeforces.com/problemset/problem/1924/A 1500 输出方案
- [1686. 石子游戏 VI](https://leetcode.cn/problems/stone-game-vi/) 2001
    - https://codeforces.com/contest/1914/problem/E2 1400
- todo 复习 [2193. 得到回文串的最少操作次数](https://leetcode.cn/problems/minimum-number-of-moves-to-make-palindrome/) 2091
- todo 复习 [659. 分割数组为连续子序列](https://leetcode.cn/problems/split-array-into-consecutive-subsequences/)
https://codeforces.com/problemset/problem/1920/B 1100
https://codeforces.com/problemset/problem/1993/B 1100
https://codeforces.com/problemset/problem/2047/B ~1200 做到 O(n)
https://codeforces.com/problemset/problem/545/D 1300
https://codeforces.com/problemset/problem/1443/B 1300
https://codeforces.com/problemset/problem/1498/B 1300 从大到小贪心
https://codeforces.com/problemset/problem/1902/C 1300
https://codeforces.com/problemset/problem/388/A 1400
https://codeforces.com/problemset/problem/437/C 1400 排序不等式/交换论证法
https://codeforces.com/problemset/problem/492/C 1400
https://codeforces.com/problemset/problem/1369/C 1400
	提示 1：前 k 大的数一定可以作为最大值。且尽量把大的数放在 w[i] = 1 的组中，这样可以计入答案两次。
	如果某个前 k 大的数 x 没有作为最大值（其中一个组的最大值是不在前 k 大中的 y），那么把 x 和 y 交换，
	如果 x 是某个组的最小值，那么交换后 y 必然也是最小值，此时答案不变。
	如果 x 不是某个组的最小值（这个组的最小值是 z）：
		   如果 y 交换后变成了最小值，那么答案变大了 x-z。
		   如果 y 交换后也不是最小值，那么答案变大了 x-y。
	无论如何，这样交换都不会使答案变小，因此前 k 大的数一定可以作为最大值。
	提示 2：然后来说最小值。a 的最小值必然要分到某个组中，为了「跳过」尽量多的较小的数，优先把 a 中较小的数分到 w 较大的组中。所以 a 从小到大遍历，w 从大到小遍历。
https://codeforces.com/problemset/problem/1443/C 1400
https://codeforces.com/problemset/problem/1691/C 1400
https://codeforces.com/problemset/problem/1895/C 1400
https://codeforces.com/problemset/problem/1896/C 1400
https://codeforces.com/problemset/problem/864/D 1500
https://codeforces.com/problemset/problem/985/C 1500
https://codeforces.com/problemset/problem/1659/C 1500
https://codeforces.com/problemset/problem/1759/E 1500
https://codeforces.com/problemset/problem/1873/G 1500
https://codeforces.com/problemset/problem/1924/A 1500
https://codeforces.com/problemset/problem/913/C 1600
https://codeforces.com/problemset/problem/1707/A 1600 倒序思维
https://codeforces.com/problemset/problem/1157/C2 1700
https://codeforces.com/problemset/problem/1661/C 1700 奇数天+1 偶数天 +2
https://codeforces.com/problemset/problem/1995/B2 1700
https://codeforces.com/problemset/problem/2035/D 1800
https://codeforces.com/problemset/problem/2042/C 1800
https://codeforces.com/problemset/problem/3/B 1900
https://codeforces.com/problemset/problem/1479/B1 1900
https://codeforces.com/problemset/problem/1804/D 2000
https://codeforces.com/problemset/problem/1029/E 2100 树
https://codeforces.com/problemset/problem/1479/B2 2100
    https://www.luogu.com.cn/blog/wsyhb/post-ti-xie-cf1479b1-painting-the-array-i
https://codeforces.com/problemset/problem/442/C 2500
    如果 x>=y<=z，那么删除 y 最优
    结束后剩下一个长为 m 的 /\ 形状的序列，由于无法取到最大值和次大值，那么加上剩下最小的 m-2 个数
https://atcoder.jp/contests/arc147/tasks/arc147_e 难
https://www.luogu.com.cn/problem/P1016
https://www.luogu.com.cn/problem/UVA11384 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=25&page=show_problem&problem=2379

数学思维
https://codeforces.com/problemset/problem/23/C 2500
- https://codeforces.com/problemset/problem/798/D 2400

乘法贪心
https://codeforces.com/problemset/problem/45/I 1400
https://codeforces.com/problemset/problem/934/A 1400
https://atcoder.jp/contests/abc173/tasks/abc173_e 1607=CF1926 k 个数的最大乘积
删除一个数后，最小化 k 个数的最大乘积

中位数贪心（右边数字为难度分） // 注：算长度用左闭右开区间思考，算中间值用闭区间思考    两个中位数分别是 a[(n-1)/2] 和 a[n/2]
有两种证明方法，见 https://leetcode.cn/problems/5TxKeK/solution/zhuan-huan-zhong-wei-shu-tan-xin-dui-din-7r9b/
【思考题】插入一个数再选定一个 x，每次操作 +x/-x，计算最小操作次数
https://codeforces.com/problemset/problem/710/B 1400
中位数相关 https://codeforces.com/problemset/problem/166/C 1500 *可以做到对不同的 x 用 O(log n) 时间回答

排序不等式
https://codeforces.com/problemset/problem/276/C 1500
https://codeforces.com/problemset/problem/1165/E 1600

相邻不同
每次取两个数减一，最后剩下的数最小 / 操作次数最多 https://cs.stackexchange.com/a/145450
- [1953. 你可以工作的最大周数](https://leetcode.cn/problems/maximum-number-of-weeks-for-which-you-can-work/) 1804
   - https://codeforces.com/problemset/problem/1579/D 1400
https://codeforces.com/problemset/problem/1521/E 2700 二维+对角不同

每次取数组中大于 0 的连续一段同时减 1，求使数组全为 0 的最少操作次数
https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/solutions/371326/xing-cheng-mu-biao-shu-zu-de-zi-shu-zu-zui-shao-ze/
https://codeforces.com/problemset/problem/448/C

邻项交换（最小代价排序/字典序最小） Exchange Arguments
https://codeforces.com/blog/entry/63533
某些题目和逆序对有关
LC1665. 完成所有任务的最少初始能量 https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/ 1901
- https://www.luogu.com.cn/problem/P3619
https://codeforces.com/problemset/problem/1638/B 1100
https://codeforces.com/problemset/problem/920/C 1400
https://codeforces.com/problemset/problem/435/B 1400
https://codeforces.com/contest/246/problem/A 900
https://atcoder.jp/contests/arc147/tasks/arc147_b
https://atcoder.jp/contests/abc268/tasks/abc268_f
相邻两数之差的绝对值为 1 https://ac.nowcoder.com/acm/contest/65259/C

非邻项交换（最小代价排序/字典序最小）
某些题目可以在 i 到 a[i] 之间连边建图
LC1202 https://leetcode.cn/problems/smallest-string-with-swaps/ 1855
LC2948 https://leetcode.cn/problems/make-lexicographically-smallest-array-by-swapping-elements/ 2047
https://codeforces.com/contest/252/problem/B
https://codeforces.com/problemset/problem/1768/D 1800
https://codeforces.com/contest/109/problem/D 2000
shift+reverse https://codeforces.com/contest/1907/problem/F

区间与点的最大匹配/覆盖问题
https://www.luogu.com.cn/problem/P2887
https://codeforces.com/problemset/problem/555/B
https://codeforces.com/problemset/problem/863/E

倒序
LC2718 https://leetcode.cn/problems/sum-of-matrix-after-queries/
- 加强版 https://www.luogu.com.cn/problem/P9715        ?contestId=126251

思维：观察、结论
- [2498. 青蛙过河 II](https://leetcode.cn/problems/frog-jump-ii/) 1759
- [782. 变为棋盘](https://leetcode.cn/problems/transform-to-chessboard/) 2430
https://codeforces.com/problemset/problem/1811/C 1100
https://codeforces.com/problemset/problem/1822/D 1200
https://codeforces.com/problemset/problem/1077/C 1300
https://codeforces.com/problemset/problem/1364/B 1300
https://codeforces.com/problemset/problem/1844/C 1300 假设答案是某些数之和，经过什么样的操作可以得到这些数？
https://codeforces.com/problemset/problem/1765/K 1500
https://codeforces.com/problemset/problem/1990/C 1500
https://codeforces.com/problemset/problem/1608/C 1700
https://codeforces.com/problemset/problem/1442/A 1800
https://codeforces.com/problemset/problem/558/C  1900
https://codeforces.com/problemset/problem/1744/F 2000
https://codeforces.com/problemset/problem/1610/E 2300
https://codeforces.com/problemset/problem/2004/F 2600

思维：脑筋急转弯
LC1503 https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/
LC2731 https://leetcode.cn/problems/movement-of-robots/
LC280 https://leetcode.cn/problems/wiggle-sort/
LC3012 https://leetcode.cn/problems/minimize-length-of-array-using-operations/
https://www.codechef.com/problems/CLEARARR 2037
https://codeforces.com/problemset/problem/2049/A 800
https://codeforces.com/problemset/problem/1632/B 1000 位运算 XOR
https://codeforces.com/problemset/problem/1708/B 1100
https://codeforces.com/problemset/problem/2044/D 1100
https://codeforces.com/problemset/problem/1009/B 1400
https://codeforces.com/problemset/problem/1883/F 1400
https://codeforces.com/problemset/problem/1904/C 1400
https://codeforces.com/problemset/problem/1012/A 1500
https://codeforces.com/problemset/problem/1169/B 1500
https://codeforces.com/problemset/problem/500/C 1600
https://codeforces.com/problemset/problem/601/A 1600
https://codeforces.com/problemset/problem/1763/C 2000
https://atcoder.jp/contests/abc194/tasks/abc194_e
https://atcoder.jp/contests/abc196/tasks/abc196_e
https://www.luogu.com.cn/problem/UVA10881 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1822
- [LCS 01. 下载插件](https://leetcode.cn/problems/Ju9Xwi/)

注意值域
LC2653 https://leetcode.cn/problems/sliding-subarray-beauty/ 1786
LC2250 https://leetcode.cn/problems/count-number-of-rectangles-containing-each-point/ 1998
LC2857 https://leetcode.cn/problems/count-pairs-of-points-with-distance-k/ 2082
LC1906 https://leetcode.cn/problems/minimum-absolute-difference-queries/ 2147
LC1766 https://leetcode.cn/problems/tree-of-coprimes/ 2232
LC2198 https://leetcode.cn/problems/number-of-single-divisor-triplets/（会员题）

注意指数/对数
LC2188 https://leetcode.cn/problems/minimum-time-to-finish-the-race/ 2315
LC2920 https://leetcode.cn/problems/maximum-points-after-collecting-coins-from-all-nodes/ 2351

枚举答案
https://codeforces.com/contest/1977/problem/C

构造
题单 https://www.luogu.com.cn/training/14#problems
LC767 https://leetcode.cn/problems/reorganize-string/
LC667 https://leetcode.cn/problems/beautiful-arrangement-ii/
LC2745 https://leetcode.cn/problems/construct-the-longest-new-string/ 1607
LC2573 https://leetcode.cn/problems/find-the-string-with-lcp/ 2682
LC3311 https://leetcode.cn/problems/construct-2d-grid-matching-graph-layout/
构造反例 https://leetcode.cn/problems/parallel-courses-iii/solution/tuo-bu-pai-xu-dong-tai-gui-hua-by-endles-dph6/2310439
构造 TLE 数据 https://leetcode.cn/problems/maximum-total-reward-using-operations-ii/solutions/2805413/bitset-you-hua-0-1-bei-bao-by-endlessche-m1xn/comments/2320111
https://codeforces.com/problemset/problem/1927/B   900
https://codeforces.com/problemset/problem/1772/C  1000
https://codeforces.com/problemset/problem/1998/B  1000
https://codeforces.com/problemset/problem/2037/C  1000
https://codeforces.com/problemset/problem/2039/B  1000 分析性质
https://codeforces.com/problemset/problem/2040/B  1000
https://codeforces.com/problemset/problem/2044/D  1100 脑筋急转弯
https://atcoder.jp/contests/keyence2020/tasks/keyence2020_c 625=CF1183
https://codeforces.com/problemset/problem/1028/B  1200
https://codeforces.com/problemset/problem/1713/C  1200
https://codeforces.com/problemset/problem/1717/C  1300
https://codeforces.com/problemset/problem/1788/C  1300
https://codeforces.com/problemset/problem/1815/A  1300
https://codeforces.com/problemset/problem/1978/C  1300
https://codeforces.com/problemset/problem/2031/C  1300 数学
https://codeforces.com/problemset/problem/803/A   1400
https://codeforces.com/problemset/problem/1838/C  1400
https://codeforces.com/problemset/problem/1863/D  1400
https://codeforces.com/problemset/problem/1896/C  1400
https://codeforces.com/problemset/problem/1974/D  1400
https://codeforces.com/problemset/problem/1630/A  1500
https://codeforces.com/problemset/problem/1710/A  1500
https://codeforces.com/problemset/problem/1722/G  1500
https://codeforces.com/problemset/problem/1809/C  1500
https://codeforces.com/problemset/problem/1968/E  1600
https://codeforces.com/problemset/problem/201/A   1700
https://codeforces.com/problemset/problem/584/C   1700 分类讨论
https://codeforces.com/problemset/problem/1332/D  1700 给你一个错误代码，构造 hack 数据
https://codeforces.com/problemset/problem/1893/B  1700
https://codeforces.com/problemset/problem/142/B   1800 棋盘放最多的马
https://codeforces.com/problemset/problem/847/C   1800
https://codeforces.com/problemset/problem/1156/B  1800 相邻字母在字母表中不相邻
https://codeforces.com/problemset/problem/1267/L  1800
https://codeforces.com/problemset/problem/1304/D  1800 最短/最长 LIS
https://codeforces.com/problemset/problem/1554/D  1800
https://codeforces.com/problemset/problem/1965/B  1800 二进制分解
https://codeforces.com/problemset/problem/118/C   1900 贪心
https://codeforces.com/problemset/problem/327/D   1900
https://codeforces.com/problemset/problem/388/B   1900 两点间恰好 k 条最短路径
https://codeforces.com/problemset/problem/550/D   1900 度数均为 k 且至少（恰好）有一条割边
https://codeforces.com/problemset/problem/708/B   1900 分类讨论
https://codeforces.com/problemset/problem/1823/D  1900
https://codeforces.com/problemset/problem/1854/A2 1900 分类讨论
https://codeforces.com/problemset/problem/2040/D  1900 树 质数
https://atcoder.jp/contests/arc088/tasks/arc088_b 1646=CF1956
https://codeforces.com/problemset/problem/515/D   2000
https://codeforces.com/problemset/problem/1558/C  2000
https://codeforces.com/problemset/problem/1789/D  2200 推荐 位运算 把 X 变成 Y 不断靠近答案
https://codeforces.com/problemset/problem/1761/E  2400
https://codeforces.com/problemset/problem/1227/G  2600 证明是亮点
https://codeforces.com/problemset/problem/1521/E  2700 二维相邻不同
https://codeforces.com/problemset/problem/1838/F  3000 交互 二分
https://atcoder.jp/contests/arc145/tasks/arc145_a
https://atcoder.jp/contests/agc015/tasks/agc015_d bit OR

不好想到的构造
https://codeforces.com/contest/1659/problem/D
https://atcoder.jp/contests/abc178/tasks/abc178_f
https://codeforces.com/problemset/problem/1689/E 脑筋急转弯
https://codeforces.com/problemset/problem/1787/E

不变量（想一想，操作不会改变什么）
https://codeforces.com/problemset/problem/1881/D 1300
https://codeforces.com/problemset/problem/1365/F 2100 仍然对称
https://codeforces.com/problemset/problem/1775/E 2100 有点差分的味道，想想前缀和
https://atcoder.jp/contests/arc119/tasks/arc119_c 操作不影响交错和

不变量 2（总和）
把一个环形数组切两刀，分成两段，要求相等，求方案数 => 和为 sum(a)/2 的子数组个数
LC494 https://leetcode.cn/problems/target-sum/

行列独立
LC3189 https://leetcode.cn/problems/minimum-moves-to-get-a-peaceful-board/

分类讨论（部分题是易错题）
https://codeforces.com/problemset/problem/2039/B 1000
https://codeforces.com/problemset/problem/1364/A 1200
https://codeforces.com/problemset/problem/870/C 1300
https://codeforces.com/problemset/problem/1698/C 1300
https://codeforces.com/problemset/problem/30/A 1400
https://codeforces.com/problemset/problem/45/I 1400
https://codeforces.com/problemset/problem/489/C 1400
https://codeforces.com/problemset/problem/934/A 1400
https://codeforces.com/problemset/problem/1009/B 1400 脑筋急转弯
https://codeforces.com/problemset/problem/1251/B 1400
https://codeforces.com/problemset/problem/1292/A 1400 也有简单写法
https://codeforces.com/problemset/problem/1605/C 1400
https://codeforces.com/problemset/problem/115/B 1500
https://codeforces.com/problemset/problem/960/B 1500
https://codeforces.com/problemset/problem/1051/C 1500
https://codeforces.com/problemset/problem/1180/B 1500
https://codeforces.com/problemset/problem/1250/L 1500
https://codeforces.com/problemset/problem/750/C 1600 *也有偏数学的做法
https://codeforces.com/problemset/problem/898/E 1600
https://codeforces.com/problemset/problem/1822/E 1600 样例给的挺良心的
https://codeforces.com/problemset/problem/1861/C 1600 好题！
https://codeforces.com/problemset/problem/1976/C 1600
https://codeforces.com/problemset/problem/1978/D 1600
https://atcoder.jp/contests/abc173/tasks/abc173_e 1607=CF1926 乘法
https://codeforces.com/problemset/problem/193/A 1700
https://codeforces.com/problemset/problem/382/C 1700
https://codeforces.com/problemset/problem/411/C 1700
https://codeforces.com/problemset/problem/1516/C 1700
https://codeforces.com/problemset/problem/1799/C 1700
https://codeforces.com/problemset/problem/1468/J 1800 MST
https://codeforces.com/problemset/problem/1833/G 1800 样例给的挺良心的
https://codeforces.com/problemset/problem/796/C 1900
https://codeforces.com/problemset/problem/1095/E 1900
https://codeforces.com/problemset/problem/1714/F 1900 锻炼代码实现技巧的好题
https://codeforces.com/problemset/problem/1914/F 1900
https://codeforces.com/problemset/problem/1088/D 2000
https://codeforces.com/problemset/problem/1763/C 2000
https://codeforces.com/problemset/problem/1978/E 2000
https://codeforces.com/problemset/problem/2051/F 2000
https://codeforces.com/problemset/problem/1811/F 2100
https://codeforces.com/problemset/problem/1798/E 2300
https://codeforces.com/problemset/problem/209/C 2400
https://codeforces.com/problemset/problem/1594/F 2400
https://codeforces.com/problemset/problem/1736/C2 2400
https://codeforces.com/problemset/problem/1761/E 2400
https://codeforces.com/problemset/problem/1832/D2 2400
https://codeforces.com/problemset/problem/599/E 2600
https://codeforces.com/problemset/problem/1016/F 2600
https://codeforces.com/problemset/problem/1730/E 2700
https://codeforces.com/gym/105139/problem/L
https://atcoder.jp/contests/diverta2019/tasks/diverta2019_c
https://atcoder.jp/contests/abc155/tasks/abc155_d
https://atcoder.jp/contests/abc125/tasks/abc125_d
https://atcoder.jp/contests/arc134/tasks/arc134_d 1998
- [335. 路径交叉](https://leetcode.cn/problems/self-crossing/)
- [2162. 设置时间的最少代价](https://leetcode.cn/problems/minimum-cost-to-set-cooking-time/) 1852
https://leetcode.cn/problems/maximize-the-number-of-partitions-after-operations/
https://leetcode.cn/problems/count-the-number-of-houses-at-a-certain-distance-ii/

大量分类讨论
https://codeforces.com/problemset/problem/2045/A 1700 做到 O(n)
https://codeforces.com/problemset/problem/796/C 1900
https://codeforces.com/problemset/problem/1647/D 1900
https://codeforces.com/problemset/problem/356/C 2100
https://codeforces.com/problemset/problem/460/D 2300
https://codeforces.com/problemset/problem/1527/D 2400
https://codeforces.com/problemset/problem/1374/E2 2500
https://atcoder.jp/contests/arc153/tasks/arc153_c +构造
https://atcoder.jp/contests/agc015/tasks/agc015_d
https://atcoder.jp/contests/abc164/tasks/abc164_f
LC420 https://leetcode.cn/problems/strong-password-checker/
LC3348 https://leetcode.cn/problems/smallest-divisible-digit-product-ii/ 可以避免大量分类讨论
LC3366 https://leetcode.cn/problems/minimum-array-sum/ 做到 O(nlogn)
LC1534 https://leetcode.cn/problems/count-good-triplets/ 做到 O(nlogU) 或者 O(nlogn)

贡献法
见数学题单 §2.5 节 https://leetcode.cn/circle/discuss/IYT3ss/
LC2681 https://leetcode.cn/problems/power-of-heroes/
- https://atcoder.jp/contests/arc116/tasks/arc116_b
LC2763 https://leetcode.cn/problems/sum-of-imbalance-numbers-of-all-subarrays/
- https://atcoder.jp/contests/abc390/tasks/abc390_f
更多贡献法题目，见 monotone_stack.go
https://codeforces.com/problemset/problem/2019/B 1200
https://codeforces.com/problemset/problem/1648/A 1400
https://codeforces.com/problemset/problem/1691/C 1400
https://codeforces.com/problemset/problem/1789/C 1500 好题！
https://codeforces.com/problemset/problem/383/A 1600 好题
https://codeforces.com/problemset/problem/1165/E 1600
https://codeforces.com/problemset/problem/1715/C 1700 也可以用增量法
https://atcoder.jp/contests/abc356/tasks/abc356_e 1506=CF1700
https://codeforces.com/problemset/problem/1777/D 1900 树
https://codeforces.com/problemset/problem/1788/D 2000 好题！
https://atcoder.jp/contests/abc390/tasks/abc390_f 1801=CF2073
https://codeforces.com/problemset/problem/912/D 2100
https://codeforces.com/problemset/problem/1808/D 2100
https://codeforces.com/problemset/problem/520/E 2200
https://codeforces.com/problemset/problem/1208/E 2200
https://codeforces.com/problemset/problem/2063/E 2300
https://codeforces.com/problemset/problem/749/E 2400
https://codeforces.com/problemset/problem/915/F 2400
https://codeforces.com/problemset/problem/2004/F 2600
https://atcoder.jp/contests/abc290/tasks/abc290_e 好题！
https://atcoder.jp/contests/abc159/tasks/abc159_f 与 0-1 背包结合
^+ https://atcoder.jp/contests/abc201/tasks/abc201_e
https://atcoder.jp/contests/abc384/tasks/abc384_f 恰好 -> 至少
https://atcoder.jp/contests/abc127/tasks/abc127_e 1938
https://atcoder.jp/contests/abc104/tasks/abc104_d 1739=CF2026
- 抄袭？https://codeforces.com/problemset/problem/1426/F 2000
https://www.lanqiao.cn/problems/12467/learning/?contest_id=167
https://codeforces.com/group/BJlsDCvlJO/contest/586547/problem/I

增量法
- [2262. 字符串的总引力](https://leetcode.cn/problems/total-appeal-of-a-string/) 2033
      结合线段树优化 DP https://codeforces.com/contest/833/problem/B 2200
- [828. 统计子串中的唯一字符](https://leetcode.cn/problems/count-unique-characters-of-all-substrings-of-a-given-string/) 2034
- [2916. 子数组不同元素数目的平方和 II](https://leetcode.cn/problems/subarrays-distinct-element-sum-of-squares-ii/) 2816
https://codeforces.com/problemset/problem/1715/C 1700 也可以用贡献法
https://codeforces.com/problemset/problem/1428/F 2400

小模拟
LC2534 https://leetcode.cn/problems/time-taken-to-cross-the-door/
https://atcoder.jp/contests/abc279/tasks/abc279_f

中模拟
https://atcoder.jp/contests/abc319/tasks/abc319_f

其他
删除一个字符 + 删除最长连续前缀 https://codeforces.com/problemset/problem/1430/D
https://codeforces.com/problemset/problem/521/D

先撤销，再恢复
LC3187 https://leetcode.cn/problems/peaks-in-array/

合法括号字符串 (Regular Bracket Sequence, RBS)
https://codeforces.com/problemset/problem/1097/C 1400
https://codeforces.com/problemset/problem/1837/D 1400
https://codeforces.com/problemset/problem/990/C 1500
https://codeforces.com/problemset/problem/847/C 1800 构造
https://codeforces.com/problemset/problem/1821/E 2100
https://codeforces.com/problemset/problem/1830/C 2400
https://codeforces.com/problemset/problem/3/D 2600 反悔贪心（反悔堆）

= 变成 <= 或者 >=
求前缀和/后缀和
https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-substrings/

连续性 + 上下界
https://atcoder.jp/contests/arc137/tasks/arc137_b
https://codeforces.com/contest/1695/problem/C

异类双变量：固定某变量统计另一变量的 [0,n)
    EXTRA: 值域上的双变量，见 https://codeforces.com/contest/486/problem/D
同类双变量①：固定 i 统计 [0,n)
同类双变量②：固定 i 统计 [0,i-1]
套路：预处理数据（按照某种顺序排序/优先队列/BST/...），或者边遍历边维护，
     然后固定变量 i，用均摊 O(1)~O(logn) 的复杂度统计范围内的另一变量 j
这样可以将复杂度从 O(n^2) 降低到 O(n) 或 O(nlogn)
https://codeforces.com/problemset/problem/1194/E
进阶：https://codeforces.com/problemset/problem/1483/D
删除一段的最长连续递增 CERC10D https://codeforces.com/gym/101487
统计量是二元组的情形 https://codeforces.com/problemset/problem/301/D
好题 空间优化 https://codeforces.com/contest/1830/problem/B

双变量+下取整：枚举分母，然后枚举分子的范围，使得在该范围内的分子/分母是一个定值
LC1862 https://leetcode.cn/problems/sum-of-floored-pairs/
https://codeforces.com/problemset/problem/1706/D2

利用前缀和实现巧妙的构造 https://www.luogu.com.cn/blog/duyi/qian-zhui-he
邻项修改->前缀和->单项修改 https://codeforces.com/problemset/problem/1254/B2 https://ac.nowcoder.com/acm/contest/7612/C

二进制枚举
https://www.luogu.com.cn/problem/UVA11464 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=26&page=show_problem&problem=2459
横看成岭侧成峰
转换为距离的众数 https://codeforces.com/problemset/problem/1365/C
转换为差分数组 https://codeforces.com/problemset/problem/1110/E 2200
             https://codeforces.com/problemset/problem/1442/A 1800
             https://codeforces.com/problemset/problem/1700/C 1700
             https://codeforces.com/problemset/problem/1779/D 改成修改长为 x 的数组？
             https://www.luogu.com.cn/problem/P4552
转换为差 http://www.51nod.com/Challenge/Problem.html#problemId=1217
考虑每个点产生的贡献 https://codeforces.com/problemset/problem/1009/E
考虑每条边产生的负贡献 https://atcoder.jp/contests/abc173/tasks/abc173_f
考虑符合范围要求的贡献 https://codeforces.com/problemset/problem/1151/E
和式的另一视角。若每一项的值都在一个范围，不妨考虑另一个问题：值为 x 的项有多少个？https://atcoder.jp/contests/abc162/tasks/abc162_e
对所有排列考察所有子区间的性质，可以转换成对所有子区间考察所有排列。将子区间内部的排列和区间外部的排列进行区分，内部的性质单独研究，外部的当作 (n-(r-l))! 个排列 https://codeforces.com/problemset/problem/1284/C
从最大值入手 https://codeforces.com/problemset/problem/1381/B
等效性 LC1183 https://leetcode.cn/problems/maximum-number-of-ones/
LC1526 https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
置换 https://atcoder.jp/contests/abc250/tasks/abc250_e
排序+最小操作次数 https://codeforces.com/contest/1367/problem/F2
https://codeforces.com/contest/1830/problem/A
从绝对值最大的开始思考 https://codeforces.com/contest/351/problem/E
https://codeforces.com/problemset/problem/777/C 1600

棋盘染色 LC2577 https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/
        https://codeforces.com/contest/1848/problem/A

others https://codeforces.com/blog/entry/118706

离线
由于所有的询问数据都给出了，我们可以通过修改询问的顺序，达到降低时间复杂度的效果。相应的，在线算法就是按照输入的顺序处理，来一个处理一个。
https://codeforces.com/problemset/problem/1932/C 1400
https://atcoder.jp/contests/abc375/tasks/abc375_f
https://atcoder.jp/contests/abc379/tasks/abc379_f

逆向思维 / 正难则反
不可行方案通常比可行方案好求
- [2171. 拿出最少数目的魔法豆](https://leetcode.cn/problems/removing-minimum-number-of-magic-beans/) 1748
- [1354. 多次求和构造目标数组](https://leetcode.cn/problems/construct-target-array-with-multiple-sums/) 2015
LC803 https://leetcode.cn/problems/bricks-falling-when-hit/
LC936 https://leetcode.cn/problems/stamping-the-sequence/
LC1199 https://leetcode.cn/problems/minimum-time-to-build-blocks/
LC2382 https://leetcode.cn/problems/maximum-segment-sum-after-removals/
LCP52 https://leetcode.cn/problems/QO5KpG/
https://codeforces.com/problemset/problem/1792/C 1500
- 相似题目 https://codeforces.com/problemset/problem/1367/F1 2100
https://codeforces.com/problemset/problem/1882/B
https://codeforces.com/problemset/problem/712/C 1600
https://codeforces.com/problemset/problem/621/C 1700
https://codeforces.com/problemset/problem/1301/C 1700
https://codeforces.com/problemset/problem/1644/D 1700
https://codeforces.com/problemset/problem/1672/D 1700
https://codeforces.com/problemset/problem/1759/G 1900 求字典序最小，通常可以从大往小思考
https://codeforces.com/problemset/problem/1638/D 2000
https://codeforces.com/problemset/problem/571/A 2100
https://codeforces.com/problemset/problem/1919/D 2100
https://codeforces.com/problemset/problem/369/E 2200

删除变添加
https://codeforces.com/problemset/problem/295/B
https://leetcode.cn/problems/maximum-segment-sum-after-removals/

奇偶性
https://codeforces.com/problemset/problem/763/B
https://codeforces.com/problemset/problem/1270/E
https://codeforces.com/problemset/problem/1332/E 配对法：将合法局面与非法局面配对
LC932 https://leetcode.cn/problems/beautiful-array/ 分治

相邻 传递性
https://codeforces.com/problemset/problem/1582/E

归纳：solve(n)->solve(n-1) 或者 solve(n-1)->solve(n)
https://codeforces.com/problemset/problem/1517/C
https://codeforces.com/problemset/problem/412/D
https://codeforces.com/problemset/problem/266/C

见微知著：考察单个点的规律，从而推出全局规律
https://codeforces.com/problemset/problem/1510/K
LC1806 https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/ 1491

「恰好」转换成「至少/至多」
https://codeforces.com/problemset/problem/1188/C

反悔贪心
另见 heap.go 中的「反悔堆」
https://djy-juruo.blog.luogu.org/qian-tan-fan-hui-tan-xin
https://www.jvruo.com/archives/1844/
https://www.cnblogs.com/nth-element/p/11768155.html
题单 https://www.luogu.com.cn/training/8793
LC1388 双向链表反悔贪心 https://leetcode.cn/problems/pizza-with-3n-slices/
LC2813 https://leetcode.cn/problems/maximum-elegance-of-a-k-length-subsequence/

集合哈希
https://codeforces.com/problemset/problem/1394/B
https://www.luogu.com.cn/problem/P6688

操作树
和莫队类似，通过改变查询的顺序来优化复杂度
https://codeforces.com/problemset/problem/707/D

Golang 卡常技巧（注：关于 IO 的加速见 io.go）
对于存在海量小对象的情况（如 trie, treap 等），使用 debug.SetGCPercent(-1) 来禁用 GC，能明显减少耗时
对于可以回收的情况（如 append 在超过 cap 时），使用 debug.SetGCPercent(-1) 虽然会减少些许耗时，但若有大量内存没被回收，会有 MLE 的风险
其他情况下使用 debug.SetGCPercent(-1) 对耗时和内存使用无明显影响
对于多组数据的情况，若禁用 GC 会 MLE，可在每组数据的开头或末尾调用 runtime.GC() 或 debug.FreeOSMemory() 手动 GC
参考 https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/
    https://zhuanlan.zhihu.com/p/77943973

128MB ~1e7 个 int64
256MB ~3e7 个 int64
512MB ~6e7 个 int64
1GB   ~1e8 个 int64

如果没有禁用 GC 但 MLE，可以尝试 1.19 新增的 debug.SetMemoryLimit
例如 debug.SetMemoryLimit(200<<20)，其中 200 可以根据题目的约束来修改
具体见如下测试：
180<<20 1996ms 255100KB https://codeforces.com/contest/1800/submission/203769679
195<<20  779ms 257800KB https://codeforces.com/contest/1800/submission/203768086
200<<20  654ms 259300KB https://codeforces.com/contest/1800/submission/203768768
205<<20  764ms 220100KB https://codeforces.com/contest/1800/submission/203771041
210<<20        MLE
参考 https://go.dev/doc/gc-guide#Memory_limit

对于二维矩阵，以 make([][mx]int, n) 的方式使用，比 make([][]int, n) 嵌套 make([]int, m) 更高效（100MB 以上时可以快 ~150ms）
但需要注意这种方式可能会向 OS 额外申请一倍的内存
对比 https://codeforces.com/problemset/submission/375/118043978
    https://codeforces.com/problemset/submission/375/118044262

函数内的递归 lambda 会额外消耗非常多的内存（100~200MB / 1e6 递归深度）
写在 main 里面 + slice MLE      https://codeforces.com/contest/767/submission/174193385
写在 main 外面 + slice 188364KB https://codeforces.com/contest/767/submission/174194380
附：
写在 main 里面 + array 257424KB https://codeforces.com/contest/767/submission/174194515
写在 main 外面 + array 154500KB https://codeforces.com/contest/767/submission/174193693

在特殊情况下，改为手动模拟栈可以减少 > 100MB 的内存
见这题的 Go 提交记录 https://codeforces.com/problemset/problem/163/E

测试：哈希表用时是数组的 13 倍（本题瓶颈）
slice    249ms https://codeforces.com/problemset/submission/570/209063267
hashmap 3259ms https://codeforces.com/problemset/submission/570/209063603

bool to int
int(*(*uint8)(unsafe.Pointer(&boolVal)))

[]int to []int64
*(*[]int64)(unsafe.Pointer(&nums))

*/

// slice 作为 map 的 key
// 长度为 0 的 slice 对应空字符串
func intSliceAsMapKeyExample(cnt map[string]int, a []int) {
	// 如果后面还会修改 a，可以先 copy 一份
	// a = slices.Clone(a)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	sh.Len *= bits.UintSize / 8 // 装作 byte slice
	s := *(*string)(unsafe.Pointer(sh))
	cnt[s]++
}

func _() {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	pow10 := func(x int) int { return int(math.Pow10(x)) } // 底层实现是查表，不需要 round

	// TIPS: dir4[i] 和 dir4[i^1] 互为相反方向
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右（网格）
	// TIPS: dir4[i] 和 dir4[i^2] 互为相反方向
	dir4 = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}     // 右下左上（网格，顺时针）
	dir4 = []struct{ x, y int }{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}     // 右上左下（网格，逆时针）
	dir4 = []struct{ x, y int }{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}     // 右下左上（坐标系，顺时针）
	dir4 = []struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}     // 右上左下（坐标系，逆时针）
	dir4R := []struct{ x, y int }{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}} // 斜向

	/* 方向
	- [1041. 困于环中的机器人](https://leetcode.cn/problems/robot-bounded-in-circle/) 1521
	- [874. 模拟行走机器人](https://leetcode.cn/problems/walking-robot-simulation/) 1846
	- [2069. 模拟行走机器人 II](https://leetcode.cn/problems/walking-robot-simulation-ii/) 1919
	- [3443. K 次修改后的最大曼哈顿距离](https://leetcode.cn/problems/maximum-manhattan-distance-after-k-changes/) ~1900
	*/
	dir4 = []struct{ x, y int }{'W': {-1, 0}, 'E': {1, 0}, 'S': {0, -1}, 'N': {0, 1}} // 西东南北（坐标系）
	dir4 = []struct{ x, y int }{'W': {0, -1}, 'E': {0, 1}, 'S': {1, 0}, 'N': {-1, 0}} // 西东南北（网格）
	dir4 = []struct{ x, y int }{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}} // 左右下上（坐标系）
	dir4 = []struct{ x, y int }{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}} // 左右下上（网格）

	dir8 := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}  // 逆时针（坐标系）
	dir8 = []struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}   // 顺时针（矩阵）
	dir8 = []struct{ x, y int }{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}} // 马走日

	// https://codeforces.com/problemset/problem/1983/C 1400
	perm3 := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	perm4 := [][]int{
		{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1},
		{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0},
		{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0},
		{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0},
	}

	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	/*
		关于上取整的计算，当 $a$ 和 $b$ 均为正整数时，我们有

		$$
		\left\lceil\dfrac{a}{b}\right\rceil = \left\lfloor\dfrac{a-1}{b}\right\rfloor + 1
		$$

		讨论 $a$ 被 $b$ 整除，和不被 $b$ 整除两种情况，可以证明上式的正确性。
	*/
	// - [1936. 新增的最少台阶数](https://leetcode.cn/problems/add-minimum-number-of-rungs/) 1323
	// - [1785. 构成特定和需要添加的最少元素](https://leetcode.cn/problems/minimum-elements-to-add-to-form-a-given-sum/) 1432
	ceil := func(a, b int) int {
		// assert a >= 0 && b > 0
		if a == 0 {
			return 0
		}
		return (a-1)/b + 1
	}
	// 另一种写法，无需考虑 a 为 0 的情况
	ceil = func(a, b int) int {
		return (a + b - 1) / b
	}

	// 顺时针旋转矩阵 90°
	// 返回一个拷贝
	// Python3: list(zip(*reversed(a))) 或者 list(zip(*a[::-1]))
	// 注：逆时针旋转矩阵 90° 就是 list(zip(*a))[::-1]
	rotateCopy := func(a [][]int) [][]int {
		n, m := len(a), len(a[0])
		b := make([][]int, m)
		for i := range b {
			b[i] = make([]int, n)
		}
		for i, r := range a {
			for j, v := range r {
				b[j][n-1-i] = v
			}
		}
		return b
	}
	rotate := rotateCopy

	// 转置
	transposeCopy := func(a [][]int) [][]int {
		n, m := len(a), len(a[0])
		b := make([][]int, m)
		for i := range b {
			b[i] = make([]int, n)
			for j, r := range a {
				b[i][j] = r[i]
			}
		}
		return b
	}

	// 按顺序从小到大生成所有回文数
	// https://oeis.org/A002113
	// LC2967 https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/
	// LC906 https://leetcode.cn/problems/super-palindromes/
	// LC2081 https://leetcode.cn/problems/sum-of-k-mirror-numbers/
	// LC3272 https://leetcode.cn/problems/find-the-count-of-good-integers/
	// EXTRA: 单个数字的情况 LC564 https://leetcode.cn/problems/find-the-closest-palindrome/
	// https://codeforces.com/problemset/problem/897/B 1300
	initPalindromeNumber := func() {
		const mx int = 1e9
		pal := []int{}

		// 哨兵。根据题目来定，也可以设置成 -2e9 等
		pal = append(pal, 0)

	outer:
		for base := 1; ; base *= 10 {
			// 生成奇数长度回文数，例如 base = 10，生成的范围是 101 ~ 999
			for i := base; i < base*10; i++ {
				x := i
				for t := i / 10; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
			// 生成偶数长度回文数，例如 base = 10，生成的范围是 1001 ~ 9999
			for i := base; i < base*10; i++ {
				x := i
				for t := i; t > 0; t /= 10 {
					x = x*10 + t%10
				}
				if x > mx {
					break outer
				}
				pal = append(pal, x)
			}
		}

		// 哨兵。根据 mx 调整，如果 mx 是 2e9 的话要写成 mx+2
		pal = append(pal, mx+1)
	}

	// 合并有序数组，保留重复元素
	// a b 必须是有序的（可以为空）
	merge := func(a, b []int) []int {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]int, 0, n+m)
		for {
			if i == n {
				return append(res, b[j:]...)
			}
			if j == m {
				return append(res, a[i:]...)
			}
			if a[i] < b[j] { // 改成 > 为降序
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
	}

	// 另一种写法
	merge2 := func(a, b []int) []int {
		res := make([]int, 0, len(a)+len(b))
		i, j := 0, 0
		for i < len(a) || j < len(b) {
			if j == len(b) || i < len(a) && a[i] < b[j] {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
		return res
	}

	// 合并有序数组，保留至多 k 个元素
	// https://codeforces.com/problemset/problem/587/C
	// https://codeforces.com/problemset/problem/1665/E
	mergeWithLimit := func(a, b []int, k int) []int {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]int, 0, min(n+m, k))
		for len(res) < k {
			if i == n {
				res = append(res, b[j:min(j+k-len(res), m)]...)
				break
			}
			if j == m {
				res = append(res, a[i:min(i+k-len(res), n)]...)
				break
			}
			if a[i] < b[j] {
				res = append(res, a[i])
				i++
			} else {
				res = append(res, b[j])
				j++
			}
		}
		return res
	}

	// 返回 a 的各个子集的元素和
	// 复杂度为 O(1+2+4+...+2^(n-1)) = O(2^n)
	// https://codeforces.com/contest/1209/problem/E2
	// LC3116 https://leetcode.cn/problems/kth-smallest-amount-with-single-denomination-combination/
	// LC3444 https://leetcode.cn/problems/minimum-increments-for-target-multiples-in-an-array/
	subsetSum := func(a []int) []int {
		sum := make([]int, 1<<len(a))
		// lcms[0] = 1
		for i, v := range a {
			highBit := 1 << i
			for mask, s := range sum[:highBit] {
				sum[highBit|mask] = s + v // + 可以换成其它运算，例如 gcd lcm 等
			}
		}
		return sum
	}

	// 另一种写法（benchmark 结果和上面差不多）
	subsetSum = func(a []int) []int {
		sum := make([]int, 1<<len(a))
		for i, v := range a {
			sum[1<<i] = v
		}
		for i := range sum {
			lb := i & -i // 随便取一个 i 中的元素（甚至非空真子集也可以）
			sum[i] = sum[i^lb] + sum[lb]
		}
		return sum
	}

	// 应用：给出由非负整数组成的数组 a 的子集和 sum，返回 a
	// 保证输入有解且 len(sum) = 2^len(a)
	// 变形：sum 包含负数 LC1982 https://leetcode.cn/problems/find-array-given-subset-sums/ 2872
	// 做法是给所有 sum[i] 加上 -min(sum)，这会导致：
	// - 若 sum[i] 包含负数 a[i]，则新的 sum'[i] 就不包含 a[i]
	// - 若 sum[i] 不包含负数 a[i]，则新的 sum'[i] 会包含 -a[i]
	// 所以新的 sum' 也一样有解
	// 对 sum' 求出 a'
	// 由于 -min(sum) 是 a 的所有负数之和，所以找到一个 a' 的子集和，若其等于 -min(sum)，则将该子集在 a' 中的元素取相反数，就得到了 a
	recoverArrayFromSubsetSum := func(sum []int) []int {
		slices.Sort(sum)
		n := bits.TrailingZeros(uint(len(sum)))
		skip := map[int]int{}
		ans := make([]int, 0, n)
		for j := 0; n > 0; n-- {
			for j++; skip[sum[j]] > 0; j++ {
				skip[sum[j]]--
			}
			s := sum[j]
			_s := make([]int, 1<<len(ans))
			for i, v := range ans {
				for m, b := 0, 1<<i; m < b; m++ {
					_s[b|m] = _s[m] + v
					skip[_s[b|m]+s]++
				}
			}
			ans = append(ans, s)
		}
		return ans
	}

	// 返回 a 的各个子集的元素和的排序后的结果
	// 若已求出前 i-1 个数的有序子集和 b，那么前 i 个数的有序子集和可以由 b 和 {b 的每个数加上 a[i]} 归并得到
	// 复杂度为 O(1+2+4+...+2^(n-1)) = O(2^n)
	// 参考 https://leetcode.cn/problems/closest-subsequence-sum/solution/o2n2de-zuo-fa-by-heltion-0yn7/
	subSumSorted := func(a []int) []int {
		sum := []int{0}
		for _, v := range a {
			b := make([]int, len(sum))
			for i, w := range sum {
				b[i] = w + v
			}
			sum = merge(sum, b)
		}
		return sum
	}

	// 前缀和
	// LC303 https://leetcode.cn/problems/range-sum-query-immutable/
	// https://codeforces.com/problemset/problem/1922/C 1300
	// https://codeforces.com/problemset/problem/2033/D 1300
	// https://codeforces.com/problemset/problem/1923/C 1400 构造
	// https://codeforces.com/problemset/problem/2009/F 1700
	// https://codeforces.com/problemset/problem/901/C 2300
	prefixSum := func(arr []int) {
		slices.Sort(arr)
		sum := make([]int, len(arr)+1)
		for i, v := range arr {
			sum[i+1] = sum[i] + v
		}

		// 返回 a 的所有数到 target 的距离之和，即 sum(abs(a[i]-target))
		// ！需要保证 a 是有序的
		// LC2602 https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/
		// - 原题是 https://atcoder.jp/contests/abc255/tasks/abc255_d
		distanceSum := func(target int) int {
			i := sort.SearchInts(arr, target)
			s1 := target*i - sum[i]
			s2 := sum[len(arr)] - sum[i] - target*(len(arr)-i)
			return s1 + s2
		}

		// LC2968 https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/
		// LC3086 https://leetcode.cn/problems/minimum-moves-to-pick-k-ones/ 2673
		// 返回下标在左闭右开区间 [left,right) 内的所有 a[i] 到 target 的距离之和
		// ！需要保证 a 是有序的
		distanceSumRange := func(left, right, target int) int {
			i := sort.SearchInts(arr, target)
			if i <= left { // target <= a[i] <= a[left]
				return sum[right] - sum[left] - target*(right-left)
			}
			if i >= right { // target > a[i-1] >= a[right-1]
				return target*(right-left) - (sum[right] - sum[left])
			}
			s1 := target*(i-left) - (sum[i] - sum[left])
			s2 := sum[right] - sum[i] - target*(right-i)
			return s1 + s2
		}

		// 返回 a 的所有数移动到 [low, high] 内的移动距离之和
		// ！需要保证 a 是有序的
		moveSum := func(low, high int) int {
			i := sort.SearchInts(arr, low)
			j := sort.SearchInts(arr, high+1)
			s1 := low*i - sum[i]
			s2 := sum[len(arr)] - sum[j] - high*(len(arr)-j)
			return s1 + s2
		}

		// +1 操作执行 maxOp 次，最大化最小值   fill
		// ！需要保证 a 是有序的
		// - [3081. 替换字符串中的问号使分数最小](https://leetcode.cn/problems/replace-question-marks-in-string-to-minimize-its-value/) 1905
		maxLow := func(maxOp int) int {
			i := sort.Search(len(arr), func(i int) bool { return arr[i]*i-sum[i] > maxOp })
			// low * i - sum[i] <= maxOp, i=n 时也适用
			low := (sum[i] + maxOp) / i
			// 注：如果存在 sum[i] + maxOp < 0 的情况，直接用 % i == 0 if else 实现
			return low
		}

		// -1 操作执行 maxOp 次，最小化最大值    remove
		// ！需要保证 a 是有序的
		// - [1300. 转变数组后最接近目标值的数组和](https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target/) 1607
		// 类似题目 https://codeforces.com/problemset/problem/1065/C 1600
		minHigh := func(maxOp int) int {
			n := len(arr)
			i := sort.Search(n, func(i int) bool { return sum[n]-sum[n-i]-arr[n-1-i]*i > maxOp })
			// 注意：这里的 i 表示，把最大的 i 个数减小到 high
			// sum[n] - sum[n-i] - high * i <= maxOp, i=n 时也适用
			high := (sum[n] - sum[n-i] - maxOp + i - 1) / i
			// 注：如果存在 sum[n] - sum[n-i] - maxOp < 0 的情况，直接用 % i == 0 if else 实现
			return high
		}

		// EXTRA: 青蛙跳井
		// 一次询问（下标从 1 开始）https://codeforces.com/problemset/problem/1141/E
		// 多次询问（下标从 0 开始）https://codeforces.com/problemset/problem/1490/G

		_ = []any{distanceSum, distanceSumRange, moveSum, maxLow, minHigh}
	}

	// 同余前缀和，a 的下标从 0 开始，md 为模数
	// 求 a[i]+a[i+md]+a[i+2*md]+...
	// 具体用法见 query 上的注释
	// LC1664 https://leetcode.cn/problems/ways-to-make-a-fair-array/
	// LC2902 https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/
	// https://atcoder.jp/contests/abc288/tasks/abc288_d
	groupPrefixSum := func(arr []int, md int) {
		_sum := make([]int, len(arr)+md)
		for i, v := range arr {
			_sum[i+md] = _sum[i] + v
		}
		_pre := func(x, t int) int {
			if x%md <= t {
				return _sum[x/md*md+t]
			}
			return _sum[(x+md-1)/md*md+t]
		}
		// 求下标在 [l,r) 范围内，并且下标模 md 同余于 rem 的所有元素之和
		query := func(l, r, rem int) int {
			rem %= md
			return _pre(r, rem) - _pre(l, rem) // % mod
		}
		_ = query
	}

	// 无限循环数组的区间和 [l,r)
	circularRangeSum := func(arr []int, md int) {
		_n := len(arr)
		sum := make([]int, _n+1)
		for i, v := range arr {
			sum[i+1] = sum[i] + v
		}
		pre := func(p int) int { return sum[_n]*(p/_n) + sum[p%_n] }
		// [l,r)
		query := func(l, r int) int { return pre(r) - pre(l) }

		_ = query
	}

	// 带权前缀和，权重是等差数列
	// 视频讲解：https://www.bilibili.com/video/BV1hQ4y1L7Tk/
	// https://codeforces.com/problemset/problem/1016/C 1800
	// https://codeforces.com/problemset/problem/1921/F 1900 +分块 
	// https://codeforces.com/problemset/problem/2026/D 1900
	// 二维版本 https://codeforces.com/problemset/problem/2044/H
	weightedPrefixSum := func(a []int) {
		n := len(a)
		sum := make([]int, n+1)
		iSum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
			iSum[i+1] = iSum[i] + (i+1)*v
		}
		query := func(l, r int) int { return iSum[r] - iSum[l] - l*(sum[r]-sum[l]) } // [l,r)

		_ = query
	}

	// 拆位前缀和
	bitSum := func(a []int) {
		m := bits.Len(uint(slices.Max(a)))
		sum := make([][]int, len(a))
		for i, v := range a {
			sum[i] = make([]int, m)
			for j := range sum[i] {
				sum[i+1][j] = sum[i][j] + v>>j&1
			}
		}

		// 返回 [l,r] 内的 AND
		queryAnd := func(l, r int) (res int) {
			for j, s := range sum[r+1] {
				if s-sum[l][j] == r-l+1 { // 全是 1
					res |= 1 << j
				}
			}
			return
		}

		// 返回 [l,r] 内的 OR
		queryOr := func(l, r int) (res int) {
			for j, s := range sum[r+1] {
				if s > sum[l][j] { // 有 1
					res |= 1 << j
				}
			}
			return
		}

		_ = []any{queryAnd, queryOr}
	}

	// 二维前缀和 sum2d
	// - [304. 二维区域和检索 - 矩阵不可变](https://leetcode.cn/problems/range-sum-query-2d-immutable/) *模板题
	// - [1314. 矩阵区域和](https://leetcode.cn/problems/matrix-block-sum/) 1484
	// - [3070. 元素和小于等于 k 的子矩阵的数目](https://leetcode.cn/problems/count-submatrices-with-top-left-element-and-sum-less-than-k/) 1499
	// - [1277. 统计全为 1 的正方形子矩阵](https://leetcode.cn/problems/count-square-submatrices-with-all-ones/) 1613  也可以 DP
	// - [1292. 元素和小于等于阈值的正方形的最大边长](https://leetcode.cn/problems/maximum-side-length-of-a-square-with-sum-less-than-or-equal-to-threshold/) 1735
	// - [221. 最大正方形](https://leetcode.cn/problems/maximal-square/) 也可以 DP
	//     - https://codeforces.com/problemset/problem/1301/E 2500
	//     - 思考：如果是菱形怎么 DP
	// - [1504. 统计全 1 子矩形](https://leetcode.cn/problems/count-submatrices-with-all-ones/) 1845
	// - [1074. 元素和为目标值的子矩阵数量](https://leetcode.cn/problems/number-of-submatrices-that-sum-to-target/) 2189
	// 自加写法 https://codeforces.com/contest/835/submission/120031673
	// https://codeforces.com/contest/1107/problem/D
	// https://codeforces.com/problemset/problem/1731/D
	// https://codeforces.com/problemset/problem/611/C
	// 二维带权前缀和 https://codeforces.com/problemset/problem/2044/H
	// 分类讨论 https://www.luogu.com.cn/problem/P3625
	// - https://atcoder.jp/contests/abc347/tasks/abc347_f
	// 转换 https://atcoder.jp/contests/agc015/tasks/agc015_c
	matrixSum2D := func(a [][]int) {
		n, m := len(a), len(a[0])
		// sum 第一行和第一列都是 0
		sum := make([][]int, n+1)
		for i := range sum {
			sum[i] = make([]int, m+1)
			//for j := range sum[i] { sum[i][j] = 1e18 }
		}
		for i, row := range a {
			for j, v := range row {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
			}
		}
		// 左闭右开写法 r1<=r<r2 && c1<=c<c2
		query := func(r1, c1, r2, c2 int) int {
			return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
		}
		// 左闭右闭写法 r1<=r<=r2 && c1<=c<=c2
		query2 := func(r1, c1, r2, c2 int) int {
			return sum[r2+1][c2+1] - sum[r2+1][c1] - sum[r1][c2+1] + sum[r1][c1]
		}

		_ = []any{query, query2}
	}

	/* 菱形（曼哈顿距离）区域和
	原矩阵 n 行 m 列

	原坐标顺时针旋转 45°，再上移 n-1
	(x,y) -> (x+y,y-x+n-1)
	复原成原坐标（下移 n-1，再逆时针旋转 45°）
	设 y' = y - n + 1
	(x,y') -> ((x-y')/2, (x+y')/2)

	映射前（坐标系顺序，不是矩阵顺序）
	n=4, m=3
	(0,2) (1,2) (2,2) (3,2)
	(0,1) (1,1) (2,1) (3,1)
	(0,0) (1,0) (2,0) (3,0)

	映射后
	横纵坐标范围 [0, n+m-2]
	设 y' = y - n + 1
	满足 -x <= y' <= x 且 -(2n-2-x) <= y' <= 2m-2-x 且 (x+y')%2 == 0
	y^
	5|            (0,2)
	4|      (0,1)       (1,2)
	3|(0,0)       (1,1)       (2,2)
	2|      (1,0)       (2,1)       (3,2)   <- (n+m-2, m-1)
	1|            (2,0)       (3,1)
	0|                  (3,0)
	 |------------------------------------> x
	    0     1     2     3     4     5

	若按照 for x for y 的顺序，会从左到右一列一列遍历，每列从下到上

	映射前：以 (x,y) 为中心，曼哈顿距离为 r 的菱形区域
	映射后：左上角在 (x'-r,y'-r)，右下角在 (x'+r,y'+r) 的正方形区域
	注意下标和 0 取 max，和 n+m-2 取 min
	*/
	rhombusSum := func(a [][]int) {
		n, m := len(a), len(a[0])
		size := n + m - 1
		b := make([][]int, size)
		for i := range b {
			b[i] = make([]int, size)
		}
		for i, r := range a {
			for j, v := range r {
				b[i+j][j-i+n-1] = v // 映射
			}
		}

		sum := make([][]int, size+1)
		for i := range sum {
			sum[i] = make([]int, size+1)
		}
		for i, row := range b {
			for j, v := range row {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
			}
		}
		// 返回到 (x,y) 曼哈顿距离不超过 r 的格子元素和
		// 左闭右闭写法 r1<=r<=r2 && c1<=c<=c2
		query := func(x, y, r int) int {
			x, y = x+y, y-x+n-1
			r1, c1 := max(x-r, 0), max(y-r, 0)
			r2, c2 := min(x+r, size-1), min(y+r, size-1)
			return sum[r2+1][c2+1] - sum[r2+1][c1] - sum[r1][c2+1] + sum[r1][c1]
		}

		_ = query
	}

	// 矩阵每行每列的前缀和
	rowColSum := func(a [][]int) (sumR, sumC [][]int) {
		n, m := len(a), len(a[0])
		sumR = make([][]int, n)
		for i, row := range a {
			sumR[i] = make([]int, m+1)
			for j, v := range row {
				sumR[i][j+1] = sumR[i][j] + v
			}
		}
		sumC = make([][]int, n+1)
		for i := range sumC {
			sumC[i] = make([]int, m)
		}
		for j := 0; j < m; j++ {
			for i, row := range a {
				sumC[i+1][j] = sumC[i][j] + row[j]
			}
		}
		// 用法：
		// (i,j) 向右连续 k 个数：sumR[i][j+k] - sumR[i][j]
		// (i,j) 向下连续 k 个数：sumC[i+k][j] - sumC[i][j]
		return
	}

	// 矩阵斜向前缀和 / 菱形边界和  ⃟ 
	// 菱形区域和见下面的 rhombusSum
	// LC1878 https://leetcode.cn/problems/get-biggest-three-rhombus-sums-in-a-grid/ 1898
	diagonalSum := func(a [][]int) {
		n, m := len(a), len(a[0])
		ds := make([][]int, n+1) // 主对角线方向 ↘ 前缀和
		as := make([][]int, n+1) // 反对角线方向 ↙ 前缀和
		for i := range ds {
			ds[i] = make([]int, m+1)
			as[i] = make([]int, m+1)
		}
		for i, r := range a {
			for j, v := range r {
				ds[i+1][j+1] = ds[i][j] + v // ↘
				as[i+1][j] = as[i][j+1] + v // ↙
			}
		}
		// 从 (x,y) 开始，向 ↘ 连续 k 个数的和（需要保证 ↘ 至少有 k 个数）
		queryDiagonal := func(x, y, k int) int { return ds[x+k][y+k] - ds[x][y] }
		// 从 (x,y) 开始，向 ↙ 连续 k 个数的和（需要保证 ↙ 至少有 k 个数）
		queryAntiDiagonal := func(x, y, k int) int { return as[x+k][y+1-k] - as[x][y+1] }

		// 中心在 (x,y)，向外扩展 k 个单位的菱形边界和  ⃟ 
		// 菱形上顶点 (x-k,y)
		// 菱形下顶点 (x+k,y)
		// 菱形左顶点 (x,y-k)
		// 菱形右顶点 (x,y+k)
		// ！必须保证四个顶点都在矩阵内
		// k=0 时返回 a[x][y]
		// k=1 时返回 (x,y) 上下左右四个格子的和，依此类推
		// 相当于菱形边界长度是 k+1
		queryRhombus := func(x, y, k int) int {
			if !(k <= x && x+k < n && k <= y && y+k < m) {
				panic(-1) // 出界
			}
			if k == 0 {
				return a[x][y]
			}
			s1 := queryDiagonal(x-k, y, k)           // 菱形右上斜边 ↘
			s2 := queryDiagonal(x, y-k, k)           // 菱形左下斜边 ↘
			s3 := queryAntiDiagonal(x-k+1, y-1, k-1) // 菱形左上斜边 ↙
			s4 := queryAntiDiagonal(x, y+k, k+1)     // 菱形右下斜边 ↙
			return s1 + s2 + s3 + s4
		}

		_ = queryRhombus
	}

	// ◣ 等腰直角三角形区域和 / ▲ 金字塔区域和 / ⯁ 菱形区域和 / 风车区域和
	// 金字塔 LC2088 https://leetcode.cn/problems/count-fertile-pyramids-in-a-land/ 2105
	// 菱形 https://codeforces.com/problemset/problem/1393/D 2100
	rightTriangleSum := func(a [][]int) {
		n, m := len(a), len(a[0])

		// 矩形二维前缀和
		sumRect := make([][]int, n+1)
		sumRect[0] = make([]int, m+1)
		for i, row := range a {
			sumRect[i+1] = make([]int, m+1)
			for j, v := range row {
				sumRect[i+1][j+1] = sumRect[i+1][j] + sumRect[i][j+1] - sumRect[i][j] + v
			}
		}
		// 任意矩形区域和
		// 左上 (r1,c1)，右下 (r2,c2)
		queryRect := func(r1, c1, r2, c2 int) int {
			return sumRect[r2+1][c2+1] - sumRect[r2+1][c1] - sumRect[r1][c2+1] + sumRect[r1][c1]
		}

		// 按照四象限划分
		// ◢ ◣
		// ◥ ◤

		// 第一类：◣ 右顶点
		// ts1[i+1][j+1] = 上顶点在第一行或第一列，右顶点在 (i,j) 的 ◣ 区域和
		// todo https://www.codechef.com/problems/TRIQUERY
		ts1 := make([][]int, n+1)
		for i := range ts1 {
			ts1[i] = make([]int, m+1)
		}
		for i, row := range a {
			s := 0
			for j, v := range row {
				s += v
				ts1[i+1][j+1] = ts1[i][j] + s
				if j >= i {
					s -= row[j-i]
				}
			}
		}

		// 第二类：◢ 左顶点
		// ts2[i+1][j] = 上顶点在第一行或最后一列，左顶点在 (i,j) 的 ◢ 区域和
		ts2 := make([][]int, n+1)
		for i := range ts2 {
			ts2[i] = make([]int, m+1)
		}
		for i, row := range a {
			s := 0
			for j := m - 1; j >= 0; j-- {
				s += row[j]
				ts2[i+1][j] = ts2[i][j+1] + s
				if i+j < m {
					s -= row[j+i]
				}
			}
		}

		// 第三类：◥ 左顶点
		// ts3[i][j] = 下顶点在最后一行或最后一列，左顶点在 (i,j) 的 ◥ 区域和
		ts3 := make([][]int, n+1)
		for i := range ts3 {
			ts3[i] = make([]int, m+1)
		}
		for i := n - 1; i >= 0; i-- {
			row := a[i]
			s := 0
			for j := m - 1; j >= 0; j-- {
				s += row[j]
				ts3[i][j] = ts3[i+1][j+1] + s
				if m-j >= n-i {
					s -= row[j+(n-1-i)]
				}
			}
		}

		// 第四类：◤ 右顶点
		// ts4[i][j+1] = 下顶点在最后一行或第一列，右顶点在 (i,j) 的 ◤ 区域和
		ts4 := make([][]int, n+1)
		for i := range ts4 {
			ts4[i] = make([]int, m+1)
		}
		for i := n - 1; i >= 0; i-- {
			row := a[i]
			s := 0
			for j, v := range row {
				s += v
				ts4[i][j+1] = ts4[i+1][j] + s
				if j >= n-1-i {
					s -= row[j-(n-1-i)]
				}
			}
		}

		// 任意 ◣ 区域和
		// 直角顶点为 (i,j)，右顶点为 (i,j+k)，上顶点为 (i-k,j)
		queryTri1 := func(i, j, k int) int {
			if !(k <= i && i < n && 0 <= j && j+k < m) {
				panic(-1)
			}
			// (i-k,j+1) 向 ↘ 走到纵坐标 min(j+1+n-1-(i-k),m-1)
			s1 := queryRect(i-k, j, i, min(j+n-i+k, m-1))
			s2 := ts3[i-k][j+1] // (i-k,j+1)
			if j+k+2 >= m {
				return s1 - s2
			}
			s3 := ts3[i+1][j+k+2] // (i+1, j+k+2)
			return s1 - s2 + s3
		}

		// 任意 ◢ 区域和
		// 直角顶点为 (i,j)，左顶点为 (i,j-k)，上顶点为 (i-k,j)
		queryTri2 := func(i, j, k int) int {
			if !(k <= i && i < n && k <= j && j < m) {
				panic(-1)
			}
			// (i-k,j-1) 向 ↙ 走到纵坐标 max(j-1-(n-1-(i-k)),0)
			s1 := queryRect(i-k, max(j-n+i-k, 0), i, j)
			s2 := ts4[i-k][j] // (i-k,j-1)
			if j-k <= 1 {
				return s1 - s2
			}
			s3 := ts4[i+1][j-k-1] // (i+1, j-k-2)
			return s1 - s2 + s3
		}

		// 任意 ◥ 区域和
		// 直角顶点为 (i,j)，左顶点为 (i,j-k)，下顶点为 (i+k,j)
		queryTri3 := func(i, j, k int) int {
			if !(0 <= i && i+k < n && k <= j && j < m) {
				panic(-1)
			}
			// (i+k,j-1) 向 ↖ 走到纵坐标 max(j-1-i-k,0)
			s1 := queryRect(i, max(j-1-i-k, 0), i+k, j)
			s2 := ts1[i+k+1][j] // (i+k, j-1)
			if j-k <= 1 {
				return s1 - s2
			}
			s3 := ts1[i][j-k-1] // (i-1, j-k-2)
			return s1 - s2 + s3
		}

		// 任意 ◤ 区域和
		// 直角顶点为 (i,j)，右顶点为 (i,j+k)，下顶点为 (i+k,j)
		queryTri4 := func(i, j, k int) int {
			if !(0 <= i && i+k < n && 0 <= j && j+k < m) {
				panic(-1)
			}
			// (i+k,j+1) 向 ↗ 走到纵坐标 min(j+1+i+k,m-1)
			s1 := queryRect(i, j, i+k, min(j+1+i+k, m-1))
			s2 := ts2[i+k+1][j+1] // (i+k, j+1)
			if j+k+2 >= m {
				return s1 - s2
			}
			s3 := ts2[i][j+k+2] // (i-1, j+k+2)
			return s1 - s2 + s3
		}

		// ▲ 金字塔区域和
		// 上顶点为 (i,j)，高度为 h，即最下面一层为 i+h-1
		// 1 <= h <= min(n-i, j+1, m-j)
		pyramidSum := func(i, j, h int) int {
			if !(0 <= i && i+h <= n && h-1 <= j && j+h <= m) {
				panic(-1)
			}
			if h == 1 {
				return a[i][j]
			}
			s1 := queryTri2(i+h-1, j, h-1)
			s2 := queryTri1(i+h-1, j+1, h-2)
			return s1 + s2
		}

		// ▼ 倒金字塔区域和
		// 下顶点为 (i,j)，高度为 h，即最上面一层为 i-h+1
		// 1 <= h <= min(i+1, j+1, m-j)
		invPyramidSum := func(i, j, h int) int {
			if !(h-1 <= i && i < n && h-1 <= j && j+h <= m) {
				panic(-1)
			}
			if h == 1 {
				return a[i][j]
			}
			s1 := queryTri3(i-h+1, j, h-1)
			s2 := queryTri4(i-h+1, j+1, h-2)
			return s1 + s2
		}

		// ⯁ 菱形区域和 
		// 菱形中心点 (i,j)
		// 菱形上顶点 (i-k,j)
		// 菱形下顶点 (i+k,j)
		// 菱形左顶点 (i,j-k)
		// 菱形右顶点 (i,j+k)
		// 0 <= k <= min(i, n-i-1, j, m-1-j)
		rhombusSum := func(i, j, k int) int {
			if !(k <= i && i+k < n && k <= j && j+k < m) {
				panic(-1)
			}
			if k == 0 {
				return a[i][j]
			}
			s1 := queryTri1(i, j+1, k-1)
			s2 := queryTri2(i-1, j, k-1)
			s3 := queryTri3(i, j-1, k-1)
			s4 := queryTri4(i+1, j, k-1)
			return a[i][j] + s1 + s2 + s3 + s4
		}

		// todo 允许出界（出界元素为 0）

		_ = []any{pyramidSum, invPyramidSum, rhombusSum}
	}

	// 利用每个数产生的贡献计算 ∑|ai-aj|, i!=j
	// https://codeforces.com/contest/1311/problem/F
	contributionSum := func(a []int) (sum int) {
		slices.Sort(a)
		for i, v := range a {
			sum += v * (2*i + 1 - len(a))
		}
		return
	}

	/* 差分数组
	请看 https://leetcode.cn/circle/discuss/FfMCgb/
	- [1893. 检查是否区域内所有整数都被覆盖](https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered/) 1307（暴力也可）
	- [1094. 拼车](https://leetcode.cn/problems/car-pooling/) 1441
	- [1109. 航班预订统计](https://leetcode.cn/problems/corporate-flight-bookings/) 1570
	- [2406. 将区间分为最少组数](https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/) 1713
	- [2381. 字母移位 II](https://leetcode.cn/problems/shifting-letters-ii/) 1793
	- [995. K 连续位的最小翻转次数](https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips/) 1835
	- [1943. 描述绘画结果](https://leetcode.cn/problems/describe-the-painting/) 1969
	- [2251. 花期内花的数目](https://leetcode.cn/problems/number-of-flowers-in-full-bloom/) 2022
	- [2772. 使数组中的所有元素都等于零](https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero/) 2029
	- [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/) 2236
	- [370. 区间加法](https://leetcode.cn/problems/range-addition/)（会员题）
	- [2237. 计算街道上满足所需亮度的位置数量](https://leetcode.cn/problems/count-positions-on-street-with-required-brightness/)（会员题）
	- [3009. 折线图上的最大交点数量](https://leetcode.cn/problems/maximum-number-of-intersections-on-the-chart/)（会员题）
	https://codeforces.com/problemset/problem/816/B 1400
	https://codeforces.com/problemset/problem/276/C 1500
	https://codeforces.com/problemset/problem/1700/C 1700
	https://codeforces.com/problemset/problem/1955/E 1700
	https://codeforces.com/problemset/problem/2037/F 2100
	https://codeforces.com/problemset/problem/1634/F 2700 差分思想
	https://atcoder.jp/contests/abc274/tasks/abc274_f 浮点数差分（也可以用扫描线）
	https://www.lanqiao.cn/problems/17164/learning/?contest_id=179
	*/

	// 二阶差分（金字塔式更新）
	// https://ac.nowcoder.com/acm/contest/56446/C
	// https://www.luogu.com.cn/problem/U318099?contestId=123900
	// https://codeforces.com/problemset/problem/1661/D 1900 简化
	// https://codeforces.com/problemset/problem/1710/B 2100
	// todo LC2735 https://leetcode.cn/problems/collecting-chocolates/
	diffOfDiff := func(n int) {
		diff := make([]int, n+1)
		diff2 := make([]int, n+1)

		// 下标从 0 开始
		// a[l]+=base
		// a[l+1]+=base+step
		// a[l+2]+=base+step*2
		// ...
		// a[r] += base+step*(r-l)
		// 一般题目中的 step 会取 1 或者 -1
		update := func(l, r, base, step int) {
			diff[l] += base
			diff[r+1] -= base + step*(r-l) // 修正

			// 差分数组从 l+1 到 r 都加上了 step
			diff2[l+1] += step
			diff2[r+1] -= step
		}

		// 金字塔式更新
		// 下标从 0 开始
		// a[j] += max(base-|i-j|, 0)
		update2 := func(i, base int) {
			update(max(i-base+1, 0), i, max(base-i, 1), 1)
			if base > 1 && i < n-1 {
				update(i+1, min(i+base-1, n-1), base-1, -1)
			}
		}

		// 更新完后，恢复原数组
		sd2, sd := 0, 0
		ori := make([]int, n)
		for i := range ori {
			sd2 += diff2[i]
			sd += diff[i] + sd2
			ori[i] = sd
		}

		_ = update2
	}

	// 离散差分，传入闭区间列表 a，不要求有序
	// 例如，求被 n 个区间覆盖的整点个数：https://leetcode.cn/problems/count-days-without-meetings/
	// https://codeforces.com/problemset/problem/1420/D 1800
	diffMap := func(a []struct{ l, r int }) {
		diff := map[int]int{}
		for _, p := range a {
			diff[p.l]++
			diff[p.r+1]--
		}
		xs := slices.Sorted(maps.Keys(diff))

		// 左闭右开区间 [posAndCnt[i].x, posAndCnt[i+1].x) 中的值都是 posAndCnt[i].c
		// 特别地：
		// posAndCnt[0].x 是最小的区间左端点
		// posAndCnt[len(posAndCnt)-1].x-1 是最大的区间右端点
		type _pair struct{ x, c int }
		posAndCnt := make([]_pair, len(xs))
		sd := 0
		for _, x := range xs {
			sd += diff[x]
			posAndCnt = append(posAndCnt, _pair{x, sd})
		}
		// 返回 x 被多少个 a 中的区间包含（由于 a 是闭区间，端点也算包含）
		query := func(x int) int {
			i := sort.Search(len(posAndCnt), func(i int) bool { return posAndCnt[i].x > x }) - 1
			if i < 0 {
				return 0
			}
			return posAndCnt[i].c
		}

		{
			// 有多少个整点被至少一个区间覆盖？
			coveredCnt := 0
			for i := 0; i < len(posAndCnt)-1; i++ {
				if posAndCnt[i].c > 0 {
					coveredCnt += posAndCnt[i+1].x - posAndCnt[i].x
				}
			}
		}

		{
			// 如果只对左端点感兴趣，可以改为如下写法
			preCnt := make(map[int]int, len(xs)) // 前缀和
			sd := 0
			for _, x := range xs {
				sd += diff[x]
				preCnt[x] = sd
			}
		}

		_ = query
	}

	// 二维差分
	// 【图解】从一维差分到二维差分 https://leetcode.cn/problems/stamping-the-grid/solution/wu-nao-zuo-fa-er-wei-qian-zhui-he-er-wei-zwiu/
	// 模板题 LC2536 https://leetcode.cn/problems/increment-submatrices-by-one/ 1583
	// LC2132 https://leetcode.cn/problems/stamping-the-grid/ 2364（也可以不用差分）
	// https://www.luogu.com.cn/problem/P3397
	// LCP74 离散化 https://leetcode.cn/problems/xepqZ5/
	// https://codeforces.com/problemset/problem/1985/H2
	diff2D := func(n, m int) {
		diff := make([][]int, n+2)
		for i := range diff {
			diff[i] = make([]int, m+2)
		}
		// 将区域 r1<=r<=r2 && c1<=c<=c2 内的数都加上 val（额外 +1 是为了方便求前缀和）
		update := func(r1, c1, r2, c2, val int) {
			diff[r1+1][c1+1] += val
			diff[r1+1][c2+2] -= val
			diff[r2+2][c1+1] -= val
			diff[r2+2][c2+2] += val
		}
		// 直接在 diff 上还原原始矩阵（计算二维前缀和）
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
				// 此时 diff[i+1][j+1] 就是区域更新后的矩阵的 a[i][j]
			}
		}
		// 切出中间 n*m 的结果矩阵
		diff = diff[1 : n+1]
		for i, row := range diff {
			diff[i] = row[1 : m+1]
		}

		// EXTRA: 计算矩阵 a 的二维差分矩阵 d
		// https://codeforces.com/problemset/problem/1592/F1 2600
		diffA := func(a [][]int) [][]int {
			n, m := len(a), len(a[0])
			d := make([][]int, n)
			for i := range d {
				d[i] = make([]int, m)
			}
			d[0][0] = a[0][0]
			for j := 1; j < m; j++ {
				d[0][j] = a[0][j] - a[0][j-1]
			}
			// NOTE: 根据题目的不变量，可能要考虑 i=n 和 j=m 的差分值
			for i := 1; i < n; i++ {
				d[i][0] = a[i][0] - a[i-1][0]
				for j := 1; j < m; j++ {
					d[i][j] = a[i][j] - a[i][j-1] - a[i-1][j] + a[i-1][j-1]
				}
			}
			return d
		}

		_ = []any{update, diffA}
	}

	// 菱形（曼哈顿距离）差分
	// 原矩阵 n 行 m 列
	// 原理见上面的 rhombusSum
	// https://ac.nowcoder.com/acm/contest/83687/F
	// https://codeforces.com/problemset/problem/1864/D 也有更简单的做法
	diffRhombus := func(n, m, q int) [][]int {
		size := n + m - 1
		diff := make([][]int, size+2)
		for i := range diff {
			diff[i] = make([]int, size+2)
		}
		// 将映射后的区域 r1<=r<=r2 && c1<=c<=c2 内的数都加上 val（额外 +1 是为了方便求前缀和）
		update := func(r1, c1, r2, c2, val int) {
			diff[r1+1][c1+1] += val
			diff[r1+1][c2+2] -= val
			diff[r2+2][c1+1] -= val
			diff[r2+2][c2+2] += val
		}
		for ; q > 0; q-- {
			var x, y, r, val int
			//Fscan(in, &x, &y, &r, &val);x--;y--
			x, y = x+y, y-x+n-1 // 映射
			update(max(x-r, 0), max(y-r, 0), min(x+r, size-1), min(y+r, size-1), val)
		}

		ans := make([][]int, n)
		for i := range ans {
			ans[i] = make([]int, m)
		}
		for i := 0; i < size; i++ {
			for j := 0; j < size; j++ {
				diff[i+1][j+1] += diff[i+1][j] + diff[i][j+1] - diff[i][j]
				x, y := i, j-(n-1)
				if -x <= y && y <= x && -(2*n-2-x) <= y && y <= 2*n-2-x && (x+y)%2 == 0 {
					ans[(x-y)/2][(x+y)/2] = diff[i+1][j+1] // 复原
				}
			}
		}
		return ans
	}

	// 求差集 A-B, B-A 和交集 A∩B
	// EXTRA: 求并集 union: A∪B = A-B+A∩B = merge(differenceA, intersection) 或 merge(differenceB, intersection)
	// EXTRA: 求对称差 symmetric_difference: A▲B = A-B ∪ B-A = merge(differenceA, differenceB)
	// a b 必须是有序的（可以为空）
	// 与图论结合 https://codeforces.com/problemset/problem/243/B
	splitDifferenceAndIntersection := func(a, b []int) (differenceA, differenceB, intersection []int) {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				differenceB = append(differenceB, b[j:]...)
				return
			}
			if j == m {
				differenceA = append(differenceA, a[i:]...)
				return
			}
			x, y := a[i], b[j]
			if x < y { // 改成 > 为降序
				differenceA = append(differenceA, x)
				i++
			} else if x > y { // 改成 < 为降序
				differenceB = append(differenceB, y)
				j++
			} else {
				intersection = append(intersection, x)
				i++
				j++
			}
		}
	}

	// 求交集简洁写法
	intersection := func(a, b []int) []int {
		mp := map[int]bool{}
		for _, v := range a {
			mp[v] = true
		}
		mp2 := map[int]bool{}
		for _, v := range b {
			if mp[v] {
				mp2[v] = true
			}
		}
		mp = mp2

		keys := make([]int, 0, len(mp))
		for k := range mp {
			keys = append(keys, k)
		}
		slices.Sort(keys)
		return keys
	}

	// a 是否为 b 的子集（相当于 differenceA 为空）
	// a b 需要是有序的
	isSubset := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				return true
			}
			if j == m {
				return false
			}
			x, y := a[i], b[j]
			if x < y { // 改成 > 为降序
				return false
			} else if x > y { // 改成 < 为降序
				j++
			} else {
				i++
				j++
			}
		}
	}

	// 是否为不相交集合（相当于 intersection 为空）
	// a b 需要是有序的
	isDisjoint := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n || j == m {
				return true
			}
			x, y := a[i], b[j]
			if x < y { // 改成 > 为降序
				i++
			} else if x > y { // 改成 < 为降序
				j++
			} else {
				return false
			}
		}
	}

	xorSet := func(x, y map[int]bool) map[int]bool { // xorMap
		res := make(map[int]bool, len(x)+len(y))
		for v := range x {
			res[v] = true
		}
		for v := range y {
			if res[v] {
				delete(res, v)
			} else {
				res[v] = true
			}
		}
		return res
	}

	//

	// 离散化 · 其一（排序+去重+二分查找）
	// 如果不要求下标连续，也可以不去重
	// https://codeforces.com/problemset/problem/1045/G
	discrete := func(a []int, startIndex int) []int {
		b := slices.Clone(a)
		slices.Sort(b)
		b = slices.Compact(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + startIndex
		}
		return a
	}

	// 离散化 · 其二（不用二分）
	// 返回离散化后的序列（名次）
	// discrete2([]int{100,20,50,50}, 1) => []int{3,1,2,2}
	// 对比，相差 ~10%（Go 1.14.1）
	// discrete   333ms/11748KB https://atcoder.jp/contests/abc221/submissions/35791225
	// discrete2  296ms/14952KB https://atcoder.jp/contests/abc221/submissions/35791381
	// 有些题目需要把 0 加进去离散化，请特别注意 https://atcoder.jp/contests/jsc2021/tasks/jsc2021_f
	// LC1331 https://leetcode.cn/problems/rank-transform-of-an-array/
	discrete2 := func(a []int, startIndex int) (kth []int) {
		type vi struct{ v, i int }
		ps := make([]vi, len(a))
		for i, v := range a {
			ps[i] = vi{v, i}
		}
		slices.SortFunc(ps, func(a, b vi) int { return a.v - b.v }) // or SortStableFunc
		kth = make([]int, len(a))

		// a 有重复元素
		k := startIndex
		for i, p := range ps {
			if i > 0 && p.v != ps[i-1].v {
				k++
			}
			kth[p.i] = k
		}

		// 若需要用 kth 值访问原始值，可以将 ps 去重后求 kth

		// a 无重复元素，或者给相同元素也加上顺序（例如某些求 kth 的题目）
		for i, p := range ps {
			kth[p.i] = i + startIndex
		}

		return
	}

	// 离散化，返回一个名次 map
	// discreteMap([]int{100,20,20,50}, 1) => map[int]int{20:1, 50:2, 100:3}
	// 例题：LC327 https://leetcode.cn/problems/count-of-range-sum/
	discreteMap := func(a []int, startIndex int) (kth map[int]int) {
		sorted := slices.Clone(a)
		slices.Sort(sorted)

		// 有重复元素
		kth = map[int]int{}
		curIdx := startIndex
		for i, v := range sorted {
			if i == 0 || v != sorted[i-1] {
				kth[v] = curIdx
				curIdx++
			}
		}

		// 无重复元素
		kth = make(map[int]int, len(sorted))
		for i, v := range sorted {
			kth[v] = i + startIndex
		}

		// EXTRA: 第 k 小元素在原数组中的下标 kthPos
		pos := make(map[int][]int, curIdx-startIndex)
		for i, v := range a {
			pos[v] = append(pos[v], i)
		}
		kthPos := make([][]int, curIdx+1)
		for v, k := range kth {
			kthPos[k] = pos[v]
		}

		return
	}

	// 哈希编号，也可以理解成另一种离散化（无序）
	// 编号从 0 开始
	indexMap := func(a []string) map[string]int {
		mp := map[string]int{}
		for _, v := range a {
			if _, ok := mp[v]; !ok {
				mp[v] = len(mp)
			}
		}
		return mp
	}

	// a 相对于 [0,n) 的补集
	// a 必须是升序且无重复元素
	complement := func(n int, a []int) (res []int) {
		j := 0
		for i := 0; i < n; i++ {
			if j == len(a) || i < a[j] {
				res = append(res, i)
			} else {
				j++
			}
		}
		return
	}

	// 数组第 k 小 (Quick Select)       kthElement nthElement
	// 0 <= k < len(a)
	// 调用会改变数组中元素顺序
	// 代码实现参考算法第四版 p.221
	// 算法的平均比较次数为 ~2n+2kln(n/k)+2(n-k)ln(n/(n-k))
	// https://en.wikipedia.org/wiki/Quickselect
	// https://www.geeksforgeeks.org/quickselect-algorithm/
	// 模板题 LC215 https://leetcode.cn/problems/kth-largest-element-in-an-array/
	//       LC973 https://leetcode.cn/problems/k-closest-points-to-origin/submissions/
	// 模板题 https://codeforces.com/contest/977/problem/C
	quickSelect := func(a []int, k int) int {
		//k = len(a) - 1 - k // 求第 k 大
		rand.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
		for l, r := 0, len(a)-1; l < r; {
			v := a[l] // 切分元素
			i, j := l, r+1
			for {
				for i++; i < r && a[i] < v; i++ { // less(i, l)
				}
				for j--; j > l && a[j] > v; j-- { // less(l, j)
				}
				if i >= j {
					break
				}
				a[i], a[j] = a[j], a[i]
			}
			a[l], a[j] = a[j], v
			if j == k {
				break
			} else if j < k {
				l = j + 1
			} else {
				r = j - 1
			}
		}
		return a[k] //  a[:k+1]  a[k:]
	}

	// 扫描线 Events Sorting + Sweep Line
	// 常与树状数组、线段树、平衡树等数据结构结合
	// https://en.wikipedia.org/wiki/Sweep_line_algorithm
	// https://cses.fi/book/book.pdf 30.1
	// TODO 窗口的星星 https://www.luogu.com.cn/problem/P1502
	// TODO 矩形周长 https://www.luogu.com.cn/problem/P1856
	// 天际线问题 LC218 https://leetcode.cn/problems/the-skyline-problem/
	// TODO 矩形面积并 LC850 https://leetcode.cn/problems/rectangle-area-ii/ 《算法与实现》5.4.3
	//  矩形周长并 http://poj.org/problem?id=1177
	// 经典题 https://codeforces.com/problemset/problem/1000/C
	// https://codeforces.com/problemset/problem/1379/D
	// 转换求解目标 https://codeforces.com/problemset/problem/1285/E
	// 线段相交统计（栈）https://codeforces.com/problemset/problem/1278/D
	// 统计水平方向的线段与垂直方向的线段的交点个数 https://codeforces.com/problemset/problem/610/D
	// 力扣套题 https://leetcode.cn/tag/line-sweep/
	// http://poj.org/problem?id=2932
	// 转换 https://atcoder.jp/contests/arc068/tasks/arc068_c
	sweepLine := func(ranges [][]int) {
		n := len(ranges)
		type event struct{ pos, delta int }
		events := make([]event, 0, 2*n)
		for _, p := range ranges {
			l, r := p[0], p[1]
			events = append(events, event{l, 1}, event{r, -1})
		}
		sort.Slice(events, func(i, j int) bool {
			a, b := events[i], events[j]
			return a.pos < b.pos || a.pos == b.pos && a.delta < b.delta // 先出后进。改成 a.delta > b.delta 为先进后出
		})

		for _, e := range events {
			if e.delta > 0 {

			} else {

			}
		}
	}

	// 扫描线另一种写法，把 delta 压缩进 pos
	// 这样可以避免写一个复杂的排序
	sweepLine2 := func(ranges [][]int) {
		n := len(ranges)
		events := make([]int, 0, 2*n)
		for _, p := range ranges {
			l, r := p[0], p[1]
			events = append(events, l<<1|1, r<<1) // 先出后进
			//events = append(events, l<<1, r<<1|1) // 先进后出
		}
		slices.Sort(events)

		for _, e := range events {
			pos, delta := e>>1, e&1
			_ = pos
			if delta > 0 { // 根据上面的写法来定义何为出何为进

			} else {

			}
		}
	}

	// 扫描线：一维格点刷漆，返回被刷到的格点数
	countCoveredPoints := func(ranges [][]int) int {
		type pair struct{ p, d int }
		m := len(ranges)
		es := make([]pair, 0, 2*m)
		for _, p := range ranges {
			l, r := p[0], p[1]
			es = append(es, pair{l, 1}, pair{r, -1})
		}
		// assert len(es) > 0
		sort.Slice(es, func(i, j int) bool { return es[i].p < es[j].p })
		ans := es[len(es)-1].p - es[0].p + 1
		// 减去没被刷到的格点
		eventCnt, st := 0, es[0].p
		for _, e := range es {
			if eventCnt == 0 {
				if d := e.p - st - 1; d > 0 {
					ans -= d
				}
			}
			eventCnt += e.d
			if eventCnt == 0 {
				st = e.p
			}
		}
		return ans
	}

	// 二维离散化
	// 代码来源 https://atcoder.jp/contests/abc168/tasks/abc168_f
	discrete2D := func(n, m int) (ans int) {
		type line struct{ a, b, c int }
		lr := make([]line, n)
		du := make([]line, m)
		// read ...

		xs := []int{-2e9, 0, 2e9}
		ys := []int{-2e9, 0, 2e9}
		for _, l := range lr {
			a, b, c := l.a, l.b, l.c
			xs = append(xs, a, b)
			ys = append(ys, c)
		}
		for _, l := range du {
			a, b, c := l.a, l.b, l.c
			xs = append(xs, a)
			ys = append(ys, b, c)
		}
		slices.Sort(xs)
		xs = slices.Compact(xs)
		xi := discreteMap(xs, 0) // todo
		slices.Sort(ys)
		ys = slices.Compact(ys)
		yi := discrete(ys, 0) // todo

		lx, ly := len(xi), len(yi)
		glr := make([][]int, lx)
		gdu := make([][]int, lx)
		vis := make([][]bool, lx)
		for i := range glr {
			glr[i] = make([]int, ly)
			gdu[i] = make([]int, ly)
			vis[i] = make([]bool, ly)
		}
		for _, p := range lr {
			glr[xi[p.a]][yi[p.c]]++
			glr[xi[p.b]][yi[p.c]]--
		}
		for _, p := range du {
			gdu[xi[p.a]][yi[p.b]]++
			gdu[xi[p.a]][yi[p.c]]--
		}
		for i := 1; i < lx-1; i++ {
			for j := 1; j < ly-1; j++ {
				glr[i][j] += glr[i-1][j]
				gdu[i][j] += gdu[i][j-1]
			}
		}

		type pair struct{ x, y int }
		q := []pair{{xi[0], yi[0]}}
		for len(q) > 0 {
			p := q[0]
			q = q[1:]
			x, y := p.x, p.y
			if x == 0 || x == lx-1 || y == 0 || y == ly-1 {
				return -1 // 无穷大
			}
			if !vis[x][y] {
				vis[x][y] = true
				ans += (xs[x+1] - xs[x]) * (ys[y+1] - ys[y])
				if glr[x][y] == 0 {
					q = append(q, pair{x, y - 1})
				}
				if glr[x][y+1] == 0 {
					q = append(q, pair{x, y + 1})
				}
				if gdu[x][y] == 0 {
					q = append(q, pair{x - 1, y})
				}
				if gdu[x+1][y] == 0 {
					q = append(q, pair{x + 1, y})
				}
			}
		}
		return
	}

	// https://leetcode.cn/problems/maximum-value-sum-by-placing-three-rooks-ii/
	// 矩阵题目，维护前 i 排的最大、次大、第三大的格子及其列号，并保证三个列号互不相同
	maintainTop3 := func(board [][]int) {
		type pair struct{ x, j int }
		p := [3]pair{} // 最大、次大、第三大
		for i := range p {
			p[i].x = math.MinInt / 3
			p[i].j = -i - 1
		}
		update := func(row []int) {
			for j, x := range row {
				if x > p[0].x {
					if p[0].j != j {
						if p[1].j != j {
							p[2] = p[1]
						}
						p[1] = p[0]
					}
					p[0] = pair{x, j}
				} else if x > p[1].x && j != p[0].j {
					if p[1].j != j {
						p[2] = p[1]
					}
					p[1] = pair{x, j}
				} else if x > p[2].x && j != p[0].j && j != p[1].j {
					p[2] = pair{x, j}
				}
			}
		}

		// 另一种写法，支持维护 top N
		update = func(row []int) {
			for j, x := range row {
			outer:
				for k := range p {
					if x <= p[k].x {
						continue
					}
					for k2 := range p[:k] {
						if j == p[k2].j {
							continue outer
						}
					}
					p[k].x, x = x, p[k].x
					p[k].j, j = j, p[k].j
				}
			}
		}

		_ = update
	}

	//

	// 维护滑动窗口的 max(cnt.values())
	// 若用有序集合维护，需要 O(nlogn)，而下面的写法只需要 O(n)
	// 这里是定长滑窗，也支持不定长滑窗
	// https://codeforces.com/problemset/problem/2009/G2 2200
	slidingWindowMaxFreq := func(a []int, windowSize int) []int {
		n := len(a)
		ans := make([]int, n-windowSize+1)
		cnt := map[int]int{}
		cc := make([]int, n+1)
		maxC := 0 // max(cnt.values())
		for r, v := range a {
			cc[cnt[v]]--
			cnt[v]++
			cc[cnt[v]]++
			maxC = max(maxC, cnt[v])

			l := r + 1 - windowSize
			if l < 0 {
				continue
			}
			ans[l] = maxC // 记录每个窗口的 max(cnt.values())

			v = a[l]
			cc[cnt[v]]--
			cnt[v]--
			cc[cnt[v]]++
			if cc[maxC] == 0 { // 注意不用写 for，因为 max(cnt.values()) 只会一点一点变化
				maxC--
			}
		}
		return ans
	}

	// 滑动窗口还可以维护没有逆运算的运算（但是有单调性），例如 OR AND GCD LCM
	// 时间复杂度 O(n)，考虑每个元素入栈出栈各至多一次
	// 以 LC3171 为例 https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/
	// 如果有二分的需求，可以对从 bottom+1 到 right 的这段单独开个栈（右栈）
	slidingWindowWithStack := func(a []int, k int) int {
		ans := math.MaxInt
		var left, bottom, rightOr int
		for right, x := range a {
			rightOr |= x
			for left <= right && a[left]|rightOr > k {
				ans = min(ans, (a[left]|rightOr)-k)
				if bottom <= left {
					// 重新构建一个栈
					// 由于 left 即将移出窗口，只需计算到 left+1
					for i := right - 1; i > left; i-- {
						a[i] |= a[i+1]
					}
					bottom = right
					rightOr = 0
				}
				left++
			}
			if left <= right {
				ans = min(ans, k-(a[left]|rightOr))
			}
		}
		return ans
	}

	// 注：下面这个代码不是 O(n)，而是 O(nlogU) 
	// - https://leetcode.cn/problems/find-subarray-with-bitwise-or-closest-to-k/submissions/571152848/
	// - 构造方法：左边一串 2^i(i=2~m-1), 中间很多 2^m，右边一串 2^i(i=1~m-1)    这里「一串」指每个 2^i 出现一次

	_ = []any{
		alphabet,
		pow10, dir4, dir4R, dir8, perm3, perm4,
		abs, ceil,
		rotateCopy, rotate, transposeCopy,
		initPalindromeNumber, // 回文数

		subsetSum, recoverArrayFromSubsetSum, subSumSorted,
		prefixSum, groupPrefixSum, circularRangeSum, weightedPrefixSum, bitSum,
		matrixSum2D, rhombusSum, rowColSum, diagonalSum, rightTriangleSum, contributionSum,

		diffOfDiff, diffMap, diff2D, diffRhombus,

		merge, merge2, mergeWithLimit, splitDifferenceAndIntersection, intersection, isSubset, isDisjoint, xorSet,

		discrete, discrete2, discreteMap, indexMap,
		complement, quickSelect, sweepLine, sweepLine2, countCoveredPoints,
		discrete2D,
		maintainTop3,

		slidingWindowMaxFreq, slidingWindowWithStack,
	}
}
