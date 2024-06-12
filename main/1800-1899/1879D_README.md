我将带你一步步思考本题。

首先，完成

560. 和为 K 的子数组
     https://leetcode.cn/problems/subarray-sum-equals-k/

我的题解
https://leetcode.cn/problems/subarray-sum-equals-k/solutions/2781031/qian-zhui-he-ha-xi-biao-cong-liang-ci-bi-4mwr/

在我的题解中，第三个思考题是：
计算和为 k 的所有子数组的长度之和。

这也可以用前缀和来算，即所有满足 s[r]-s[l] = k 的 r-l 之和。
枚举 r，假如有三个 l 满足 s[l] = s[r] - k，那么 r-l 之和就是
(r-l1) + (r-l2) + (r-l3)
= r * 3 - (l1 + l2 + l3)
所以我们需要维护两类数据：
1. s[i] 的出现次数，即上式中的 3。
2. s[i] 的 i 之和，也就是值等于 s[i] 的下标之和，即上式中的 l1 + l2 + l3。

回到本题，计算异或前缀和，问题变成 (s[r]^s[l])*(r-l) 之和（利用半闭半开区间，把原题公式中的 +1 去掉）。
拆位，考虑每一位的贡献。
最低位，0 没有贡献，只需要考虑 1，即所有 s[r]^s[l] 最低位 = 1 的 r-l 之和。
这正是我们前面思考的「和为 k 的所有子数组的长度之和」！
次低位，0 没有贡献，只需要考虑 1，即所有 s[r]^s[l] 次低位 = 1 的 r-l 之和，再乘上 1 << 1。
第三位，0 没有贡献，只需要考虑 1，即所有 s[r]^s[l] 第三位 = 1 的 r-l 之和，再乘上 1 << 2。
依此类推，一直到：所有 ((s[r]^s[l]) >> 29 & 1) == 1 的 r-l 之和，再乘上 1 << 29。

注意取模。

https://codeforces.com/contest/1879/submission/264665078