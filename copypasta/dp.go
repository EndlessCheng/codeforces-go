package copypasta

import (
	"container/heap"
	"math"
	"math/bits"
	"slices"
	"sort"
	"strconv"
)

/* 动态规划（Dynamic Programming，DP）

题单：入门/背包/状态机/划分/区间/状压/数位/数据结构优化/树形/博弈/概率期望
https://leetcode.cn/circle/discuss/tXLS3i/

入门视频
https://www.bilibili.com/video/BV1Xj411K7oF/

如何用记忆化搜索打印具体方案
https://leetcode.cn/problems/shortest-common-supersequence/solution/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-auy8z/

① 前缀/后缀之间的转移，例如从 f[i-1] 转移到 f[i]，或者从 f[j] 转移到 f[i]
LC70 爬楼梯 https://leetcode.cn/problems/climbing-stairs/
- LC509 https://leetcode.cn/problems/fibonacci-number/
- LC1137 https://leetcode.cn/problems/n-th-tribonacci-number/ 1143
- 变形：有障碍物 https://atcoder.jp/contests/abc129/tasks/abc129_c
- 变形：有花费 LC746 https://leetcode.cn/problems/min-cost-climbing-stairs/
- LC2466 https://leetcode.cn/problems/count-ways-to-build-good-strings/ 1694
- LC2533 https://leetcode.cn/problems/number-of-good-binary-strings/
- LC377 https://leetcode.cn/problems/combination-sum-iv/ 每次可以往上爬 nums[i] 步
LC198 打家劫舍 https://leetcode.cn/problems/house-robber/
- LC740 https://leetcode.cn/problems/delete-and-earn/
- 变形：恰好选 floor(n/2) 个 https://atcoder.jp/contests/abc162/tasks/abc162_f
- 变形：矩阵打家劫舍 https://codeforces.com/problemset/problem/1195/C
LC213 环形打家劫舍 https://leetcode.cn/problems/house-robber-ii/
- 相似题目 https://atcoder.jp/contests/abc251/tasks/abc251_e
LC276 https://leetcode.cn/problems/paint-fence/
LC368 https://leetcode.cn/problems/largest-divisible-subset/
LC2369 https://leetcode.cn/problems/check-if-there-is-a-valid-partition-for-the-array/ 1780
- 变形：改成环形数组要怎么做
- 相似题目 https://codeforces.com/problemset/problem/1624/E 2000
LC983 https://leetcode.cn/problems/minimum-cost-for-tickets/ 1786
LC1416 https://leetcode.cn/problems/restore-the-array/ 1920
LC2944 https://leetcode.cn/problems/minimum-number-of-coins-for-fruits/
LC2297 https://leetcode.cn/problems/jump-game-viii/
LCR165 https://leetcode.cn/problems/ba-shu-zi-fan-yi-cheng-zi-fu-chuan-lcof/
https://codeforces.com/problemset/problem/1547/E 1500
https://codeforces.com/problemset/problem/1881/E 1500
https://codeforces.com/problemset/problem/1875/D 1600
https://codeforces.com/problemset/problem/1627/E 2200 刷表法 双指针
另见「最长递增子序列」

② 双序列问题，一般定义 f[i][j] 表示对子问题 (s1[:i],s2[:j]) 的求解结果
见下面的「最长公共子序列」，包含大量扩展题目
https://codeforces.com/problemset/problem/2050/E 1500

③ 划分型 DP：将数组分成恰好（或至多）k 个连续子数组，求解与这些子数组有关的最优值
一般定义 f[i][j] 表示将 a[:j] 分成 i 个连续子数组所得到的最优解
此时可以枚举最后一个子数组的左端点 L，从 f[i-1][L] 转移到 f[i][j]，并考虑 a[L:j] 对最优解的影响
- [410. 分割数组的最大值](https://leetcode.cn/problems/split-array-largest-sum/) 做法不止一种
- [813. 最大平均值和的分组](https://leetcode.cn/problems/largest-sum-of-averages/) 1937
- [1278. 分割回文串 III](https://leetcode.cn/problems/palindrome-partitioning-iii/) 1979
- [1335. 工作计划的最低难度](https://leetcode.cn/problems/minimum-difficulty-of-a-job-schedule/) 2035
- [2478. 完美分割的方案数](https://leetcode.cn/problems/number-of-beautiful-partitions/) 2344
- [3077. K 个不相交子数组的最大能量值](https://leetcode.cn/problems/maximum-strength-of-k-disjoint-subarrays/) 2557
- [2911. 得到 K 个半回文串的最少修改次数](https://leetcode.cn/problems/minimum-changes-to-make-k-semi-palindromes/) 2608
https://www.luogu.com.cn/problem/P2679 2023.11.30 茶
https://codeforces.com/problemset/problem/1969/C 1700
https://codeforces.com/problemset/problem/797/F 2600

④ 划分型 DP：最小化/最大化分割出的子数组个数等
- [132. 分割回文串 II](https://leetcode.cn/problems/palindrome-partitioning-ii/)
    至多 k 个 https://codeforces.com/problemset/problem/137/D
- [2707. 字符串中的额外字符](https://leetcode.cn/problems/extra-characters-in-a-string/) 1736
- [2767. 将字符串分割为最少的美丽子字符串](https://leetcode.cn/problems/partition-string-into-minimum-beautiful-substrings/) 1865
- [1105. 填充书架](https://leetcode.cn/problems/filling-bookcase-shelves/) 2014
- [2547. 拆分数组的最小代价](https://leetcode.cn/problems/minimum-cost-to-split-an-array/) 2020
- [2463. 最小移动总距离](https://leetcode.cn/problems/minimum-total-distance-traveled/) 2454
- [2977. 转换字符串的最小成本 II](https://leetcode.cn/problems/minimum-cost-to-convert-string-ii/) 2696
- [2052. 将句子分隔成行的最低成本](https://leetcode.cn/problems/minimum-cost-to-separate-sentence-into-rows/)（会员题）
https://codeforces.com/problemset/problem/1005/D 1500
https://codeforces.com/problemset/problem/1714/D 1600 允许重叠 输出方案
https://www.luogu.com.cn/problem/P1874

⑤ 多维 / 额外状态
LC1223 https://leetcode.cn/problems/dice-roll-simulation/ 2008
LC2919 https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/ 2031 状态设计的好题
LC2209 https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/ 2106
LC956 https://leetcode.cn/problems/tallest-billboard/ 2381
LC920 https://leetcode.cn/problems/number-of-music-playlists/ 2400
LC1531 看起来是区间 DP，仔细分析后是线性 DP https://leetcode.cn/problems/string-compression-ii/ 2576
LC2464 https://leetcode.cn/problems/minimum-subarrays-in-a-valid-split/ 枚举选哪个
https://codeforces.com/problemset/problem/2027/D1 1700
https://codeforces.com/contest/404/problem/D 1900
https://codeforces.com/problemset/problem/1920/E 2000
https://codeforces.com/problemset/problem/2027/D2 2200 在 DP 数组上滑窗
https://codeforces.com/problemset/problem/2045/H 2200
https://codeforces.com/problemset/problem/6/D 2600

从 X 操作到 Y（部分题目也可以用 BFS）
+1 -1 /2 [397. 整数替换](https://leetcode.cn/problems/integer-replacement/)
-1 +1 /5 /11 [2998. 使 X 和 Y 相等的最少操作次数](https://leetcode.cn/problems/minimum-number-of-operations-to-make-x-and-y-equal/) 1795
+a[i] -a[i] ^a[i] [2059. 转化数字的最小运算数](https://leetcode.cn/problems/minimum-operations-to-convert-number/) 1850
-1 *2 [991. 坏了的计算器](https://leetcode.cn/problems/broken-calculator/) 1909
/2 /3 [1553. 吃掉 N 个橘子的最少天数](https://leetcode.cn/problems/minimum-number-of-days-to-eat-n-oranges/) 2048
-1 /k https://codeforces.com/problemset/problem/940/B 1400
[LCP 09. 最小跳跃次数](https://leetcode.cn/problems/zui-xiao-tiao-yue-ci-shu/)
[LCP 20. 快速公交](https://leetcode.cn/problems/meChtZ/)
*5 /6 https://ac.nowcoder.com/acm/contest/71512/D

预处理
LC2638 https://leetcode.cn/problems/count-the-number-of-k-free-subsets/

todo 题单 https://www.luogu.com.cn/training/83815#problems
跳台阶+禁入点 https://atcoder.jp/contests/abc289/tasks/abc289_d
入门计数 DP https://atcoder.jp/contests/abc248/tasks/abc248_c
https://atcoder.jp/contests/abc281/tasks/abc281_d
选或不选 [1800·hot10] https://codeforces.com/contest/1525/problem/D
https://codeforces.com/problemset/problem/176/B
https://codeforces.com/problemset/problem/1324/E
https://codeforces.com/problemset/problem/505/C
https://atcoder.jp/contests/abc267/tasks/abc267_d
由 n 个值互不相同的点组成的高度不小于 h 的 BST 有多少个 https://codeforces.com/problemset/problem/9/D
https://codeforces.com/problemset/problem/38/E
好题：涉及到相邻状态先后关系的 DP（喂兔子） https://codeforces.com/problemset/problem/358/D
https://codeforces.com/problemset/problem/446/A
https://codeforces.com/problemset/problem/603/A
处理区间元素不能在区间外面的技巧 https://codeforces.com/problemset/problem/811/C https://codeforces.com/contest/811/submission/174568255
https://codeforces.com/problemset/problem/1120/C
与 KMP 结合 https://codeforces.com/problemset/problem/1163/D
https://codeforces.com/problemset/problem/1168/C
https://codeforces.com/problemset/problem/1542/D
https://codeforces.com/problemset/problem/1845/E
LC2143 https://leetcode.cn/problems/choose-numbers-from-two-arrays-in-range/

不相交区间 DP
- [2830. 销售利润最大化](https://leetcode.cn/problems/maximize-the-profit-as-the-salesman/) 1851
- [2008. 出租车的最大盈利](https://leetcode.cn/problems/maximum-earnings-from-taxi/) 1872
- [1235. 规划兼职工作](https://leetcode.cn/problems/maximum-profit-in-job-scheduling/) 2023
- [1751. 最多可以参加的会议数目 II](https://leetcode.cn/problems/maximum-number-of-events-that-can-be-attended-ii/) 2041
https://codeforces.com/problemset/problem/1801/C
LC2054 贪心 https://leetcode.cn/problems/two-best-non-overlapping-events/

排列型/插入型
LC629 https://leetcode.cn/problems/k-inverse-pairs-array/ https://www.luogu.com.cn/problem/P2513
https://www.lanqiao.cn/problems/240/learning/
https://atcoder.jp/contests/abc282/tasks/abc282_g

介值定理
https://codeforces.com/contest/1695/problem/C
https://codeforces.com/contest/2043/problem/C

网格路径问题 网格图 DP
#### 练习 1
- [62. 不同路径](https://leetcode.cn/problems/unique-paths/)
- [63. 不同路径 II](https://leetcode.cn/problems/unique-paths-ii/)
- [64. 最小路径和](https://leetcode.cn/problems/minimum-path-sum/)
	- 变形：连续性 & 上下界思想 https://codeforces.com/contest/1695/problem/C
    - https://atcoder.jp/contests/arc137/tasks/arc137_b 也用到了这个思想
- [120. 三角形最小路径和](https://leetcode.cn/problems/triangle/)
	- https://www.luogu.com.cn/problem/P1216
    - 额外状态 https://www.luogu.com.cn/problem/P1544
- [2684. 矩阵中移动的最大次数](https://leetcode.cn/problems/maximum-number-of-moves-in-a-grid/) 1626
- [1301. 最大得分的路径数目](https://leetcode.cn/problems/number-of-paths-with-max-score/) 1853
#### 练习 2
- [329. 矩阵中的最长递增路径](https://leetcode.cn/problems/longest-increasing-path-in-a-matrix/)
- [2328. 网格图中递增路径的数目](https://leetcode.cn/problems/number-of-increasing-paths-in-a-grid/) 2001
#### 练习 3
- [1289. 下降路径最小和 II](https://leetcode.cn/problems/minimum-falling-path-sum-ii/) 1697
- [2435. 矩阵中和能被 K 整除的路径](https://leetcode.cn/problems/paths-in-matrix-whose-sum-is-divisible-by-k/) 1952
- [741. 摘樱桃](https://leetcode.cn/problems/cherry-pickup/)
   - https://codeforces.com/problemset/problem/213/C 2000
   - https://www.luogu.com.cn/problem/P1004 [NOIP2000 提高组] 方格取数
   - https://atcoder.jp/contests/abc320/tasks/abc320_f 2042
- [1463. 摘樱桃 II](https://leetcode.cn/problems/cherry-pickup-ii/) 1957
   - 回文串 https://codeforces.com/problemset/problem/570/E
每行至多选三个 https://atcoder.jp/contests/abc175/tasks/abc175_e
https://atcoder.jp/contests/abc147/tasks/abc147_e 1713

图 DP
https://codeforces.com/problemset/problem/346/D 2600

思考过程：
1. 把原问题重新复述一遍，例如「从前 n 个数中选择若干个数，这些数的和为 m 的方案数」。
2. 根据题意，尝试「缩小」问题的规模，我们可以怎样缩小？
   - 这里有两个变量 n 和 m，有什么方法可以把它们缩小？
3. 尝试「原子」操作（考虑其中「一个」数选或者不选，例如第 n 个数）：
   - 不选第 n 个数，问题变为「从前 n-1 个数中选择若干个数，这些数的和为 m 的方案数」。
   - 选第 n 个数，问题变为「从前 n-1 个数中选择若干个数，这些数的和为 m-a[n] 的方案数」。
   - 原问题可以不重不漏地分解成这两种情况。
   - 根据加法原理，原问题为这两种方案的和。
4. 这可以用记忆化搜索写。终点是什么？
   - n=0。（这里数组下标是从 1 开始的）
5. 如果用递推来思考，要怎么写？空间能否压缩？
   - 自底向上思考记忆化搜索的过程。
6.（进阶）如果复杂度过高，如何根据状态转移方程来优化？
7.（进阶）状态不好确定时，尝试转化问题模型、逆序思考、增加维度等等。（试试下面的题目）

题目已经分类整理好：试试搜索「最大子段和」等。

常规题目
预处理 https://codeforces.com/contest/1932/problem/F

记忆化搜索
- [1387. 将整数按权重排序](https://leetcode.cn/problems/sort-integers-by-the-power-value/) 1507

如何设计状态
https://codeforces.com/problemset/problem/1461/B 1400
https://codeforces.com/problemset/problem/553/A 1500
https://codeforces.com/problemset/problem/1286/A 1800
https://codeforces.com/problemset/problem/1987/D 1800
https://codeforces.com/problemset/problem/14/E 1900
https://codeforces.com/problemset/problem/452/D 1900 题解 https://www.luogu.com.cn/blog/endlesscheng/solution-cf452d
https://codeforces.com/problemset/problem/687/C 1900
https://codeforces.com/problemset/problem/1012/C 1900
https://codeforces.com/problemset/problem/360/B 2000
https://codeforces.com/problemset/problem/461/B 2000
todo https://codeforces.com/problemset/problem/571/B 2000
https://codeforces.com/problemset/problem/1408/D 2000
https://codeforces.com/problemset/problem/1783/D 2000 推公式
https://codeforces.com/problemset/problem/1025/D 2100
https://codeforces.com/problemset/problem/1027/E 2100
https://codeforces.com/problemset/problem/1579/G 2200
todo https://codeforces.com/problemset/problem/441/E 2400 考虑 x+i 的尾零个数
https://codeforces.com/contest/1927/problem/G
https://atcoder.jp/contests/arc115/tasks/arc115_e 容斥
- https://codeforces.com/contest/1591/problem/F
todo https://codeforces.com/problemset/problem/744/C 2400
https://codeforces.com/problemset/problem/840/C 2500
https://atcoder.jp/contests/abc237/tasks/abc237_f
https://atcoder.jp/contests/abc232/tasks/abc232_e
混合逆序对 https://atcoder.jp/contests/arc097/tasks/arc097_c
寻找子问题 https://atcoder.jp/contests/arc116/tasks/arc116_d
todo https://atcoder.jp/contests/abc200/tasks/abc200_e
SEERC05，紫书例题 9-3，UVa 1347 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=446&page=show_problem&problem=4093
LC2919 https://leetcode.cn/problems/minimum-increment-operations-to-make-array-beautiful/ 2031
LC3193 https://leetcode.cn/problems/count-the-number-of-inversions/ 利用排列的性质
LC956 https://leetcode.cn/problems/tallest-billboard/ 2381
LC1388 https://leetcode.cn/problems/pizza-with-3n-slices/ 2410
LC903 DI 序列的有效排列 https://leetcode.cn/problems/valid-permutations-for-di-sequence/ 2433
LC2638 https://leetcode.cn/problems/count-the-number-of-k-free-subsets/
https://www.luogu.com.cn/problem/P9688?contestId=133572
《挑战》pp.62-64 多重部分和问题

如何消除后效性（通过巧妙地设计状态/发现性质）
LC2896 执行操作使两个字符串相等 https://leetcode.cn/problems/apply-operations-to-make-two-strings-equal/
LC312 戳气球 https://leetcode.cn/problems/burst-balloons/
LC546 消消乐 https://leetcode.cn/problems/remove-boxes/ https://leetcode.com/contest/leetcode-weekly-contest-25
UVa1625 合并序列 最小化所有元素第一次和最后一次出现的位置差的和 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4500
涉及到相邻状态先后关系（喂兔子）https://codeforces.com/problemset/problem/358/D

状态优化
LC935 减少状态个数 https://leetcode.cn/problems/knight-dialer/
妙！https://atcoder.jp/contests/abc243/tasks/abc243_g

输出方案
LC1092 https://leetcode.cn/problems/shortest-common-supersequence/
- 题解：https://leetcode.cn/problems/shortest-common-supersequence/solution/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-auy8z/
LC2212 https://leetcode.cn/problems/maximum-points-in-an-archery-competition/

值域 DP
常见于递增子序列相关的题目
LC3041 https://leetcode.cn/problems/maximize-consecutive-elements-in-an-array-after-modification/
https://codeforces.com/problemset/problem/1582/F1

决策单调性
todo 决策单调性优化总结 https://www.luogu.com.cn/blog/command-block/dp-di-jue-ce-dan-diao-xing-you-hua-zong-jie
https://codeforces.com/problemset/problem/229/D

增量法
见 common.go

思维转换
谁来当 DP 对象 LC1434 https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other/
扔蛋问题 LC887 https://leetcode.cn/problems/super-egg-drop/ https://www.bilibili.com/video/BV1KE41137PK
LC920* https://leetcode.cn/problems/number-of-music-playlists/ 注：官方题解给出了一种生成函数的做法
状态优化 https://codeforces.com/problemset/problem/838/E
「排序」题的转换 https://codeforces.com/problemset/problem/1223/D
https://codeforces.com/problemset/problem/1542/D
https://codeforces.com/problemset/problem/520/E
https://codeforces.com/problemset/problem/883/I
路径计数+推箱子 https://codeforces.com/problemset/problem/1225/E
找关键元素+状态机DP https://codeforces.com/problemset/problem/623/B
https://codeforces.com/problemset/problem/1624/E

贪心+DP
https://leetcode.cn/problems/minimum-time-to-make-array-sum-at-most-x/
https://codeforces.com/problemset/problem/1799/F 2700

NOTE: 无后效性是指当前的决策只与过去的结果有关，而与过去的决策无关
NOTE: 若状态转移不构成 DAG，请尝试建图+BFS，见：
	https://ac.nowcoder.com/acm/contest/6218/B
	https://codeforces.com/problemset/problem/283/B 活用 012 染色
    https://codeforces.com/problemset/problem/1272/F
    - 也可以在记忆化搜索的过程中，提前设置 memo 的值，来避免陷入死循环 https://codeforces.com/problemset/submission/1272/208121980
NOTE: 递归套递归打印方案 https://leetcode.cn/problems/shortest-common-supersequence/solutions/2194615/cong-di-gui-dao-di-tui-jiao-ni-yi-bu-bu-auy8z/
NOTE: 若使用滚动数组，注意复用时可能要初始化
NOTE:（区间 DP）正向计算不易时，试着反向计算
TIPS: 若转移是若干相邻项之和，可以考虑 f(p) - f(p-1) 的值，用滑动窗口来维护区间和，从而优化转移
      例题 LC837 https://leetcode.cn/problems/new-21-game/
递归打印路径：https://codeforces.com/problemset/problem/2/B
需要补充额外的状态 https://codeforces.com/problemset/problem/682/D

todo Non-trivial DP Tricks and Techniques https://codeforces.com/blog/entry/47764

交替 DP
https://codeforces.com/problemset/problem/1479/B2
思路二 https://www.luogu.com.cn/blog/wsyhb/post-ti-xie-cf1479b2-painting-the-array-ii

计数 DP
另见 math_comb.go 中的「一些组合问题」
入门计数 DP https://atcoder.jp/contests/abc248/tasks/abc248_c
入门计数 DP LC1079 https://leetcode.cn/problems/letter-tile-possibilities/
- 相似技巧 LC1681 https://leetcode.cn/problems/minimum-incompatibility/discussion/comments/2051770
https://codeforces.com/problemset/problem/414/B 1400
https://codeforces.com/problemset/problem/1794/D 1900
https://codeforces.com/problemset/problem/1767/C 2100 带约束的计数 DP
https://codeforces.com/problemset/problem/2060/F 2200
https://codeforces.com/problemset/problem/626/F 2400 转换
https://codeforces.com/problemset/problem/1237/F 2600
todo https://atcoder.jp/contests/abc234/tasks/abc234_f

多重组合
- 见「多重背包 - 求方案数 - 同余前缀和优化」

多重排列
- f[i][j] 表示前 i 类数字组成长为 j 的排列个数
- f[i][j] = ∑f[i-1][k]*C(j,k), 0<=k<=min(j,cnt[i])
- 边界 f[0][0] = 1

贪心优化 DP
https://codeforces.com/problemset/problem/864/E

双指针优化 DP
https://codeforces.com/problemset/problem/883/I
https://training.olinfo.it/#/task/preoii_yutaka/statement

我的视频讲解：
https://www.bilibili.com/video/BV1Xj411K7oF 从记忆化搜索到递推
https://www.bilibili.com/video/BV16Y411v7Y6 背包问题
https://www.bilibili.com/video/BV1TM4y1o7ug LCS
https://www.bilibili.com/video/BV1ub411Q7sB LIS
https://www.bilibili.com/video/BV1ho4y1W7QK 状态机 DP
https://www.bilibili.com/video/BV1Gs4y1E7EU 区间 DP
https://www.bilibili.com/video/BV17o4y187h1 树形 DP
https://www.bilibili.com/video/BV1vu4y1f7dn 树形 DP
https://www.bilibili.com/video/BV1oF411U7qL 树形 DP

https://www.bilibili.com/video/BV1gf4y1i78H 动态规划的套路 by wisdompeak
https://www.bilibili.com/video/av70148899 DP 入门，01 背包，完全背包，多重背包
https://www.bilibili.com/video/av77393700 LCS LIS
https://www.bilibili.com/video/av83939419 区间 DP
https://www.bilibili.com/video/av93356551 状态压缩 DP
https://www.bilibili.com/video/av98090640 树形 DP
https://www.bilibili.com/video/av85636122 动态规划 · 零 - Introduction
https://www.bilibili.com/video/av86983419 动态规划 · 一 - 序列型
https://www.bilibili.com/video/av89052674 动态规划 · 二 - 坐标、双序列、划分 & 状态压缩

套题/总结：
推荐 AtCoder 上的经典 DP 场 https://atcoder.jp/contests/dp
    题解 https://www.cnblogs.com/shanxieng/p/10232228.html
        https://codeforces.com/blog/entry/92170
        https://www.hamayanhamayan.com/entry/2019/01/12/163853
    讨论 https://codeforces.com/blog/entry/64250
《挑战程序设计竞赛》上的练习题（均为 POJ）
2.3 节
3176 https://www.luogu.com.cn/problem/P1216 数字三角形
2229 https://www.luogu.com.cn/problem/P6065 将 n 分拆为若干个 2 的次幂的和的方法数 https://oeis.org/A018819
2385 https://www.luogu.com.cn/problem/P2690 f[i分钟][j移动次数] = max(f[i-1][j], f[i-1][j-1]) + 当前分钟是否有苹果落在 j 次移动后的位置   最后答案为 max{f[n-1]}
3616 https://www.luogu.com.cn/problem/P2889 DAG 最长路
3280 https://www.luogu.com.cn/problem/P2890 增删取 min，跑区间 DP
1742 http://acm.hdu.edu.cn/showproblem.php?pid=2844 多重背包
3046 http://poj.org/problem?id=3046 todo
3181 https://www.luogu.com.cn/problem/P6205 完全背包
1065 http://acm.hdu.edu.cn/showproblem.php?pid=1051 n 轮 LIS
1631 http://acm.hdu.edu.cn/showproblem.php?pid=1950 转换成 LIS
3666 https://www.luogu.com.cn/problem/P2893
     https://codeforces.com/problemset/problem/13/C
     https://codeforces.com/problemset/problem/713/C
     https://www.luogu.com.cn/problem/P4597 加强版
2392 https://www.luogu.com.cn/problem/P6771 多重背包，按高度限制排序。高度既是价值也是体积
2184 https://www.luogu.com.cn/problem/P2340 把 IQ 看成体积，EQ 看成价值，注意把负数偏移到非负数，以及负数的转移写法
todo 3.4 节
2686 https://www.luogu.com.cn/problem/SP1700
1769 https://www.luogu.com.cn/problem/SP90 https://www.luogu.com.cn/problem/UVA1322
2441
3254 https://www.luogu.com.cn/problem/P1879
2836
1795 https://www.luogu.com.cn/problem/SP1776
3411 https://www.luogu.com.cn/problem/SP3953
3420
3735
3171 https://www.luogu.com.cn/problem/P4644 见 graph.shortestPathDijkstra
CSES DP section editorial https://codeforces.com/blog/entry/70018
力扣上的 DP 问题
    分类汇总 https://zhuanlan.zhihu.com/p/126546914
    https://leetcode.com/discuss/general-discussion/458695/dynamic-programming-patterns
    https://github.com/CyC2018/CS-Notes/blob/master/notes/Leetcode%20%E9%A2%98%E8%A7%A3%20-%20%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92.md
    https://leetcode.com/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/discuss/108870/Most-consistent-ways-of-dealing-with-the-series-of-stock-problems
    https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/solution/yi-ge-tong-yong-fang-fa-tuan-mie-6-dao-gu-piao-w-5/
    https://leetcode.cn/tag/dynamic-programming/
信息学奥赛一本通 第二部分 基础算法 --> 第九章 动态规划 http://ybt.ssoier.cn:8088/index.php
算法竞赛专题解析（11）：DP概述和常见DP面试题 https://blog.csdn.net/weixin_43914593/article/details/105444090
todo 题目推荐 https://www.luogu.com.cn/blog/wyy2020/dp-qian-tan
  https://www.cnblogs.com/flashhu/p/9480669.html

其他资料：
https://github.com/hzwer/shareOI/tree/master/%E5%8A%A8%E6%80%81%E8%A7%84%E5%88%92
https://oi-wiki.org/dp/
https://cp-algorithms.com/dynamic_programming/divide-and-conquer-dp.html
https://wenku.baidu.com/view/7c9de809581b6bd97f19ea72.html 算法合集之《从《鹰蛋》一题浅析对动态规划算法的优化》
*/
func _(abs func(int) int) {
	// 如果用（多维）数组记忆化
	// 小技巧，可以在保存的时候把计算结果 +1（或者 +1e18），取出来的时候 -1（或者 -1e18）
	// 这样原来的 *p != -1 可以改成 *p > 0，避免初始化 memo 数组时的多重循环

	// 由于数据范围的原因，采用 map 记忆化         dpMap
	// LC1553 https://leetcode.cn/problems/minimum-number-of-days-to-eat-n-oranges/
	// LC2998 https://leetcode.cn/problems/minimum-number-of-operations-to-make-x-and-y-equal/
	// LC638 https://leetcode.cn/problems/shopping-offers/
	// https://codeforces.com/problemset/problem/510/D
	// https://codeforces.com/problemset/problem/1746/D
	// 如何估计时间复杂度 https://atcoder.jp/contests/abc275/tasks/abc275_d
	mapDP := func(n int) (ans int) {
		{
			// 一维（多维见下）
			memo := map[int]int{}
			var f func(int) int
			f = func(x int) (res int) {
				//if x == 0 {
				//	return
				//}
				if v, ok := memo[x]; ok {
					return v
				}
				defer func() { memo[x] = res }()

				return
			}
			ans = f(n)
		}

		{
			// 多维
			type pair struct{ x, y int }
			memo := map[pair]int{}
			var f func(int, int) int
			f = func(x, y int) (res int) {
				//if x == n {
				//  return
				//}
				p := pair{x, y}
				if v, ok := memo[p]; ok {
					return v
				}
				defer func() { memo[p] = res }()

				return
			}
			ans = f(0, 0)
		}

		return
	}

	// 最大子段和 / 最大子数组和
	// 有三种思路
	// 1. Kadane 算法：定义状态 f[i] 表示以 a[i] 结尾的最大子段和，则有状态转移方程 f[i]=max(f[i−1],0)+a[i]，答案为 max(f)
	// 2. 前缀和：遍历 a 的同时维护前缀和的最小值，则遍历到 a[i] 时，当前最大子段和为 sum[i]-min(sum[j]), j<i
	//    - 我的题解：https://leetcode.cn/problems/maximum-subarray/solution/qian-zhui-he-zuo-fa-ben-zhi-shi-mai-mai-abu71/
	// 3. 分治：通常用于一些带修改的题目，需要用到线段树 https://www.luogu.com.cn/problem/P4513 https://codeforces.com/contest/1843/problem/F2
	//
	// LC53 https://leetcode.cn/problems/maximum-subarray/ 
	// https://www.luogu.com.cn/problem/P1115
	// LC2606 https://leetcode.cn/problems/find-the-substring-with-maximum-cost/
	//
	// 二维版本
	// https://leetcode.cn/problems/max-submatrix-lcci/
	// https://www.luogu.com.cn/problem/P1719
	//
	// 《算法导论》练习 4.1-5
	// [题型总结] 关于最大子段和及其变式 https://www.luogu.com.cn/blog/wey-yzyl/zui-tai-zi-duan-hu-ji-ji-bian-shi-di-qi-shi
	// 子段长度有上限的最大子段和：见单调队列，题目为 https://ac.nowcoder.com/acm/contest/1006/D
	// 子段长度有下限的最大子段和：转换为前缀和之差 sum[i]-sum[j]，i-j>=K，维护 mn=min(mn,sum[j])，同时更新 sum[i]-mn 的最大值（另见 sort.go 中的 0-1 分数规划）
	// - https://www.luogu.com.cn/problem/P1404
	// 一道题考察上面两个限制 https://codeforces.com/problemset/problem/1796/D
	// - O(n) 做法 https://codeforces.com/contest/1796/submission/263191415
	// 子段和有上限的最大子段和：转换为前缀和之差 sum[i]-sum[j]<=K，在平衡树上二分 sum[j] 
	// - LC363 https://leetcode.cn/problems/max-sum-of-rectangle-no-larger-than-k/
	// 最大两段子段和：求每个位置上的前缀最大子段和和后缀最大子段和 https://www.luogu.com.cn/problem/P2642
	// - 等价题目：允许翻转一段子区间的最大子段和
	// 最大三段子段和 LC689 https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/ *长度固定为 k
	// 删除至多一个元素后的最大子段和 LC1186 https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
	// - 也可以前后缀分解
	// 最大 m 段子段和 http://acm.hdu.edu.cn/showproblem.php?pid=1024
	// 环状最大子段和：转换为 max(最大子段和, 总和减去最小子段和) LC918 https://leetcode.cn/problems/maximum-sum-circular-subarray/
	// 环状最大两段子段和：思路类似，注意取反后需要传入 a[1:n-1] https://www.luogu.com.cn/problem/P1121 https://ac.nowcoder.com/acm/contest/7738/B
	// 去掉一个最大值的最大子段和（值域比较小）https://codeforces.com/contest/1359/problem/D
	// 变形题：
	// LC2321 https://leetcode.cn/problems/maximum-score-of-spliced-array/
	// LC1749 https://leetcode.cn/problems/maximum-absolute-sum-of-any-subarray/
	// - 另一种做法是计算前缀和的最大值与最小值的差
	// LC1191 重复 k 次 https://leetcode.cn/problems/k-concatenation-maximum-sum/
	// https://codeforces.com/problemset/problem/1285/B 1300
	// https://codeforces.com/problemset/problem/788/A 1600
	// https://codeforces.com/problemset/problem/1373/D 1600
	// https://codeforces.com/problemset/problem/2043/C 1600 介值定理
	// https://codeforces.com/problemset/problem/33/C 1800
	// https://codeforces.com/problemset/problem/1845/D 1800 本质是去掉一个最小的子段
	// https://codeforces.com/problemset/problem/1155/D 1900
	// https://codeforces.com/problemset/problem/1197/D 1900 思路 https://docs.qq.com/sheet/DWGFoRGVZRmxNaXFz 里面搜本题链接
	// https://codeforces.com/problemset/problem/1082/E 2000 需要一些转换技巧
	// https://atcoder.jp/contests/arc137/tasks/arc137_b
	// https://codeforces.com/problemset/problem/75/D 2000 多个小数组合并
	// - 这题做法需要用到上面说到的第二种思路
	// https://codeforces.com/problemset/problem/1370/E 2100 转换
	// https://codeforces.com/problemset/problem/1843/F2 2300 树上的情况 
	// 二维的情况（最大子阵和）可以枚举上下边界，转换成一维   O(n^3)
	maxSubarraySum := func(a []int) int {
		if len(a) == 0 { // 根据题意返回
			return 0
		}
		maxS, sum := math.MinInt, 0
		for _, v := range a {
			sum = max(sum, 0) + v
			maxS = max(maxS, sum)
		}
		if maxS < 0 { // 根据题意返回
			//return 0
		}
		return maxS
	}

	// 除了返回最大子段和外，还返回最大子段和对应的子段 [l,r]
	// https://codeforces.com/contest/1692/problem/H
	maxSubarraySumWithRange := func(a []int) (maxS, l, r int) {
		if len(a) == 0 { // 根据题意返回
			return 0, -1, -1
		}
		maxS = a[0] // 注意 l 和 r 默认为 0，即 a[:1]
		for i, sum, st := 1, a[0], 0; i < len(a); i++ {
			if sum < 0 {
				sum, st = 0, i // 重新开始
			}
			sum += a[i]
			if sum > maxS {
				maxS, l, r = sum, st, i
			}
		}
		if maxS < 0 { // 根据题意返回
			//return 0, -1, -1
		}
		return
	}

	// 维护前缀和的最小值的写法
	// https://codeforces.com/contest/1692/problem/H
	maxSubarraySumWithRange = func(a []int) (maxS, l, r int) {
		if len(a) == 0 { // 根据题意返回
			return 0, -1, -1
		}
		maxS = a[0] // 注意 l 和 r 默认为 0，即 a[:1]
		sum, minS, minI := 0, 0, -1
		for i, v := range a {
			sum += v
			if sum-minS > maxS {
				maxS, l, r = sum-minS, minI+1, i
			}
			if sum < minS {
				minS, minI = sum, i
			}
		}
		if maxS < 0 { // 根据题意返回
			//return 0, -1, -1
		}
		return
	}

	// 最大两段子段和（两段必须间隔至少 gap 个数）
	maxTwoSubarraySum := func(a []int, gap int) int {
		// 注意下界
		n := len(a)
		suf := make([]int, n)
		suf[n-1] = a[n-1]
		curSum := a[n-1]
		for i := n - 2; i >= 0; i-- {
			v := a[i]
			curSum = max(curSum+v, v)
			suf[i] = max(suf[i+1], curSum)
		}
		curSum, pre := a[0], a[0]
		ans := pre + suf[1+gap]
		for i := 1; i < n-1-gap; i++ {
			v := a[i]
			curSum = max(curSum+v, v)
			pre = max(pre, curSum)
			ans = max(ans, pre+suf[i+1+gap])
		}
		return ans
	}

	// 最大子序列交替和（买卖股票）
	// 有两种思路：
	// 1. 动态规划，具体见我的题解 https://leetcode.cn/problems/maximum-alternating-subsequence-sum/solution/dong-tai-gui-hua-by-endlesscheng-d92a/
	// 2. 贪心，由于第一个值需要取正，将开头补上 0，就变成买卖股票问题了，只需关心波峰和波谷的值，即 ∑max(0,a[i+1]-a[i])
	// LC1911 https://leetcode.cn/problems/maximum-alternating-subsequence-sum/
	// 变形 LC1526 https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
	// LC122 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
	// 扩展：O(1) 回答交换其中两个元素后的最大子序列交替和 https://codeforces.com/problemset/problem/1420/C2
	maxAlternatingSumDP := func(a []int) int {
		f := [2]int{0, -1e9}
		for _, v := range a {
			f = [2]int{max(f[0], f[1]-v), max(f[1], f[0]+v)}
		}
		return f[1]
	}

	maxAlternatingSumGreedy := func(a []int) (ans int) {
		a = append([]int{0}, a...)
		for i := 1; i < len(a); i++ {
			ans += max(0, a[i]-a[i-1])
		}
		return
	}

	// 修改序列为非降或非增的最小修改次数
	// - 单次修改可以把某个数 +1 或 -1
	// https://writings.sh/post/slope-trick-mono-sequence
	//
	// 通过一个例子来解释这个基于堆的算法：1 5 10 4 2 2 2 2
	// 假设当前维护的是非降序列，前三个数直接插入，不需要任何修改
	// 插入 4 的时候，可以修改为 1 5 5 5，或 1 5 6 6，或... 1 5 10 10，修改次数均为 6
	// 但我们也可以把修改后的序列视作 1 5 4 4，虽然序列不为非降序列，但修改的次数仍然为 6
	// 接下来插入 2，基于 1 5 5 5 的话，修改后的序列就是 1 5 5 5 5，总的修改次数为 9
	// 但我们也可以把修改后的序列视作 1 2 4 4 2，总的修改次数仍然为 9
	// 接下来插入 2，如果基于 1 5 5 5 5 变成 1 5 5 5 5 5，会得到错误的修改次数 12
	// 但是实际上有更优的修改 1 4 4 4 4 4，总的修改次数为 11
	// 同上，把这个序列视作 1 2 2 4 2 2，总的修改次数仍然为 11
	// ...
	// LC2263 https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing/
	// https://codeforces.com/problemset/problem/13/C 2200
	// https://codeforces.com/problemset/problem/713/C 2300 严格单调递增 https://codeforces.com/blog/entry/47094?#comment-315161
	//     这道题做了一个 a[i]-=i 的操作（i 从 1 开始），把严格单调递增变成了非降的情况，从而可以应用该算法
	//     这一技巧的原理是，对于整数来说，单调递增的最小情况是 y=x+C，减去这一函数，就得到了非降序列的最小情况 y=C
	// https://www.luogu.com.cn/problem/P4597
	// https://www.luogu.com.cn/problem/P2893
	// http://poj.org/problem?id=3666
	slopeTrick := func(a []int) int {
		h := hp{} // 大根堆
		ans := 0
		for _, v := range a {
			h.push(v)
			if d := h.IntSlice[0] - v; d > 0 {
				ans += d
				h.IntSlice[0] = v
				heap.Fix(&h, 0)
			}
		}
		return ans
	}

	// 返回一个矩阵 f，其中 f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
	// https://leetcode.cn/problems/find-the-minimum-area-to-cover-all-ones-ii/
	// 类似题目【分三块】：
	// https://www.luogu.com.cn/problem/P3625 
	// - https://atcoder.jp/contests/abc347/tasks/abc347_f
	// https://atcoder.jp/contests/abc062/tasks/arc074_a
	minimumArea := func(a [][]int) [][]int {
		n, m := len(a), len(a[0])
		// f[i+1][j+1] 表示包含【左上角为 (0,0) 右下角为 (i,j) 的子矩形】中的所有 1 的最小矩形面积
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, m+1)
		}
		type data struct{ top, left, right int }
		border := make([]data, m)
		for j := range border {
			border[j].top = -1 // 无
		}

		for i, row := range a {
			left, right := -1, 0
			for j, x := range row {
				if x > 0 {
					if left < 0 {
						left = j
					}
					right = j
				}
				preB := border[j]
				if left < 0 { // 这一排目前全是 0
					f[i+1][j+1] = f[i][j+1] // 等于上面的结果
				} else if preB.top < 0 { // 这一排有 1，上面全是 0
					f[i+1][j+1] = right - left + 1
					border[j] = data{i, left, right}
				} else { // 这一排有 1，上面也有 1
					l, r := min(preB.left, left), max(preB.right, right)
					f[i+1][j+1] = (r - l + 1) * (i - preB.top + 1)
					border[j] = data{preB.top, l, r}
				}
			}
		}
		return f
	}

	// 最长公共子序列 (LCS)
	// 视频讲解：https://www.bilibili.com/video/BV1TM4y1o7ug/
	// 更快的做法（位运算）见 SPOJ LCS0 https://www.luogu.com.cn/problem/SP12076
	//
	// LC1143 模板题 https://leetcode.cn/problems/longest-common-subsequence/
	// LC72 编辑距离 https://leetcode.cn/problems/edit-distance/    
	// - 热身 LC161 https://leetcode.cn/problems/one-edit-distance/
	// - 进阶：编辑距离 + 可以交换相邻字母 https://codeforces.com/problemset/problem/67/C 2600
	// LC97   https://leetcode.cn/problems/interleaving-string/
	// LC115  https://leetcode.cn/problems/distinct-subsequences/
	// LC583  https://leetcode.cn/problems/delete-operation-for-two-strings/
	// LC712  https://leetcode.cn/problems/minimum-ascii-delete-sum-for-two-strings/
	// LC1035 https://leetcode.cn/problems/uncrossed-lines/ 1806
	// LC1458 https://leetcode.cn/problems/max-dot-product-of-two-subsequences/ 1824
	// LC1092 最短公共超序列 (SCS) https://leetcode.cn/problems/shortest-common-supersequence/ 1977
	// LC1639 https://leetcode.cn/problems/number-of-ways-to-form-a-target-string-given-a-dictionary/ 2082
	// LC1713 https://leetcode.cn/problems/minimum-operations-to-make-a-subsequence/ 2351 若其中一个序列无重复元素，可以转换成 LIS
	// - https://www.luogu.com.cn/problem/P1439 
	// https://codeforces.com/problemset/problem/163/A 1700 其中一个改为子串 
	// https://codeforces.com/problemset/problem/1446/B 1800
	// https://codeforces.com/problemset/problem/463/D 1900 多个排列的 LCS 
	// - https://www.luogu.com.cn/problem/P2364 三个字符串的 LCS + 输出方案 
	// https://codeforces.com/problemset/problem/1114/D 1900 转换【巧妙】
	// https://codeforces.com/problemset/problem/346/B 2000 与 KMP 结合 
	// - follow up: 要求某个子串 sub 一定在 LCS 中
	// https://codeforces.com/problemset/problem/1584/F 2600 多串 + 每种字母至多出现两次 
	// - f[c][mask]，其中 mask 记录字母 c 在每个字符串中的出现位置，0 表示左边那个，1 表示右边那个
	// https://atcoder.jp/contests/abc185/tasks/abc185_e 权值 
	// https://atcoder.jp/contests/abc130/tasks/abc130_e 相同子序列个数
	// 2020 多校第二场 https://acm.hdu.edu.cn/showproblem.php?pid=6774
	lcs := func(s, t []byte) int {
		// f[i][j] = LCS(s[:i], t[:j])
		n, m := len(s), len(t)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, m+1)
		}
		for i, v := range s {
			for j, w := range t {
				if v == w {
					f[i+1][j+1] = f[i][j] + 1
				} else {
					f[i+1][j+1] = max(f[i][j+1], f[i+1][j])
				}
			}
		}

		{
			// EXTRA: 某些 dp 非单调性的题目需要计算全局最值
			allMax := 0
			for _, row := range f {
				for _, v := range row {
					allMax = max(allMax, v)
				}
			}
		}

		return f[n][m]
	}

	lcsPath := func(s, t []byte) []byte {
		n, m := len(s), len(t)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, m+1)
		}
		fa := make([][]int8, n+1)
		for i := range fa {
			fa[i] = make([]int8, m+1)
		}
		for i, v := range s {
			for j, w := range t {
				if v == w {
					f[i+1][j+1] = f[i][j] + 1
					fa[i+1][j+1] = 1
				} else {
					if f[i][j+1] > f[i+1][j] {
						f[i+1][j+1] = f[i][j+1]
						fa[i+1][j+1] = 2
					} else {
						f[i+1][j+1] = f[i+1][j]
						fa[i+1][j+1] = 3
					}
				}
			}
		}
		lcs := make([]byte, 0, f[n][m])
		var makeLCS func(i, j int)
		makeLCS = func(i, j int) {
			if i == 0 || j == 0 {
				return
			}
			if fa[i][j] == 1 {
				makeLCS(i-1, j-1)
				lcs = append(lcs, s[i-1])
			} else if fa[i][j] == 2 {
				makeLCS(i-1, j)
			} else {
				makeLCS(i, j-1)
			}
		}
		makeLCS(n, m)
		return lcs
	}

	// LCS 个数
	// https://www.luogu.com.cn/problem/P2516
	lcsCount := func(s, t []byte) int {
		const mod = 1_000_000_007
		n, m := len(s), len(t)
		lcs := make([][]int, n+1)
		count := make([][]int, n+1)
		for i := 0; i <= n; i++ {
			lcs[i] = make([]int, m+1)
			count[i] = make([]int, m+1)
		}
		for i := 0; i <= n; i++ {
			count[i][0] = 1
		}
		for j := 0; j <= m; j++ {
			count[0][j] = 1
		}
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				if s[i-1] == t[j-1] {
					lcs[i][j] = lcs[i-1][j-1] + 1
					count[i][j] = count[i-1][j-1] // 延续之前的子序列数
				} else {
					lcs[i][j] = max(lcs[i-1][j], lcs[i][j-1])
				}
				if lcs[i][j] == lcs[i-1][j] {
					count[i][j] += count[i-1][j] // 加上不考虑 s2[j-1] 的子序列数
				}
				if lcs[i][j] == lcs[i][j-1] {
					count[i][j] += count[i][j-1] // 加上不考虑 s1[i-1] 的子序列数
				}
				if s[i-1] != t[j-1] && lcs[i][j] == lcs[i-1][j-1] {
					count[i][j] -= count[i-1][j-1] // 避免重复子序列的统计
				}
				count[i][j] %= mod
			}
		}
		return (count[n][m] + mod) % mod
	}

	// 最长回文子序列 (LPS)
	// 见下面的「区间 DP」

	// 最长上升子序列 (LIS)
	// 视频讲解：https://www.bilibili.com/video/BV1ub411Q7sB/
	// 这种写法适用于一些定义比较复杂的变形题
	// O(n^2) - 定义 f[i] 为以 a[i] 为末尾的 LIS 的长度
	//          可以把此问题想象成一个「跳跃游戏」，任选一个初始位置向右跳跃，每次只能跳到比当前位置更高的位置，问最多能跳多少次（最后答案加一）
	//          这样能更容易地看出转移的顺序，然后变成一个 DAG 上求最长路的问题
	// 转换 http://acm.hdu.edu.cn/showproblem.php?pid=1950
	// 变体 https://codeforces.com/problemset/problem/1350/B 1400
	// todo 转换 https://codeforces.com/problemset/problem/1562/E
	//【网络流 24 题】能取出多少个长为 len(LIS) 的不相交子序列 https://loj.ac/p/6005 https://www.luogu.com.cn/problem/P2766
	lisSlow := func(a []int) int {
		n := len(a)
		f := make([]int, n)
		for i, v := range a {
			f[i] = 1
			for j, w := range a[:i] {
				if w < v { // 改成 <= 为非降
					f[i] = max(f[i], f[j]+1)
				}
			}
		}
		return slices.Max(f)
	}

	// 最长上升子序列 (LIS)   最长递增子序列
	// 视频讲解：https://www.bilibili.com/video/BV1ub411Q7sB/
	// 方法一：二分
	// O(nlogn) - 定义 g[i] 为长度为 i+1 的上升子序列的末尾元素的最小值（技巧：交换 O(n^2) 定义中的状态与状态值）
	// 求下降，可以考虑取相反数
	// https://oi-wiki.org/dp/basic/#_12
	// 最小划分数 / 狄尔沃斯定理（Dilworth's theorem）https://en.wikipedia.org/wiki/Dilworth%27s_theorem
	//    偏序集的最少反链划分数等于最长链的长度
	// 随机排列 LIS 的长度期望 https://www.zhihu.com/question/266958886
	// On Range LIS Queries https://codeforces.com/blog/entry/111625 https://codeforces.com/blog/entry/111807 https://arxiv.org/pdf/0707.3619
	//
	// LC300 https://leetcode.cn/problems/longest-increasing-subsequence/
	// LC1964 https://leetcode.cn/problems/find-the-longest-valid-obstacle-course-at-each-position/ 1933
	// 建模 https://codeforces.com/problemset/problem/269/B 1700
	// 经典转换（最多相交问题） https://codeforces.com/problemset/problem/67/D https://atcoder.jp/contests/arc126/tasks/arc126_b
	// 最小划分数（导弹拦截）https://www.luogu.com.cn/problem/P1020
	// - https://atcoder.jp/contests/abc134/tasks/abc134_e
	// - LC3231 https://leetcode.cn/problems/minimum-number-of-increasing-subsequence-to-be-removed/
	// 转化成最小划分数+打印划分方案 https://codeforces.com/problemset/problem/1296/E2
	// 合唱队形 https://www.luogu.com.cn/problem/P1091
	// LC1671 合唱队形（至少有升有降）https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/ 1913
	// 二维 LIS LC354 https://leetcode.cn/problems/russian-doll-envelopes/
	// 二维 LIS + 打印方案 https://codeforces.com/problemset/problem/4/D 1700
	// 重复数组的 LIS https://codeforces.com/problemset/problem/582/B 1900 非降
	// - todo https://codeforces.com/problemset/problem/261/D 2600 严格递增
	// - todo https://codeforces.com/contest/261/submission/31860735
	// 若其中一个序列无重复元素，LCS 可以转换成 LIS https://www.luogu.com.cn/problem/P1439 
	// - LC1713 https://leetcode.cn/problems/minimum-operations-to-make-a-subsequence/ 2351
	// 在一维 LIS 的基础上，a[i] 可以从多个数中选一个，问 LIS 最长可以多长
	// - 思路：将各个 a[i] 的可选项从大到小排序，然后拼接成一个序列，求 LIS 即可（关键：从大到小排序避免了在同一个可选项中选择多个元素）
	// 插入后最小化 LIS https://codeforces.com/problemset/problem/1893/B 1700
	// 图上的路径的 LIS https://codeforces.com/problemset/problem/960/F 2100
	// LaIS 与单调栈结合 https://codeforces.com/problemset/problem/1468/A 2200
	// 将所有 a[i] 分成三类：不在任何 LIS / 在至少一个 LIS / 在所有 LIS https://codeforces.com/problemset/problem/486/E 2200
	// - 详细证明 https://www.luogu.com.cn/article/m4267u8v
	// - 弱化版 https://atcoder.jp/contests/abc354/tasks/abc354_f
	// - 删除元素+Dilworth https://ac.nowcoder.com/acm/contest/3782/C
	// - 单点修改后计算 LIS https://codeforces.com/problemset/problem/650/D 2600
	// 状态设计 LIS 计数 https://atcoder.jp/contests/abc237/tasks/abc237_f
	// 逆向题：输入 LIS 返回字典序最小的排列 a https://atcoder.jp/contests/arc125/tasks/arc125_c
	// 反向构造：构造一个 LIS 个数是 x 的数组
	// - 这里把 x 定义成非空 LIS 的个数。把 x 二进制拆分成 2^m1 + 2^m2 + 2^m3 + ... 
	// - 例如 13 = 2^3 + 2^2 + 2^0，我们可以构造 91,91,92,92,93,93,81,82,82,83,83,71,72,73，
	// - 看成三段，每一段的贡献就是前面拆分出的二进制数（这里只是举了个例子，每一段的 gap 可以调大一些以满足构造要求）
	// bitset 优化 https://codeforces.com/contest/1826/problem/E
	// 思想 https://codeforces.com/problemset/problem/1582/F1
	// https://atcoder.jp/contests/arc159/tasks/arc159_d 区间 LIS
	lis := func(a []int) int {
		g := []int{}
		for _, v := range a {
			j := sort.SearchInts(g, v) // 改成 v+1 为非严格递增（即 upper_bound）
			if j < len(g) {
				g[j] = v
			} else {
				g = append(g, v)
			}
		}
		return len(g)
	}

	// 方法二：线段树优化 DP
	// 在值域上建一棵线段树，单点维护的是 a[i] 对应的 dp 值，区间维护的就是一段值域的 dp 的最大值
	// 转移时，查询 < a[i] 的最大值，单点更新到线段树的 a[i] 上
	// 这种做法也可以做到 O(nlogn)，且更加灵活
	// https://www.acwing.com/problem/content/description/3665/
	// LC2407 https://leetcode.cn/problems/longest-increasing-subsequence-ii/

	// 方法三：平衡树
	// todo 参考 https://leetcode.cn/problems/longest-increasing-subsequence-ii/solution/jianjie-by-xing-chen-26-ydqp/

	// 方法四：分治 + 单调队列
	// todo 参考 https://leetcode.cn/problems/longest-increasing-subsequence-ii/solution/fen-zhi-by-heltion-h31y/

	// 每个前缀的 LIS
	// LC1964 https://leetcode.cn/problems/find-the-longest-valid-obstacle-course-at-each-position/ 1933
	lisAll := func(a []int) []int {
		n := len(a)
		lis := make([]int, n)
		g := []int{}
		for i, v := range a {
			p := sort.SearchInts(g, v) // 改成 v+1 为非严格递增（即 upper_bound）
			if p < len(g) {
				g[p] = v
			} else {
				g = append(g, v)
			}
			lis[i] = p + 1
		}
		return lis
	}

	// LIS 方案数 O(nlogn)
	// 原理见下面这题官方题解的方法二
	// LC673 https://leetcode.cn/problems/number-of-longest-increasing-subsequence/
	cntLis := func(a []int) int {
		g := [][]int{}   // 保留所有历史信息
		cnt := [][]int{} // 个数前缀和
		for _, v := range a {
			p := sort.Search(len(g), func(i int) bool { return g[i][len(g[i])-1] >= v })
			c := 1
			if p > 0 {
				// 根据 g[p-1] 来计算 cnt
				i := sort.Search(len(g[p-1]), func(i int) bool { return g[p-1][i] < v })
				c = cnt[p-1][len(cnt[p-1])-1] - cnt[p-1][i]
			}
			if p == len(g) {
				g = append(g, []int{v})
				cnt = append(cnt, []int{0, c})
			} else {
				g[p] = append(g[p], v)
				cnt[p] = append(cnt[p], cnt[p][len(cnt[p])-1]+c)
			}
		}
		c := cnt[len(cnt)-1]
		return c[len(c)-1]
	}

	// LIS 相关构造题
	// https://codeforces.com/problemset/problem/1304/D
	// https://atcoder.jp/contests/arc091/tasks/arc091_c

	// 最大上升子序列和
	// 按值从小到大排序，值相同的下标从大到小排序
	// 然后用树状数组或线段树：单点更新，维护前缀最大值
	// https://www.acwing.com/problem/content/3665/

	// 最长公共上升子序列 (LCIS)
	// https://www.acwing.com/problem/content/274/
	// https://codeforces.com/problemset/problem/10/D
	lcis := func(a, b []int) int {
		n, m := len(a), len(b)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, m)
		}
		for i, v := range a {
			mx := 0
			for j, w := range b {
				if v == w {
					f[i+1][j] = mx + 1
				} else {
					f[i+1][j] = f[i][j]
				}
				if w < v {
					mx = max(mx, f[i][j])
				}
			}
		}
		return slices.Max(f[n])
	}

	// LCIS 打印方案
	lcisPath := func(a, b []int) (ans int, lcis []int) {
		n, m := len(a), len(b)
		f := make([][]int, n+1)
		fa := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, m)
			fa[i] = make([]int, m)
		}
		for i, v := range a {
			mx, k := 0, -1
			for j, w := range b {
				if v == w {
					f[i+1][j] = mx + 1
					fa[i+1][j] = k // k < j
				} else {
					f[i+1][j] = f[i][j]
					fa[i+1][j] = j
				}
				if w < v && f[i][j] > mx {
					mx, k = f[i][j], j
				}
			}
		}
		ansJ := 0
		for j, fv := range f[n] {
			if fv > f[n][ansJ] {
				ansJ = j
			}
		}
		ans = f[n][ansJ]
		var getLCIS func(i, j int)
		getLCIS = func(i, j int) {
			if i == 0 || j < 0 {
				return
			}
			getLCIS(i-1, fa[i][j])
			if fa[i][j] < j {
				lcis = append(lcis, b[j])
			}
		}
		getLCIS(n, ansJ)
		return
	}

	// 长度为 m 的 LIS 个数
	// 赤壁之战 https://www.acwing.com/problem/content/299/
	// 定义 f[i][j] 表示 a[:j+1] 的长度为 i 且以 a[j] 结尾的 LIS
	// 则有 f[i][j] = ∑f[i-1][k]  (k<j && a[k]<a[j])
	// 注意到当 j 增加 1 时，只多了 k=j 这一个新决策，这样可以用树状数组来维护
	// 复杂度 O(mnlogn)
	countLIS := func(a []int, m int) int {
		// 将 a 离散化成从 2 开始的序列
		b := append([]int(nil), a...)
		sort.Ints(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + 2
		}

		n := len(a)
		tree := make([]int, n+2)
		add := func(i, val int) {
			for ; i < len(tree); i += i & -i {
				tree[i] = (tree[i] + val) % mod
			}
		}
		sum := func(i int) (res int) {
			for ; i > 0; i &= i - 1 {
				res = (res + tree[i]) % mod
			}
			return
		}

		f := make([][]int, m+1)
		for i := range f {
			f[i] = make([]int, n)
		}
		for i := 1; i <= m; i++ {
			tree = make([]int, n+2)
			if i == 1 {
				add(1, 1)
			}
			for j, v := range a {
				f[i][j] = sum(v - 1)
				add(v, f[i-1][j])
			}
		}
		ans := 0
		for _, v := range f[m] {
			ans = (ans + v) % mod
		}
		return ans
	}

	// 本质不同非空子序列个数
	// 详细讲解见 https://leetcode.cn/problems/distinct-subsequences-ii/solution/xi-fen-wen-ti-fu-za-du-you-hua-pythonjav-1ihu/
	// - [940. 不同的子序列 II](https://leetcode.cn/problems/distinct-subsequences-ii/) 1985
	// - [1987. 不同的好子序列数目](https://leetcode.cn/problems/number-of-unique-good-subsequences/) 2422 倒序遍历
	// 需要一点构造能力 https://codeforces.com/problemset/problem/645/E
	distinctSubsequence := func(s string) int {
		f := [26]int{}
		sumF := 0
		for _, b := range s {
			b -= 'a'
			tmp := (sumF + mod - f[b]) % mod
			f[b] = (sumF + 1) % mod
			sumF = (tmp + f[b]) % mod
		}
		// 把空的也算上
		//sumF = (sumF + 1) % mod
		return sumF
	}

	// 扩展：长度为 k 的本质不同子序列个数
	// 返回一个数组 f，f[k] 表示长度为 k 的本质不同子序列个数
	// https://codeforces.com/problemset/problem/1183/H
	// https://ac.nowcoder.com/acm/contest/4853/C 题解 https://ac.nowcoder.com/discuss/394080
	//distinctSubsequenceWithFixedLength := func(s string) []int {
	//	const mod int = 1e9 + 7
	//
	//	panic("impl me")
	//	//f := [26]int{}
	//	//
	//	//return sumF
	//}

	// 滚动数组写法
	distinctSubsequence = func(s string) int {
		last := make([]int, 26)
		f := 1
		for _, v := range s {
			v -= 'a'
			res := f - last[v]
			if res < 0 {
				res += mod
			}
			f = (f + res) % mod
			last[v] = (last[v] + res) % mod
		}
		return (f + mod - 1) % mod // 去掉空序列
	}

	// 回文串：中心扩展法
	// 原理见 https://leetcode.cn/problems/palindromic-substrings/solutions/379987/hui-wen-zi-chuan-by-leetcode-solution/
	// LC647 https://leetcode.cn/problems/palindromic-substrings/
	// LC2472 https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/
	// https://ac.nowcoder.com/acm/contest/64272/D
	palindromeO1Space := func(s string) {
		n := len(s)
		for i := 0; i < 2*n-1; i++ { // i 为偶数表示奇回文串，i 为奇数表示偶回文串
			l, r := i/2, i/2+i%2
			// 从 s[i/2..i/2(+1)] 开始扩展
			// do init ...

			for l >= 0 && r < n && s[l] == s[r] {
				// do s[l..r] ...

				l--
				r++
			}
		}
	}

	// O(n^2) 求每个子串是否是回文的
	// 一般用于 DP 预处理
	// LC132 https://leetcode.cn/problems/palindrome-partitioning-ii/
	// LC2472 https://leetcode.cn/problems/maximum-number-of-non-overlapping-palindrome-substrings/
	// https://codeforces.com/problemset/problem/835/D
	isPalindrome := func(s string) [][]bool {
		n := len(s)
		isP := make([][]bool, n)
		for i := range isP {
			isP[i] = make([]bool, n)
		}
		for l := n - 1; l >= 0; l-- {
			for r := l; r < n; r++ {
				isP[l][r] = s[l] == s[r] && (r-l < 3 || isP[l+1][r-1])
			}
		}
		return isP
	}

	// 回文串最小分割次数
	// 紫书例题 9-7，UVa 11584 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=27&page=show_problem&problem=2631
	// LC132 https://leetcode.cn/problems/palindrome-partitioning-ii/
	minPalindromeCut := func(s string) int {
		n := len(s)
		isP := make([][]bool, n)
		for i := range isP {
			isP[i] = make([]bool, n)
		}
		for l := n - 1; l >= 0; l-- {
			for r := l; r < n; r++ {
				isP[l][r] = s[l] == s[r] && (r-l < 3 || isP[l+1][r-1])
			}
		}

		f := make([]int, n)
		for i, prefixP := range isP[0] {
			if prefixP { // f[i] = 0
				continue
			}
			f[i] = int(1e9)
			for j := 0; j < i; j++ {
				if isP[j+1][i] {
					f[i] = min(f[i], f[j]+1)
				}
			}
		}
		return f[n-1]
	}

	/* 背包问题

	0-1 背包 0-1 Knapsack
	完全背包  Unbounded Knapsack
	多重背包  Bounded Knapsack

	这类问题可以从物品选择次序的无后效性入手
	子区间 -> 前缀和
	子序列 -> 背包
	https://en.wikipedia.org/wiki/Knapsack_problem
	https://codeforces.com/blog/entry/59606
	浅谈 ZKP 问题 https://www.luogu.com.cn/blog/xww666/qian-tan-zkp-wen-ti-gai-post
	另见 math_ntt.go 中的生成函数

	O(n√nw) 的 shuffle 做法（这里 w=max(wi)）
	https://codeforces.com/blog/entry/50036 https://codeforces.com/contest/739/problem/E
	https://arxiv.org/pdf/2308.11307.pdf
	http://acm.hdu.edu.cn/showproblem.php?pid=6804

	NOTE: 若求能否凑成 1,2,3,...,M，只需判断 f[i] 是否为正 LC1049 https://leetcode.cn/problems/last-stone-weight-ii/
	套题 https://www.acwing.com/problem/
	混合背包 https://www.luogu.com.cn/problem/P1833
	*/

	// 0-1 背包 (n 个物品，背包容量为 maxW)   0-1 Knapsack
	// 视频讲解：https://www.bilibili.com/video/BV16Y411v7Y6/
	// 状态：从前 i 个物品中选择若干个，当容量限制为 j 时能获得的最大价值和  i∈[0,n-1], j∈[0,maxW]
	// 初始值：f(0,j)=0  j∈[0,maxW]
	// 除开初始状态，每个状态有两个来源，决策为 max：
	// - 不选第 i 个物品：f(i-1,j) -> f(i,j)
	// - 选第 i 个物品：f(i-1,j-wi)+vi -> f(i,j)   j≥wi
	// 最优解为 f(n-1,maxW)
	// https://oi-wiki.org/dp/knapsack/
	// 关于「超大背包问题」见 search.go
	//
	// 模板题 
	// - https://www.luogu.com.cn/problem/P1048 
	// - https://www.luogu.com.cn/problem/P2871
	// - https://atcoder.jp/contests/dp/tasks/dp_d
	// 恰好装满 LC2915 https://leetcode.cn/problems/length-of-the-longest-subsequence-that-sums-to-target/
	// LC2787 https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/
	// LC2291 https://leetcode.cn/problems/maximum-profit-from-trading-stocks/
	// 转换 LC494 https://leetcode.cn/problems/target-sum/
	//            https://atcoder.jp/contests/abc274/tasks/abc274_d
	// 转换 LC1049 https://leetcode.cn/problems/last-stone-weight-ii/
	// LC2431 https://leetcode.cn/problems/maximize-total-tastiness-of-purchased-fruits/
	// 状压 LC1125 https://leetcode.cn/problems/smallest-sufficient-team/
	// 转换 https://leetcode.com/discuss/interview-question/2677093/Snowflake-oror-Tough-OA-question-oror-How-to-solve
	// 转换 https://atcoder.jp/contests/dp/tasks/dp_x
	// https://atcoder.jp/contests/abc281/tasks/abc281_d 二维
	// https://atcoder.jp/contests/abc192/tasks/abc192_f 枚举+二维
	// 转换 https://codeforces.com/problemset/problem/1516/C 1700
	// 转换 https://codeforces.com/problemset/problem/1381/B 1800
	// 抽屉原理 https://codeforces.com/contest/577/problem/B 1900
	// 恰好组成 k 的数中能恰好组成哪些数 https://codeforces.com/problemset/problem/687/C 1900
	// 打印方案 https://codeforces.com/problemset/problem/864/E 2000
	// 转移对象是下标 https://codeforces.com/problemset/problem/981/E 2200
	// 排序+转换 https://codeforces.com/problemset/problem/1203/F2 2300
	// 转移对象是下标 https://codeforces.com/edu/course/2/lesson/9/3/practice/contest/307094/problem/I
	// - f[i][j] 表示前 i 个数，凑成 j 的所有方案中，最小下标的最大值	// 变形，需要多加一个维度 https://atcoder.jp/contests/abc275/tasks/abc275_f
	// 正难则反 https://atcoder.jp/contests/tenka1-2019/tasks/tenka1_2019_d
	// 贡献 https://atcoder.jp/contests/abc159/tasks/abc159_f
	// https://atcoder.jp/contests/agc020/tasks/agc020_c 所有非空子集和的中位数
	// NOIP06·提高 金明的预算方案（也可以用树上背包做）https://www.luogu.com.cn/problem/P1064
	// EXTRA: 恰好装满（相当于方案数不为 0）LC416 https://leetcode.cn/problems/partition-equal-subset-sum/
	//        必须定义成恰好装满（紫书例题 9-5，UVa 12563）https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=441&page=show_problem&problem=4008
	// EXTRA: 背包容量为 0 https://codeforces.com/problemset/problem/366/C 1900
	// EXTRA: 二维费用 LC474 https://leetcode.cn/problems/ones-and-zeroes/ https://www.luogu.com.cn/problem/P1507 
	// EXTRA: 把一个维度转换成 DP 的定义 https://codeforces.com/problemset/problem/837/D
	// EXTRA: 离散化背包 https://codeforces.com/contest/366/submission/61452111
	// EXTRA: 位运算背包 https://ac.nowcoder.com/acm/problem/270788 AND 和 OR 思路是一样的
	zeroOneKnapsack := func(values, weights []int, maxW int) int {
		// 至多装满
		f := make([]int, maxW+1)
		for i, w := range weights {
			v := values[i]
			// 这里 j 的初始值可以优化成前 i 个物品的重量之和（但不能超过 maxW）
			for j := maxW; j >= w; j-- {
				f[j] = max(f[j], f[j-w]+v)
			}
		}
		return f[maxW]
	}

	// 0-1 背包 EXTRA: 恰好装满
	// https://leetcode.cn/contest/sf-tech/problems/cINqyA/
	// 转换 二维费用 https://codeforces.com/problemset/problem/730/J
	zeroOneKnapsackExactlyFull := func(values, weights []int, maxW int) {
		f := make([]int, maxW+1)
		for i := range f {
			f[i] = -1e9 // -1e18
		}
		f[0] = 0
		for i, w := range weights {
			v := values[i]
			for j := maxW; j >= w; j-- {
				f[j] = max(f[j], f[j-w]+v)
			}
		}
		for i := maxW; i >= 0; i-- {
			if f[i] >= 0 { // 能恰好装满 i，此时背包物品价值和的最大值是 f[i]
				// ...
			}
		}
	}

	// 0-1 背包 EXTRA: 至少装入重量和为 maxW 的物品，求价值和的最小值 https://www.luogu.com.cn/problem/P4377
	// f[0] 表示至少为 0 的情况，也表示没有任何约束的情况
	// 比如选第 i 个物品后容量 <=0 了，那就表示前面的 i-1 个物品可以不受约束地随意选或不选了
	// 转换 https://codeforces.com/problemset/problem/19/B LC2742 https://leetcode.cn/problems/painting-the-walls/
	// 二维费用的情况+价值最小 https://ac.nowcoder.com/acm/contest/6218/C
	zeroOneKnapsackAtLeastFillUp := func(values, weights []int, maxW int) int {
		f := make([]int, maxW+1)
		for i := range f {
			f[i] = 1e9 // 1e18
		}
		f[0] = 0 // 无任何限制
		for i, v := range values {
			w := weights[i]
			for j := maxW; j >= 0; j-- {
				f[j] = min(f[j], f[max(j-w, 0)]+v) // 如果 j <= w，那么选之后没有任何限制，所以从 f[0] 转移过来
			}
		}

		{
			// 另一种写法
			for i, v := range values {
				w := weights[i]
				for j := maxW; j >= 0; j-- {
					k := min(j+w, maxW)
					f[k] = min(f[k], f[j]+v)
				}
			}
		}

		return f[maxW]
	}

	// 0-1 背包 EXTRA: 从序列 a 中选若干个数，使其总和为 sum 的方案数
	// 常见题目是算严格分拆（必须用不同数字） LC2787 https://leetcode.cn/problems/ways-to-express-an-integer-as-sum-of-powers/
	// - https://oeis.org/A000009
	// NOTE: 1,1,1,...1(32个1),s-32,s-31,...,s 可以让方案数恰好为 2^32
	// 二维+上限+下限 LC879 https://leetcode.cn/problems/profitable-schemes/
	// https://atcoder.jp/contests/arc060/tasks/arc060_a
	// https://codeforces.com/problemset/problem/1673/C
	// 转换 https://atcoder.jp/contests/abc169/tasks/abc169_f
	// 转换 https://codeforces.com/problemset/problem/478/D
	// 转换 LC494 https://leetcode.cn/problems/target-sum/
	// 转换 LC1434 https://leetcode.cn/problems/number-of-ways-to-wear-different-hats-to-each-other/
	// 由于顺序不同也算方案，所以这题需要正序递推（爬楼梯）LC377 https://leetcode.cn/problems/combination-sum-iv/
	zeroOneWaysToSum := func(a []int, sum int) int {
		f := make([]int, sum+1)
		f[0] = 1
		for _, v := range a {
			for j := sum; j >= v; j-- {
				f[j] += f[j-v] // % mod
			}
		}
		return f[sum]
	}

	// 0-1 背包 EXTRA: 打印字典序最小的方案
	// 倒序遍历物品，同时用 fa 数组记录转移来源，这样跑完 DP 后，从第一个物品开始即可得到字典序最小的方案
	// https://www.acwing.com/problem/content/description/12/
	zeroOneKnapsackLexicographicallySmallestResult := func(values, weights []int, maxW int) (ans []int) {
		n := len(values)
		f := make([]int, maxW+1) // fill
		//f[0] = 0
		fa := make([][]int, n)
		for i := n - 1; i >= 0; i-- {
			fa[i] = make([]int, maxW+1)
			for j := range fa[i] {
				fa[i][j] = j // 注意：<w 的转移来源也要标上！
			}
			v, w := values[i], weights[i]
			for j := maxW; j >= w; j-- {
				if f[j-w]+v >= f[j] { // 注意这里要取等号，从而保证尽可能地从字典序最小的方案转移过来
					f[j] = f[j-w] + v
					fa[i][j] = j - w
				}
			}
		}
		for i, j := 0, maxW; i < n; {
			if fa[i][j] == j { // &&  weights[i] > 0      考虑重量为 0 的情况，必须都选上
				i++
			} else {
				ans = append(ans, i+1) // 下标从 1 开始
				j = fa[i][j]
				i++ // 完全背包的情况，这行去掉
			}
		}
		return
	}

	// 0-1 背包 EXTRA: 价值主导的 0-1 背包
	// 适用于背包容量很大，但是物品价值不高的情况
	// 把重量看成价值，价值看成重量，求同等价值下能得到的最小重量，若该最小重量不超过背包容量，则该价值合法。所有合法价值的最大值即为答案
	// 时间复杂度 O(n * sum(values)) 或 O(n^2 * maxV)
	// https://atcoder.jp/contests/dp/tasks/dp_e
	// https://codeforces.com/contest/1974/problem/E 1800
	// https://codeforces.com/contest/1650/problem/F 2200
	zeroOneKnapsackByValue := func(values, weights []int, maxW int) int {
		totValue := 0
		for _, v := range values {
			totValue += v
		}
		f := make([]int, totValue+1)
		for i := range f {
			f[i] = 1e18
		}
		f[0] = 0
		totValue = 0
		for i, v := range values {
			w := weights[i]
			totValue += v
			for j := totValue; j >= v; j-- {
				f[j] = min(f[j], f[j-v]+w)
			}
		}
		for i := totValue; ; i-- {
			if f[i] <= maxW {
				return i
			}
		}
	}

	// todo 回退背包

	// 完全背包   Unbounded Knapsack
	// 视频讲解 https://www.bilibili.com/video/BV16Y411v7Y6/
	// 更快的做法 https://www.zhihu.com/question/26547156/answer/1181239468
	// https://github.com/hqztrue/shared_materials/blob/master/codeforces/101064%20L.%20The%20Knapsack%20problem%20156ms_short.cpp
	// https://www.luogu.com.cn/problem/P1616
	// 至少 https://www.luogu.com.cn/problem/P2918
	// 恰好装满 LC322 https://leetcode.cn/problems/coin-change/
	// EXTRA: 恰好装满+打印方案 LC1449 https://leetcode.cn/problems/form-largest-integer-with-digits-that-add-up-to-target/
	// 【脑洞】求极限：lim_{maxW->∞} f[maxW]/maxW
	unboundedKnapsack := func(values, weights []int, maxW int) int {
		f := make([]int, maxW+1) // fill
		//f[0] = 0
		for i, v := range values {
			w := weights[i]
			for j := w; j <= maxW; j++ {
				f[j] = max(f[j], f[j-w]+v)
			}
		}
		return f[maxW]
	}

	// 完全背包 EXTRA: 方案数
	// LC518 https://leetcode.cn/problems/coin-change-ii/
	// https://codeforces.com/problemset/problem/1673/C 1500
	// https://www.luogu.com.cn/problem/P1832
	// https://www.luogu.com.cn/problem/P6205（需要高精）
	unboundedWaysToSum := func(a []int, total int) int {
		f := make([]int, total+1)
		f[0] = 1
		for _, v := range a {
			for j := v; j <= total; j++ {
				f[j] = (f[j] + f[j-v]) % mod
			}
		}
		return f[total]
	}

	// 完全背包 EXTRA: 二维费用方案数
	// 注意：「恰好使用 m 个物品」这个条件要当成一种费用来看待
	// https://codeforces.com/problemset/problem/543/A

	// 多重背包   Bounded Knapsack
	// 模板题 https://codeforces.com/problemset/problem/106/C
	//       https://www.luogu.com.cn/problem/P1776
	// todo 多重背包+完全背包 https://www.luogu.com.cn/problem/P1782 https://www.luogu.com.cn/problem/P1833 https://www.luogu.com.cn/problem/P2851
	// http://acm.hdu.edu.cn/showproblem.php?pid=2844 http://poj.org/problem?id=1742
	// https://www.luogu.com.cn/problem/P6771 http://poj.org/problem?id=2392
	// https://codeforces.com/contest/999/problem/F
	// https://codeforces.com/problemset/problem/95/E
	// todo 打印方案

	// 多重背包 - 未优化
	// 转换（价值主导）https://codeforces.com/problemset/problem/922/E（由于要取 min 所以不能用二进制优化）
	boundedKnapsack := func(stocks, values, weights []int, maxW int) int {
		n := len(stocks)
		f := make([][]int, n+1)
		for i := range f {
			f[i] = make([]int, maxW+1)
		}
		for i, num := range stocks {
			v, w := values[i], weights[i]
			for j := range f[i] {
				// 枚举选了 k=0,1,2,...num 个第 i 种物品
				for k := 0; k <= num && k*w <= j; k++ {
					f[i+1][j] = max(f[i+1][j], f[i][j-k*w]+k*v)
				}
			}
		}
		return f[n][maxW]
	}

	// 多重背包 - 优化 1 - 二进制优化
	// 由于不需要队列，常数比多重背包小，可能比多重背包更快
	// 注意循环上界可以从 maxW 优化成 min(sum(w*num), maxW)
	boundedKnapsackBinary := func(stocks, values, weights []int, maxW int) int {
		f := make([]int, maxW+1) // 求最小就 fill inf
		for i, num := range stocks {
			v, w := values[i], weights[i]
			for k1 := 1; num > 0; k1 <<= 1 {
				k := min(k1, num)
				for j := maxW; j >= k*w; j-- {
					f[j] = max(f[j], f[j-k*w]+k*v) // 求最小改成 min
				}
				num -= k
			}
		}
		return f[maxW]
	}

	// 多重背包 - 优化 2 - 单调队列优化
	// 参考挑战 p.340
	// 时间复杂度 O(n*maxW)
	// 注意循环上界可以从 maxW 优化成 min(sum(w*num), maxW)
	boundedKnapsackMonotoneQueue := func(stocks, values, weights []int, maxW int) int {
		f := make([]int, maxW+1) // 求最小就 fill inf
		type pair struct{ maxF, j int }
		q := []pair{}
		for i, num := range stocks {
			v, w := values[i], weights[i]
			for rem := 0; rem < w; rem++ { // 按照 j%w 的结果，分组转移，rem 表示 remainder
				q = q[:0]
				// 为什么压缩维度了还可以正着枚举？因为转移来源都存到单调队列里面了，正序倒序都可以
				// 并且这样相比倒着枚举，不需要先往队列里面塞 num 个数据，更加简洁
				for j := 0; j*w+rem <= maxW; j++ {
					t := f[j*w+rem] - j*v
					for len(q) > 0 && t >= q[len(q)-1].maxF { // 求最小这里改成 <=
						q = q[:len(q)-1] // 及时去掉无用数据
					}
					q = append(q, pair{t, j})
					// 本质是查表法，q[0].maxF 就表示 f[(j-1)*w+r]-(j-1)*v, f[(j-2)*w+r]-(j-2)*v, …… 这些转移来源的最大值
					f[j*w+rem] = q[0].maxF + j*v // 把物品个数视作两个 j 的差（前缀和思想）
					if j-q[0].j == num {         // 至多选 num 个物品
						q = q[1:] // 及时去掉无用数据
					}
				}
			}
		}
		return f[maxW]
	}

	// 多重背包 - 求方案数 - 前缀和优化
	// 每个物品的体积都是 1
	// 如果题目要求每种物品至少选一个，可以把每个 cnts[i] 都减一，maxW 减去 len(cnts)，这样就转换成了每种物品至少选 0 个的情况了
	// 讲解 https://leetcode.cn/problems/find-the-original-typed-string-ii/solutions/2966856/zheng-nan-ze-fan-qian-zhui-he-you-hua-dp-5mi9/
	// 挑战 pp.68-69 多重集组合数
	// 另见 math_comb.go 中的「多重集组合数」容斥做法
	// LC3333 https://leetcode.cn/problems/find-the-original-typed-string-ii/
	// LC2902 https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/ 2759
	// LC2585 https://leetcode.cn/problems/number-of-ways-to-earn-points/ 1910
	// LC1155 https://leetcode.cn/problems/number-of-dice-rolls-with-target-sum/
	// LCR185 https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/
	boundedKnapsackWays := func(cnts []int, maxW int) int {
		// 从 len(cnts) 种物品中选出【至多】maxW 个物品的方案数
		// f[i+1][j] 表示从前 i 种物品中选恰好 j 个物品的方案数（第一维度优化掉）
		f := make([]int, maxW+1)
		f[0] = 1
		for _, c := range cnts {
			// 原地计算 f 的前缀和
			for j := 1; j <= maxW; j++ {
				f[j] = (f[j] + f[j-1]) % mod
			}
			// 计算子数组和
			for j := maxW; j > c; j-- {
				f[j] -= f[j-c-1] // 可能会算出负数，所以最后返回的时候要调整一下
			}
		}

		ans := 0
		for _, x := range f {
			ans += x
		}
		return (ans%mod + mod) % mod
	}

	// 多重背包 - 求方案数 - 同余前缀和优化
	// 第 i 种物品有 c 个，每个物品的体积都是 x，记录在 cnt 中
	// 讲解 https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/solution/duo-zhong-bei-bao-fang-an-shu-cong-po-su-f5ay/
	boundedKnapsackWays2 := func(a []int) []int {
		total := 0
		cnt := map[int]int{}
		for _, x := range a {
			total += x
			cnt[x]++
		}

		// total 表示背包容量
		f := make([]int, total+1)
		f[0] = cnt[0] + 1
		delete(cnt, 0)

		maxS := 0
		for x, c := range cnt {
			maxS += x * c
			// 同余前缀和
			for j := x; j <= maxS; j++ {
				f[j] = (f[j] + f[j-x]) % mod
			}
			for j := maxS; j >= x*(c+1); j-- {
				f[j] = (f[j] - f[j-x*(c+1)] + mod) % mod
			}
		}

		// f[j] 表示所选物品体积之和恰好为 j 的方案数
		// 注意 f[0]
		return f
	}

	// 分组背包·每组至多选一个（恰好选一个见后面）
	// https://www.acwing.com/problem/content/9/
	// https://www.luogu.com.cn/problem/P1757
	// LC2218 https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/
	// https://codeforces.com/problemset/problem/148/E
	// todo 进一步优化 https://codeforces.com/problemset/problem/1442/D
	// 方案数（可以用前缀和优化）https://www.luogu.com.cn/problem/P1077
	groupKnapsack := func(groups [][]struct{ value, weight int }, maxW int) int {
		f := make([]int, maxW+1)
		for _, g := range groups {
			// 这里 j 的初始值可以优化成前 i 个组的每组最大重量之和（但不能超过 maxW）
			for j := maxW; j >= 0; j-- {
				for _, it := range g {
					v, w := it.value, it.weight
					if w <= j {
						// ！如果 it.w 可能为 0 则需要用 f[2][] 来滚动（或者保证每组至多一个 0 且 0 在该组最前面）
						f[j] = max(f[j], f[j-w]+v)
					}
				}
			}
		}
		return f[maxW]
	}

	// todo 撤销计数
	//  https://leetcode.cn/circle/article/YnZBve/

	// 分组背包·每组恰好选一个
	// 允许物品重量为 0
	// https://atcoder.jp/contests/abc240/tasks/abc240_c
	// LC1981 https://leetcode.cn/problems/minimize-the-difference-between-target-and-chosen-elements/
	// 与二分图染色结合 https://codeforces.com/problemset/problem/1354/E
	// 转换 https://codeforces.com/problemset/problem/1637/D
	groupKnapsackFill := func(groups [][]int, maxW int) []bool {
		f := make([]bool, maxW+1) // f[i][j] 表示能否从前 i 组物品中选出重量恰好为 j 的，且每组都恰好选一个物品
		f[0] = true
		for _, g := range groups {
		next:
			for j := maxW; j >= 0; j-- { // 这里 j 的初始值可以优化至前 i 组的最大元素值之和
				for _, w := range g {
					if w <= j && f[j-w] {
						f[j] = true
						continue next
					}
				}
				f[j] = false // 由于我们是滚动数组的写法，f[i][j] 无法满足时要标记成 false
			}
		}
		return f // f[j] 表示从每组恰好选一个，能否凑成重量 j
	}

	// 树上背包/树形背包/依赖背包
	// 时间复杂度 O(n^2)，因为一对点只在 LCA 处被合并一次
	// 树上背包的上下界优化 https://ouuan.github.io/post/%E6%A0%91%E4%B8%8A%E8%83%8C%E5%8C%85%E7%9A%84%E4%B8%8A%E4%B8%8B%E7%95%8C%E4%BC%98%E5%8C%96/
	// 子树合并背包的复杂度证明 https://blog.csdn.net/lyd_7_29/article/details/79854245
	// 复杂度 https://codeforces.com/blog/entry/100910
	// 复杂度 https://leetcode.cn/circle/discuss/t7l62c/
	// https://www.cnblogs.com/shaojia/p/15520224.html
	// https://snuke.hatenablog.com/entry/2019/01/15/211812
	// 复杂度优化 https://loj.ac/d/3144
	// https://zhuanlan.zhihu.com/p/103813542
	//
	// https://loj.ac/p/160
	// https://www.luogu.com.cn/problem/P2014 https://www.acwing.com/problem/content/10/ https://www.acwing.com/problem/content/288/
	// 加强版 https://www.luogu.com.cn/problem/U53204
	// https://www.luogu.com.cn/problem/P1272
	// https://www.luogu.com.cn/problem/P1273
	// 加强版 https://www.luogu.com.cn/problem/U53878
	// https://www.luogu.com.cn/problem/P3177
	// NOIP06·提高 金明的预算方案 https://www.luogu.com.cn/problem/P1064
	treeKnapsack := func(g [][]int, items []struct{ value, weight int }, root, maxW int) int {
		var dfs func(int) []int
		dfs = func(v int) []int {
			it := items[v]
			f := make([]int, maxW+1)
			for i := it.weight; i <= maxW; i++ {
				f[i] = it.value // 根节点必须选
			}
			for _, to := range g[v] {
				ft := dfs(to)
				for j := maxW; j >= it.weight; j-- {
					// 类似分组背包，枚举分给子树 to 的容量 w，对应的子树的最大价值为 dt[w]
					// w 不可超过 j-it.w，否则无法选择根节点
					for w := 0; w <= j-it.weight; w++ {
						f[j] = max(f[j], f[j-w]+ft[w])
					}
				}
			}
			return f
		}
		return dfs(root)[maxW]
	}

	/* 区间 DP
	一般来说转移是合并区间或者分解区间
	套路 https://www.luogu.com.cn/blog/BreakPlus/ou-jian-dp-zong-jie-ti-xie

	求解关于某个序列的最优性质，要求大区间的最优解可以依赖于小区间的最优解
	一般定义 f[i][j] 表示子数组 a[i] 到 a[j] 的最优解
	LC375 https://leetcode.cn/problems/guess-number-higher-or-lower-ii/
	LC312 戳气球 https://leetcode.cn/problems/burst-balloons/
	LC664 打印机 https://leetcode.cn/problems/strange-printer/
	LC678 也有非 DP 做法 https://leetcode.cn/problems/valid-parenthesis-string/
	LC1312 插入形成回文 https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/ 1787
	- https://www.luogu.com.cn/problem/P1435
	LC3040 https://leetcode.cn/problems/maximum-number-of-operations-with-the-same-score-ii/
	LC1130 https://leetcode.cn/problems/minimum-cost-tree-from-leaf-values/ 1919
	LC1770 https://leetcode.cn/problems/maximum-score-from-performing-multiplication-operations/ 2068
	LC1547 https://leetcode.cn/problems/minimum-cost-to-cut-a-stick/ 2116
	LC1039 最优三角剖分 https://leetcode.cn/problems/minimum-score-triangulation-of-polygon/ 2130
	- 方案数 不保证凸 https://codeforces.com/problemset/problem/437/E 2500
	LC546 移除盒子 状态定义 https://leetcode.cn/problems/remove-boxes/ 从 CF 难度来看，这题可以评 2900（力扣难度）
	- 原题是紫书例题 9-27 https://www.luogu.com.cn/problem/UVA10559 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=17&page=show_problem&problem=1500
	- https://codeforces.com/problemset/problem/1107/E 2400
	LC3018 https://leetcode.cn/problems/maximum-number-of-removal-queries-that-can-be-processed-i/ 会员题
	https://atcoder.jp/contests/abc163/tasks/abc163_e 贪心
	另见 LPS

	最短括号超序列 UVa1626 紫书例题 9-10 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=825&page=show_problem&problem=4501
	容斥 https://atcoder.jp/contests/abc106/tasks/abc106_d
	https://codeforces.com/problemset/problem/245/H 1800
	https://codeforces.com/problemset/problem/1509/C 1800
	https://codeforces.com/problemset/problem/149/D 1900
	https://codeforces.com/problemset/problem/607/B 1900 回文消除
	- https://leetcode.cn/problems/palindrome-removal/ LC 抄袭 CF
	- https://codeforces.com/problemset/problem/150/D 2600 回文消除 每个长度有不同的权值
	染色【套路】https://codeforces.com/problemset/problem/1114/D 1900
	同色消除【套路】https://www.luogu.com.cn/problem/P4170
	             https://codeforces.com/problemset/problem/1132/F 2000
	https://codeforces.com/problemset/problem/1025/D 2100 状态设计
	https://codeforces.com/problemset/problem/1336/C 2200 状态设计
	- https://www.luogu.com.cn/problem/P3205
	https://atcoder.jp/contests/arc183/tasks/arc183_c 2018=CF2237
	https://codeforces.com/problemset/problem/1198/D 2300 二维区间 DP
	https://codeforces.com/problemset/problem/1107/E 2400
	https://codeforces.com/problemset/problem/1863/F 2600
	todo https://atcoder.jp/contests/abc159/tasks/abc159_f
	https://blog.csdn.net/weixin_43914593/article/details/106163859 算法竞赛专题解析（14）：DP应用--区间DP
	*/

	// 最长回文子序列 (LPS)
	// 即 LCS(s, reverse(s))
	// 视频讲解 https://www.bilibili.com/video/BV1Gs4y1E7EU/
	// 回文子串见下面的 isPalindrome 或者 strings.go 的 manacher
	// LC516  https://leetcode.cn/problems/longest-palindromic-subsequence/
	// LC1682 https://leetcode.cn/problems/longest-palindromic-subsequence-ii/
	// LC730 求个数 https://leetcode.cn/problems/count-different-palindromic-subsequences/
	// LC1216 https://leetcode.cn/problems/valid-palindrome-iii/ 1754
	// LC1312 https://leetcode.cn/problems/minimum-insertion-steps-to-make-a-string-palindrome/ 1787
	//        - https://www.luogu.com.cn/problem/P1435
	// LC1771 https://leetcode.cn/problems/maximize-palindrome-length-from-subsequences/ 2182
	// LC1246 https://leetcode.cn/problems/palindrome-removal/ 2203
	// todo 【需要复习】 树上路径 LPS https://codeforces.com/problemset/problem/1771/D 2100
	longestPalindromeSubsequence := func(s string) int {
		n := len(s)
		f := make([][]int, n)
		for i := range f {
			f[i] = make([]int, n)
		}
		for i := n - 1; i >= 0; i-- {
			f[i][i] = 1
			for j := i + 1; j < n; j++ {
				if s[i] == s[j] {
					f[i][j] = f[i+1][j-1] + 2
				} else {
					f[i][j] = max(f[i+1][j], f[i][j-1])
				}
			}
		}
		return f[0][n-1]
	}

	// 统计区间内回文串个数
	// 返回一个二维数组 f, f[i][j] 表示 s[i:j+1] 内的回文串的个数
	// https://codeforces.com/problemset/problem/245/H
	countPalindromes := func(s string) [][]int {
		n := len(s)
		f := make([][]int, n)
		for i := range f {
			f[i] = make([]int, n)
			f[i][i] = 1
			if i+1 < n && s[i] == s[i+1] {
				f[i][i+1] = 1
			}
		}
		for i := n - 3; i >= 0; i-- {
			for j := i + 2; j < n; j++ {
				if s[i] == s[j] {
					f[i][j] = f[i+1][j-1]
				}
			}
		}
		// 到这里为止，f[i][j] = 1 表示 s[i:j+1] 是回文串
		for i := n - 2; i >= 0; i-- {
			for j := i + 1; j < n; j++ {
				f[i][j] += f[i][j-1] + f[i+1][j] - f[i+1][j-1] // 容斥
			}
		}
		return f
	}

	// 石子合并
	// https://atcoder.jp/contests/dp/tasks/dp_n
	// https://ac.nowcoder.com/acm/contest/1043/A https://ac.nowcoder.com/acm/problem/51170
	// 环形的情况 https://www.luogu.com.cn/problem/P1880
	// 相邻 k 堆的情况（综合①②）LC1000 https://leetcode.cn/problems/minimum-cost-to-merge-stones/ 2423
	mergeStones := func(a []int) int {
		n := len(a)
		sum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		f := make([][]int, n)
		for i := range f {
			f[i] = make([]int, n)
			for j := range f[i] {
				f[i][j] = 1e9
			}
			f[i][i] = 0
		}
		for sz := 2; sz <= n; sz++ {
			for l := 0; l+sz <= n; l++ {
				r := l + sz - 1
				for i := l; i < r; i++ {
					f[l][r] = min(f[l][r], f[l][i]+f[i+1][r])
				}
				f[l][r] += sum[r+1] - sum[l]
			}
		}
		return f[0][n-1]
	}

	/* 环形 DP
	两种解题策略：
	一种是假设在 0 处断开，求一遍 DP，然后强制让 0 和 n-1 上是连通的，再求一遍 DP，取二者最值
	另一种是倍增链
	休息时间 https://www.luogu.com.cn/problem/P6064 https://www.acwing.com/problem/content/290/
	环路运输 https://www.acwing.com/problem/content/291/
	https://www.luogu.com.cn/problem/P1453
	*/

	/* 状态机 DP
	可以用网格图 DP 形象地理解，一般状态机都是 0->1->2 这种，类似只能向右/右下的网格图 DP
	LC123 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iii/
	LC188 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
	LC309 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-cooldown/
	LC1186 https://leetcode.cn/problems/maximum-subarray-sum-with-one-deletion/
	https://codeforces.com/problemset/problem/1178/B 1300 入门
	https://codeforces.com/problemset/problem/698/A 1400 入门
	https://codeforces.com/problemset/problem/2033/C 1400
	https://codeforces.com/problemset/problem/855/B 1500
	https://codeforces.com/problemset/problem/1826/D 1700 式子变形
	https://codeforces.com/problemset/problem/2029/C 1700
	https://codeforces.com/problemset/problem/404/D 1900
	https://codeforces.com/problemset/problem/1613/D 1900 爽
	https://codeforces.com/problemset/problem/623/B 2300
	*/

	/* 分治 DP
	视频讲解 https://www.bilibili.com/video/BV1ne4y177wN/ 第四题
	https://leetcode.cn/contest/tianchi2022/problems/tRZfIV/
	- https://codeforces.com/contest/1442/problem/D
	*/

	/* 博弈 DP
	另见 games.go
	转移：让「自己与对手的分差」最大
	图上博弈 https://codeforces.com/problemset/problem/917/B
	LC464 https://leetcode.cn/problems/can-i-win/
	LC486 https://leetcode.cn/problems/predict-the-winner/
	LC877 https://leetcode.cn/problems/stone-game/ https://nanti.jisuanke.com/t/48
	LC913 https://leetcode.cn/problems/cat-and-mouse/
	LC1025 https://leetcode.cn/problems/divisor-game/
	LC1140 https://leetcode.cn/problems/stone-game-ii/
	LC1406 https://leetcode.cn/problems/stone-game-iii/
	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp%2Cgames
	【记忆化搜索模板】https://codeforces.com/problemset/problem/1738/C
	https://www.lanqiao.cn/problems/8051/learning/?contest_id=146
	*/

	/* 概率 DP / 期望 DP
	https://oi-wiki.org/dp/probability/
	https://en.wikipedia.org/wiki/Probability
	https://en.wikipedia.org/wiki/Expected_value
	https://en.wikipedia.org/wiki/Variance
	https://en.wikipedia.org/wiki/Optional_stopping_theorem
	todo https://codeforces.com/blog/entry/62690
	     https://codeforces.com/blog/entry/62792
	 https://www.luogu.com.cn/blog/Troverld/gai-shuai-ji-wang-xue-xi-bi-ji
	 一类概率期望问题的杀器：势函数和鞅的停时定理 https://www.cnblogs.com/TinyWong/p/12887591.html https://codeforces.com/blog/entry/87598 最后一题
	 鞅与停时定理学习笔记 https://www.luogu.com.cn/blog/gxy001/yang-yu-ting-shi-ding-li-xue-xi-bi-ji
	todo 生成函数与期望 http://www.matrix67.com/blog/archives/4534

	期望的可加性
	https://zhidao.baidu.com/question/259203053.html

	马尔可夫链 Markov chain https://en.wikipedia.org/wiki/Markov_chain
	吸收马尔可夫链 Absorbing Markov chain https://en.wikipedia.org/wiki/Absorbing_Markov_chain
	https://www.bilibili.com/video/BV14y4y1S7ve
	切尔诺夫界 Chernoff bound https://en.wikipedia.org/wiki/Chernoff_bound
	https://leetcode.cn/problems/soup-servings/solutions/1982989/shou-lian-su-du-by-hqztrue-afba/

	一个比较有用的公式（应用：CF1623D）
	E(x) = ∑i*P(x=i) = ∑P(x>=i)

	方差
	σ²(x) = sum(x²)/n - (sum(x)/n)²

	概率 DP
	https://codeforces.com/problemset/problem/16/E
	https://codeforces.com/problemset/problem/540/D
	https://codeforces.com/problemset/problem/678/E
	https://codeforces.com/problemset/problem/2028/E 2300
	LC688 https://leetcode.cn/problems/knight-probability-in-chessboard/
	LC808 https://leetcode.cn/problems/soup-servings/
	LC837 https://leetcode.cn/problems/new-21-game/
	LC1227 数学题 https://leetcode.cn/problems/airplane-seat-assignment-probability/
	LC1230 https://leetcode.cn/problems/toss-strange-coins/
	LC1467 https://leetcode.cn/problems/probability-of-a-two-boxes-having-the-same-number-of-distinct-balls/
	剑指 Offer 60 https://leetcode.cn/problems/nge-tou-zi-de-dian-shu-lcof/

	期望 DP
	入门题 https://atcoder.jp/contests/abc280/tasks/abc280_e
	https://atcoder.jp/contests/abc350/tasks/abc350_e
	- 如果状态转移左右两边都包含 f[i]，则需要移项化简
	https://atcoder.jp/contests/dp/tasks/dp_j
	DAG https://www.luogu.com.cn/problem/P4316
	后缀和优化 https://atcoder.jp/contests/abc326/tasks/abc326_e
	todo https://codeforces.com/contest/1842/problem/G https://codeforces.com/blog/entry/117640
	https://codeforces.com/problemset/problem/235/B 2000
	https://codeforces.com/problemset/problem/1753/C 2000
	https://codeforces.com/problemset/problem/908/D 2200
	https://codeforces.com/problemset/problem/1097/D 2200
	https://codeforces.com/problemset/problem/1623/D 2300
	https://codeforces.com/problemset/problem/1824/B2 2300
	todo https://codeforces.com/problemset/problem/494/C 2600
	todo https://codeforces.com/problemset/problem/1172/C2 2600
	Kick Start 2020 Round F Yeetzhee https://codingcompetitions.withgoogle.com/kickstart/round/000000000019ff48/00000000003f4dea
	todo https://leetcode.cn/contest/ubiquant2022/problems/I3Gm2h/
	 https://ac.nowcoder.com/acm/contest/76681/J

	如下题目来自 https://www.luogu.com/paste/3cp2ob1v
	https://ac.nowcoder.com/acm/contest/80259/F
	https://atcoder.jp/contests/dwacon6th-prelims/tasks/dwacon6th_prelims_b
	https://www.lanqiao.cn/problems/10009/learning/?contest_id=157
	https://ac.nowcoder.com/acm/problem/17412
	https://codefun2000.com/p/P1617
	https://www.codechef.com/problems/LEBOBBLE?tab=statement
	https://yukicoder.me/problems/no/1904
	https://atcoder.jp/contests/abc276/tasks/abc276_f
	https://www.luogu.com.cn/problem/P9217
	https://codeforces.com/problemset/problem/453/A
	https://www.nowcoder.com/exam/test/79818590/detail?pid=23354036&pageSource=testHistory
	https://codeforces.com/problemset/problem/280/C
	https://ac.nowcoder.com/acm/contest/275/G
	https://ac.nowcoder.com/acm/contest/275/I
	https://atcoder.jp/contests/abc149/tasks/abc149_f
	https://atcoder.jp/contests/typical90/tasks/typical90_bn
	https://www.luogu.com.cn/problem/P4550
	https://atcoder.jp/contests/abc326/tasks/abc326_e
	https://ac.nowcoder.com/acm/contest/73760/F
	https://ac.nowcoder.com/acm/contest/80743/F
	https://ac.nowcoder.com/acm/contest/28263/B
	https://ac.nowcoder.com/acm/contest/60771/A
	https://ac.nowcoder.com/acm/contest/60771/B
	https://atcoder.jp/contests/abc184/tasks/abc184_d
	https://atcoder.jp/contests/abc194/tasks/abc194_d
	https://atcoder.jp/contests/abc266/tasks/abc266_e
	https://atcoder.jp/contests/abc263/tasks/abc263_e
	https://www.luogu.com.cn/problem/AT_tdpc_ball
	https://codeforces.com/problemset/problem/1753/C
	https://atcoder.jp/contests/s8pc-4/tasks/s8pc_4_d
	https://ac.nowcoder.com/acm/contest/76681/J
	https://www.acwing.com/problem/content/description/4854/
	https://codefun2000.com/p/P1397
	https://www.nowcoder.com/exam/test/79742799/detail?pid=38431372&pageSource=testHistory
	https://www.lanqiao.cn/problems/9493/learning/?contest_id=152
	https://ac.nowcoder.com/acm/contest/51458/E
	https://yukicoder.me/problems/no/1478
	https://atcoder.jp/contests/abc314/tasks/abc314_e
	https://atcoder.jp/contests/arc016/tasks/arc016_3
	https://atcoder.jp/contests/abc342/tasks/abc342_f
	https://atcoder.jp/contests/past201912-open/tasks/past201912_o
	https://yukicoder.me/problems/no/76
	https://codeforces.com/problemset/problem/1139/D
	https://ac.nowcoder.com/acm/contest/85687/E
	https://codeforces.com/problemset/problem/235/B
	https://www.luogu.com.cn/problem/P1654
	https://www.lanqiao.cn/problems/4519/learning/?page=1&first_category_id=1&tags=概率&sort=difficulty&asc=1
	https://www.lanqiao.cn/problems/8422/learning/?page=2&first_category_id=1&sort=difficulty&tags=%E6%9C%9F%E6%9C%9BDP&asc=1
	https://www.codechef.com/problems/DIAMOND?tab=statement
	https://codeforces.com/problemset/problem/518/D
	https://www.nowcoder.com/exam/test/79718811/detail?pid=31381957
	https://atcoder.jp/contests/abc008/tasks/abc008_3
	https://atcoder.jp/contests/tkppc3/tasks/tkppc3_e
	https://www.lanqiao.cn/problems/3339/learning/?page=1&first_category_id=1&sort=difficulty&tags=%E6%9C%9F%E6%9C%9BDP&asc=1
	https://www.lanqiao.cn/problems/4405/learning/?page=1&first_category_id=1&sort=difficulty&tags=%E6%9C%9F%E6%9C%9BDP&asc=1
	https://yukicoder.me/problems/no/1688
	*/

	/* 状压 DP
	常用于处理包含排列的问题等
	NOTE: 若问题无法划分成小问题，必须考虑各种可能的情况，则可能是 NP 完全问题

	教你一步步思考状压 DP：从记忆化搜索到递推，附题单 https://leetcode.cn/problems/beautiful-arrangement/solution/jiao-ni-yi-bu-bu-si-kao-zhuang-ya-dpcong-c6kd/

	状压 DP 本质上就是在集合与集合之间转移，所以一定要能熟练地把集合语言翻译成位运算
	推荐阅读《从集合论到位运算，常见位运算技巧分类总结！》https://leetcode.cn/circle/discuss/CaOJ45/

	浅谈状压 DP https://www.luogu.com.cn/blog/yijan/zhuang-ya-dp
	https://blog.csdn.net/weixin_43914593/article/details/106432695 算法竞赛专题解析（15）：DP应用--状态压缩DP

	todo 题单 https://www.luogu.com.cn/training/215#problems
	     题单 https://ac.nowcoder.com/acm/problem/collection/808
	     题单 https://ac.nowcoder.com/acm/problem/collection/810

	https://codeforces.com/problemset/problem/991/D 迷你状压
	https://atcoder.jp/contests/abc142/tasks/abc142_e 基础状压 DP
	https://atcoder.jp/contests/abc359/tasks/abc359_d
	LC691 https://leetcode.cn/problems/stickers-to-spell-word/
	LC943 https://leetcode.cn/problems/find-the-shortest-superstring/
	LC1125 状压 0-1 背包 https://leetcode.cn/problems/smallest-sufficient-team/
	LC1411 https://leetcode.cn/problems/number-of-ways-to-paint-n-3-grid/
	LC1931 https://leetcode.cn/problems/painting-a-grid-with-three-different-colors/
	LC2184 https://leetcode.cn/problems/number-of-ways-to-build-sturdy-brick-wall/
	LC2247 https://leetcode.cn/problems/maximum-cost-of-trip-with-k-highways/ 会员题
	LCP53 https://leetcode.cn/problems/EJvmW4/
	LCP69 https://leetcode.cn/problems/rMeRt2/
	LCP76 https://leetcode.cn/problems/1ybDKD/
	todo LCP82 https://leetcode.cn/problems/cnHoX6/
	https://www.luogu.com.cn/problem/P1879
	https://codeforces.com/problemset/problem/16/E 1900 与概率 DP 结合
	https://codeforces.com/problemset/problem/401/D 2000
	https://codeforces.com/problemset/problem/453/B 2000 与质因数分解结合
	https://codeforces.com/problemset/problem/895/C 2000
	https://codeforces.com/problemset/problem/903/F 2200 轮廓线
	https://codeforces.com/problemset/problem/1316/E 2300 与排序贪心结合
	https://codeforces.com/problemset/problem/1955/H 2300
	https://codeforces.com/problemset/problem/1209/E2 2500 循环移位
	https://codeforces.com/problemset/problem/599/E 2600 树上子集状压 DP
	https://codeforces.com/problemset/problem/1430/G 2600
	https://atcoder.jp/contests/abc359/tasks/abc359_d
	https://www.luogu.com.cn/problem/P5369 状态设计的好题

	todo 汉密尔顿路径/回路 Hamiltonian path
	https://en.wikipedia.org/wiki/Hamiltonian_path
	https://en.wikipedia.org/wiki/Hamiltonian_path_problem

	求最大团/最大独立集的记忆化写法见 graph.go
	*/

	// 注：下面的任意排列 DP，也适用于非全排列（子集排列）的情况，因为在计算过程中，我们也计算出了子集排列的 DP

	// 任意排列 DP - 相邻无关 （刷表法）
	// 适用于不需要知道上一个数的场景
	// 时间复杂度通常是 O(n*2^n) 下面的写法常数是 1/2
	// LC526 https://leetcode.cn/problems/beautiful-arrangement/
	//    - https://oeis.org/A320843 Number of permutations sigma of {1,2,...,n} such that sigma(i) divides i or i divides sigma(i) for 1 <= i <= n
	// LC1879 https://leetcode.cn/problems/minimum-xor-sum-of-two-arrays/ 2145
	// LC2172 https://leetcode.cn/problems/maximum-and-sum-of-array/ 2392
	// LC2403 https://leetcode.cn/problems/minimum-time-to-kill-all-monsters/
	// LC2992 https://leetcode.cn/problems/number-of-self-divisible-permutations/
	// https://atcoder.jp/contests/dp/tasks/dp_o
	// https://atcoder.jp/contests/abc381/tasks/abc381_f 1739=CF2026 状态设计
	// https://atcoder.jp/contests/abc199/tasks/abc199_e 1814=CF2083
	// https://codeforces.com/problemset/problem/377/C 2200 枚举来源
	// https://codeforces.com/problemset/problem/743/E 2200 状态设计
	// https://codeforces.com/problemset/problem/1215/E 2200
	// https://codeforces.com/problemset/problem/1238/E 2200 式子变形
	// https://codeforces.com/problemset/problem/327/E 2300 卡常优化 另一种做法是折半枚举
	// todo 状态设计 https://codeforces.com/problemset/problem/744/C 2400
	// https://codeforces.com/problemset/problem/1550/E 2500 状态设计
	permDP := func(a []int, check func(int, int) bool) int {
		n := len(a)
		f := make([]int, 1<<n)
		// 求最小值的题目
		//for i := 1; i < len(f); i++ { f[i] = 1e18 }
		// 计数题目
		f[0] = 1
		for s, fs := range f { // 前面选的下标集合是 s
			if fs == 0 { // 剪枝：用在计数题目上
				continue
			}
			// 考虑第 i 个位置怎么填
			i := bits.OnesCount(uint(s))
			// g[i]&^s
			for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
				lb = cus & -cus
				ns := s | lb
				// 枚举（第 i 个位置）填第 p 个 ...
				p := bits.TrailingZeros(uint(lb))
				v := a[p]
				if check(i, v) {
					f[ns] = (f[ns] + fs) % mod
				}
			}
		}
		return f[len(f)-1]
	}

	// 任意排列 DP - 相邻相关
	// 适用于需要知道上一个数的场景
	// 时间复杂度通常是 O(n^2*2^n) 下面的写法常数约为 1/4 https://oeis.org/A001815
	// LC996 最后答案需要除相同元素个数的阶乘 https://leetcode.cn/problems/number-of-squareful-arrays/ 1932
	// LC2741 https://leetcode.cn/problems/special-permutations/ 2021
	// LC1681 https://leetcode.cn/problems/minimum-incompatibility/ 2390
	// https://codeforces.com/problemset/problem/1950/G 1900
	// https://codeforces.com/problemset/problem/1185/G1 2100
	// https://codeforces.com/problemset/problem/2051/G 2100
	permDP2 := func(a []int, check func(int, int) bool) int {
		n := len(a)
		f := make([][]int, 1<<n)
		for i := range f {
			f[i] = make([]int, n)
			// 求最小值的题目
			//for j := range f[i] { f[i][j] = 1e18 }
		}
		for j := range f[0] {
			f[1<<j][j] = 1 // 排列的第一个数
		}
		for s, fs := range f {
			for _s := uint(s); _s > 0; _s &= _s - 1 {
				i := bits.TrailingZeros(_s)
				if fs[i] == 0 { // 剪枝：用于计数题目，或子集排列题目
					continue
				}
				pre := a[i] // 枚举上一个选的数
				// g[i]&^s
				for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
					lb = cus & -cus
					ns := s | lb
					j := bits.TrailingZeros(uint(lb))
					cur := a[j] // 枚举当前选的数
					if check(pre, cur) {
						f[ns][j] = (f[ns][j] + fs[i]) % mod
					}
				}
			}
		}
		ans := 0
		for _, dv := range f[len(f)-1] {
			ans = (ans + dv) % mod
		}
		return ans
	}

	// 旅行商问题  travelling salesman problem  TSP
	// 图论中的一个等价形式是：给定一个加权完全图（顶点表示城市，边表示道路，权重是道路的距离），求一权值和最小的哈密尔顿回路。
	// 返回一个 ans 数组，ans[i] 表示从 st 出发，访问完所有位置且最后停在 i 的最短路径（注意可能要特判 i==st 的情况）
	// 做法：定义 f[s][i] 表示已访问的集合为 s，最后一个访问的位置是 i 时的最小花费
	//      则有 f[s|1<<j][j] = min(f[s|1<<j][j], f[s][i]+dist[i][j])
	//      枚举 i 和 j 时可以用 TrailingZeros 来直接枚举每个 1 和 0 的位置
	// https://en.wikipedia.org/wiki/Travelling_salesman_problem
	// https://en.wikipedia.org/wiki/Hamiltonian_path HCP
	//
	// 模板题 https://www.luogu.com.cn/problem/P1171 
	//       https://www.luogu.com.cn/problem/P1433 
	//       回路 https://atcoder.jp/contests/abc180/tasks/abc180_e
	//       https://www.acwing.com/problem/content/93/
	//
	// 建模转换题 LC943 https://leetcode.cn/problems/find-the-shortest-superstring/ 2186
	//          LCP13 https://leetcode.cn/problems/xun-bao/
	// LC847 https://leetcode.cn/problems/shortest-path-visiting-all-nodes/ 2201
	//
	// 额外的起点和终点 https://atcoder.jp/contests/abc301/tasks/abc301_e
	// 略微变形 吃加速药水 https://atcoder.jp/contests/abc274/tasks/abc274_e
	// 恰好访问 m 个点 https://codeforces.com/contest/580/problem/D 1800
	// 变体+打印路径 https://codeforces.com/problemset/problem/8/C 2000
	// 矩阵交换行问题 https://codeforces.com/problemset/problem/1102/F
	// EXTRA: 固定起点终点的问题，视问题情况有两种方法：
	//        添加一个节点 https://stackoverflow.com/questions/14527815/how-to-fix-the-start-and-end-points-in-travelling-salesmen-problem
	//        设置距离 https://stackoverflow.com/questions/36086406/traveling-salesman-tsp-with-set-start-and-end-point
	tsp := func(dist [][]int, st int) []int {
		const inf int = 1e9 // 1e18
		f := make([][]int, 1<<len(dist))
		for i := range f {
			f[i] = make([]int, len(dist))
			for j := range f[i] {
				f[i][j] = inf
			}
		}
		// 如果有多个起点的话就初始化多个 f[1<<st[i]][st[i]] = 0
		// 如果有超级源点 src，那么初始化 f[1<<st[i]][st[i]] = dis[src][st[i]]
		// 如果有超级汇点 dst，最后答案为 min(f[-1][i]+dis[i][dst])
		f[1<<st][st] = 0
		for s, fs := range f {
			// 利用位运算快速求出 s 中 1 的位置 i，以及 s 中 0 的位置 j（通过 s 的补集中的 1 的位置求出）
			for _s := uint(s); _s > 0; _s &= _s - 1 {
				i := bits.TrailingZeros(_s)
				if fs[i] == inf {
					continue
				}
				for cus, lb := len(f)-1^s, 0; cus > 0; cus ^= lb {
					lb = cus & -cus
					ns := s | lb
					j := bits.TrailingZeros(uint(lb))
					f[ns][j] = min(f[ns][j], fs[i]+dist[i][j])
				}
			}
		}
		return f[len(f)-1]
	}

	// 无向图最长简单路径
	// 输入的 g 是邻接矩阵，第二维度压缩成一个 int
	// https://codeforces.com/contest/1950/problem/G
	longestSimplePath := func(g []int) int {
		n := len(g)
		f := make([][]int, 1<<n)
		for i := range f {
			f[i] = make([]int, n)
		}
		for j := range f[0] {
			f[1<<j][j] = 1
		}
		ans := 1
		for s, fs := range f {
			for _s := uint(s); _s > 0; _s &= _s - 1 {
				i := bits.TrailingZeros(_s)
				if fs[i] == 0 {
					continue
				}
				for cus, lb := g[i]&^s, 0; cus > 0; cus ^= lb {
					lb = cus & -cus
					ns := s | lb
					j := bits.TrailingZeros(uint(lb))
					f[ns][j] = max(f[ns][j], fs[i]+1)
					ans = max(ans, f[ns][j])
				}
			}
		}
		return ans
	}

	// 无向图简单环数量
	// https://blog.csdn.net/fangzhenpeng/article/details/49078233
	// https://codeforces.com/problemset/problem/11/D
	countCycle := func(g [][]int, n, m int) int {
		ans := 0
		// 取集合 s 的最小值作为起点
		f := make([][]int, 1<<n)
		for i := range f {
			f[i] = make([]int, n)
		}
		for i := 0; i < n; i++ {
			f[1<<i][i] = 1
		}
		for s := range f {
			for v, fs := range f[s] {
				if fs == 0 {
					continue
				}
				for _, w := range g[v] {
					if 1<<w < s&-s {
						continue
					}
					if 1<<w&s == 0 {
						f[s|1<<w][w] += fs
					} else if 1<<w == s&-s {
						ans += fs
					}
				}
			}
		}
		return ans - m/2
	}

	// 枚举子集的子集
	// 复杂度 O(3^n)，证明：元素个数为 k 的集合有 C(n,k) 个，其子集有 2^k 个，故有 ∑C(n,k)*2^k = (2+1)^n = 3^n
	// 例如：f[set] = max{f[set^sub] + sum of sub} for all valid sub
	//
	// 模板题 https://atcoder.jp/contests/dp/tasks/dp_u
	// - [2305. 公平分发饼干](https://leetcode.cn/problems/fair-distribution-of-cookies/) 1887
	// - [1986. 完成任务的最少工作时间段](https://leetcode.cn/problems/minimum-number-of-work-sessions-to-finish-the-tasks/) 1995
	// - [1494. 并行课程 II](https://leetcode.cn/problems/parallel-courses-ii/) 2082
	// - [1723. 完成所有工作的最短时间](https://leetcode.cn/problems/find-minimum-time-to-finish-all-jobs/) 2284 *子集划分型
	// - [1655. 分配重复整数](https://leetcode.cn/problems/distribute-repeating-integers/) 2307
	// - [1349. 参加考试的最大学生数](https://leetcode.cn/problems/maximum-students-taking-exam/) 2386
	//    - https://codeforces.com/contest/1926/problem/F
	// - [1681. 最小不兼容性](https://leetcode.cn/problems/minimum-incompatibility/) 2390 有 O(n^2*2^n) 做法
	//    - 相关（不是状压） https://codeforces.com/contest/626/problem/F
	// - [2572. 无平方子集计数](https://leetcode.cn/problems/count-the-number-of-square-free-subsets/) 2420
	//    - 也可以用 01 背包 / 枚举 square-free
	//    - 更快的做法 https://leetcode.cn/problems/count-the-number-of-square-free-subsets/solution/shu-zhi-fan-wei-zai-da-yi-dian-de-hua-ze-56w3/
	//    - 把相同的大质数归类，避免重复选择（分组背包）
	// - [1994. 好子集的数目](https://leetcode.cn/problems/the-number-of-good-subsets/) 2465
	// - [LCP 53. 守护太空城](https://leetcode.cn/problems/EJvmW4/)
	// - [465. 最优账单平衡](https://leetcode.cn/problems/optimal-account-balancing/)（会员题）
	// todo https://codeforces.com/problemset/problem/1556/F 2500
	// 训练指南第一章例题 29，UVa11825 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=226&page=show_problem&problem=2925
	//    - 将 n 个集合分成尽量多组，使得对于每组，组内所有集合的并集等于全集
	// 训练指南第一章例题 32，WF10，UVa1099 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=245&page=show_problem&problem=3540
	// https://codeforces.com/problemset/problem/599/E 2600 树上子集状压 DP
	subsubDP := func(a []int) int {
		n := len(a)
		m := 1 << n
		// 预处理每个子集的子集和
		sum := make([]int, m)
		for i := range sum {
			for s := uint(i); s > 0; s &= s - 1 {
				v := a[bits.TrailingZeros(s)]
				sum[i] += v
			}
		}
		f := make([]int, m)
		for s, fs := range f {
			t := m - 1 ^ s
			// 枚举补集的非空子集
			for sub := t; sub > 0; sub = (sub - 1) & t {
				ss := s | sub
				f[ss] = max(f[ss], fs+sum[sub])
			}
		}
		return f[m-1]
	}

	// 至多分成 k 组
	// https://atcoder.jp/contests/abc332/tasks/abc332_e
	subsubDP2 := func(a []int, k int) int {
		n := len(a)
		m := 1 << n
		// 预处理每个子集的子集和
		sum := make([]int, m)
		for i := range sum {
			for s := uint(i); s > 0; s &= s - 1 {
				v := a[bits.TrailingZeros(s)]
				sum[i] += v
			}
		}
		f := make([]int, m)
		for i := 1; i < m; i++ {
			f[i] = 1e18
		}
		for i := 0; i < k; i++ {
			for s := m - 1; s >= 0; s-- {
				t := m - 1 ^ s
				// 枚举补集的非空子集
				for sub := t; sub > 0; sub = (sub - 1) & t {
					f[s|sub] = min(f[s|sub], f[s]+sum[sub])
				}
				f[s] = 1e18
			}
		}
		return f[m-1]
	}

	// 记忆化写法
	subsubDPMemo := func(a []int) int {
		n := len(a)
		m := 1 << n
		// 预处理每个子集的子集和
		sum := make([]int, m)
		for i := range sum {
			for s := uint(i); s > 0; s &= s - 1 {
				sum[i] += a[bits.TrailingZeros(s)]
			}
		}
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, 1<<m)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}
		const inf int = 1e9 // 1e18
		var dfs func(p, set int) int
		dfs = func(p, set int) (res int) {
			if p == n {
				if set > 0 {
					return inf
				}
				return
			}
			dv := &memo[p][set]
			if *dv != -1 {
				return *dv
			}
			defer func() { *dv = res }()
			res = inf

			// 所有子集
			for sub, ok := set, true; ok; ok = sub != set {
				r := dfs(p+1, set^sub)
				res = min(res, r+sum[sub])
				sub = (sub - 1) & set
			}

			// 所有非空子集
			for sub := set; sub > 0; sub = (sub - 1) & set {

			}

			return
		}
		return dfs(0, 1<<m-1)
	}

	// 高维前缀和 SOS DP (Sum over Subsets)
	// 给一个全集 U，对 U 的所有子集 S，计算 S 的所有子集 T 之和（这个「和」不一定是加法，可以是其它的满足合并性质的统计量，例如 max 等）
	// https://codeforces.com/blog/entry/45223
	// Some SOS DP Insights https://codeforces.com/blog/entry/105247
	// 大量习题 https://blog.csdn.net/weixin_38686780/article/details/100109753
	//
	// LC2732 https://leetcode.cn/problems/find-a-good-subset-of-the-matrix/
	//    最简单的 SOSDP，判断是否有两个子集满足 x&y=0 且 idx[x]>=0 且 idx[y]>=0
	// https://codeforces.com/problemset/problem/1234/F
	//    求满足 ai&aj=0 的 ai|aj 的二进制 1 的个数的最大值
	//    由于 ai 的补集一定满足和 ai 的 & = 0
	//    所以转换成求每个 ai 的补集的 SOS，维护子集二进制 1 的个数的最大值
	// https://www.hackerearth.com/zh/problem/algorithm/special-pairs-5-3ee6b3fe-3d8a1606/
	//    求 ai&aj=0 的 (i,j) 对数，0<=ai<=1e6
	//    思路和上面类似，转换成求每个 ai 的补集的 SOS
	//    注：另一种解法是求 FWT(cnt)[0]
	// todo https://codeforces.com/problemset/problem/1208/F
	//    求 ai|(aj&ak) 的最大值，其中 i<j<k
	//    根据 a|b = (^a)&b + a
	//    问题变成 (^ai)&aj&ak + ai 的最大值
	// 转换成求集合中最大次大 https://atcoder.jp/contests/arc100/tasks/arc100_c
	// 求下标最大次大，且不需要在乎 k 的上限的写法 https://codeforces.com/problemset/problem/1554/B
	// https://codeforces.com/problemset/problem/165/E
	// 容斥 https://codeforces.com/problemset/problem/449/D
	// todo https://codeforces.com/problemset/problem/800/D
	//  https://codeforces.com/problemset/problem/383/E
	//  https://www.luogu.com.cn/problem/P6442
	// https://codeforces.com/problemset/problem/1523/D
	// 十进制 https://atcoder.jp/contests/arc136/tasks/arc136_d
	sosDP := func(a []int) []int {
		// 从子集转移的写法
		const mx = 20 // bits.Len(uint(max(a))
		f := make([]int, 1<<mx)
		for _, v := range a {
			f[v]++
		}
		for i := 0; i < mx; i++ {
			for s := 0; s < 1<<mx; s++ {
				s |= 1 << i
				// 将 s 的子集 s^1<<i 的统计量合并到 s 中
				f[s] += f[s^1<<i]
			}
		}

		{
			// 从超集转移的写法
			for i := 0; i < mx; i++ {
				for s := 1<<mx - 1; s >= 0; s-- {
					if s>>i&1 == 0 {
						f[s] += f[s|1<<i]
					}
				}
			}
		}

		{
			// 维护集合最大和次大的写法
			type pair struct{ fi, se int }
			f := make([]pair, 1<<mx)
			for i := 0; i < mx; i++ {
				for s := 0; s < 1<<mx; s++ {
					s |= 1 << i
					p, q := f[s], f[s^1<<i]
					if q.se > p.fi {
						f[s] = q
					} else if q.fi > p.fi {
						f[s] = pair{q.fi, p.fi}
					} else if q.fi > p.se {
						f[s].se = q.fi
					}
				}
			}
		}

		return f
	}

	/* 插头 DP（Plug DP）/ 轮廓线 DP（Broken Profile DP）
	轮廓线：已决策格子和未决策格子的分界线
	陈丹琦《基于连通性状态压缩的动态规划问题》
	训练指南 6.1
	https://oi-wiki.org/dp/plug/（花絮 https://zhuanlan.zhihu.com/p/133761303）
	https://cp-algorithms.com/dynamic_programming/profile-dynamics.html
	图解 https://www.luogu.com.cn/blog/GNAQ/solution-p5056
	An Introduction to Plug DP https://codeforces.com/blog/entry/90841
	todo 大量题目 https://www.luogu.com.cn/blog/efforts-will-pay-off/du-liu-dong-gui-cha-tou-dp
	模板题 https://www.luogu.com.cn/problem/P5056
	https://www.luogu.com.cn/problem/P3272
	https://www.luogu.com.cn/problem/P5074
	https://www.luogu.com.cn/problem/P1933
	https://codeforces.com/problemset/problem/903/F 2200
	LCP4 https://leetcode.cn/problems/broken-board-dominoes/
	*/
	plugDP := func(grids [][]byte) int {
		n, m := len(grids), len(grids[0])
		var endX, endY int
		g := make([][]bool, n)
		for i, row := range grids {
			g[i] = make([]bool, m)
			for j, b := range row {
				if b == '.' {
					g[i][j] = true
					endX, endY = i, j
				}
			}
		}

		// 四进制状态
		get := func(s, k int) int { return s >> (k << 1) & 3 }
		set := func(k, v int) int { return 1 << (k << 1) * v }

		ans := 0
		f := map[int]int{0: 1}
		for i, row := range g {
			tmp := f
			f = make(map[int]int, len(tmp))
			for s, dv := range tmp {
				f[s<<2] = dv // 轮廓线移动到当前行
			}
			for j, empty := range row {
				tmp := f
				f = make(map[int]int, len(tmp))
				for s, dv := range tmp {
					switch x, y := get(s, j), get(s, j+1); {
					case !empty: // 障碍格
						if x == 0 && y == 0 { // 空
							f[s] += dv
						}
					case x == 0 && y == 0: // ┌ 单独形成一对括号
						if j+1 < m && row[j+1] && i+1 < n && g[i+1][j] {
							f[s|set(j, 1)|set(j+1, 2)] += dv
						}
					case x == 0 && y > 0:
						if j+1 < m && row[j+1] { // └
							f[s] += dv
						}
						if i+1 < n && g[i+1][j] { // │
							f[s|set(j, y)^set(j+1, y)] += dv
						}
					case x > 0 && y == 0:
						if j+1 < m && row[j+1] { // ─
							f[s^set(j, x)|set(j+1, x)] += dv
						}
						if i+1 < n && g[i+1][j] { // ┐
							f[s] += dv
						}
					case x == 1 && y == 1: // ┘ 消去 x 和 y，并找到和 y 匹配的右括号，将其改成左括号
						// 注：这里和下边的 k 的位置可以事先预处理出来
						for k, c := j+2, 1; ; k++ {
							if t := get(s, k); t == 1 {
								c++
							} else if t == 2 {
								if c--; c == 0 {
									f[s^set(j, x)^set(j+1, y)^set(k, 3)] += dv // 将 2 改成 1 要异或 3
									break
								}
							}
						}
					case x == 2 && y == 2: // ┘ 消去 x 和 y，并找到和 x 匹配的左括号，将其改成右括号
						for k, c := j-1, 1; ; k-- {
							if t := get(s, k); t == 2 {
								c++
							} else if t == 1 {
								if c--; c == 0 {
									f[s^set(j, x)^set(j+1, y)^set(k, 3)] += dv // 将 1 改成 2 要异或 3
									break
								}
							}
						}
					case x == 2 && y == 1: // ┘ 消去右括号和左括号，连接两个插头
						f[s^set(j, x)^set(j+1, y)] += dv
					default: // ┘ x == 1 && y == 2
						// 此时封闭整个路径，这只应当发生在最后一个合法格子上
						if i == endX && j == endY {
							ans += dv
						}
					}
				}
			}
		}
		// 若需要取模则中间记得取模（若结果不大可以仅在循环结束时取模）
		return ans
	}

	/* 数位 DP
	v1.0 模板视频讲解 https://www.bilibili.com/video/BV1rS4y1s721/?t=19m36s
	v2.0 模板视频讲解 https://www.bilibili.com/video/BV1Fg4y1Q7wv/?t=31m28s
	https://zhuanlan.zhihu.com/p/348851463
	https://www.bilibili.com/video/BV1MT4y1376C

	入门题 https://atcoder.jp/contests/abc154/tasks/abc154_e
	      https://atcoder.jp/contests/dp/tasks/dp_s
	      https://codeforces.com/problemset/problem/1036/C
	二进制 1 的个数恰为 k 的数字个数 https://codeforces.com/problemset/problem/431/D https://www.acwing.com/problem/content/1083/
	是 m 的倍数且偶数位为 d 且奇数位不为 d 的数字个数 https://codeforces.com/problemset/problem/628/D
	所有数字均出现偶数次的数字个数 https://codeforces.com/problemset/problem/855/E
	相邻数字约束 SC09 https://www.luogu.com.cn/problem/P2657
	数位统计
	- 入门题 [2719. 统计整数数目](https://leetcode.cn/problems/count-of-integers/)（[题解](https://leetcode.cn/problems/count-of-integers/solution/shu-wei-dp-tong-yong-mo-ban-pythonjavacg-9tuc/)）
	- [788. 旋转数字](https://leetcode.cn/problems/rotated-digits/)（[题解](https://leetcode.cn/problems/rotated-digits/solution/by-endlesscheng-9b96/)）
	- [902. 最大为 N 的数字组合](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/)（[题解](https://leetcode.cn/problems/numbers-at-most-n-given-digit-set/solution/shu-wei-dp-tong-yong-mo-ban-xiang-xi-zhu-e5dg/)）1990
	    - 贪心：求 <= N 的最大数
	- [233. 数字 1 的个数](https://leetcode.cn/problems/number-of-digit-one/)（[题解](https://leetcode.cn/problems/number-of-digit-one/solution/by-endlesscheng-h9ua/)）
	- [面试题 17.06. 2 出现的次数](https://leetcode.cn/problems/number-of-2s-in-range-lcci/)（[题解](https://leetcode.cn/problems/number-of-2s-in-range-lcci/solution/by-endlesscheng-x4mf/)）
	    - 0~9 的个数 https://www.luogu.com.cn/problem/P2602
	    - http://acm.hdu.edu.cn/showproblem.php?pid=3555
	    - http://acm.hdu.edu.cn/showproblem.php?pid=2089
	- [600. 不含连续 1 的非负整数](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/)（[题解](https://leetcode.cn/problems/non-negative-integers-without-consecutive-ones/solution/by-endlesscheng-1egu/)）
	- [2376. 统计特殊整数](https://leetcode.cn/problems/count-special-integers/)（[题解](https://leetcode.cn/problems/count-special-integers/solution/shu-wei-dp-mo-ban-by-endlesscheng-xtgx/)）2120
	- [1012. 至少有 1 位重复的数字](https://leetcode.cn/problems/numbers-with-repeated-digits/)（[题解](https://leetcode.cn/problems/numbers-with-repeated-digits/solution/by-endlesscheng-c5vg/)）2230
	- [357. 统计各位数字都不同的数字个数](https://leetcode.cn/problems/count-numbers-with-unique-digits/)
	- [3007. 价值和小于等于 K 的最大数字](https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/) 2258
	    - 有更简单的做法，见 https://leetcode.cn/problems/maximum-number-that-sum-of-the-prices-is-less-than-or-equal-to-k/solutions/2603673/er-fen-da-an-shu-wei-dpwei-yun-suan-pyth-tkir/
	- [2827. 范围中美丽整数的数目](https://leetcode.cn/problems/number-of-beautiful-integers-in-the-range/) 2324
	    - 整除 https://atcoder.jp/contests/abc317/tasks/abc317_f
	    - [SCOI2009] windy 数 https://www.luogu.com.cn/problem/P2657
	- [2999. 统计强大整数的数目](https://leetcode.cn/problems/count-the-number-of-powerful-integers/) 2351
	- [2801. 统计范围内的步进数字数目](https://leetcode.cn/problems/count-stepping-numbers-in-range/) 2367
	- [1397. 找到所有好字符串](https://leetcode.cn/problems/find-all-good-strings/) 2667
		- https://www.luogu.com.cn/problem/P3193
		- https://atcoder.jp/contests/abc295/tasks/abc295_f
		- 与 AC 自动机结合 https://ac.nowcoder.com/acm/problem/20366
	- [1215. 步进数](https://leetcode.cn/problems/stepping-numbers/)（会员题）
	- [1067. 范围内的数字计数](https://leetcode.cn/problems/digit-count-in-range/)（会员题） *LC233
	- [1742. 盒子中小球的最大数量](https://leetcode.cn/problems/maximum-number-of-balls-in-a-box/) *非暴力做法 枚举数位和+DP
	- [2843. 统计对称整数的数目](https://leetcode.cn/problems/count-symmetric-integers/) *非暴力做法
	https://codeforces.com/contest/1710/problem/C
	数位和 digsum(n)|n https://www.luogu.com.cn/problem/P4127
	- https://atcoder.jp/contests/abc336/tasks/abc336_e
	- https://ac.nowcoder.com/acm/contest/28262/E
	- D 的倍数 https://atcoder.jp/contests/tdpc/tasks/tdpc_number
	数位和是最后一位的倍数 https://www.lanqiao.cn/problems/5891/learning/?contest_id=145
	数位乘积不超过 k https://atcoder.jp/contests/abc208/tasks/abc208_e
	恰好用到了 k 个不同数位 https://atcoder.jp/contests/abc194/tasks/abc194_f 2197=CF2373
	被每个非零数位都整除的数字个数 https://codeforces.com/problemset/problem/55/D 2500
	https://codeforces.com/gym/104337/problem/B 【妙】数位众数，把 freq 排序作为 key
	todo https://codeforces.com/problemset/problem/1245/F
	【转换】选两个不超过 U 的数，满足异或和为 target https://atcoder.jp/contests/arc133/tasks/arc133_d 2658
	https://lightoj.com/problem/investigation
	http://acm.hdu.edu.cn/showproblem.php?pid=4507
	http://acm.hdu.edu.cn/showproblem.php?pid=3886
	http://acm.hdu.edu.cn/showproblem.php?pid=6796
	注：一些第 k 小的题目需要与二分结合，或者用试填法（见后面的 kth666）
	todo 套题 https://www.luogu.com.cn/blog/s-r-f/oi-bi-ji-shuo-wei-dp-ge-ji-dui-shuo-wei-dp-di-yi-dian-li-xie
	todo 套题 https://codeforces.com/blog/entry/53960
	对称
	LC248 https://leetcode.cn/problems/strobogrammatic-number-iii/
	https://leetcode.cn/problems/confusing-number-ii/ 2077
	*/

	// 只做一次记忆化搜索的写法
	// TIPS：如果题目要求计算 < high 的方案数（high 是个字符串），
	//       可以在递归到 i=n 时，判断 limitHigh 是否为 true，如果是 true 则表示填入的数字等于 high，返回 0
	// 举例 https://atcoder.jp/contests/abc387/tasks/abc387_c
	// 代码 https://atcoder.jp/contests/abc387/submissions/61401082
	digitDP := func(low, high int, sumUpper int) int {
		lowS := strconv.Itoa(low) // 不加前导零
		highS := strconv.Itoa(high)
		n := len(highS)
		diffLH := n - len(lowS)
		memo := make([][]int, n)
		for i := range memo {
			memo[i] = make([]int, sumUpper+1)
			for j := range memo[i] {
				memo[i][j] = -1
			}
		}

		// 第一种写法（前导零不影响答案）
		var f func(int, int, bool, bool) int
		f = func(i, sum int, limitLow, limitHigh bool) (res int) {
			if i == n {
				// 不合法
				if sum > sumUpper {
					return 0
				}
				// 合法
				return 1
			}
			if !limitLow && !limitHigh {
				dv := &memo[i][sum]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}

			// 注：不要修改这里！如果对数位有其它限制，应当写在下面 for 循环中
			lo := 0
			if limitLow && i >= diffLH {
				lo = int(lowS[i-diffLH] - '0')
			}
			hi := 9
			if limitHigh {
				hi = int(highS[i] - '0')
			}

			for d := lo; d <= hi; d++ {
				res += f(i+1, sum+d, limitLow && d == lo, limitHigh && d == hi)
				res %= mod
			}
			return
		}
		//ans := f(0, 0, true, true)

		// 第二种写法（前导零影响答案）
		// 注意，仍然无需使用 isNum
		// 下面是计算每个数位都出现偶数次的方案数，这种情况下就要区分【前导零】和【数字中的零】了，
		// 前导零是不能统计其出现次数的，而数字中的零，例如 110220 中的 0 是要统计的
		var dfs func(int, int, bool, bool) int
		dfs = func(i, mask int, limitLow, limitHigh bool) (res int) {
			if i == n {
				// 注意：如果左边界 low=0，那么 0 是题目允许的吗？
				//if limitLow {
				//	return 0
				//}
				if mask > 0 { // 有数字出现奇数次
					return 0
				}
				return 1
			}
			if !limitLow && !limitHigh {
				dv := &memo[i][mask]
				if *dv >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}

			// 注：不要修改这里！如果对数位有其它限制，应当写在下面 for 循环中
			lo := 0
			if limitLow && i >= diffLH {
				lo = int(lowS[i-diffLH] - '0')
			}
			hi := 9
			if limitHigh {
				hi = int(highS[i] - '0')
			}

			d := lo
			// 这样就可以表示当前是否处于前导零状态了，无需 isNum 参数
			if limitLow && i < diffLH {
				// 什么也不填
				res = dfs(i+1, mask, true, false)
				d++ // d = 1
			}
			for ; d <= hi; d++ {
				res += dfs(i+1, mask^1<<d, limitLow && d == lo, limitHigh && d == hi)
				res %= mod
			}
			return
		}
		ans := dfs(0, 0, true, true)
		return ans
	}

	//（旧版写法）做两次记忆化搜索（或者题目只需要考虑上界）
	digitDP2 := func(low, high string, sumUpper int) int {
		// 返回 <=s 的符合要求的字符串数目
		// TIPS: 某些情况下思考补集会更加容易，即求不符合要求的字符串数目
		calc := func(s string) int {
			// 注：如果参数太多可以用 map + struct
			memo := make([][]int, len(s))
			for i := range memo {
				memo[i] = make([]int, sumUpper+1)
				for j := range memo[i] {
					memo[i][j] = -1
				}
			}

			// 第一种写法（前导零不影响答案）
			var f func(int, int, bool) int
			f = func(i int, sum int, isLimit bool) (res int) {
				if i == len(s) {
					return 1
				} // sum
				if !isLimit {
					dv := &memo[i][sum]
					if *dv >= 0 {
						return *dv
					} // *dv + sum*int(math.Pow10(len(s)-i))
					defer func() { *dv = res }()
				}
				up := 9 // 25
				if isLimit {
					up = int(s[i] - '0') // 'a'
				}
				for d := 0; d <= up; d++ {
					tmp := sum

					cnt := f(i+1, tmp, isLimit && d == up)
					res = (res + cnt) % mod
				}
				return
			}
			//res := f(0, 0, true)

			// 第二种写法（前导零影响答案）
			// 对于需要判断/禁止前导零的情况，可以加一个额外的维度 isNum，表示已经填入了数字（没有前导零的合法状态），最后 p=n 的时候可以根据情况返回 1 或者 0
			// 例如 https://codeforces.com/contest/855/submission/125651587
			// 以下代码以 https://www.luogu.com.cn/problem/P2657 为例
			var dfs func(int, int, bool, bool) int
			dfs = func(i int, pre int, isLimit, isNum bool) (res int) {
				if i == len(s) {
					if !isNum {
						return 0
					}
					return 1
				}
				if !isLimit && isNum {
					dv := &memo[i][pre]
					if *dv >= 0 {
						return *dv
					}
					defer func() { *dv = res }()
				}
				if !isNum {
					res += dfs(i+1, pre, false, false)
				}
				up := 9 // 25
				if isLimit {
					up = int(s[i] - '0') // - 'a'
				}
				d := 0
				if !isNum {
					d = 1
				}
				for ; d <= up; d++ {
					//if ... {
					//	continue
					//}
					res += dfs(i+1, pre+d /* nextD */, isLimit && d == up, true)
					res %= mod
				}
				return
			}
			res := dfs(0, 0, true, false)
			return res
		}
		ansHigh := calc(high) // 上界
		ansLow := calc(low)   // 下界（注意下面单独特判 low）
		ans := ansHigh - ansLow
		ans = (ans%mod + mod) % mod

		// low 是否算上
		lowIsValid := true
		for i := 1; i < len(low); i++ {
			pre, cur := int(low[i-1]), int(low[i])
			if abs(pre-cur) != 1 { // 不合法
				lowIsValid = false
				break
			}
		}
		if lowIsValid {
			ans++
		}
		ans = (ans%mod + mod) % mod
		return ans
	}

	// 若需要计算的不是合法数字个数，而是合法数字之和，则需要在计算时考虑单个数位的贡献
	// 以下代码以 https://codeforces.com/problemset/problem/1073/E 为例
	calcSum := func(s string, k int) int {
		n := len(s)
		type pair struct{ cnt, sum int }
		memo := make([][1 << 10]pair, n)
		for i := range memo {
			for j := range memo[i] {
				memo[i][j] = pair{-1, -1}
			}
		}
		var f func(int, uint16, bool, bool) pair
		f = func(p int, mask uint16, limitUp, fill bool) (res pair) {
			if p == n {
				if !fill {
					return
				}
				return pair{1, 0}
			}
			if !limitUp && fill {
				dv := &memo[p][mask]
				if dv.cnt >= 0 {
					return *dv
				}
				defer func() { *dv = res }()
			}
			up := 9
			if limitUp {
				up = int(s[p] & 15)
			}
			for d := 0; d <= up; d++ {
				tmp := mask
				if fill || d > 0 {
					tmp |= 1 << d
				}
				if bits.OnesCount16(tmp) <= k {
					pr := f(p+1, tmp, limitUp && d == up, fill || d > 0)
					res.cnt = (res.cnt + pr.cnt) % mod
					res.sum = (res.sum + int(math.Pow10(n-1-p))%mod*pr.cnt%mod*int(d) + pr.sum) % mod
				}
			}
			return
		}
		return f(0, 0, true, false).sum
	}

	// 试填法
	// 第 k 个包含 3 个连续的 6 的数 https://www.acwing.com/problem/content/312/
	kth666 := func(k int) (ans []byte) {
		// f[i][3] 表示由 i 位数字构成的魔鬼数的个数
		// f[i][j] (j<3) 表示 i 位数字构成的、开头有连续 j 个 6 的非魔鬼数的个数
		const mx = 30  // 长度上限
		const cont = 3 // 连续 3 个数才算符合要求
		f := [mx][cont + 1]int{}
		f[0][0] = 1
		for i := 1; i < mx; i++ {
			for j := 0; j < cont; j++ {
				f[i][0] += f[i-1][j] * 9 // 开头无 6，直接转移（0-9 中除去 6 共 9 个数）
				f[i][j+1] = f[i-1][j]    // 开头有 j+1 个 6，下一个有 j 个 6
			}
			f[i][cont] += f[i-1][cont] * 10
		}

		const tarDigit byte = '6'
		n := 1
		for f[n][cont] < k {
			n++
		}
		has := 0
		for i := 1; i <= n; i++ {
			for digit := byte('0'); digit <= '9'; digit++ { // 试填当前位
				need := cont
				if has == cont {
					need = 0
				} else if digit == tarDigit {
					need = cont - 1 - has
				}
				sum := 0
				for j := need; j <= cont; j++ {
					sum += f[n-i][j]
				}
				if sum >= k { // 填入
					ans = append(ans, digit)
					if has < cont {
						if digit == tarDigit {
							has++
						} else {
							has = 0
						}
					}
					break
				}
				k -= sum
			}
		}
		return
	}

	/* 倍增优化 DP
	https://codeforces.com/problemset/problem/1175/E 模板题
	https://codeforces.com/problemset/problem/1516/D
	https://atcoder.jp/contests/arc060/tasks/arc060_c
	https://www.luogu.com.cn/problem/P1081 开车旅行
	https://www.luogu.com.cn/problem/P3147 合并数字
	计算重复 https://www.acwing.com/problem/content/296/
	*/
	binaryLifting := func(segs, qs []struct{ l, r int }) []int {
		// 以 CF1175E 为例
		const mx = 19
		f := make([][mx]int, 5e5+1)
		for _, s := range segs {
			l, r := s.l, s.r
			f[l][0] = max(f[l][0], r)
		}
		// 前缀最大值（最右）
		for i := 1; i < len(f); i++ {
			f[i][0] = max(f[i][0], f[i-1][0])
		}
		// 倍增
		for i := 0; i+1 < mx; i++ {
			for p := range f {
				f[p][i+1] = f[f[p][i]][i]
			}
		}

		ans := make([]int, len(qs))
		for qi, q := range qs {
			l, r := q.l, q.r
			res := 0
			for i := mx - 1; i >= 0; i-- {
				if f[l][i] < r {
					l = f[l][i]
					res |= 1 << i
				}
			}
			if f[l][0] >= r {
				ans[qi] = res + 1
			} else {
				ans[qi] = -1
			}
		}
		return ans
	}

	/* 数据结构优化 DP
	https://codeforces.com/problemset?order=BY_RATING_ASC&tags=data+structures%2Cdp
	入门：数组优化 https://codeforces.com/contest/1842/problem/C
	             LC2713 https://leetcode.cn/problems/maximum-strictly-increasing-cells-in-a-matrix/
	线段树 LC2407 https://leetcode.cn/problems/longest-increasing-subsequence-ii/ 2280
	变量优化 O(n) LC2746 https://leetcode.cn/problems/decremental-string-concatenation/
	长为 k 的上升子序列个数 https://codeforces.com/problemset/problem/597/C
	Lazy 线段树 https://atcoder.jp/contests/dp/tasks/dp_w
	值域线段树 https://atcoder.jp/contests/abc339/tasks/abc339_e
	todo https://codeforces.com/problemset/problem/1667/B
	https://atcoder.jp/contests/arc073/tasks/arc073_d https://www.luogu.com.cn/problem/T190609?contestId=48376 https://www.luogu.com.cn/blog/abruce-home/ti-xie-nao-zhong
	todo https://codeforces.com/problemset/problem/66/E
	todo https://codeforces.com/problemset/problem/1788/E
	https://codeforces.com/contest/833/problem/B 2200
	https://codeforces.com/problemset/problem/1295/E 2200
	https://codeforces.com/contest/1842/problem/E 2300
	https://codeforces.com/problemset/problem/1609/E 2400 分治 线段树

	维护最大次大
	LC3122 https://leetcode.cn/problems/minimum-number-of-operations-to-satisfy-conditions/ 1905
	https://codeforces.com/problemset/problem/264/C 2000

	滑动窗口优化 DP
	https://codeforces.com/problemset/problem/985/E 2100
	- 把 a 从小到大排序后，如果分组方案存在，那么一定有一个分组方案是，同一组的数在 a 中是连续的
	- 证明：因为交换两个不同组的数，一定不会变得更优

	前缀和优化 DP
	优化空间时，可以直接把前缀和保存到 f 数组上，然后倒序遍历，计算实际的 f
	- 例如 https://leetcode.cn/problems/count-the-number-of-inversions/
	LC1977 https://leetcode.cn/problems/number-of-ways-to-separate-numbers/
	- [1997. 访问完所有房间的第一天](https://leetcode.cn/problems/first-day-where-you-have-been-in-all-the-rooms/) 2260
	     - https://codeforces.com/problemset/problem/1552/F 2200 数据加强
	LC2327 https://leetcode.cn/problems/number-of-people-aware-of-a-secret/
	LC2463 https://leetcode.cn/problems/minimum-total-distance-traveled/
	LC2478 https://leetcode.cn/problems/number-of-beautiful-partitions/
	https://codeforces.com/problemset/problem/46/E 1900 前缀最大值/后缀最大值
	https://codeforces.com/problemset/problem/479/E 1900
	- https://atcoder.jp/contests/abc253/tasks/abc253_e
	https://atcoder.jp/contests/abc248/tasks/abc248_c
	终极 BOSS https://atcoder.jp/contests/diverta2019/tasks/diverta2019_e

	其他
	https://codeforces.com/problemset/problem/1863/F 2600

	数论分块优化 DP
	https://codeforces.com/problemset/problem/1603/C

	动态 DP
	https://oi-wiki.org/dp/dynamic/
	线段树维护转移矩阵 https://www.cnblogs.com/Miracevin/p/9124511.html
	1. 先不考虑修改，不考虑区间，直接列出整个区间的 DP 转移。
	2. 列出转移矩阵。由于有很多修改操作，我们将数据集中在一起处理，还可以利用矩阵结合律，并且区间比较好提取（找一段矩阵就好了），修改也方便。
	3. 线段树维护矩阵。对于修改，我们就是在矩阵上进行修改。对于不同的题目，我们要用不同的修改方式和记录手段，但都是用线段树维护对应区间内的矩阵信息。
	   如矩阵乘积，矩阵和等等。线段树的区间操作可以应对区间修改问题。
	todo https://www.cnblogs.com/alex-wei/p/DP_Involution.html
	https://codeforces.com/problemset/problem/1380/F
	https://codeforces.com/problemset/problem/718/C
	https://codeforces.com/problemset/problem/750/E
	https://codeforces.com/problemset/problem/1149/B
	https://codeforces.com/problemset/problem/2042/F 2600
	todo 与树剖结合 https://www.luogu.com.cn/problem/P4719 https://www.luogu.com.cn/problem/P4751
	*/

	// 单调队列优化 DP
	// 见 monotone_queue.go

	// 斜率优化 / 凸包优化 (Convex Hull Trick, CHT)
	//
	// 若状态转移方程具有类似于 f[i] = min{f[j]-a[i]*b[j]}, j<i 的形式，方程中包含一个 i 和 j 的乘积项，且序列 a 和 b 均单调递增
	// 若将 (b[j],f[j]) 看作二维平面上的点，则 f[i] 就是所有斜率为 a[i] 且过其中一点的直线中，与 y 轴的最小截距
	// 我们可以用一个单调队列来维护 (b[j],f[j]) 的相邻点所构成的下凸包
	// 对于斜率 a[i]，我们需要在队列中寻找一个位置 k，其左侧斜率小于 a[i]，右侧斜率大于 a[i]，此时经过点 (b[k],f[k]) 能取到最小截距
	//
	// 具体到实现，设两转移来源的下标为 j 和 k，若 k < j 且 f[k]-a[i]*b[k] < f[j]-a[i]*b[j]
	// 则有 (f[j]-f[k])/(b[j]-b[k]) > a[i]
	// 据此式，用单调队列来维护斜率（下凸包）
	// 转移前，在单调队列中找到斜率 a[i] 的对应位置，然后代入转移方程，求出 f[i]
	// 转移后，将点 (b[i],f[i]) 加入单调队列中
	//
	// https://oi-wiki.org/dp/opt/slope/
	// https://cp-algorithms.com/geometry/convex_hull_trick.html
	// todo https://www.bilibili.com/video/BV178411W7Aj/
	// https://www.luogu.com.cn/blog/ChenXingLing/post-xue-xi-bi-ji-dong-tai-gui-hua-xie-shuai-you-hua-dp-chao-yang-x
	// https://www.luogu.com.cn/blog/ningago-lsh/xie-lv-you-hua-dp
	// https://blog.csdn.net/weixin_43914593/article/details/105560357 算法竞赛专题解析（12）：DP优化(2)--斜率(凸壳)优化
	// https://zhuanlan.zhihu.com/p/558522044
	// https://zhuanlan.zhihu.com/p/363772434
	// https://codeforces.com/blog/entry/63823
	//
	// todo - [LCP 59. 搭桥过河](https://leetcode.cn/problems/NfY1m5/)
	// - [2263. 数组变为有序的最小操作次数](https://leetcode.cn/problems/make-array-non-decreasing-or-non-increasing/)（会员题）
	// https://codeforces.com/problemset/problem/319/C 2100
	// https://www.luogu.com.cn/problem/P2365 https://www.luogu.com.cn/problem/P5785 http://poj.org/problem?id=1180
	// todo https://atcoder.jp/contests/dp/tasks/dp_z
	// todo https://www.luogu.com.cn/problem/P2900
	//  https://www.luogu.com.cn/problem/P3195 https://loj.ac/p/10188
	//  http://poj.org/problem?id=3709
	//  https://codeforces.com/problemset/problem/311/B 2400
	//  https://codeforces.com/problemset/problem/1715/E 2400
	//  https://codeforces.com/problemset/problem/631/E 2600
	//  结合李超线段树 https://codeforces.com/contest/1175/problem/G 3000
	cht := func(a, b []int) int {
		n := len(a)
		f := make([]int, n)
		// 计算两点间的斜率，若分子分母均在 32 位整数范围内，可以去掉浮点，改用乘法（或者用 lessPair）
		slope := func(i, j int) float64 {
			if b[i] == b[j] { // 若保证不相等则去掉
				if f[j] > f[i] {
					return 1e99
				}
				return -1e99
			}
			return float64(f[j]-f[i]) / float64(b[j]-b[i])
		}
		q := []int{0}
		for i := 1; i < n; i++ {
			k := a[i]
			// 若斜率 k 随 i 单调递增，则可以直接将单调队列中小于 k 的斜率弹出
			for len(q) > 1 && slope(q[0], q[1]) < float64(k) {
				q = q[1:]
			}
			// ……之后，队首 q[0] 就是最优决策的下标
			j := q[0]
			// ……否则，需要在单调队列中二分得到最优决策的位置
			j = sort.Search(len(q)-1, func(j int) bool { return slope(j, j+1) > float64(k) })

			// 转移
			f[i] = f[j] - a[i]*b[j]

			// 然后，将点 (b[i],f[i]) 加入单调队列中
			for len(q) > 1 && slope(q[len(q)-1], i) < slope(q[len(q)-2], q[len(q)-1]) {
				q = q[:len(q)-1]
			}
			q = append(q, i)
		}
		return f[n-1]
	}

	// WQS 二分 / 凸优化 DP / 带权二分 / Alien Trick / Alien DP / Monge グラフ上のd-辺最短路長を計算するアルゴリズム 
	// 原文：《浅析一类二分方法》
	// 把最多选 k 个物品的问题（时间复杂度高）转换成选任意个物品的问题（更容易解决，时间复杂度低）
	// 要求满足性质：k 越大，额外产生的收益是单调递减的
	// 具体请看下面的代码注释
	// https://zhuanlan.zhihu.com/p/340514421
	// https://www.cnblogs.com/alex-wei/p/DP_Involution.html
	// https://www.cnblogs.com/CreeperLKF/p/9045491.html
	// https://taodaling.github.io/blog/2020/07/31/WQS%E4%BA%8C%E5%88%86/
	// https://www.luogu.com.cn/blog/daniu/wqs-er-fen
	// https://www.luogu.com.cn/blog/Flying2018/wqs-er-fen-min-ke-fu-si-ji-hu-xue-xi-bi-ji
	// https://www.luogu.com.cn/blog/juruoforever/wqs-er-fen-qian-xi
	// https://noshi91.github.io/algorithm-encyclopedia/d-edge-shortest-path-monge
	//
	// 题单 https://www.luogu.com.cn/training/3495#problems
	// 种树（打家劫舍）https://www.luogu.com.cn/problem/P1484
	// LC188 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
	// LC2209 https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/
	// todo https://codeforces.com/problemset/problem/125/E 2400 单度限制最小生成树（恰好）
	//  https://codeforces.com/problemset/problem/321/E 2600
	//  https://codeforces.com/problemset/problem/739/E 3000 也可以费用流
	//  结合四边形不等式 IOI00 邮局 https://www.luogu.com.cn/problem/P4767 
	//  - https://www.luogu.com.cn/problem/P6246
	//  https://www.luogu.com.cn/problem/P5308
	//  IOI16 aliens https://www.luogu.com.cn/problem/P5896
	//  https://www.luogu.com.cn/problem/U72600
	wqs := func(a []int, k int) (ans int) {
		// 代码分为两部分：dpWithFee 和 sort.Search
		// 前者是需要实现的 DP，后者是固定的模板，一般只需要注意二分上界

		// dpWithFee 是一个更容易解决的 DP：每次产生一次「选择」的时候（例如完成一次买卖），额外有手续费 fee，但没有至多 k 次的限制了
		// 返回最大收益，以及在收益最大的前提下，「选择」次数的最大值
		// LC714 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
		dpWithFee := func(fee int) (res, cnt int) {
			f0, f1 := 0, math.MinInt/2
			cnt0, cnt1 := 0, 0
			for _, p := range a {
				if f0-p >= f1 { // 取等号，让交易次数尽量多
					f1 = f0 - p
					cnt1 = cnt0
				}
				if f1+p-fee >= f0 { // 取等号，让交易次数尽量多
					f0 = f1 + p - fee
					cnt0 = cnt1 + 1 // 卖出才算完整交易
				}
			}
			return f0, cnt0
		}

		// 下面是 WQS 二分模板
		// 以 LC188 为例 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-iv/
		// 二分交易「手续费」fee，做一个 LC714 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-with-transaction-fee/
		// 那么 fee 越小，交易次数越多；fee 越大，交易次数越小
		// 如果交易次数小于 k，说明 fee 取大了，反之 fee 取小了
		// 如果某个 fee 对应着恰好 k 次交易，就得到了正确答案
		//
		// fee 最大为 slices.Max(a)，此时可以保证收益为 0（但仍然可以完成交易，比如卖出价格-买入价格-fee=0）
		// +1 可以保证至少触发一次 ans 赋值
		sort.Search(slices.Max(a)+1, func(fee int) bool {
			res, cnt := dpWithFee(fee)
			if cnt >= k { // 至少 k 次交易
				ans = res + k*fee // 直接算，因为最终一定会二分到恰好 k 次交易
				return false
			}
			if fee == 0 { // 说明无论如何，交易次数都小于 k
				ans = res // fee 为 0 时，res 最大，也就是答案
			}
			return true
		})

		// 下面是另一个例子 https://www.luogu.com.cn/problem/P1484
		// 这题的 dpWithFee 就只需要实现一个打家劫舍的 DP
		dpWithFee = func(fee int) (res, cnt int) {
			var f0, f1, cnt0, cnt1 int
			for _, v := range a {
				v -= fee
				if f1 > f0+v { // 不选
					f0 = f1
					cnt0 = cnt1
				} else { // 选（相等时也选）
					f0, f1 = f1, f0+v
					cnt0, cnt1 = cnt1, cnt0+1
				}
			}
			return f1, cnt1
		}
		// WQS 模板（代码和上面一样）
		// fee 最大为 slices.Max(a)，此时可以保证收益为 0
		sort.Search(slices.Max(a)+1, func(fee int) bool {
			res, cnt := dpWithFee(fee)
			if cnt >= k {
				ans = res + k*fee
				return false
			}
			if fee == 0 {
				ans = res
			}
			return true
		})

		return
	}

	// 四边形不等式优化 Knuth's Optimization / Knuth-Yao speedup
	// 满足该性质的 DAG 叫做 Monge 图
	// https://oi-wiki.org/dp/opt/quadrangle/
	// https://jeffreyxiao.me/blog/knuths-optimization
	// todo https://blog.csdn.net/weixin_43914593/article/details/105150937 算法竞赛专题解析（10）：DP优化(1)--四边形不等式
	//      决策单调性优化讲解 https://www.luogu.com.cn/blog/83547/zong-dong-tai-gui-hua-di-ben-zhi-kan-si-bian-xing-fou-deng-shi-you-hua
	// 扔蛋问题 LC887 https://leetcode.cn/problems/super-egg-drop/

	// CDQ 分治优化 DP
	// todo https://oi-wiki.org/misc/cdq-divide/
	//  推荐 https://blog.nowcoder.net/n/f44d4aada5a24f619442dd6ddffa7320
	//  推荐 https://zhuanlan.zhihu.com/p/332996578
	//  https://www.bilibili.com/video/BV1mC4y1s7ic
	//  [学习笔记]CDQ分治和整体二分 https://www.luogu.com.cn/blog/Owencodeisking/post-xue-xi-bi-ji-cdq-fen-zhi-hu-zheng-ti-er-fen
	//  https://www.luogu.com.cn/blog/ljc20020730/cdq-fen-zhi-xue-xi-bi-ji
	//  动态逆序对 https://www.luogu.com.cn/problem/P3157 https://www.luogu.com.cn/problem/UVA11990
	//  CDQ 优化 DP https://www.luogu.com.cn/problem/P2487
	//  https://codeforces.com/problemset/problem/848/C
	//  数据结构 https://codeforces.com/problemset/problem/1045/G https://codeforces.com/problemset/problem/762/E

	/* 树形 DP
	思考方向：
	每个节点需要计算的信息，是否只取决于邻居？
	如果不能，如何把子树的信息归纳到邻居上？

	一般是从自底向上计算的，也就是根据子树返回值来计算父节点的值
	也有自顶向下的写法，见后面

	https://blog.csdn.net/weixin_43914593/article/details/107145592
	https://codeforces.com/blog/entry/20935
	https://codeforces.com/blog/entry/63257

	LC337 https://leetcode.cn/problems/house-robber-iii/
	LC1372 https://leetcode.cn/problems/longest-zigzag-path-in-a-binary-tree/
	LC2925 https://leetcode.cn/problems/maximum-score-after-applying-operations-on-a-tree/ 1940
	LC2920 https://leetcode.cn/problems/maximum-points-after-collecting-coins-from-all-nodes/ 2351
	简单 DFS https://leetcode.cn/problems/find-number-of-coins-to-place-in-tree-nodes/
	https://atcoder.jp/contests/abc259/tasks/abc259_f
	https://atcoder.jp/contests/abc239/tasks/abc239_e

	CF tag https://codeforces.com/problemset?order=BY_RATING_ASC&tags=dp%2Ctrees
	todo 题单 https://ac.nowcoder.com/acm/problem/collection/807
	     题单 https://ac.nowcoder.com/acm/problem/collection/809
	https://codeforces.com/problemset/problem/982/C 1500
	https://codeforces.com/problemset/problem/369/C 1600
	https://codeforces.com/problemset/problem/1689/C 1600 二叉树
	https://codeforces.com/problemset/problem/743/D 1800
	https://codeforces.com/problemset/problem/1083/A 1800
	https://codeforces.com/problemset/problem/1833/G 1800
	https://codeforces.com/problemset/problem/1739/D 1900 二分答案
	https://codeforces.com/problemset/problem/1926/G 1900
	https://codeforces.com/problemset/problem/855/C 2000
	https://codeforces.com/problemset/problem/461/B 2000 如何定义状态
	https://codeforces.com/problemset/problem/1923/E 2000 两端点颜色相同，中间节点颜色不等于两端点
	- 对于每个点，统计往上走能到哪些同色点（中途不能经过同色点）
	https://codeforces.com/problemset/problem/1988/D 2000
	https://codeforces.com/problemset/problem/734/E 2100 巧妙的转换
	https://codeforces.com/problemset/problem/538/E 2200 **如何转移
	https://codeforces.com/problemset/problem/1220/E 2200 可以重复走
	https://codeforces.com/problemset/problem/1249/F 2200 好题
	https://codeforces.com/problemset/problem/1292/C 2300
	https://codeforces.com/problemset/problem/1453/E 2300 好题
	https://codeforces.com/problemset/problem/1059/E 2400 取往上冲的最高的点（子树）
	https://ac.nowcoder.com/acm/contest/63585/d

	自顶向下
	LCP64 https://leetcode.cn/problems/U7WvvU/
	- 题解 https://leetcode.cn/problems/U7WvvU/solution/shu-xing-dp-by-endlesscheng-isuo/
	*/

	// 树的直径（两遍 DFS 求法另见 graph_tree.go 中的 diameter）
	// LC1245 https://leetcode.cn/problems/tree-diameter/
	// 带权版本 https://atcoder.jp/contests/abc361/tasks/abc361_e
	// 变形 LC2246 https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/
	// 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/
	// - 变形 LC2385 https://leetcode.cn/problems/amount-of-time-for-binary-tree-to-be-infected/
	// https://codeforces.com/problemset/problem/1881/F 1700
	// https://codeforces.com/problemset/problem/1238/F 2200
	// https://codeforces.com/problemset/problem/1042/F 2400 贪心
	// 虚树直径 https://www.luogu.com.cn/problem/P4103
	diameter := func(st int, g [][]int) (diameter int) {
		var dfs func(int, int) int
		dfs = func(v, fa int) (maxL int) {
			for _, w := range g[v] {
				if w != fa {
					subL := dfs(w, v) + 1 // wt
					diameter = max(diameter, maxL+subL)
					maxL = max(maxL, subL)
				}
			}
			return
		}
		dfs(st, -1)
		return
	}

	// 树的直径及其个数
	// http://acm.hdu.edu.cn/showproblem.php?pid=3534
	// https://ac.nowcoder.com/acm/contest/view-submission?submissionId=45988692
	countDiameter := func(st int, g [][]int) (diameter, diameterCnt int) {
		var dfs func(int, int) (int, int)
		dfs = func(v, fa int) (int, int) {
			mxDep, cnt := 0, 1
			for _, w := range g[v] {
				if w != fa {
					d, c := dfs(w, v)
					if l := mxDep + d; l > diameter {
						diameter, diameterCnt = l, cnt*c
					} else if l == diameter {
						diameterCnt += cnt * c
					}
					if d > mxDep {
						mxDep, cnt = d, c
					} else if d == mxDep {
						cnt += c
					}
				}
			}
			return mxDep + 1, cnt
		}
		dfs(st, -1)
		return
	}

	// 统计树上所有简单路径长度及其个数 O(n^3)
	// 返回一个长为 n 的数组 ans，ans[i] 表示长为 i 的简单路径个数
	// O(n^2) 见 https://github.com/hqztrue/LeetCodeSolutions/blob/master/1601-1700/1617.%20Count%20Subtrees%20With%20Max%20Distance%20Between%20Cities.pdf
	// LC1617 https://leetcode.cn/problems/count-subtrees-with-max-distance-between-cities/
	countPath := func(g [][]int) []int {
		n := len(g)
		// 计算树上任意两点的距离
		dis := make([][]int, n)
		for i := range dis {
			// 计算 i 到其余点的距离
			dis[i] = make([]int, n)
			var dfs func(int, int)
			dfs = func(x, fa int) {
				for _, y := range g[x] {
					if y != fa {
						dis[i][y] = dis[i][x] + 1 // 自顶向下
						dfs(y, x)
					}
				}
			}
			dfs(i, -1)
		}

		ans := make([]int, n)
		ans[0] = n
		for i, di := range dis {
			for j := i + 1; j < n; j++ {
				dj := dis[j]
				dij := di[j]
				var dfs func(int, int) int
				dfs = func(x, fa int) int {
					// 能递归到这，说明 x 可以选
					cnt := 1 // 选 x
					for _, y := range g[x] {
						if y != fa &&
							(di[y] < dij || di[y] == dij && y > j) &&
							(dj[y] < dij || dj[y] == dij && y > i) { // 满足这些条件就可以选
							cnt *= dfs(y, x) // 每棵子树互相独立，采用乘法原理
						}
					}
					if di[x]+dj[x] > dij { // x 是可选点
						cnt++ // 不选 x
					}
					return cnt
				}
				ans[dij] += dfs(i, -1)
			}
		}
		return ans
	}

	// 树的直径及在直径上的节点个数
	// https://ac.nowcoder.com/acm/contest/view-submission?submissionId=45987468
	// 注意这里的 cnt 初始化与 countDiameter 的不同之处
	countVerticesOnDiameter := func(st int, g [][]int) (diameter, verticesCnt int) {
		var f func(int, int) (int, int)
		f = func(v, fa int) (int, int) {
			mxDep, cnt := 0, 0
			for _, w := range g[v] {
				if w != fa {
					d, c := f(w, v)
					if l := mxDep + d; l > diameter {
						diameter, verticesCnt = l, cnt+c+1 // 最长的链 + 当前链 + 当前节点
					} else if l == diameter {
						verticesCnt += c
					}
					if d > mxDep {
						mxDep, cnt = d, c
					} else if d == mxDep {
						cnt += c
					}
				}
			}
			return mxDep + 1, cnt + 1
		}
		f(st, -1)
		return
	}

	// 最大路径和 最大路径点权和
	// 变形 LC2538 https://leetcode.cn/problems/difference-between-maximum-and-minimum-price-sum/
	// https://codeforces.com/problemset/problem/633/F 2600 两条不相交路径的最大点权和
	maxPathSum := func(st int, g [][]int, a []int) (ans int) {
		// 点权
		var f func(int, int) int
		f = func(v, fa int) int {
			val := a[v]
			ans = max(ans, val)
			maxS := val // 从下面到 v 的最大【链】点权和
			for _, w := range g[v] {
				if w != fa {
					s := f(w, v)
					ans = max(ans, maxS+s)
					maxS = max(maxS, s+val)
				}
			}
			return maxS
		}
		f(st, -1)

		{
			// 边权
			type nb struct{ to, wt int }
			var g [][]nb // read...
			var f func(int, int) int
			f = func(v, fa int) int {
				maxS := 0
				for _, e := range g[v] {
					w := e.to
					if w != fa {
						s := f(w, v) + e.wt
						ans = max(ans, maxS+s)
						maxS = max(maxS, s)
					}
				}
				return maxS
			}
			f(st, -1)
		}

		{
			// 点权+边权
			type nb struct{ to, wt int }
			var g [][]nb // read...
			var f func(int, int) int
			f = func(v, fa int) int {
				val := a[v]
				ans = max(ans, val)
				maxS := val
				for _, e := range g[v] {
					w := e.to
					if w != fa {
						s := f(w, v) + e.wt
						ans = max(ans, maxS+s)
						maxS = max(maxS, s+val)
					}
				}
				return maxS
			}
			f(st, -1)
		}

		return
	}

	// 树上最大独立集
	// 返回最大点权和（最大独立集的情形即所有点权均为一）
	// 每个点有选和不选两种决策，接受子树转移时，选的决策只能加上不选子树，而不选的决策可以加上 max{不选子树, 选子树}
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	// https://stackoverflow.com/questions/13544240/algorithm-to-find-max-independent-set-in-a-tree
	// 经典题：没有上司的舞会 LC337 https://leetcode.cn/problems/house-robber-iii/ https://www.luogu.com.cn/problem/P1352 https://ac.nowcoder.com/acm/problem/51178
	// - EXTRA：可以修改树上的点权 https://www.luogu.com.cn/problem/P4719 https://www.luogu.com.cn/problem/P4751
	// 边独立集 LC2378 https://leetcode.cn/problems/choose-edges-to-maximize-score-in-a-tree/
	// 变形 LC2646 https://leetcode.cn/problems/minimize-the-total-price-of-the-trips/
	// 方案是否唯一 Tehran06，紫书例题 9-13，UVa 1220 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=247&page=show_problem&problem=3661
	// 变形：同一层至多选一个节点（蓝桥杯）
	// 相似，但不完全是 https://codeforces.com/problemset/problem/1988/D
	maxIndependentSetOfTree := func(n int, g [][]int, a []int) int { // 无根树
		var f func(int, int) (int, int)
		f = func(v, fa int) (notChosen, chosen int) {
			chosen = a[v] // 1
			for _, w := range g[v] {
				if w != fa {
					nc, c := f(w, v)
					notChosen += max(nc, c)
					chosen += nc
				}
			}
			return
		}
		return max(f(0, -1))
	}

	// 树上最小顶点覆盖（每条边至少有一个端点被覆盖，也就是在已选的点集中）
	// 代码和树上最大独立集类似
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	// 经典题：战略游戏 https://www.luogu.com.cn/problem/P2016
	// 训练指南第一章例题 30，UVa10859 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1800
	// - 求最小顶点覆盖，以及所有最小顶点覆盖中，两端点都被覆盖的边的最大个数
	// 构造 https://codeforces.com/problemset/problem/959/C
	minVertexCoverOfTree := func(n int, g [][]int, a []int) int { // 无根树
		var f func(int, int) (notChosen, chosen int)
		f = func(v, fa int) (notChosen, chosen int) {
			chosen = a[v] // 1
			for _, w := range g[v] {
				if w != fa {
					nc, c := f(w, v)
					notChosen += c
					chosen += min(nc, c)
				}
			}
			return
		}
		return min(f(0, -1))
	}

	// 树上最小支配集
	// 返回最小点权和（最小支配集的情形即所有点权均为一）
	// 下面的定义省去了（……时的最小支配集的元素个数）   w 为 i 的儿子
	// 视频讲解：https://www.bilibili.com/video/BV1oF411U7qL/
	// f[i][0]：选 i = a[i]+∑min(f[w][0],f[w][1],f[w][2])
	// f[i][1]：不选 i，且 i 被儿子支配 = ∑min(f[w][0],f[w][1]) + max(min{f[w][0]-f[w][1]}, 0)
	// f[i][2]：不选 i，且 i 被父亲支配 = ∑min(f[w][0],f[w][1])
	// https://brooksj.com/2019/06/20/%E6%A0%91%E7%9A%84%E6%9C%80%E5%B0%8F%E6%94%AF%E9%85%8D%E9%9B%86%EF%BC%8C%E6%9C%80%E5%B0%8F%E7%82%B9%E8%A6%86%E7%9B%96%E9%9B%86%EF%BC%8C%E6%9C%80%E5%A4%A7%E7%82%B9%E7%8B%AC%E7%AB%8B%E9%9B%86/
	//
	// 监控二叉树 LC968 https://leetcode.cn/problems/binary-tree-cameras/
	// - https://codeforces.com/problemset/problem/1029/E
	// 保安站岗 https://www.luogu.com.cn/problem/P2458
	// 手机网络 https://www.luogu.com.cn/problem/P2899
	// https://ac.nowcoder.com/acm/problem/24953
	// todo EXTRA: 消防局的设立（支配距离为 2） https://www.luogu.com.cn/problem/P2279
	// todo EXTRA: 将军令（支配距离为 k） https://www.luogu.com.cn/problem/P3942
	//                                https://atcoder.jp/contests/arc116/tasks/arc116_e
	minDominatingSetOfTree := func(n int, g [][]int, a []int) int { // 无根树
		const inf int = 1e18
		var f func(int, int) (chosen, bySon, byFa int)
		f = func(v, fa int) (chosen, bySon, byFa int) {
			chosen = a[v] // 1
			extra := inf
			for _, w := range g[v] {
				if w != fa {
					c, bs, bf := f(w, v)
					m := min(c, bs)
					chosen += min(m, bf)
					bySon += m
					byFa += m
					extra = min(extra, c-bs)
				}
			}
			bySon += max(extra, 0)
			return
		}
		chosen, bySon, _ := f(0, -1)
		return min(chosen, bySon)
	}

	// EXTRA: 每个被支配的点，仅被一个点支配
	// Kaoshiung06，紫书例题 9-14，UVa 1218 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=247&page=show_problem&problem=3659

	// 树上最大匹配
	// g[v] = ∑{max(f[son],g[son])}
	// f[v] = max{1+g[son]+g[v]−max(f[son],g[son])}
	// https://codeforces.com/blog/entry/2059
	// https://blog.csdn.net/lycheng1215/article/details/78368002
	// https://vijos.org/p/1892
	maxMatchingOfTree := func(n int, g [][]int) int { // 无根树
		cover, nonCover := make([]int, n), make([]int, n)
		var f func(int, int)
		f = func(v, fa int) {
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
					nonCover[v] += max(cover[w], nonCover[w])
				}
			}
			for _, w := range g[v] {
				cover[v] = max(cover[v], 1+nonCover[w]+nonCover[v]-max(cover[w], nonCover[w]))
			}
		}
		f(0, -1)
		return max(cover[0], nonCover[0])
	}

	// 树上完美匹配子集个数
	// https://codeforces.com/problemset/problem/1032/F

	// todo 给一棵树和树上的一些关键节点，选 m 个点，使得关键节点到这些点中距离的最小值的最大值最小，求这个值
	//      https://www.luogu.com.cn/problem/P3523

	// 换根 DP / 二次扫描法
	// 【图解】一张图秒懂换根 DP！（Python/Java/C++/Go/JS）https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/
	// https://codeforces.com/blog/entry/20935
	// https://ei1333.hateblo.jp/entry/2017/04/10/224413
	//
	// todo 题集 https://atcoder-tags.herokuapp.com/tags/Dynamic-Programming/Every-Direction-DP
	//
	// LC310 也可以用拓扑排序的思想 https://leetcode.cn/problems/minimum-height-trees/
	// - https://codeforces.com/problemset/problem/1881/F
	// LC834 https://leetcode.cn/problems/sum-of-distances-in-tree
	// https://codeforces.com/problemset/problem/763/A 1600（有更巧妙的做法）
	// https://codeforces.com/problemset/problem/219/D 1700
	// - LC2858 https://leetcode.cn/problems/minimum-edge-reversals-so-every-node-is-reachable/
	// - LC2581 https://leetcode.cn/problems/count-number-of-possible-root-nodes/
	// https://codeforces.com/problemset/problem/1822/F 1700
	// https://codeforces.com/problemset/problem/1324/F 1800 类似最大子数组和
	// https://codeforces.com/problemset/problem/109/C 1900 也有组合数学做法
	// https://codeforces.com/problemset/problem/1092/F 1900
	// https://codeforces.com/problemset/problem/1882/D 1900
	// https://codeforces.com/problemset/problem/337/D 2000
	// https://codeforces.com/problemset/problem/791/D 2100
	// https://codeforces.com/problemset/problem/1187/E 2100
	// https://codeforces.com/problemset/problem/543/D 2300 注意不存在逆元的情形
	// https://codeforces.com/problemset/problem/1626/E 2400
	// https://codeforces.com/problemset/problem/1794/E 2400
	// https://codeforces.com/problemset/problem/1691/F 2500 计数
	// https://codeforces.com/problemset/problem/1320/E 3000 虚树
	// https://atcoder.jp/contests/dp/tasks/dp_v
	// https://atcoder.jp/contests/abc222/tasks/abc222_f 还可以用直径做 
	// todo https://atcoder.jp/contests/ttpc2019/tasks/ttpc2019_m
	//  https://atcoder.jp/contests/abc337/tasks/abc337_g
	// https://www.luogu.com.cn/problem/P3047 对于每个点 v，计算到点 v 距离为 k 的点权和（k 是定值）
	// https://www.luogu.com.cn/problem/P3478
	// https://www.luogu.com.cn/problem/P2986
	// https://ac.nowcoder.com/acm/contest/59717/F

	// 换根 DP · 其一（简单情况）
	// 第一次 DFS 算出以 0 为根的答案 ans0（一般是自底向上）
	// 第二次 DFS 基于 ans0，算出从节点 x 换到子节点 y 的答案的「变化量」，从而计算出其它节点的答案（一般是自顶向下）
	//
	// 每个点到其余点的距离之和 LC834 https://leetcode.cn/problems/sum-of-distances-in-tree
	// - 【图解】一张图秒懂换根 DP！https://leetcode.cn/problems/sum-of-distances-in-tree/solution/tu-jie-yi-zhang-tu-miao-dong-huan-gen-dp-6bgb/
	// - 变形：把距离之和改成每个距离的平方之和、立方之和
	// - 记录子树大小 size[v] 和子树每个节点的深度之和 sum(dep[sub])
	// https://atcoder.jp/contests/abc220/tasks/abc220_f
	// https://codeforces.com/problemset/problem/1324/F 1800 类似最大子数组和
	// https://codeforces.com/problemset/problem/109/C 1900 也有组合数学做法
	// https://codeforces.com/problemset/problem/791/D 2100 任意两点距离除以 k 的上取整之和
	// https://atcoder.jp/contests/abc160/tasks/abc160_f 2048=CF2260
	reroot1 := func(g [][]int) []int {
		ans := make([]int, len(g))
		size := make([]int, len(g))
		var dfs func(int, int, int)
		dfs = func(x, fa, depth int) {
			ans[0] += depth // 
			size[x] = 1
			for _, y := range g[x] {
				if y != fa {
					dfs(y, x, depth+1)
					// ans[0] += ...
					size[x] += size[y]
				}
			}
		}
		dfs(0, -1, 0)

		var reroot func(int, int)
		reroot = func(x, fa int) {
			for _, y := range g[x] {
				if y != fa {
					ans[y] = ans[x] + len(g) - size[y]*2 //
					reroot(y, x)
				}
			}
		}
		reroot(0, -1)
		return ans
	}

	// 换根 DP · 其二（维护最大次大）
	// LC3241 https://leetcode.cn/problems/time-taken-to-mark-all-nodes/
	// https://codeforces.com/problemset/problem/1822/F 1700
	// https://codeforces.com/problemset/problem/633/F 2600 计算最大次大第三大（也可以直接树形 DP，无需换根）
	reroot2 := func(g [][]struct{ to, wt int }) []int {
		nodes := make([]struct{ fi, se, fiW int }, len(g))
		var dfs func(int, int) int
		dfs = func(v, fa int) int {
			p := &nodes[v]
			for _, e := range g[v] {
				w := e.to
				if w == fa {
					continue
				}
				d := dfs(w, v) + e.wt // 从 v 出发，往 w 方向的最大链和
				if d > p.fi {
					p.se = p.fi
					p.fi = d
					p.fiW = w
				} else if d > p.se {
					p.se = d
				}
			}
			return p.fi
		}
		dfs(0, -1)

		ans := make([]int, len(g))
		var reroot func(int, int, int)
		reroot = func(v, fa, fromUp int) {
			p := nodes[v]
			ans[v] = max(fromUp, p.fi) // 从 v 出发的最大链和
			for _, e := range g[v] {
				w := e.to
				if w == fa {
					continue
				}
				exceptW := p.fi
				if w == p.fiW {
					exceptW = p.se // 对于 w 来说，上面要选次大的
				}
				reroot(w, v, max(fromUp, exceptW)+e.wt)
			}
		}
		reroot(0, -1, 0)
		return ans
	}

	// 换根 DP · 其三（前后缀分解写法，适用性最广）
	// 使用时根据题目修改 data unit moveEdge merge
	// https://nyaannyaan.github.io/library/tree/rerooting.hpp.html
	// https://qiita.com/keymoon/items/2a52f1b0fb7ef67fb89e
	// https://atcoder.jp/contests/dp/tasks/dp_v 母题
	// https://codeforces.com/contest/1822/problem/F 1700
	// https://codeforces.com/problemset/problem/543/D 2300
	// https://atcoder.jp/contests/abc160/tasks/abc160_f
	// https://atcoder.jp/contests/abc222/tasks/abc222_f
	rerootPreSuf := func(g [][]int, root int) {
		// type data struct{ x, y int }
		type data int
		const unit data = 0
		// 返回 d 在通过 v-w 边之后的结果    *也可以传入边权
		// swap=true 表示 v-w 是换根时的那条边
		moveEdge := func(d data, v, w int, swap bool) data {
			return d + 1 // add weight from v to w
		}
		// 返回 p 和 q 合并后的结果（p 和 q 已经包含边的信息）
		merge := func(p, q data) data {
			return max(p, q) // p + q
		}

		// 以 root 为根时的子树信息
		subData := make([]data, len(g))
		var dfs func(int, int)
		dfs = func(v, fa int) {
			res := unit
			for _, w := range g[v] {
				if w == fa {
					continue
				}
				dfs(w, v)
				res = merge(res, moveEdge(subData[w], v, w, false)) // v-w 边
			}
			subData[v] = res
		}
		dfs(root, -1)

		ansAtRoot := make([]data, len(g))
		var reroot func(int, int, data)
		reroot = func(v, fa int, movedFaData data) {
			// 必要时特判 fa < 0 的情况
			ansAtRoot[v] = merge(movedFaData, subData[v])

			// suf 是 g[v] 的子树后缀汇总信息（已经包含 v-g[v][i] 边）
			ngv := len(g[v])
			suf := make([]data, ngv+1)
			suf[ngv] = unit
			for i := ngv - 1; i >= 0; i-- {
				w := g[v][i]
				if w != fa {
					suf[i] = merge(suf[i+1], moveEdge(subData[w], v, w, false)) // v-w 边
				} else {
					suf[i] = suf[i+1]
				}
			}

			// pre 是 g[v] 子树前缀汇总信息（已经包含 v-g[v][i] 边）
			pre := unit
			for i, w := range g[v] {
				if w == fa {
					continue
				}
				// mergedData 是除了 subData[w] 以外的子树汇总信息（已经包含 v-g[v][i] 边）
				mergedData := merge(movedFaData, merge(pre, suf[i+1]))
				reroot(w, v, moveEdge(mergedData, w, v, true))      // w-v 边（以 w 为根）
				pre = merge(pre, moveEdge(subData[w], v, w, false)) // v-w 边
			}
		}
		reroot(root, -1, unit)

		for _, res := range ansAtRoot {
			_ = res
			// ...
		}
	}

	// 树上所有路径的位运算与(&)的和
	// 单个点也算路径
	// 解法：对每一位，统计仅含 1 的路径个数
	// a[i] <= 2^20
	// https://ac.nowcoder.com/acm/contest/10167/C
	andPathSum := func(g [][]int, a []int) int {
		const mx = 21
		ans := 0
		for i := 0; i < mx; i++ {
			cntOnePath := 0
			var f func(v, fa int) int
			f = func(v, fa int) int {
				one := a[v] >> i & 1
				cntOnePath += one
				for _, w := range g[v] {
					if w != fa {
						o := f(w, v)
						if one > 0 {
							cntOnePath += one * o
							one += o
						}
					}
				}
				return one
			}
			f(0, -1)
			ans += 1 << i * cntOnePath
		}

		{
			// 另一种做法是对每一位，用并查集求出 1 组成的连通分量，每个连通分量对答案的贡献是 sz*(sz+1)/2
			n := len(a)
			fa := make([]int, n)
			var find func(int) int
			find = func(x int) int {
				if fa[x] != x {
					fa[x] = find(fa[x])
				}
				return fa[x]
			}
			merge := func(from, to int) { fa[find(from)] = find(to) }

			ans := 0
			for i := 0; i < mx; i++ {
				for j := range fa {
					fa[j] = j
				}
				sz := make([]int, n)
				for v, vs := range g {
					for _, w := range vs {
						if a[v]>>i&1 > 0 && a[w]>>i&1 > 0 {
							merge(v, w)
						}
					}
				}
				for j := 0; j < n; j++ {
					sz[find(j)]++
				}
				for j, f := range fa {
					if j == f && a[j]>>i&1 > 0 {
						ans += 1 << i * sz[j] * (sz[j] + 1) / 2
					}
				}
			}
		}
		return ans
	}

	// 树上所有路径的位运算或(|)的和
	// 单个点也算路径
	// 做法和上面类似，求出仅含 0 的路径的个数，然后用路径总数 n*(n+1) 减去该个数就得到了包含至少一个 1 的路径个数
	// 也可以用并查集求出 0 组成的连通分量

	// 树上所有路径的位运算异或(^)的和
	// https://atcoder.jp/contests/abc201/tasks/abc201_e
	// 上面链接是边权，这里改成点权，且路径至少有两个点
	// 解法：由于任意路径异或和可以用从根节点出发的路径异或和表示
	// 对每一位，统计从根节点出发的路径异或和在该位上的 0 的个数和 1 的个数，
	// 只有当 0 与 1 异或时才对答案有贡献，所以贡献即为这两个个数之积
	xorPathSum := func(g [][]int, a []int) int {
		n := len(a)
		const mx = 30
		cnt := [mx]int{}
		xor := 0
		var f func(v, fa int)
		f = func(v, fa int) {
			xor ^= a[v]
			for _, w := range g[v] {
				if w != fa {
					f(w, v)
				}
			}
			for i := 0; i < mx; i++ {
				cnt[i] += xor >> i & 1
			}
			xor ^= a[v]
		}
		f(0, -1)
		ans := 0
		for i, c := range cnt {
			ans += 1 << i * c * (n - c)
		}
		return ans
	}

	// 树上所有路径的位运算异或(^)的异或和
	// 这里的路径至少有两个点
	// 方法是考虑每个点出现在多少条路径上，若数目为奇数则对答案有贡献
	// 路径分两种情况，一种是没有父节点参与的，树形 DP 一下就行了；另一种是父节点参与的，个数就是 子树*(n-子树)
	// https://ac.nowcoder.com/acm/contest/272/B
	xorPathXorSum := func(g [][]int, a []int) int {
		n := len(a)
		ans := 0
		var f func(v, fa int) int
		f = func(v, fa int) int {
			cnt := 0
			sz := 1
			for _, w := range g[v] {
				if w != fa {
					s := f(w, v)
					cnt += sz * s
					sz += s
				}
			}
			cnt += sz * (n - sz)
			// 若一个点也算路径，那就再加一。或者在递归结束后把 ans^=a[0]^...^a[n-1]
			if cnt&1 > 0 {
				ans ^= a[v]
			}
			return sz
		}
		f(0, -1)
		return ans
	}

	_ = []any{
		mapDP,
		maxSubarraySum, maxSubarraySumWithRange, maxTwoSubarraySum,
		maxAlternatingSumDP, maxAlternatingSumGreedy,
		slopeTrick,
		minimumArea,

		lcs, lcsPath, lcsCount,
		lisSlow, lis, lisAll, cntLis, lcis, lcisPath, countLIS,
		distinctSubsequence,
		palindromeO1Space, isPalindrome, minPalindromeCut,

		zeroOneKnapsack, zeroOneKnapsackExactlyFull, zeroOneKnapsackAtLeastFillUp, zeroOneWaysToSum, zeroOneKnapsackLexicographicallySmallestResult, zeroOneKnapsackByValue,
		unboundedKnapsack, unboundedWaysToSum,
		boundedKnapsack, boundedKnapsackBinary, boundedKnapsackMonotoneQueue, boundedKnapsackWays, boundedKnapsackWays2,
		groupKnapsack, groupKnapsackFill,
		treeKnapsack,

		longestPalindromeSubsequence, mergeStones, countPalindromes,

		// 期望 DP

		permDP, permDP2, tsp, longestSimplePath, countCycle,
		subsubDP, subsubDP2, subsubDPMemo, sosDP, plugDP,

		digitDP, digitDP2, calcSum, kth666, // 数位 DP

		binaryLifting,

		// 数据结构优化 DP

		cht, wqs,

		diameter, countDiameter, countPath, countVerticesOnDiameter, maxPathSum,
		maxIndependentSetOfTree, minVertexCoverOfTree, minDominatingSetOfTree, maxMatchingOfTree,
		reroot1, reroot2, rerootPreSuf,
		andPathSum, xorPathSum, xorPathXorSum,
	}
}
