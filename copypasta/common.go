package copypasta

import (
	"math"
	"math/bits"
	"math/rand"
	"reflect"
	"slices"
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
力扣题目分类汇总
https://leetcode.cn/circle/article/04PVPY/
https://leetcode.cn/circle/discuss/vEFf96/

枚举右，维护左
- [1. 两数之和](https://leetcode.cn/problems/two-sum/)
   - https://codeforces.com/problemset/problem/702/B
- [219. 存在重复元素 II](https://leetcode.cn/problems/contains-duplicate-ii/)
- [121. 买卖股票的最佳时机](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)
- [1512. 好数对的数目](https://leetcode.cn/problems/number-of-good-pairs/) 1161
- [2815. 数组中的最大数对和](https://leetcode.cn/problems/max-pair-sum-in-an-array/) 1295
- [2748. 美丽下标对的数目](https://leetcode.cn/problems/number-of-beautiful-pairs/) 1301
- [2342. 数位和相等数对的最大和](https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits/) 1309
- [1010. 总持续时间可被 60 整除的歌曲](https://leetcode.cn/problems/pairs-of-songs-with-total-durations-divisible-by-60/) 1377
- [2874. 有序三元组中的最大值 II](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii/) 1583
    巧妙安排更新顺序，使得 ans，pre_max 只能使用之前的值，从而符合 i<j<k 的要求
- [454. 四数相加 II](https://leetcode.cn/problems/4sum-ii/)
- https://leetcode.cn/problems/find-polygon-with-the-largest-perimeter/
https://leetcode.com/discuss/interview-question/3685049/25-variations-of-Two-sum-question

哈希表
- [2260. 必须拿起的最小连续卡牌数](https://leetcode.cn/problems/minimum-consecutive-cards-to-pick-up/) 1365
- [982. 按位与为零的三元组](https://leetcode.cn/problems/triples-with-bitwise-and-equal-to-zero/) 2085
- [面试题 16.21. 交换和](https://leetcode.cn/problems/sum-swap-lcci/)

前缀和
- [1732. 找到最高海拔](https://leetcode.cn/problems/find-the-highest-altitude/)
- [303. 区域和检索 - 数组不可变](https://leetcode.cn/problems/range-sum-query-immutable/)
- [1310. 子数组异或查询](https://leetcode.cn/problems/xor-queries-of-a-subarray/)
- [2615. 等值距离和](https://leetcode.cn/problems/sum-of-distances/) 1793
- [2602. 使数组元素全部相等的最少操作次数](https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/) 1903
- [2955. Number of Same-End Substrings](https://leetcode.cn/problems/number-of-same-end-substrings/)（会员题）

前缀和+哈希表（双变量思想）
- [930. 和相同的二元子数组](https://leetcode.cn/problems/binary-subarrays-with-sum/) 1592  *同 560
- [560. 和为 K 的子数组](https://leetcode.cn/problems/subarray-sum-equals-k/)
- [1524. 和为奇数的子数组数目](https://leetcode.cn/problems/number-of-sub-arrays-with-odd-sum/) 1611
- [974. 和可被 K 整除的子数组（数目）](https://leetcode.cn/problems/subarray-sums-divisible-by-k/) 1676
   - 变形：乘积可以被 k 整除
   - a[i] = gcd(a[i], k) 之后窗口乘积是 k 的倍数就行，不会乘爆
- [523. 连续的子数组和（长度至少为 2 且可被 K 整除）](https://leetcode.cn/problems/continuous-subarray-sum/)
- [1590.（移除最短子数组）使数组和能被 P 整除](https://leetcode.cn/problems/make-sum-divisible-by-p/) 2039
- [525. 连续数组](https://leetcode.cn/problems/contiguous-array/) *转换
- [1124. 表现良好的最长时间段](https://leetcode.cn/problems/longest-well-performing-interval/) 1908 *转换
- [2488. 统计中位数为 K 的子数组](https://leetcode.cn/problems/count-subarrays-with-median-k/) 1999 *转换
- [2949. 统计美丽子字符串 II](https://leetcode.cn/problems/count-beautiful-substrings-ii/)
- [面试题 17.05. 字母与数字](https://leetcode.cn/problems/find-longest-subarray-lcci/)
- [1983. 范围和相等的最宽索引对](https://leetcode.cn/problems/widest-pair-of-indices-with-equal-range-sum/)（会员题）
- [2489. 固定比率的子字符串数](https://leetcode.cn/problems/number-of-substrings-with-fixed-ratio/)（会员题）
https://atcoder.jp/contests/abc233/tasks/abc233_d
交错前缀和 https://codeforces.com/contest/1915/problem/E

前缀和思想 LC1523 https://leetcode.cn/problems/count-odd-numbers-in-an-interval-range/
有点数形结合 https://codeforces.com/problemset/problem/1748/C

前缀和的前缀和（二重前缀和）
LC2281 https://leetcode.cn/problems/sum-of-total-strength-of-wizards/
https://atcoder.jp/contests/abc058/tasks/arc071_b

前缀和+异或
- [1177. 构建回文串检测](https://leetcode.cn/problems/can-make-palindrome-from-substring/)
- [1371. 每个元音包含偶数次的最长子字符串](https://leetcode.cn/problems/find-the-longest-substring-containing-vowels-in-even-counts/)
- [1542. 找出最长的超赞子字符串](https://leetcode.cn/problems/find-longest-awesome-substring/)
- [1915. 最美子字符串的数目](https://leetcode.cn/problems/number-of-wonderful-substrings/)，[题解](https://leetcode.cn/problems/number-of-wonderful-substrings/solution/qian-zhui-he-chang-jian-ji-qiao-by-endle-t57t/)
https://atcoder.jp/contests/abc295/tasks/abc295_d
模 3 & 字符集大小为 n https://codeforces.com/problemset/problem/1418/G

https://leetcode.cn/problems/find-longest-subarray-lcci/
https://codeforces.com/problemset/problem/1296/C

前后缀分解（右边数字为难度分）
- [42. 接雨水](https://leetcode.cn/problems/trapping-rain-water/)（[视频讲解](https://www.bilibili.com/video/BV1Qg411q7ia/?t=3m05s)）
  注：带修改的接雨水 https://codeforces.com/gym/104821/problem/M
  - https://www.zhihu.com/question/627281278/answer/3280684055
- [2256. 最小平均差](https://leetcode.cn/problems/minimum-average-difference/) 1395
- [2909. 元素和最小的山形三元组 II](https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-ii/) 1479
- [2483. 商店的最少代价](https://leetcode.cn/problems/minimum-penalty-for-a-shop/) 1495
- [2874. 有序三元组中的最大值 II](https://leetcode.cn/problems/maximum-value-of-an-ordered-triplet-ii/) 1583
- [2420. 找到所有好下标](https://leetcode.cn/problems/find-all-good-indices/) 1695
- [1671. 得到山形数组的最少删除次数](https://leetcode.cn/problems/minimum-number-of-removals-to-make-mountain-array/) 1913 *DP
- [238. 除自身以外数组的乘积](https://leetcode.cn/problems/product-of-array-except-self/) ~2000
- [2906. 构造乘积矩阵](https://leetcode.cn/problems/construct-product-matrix/) 2075
- [2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/) 2219 *DP
- [2484. 统计回文子序列数目](https://leetcode.cn/problems/count-palindromic-subsequences/) 2223
- [2565. 最少得分子序列](https://leetcode.cn/problems/subsequence-with-the-minimum-score/) 2432
- [2552. 统计上升四元组](https://leetcode.cn/problems/count-increasing-quadruplets/) 2433

#### 定长滑动窗口（右边数字为难度分）
- [1456. 定长子串中元音的最大数目](https://leetcode.cn/problems/maximum-number-of-vowels-in-a-substring-of-given-length/) 1263
- [2269. 找到一个数字的 K 美丽值](https://leetcode.cn/problems/find-the-k-beauty-of-a-number/) 1280
- [1984. 学生分数的最小差值](https://leetcode.cn/problems/minimum-difference-between-highest-and-lowest-of-k-scores/) 1306
- [643. 子数组最大平均数 I](https://leetcode.cn/problems/maximum-average-subarray-i/)
- [1343. 大小为 K 且平均值大于等于阈值的子数组数目](https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/) 1317
- [2090. 半径为 k 的子数组平均值](https://leetcode.cn/problems/k-radius-subarray-averages/) 1358
- [2379. 得到 K 个黑块的最少涂色次数](https://leetcode.cn/problems/minimum-recolors-to-get-k-consecutive-black-blocks/) 1360
- [1052. 爱生气的书店老板](https://leetcode.cn/problems/grumpy-bookstore-owner/) 1418
- [2841. 几乎唯一子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/) 1546
- [2461. 长度为 K 子数组中的最大和](https://leetcode.cn/problems/maximum-sum-of-distinct-subarrays-with-length-k/) 1553
- [1423. 可获得的最大点数](https://leetcode.cn/problems/maximum-points-you-can-obtain-from-cards/) 1574
- [2134. 最少交换次数来组合所有的 1 II](https://leetcode.cn/problems/minimum-swaps-to-group-all-1s-together-ii/) 1748
- [2653. 滑动子数组的美丽值](https://leetcode.cn/problems/sliding-subarray-beauty/) 1786
- [567. 字符串的排列](https://leetcode.cn/problems/permutation-in-string/)
- [438. 找到字符串中所有字母异位词](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)
- [2156. 查找给定哈希值的子串](https://leetcode.cn/problems/find-substring-with-given-hash-value/) 2063
- [2953. 统计完全子字符串](https://leetcode.cn/problems/count-complete-substrings/)
- [346. 数据流中的移动平均值](https://leetcode.cn/problems/moving-average-from-data-stream/)（会员题）
- [1100. 长度为 K 的无重复字符子串](https://leetcode.cn/problems/find-k-length-substrings-with-no-repeated-characters/)（会员题）
- [1852. 每个子数组的数字种类数](https://leetcode.cn/problems/distinct-numbers-in-each-subarray/)（会员题）
- [2067. 等计数子串的数量](https://leetcode.cn/problems/number-of-equal-count-substrings/)（会员题）
- [2107. 分享 K 个糖果后独特口味的数量](https://leetcode.cn/problems/number-of-unique-flavors-after-sharing-k-candies/)（会员题）
https://codeforces.com/problemset/problem/69/E 1800

#### 不定长滑动窗口（求最长/最大）
- [3. 无重复字符的最长子串](https://leetcode.cn/problems/longest-substring-without-repeating-characters/)
   - 翻转至多一个任意子串后的无重复字符的最长子串 https://codeforces.com/contest/1234/problem/F
- [1493. 删掉一个元素以后全为 1 的最长子数组](https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/) 1423
- [904. 水果成篮](https://leetcode.cn/problems/fruit-into-baskets/) 1516
- [1695. 删除子数组的最大得分](https://leetcode.cn/problems/maximum-erasure-value/) 1529
- [2958. 最多 K 个重复元素的最长子数组](https://leetcode.cn/problems/length-of-longest-subarray-with-at-most-k-frequency/) 1535
- [2841. 几乎唯一子数组的最大和](https://leetcode.cn/problems/maximum-sum-of-almost-unique-subarray/) 1546
- [2024. 考试的最大困扰度](https://leetcode.cn/problems/maximize-the-confusion-of-an-exam/) 1643
- [1004. 最大连续1的个数 III](https://leetcode.cn/problems/max-consecutive-ones-iii/) 1656
- [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/) 1672  *需要 SortedList
- [2401. 最长优雅子数组](https://leetcode.cn/problems/longest-nice-subarray/) 1750  *位运算
- [1658. 将 x 减到 0 的最小操作数](https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/) 1817
- [1838. 最高频元素的频数](https://leetcode.cn/problems/frequency-of-the-most-frequent-element/) 1876
- [2831. 找出最长等值子数组](https://leetcode.cn/problems/find-the-longest-equal-subarray/) 1976
- [2106. 摘水果](https://leetcode.cn/problems/maximum-fruits-harvested-after-at-most-k-steps/) 2062
- [1610. 可见点的最大数目](https://leetcode.cn/problems/maximum-number-of-visible-points/) 2147
- [159. 至多包含两个不同字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-most-two-distinct-characters/)（会员题）
- [340. 至多包含 K 个不同字符的最长子串](https://leetcode.cn/problems/longest-substring-with-at-most-k-distinct-characters/)（会员题）
与单调队列结合 https://www.luogu.com.cn/problem/P3594

#### 不定长滑动窗口（求最短/最小）
- [209. 长度最小的子数组](https://leetcode.cn/problems/minimum-size-subarray-sum/)
- [1234. 替换子串得到平衡字符串](https://leetcode.cn/problems/replace-the-substring-for-balanced-string/) 1878
- [1574. 删除最短的子数组使剩余数组有序](https://leetcode.cn/problems/shortest-subarray-to-be-removed-to-make-array-sorted/) 1932
- [76. 最小覆盖子串](https://leetcode.cn/problems/minimum-window-substring/)
- [面试题 17.18. 最短超串](https://leetcode.cn/problems/shortest-supersequence-lcci/)
https://codeforces.com/problemset/problem/701/C

#### 不定长滑动窗口（求子数组个数）
- [2799. 统计完全子数组的数目](https://leetcode.cn/problems/count-complete-subarrays-in-an-array/) 1398
- [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)
- [1358. 包含所有三种字符的子字符串数目](https://leetcode.cn/problems/number-of-substrings-containing-all-three-characters/) 1646
- [2962. 统计最大元素出现至少 K 次的子数组](https://leetcode.cn/problems/count-subarrays-where-max-element-appears-at-least-k-times/) 1701
- [2302. 统计得分小于 K 的子数组数目](https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/) 1808
- [2537. 统计好子数组的数目](https://leetcode.cn/problems/count-the-number-of-good-subarrays/) 1892
- [2762. 不间断子数组](https://leetcode.cn/problems/continuous-subarrays/) 1940
- [2743. 计算没有重复字符的子字符串数量](https://leetcode.cn/problems/count-substrings-without-repeating-character/)（会员题）
和至少为 k 的子数组个数 https://atcoder.jp/contests/abc130/tasks/abc130_d
变形：改成子数组 https://codeforces.com/problemset/problem/550/B
其它题目见【前缀和】

#### 多指针滑动窗口
- [930. 和相同的二元子数组](https://leetcode.cn/problems/binary-subarrays-with-sum/) 1592
- [1248. 统计「优美子数组」](https://leetcode.cn/problems/count-number-of-nice-subarrays/) 1624
- [2563. 统计公平数对的数目](https://leetcode.cn/problems/count-the-number-of-fair-pairs/) 1721
- [1712. 将数组分成三个子数组的方案数](https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays/) 2079
- [2444. 统计定界子数组的数目](https://leetcode.cn/problems/count-subarrays-with-fixed-bounds/) 2093
- [992. K 个不同整数的子数组](https://leetcode.cn/problems/subarrays-with-k-different-integers/) 2210
- [1989. 捉迷藏中可捕获的最大人数](https://leetcode.cn/problems/maximum-number-of-people-that-can-be-caught-in-tag/)（会员题）

入门题 https://codeforces.com/problemset/problem/602/B
入门题 https://codeforces.com/problemset/problem/279/B
https://atcoder.jp/contests/abc229/tasks/abc229_d
LC424 替换后的最长重复字符 有些特殊的滑窗 https://leetcode.cn/problems/longest-repeating-character-replacement/
LC2271 毯子覆盖的最多白色砖块数 需要多思考一点点 https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/
- https://atcoder.jp/contests/abc098/tasks/arc098_b
较为复杂 https://atcoder.jp/contests/abc294/tasks/abc294_e
      - https://ac.nowcoder.com/acm/contest/62033/D
https://codeforces.com/problemset/problem/1208/B
https://codeforces.com/problemset/problem/1765/D
多指针 https://codeforces.com/problemset/problem/895/B
https://codeforces.com/contest/1833/problem/F
计算有多少子数组，其中有至少 k 个相同的数 https://codeforces.com/problemset/problem/190/D
mex https://atcoder.jp/contests/abc194/tasks/abc194_e
https://codeforces.com/problemset/problem/165/C

双序列双指针
LC88 https://leetcode.cn/problems/merge-sorted-array/
LC360（背向双指针）https://leetcode.cn/problems/sort-transformed-array/
LC986 区间列表的交集 https://leetcode.cn/problems/interval-list-intersections/

相向双指针
LC2824 https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
LC923 https://leetcode.cn/problems/3sum-with-multiplicity/
https://www.facebook.com/codingcompetitions/hacker-cup/2023/practice-round/problems/C

同时用到同向双指针和相向双指针的题
https://atcoder.jp/contests/abc155/tasks/abc155_d
- 相似题目 https://leetcode.cn/problems/kth-smallest-product-of-two-sorted-arrays/

a[i] + b[j] = target 的方案数
a[i] + b[j] < target 的方案数    相向双指针 https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
a[i] + b[j] > target 的方案数    同上
a[i] - b[j] = target 的方案数
a[i] - b[j] < target 的方案数    滑窗
a[i] - b[j] > target 的方案数    同上
子数组元素和 = < > target 的方案数：用前缀和，转换成上面 a[i] - b[j] 的形式
子序列元素和 = < > target 的方案数：0-1 背包恰好/至多/至少，见 https://www.bilibili.com/video/BV16Y411v7Y6/ 末尾的总结

## 分组循环

https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/solution/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-zuspx/

**适用场景**：按照题目要求，数组会被分割成若干组，每一组的判断/处理逻辑是相同的。

**核心思想**：

- 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
- 内层循环负责遍历组，找出这一组最远在哪结束。

这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组（易错点）。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。

- [1446. 连续字符](https://leetcode.cn/problems/consecutive-characters/)
- [1869. 哪种连续子字符串更长](https://leetcode.cn/problems/longer-contiguous-segments-of-ones-than-zeros/)
- [1957. 删除字符使字符串变好](https://leetcode.cn/problems/delete-characters-to-make-fancy-string/)
- [228. 汇总区间](https://leetcode.cn/problems/summary-ranges/)
- [2038. 如果相邻两个颜色均相同则删除当前颜色](https://leetcode.cn/problems/remove-colored-pieces-if-both-neighbors-are-the-same-color/)
- [1759. 统计同质子字符串的数目](https://leetcode.cn/problems/count-number-of-homogenous-substrings/)
- [2110. 股票平滑下跌阶段的数目](https://leetcode.cn/problems/number-of-smooth-descent-periods-of-a-stock/)
- [1578. 使绳子变成彩色的最短时间](https://leetcode.cn/problems/minimum-time-to-make-rope-colorful/)
- [1839. 所有元音按顺序排布的最长子字符串](https://leetcode.cn/problems/longest-substring-of-all-vowels-in-order/)
- [2760. 最长奇偶子数组](https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/)
- [2765. 最长交替子序列](https://leetcode.cn/problems/longest-alternating-subarray/)
- [795. 区间子数组个数](https://leetcode.cn/problems/number-of-subarrays-with-bounded-maximum/)
LC1887 https://leetcode.cn/problems/reduction-operations-to-make-the-array-elements-equal/
LC1180（会员）https://leetcode.cn/problems/count-substrings-with-only-one-distinct-letter/
LC2257 https://leetcode.cn/problems/count-unguarded-cells-in-the-grid/
- https://atcoder.jp/contests/abc317/tasks/abc317_e
LC2495（会员）逆向思维 https://leetcode.cn/problems/number-of-subarrays-having-even-product/
https://codeforces.com/problemset/problem/1380/C 1400
https://codeforces.com/problemset/problem/620/C 1500
https://codeforces.com/problemset/problem/525/C 1600
https://codeforces.com/problemset/problem/1748/C 1600

巧妙枚举
LC939 https://leetcode.cn/problems/minimum-area-rectangle/
https://codeforces.com/problemset/problem/1181/C
https://codeforces.com/problemset/problem/1626/D
https://codeforces.com/problemset/problem/846/C

贪心及其证明
- [455. 分发饼干](https://leetcode.cn/problems/assign-cookies/)
- [1029. 两地调度](https://leetcode.cn/problems/two-city-scheduling/) 1348
- [2410. 运动员和训练师的最大匹配数](https://leetcode.cn/problems/maximum-matching-of-players-with-trainers/) 1381
- [881. 救生艇](https://leetcode.cn/problems/boats-to-save-people/) 1530
    - https://codeforces.com/problemset/problem/1690/E
    - https://www.lanqiao.cn/problems/4174/learning/?contest_id=135
    - https://codeforces.com/problemset/problem/1765/D
- [2611. 老鼠和奶酪](https://leetcode.cn/problems/mice-and-cheese/) 1663
- [2931. 购买物品的最大开销](https://leetcode.cn/problems/maximum-spending-after-buying-items/) 1822
- [2952. 需要添加的硬币的最小数量](https://leetcode.cn/problems/minimum-number-of-coins-to-be-added/)
    - LC330 https://leetcode.cn/problems/patching-array/
- [2136. 全部开花的最早一天](https://leetcode.cn/problems/earliest-possible-day-of-full-bloom/) 2033
- todo 复习 [2193. 得到回文串的最少操作次数](https://leetcode.cn/problems/minimum-number-of-moves-to-make-palindrome/) 2091
- [1505. 最多 K 次交换相邻数位后得到的最小整数](https://leetcode.cn/problems/minimum-possible-integer-after-at-most-k-adjacent-swaps-on-digits/) 2337
https://codeforces.com/problemset/problem/388/A 1400
https://codeforces.com/problemset/problem/1443/C 1400
https://codeforces.com/problemset/problem/864/D 1500
https://codeforces.com/problemset/problem/1691/C
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
https://codeforces.com/problemset/problem/1804/D
https://codeforces.com/problemset/problem/442/C
    如果 x>=y<=z，那么删除 y 最优
    结束后剩下一个长为 m 的 /\ 形状的序列，由于无法取到最大值和次大值，那么加上剩下最小的 m-2 个数
https://codeforces.com/problemset/problem/1157/C2
https://www.luogu.com.cn/problem/UVA11384 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=25&page=show_problem&problem=2379
倒序思维 https://codeforces.com/problemset/problem/1707/A
https://codeforces.com/contest/1873/problem/G
https://atcoder.jp/contests/arc147/tasks/arc147_e 难

中位数贪心（右边数字为难度分） // 注：算长度用左闭右开区间思考，算中间值用闭区间思考    两个中位数分别是 a[(n-1)/2] 和 a[n/2]
为方便描述，将 $\textit{nums}$ 简记为 $a$。
**定理**：将 $a$ 的所有元素变为 $a$ 的**中位数**是最优的。
**证明**：设 $a$ 的长度为 $n$，设要将所有 $a[i]$ 变为 $x$。假设 $a$ 已经从小到大排序。首先，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之外，那么 $x$ 向区间方向移动可以使距离和变小；同时，如果 $x$ 取在区间 $[a[0],a[n-1]]$ 之内，无论如何移动 $x$，它到 $a[0]$ 和 $a[n-1]$ 的距离和都是一个定值 $a[n-1]-a[0]$，那么去掉 $a[0]$ 和 $a[n-1]$ 这两个最左最右的数，问题规模缩小。不断缩小问题规模，如果最后剩下 $1$ 个数，那么 $x$ 就取它；如果最后剩下 $2$ 个数，那么 $x$ 取这两个数之间的任意值都可以（包括这两个数）。因此 $x$ 可以取 $a[n/2]$。
题单（右边数字为难度分）
- [462. 最小操作次数使数组元素相等 II](https://leetcode.cn/problems/minimum-moves-to-equal-array-elements-ii/)
- [2033. 获取单值网格的最小操作数](https://leetcode.cn/problems/minimum-operations-to-make-a-uni-value-grid/) 1672
- [2448. 使数组相等的最小开销](https://leetcode.cn/problems/minimum-cost-to-make-array-equal/) 2005
- [2607. 使子数组元素和相等](https://leetcode.cn/problems/make-k-subarray-sums-equal/) 2071
https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/
https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/
- [1703. 得到连续 K 个 1 的最少相邻交换次数](https://leetcode.cn/problems/minimum-adjacent-swaps-for-k-consecutive-ones/) 2467
- [LCP 24. 数字游戏](https://leetcode.cn/problems/5TxKeK/)
https://codeforces.com/problemset/problem/710/B 1400

每次取两个数减一，最后剩下的数最小
https://cs.stackexchange.com/a/145450

每次取数组中大于 0 的连续一段同时减 1，求使数组全为 0 的最少操作次数
https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/solutions/371326/xing-cheng-mu-biao-shu-zu-de-zi-shu-zu-zui-shao-ze/
https://codeforces.com/problemset/problem/448/C

邻项交换（最小代价排序/字典序最小）
某些题目和逆序对有关
LC1665 https://leetcode.cn/problems/minimum-initial-energy-to-finish-tasks/ 1901
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

观察、结论
https://codeforces.com/problemset/problem/1442/A
https://codeforces.com/problemset/problem/558/C
https://codeforces.com/problemset/problem/1610/E
https://codeforces.com/problemset/problem/1811/C
https://codeforces.com/problemset/problem/1822/D
https://codeforces.com/problemset/problem/1608/C 对拍找反例

脑筋急转弯
LC1503 https://leetcode.cn/problems/last-moment-before-all-ants-fall-out-of-a-plank/
LC2731 https://leetcode.cn/problems/movement-of-robots/
LC280 https://leetcode.cn/problems/wiggle-sort/
https://codeforces.com/problemset/problem/1009/B 1400
https://codeforces.com/problemset/problem/1169/B 1500
https://codeforces.com/problemset/problem/1763/C 2000
https://atcoder.jp/contests/abc194/tasks/abc194_e
https://atcoder.jp/contests/abc196/tasks/abc196_e
https://www.luogu.com.cn/problem/UVA10881 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=20&page=show_problem&problem=1822

构造
todo 题单 https://www.luogu.com.cn/training/14#problems
LC767 https://leetcode.cn/problems/reorganize-string/
LC667 https://leetcode.cn/problems/beautiful-arrangement-ii/
LC2573 https://leetcode.cn/problems/find-the-string-with-lcp/ 2682
https://atcoder.jp/contests/arc145/tasks/arc145_a
https://codeforces.com/contest/803/problem/A 1400
https://codeforces.com/problemset/problem/1863/D 1400
https://codeforces.com/problemset/problem/1809/C 1500
贪心 https://codeforces.com/problemset/problem/118/C
分类讨论 https://codeforces.com/problemset/problem/584/C
分类讨论 https://codeforces.com/problemset/problem/708/B
https://codeforces.com/problemset/problem/1554/D 1800
分类讨论 https://codeforces.com/contest/1854/problem/A2 1900
相邻字母在字母表中不相邻 https://codeforces.com/contest/1156/problem/B
棋盘放最多的马 https://codeforces.com/problemset/problem/142/B
两点间恰好 k 条最短路径 http://codeforces.com/problemset/problem/388/B
https://codeforces.com/problemset/problem/327/D
https://codeforces.com/problemset/problem/515/D
度数均为 k 且至少（恰好）有一条割边 https://codeforces.com/problemset/problem/550/D
最短/最长 LIS https://codeforces.com/problemset/problem/1304/D
https://codeforces.com/problemset/problem/1789/D
交互 二分 https://codeforces.com/problemset/problem/1838/F
todo https://codeforces.com/problemset/problem/1823/D
OR 构造 https://atcoder.jp/contests/agc015/tasks/agc015_d
https://codeforces.com/problemset/problem/1761/E

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
https://codeforces.com/problemset/problem/193/A
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
https://codeforces.com/problemset/problem/1763/C
https://codeforces.com/problemset/problem/960/B
https://codeforces.com/problemset/problem/1799/C
https://codeforces.com/problemset/problem/1009/B 1400 脑筋急转弯
https://codeforces.com/contest/1251/problem/B 1400
https://codeforces.com/problemset/problem/1180/B 1500
https://codeforces.com/problemset/problem/750/C 1600 *也有偏数学的做法
https://codeforces.com/problemset/problem/1861/C 1600 好题！
https://codeforces.com/contest/1914/problem/F 1900
https://codeforces.com/problemset/problem/1761/E 2400
https://atcoder.jp/contests/diverta2019/tasks/diverta2019_c
https://atcoder.jp/contests/abc155/tasks/abc155_d
https://atcoder.jp/contests/abc125/tasks/abc125_d
https://atcoder.jp/contests/arc134/tasks/arc134_d 1998

大量分类讨论
https://codeforces.com/problemset/problem/356/C
https://codeforces.com/problemset/problem/460/D
https://codeforces.com/problemset/problem/796/C
https://codeforces.com/problemset/problem/1374/E2
https://codeforces.com/problemset/problem/1647/D
+构造 https://atcoder.jp/contests/arc153/tasks/arc153_c
https://atcoder.jp/contests/agc015/tasks/agc015_d

贡献法
LC979 https://leetcode.cn/problems/distribute-coins-in-binary-tree/ 1709
LC2477 https://leetcode.cn/problems/minimum-fuel-cost-to-report-to-the-capital/ 2012
LC891 https://leetcode.cn/problems/sum-of-subsequence-widths/
LC1588 https://leetcode.cn/problems/sum-of-all-odd-length-subarrays/
LC2681 https://leetcode.cn/problems/power-of-heroes/
- https://atcoder.jp/contests/arc116/tasks/arc116_b
LC2763 https://leetcode.cn/problems/sum-of-imbalance-numbers-of-all-subarrays/
更多贡献法题目，见 monotone_stack.go
https://codeforces.com/problemset/problem/1648/A 1400 维度独立
https://codeforces.com/problemset/problem/912/D
https://codeforces.com/problemset/problem/915/F
https://codeforces.com/problemset/problem/1208/E
https://codeforces.com/problemset/problem/1691/C
https://codeforces.com/problemset/problem/1777/D 树
https://codeforces.com/problemset/problem/1788/D 好题！
https://codeforces.com/problemset/problem/1789/C 好题！
https://codeforces.com/problemset/problem/1808/D
https://atcoder.jp/contests/abc290/tasks/abc290_e 好题！
https://atcoder.jp/contests/abc159/tasks/abc159_f 与 0-1 背包结合
^+ https://atcoder.jp/contests/abc201/tasks/abc201_e

小模拟
LC2534 https://leetcode.cn/problems/time-taken-to-cross-the-door/
https://atcoder.jp/contests/abc279/tasks/abc279_f

中模拟
https://atcoder.jp/contests/abc319/tasks/abc319_f

其他
删除一个字符 + 删除最长连续前缀 https://codeforces.com/problemset/problem/1430/D
https://codeforces.com/problemset/problem/521/D

合法括号字符串 Regular Bracket Sequence, RBS
https://codeforces.com/problemset/problem/1821/E
https://codeforces.com/problemset/problem/1830/C

= 变成 <= 或者 >=
求前缀和/后缀和
https://leetcode.cn/problems/maximum-product-of-the-length-of-two-palindromic-substrings/

连续性 + 上下界
https://atcoder.jp/contests/arc137/tasks/arc137_b
https://codeforces.com/contest/1695/problem/C
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

/* 二进制枚举
https://www.luogu.com.cn/problem/UVA11464 https://onlinejudge.org/index.php?option=com_onlinejudge&Itemid=8&category=26&page=show_problem&problem=2459
*/

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
等效性 LC1183 https://leetcode.cn/problems/maximum-number-of-ones/
LC1526 https://leetcode.cn/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
置换 https://atcoder.jp/contests/abc250/tasks/abc250_e
排序+最小操作次数 https://codeforces.com/contest/1367/problem/F2
https://codeforces.com/contest/1830/problem/A
从绝对值最大的开始思考 https://codeforces.com/contest/351/problem/E
https://codeforces.com/problemset/problem/777/C 1600

棋盘染色 LC2577 https://leetcode.cn/problems/minimum-time-to-visit-a-cell-in-a-grid/
        https://codeforces.com/contest/1848/problem/A
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
https://codeforces.com/problemset/problem/1792/C 1500
- 相似题目 https://codeforces.com/problemset/problem/1367/F1 2100
https://codeforces.com/problemset/problem/1882/B
https://codeforces.com/problemset/problem/712/C 1600
https://codeforces.com/problemset/problem/621/C 1700
https://codeforces.com/problemset/problem/1644/D 1700
https://codeforces.com/problemset/problem/1672/D 1700
https://codeforces.com/problemset/problem/1759/G 1900 求字典序最小，通常可以从大往小思考
https://codeforces.com/problemset/problem/1638/D 2000
https://codeforces.com/problemset/problem/571/A 2100
https://codeforces.com/problemset/problem/369/E 2200

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
LC1806 https://leetcode.cn/problems/minimum-number-of-operations-to-reinitialize-a-permutation/ 1491
*/

// 「恰好」转换成「至少/至多」https://codeforces.com/problemset/problem/1188/C

/* 反悔贪心
另见 heap.go 中的「反悔堆」
https://djy-juruo.blog.luogu.org/qian-tan-fan-hui-tan-xin
https://www.jvruo.com/archives/1844/
https://www.cnblogs.com/nth-element/p/11768155.html
题单 https://www.luogu.com.cn/training/8793
LC1388 双向链表反悔贪心 https://leetcode.cn/problems/pizza-with-3n-slices/
*/

/* 集合哈希
https://codeforces.com/problemset/problem/1394/B
https://www.luogu.com.cn/problem/P6688
*/

/* 操作树
和莫队类似，通过改变查询的顺序来优化复杂度
https://codeforces.com/problemset/problem/707/D
*/

/* Golang 卡常技巧（注：关于 IO 的加速见 io.go）
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

在特殊情况下，改为手动模拟栈可以减少 > 100MB 的内存
见这题的 Go 提交记录 https://codeforces.com/problemset/problem/163/E

测试：哈希表用时是数组的 13 倍（本题瓶颈）
slice    249ms https://codeforces.com/problemset/submission/570/209063267
hashmap 3259ms https://codeforces.com/problemset/submission/570/209063603
*/

// slice 作为 map 的 key
// 长度为 0 的 slice 对应空字符串
func intSliceAsMapKeyExample(cnt map[string]int, a []int) {
	// 如果后面还会修改 a，可以先 copy 一份
	//a = append(a[:0:0], a...)
	sh := (*reflect.SliceHeader)(unsafe.Pointer(&a))
	sh.Len *= bits.UintSize / 8 // 装作 byte slice
	s := *(*string)(unsafe.Pointer(sh))
	cnt[s]++
}

// 力扣上的 int 和 int64 是一样的，但是有些题目要求返回 []int64
// 此时可以用指针强转
func intsToInt64s(a []int) []int64 {
	int64s := *(*[]int64)(unsafe.Pointer(&a))
	return int64s
}

func _() {
	const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	pow10 := func(x int) int { return int(math.Pow10(x)) } // 不需要 round

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

	dir8 := []struct{ x, y int }{{1, 0}, {1, 1}, {0, 1}, {-1, 1}, {-1, 0}, {-1, -1}, {0, -1}, {1, -1}}  // 逆时针（坐标系）
	dir8 = []struct{ x, y int }{{-1, -1}, {-1, 0}, {-1, 1}, {0, 1}, {1, 1}, {1, 0}, {1, -1}, {0, -1}}   // 顺时针（矩阵）
	dir8 = []struct{ x, y int }{{2, 1}, {1, 2}, {-1, 2}, {-2, 1}, {-2, -1}, {-1, -2}, {1, -2}, {2, -1}} // 马走日

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
	cmp := func(a, b int) int {
		if a == b {
			return 0
		}
		if a < b {
			return -1
		}
		return 1
	}

	sort3 := func(a ...int) (x, y, z int) { slices.Sort(a); return a[0], a[1], a[2] }
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

	// 预处理所有回文数 https://oeis.org/A002113
	// LC2967 https://leetcode.cn/problems/minimum-cost-to-make-array-equalindromic/
	// LC906 https://leetcode.cn/problems/super-palindromes/
	// LC2081 https://leetcode.cn/problems/sum-of-k-mirror-numbers/
	// EXTRA: 单个数字的情况 LC564 https://leetcode.cn/problems/find-the-closest-palindrome/
	initPalindromeNumber := func() {
		const mx int = 1e9
		pal := []int{}
		// 哨兵。根据题目来定，也可以设置成 -2e9 等
		pal = append(pal, 0)

		// 严格按顺序从小到大生成所有回文数
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
	subSum := func(a []int) []int {
		sum := make([]int, 1<<len(a))
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

	// 前缀和应用：求距离和
	prefixSum := func(a []int) {
		slices.Sort(a)
		sum := make([]int, len(a)+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}

		// 返回 a 的所有数到 target 的距离之和，即 sum(abs(a[i]-target))
		// ！需要保证 a 是有序的
		// LC2602 https://leetcode.cn/problems/minimum-operations-to-make-all-array-elements-equal/
		// - 原题是 https://atcoder.jp/contests/abc255/tasks/abc255_d
		distanceSum := func(target int) int {
			i := sort.SearchInts(a, target)
			s1 := target*i - sum[i]
			s2 := sum[len(a)] - sum[i] - target*(len(a)-i)
			return s1 + s2
		}

		// LC2968 https://leetcode.cn/problems/apply-operations-to-maximize-frequency-score/
		// 返回下标在左闭右开区间 [left,right) 内的所有 a[i] 到 target 的距离之和
		// ！需要保证 a 是有序的
		distanceSumRange := func(left, right, target int) int {
			i := sort.SearchInts(a, target)
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

		// EXTRA: 青蛙跳井
		// 一次询问（下标从 1 开始）https://codeforces.com/problemset/problem/1141/E
		// 多次询问（下标从 0 开始）https://codeforces.com/problemset/problem/1490/G

		_ = []any{distanceSum, distanceSumRange}
	}

	// 同余前缀和，a 的下标从 0 开始，md 为模数
	// 求 a[i]+a[i+md]+a[i+2*md]+...
	// 具体用法见 query 上的注释
	// LC1664 https://leetcode.cn/problems/ways-to-make-a-fair-array/
	// LC2902 https://leetcode.cn/problems/count-of-sub-multisets-with-bounded-sum/
	// https://atcoder.jp/contests/abc288/tasks/abc288_d
	groupPrefixSum := func(_a []int, md int) {
		_sum := make([]int, len(_a)+md)
		for i, v := range _a {
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

	// 环形区间和 [l,r)
	circularRangeSum := func(a []int) {
		_n := len(a)
		sum := make([]int, _n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
		}
		pre := func(p int) int { return sum[_n]*(p/_n) + sum[p%_n] }
		// [l,r)
		query := func(l, r int) int { return pre(r) - pre(l) }

		_ = query
	}

	// 带权(等差数列)前缀和
	{
		var n int // read
		a := make([]int, n)
		// read a ...

		sum := make([]int, n+1)
		iSum := make([]int, n+1)
		for i, v := range a {
			sum[i+1] = sum[i] + v
			iSum[i+1] = iSum[i] + (i+1)*v
		}
		query := func(l, r int) int { return iSum[r] - iSum[l] - l*(sum[r]-sum[l]) } // [l,r)

		_ = query
	}

	// 二维前缀和     sum2d
	// LC221 https://leetcode.cn/problems/maximal-square/
	// LC1277 https://leetcode.cn/problems/count-square-submatrices-with-all-ones/
	// LC1504 https://leetcode.cn/problems/count-submatrices-with-all-ones/
	// 自加写法 https://codeforces.com/contest/835/submission/120031673
	// https://codeforces.com/contest/1107/problem/D
	// https://codeforces.com/problemset/problem/1731/D
	// https://codeforces.com/problemset/problem/611/C
	matrixSum := func(a [][]int) {
		n, m := len(a), len(a[0])
		sum := make([][]int, n+1)
		sum[0] = make([]int, m+1)
		for i, row := range a {
			sum[i+1] = make([]int, m+1)
			for j, v := range row {
				sum[i+1][j+1] = sum[i+1][j] + sum[i][j+1] - sum[i][j] + v
			}
		}
		// 类似前缀和的左闭右开
		// r1<=r<r2 && c1<=c<c2
		query := func(r1, c1, r2, c2 int) int {
			return sum[r2][c2] - sum[r2][c1] - sum[r1][c2] + sum[r1][c1]
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

	// 矩阵斜向前缀和
	// LC1878 https://leetcode.cn/problems/get-biggest-three-rhombus-sums-in-a-grid/
	diagonalSum := func(a [][]int) {
		n, m := len(a), len(a[0])
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
	contributionSum := func(a []int) (sum int) {
		n := len(a)
		slices.Sort(a)
		for i, v := range a {
			sum += v * (2*i + 1 - n)
		}
		return
	}

	// 差分数组
	// 请看 https://leetcode.cn/circle/discuss/FfMCgb/
	// - [1893. 检查是否区域内所有整数都被覆盖](https://leetcode.cn/problems/check-if-all-the-integers-in-a-range-are-covered/) 1307（暴力也可）
	// - [1094. 拼车](https://leetcode.cn/problems/car-pooling/) 1441
	// - [1109. 航班预订统计](https://leetcode.cn/problems/corporate-flight-bookings/) 1570
	// - [2406. 将区间分为最少组数](https://leetcode.cn/problems/divide-intervals-into-minimum-number-of-groups/) 1713
	// - [2381. 字母移位 II](https://leetcode.cn/problems/shifting-letters-ii/) 1793
	// - [995. K 连续位的最小翻转次数](https://leetcode.cn/problems/minimum-number-of-k-consecutive-bit-flips/) 1835
	// - [2772. 使数组中的所有元素都等于零](https://leetcode.cn/problems/apply-operations-to-make-all-array-elements-equal-to-zero/) 2029
	// - [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/) 2236
	// - [370. 区间加法](https://leetcode.cn/problems/range-addition/)（会员题）
	// https://codeforces.com/problemset/problem/816/B 1400
	// https://codeforces.com/problemset/problem/1700/C 1700
	// 浮点数差分（也可以用扫描线）https://atcoder.jp/contests/abc274/tasks/abc274_f
	// 差分思想 https://codeforces.com/problemset/problem/1634/F 2700

	// 二阶差分
	// https://ac.nowcoder.com/acm/contest/56446/C
	// https://www.luogu.com.cn/problem/U318099?contestId=123900
	// 简化 https://codeforces.com/problemset/problem/1661/D
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

		// 下标从 0 开始
		// a[j] += max(base-abs(i-j), 0)
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
		slices.Sort(xs)

		// 左闭右开区间 [_cnt[i].x, _cnt[i+1].x) 中的值都是 _cnt[i].c
		type _pair struct{ x, c int }
		_cnt := make([]_pair, len(xs))
		sd := 0
		for _, x := range xs {
			sd += diff[x]
			_cnt = append(_cnt, _pair{x, sd})
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
			sd := 0
			for _, x := range xs {
				sd += diff[x]
				_cnt[x] = sd
			}
		}

		_ = query
	}

	// 二维差分
	// https://blog.csdn.net/weixin_43914593/article/details/113782108
	// 模板题 LC2536 https://leetcode.cn/problems/increment-submatrices-by-one/
	// LC2132 https://leetcode.cn/problems/stamping-the-grid/（也可以不用差分）
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
	// LC1331 https://leetcode.cn/problems/rank-transform-of-an-array/
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

	// 另一种写法
	// 对比，相差不大（~10%，Go 1.14.1）
	// discrete  296ms/14952KB https://atcoder.jp/contests/abc221/submissions/35791381
	// discrete2 333ms/11748KB https://atcoder.jp/contests/abc221/submissions/35791225
	discrete2 := func(a []int, startIndex int) []int {
		b := slices.Clone(a)
		slices.Sort(b)
		b = slices.Compact(b)
		for i, v := range a {
			a[i] = sort.SearchInts(b, v) + startIndex
		}
		return a
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
		rd := rand.New(rand.NewSource(time.Now().UnixNano()))
		rd.Shuffle(len(a), func(i, j int) { a[i], a[j] = a[j], a[i] })
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
	// LC 套题 https://leetcode.cn/tag/line-sweep/
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
	// 这样可以避免写一个复杂的 sort.Slice
	sweepLine2 := func(ranges [][]int) {
		n := len(ranges)
		events := make([]int, 0, 2*n)
		for _, p := range ranges {
			l, r := p[0], p[1]
			// 注意移位后是否溢出
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

	_ = []any{
		alphabet,
		pow10, dir4, dir8, perm3, perm4,
		abs, ceil, cmp,
		ternaryI, ternaryS, mergeMap, xorSet, rotateCopy, transpose,
		toAnyBase, digits, initPalindromeNumber,
		subSum, recoverArrayFromSubsetSum, subSumSorted,
		prefixSum, groupPrefixSum, circularRangeSum, matrixSum, rowColSum, diagonalSum,
		contributionSum,
		diffOfDiff, diffMap, diff2D,
		sort3,
		merge, mergeWithLimit, splitDifferenceAndIntersection, intersection, isSubset, isSubSequence, isDisjoint,
		discreteSimple, discrete, discrete2, discreteMap,
		indexMap, complement, quickSelect,
		sweepLine, sweepLine2, countCoveredPoints,
		discrete2D,
	}
}
