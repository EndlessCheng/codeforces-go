package copypasta

import (
	. "fmt"
	"io"
	"math"
	"math/bits"
	"math/rand"
	"reflect"
	"sort"
	"time"
	"unsafe"
)

// Competitive Programming Roadmap (target: [gray, blue]) https://codeforces.com/blog/entry/111099
// A way to Practice Competitive Programming : From Rating 1000 to 2400+ https://codeforces.com/blog/entry/66909
// https://cin.ufpe.br/~fbma/Crack/%5BTutorial%5D%20A%20Way%20to%20Practice%20Competitive%20Programming.pdf

// 解决问题的一般方法 https://codeforces.com/blog/entry/92248?#comment-809401
// General ideas https://codeforces.com/blog/entry/48417
// 从特殊到一般：尝试修改条件或缩小题目的数据范围，先研究某个特殊情况下的思路，然后再逐渐扩大数据范围来思考怎么改进算法
// 重谈主定理及其证明 https://www.luogu.com.cn/blog/GJY-JURUO/master-theorem

/*
力扣题目分类汇总 https://leetcode.cn/circle/article/04PVPY/

双变量思想
- [1. 两数之和](https://leetcode.cn/problems/two-sum/)
- [219. 存在重复元素 II](https://leetcode.cn/problems/contains-duplicate-ii/)
- [1010. 总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/)
- [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/)

哈希表与前缀和（双变量思想）
推荐按照顺序完成
LC560 https://leetcode.cn/problems/subarray-sum-equals-k/
LC974 https://leetcode.cn/problems/subarray-sums-divisible-by-k/
LC1590 https://leetcode.cn/problems/make-sum-divisible-by-p/
LC523 https://leetcode.cn/problems/continuous-subarray-sum/
LC525 https://leetcode.cn/problems/contiguous-array/
LC930 https://leetcode-cn.com/problems/binary-subarrays-with-sum/

#### 前缀和+异或

- [1177. 构建回文串检测](https://leetcode.cn/problems/can-make-palindrome-from-substring/)
- [1371. 每个元音包含偶数次的最长子字符串](https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/)
- [1542. 找出最长的超赞子字符串](https://leetcode.cn/problems/find-longest-awesome-substring/)
- [1915. 最美子字符串的数目](https://leetcode.cn/problems/number-of-wonderful-substrings/)，[题解](https://leetcode.cn/problems/number-of-wonderful-substrings/solution/qian-zhui-he-chang-jian-ji-qiao-by-endle-t57t/)

https://leetcode.cn/problems/find-longest-subarray-lcci/
https://codeforces.com/problemset/problem/1296/C

前后缀分解
- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)（[视频讲解](https://www.bilibili.com/video/BV1Qg411q7ia/?t=3m05s)）
- [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/)
- [2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/)
- [2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/)
- [2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/)
- [2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/)
- [2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/)
- [2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/)
- [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/)

滑动窗口 双指针
入门题 https://codeforces.com/problemset/problem/602/B
LC424 https://leetcode.cn/problems/longest-repeating-character-replacement/
LC795 https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/
LC2444 https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/
LC1712 多指针 https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/
多指针 https://codeforces.com/problemset/problem/895/B
https://codeforces.com/contest/1833/problem/F
计算有多少子数组，其中有至少 k 个相同的数 https://codeforces.com/problemset/problem/190/D

分组循环 O(n)
LC2110 https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/
LC2760 https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/
LC2765 https://leetcode.cn/problems/longest-alternating-subarray/

巧妙枚举
https://codeforces.com/problemset/problem/1181/C
https://codeforces.com/problemset/problem/1626/D
https://codeforces.com/problemset/problem/846/C

贪心及其证明
- [2611. 老鼠和奶酪](https://leetcode.cn/problems/mice-and-cheese/)
- [1029. 两地调度](https://leetcode.cn/problems/two-city-scheduling/)
https://codeforces.com/problemset/problem/1369/C
	提示 1：前 k 大的数一定可以作为最大值。且尽量把大的数放在 w[i] = 1 的组中，这样可以计入答案两次。
	如果某个前 k 大的数 x 没有作为最大值（其中一个组的最大值是不在前 k 大中的 y），那么把 x 和 y 交换，
	如果 x 是某个组的最小值，那么交换后 y 必然也是最小值，此时答案不变。
	如果 x 不是某个组的最小值（这个组的最小值是 z）：
		   如果 y 交换后变成了最小值，那么答案变大了 x-z。
		   如果 y 交换后也不是最小值，那么答案变大了 x-y。
	无论如何，这样交换都不会使答案变小，因此前 k 大的数一定可以作为最大值。
	提示 2：然后来说最小值。a 的最小值必然要分到某个组中，为了「跳过」尽量多的较小的数，优先把 a 中较小的数分到 w 较大的组中。所以 a 从小到大遍历，w 从大到小遍历。
https://codeforces.com/problemset/problem/1479/B1 https://codeforces.com/problemset/problem/1479/B2
    https://www.luogu.com.cn/blog/wsyhb/post-ti-xie-cf1479b1-painting-the-array-i
1505. 最多 K 次交换相邻数位后得到的最小整数 https://leetcode.cn/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/
https://codeforces.com/problemset/problem/1804/D
https://codeforces.com/problemset/problem/442/C
    如果 x>=y<=z，那么删除 y 最优
    结束后剩下一个长为 m 的 /\ 形状的序列，由于无法取到最大值和次大值，那么加上剩下最小的 m-2 个数
https://codeforces.com/problemset/problem/1157/C2

每次取数组中大于 0 的连续一段同时减 1，求使数组全为 0 的最少操作次数
https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/solutions/371326/xing-cheng-mu-biao-shu-zu-de-zi-shu-zu-zui-shao-ze/
https://codeforces.com/problemset/problem/448/C

邻项交换
https://codeforces.com/problemset/problem/435/B
LC1665 完成所有任务的最少初始能量 https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/
https://atcoder.jp/contests/arc147/tasks/arc147_b
https://atcoder.jp/contests/abc268/tasks/abc268_f

区间与点的最大匹配/覆盖问题
https://www.luogu.com.cn/problem/P2887
https://codeforces.com/problemset/problem/555/B
https://codeforces.com/problemset/problem/863/E

观察、结论
https://codeforces.com/problemset/problem/1442/A
https://codeforces.com/problemset/problem/558/C
https://codeforces.com/problemset/problem/1610/E
https://codeforces.com/problemset/problem/1811/C
https://codeforces.com/problemset/problem/1822/D
https://codeforces.com/problemset/problem/1608/C 对拍找反例

构造
LC767 https://leetcode.cn/problems/reorganize-string/
LC667 https://leetcode.cn/problems/beautiful-arrangement-ii/
1500 https://codeforces.com/problemset/problem/1809/C
https://atcoder.jp/contests/arc145/tasks/arc145_a
+贪心 https://codeforces.com/problemset/problem/118/C
+分类讨论 https://codeforces.com/problemset/problem/584/C
+分类讨论 https://codeforces.com/problemset/problem/708/B
[1800·hot10] https://codeforces.com/problemset/problem/1554/D
相邻字母在字母表中不相邻 https://codeforces.com/contest/1156/problem/B
棋盘放最多的马 https://codeforces.com/problemset/problem/142/B
两点间恰好 k 条最短路径 http://codeforces.com/problemset/problem/388/B
https://codeforces.com/problemset/problem/327/D
https://codeforces.com/problemset/problem/515/D
度数均为 k 且至少（恰好）有一条割边 https://codeforces.com/problemset/problem/550/D
最短/最长 LIS https://codeforces.com/problemset/problem/1304/D
https://codeforces.com/problemset/problem/1789/D
交互 二分 https://codeforces.com/problemset/problem/1838/F

不好想到的构造
https://codeforces.com/contest/1659/problem/D
https://atcoder.jp/contests/abc178/tasks/abc178_f
https://codeforces.com/problemset/problem/1689/E 脑筋急转弯
https://codeforces.com/problemset/problem/1787/E

不变量（想一想，操作不会改变什么）
https://codeforces.com/contest/1775/problem/E 有点差分的味道，想想前缀和
https://atcoder.jp/contests/arc119/tasks/arc119_c 操作不影响交错和
https://codeforces.com/problemset/problem/1365/F 仍然对称

不变量 2（总和）
把一个环形数组切两刀，分成两段，要求相等，求方案数 => 和为 sum(a)/2 的子数组个数
LC494 https://leetcode.cn/problems/target-sum/

分类讨论（部分题是易错题）
https://codeforces.com/problemset/problem/489/C
https://codeforces.com/problemset/problem/1605/C
https://codeforces.com/problemset/problem/382/C
https://codeforces.com/problemset/problem/1051/C
https://codeforces.com/problemset/problem/1095/E
https://codeforces.com/problemset/problem/796/C
https://codeforces.com/problemset/problem/1594/F
https://codeforces.com/problemset/problem/1798/E
https://codeforces.com/problemset/problem/1811/F
https://codeforces.com/problemset/problem/1714/F 锻炼代码实现技巧的好题
https://codeforces.com/problemset/problem/1832/D2
https://codeforces.com/contest/1833/problem/G 样例给的挺良心的
https://codeforces.com/problemset/problem/411/C
https://codeforces.com/contest/1516/problem/C

大量分类讨论
https://codeforces.com/problemset/problem/356/C
https://codeforces.com/problemset/problem/460/D
https://codeforces.com/problemset/problem/1374/E2
https://codeforces.com/problemset/problem/1647/D
+构造 https://atcoder.jp/contests/arc153/tasks/arc153_c

贡献法
LC891 https://leetcode.cn/problems/sum-of-subsequence-widths/
LC2681 https://leetcode.cn/problems/power-of-heroes/
LC2763 https://leetcode.cn/problems/sum-of-imbalance-numbers-of-all-subarrays/
更多贡献法题目，见 monotone_stack.go
https://codeforces.com/problemset/problem/912/D
https://codeforces.com/problemset/problem/915/F
https://codeforces.com/problemset/problem/1208/E
https://codeforces.com/problemset/problem/1777/D 树
https://codeforces.com/problemset/problem/1788/D 好题！
https://codeforces.com/problemset/problem/1789/C 好题！
https://codeforces.com/problemset/problem/1808/D

其他
删除一个字符 + 删除最长连续前缀 https://codeforces.com/problemset/problem/1430/D
https://codeforces.com/problemset/problem/521/D

合法括号字符串 Regular Bracket Sequence, RBS
https://codeforces.com/problemset/problem/1821/E
https://codeforces.com/problemset/problem/1830/C

= 变成 <= 或者 >=
求前缀和/后缀和
https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-substrings/
*/

