我从上半年的周赛中精选了部分题目，分类整理后得到了下面的表格（按题号排序），供各位练习/复习时参考。

另外我更新了表中一部分题解，完善了多种解法（例如 [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/solution/er-fen-bfspythonjavacgo-by-endlesscheng-ypp1/) 添加了不用二分的线性做法），且所有题解均包含 Python/Java/C++/Go 四种语言。

## 思维题

| 题目                                                                              | 题解                                                                                                                    | 难度                                                         | 备注     |
|-----------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|--------|
| [2211. 统计道路上的碰撞次数](https://leetcode.cn/problems/count-collisions-on-a-road/)      | [题解](https://leetcode.cn/problems/count-collisions-on-a-road/solution/jie-lun-ti-san-xing-gao-ding-by-endlessc-bvnw/) | 1581                                                       | 脑筋急转弯  |
| [2227. 加密解密字符串](https://leetcode.cn/problems/encrypt-and-decrypt-strings/)        | [题解](https://leetcode.cn/problems/encrypt-and-decrypt-strings/solution/by-endlesscheng-sm8h/)                         | 1944                                                       | 逆向思维   |
| [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/) | [题解](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/solution/by-endlesscheng-dt8h/)                    | 2304                                                       | 有技巧的枚举 |
| [2306. 公司命名](https://leetcode.cn/problems/naming-a-company/)                      | [题解](https://leetcode.cn/problems/naming-a-company/solution/by-endlesscheng-ruz8/) \| [视频](https://www.bilibili.com/video/BV1aT41157bh?t=14m10s) | 2305   |分类讨论|
| [2317. 操作后的最大异或和](https://leetcode.cn/problems/maximum-xor-after-operations/)     | [题解](https://leetcode.cn/problems/maximum-xor-after-operations/solution/yi-bu-bu-tis-by-endlesscheng-89kw/) \| [视频](https://www.bilibili.com/video/BV143411w7Wx?t=6m14s) | 1678   |位运算|


## 贪心

|题目| 题解                                                                                                                                             | 难度                                                         |备注|
|---|------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|---|
|[2136. 全部开花的最早一天](https://leetcode.cn/problems/earliest-possible-day-of-full-bloom/)| [题解](https://leetcode.cn/problems/earliest-possible-day-of-full-bloom/solution/tan-xin-ji-qi-zheng-ming-by-endlesscheng-hfwe/)                 | 2033                                                       |邻项交换法|
|[2141. 同时运行 N 台电脑的最长时间](https://leetcode.cn/problems/maximum-running-time-of-n-computers/)| [题解](https://leetcode.cn/problems/maximum-running-time-of-n-computers/solution/liang-chong-jie-fa-er-fen-da-an-pai-xu-t-grd8/)                 | 2265                                                       |贪心|
|[2234. 花园的最大总美丽值](https://leetcode.cn/problems/maximum-total-beauty-of-the-gardens/)| [题解](https://leetcode.cn/problems/maximum-total-beauty-of-the-gardens/solution/by-endlesscheng-10i7/)                                          | 2561                                                       |贪心|
|[2311. 小于等于 K 的最长二进制子序列](https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/)| [题解](https://leetcode.cn/problems/longest-binary-subsequence-less-than-or-equal-to-k/solution/fen-lei-tao-lun-tan-xin-by-endlesscheng-vnlx/) \| [视频](https://www.bilibili.com/video/BV1CW4y1k7B3?t=14m56s) |1839|贪心|


## 双指针

|题目|题解|难度|备注|
|---|---|---|---|
|[2271. 毯子覆盖的最多白色砖块数](https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/)|[题解](https://leetcode.cn/problems/maximum-white-tiles-covered-by-a-carpet/solution/by-endlesscheng-kdy9/)|2021|双指针|
|[2302. 统计得分小于 K 的子数组数目](https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/)|[题解](https://leetcode.cn/problems/count-subarrays-with-score-less-than-k/solution/by-endlesscheng-b120/)|1808|双指针|


## 二分

|题目|题解|难度|备注|
|---|---|---|---|
|[2141. 同时运行 N 台电脑的最长时间](https://leetcode.cn/problems/maximum-running-time-of-n-computers/)|[题解](https://leetcode.cn/problems/maximum-running-time-of-n-computers/solution/liang-chong-jie-fa-er-fen-da-an-pai-xu-t-grd8/)|2265|二分答案|
|[2251. 花期内花的数目](https://leetcode.cn/problems/number-of-flowers-in-full-bloom/)|[题解](https://leetcode.cn/problems/number-of-flowers-in-full-bloom/solution/chai-fen-pythonjavacgo-by-endlesscheng-wz35/)|2022|转换|
|[2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/)|[题解](https://leetcode.cn/problems/escape-the-spreading-fire/solution/er-fen-bfspythonjavacgo-by-endlesscheng-ypp1/)|2346|二分答案|


## 数学

| 题目                                                                                       | 题解                                                                                                                          |难度|备注|
|------------------------------------------------------------------------------------------|-----------------------------------------------------------------------------------------------------------------------------|---|---|
| [2183. 统计可以被 K 整除的下标对数目](https://leetcode.cn/problems/count-array-pairs-divisible-by-k/) | [题解](https://leetcode.cn/problems/count-array-pairs-divisible-by-k/solution/tong-ji-yin-zi-chu-xian-ci-shu-by-endles-t5k8/) |2246|数论|
| [2245. 转角路径的乘积中最多能有几个尾随零](https://leetcode.cn/problems/maximum-trailing-zeros-in-a-cornered-path/) | [题解](https://leetcode.cn/problems/maximum-trailing-zeros-in-a-cornered-path/solution/by-endlesscheng-7z5a/) |2036|数论、前缀和|
|[2281. 巫师的总力量和](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/)| [题解](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/solution/dan-diao-zhan-qian-zhui-he-de-qian-zhui-d9nki/) \| [视频](https://www.bilibili.com/video/BV1RY4y157nW?t=17m0s)                                                     | 2621                                                       |和式变形、前缀和|
| [2310. 个位数字为 K 的整数之和](https://leetcode.cn/problems/sum-of-numbers-with-units-digit-k/)   | [题解](https://leetcode.cn/problems/sum-of-numbers-with-units-digit-k/solution/mei-ju-da-an-by-endlesscheng-zh75/) \| [视频](https://www.bilibili.com/video/BV1CW4y1k7B3?t=7m11s)           |1558|同余|


## 动态规划

|题目| 题解                                                                                                                                               | 难度                                                         | 备注      |
|---|--------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|---------|
|[2140. 解决智力问题](https://leetcode.cn/problems/solving-questions-with-brainpower/)| [题解](https://leetcode.cn/problems/solving-questions-with-brainpower/solution/dao-xu-dp-by-endlesscheng-2qkc/)                                    | 1709                                                       | 线性 DP   |
|[2167. 移除所有载有违禁货物车厢所需的最少时间](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/)| [题解](https://leetcode.cn/problems/minimum-time-to-remove-all-cars-containing-illegal-goods/solution/qian-hou-zhui-fen-jie-dp-by-endlesscheng-6u1b/) | 2219                                                       | 线性 DP   |
|[2172. 数组的最大与和](https://leetcode.cn/problems/maximum-and-sum-of-array/)| [题解](https://leetcode.cn/problems/maximum-and-sum-of-array/solution/zhuang-tai-ya-suo-dp-by-endlesscheng-5eqn/)                                  | 2392                                                       | 状压 DP   |
|[2188. 完成比赛的最少时间](https://leetcode.cn/problems/minimum-time-to-finish-the-race/)| [题解](https://leetcode.cn/problems/minimum-time-to-finish-the-race/solution/jie-he-xing-zhi-qiao-miao-dp-by-endlessc-b963/)                       | 2315                                                       | 线性 DP   |
|[2209. 用地毯覆盖后的最少白色砖块](https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/)| [题解](https://leetcode.cn/problems/minimum-white-tiles-after-covering-with-carpets/solution/by-endlesscheng-pa3v/)                                | 2105                                                       | 线性 DP   |
|[2218. 从栈中取出 K 个硬币的最大面值和](https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/)| [题解](https://leetcode.cn/problems/maximum-value-of-k-coins-from-piles/solution/zhuan-hua-cheng-fen-zu-bei-bao-pythongoc-3xnk/)                   | 2157                                                       | 分组背包    |
|[2246. 相邻字符不同的最长路径](https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/)| [题解](https://leetcode.cn/problems/longest-path-with-different-adjacent-characters/solution/by-endlesscheng-92fw/)                                | 2126                                                       | 树形 DP   |
|[2262. 字符串的总引力](https://leetcode.cn/problems/total-appeal-of-a-string/)| [题解](https://leetcode.cn/problems/total-appeal-of-a-string/solution/by-endlesscheng-g405/)                                                       | 2033                                                       | 线性 DP   |
|[2266. 统计打字方案数](https://leetcode.cn/problems/count-number-of-texts/)| [题解](https://leetcode.cn/problems/count-number-of-texts/solution/by-endlesscheng-gj8f/)                                                          | 1856                                                       | 线性 DP   |
|[2272. 最大波动的子字符串](https://leetcode.cn/problems/substring-with-largest-variance/)| [题解](https://leetcode.cn/problems/substring-with-largest-variance/solution/by-endlesscheng-5775/)                                                | 2515                                                       | 线性 DP   |
|[2305. 公平分发饼干](https://leetcode.cn/problems/fair-distribution-of-cookies/)| [题解](https://leetcode.cn/problems/fair-distribution-of-cookies/solution/by-endlesscheng-80ao/) \| [视频](https://www.bilibili.com/video/BV1aT41157bh)          | 1886    |子集状压 DP|
|[2312. 卖木头块](https://leetcode.cn/problems/selling-pieces-of-wood/)| [题解](https://leetcode.cn/problems/selling-pieces-of-wood/solution/by-endlesscheng-mrmd/) \| [视频](https://www.bilibili.com/video/BV1CW4y1k7B3?t=23m21s) | 2363    |线性 DP|
|[2318. 不同骰子序列的数目](https://leetcode.cn/problems/number-of-distinct-roll-sequences/)| [题解](https://leetcode.cn/problems/number-of-distinct-roll-sequences/solution/by-endlesscheng-tgkn/) \| [视频](https://www.bilibili.com/video/BV143411w7Wx?t=11m33s) | 2090    |线性 DP|
|[2320. 统计放置房子的方式数](https://leetcode.cn/problems/count-number-of-ways-to-place-houses/)| [题解](https://leetcode.cn/problems/count-number-of-ways-to-place-houses/solution/d-by-endlesscheng-gybx/) \| [视频](https://www.bilibili.com/video/BV1pW4y1r7xs) | 1607    |线性 DP|
|[2321. 拼接数组的最大分数](https://leetcode.cn/problems/maximum-score-of-spliced-array/)| [题解](https://leetcode.cn/problems/maximum-score-of-spliced-array/solution/by-endlesscheng-fm8l/) \| [视频](https://www.bilibili.com/video/BV1pW4y1r7xs?t=8m18s) | 1790    |线性 DP|
|[LCP 53. 守护太空城](https://leetcode.cn/problems/EJvmW4/)| [题解](https://leetcode.cn/problems/EJvmW4/solution/by-endlesscheng-pk2q/)                                                                         | -                                                          | 子集状压 DP |



## 数据结构

|题目| 题解                                                                                                                                             | 难度                                                         |备注|
|---|------------------------------------------------------------------------------------------------------------------------------------------------|------------------------------------------------------------|---|
|[2157. 字符串分组](https://leetcode.cn/problems/groups-of-strings/)| [题解](https://leetcode.cn/problems/groups-of-strings/solution/bing-cha-ji-wei-yun-suan-by-endlesscheng-uejd/)                                   | 2499                                                       |哈希并查集|
|[2163. 删除元素后和的最小差值](https://leetcode.cn/problems/minimum-difference-in-sums-after-removal-of-elements/)| [题解](https://leetcode.cn/problems/minimum-difference-in-sums-after-removal-of-elements/solution/qian-zhui-zui-xiao-he-hou-zhui-zui-da-he-yz3d/) | 2225                                                       |堆|
|[2179. 统计数组中好三元组数目](https://leetcode.cn/problems/count-good-triplets-in-an-array/)| [题解](https://leetcode.cn/problems/count-good-triplets-in-an-array/solution/deng-jie-zhuan-huan-shu-zhuang-shu-zu-by-xmyd/)                     | 	2272                                                      |树状数组|
|[2197. 替换数组中的非互质数](https://leetcode.cn/problems/replace-non-coprime-numbers-in-array/)| [题解](https://leetcode.cn/problems/replace-non-coprime-numbers-in-array/solution/li-yong-zhan-mo-ni-gocpythonjava-by-endl-bnbv/)                | 2057                                                       |栈|
|[2213. 由单个字符重复的最长子字符串](https://leetcode.cn/problems/longest-substring-of-one-repeating-character/)| [题解](https://leetcode.cn/problems/longest-substring-of-one-repeating-character/solution/by-endlesscheng-qpbw/)                                 | 2628                                                       |线段树|
|[2276. 统计区间中的整数数目](https://leetcode.cn/problems/count-integers-in-intervals/)| [题解](https://leetcode.cn/problems/count-integers-in-intervals/solution/by-endlesscheng-clk2/)                                                  | 2222                                                       |珂朵莉树、动态开点线段树|
|[2281. 巫师的总力量和](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/)| [题解](https://leetcode.cn/problems/sum-of-total-strength-of-wizards/solution/dan-diao-zhan-qian-zhui-he-de-qian-zhui-d9nki/) \| [视频](https://www.bilibili.com/video/BV1RY4y157nW?t=17m0s)                                                     | 2621                                                       |单调栈|
|[2286. 以组为单位订音乐会的门票](https://leetcode.cn/problems/booking-concert-tickets-in-groups/)| [题解](https://leetcode.cn/problems/booking-concert-tickets-in-groups/solution/by-endlesscheng-okcu/) \| [视频](https://www.bilibili.com/video/BV18t4y1p736?t=8m7s)   | 2470                                                       |线段树二分|
|[2289. 使数组按非递减顺序排列](https://leetcode.cn/problems/steps-to-make-array-non-decreasing/)| [题解](https://leetcode.cn/problems/steps-to-make-array-non-decreasing/solution/by-endlesscheng-s2yc/) \| [视频](https://www.bilibili.com/video/BV1iF41157dG)          | 2481   |单调栈|
|[2296. 设计一个文本编辑器](https://leetcode.cn/problems/design-a-text-editor/)| [题解](https://leetcode.cn/problems/design-a-text-editor/solution/lian-biao-mo-ni-pythonjavacgo-by-endless-egw4/) \| [视频](https://www.bilibili.com/video/BV1w34y1L7yu/?t=6m29s) |1911|对顶栈|


## 图论

| 题目                                                                                                      | 题解                                                                                                                                           | 难度     | 备注              |
|---------------------------------------------------------------------------------------------------------|----------------------------------------------------------------------------------------------------------------------------------------------|--------|-----------------|
| [2127. 参加会议的最多员工数](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/)          | [题解](https://leetcode.cn/problems/maximum-employees-to-be-invited-to-a-meeting/solution/nei-xiang-ji-huan-shu-tuo-bu-pai-xu-fen-c1i1b/)      | 2449   | 基环树、拓扑排序 |
| [2146. 价格范围内最高排名的 K 样物品](https://leetcode.cn/problems/k-highest-ranked-items-within-a-price-range/)          | [题解](https://leetcode.cn/problems/k-highest-ranked-items-within-a-price-range/solution/fen-ceng-bfs-de-tong-shi-pai-xu-by-endle-ash6/)      | 1836   | BFS |
| [2172. 数组的最大与和](https://leetcode.cn/problems/maximum-and-sum-of-array/)                                 | [题解](https://leetcode.cn/problems/maximum-and-sum-of-array/solution/zhuang-tai-ya-suo-dp-by-endlesscheng-5eqn/)                              | 2392   | 最小费用最大流         |
| [2203. 得到要求路径的最小带权子图](https://leetcode.cn/problems/minimum-weighted-subgraph-with-the-required-paths/)  | [题解](https://leetcode.cn/problems/minimum-weighted-subgraph-with-the-required-paths/solution/by-endlesscheng-2mxm/)                          | 2364   | 单源最短路           |
| [2258. 逃离火灾](https://leetcode.cn/problems/escape-the-spreading-fire/)                                   | [题解](https://leetcode.cn/problems/escape-the-spreading-fire/solution/er-fen-bfspythonjavacgo-by-endlesscheng-ypp1/)                          | 2346   | 多源 BFS          |
| [2267. 检查是否有合法括号字符串路径](https://leetcode.cn/problems/check-if-there-is-a-valid-parentheses-string-path/) | [题解](https://leetcode.cn/problems/check-if-there-is-a-valid-parentheses-string-path/solution/tian-jia-zhuang-tai-hou-dfscpythonjavago-f287/) | 2084   | DFS             |
| [2290. 到达角落需要移除障碍物的最小数目](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/)        | [题解](https://leetcode.cn/problems/minimum-obstacle-removal-to-reach-corner/solution/0-1-bfs-by-endlesscheng-4pjt/) \| [视频](https://www.bilibili.com/video/BV1iF41157dG?t=21m10s) | 2137            |0-1 BFS|
| [2322. 从树中删除边的最小分数](https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/)        | [题解](https://leetcode.cn/problems/minimum-score-after-removals-on-a-tree/solution/dfs-shi-jian-chuo-chu-li-shu-shang-wen-t-x1kk/) \| [视频](https://www.bilibili.com/video/BV1pW4y1r7xs?t=19m37s) | 2391            |DFS 时间戳|




## 总结

相比往年的周赛，今年上半年，动态规划题略有减少，数据结构题有所增多，去年的数据结构题中经常遇到并查集，今年见到了不少线段树的题目。

不过动态规划依旧是周赛常客，对于目标周赛四题的选手，我仍然推荐将动态规划作为平时练习中的重点。

最后欢迎点赞评论，祝各位周周四题！
