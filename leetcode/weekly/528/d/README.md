做法同 [3878. 统计好子数组](https://leetcode.cn/problems/count-good-subarrays/)，请看 [我的题解](https://leetcode.cn/problems/count-good-subarrays/solutions/3933380/mo-ban-logtrick-ji-lu-mei-ge-yuan-su-de-otgcv/)。

只需把 3878 代码中的 `|` 改成 `&`，即可通过本题。

**问**：为什么这两题做法是一样的？

**答**：虽然 3878 只要求「按位或在子数组中」，但由于按位或的结果必然 $\ge $ 数组最大值，所以该要求等价于「子数组的按位或等于其最大值」。本题是镜像版本，按位与的结果必然 $\le $ 数组最小值。