// 异类双变量：固定某变量统计另一变量的 [0,n)
//     EXTRA: 值域上的双变量，见 https://codeforces.com/contest/486/problem/D
// 同类双变量①：固定 i 统计 [0,n)
// 同类双变量②：固定 i 统计 [0,i-1]
// 套路：预处理数据（按照某种顺序排序/优先队列/BST/...），或者边遍历边维护，
//      然后固定变量 i，用均摊 O(1)~O(logn) 的复杂度统计范围内的另一变量 j
// 这样可以将复杂度从 O(n^2) 降低到 O(n) 或 O(nlogn)
// https://codeforces.com/problemset/problem/1194/E
// 进阶：https://codeforces.com/problemset/problem/1483/D
// 删除一段的最长连续递增 CERC10D https://codeforces.com/gym/101487
// 统计量是二元组的情形 https://codeforces.com/problemset/problem/301/D
// 好题 空间优化 https://codeforces.com/contest/1830/problem/B

// 双变量+下取整：枚举分母，然后枚举分子的范围，使得在该范围内的分子/分母是一个定值
// LC1862 https://leetcode.cn/problems/sum-of-floored-pairs/
// https://codeforces.com/problemset/problem/1706/D2

// 利用前缀和实现巧妙的构造 https://www.luogu.com.cn/blog/duyi/qian-zhui-he
// 邻项修改->前缀和->单项修改 https://codeforces.com/problemset/problem/1254/B2 https://ac.nowcoder.com/acm/contest/7612/C

