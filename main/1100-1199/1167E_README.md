# CF1167E 题解

本题是【2972. 统计移除递增子数组的数目 II】
https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii/
的变形题。
由于思路和代码都非常接近，推荐先把力扣那题做了。

核心思路：删除元素后，对于剩余的每个元素，考察其出现位置的区间，这些区间必须不相交。

记录每个元素首次和最后一次出现的位置。
定义 ps[v].l 表示元素 v 首次出现的位置，ps[v].r 表示元素 v 最后一次出现的位置。

用闭区间 [ps[v].l, ps[v].r] 表示元素 v。
对于剩余元素中的两个数 v 和 w，如果 v < w，那么区间 [ps[v].l, ps[v].r] 必须在区间 [ps[w].l, ps[w].r] 的左侧，即 ps[v].r < ps[w].l。

首先考虑一个简单情况：1 到 x 中的每个数都在数组 a 中。
和力扣 2972 题代码一样，双指针，把「元素值比大小」改成 ps[v].r < ps[w].l 这样的「区间左右端点比大小」即可。

本题的一个难点是，如果有元素不在 a 中，怎么算清楚这种情况，尤其是不能重复统计。
首先把 ps 中的元素值在 a 中的区间提取出来，得到数组 b。
其中 b[i].l 和 b[i].r 是第 i 个区间的左右端点，b[i].v 是第 i 个区间对应的元素值。
看上去在 b 数组上「区间左右端点比大小」就行，怎么把不在 a 中的元素值考虑进来呢？

和 2972 题我的题解
https://leetcode.cn/problems/count-the-number-of-incremovable-subarrays-ii/solution/shuang-zhi-zhen-on-shi-jian-o1-kong-jian-2hsz/
一样，首先讨论删除后缀的情况，然后讨论一般情况。

删除后缀 b[k] 到 b[n-1]，其中 n 是 b 的长度，k <= i+1。
那么删除的范围，最小值 <= b[i+1].v，最大值 >= b[n-1].v，根据乘法原理，这一共有 b[i+1].v * (x + 1 - b[n-1].v) 个。

然后讨论删除的不是后缀的情况。
删除 b[k] 到 b[j-1]，保留 b[j] 到 b[n-1]，那么删除的范围，最小值 <= b[i+1].v，最大值在左闭右开区间 [b[j-1].v, b[j].v) 中，根据乘法原理，这一共有 b[i+1].v * (b[j].v - b[j-1].v) 个。

https://codeforces.com/contest/1167/submission/270260662