// 双指针：如果 left == right 时 while 一定为 false，则可以省略 while 中 left < right 的判断
// 双序列双指针
// 背向双指针 LC360 https://leetcode.cn/problems/sort-transformed-array/
// 双指针+DP

// 滑动窗口
// https://codeforces.com/problemset/problem/165/C

/* 横看成岭侧成峰
转换为距离的众数 https://codeforces.com/problemset/problem/1365/C
转换为差分数组 https://codeforces.com/problemset/problem/1110/E
             https://codeforces.com/problemset/problem/1442/A
             https://codeforces.com/problemset/problem/1700/C
             https://codeforces.com/problemset/problem/1779/D 改成修改长为 x 的数组？
             https://www.luogu.com.cn/problem/P4552
转换为差 http://www.51nod.com/Challenge/Problem.html#problemId=1217
考虑每个点产生的贡献 https://codeforces.com/problemset/problem/1009/E
考虑每条边产生的负贡献 https://atcoder.jp/contests/abc173/tasks/abc173_f
考虑符合范围要求的贡献 https://codeforces.com/problemset/problem/1151/E
和式的另一视角。若每一项的值都在一个范围，不妨考虑另一个问题：值为 x 的项有多少个？https://atcoder.jp/contests/abc162/tasks/abc162_e
对所有排列考察所有子区间的性质，可以转换成对所有子区间考察所有排列。将子区间内部的排列和区间外部的排列进行区分，内部的性质单独研究，外部的当作 (n-(r-l))! 个排列 https://codeforces.com/problemset/problem/1284/C
从最大值入手 https://codeforces.com/problemset/problem/1381/B
等效性 LC1183 https://leetcode-cn.com/problems/maximum-number-of-ones/
LC1526 https://leetcode-cn.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
置换 https://atcoder.jp/contests/abc250/tasks/abc250_e
排序+最小操作次数 https://codeforces.com/contest/1367/problem/F2
https://codeforces.com/contest/1830/problem/A
从绝对值最大的开始思考 https://codeforces.com/contest/351/problem/E
*/

/* 
## 练习：离线（按难度分排序）

> 由于所有的询问数据都给出了，我们可以通过修改询问的顺序，达到降低时间复杂度的效果。相应的，在线算法就是按照输入的顺序处理，来一个处理一个。

- [2343. 裁剪数字后查询第 K 小的数字](https://leetcode.cn/problems/query-kth-smallest-trimmed-number/) 1652
- [2070. 每一个查询的最大美丽值](https://leetcode.cn/problems/most-beautiful-item-for-each-query/) 1724
- [2503. 矩阵查询可获得的最大分数](https://leetcode.cn/problems/maximum-number-of-points-from-grid-queries/) 2196
- [1851. 包含每个查询的最小区间](https://leetcode.cn/problems/minimum-interval-to-include-each-query/) 2286
- [1697. 检查边长度限制的路径是否存在](https://leetcode.cn/problems/checking-existence-of-edge-length-limited-paths/) 2300
- [2747. 统计没有收到请求的服务器数目](https://leetcode.cn/problems/count-zero-request-servers/)
- [1938. 查询最大基因差](https://leetcode.cn/problems/maximum-genetic-difference-query/) 2503
- [2736. 最大和查询](https://leetcode.cn/problems/maximum-sum-queries/) 2533
 */

/* 逆向思维 / 正难则反：从终点出发 / 小学奥数告诉我们，不可行方案永远比可行方案好求
LC803 https://leetcode.cn/problems/bricks-falling-when-hit/
LC936 https://leetcode.cn/problems/stamping-the-sequence/
LC1199 https://leetcode.cn/problems/minimum-time-to-build-blocks/
LC2382 https://leetcode.cn/problems/maximum-segment-sum-after-removals/
LCP52 https://leetcode.cn/problems/QO5KpG/
https://codeforces.com/problemset/problem/712/C
https://codeforces.com/problemset/problem/621/C
https://codeforces.com/problemset/problem/571/A
https://codeforces.com/problemset/problem/369/E
https://codeforces.com/problemset/problem/1644/D
https://codeforces.com/problemset/problem/1638/D
https://codeforces.com/problemset/problem/1672/D
https://codeforces.com/problemset/problem/1759/G 求字典序最小，通常可以从大往小思考

删除变添加
https://codeforces.com/problemset/problem/295/B
https://leetcode.cn/problems/maximum-segment-sum-after-removals/
*/

/* 奇偶性
https://codeforces.com/problemset/problem/763/B
https://codeforces.com/problemset/problem/1270/E
https://codeforces.com/problemset/problem/1332/E 配对法：将合法局面与非法局面配对
LC932 https://leetcode.cn/problems/beautiful-array/ 分治
*/

/* 相邻 传递性
https://codeforces.com/problemset/problem/1582/E
*/

/* 归纳：solve(n)->solve(n-1) 或者 solve(n-1)->solve(n)
https://codeforces.com/problemset/problem/1517/C
https://codeforces.com/problemset/problem/412/D
https://codeforces.com/problemset/problem/266/C
*/

/* 见微知著：考察单个点的规律，从而推出全局规律
https://codeforces.com/problemset/problem/1510/K
https://leetcode-cn.com/problems/minimum-number-of-operations-to-reinitialize-a-permutation/
*/

// 「恰好」转换成「至少/至多」https://codeforces.com/problemset/problem/1188/C

/* todo 反悔贪心
另见 heap.go 中的「反悔堆」
https://djy-juruo.blog.luogu.org/qian-tan-fan-hui-tan-xin
https://www.cnblogs.com/nth-element/p/11768155.html
题单 https://www.luogu.com.cn/training/8793
*/

/* 集合哈希
https://codeforces.com/problemset/problem/1394/B
https://www.luogu.com.cn/problem/P6688
*/

/* 操作树
和莫队类似，通过改变查询的顺序来优化复杂度
https://codeforces.com/problemset/problem/707/D
*/

/* Golang 卡常技巧（IO 之外的部分）
对于存在海量小对象的情况（如 trie, treap 等），使用 debug.SetGCPercent(-1) 来禁用 GC，能明显减少耗时
对于可以回收的情况（如 append 在超过 cap 时），使用 debug.SetGCPercent(-1) 虽然会减少些许耗时，但若有大量内存没被回收，会有 MLE 的风险
其他情况下使用 debug.SetGCPercent(-1) 对耗时和内存使用无明显影响
对于多组数据的情况，若禁用 GC 会 MLE，可在每组数据的开头或末尾调用 runtime.GC() 或 debug.FreeOSMemory() 手动 GC
参考 https://draveness.me/golang/docs/part3-runtime/ch07-memory/golang-garbage-collector/
    https://zhuanlan.zhihu.com/p/77943973

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

函数内的递归 lambda 会额外消耗非常多的内存（~100MB / 1e6 递归深度）
写在 main 里面 + slice MLE      https://codeforces.com/contest/767/submission/174193385
写在 main 里面 + array 257424KB https://codeforces.com/contest/767/submission/174194515
写在 main 外面 + slice 188364KB https://codeforces.com/contest/767/submission/174194380
写在 main 外面 + array 154500KB https://codeforces.com/contest/767/submission/174193693

测试：哈希表用时是数组的 13 倍（本题瓶颈）
slice    249ms https://codeforces.com/problemset/submission/570/209063267
hashmap 3259ms https://codeforces.com/problemset/submission/570/209063603
*/

// slice 作为 map 的 key
// 长度为 0 的 slice 对应空字符串
func intSliceAsMapKeyExample(a []int) {
	// 如果后面还会修改 a，可以先 copy 一份
	//a = append(a[:0:0], a...)
	cnt := map[string]int{}
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	// 装作 byte slice
	sh.Len *= bits.UintSize / 8
	cnt[*(*string)(unsafe.Pointer(sh))]++
}

func _() {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	pow10 := func(x int) int64 { return int64(math.Pow10(x)) } // 不需要 round

	// TIPS: dir4[i] 和 dir4[i^1] 互为相反方向
	dir4 := []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} // 上下左右（网格）
	// TIPS: dir4[i] 和 dir4[i^2] 互为相反方向
	dir4 = []struct{ x, y int }{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}   // 右下左上（网格，顺时针）
	dir4 = []struct{ x, y int }{{0, 1}, {-1, 0}, {0, -1}, {1, 0}}   // 右上左下（网格，逆时针）
	dir4 = []struct{ x, y int }{{1, 0}, {0, -1}, {-1, 0}, {0, 1}}   // 右下左上（坐标系，顺时针）
	dir4 = []struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}   // 右上左下（坐标系，逆时针）
	dir4 = []struct{ x, y int }{{1, 1}, {-1, 1}, {-1, -1}, {1, -1}} // 斜向

	dir4 = []struct{ x, y int }{'W': {-1, 0}, 'E': {1, 0}, 'S': {0, -1}, 'N': {0, 1}} // 西东南北（坐标系）
	dir4 = []struct{ x, y int }{'W': {0, -1}, 'E': {0, 1}, 'S': {1, 0}, 'N': {-1, 0}} // 西东南北（网格）
	dir4 = []struct{ x, y int }{'L': {-1, 0}, 'R': {1, 0}, 'D': {0, -1}, 'U': {0, 1}} // 左右下上（坐标系）
	dir4 = []struct{ x, y int }{'L': {0, -1}, 'R': {0, 1}, 'U': {-1, 0}, 'D': {1, 0}} // 左右下上（网格）

	dir8 := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}} // 逆时针（坐标系）
	dir8 = []struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}  // 顺时针（矩阵）

	perm3 := [][]int{{0, 1, 2}, {0, 2, 1}, {1, 0, 2}, {1, 2, 0}, {2, 0, 1}, {2, 1, 0}}
	perm4 := [][]int{
		{0, 1, 2, 3}, {0, 1, 3, 2}, {0, 2, 1, 3}, {0, 2, 3, 1}, {0, 3, 1, 2}, {0, 3, 2, 1},
		{1, 0, 2, 3}, {1, 0, 3, 2}, {1, 2, 0, 3}, {1, 2, 3, 0}, {1, 3, 0, 2}, {1, 3, 2, 0},
		{2, 0, 1, 3}, {2, 0, 3, 1}, {2, 1, 0, 3}, {2, 1, 3, 0}, {2, 3, 0, 1}, {2, 3, 1, 0},
		{3, 0, 1, 2}, {3, 0, 2, 1}, {3, 1, 0, 2}, {3, 1, 2, 0}, {3, 2, 0, 1}, {3, 2, 1, 0},
	}

	min := func(a, b int) int {
		if a < b {
			return a
		}
		return b
	}
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	mins := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v < res {
				res = v
			}
		}
		return res
	}
	maxs := func(a ...int) int {
		res := a[0]
		for _, v := range a[1:] {
			if v > res {
				res = v
			}
		}
		return res
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
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
	bin := func(v int) []byte {
		const maxLen = 30 // 62 for int64
		s := make([]byte, maxLen+1)
		for i := range s {
			s[i] = byte(v >> (maxLen - i) & 1)
		}
		return s
	}
	cmp := func(a, b int) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		}
		return 1
	}

	sort3 := func(a ...int) (x, y, z int) { sort.Ints(a); return a[0], a[1], a[2] }
	minString := func(a, b string) string {
		if len(a) != len(b) {
			if len(a) < len(b) {
				return a
			}
			return b
		}
		if a < b {
			return a
		}
		return b
	}
	ternaryI := func(cond bool, r1, r2 int) int {
		if cond {
			return r1
		}
		return r2
	}
	ternaryS := func(cond bool, r1, r2 string) string {
		if cond {
			return r1
		}
		return r2
	}
	zip := func(a, b []int) {
		n := len(a)
		type pair struct{ x, y int }
		ps := make([]pair, n)
		for i := range ps {
			ps[i] = pair{a[i], b[i]}
		}
	}
	zipI := func(a []int) {
		n := len(a)
		type pair struct{ x, y int }
		ps := make([]pair, n)
		for i := range ps {
			ps[i] = pair{a[i], i}
		}
	}

	mergeMap := func(x, y map[int]int) map[int]int {
		res := make(map[int]int, len(x)+len(y))
		for v, c := range x {
			res[v] = c
		}
		for v, c := range y {
			res[v] += c //
		}
		return res
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

	// 顺时针旋转矩阵 90°
	// 返回一个拷贝
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
	// 转置
	transpose := func(a [][]int) [][]int {
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

	// 适用于 mod 超过 int32 范围的情况
	// 还有一种用浮点数的写法，此略
	mul := func(a, b, mod int64) (res int64) {
		for ; b > 0; b >>= 1 {
			if b&1 == 1 {
				res = (res + a) % mod
			}
			a = (a + a) % mod
		}
		return
	}

	// https://en.wikipedia.org/wiki/Exponentiation_by_squaring
	pow := func(x, n, mod int64) int64 {
		x %= mod
		res := int64(1) % mod
		for ; n > 0; n >>= 1 {
			if n&1 == 1 {
				res = res * x % mod
			}
			x = x * x % mod
		}
		return res
	}

	// 等比数列求和取模
	// 返回 (x^0 + x^1 + ... + x^n) % mod
	// https://atcoder.jp/contests/abc293/tasks/abc293_e
	gp := func(x, n, mod int64) int64 {
		var f func(p, c int64) int64
		f = func(p, c int64) int64 {
			if c == 0 {
				return 1 % mod
			}
			res := (1 + pow(p, (c+1)/2, mod)) * f(p, (c-1)/2)
			if c%2 == 0 {
				res += pow(p, c, mod)
			}
			return res % mod
		}
		return f(x, n)
	}

	// 从低位到高位
	toAnyBase := func(x, base int) (res []int) {
		for ; x > 0; x /= base {
			res = append(res, x%base)
		}
		return
	}
	digits := func(x int) (res []int) {
		for ; x > 0; x /= 10 {
			res = append(res, x%10)
		}
		return
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

	// 合并有序数组，保留至多 k 个元素
	// https://codeforces.com/problemset/problem/587/C
	mergeWithLimit := func(a, b []int, k int) []int {
		i, n := 0, len(a)
		j, m := 0, len(b)
		res := make([]int, 0, min(n+m, k))
		for len(res) < k {
			if i == n {
				if len(res)+m-j > k {
					res = append(res, b[j:j+k-len(res)]...)
				} else {
					res = append(res, b[j:]...)
				}
				break
			}
			if j == m {
				if len(res)+n-i > k {
					res = append(res, a[i:i+k-len(res)]...)
				} else {
					res = append(res, a[i:]...)
				}
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
	subSum := func(a []int) []int {
		sum := make([]int, 1<<len(a)) // int64
		for i, v := range a {
			for mask, bit := 0, 1<<i; mask < bit; mask++ {
				sv := sum[mask] + v
				sum[bit|mask] = sv
				// NOTE: 若要直接在此考察 sv（相当于遍历 sum），注意别漏了 sum[0] = 0 的情况
			}
		}
		return sum
	}

	// 应用：给出由非负整数组成的数组 a 的子集和 sum，返回 a
	// 保证输入有解且 len(sum) = 2^len(a)
	// 变形：sum 包含负数 https://leetcode-cn.com/problems/find-array-given-subset-sums/
	// 做法是给所有 sum[i] 加上 -min(sum)，这会导致：
	// - 若 sum[i] 包含负数 a[i]，则新的 sum'[i] 就不包含 a[i]
	// - 若 sum[i] 不包含负数 a[i]，则新的 sum'[i] 会包含 -a[i]
	// 所以新的 sum' 也一样有解
	// 对 sum' 求出 a'
	// 由于 -min(sum) 是 a 的所有负数之和，所以找到一个 a' 的子集和，若其等于 -min(sum)，则将该子集在 a' 中的元素取相反数，就得到了 a
	recoverArrayFromSubsetSum := func(sum []int) []int {
		sort.Ints(sum)
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
	// 参考 https://leetcode-cn.com/problems/closest-subsequence-sum/solution/o2n2de-zuo-fa-by-heltion-0yn7/
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
	prefixSum := func(a []int) {
		//sort.Ints(a) // todo

		sum := make([]int, len(a)+1) // int64
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}

		// 返回所有数到 target 的距离之和，即 sum(abs(a[i]-target))
		// 需要保证 a 是有序的
		// LC2602 https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/
		distanceSum := func(target int) int {
			i := sort.SearchInts(a, target)
			left := target*i - sum[i]
			right := sum[len(a)] - sum[i] - target*(len(a)-i)
			return left + right
		}

		// EXTRA: 青蛙跳井
		// 一次询问（下标从 1 开始）https://codeforces.com/problemset/problem/1141/E
		// 多次询问（下标从 0 开始）https://codeforces.com/problemset/problem/1490/G

		_ = distanceSum
	}

	// 同余前缀和，a 的下标从 0 开始，k 为模数
	// 具体用法见 query 上的注释
	// LC1664 https://leetcode-cn.com/problems/ways-to-make-a-fair-array/
	// https://atcoder.jp/contests/abc288/tasks/abc288_d
	groupPrefixSum := func(a []int, k int) {
		sum := make([]int, len(a)+k) // int64
		for i, v := range a {
			sum[i+k] = sum[i] + v
		}
		pre := func(x, t int) int {
			if x%k <= t {
				return sum[x/k*k+t]
			}
			return sum[(x+k-1)/k*k+t]
		}
		// 求下标在 [l,r) 范围内且下标模 k 同余于 t 的所有元素之和
		query := func(l, r, t int) int {
			t %= k
			return pre(r, t) - pre(l, t)
		}
		_ = query
	}

	// 环形区间和 [l,r) 0<=l<r
	circularRangeSum := func(a []int) {
		n := len(a)
		sum := make([]int64, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + int64(v)
		}
		pre := func(p int) int64 {
			return sum[n]*int64(p/n) + sum[p%n]
		}
		query := func(l, r int) int64 {
			return pre(r) - pre(l)
		}

		_ = query
	}

	// 带权(等差数列)前缀和
	{
		var n int // read
		a := make([]int64, n)
		// read a ...

		sum := make([]int64, n+1)
		iSum := make([]int64, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
			iSum[i+1] = iSum[i] + int64(i+1)*v
		}
		query := func(l, r int) int64 { return iSum[r] - iSum[l] - int64(l)*(sum[r]-sum[l]) } // [l,r)

		_ = query
	}

	// 二维前缀和
	// 自加写法 https://codeforces.com/contest/835/submission/120031673
	// https://codeforces.com/contest/1107/problem/D
	// https://codeforces.com/problemset/problem/1731/D
	// https://codeforces.com/problemset/problem/611/C
	var sum2d [][]int
	initSum2D := func(a [][]int) {
		n, m := len(a), len(a[0])
		sum2d = make([][]int, n+1)
		sum2d[0] = make([]int, m+1)
		for i, row := range a {
			sum2d[i+1] = make([]int, m+1)
			for j, v := range row {
				sum2d[i+1][j+1] = sum2d[i+1][j] + sum2d[i][j+1] - sum2d[i][j] + v
			}
		}
	}
	// r1<=r<=r2 && c1<=c<=c2
	querySum2D := func(r1, c1, r2, c2 int) int {
		r2++
		c2++
		return sum2d[r2][c2] - sum2d[r2][c1] - sum2d[r1][c2] + sum2d[r1][c1]
	}

	// 矩阵每行每列的前缀和
	rowColSum := func(a [][]int) (sumR, sumC [][]int) {
		n, m := len(a), len(a[0])
		sumR = make([][]int, n) // int64
		for i, row := range a {
			sumR[i] = make([]int, m+1)
			for j, v := range row {
				sumR[i][j+1] = sumR[i][j] + v
			}
		}
		sumC = make([][]int, n+1) // int64
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

	// 矩阵斜向前缀和
	// LC1878 https://leetcode-cn.com/problems/get-biggest-three-rhombus-sums-in-a-grid/
	diagonalSum := func(a [][]int) {
		n, m := len(a), len(a[0])

		// int64
		ds := make([][]int, n+1) // 主对角线方向前缀和
		as := make([][]int, n+1) // 反对角线方向前缀和
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
		// 从 (x,y) 开始，向 ↘，连续的 k 个数的和（需要保证至少有 k 个数）
		queryDiagonal := func(x, y, k int) int { return ds[x+k][y+k] - ds[x][y] }
		// 从 (x,y) 开始，向 ↙，连续的 k 个数的和（需要保证至少有 k 个数）
		queryAntiDiagonal := func(x, y, k int) int { return as[x+k][y+1-k] - as[x][y+1] }

		_, _ = queryDiagonal, queryAntiDiagonal
	}

	// 利用每个数产生的贡献计算 ∑|ai-aj|, i!=j
	// 相关题目 https://codeforces.com/contest/1311/problem/F
	contributionSum := func(a []int) (sum int64) {
		n := len(a)
		sort.Ints(a)
		for i, v := range a {
			sum += int64(v) * int64(2*i+1-n)
		}
		return
	}

	// 差分
	// https://codeforces.com/problemset/problem/1700/C
	// 浮点数差分（也可以用扫描线）https://atcoder.jp/contests/abc274/tasks/abc274_f
	// 二阶差分 https://codeforces.com/problemset/problem/1661/D

	// 离散差分，传入闭区间列表 ps，不要求有序
	// https://codeforces.com/problemset/problem/1420/D
	diffMap := func(ps []struct{ l, r int }) {
		diff := map[int]int{} // or make with cap
		for _, p := range ps {
			diff[p.l]++
			diff[p.r+1]--
		}
		xs := make([]int, 0, len(diff)) // 坐标
		for x := range diff {
			xs = append(xs, x)
		}
		sort.Ints(xs)

		// 左闭右开区间 [_cnt[i].x, _cnt[i+1].x) 中的值都是 _cnt[i].c
		type _pair struct{ x, c int }
		_cnt := make([]_pair, len(xs))
		_c := 0
		for _, x := range xs {
			_c += diff[x]
			_cnt = append(_cnt, _pair{x, _c})
		}
		// 返回 x 被多少个 ps 中的区间包含（由于 ps 是闭区间，端点也算包含）
		query := func(x int) int {
			i := sort.Search(len(_cnt), func(i int) bool { return _cnt[i].x > x }) - 1
			if i < 0 {
				return 0
			}
			return _cnt[i].c
		}

		{
			// 如果只对左端点感兴趣，可以改为如下写法
			_cnt := make(map[int]int, len(xs)) // 前缀和
			_c := 0
			for _, x := range xs {
				_c += diff[x]
				_cnt[x] = _c
			}
		}

		_ = query
	}

	// 二维差分
	// https://blog.csdn.net/weixin_43914593/article/details/113782108
	// 模板题 LC2536 https://leetcode.cn/problems/increment-submatrices-by-one/
	// LC2132 https://leetcode-cn.com/problems/stamping-the-grid/（也可以不用差分）
	// https://www.luogu.com.cn/problem/P3397
	// LCP74 离散化 https://leetcode.cn/problems/xepqZ5/
	diff2D := func(n, m int) {
		diff := make([][]int, n+2)
		for i := range diff {
			diff[i] = make([]int, m+2)
		}
		// 将区域 r1<=r<=r2 && c1<=c<=c2 上的数都加上 x
		// 多 +1 是为了方便求前缀和
		update := func(r1, c1, r2, c2, x int) {
			diff[r1+1][c1+1] += x
			diff[r1+1][c2+2] -= x
			diff[r2+2][c1+1] -= x
			diff[r2+2][c2+2] += x
		}
		// 直接在 diff 上还原原始矩阵
		for i := 1; i <= n; i++ {
			for j := 1; j <= m; j++ {
				diff[i][j] += diff[i][j-1] + diff[i-1][j] - diff[i-1][j-1]
			}
		}
		// 切出中间的 n*m 的原始矩阵
		diff = diff[1 : n+1]
		for i, row := range diff {
			diff[i] = row[1 : m+1]
		}

		_ = update
	}

	reverse := func(a []byte) []byte {
		n := len(a)
		b := make([]byte, n)
		for i, v := range a {
			b[n-1-i] = v
		}
		return b
	}
	reverseInPlace := func(a []byte) {
		for i, n := 0, len(a); i < n/2; i++ {
			a[i], a[n-1-i] = a[n-1-i], a[i]
		}
	}

	equal := func(a, b []int) bool {
		// assert len(a) == len(b)
		for i, v := range a {
			if v != b[i] {
				return false
			}
		}
		return true
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
		sort.Ints(keys)
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

	// EXTRA: a 是否为 b 的子序列
	// https://codeforces.com/problemset/problem/778/A
	isSubSequence := func(a, b []int) bool {
		i, n := 0, len(a)
		j, m := 0, len(b)
		for {
			if i == n {
				return true
			}
			if j == m {
				return false
			}
			if a[i] == b[j] {
				i++
				j++
			} else {
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

	// 去重
	// a 必须是有序的
	unique := func(a []int) (res []int) {
		for i, v := range a {
			if i == 0 || v != a[i-1] {
				res = append(res, v)
			}
		}
		//n = len(res)
		return
	}

	// 原地去重
	// a 必须是有序的
	uniqueInPlace := func(a []int) []int {
		n := len(a)
		if n == 0 {
			return nil
		}
		k := 0
		for _, w := range a[1:] {
			if a[k] != w {
				k++
				a[k] = w
			}
		}
		//n = k + 1
		return a[:k+1]
	}

	// 简单离散化，适合没有重复元素的场景
	discreteSimple := func(a []int) []int {
		id := make([]int, len(a))
		for i := range id {
			id[i] = i
		}
		sort.Slice(id, func(i, j int) bool { return a[id[i]] < a[id[j]] })
		return id
	}

	// 离散化，返回离散化后的序列（名次）
	// discrete([]int{100,20,50,50}, 1) => []int{3,1,2,2}
	// 有些题目需要把 0 加进去离散化，请特别注意 https://atcoder.jp/contests/jsc2021/tasks/jsc2021_f
	// LC1331 https://leetcode-cn.com/problems/rank-transform-of-an-array/
	discrete := func(a []int, startIndex int) (kth []int) {
		type vi struct{ v, i int }
		ps := make([]vi, len(a))
		for i, v := range a {
			ps[i] = vi{v, i}
		}
		sort.Slice(ps, func(i, j int) bool { return ps[i].v < ps[j].v }) // or SliceStable
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

	// 另一种写法，不要求值连续 [10,30,20,20] => [0,3,1,1]
	// 对比，相差不大（~10%，Go 1.14.1）
	// discrete  296ms/14952KB https://atcoder.jp/contests/abc221/submissions/35791381
	// discrete2 333ms/11748KB https://atcoder.jp/contests/abc221/submissions/35791225
	discrete2 := func(a []int, startIndex int) []int {
		b := append([]int(nil), a...)
		sort.Ints(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + startIndex
		}
		return a
	}

	// 离散化，返回一个名次 map
	// discreteMap([]int{100,20,20,50}, 1) => map[int]int{20:1, 50:2, 100:3}
	// 例题：LC327 https://leetcode-cn.com/problems/count-of-range-sum/
	discreteMap := func(a []int, startIndex int) (kth map[int]int) {
		sorted := append([]int(nil), a...)
		sort.Ints(sorted)

		// 有重复元素
		kth = map[int]int{}
		k := startIndex
		for i, v := range sorted {
			if i == 0 || v != sorted[i-1] {
				kth[v] = k
				k++
			}
		}

		// 无重复元素
		kth = make(map[int]int, len(sorted))
		for i, v := range sorted {
			kth[v] = i + startIndex
		}

		// EXTRA: 第 k 小元素在原数组中的下标 kthPos
		pos := make(map[int][]int, k-startIndex)
		for i, v := range a {
			pos[v] = append(pos[v], i)
		}
		kthPos := make([][]int, k+1)
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

	allSame := func(a ...int) bool {
		for _, v := range a[1:] {
			if v != a[0] {
				return false
			}
		}
		return true
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
	// 模板题 LC215 https://leetcode-cn.com/problems/kth-largest-element-in-an-array/
	//       LC973 https://leetcode-cn.com/problems/k-closest-points-to-origin/submissions/
	// 模板题 https://codeforces.com/contest/977/problem/C
	quickSelect := func(a []int, k int) int {
		//k = len(a) - 1 - k // 求第 k 大
		rand.Seed(time.Now().UnixNano())
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

	contains := func(a []int, x int) bool {
		for _, v := range a {
			if v == x {
				return true
			}
		}
		return false
	}

	// x 是否包含 y 中的所有元素，且顺序一致
	containsAll := func(x, y []int) bool {
		for len(y) < len(x) {
			if len(y) == 0 {
				return true
			}
			if x[0] == y[0] {
				y = y[1:]
			}
			x = x[1:]
		}
		return false
	}

	// 扫描线 Events Sorting + Sweep Line
	// 常与树状数组、线段树、平衡树等数据结构结合
	// https://en.wikipedia.org/wiki/Sweep_line_algorithm
	// https://cses.fi/book/book.pdf 30.1
	// TODO 窗口的星星 https://www.luogu.com.cn/problem/P1502
	// TODO 矩形周长 https://www.luogu.com.cn/problem/P1856
	// 天际线问题 LC218 https://leetcode-cn.com/problems/the-skyline-problem/
	// TODO 矩形面积并 LC850 https://leetcode-cn.com/problems/rectangle-area-ii/ 《算法与实现》5.4.3
	//  矩形周长并 http://poj.org/problem?id=1177
	// 经典题 https://codeforces.com/problemset/problem/1000/C
	// https://codeforces.com/problemset/problem/1379/D
	// 转换求解目标 https://codeforces.com/problemset/problem/1285/E
	// 线段相交统计（栈）https://codeforces.com/problemset/problem/1278/D
	// 统计水平方向的线段与垂直方向的线段的交点个数 https://codeforces.com/problemset/problem/610/D
	// LC 套题 https://leetcode-cn.com/tag/line-sweep/
	// http://poj.org/problem?id=2932
	// 转换 https://atcoder.jp/contests/arc068/tasks/arc068_c
	sweepLine := func(in io.Reader, n int) {
		type event struct{ pos, delta int }
		events := make([]event, 0, 2*n)
		for i := 0; i < n; i++ {
			var l, r int
			Fscan(in, &l, &r)
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
	// 这样可以避免写一个复杂的 sort.Slice
	sweepLine2 := func(in io.Reader, n int) {
		events := make([]int, 0, 2*n)
		for i := 0; i < n; i++ {
			var l, r int
			Fscan(in, &l, &r)
			// 注意移位后是否溢出
			events = append(events, l<<1|1, r<<1) // 先出后进
			//events = append(events, l<<1, r<<1|1) // 先进后出
		}
		sort.Ints(events)

		for _, e := range events {
			pos, delta := e>>1, e&1
			_ = pos
			if delta > 0 { // 根据上面的写法来定义何为出何为进

			} else {

			}
		}
	}

	// 扫描线：一维格点刷漆，返回被刷到的格点数
	countCoveredPoints := func(in io.Reader, m int) int {
		type pair struct{ p, d int }
		es := make([]pair, 0, 2*m)
		for i := 0; i < m; i++ {
			var l, r int
			Fscan(in, &l, &r)
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
		sort.Ints(xs)
		xs = unique(xs)
		xi := discreteMap(xs, 0)
		sort.Ints(ys)
		ys = unique(ys)
		yi := discrete(ys, 0)

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
				return -1
			} // 无穷大
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

	_ = []interface{}{
		pow10, dir4, dir8, perm3, perm4,
		min, mins, max, maxs, abs, ceil, bin, cmp,
		ternaryI, ternaryS, zip, zipI, mergeMap, xorSet, rotateCopy, transpose, minString,
		pow, mul, gp, toAnyBase, digits,
		subSum, recoverArrayFromSubsetSum, subSumSorted,
		prefixSum, groupPrefixSum, circularRangeSum, initSum2D, querySum2D, rowColSum, diagonalSum,
		contributionSum,
		diffMap, diff2D,
		sort3, reverse, reverseInPlace, equal,
		merge, mergeWithLimit, splitDifferenceAndIntersection, intersection, isSubset, isSubSequence, isDisjoint,
		unique, uniqueInPlace, discreteSimple, discrete, discrete2, discreteMap,
		indexMap, allSame, complement, quickSelect, contains, containsAll,
		sweepLine, sweepLine2, countCoveredPoints,
		discrete2D,
	}
}
