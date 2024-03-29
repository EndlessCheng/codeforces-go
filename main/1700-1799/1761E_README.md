分类讨论题。

按照如下顺序，依次判断：
1. 如果只有一个连通块，那么无需操作，输出 0。
2. 如果存在一个孤立点，那么只需对这个点操作，即可让整个图连通。
3. 如果存在一个不是团（完全子图）的连通块，那么只需对这个连通块的度最小的点操作。证明见下。
4. 如果有至少 3 个连通块（都是团），那么只需操作两次，随便选两个在不同连通块的点。
5. 如果只有两个连通块（都是团），那么选其中点数少的连通块操作，操作次数就是该连通块的点数。

3 的证明：
设这个度最小的点是 x，设 x 所在连通分量为 V。由于 V 不是团，所以 V 中的某些点在操作前不与 x 相连，在操作后会与 x 相连。
如果 x 不是割点，那么操作后，V 不会分成更多的连通块，并且仍然有点连着 x，所以整个图是连通的。
如果 x 是割点（例如两个三角形各有一个顶点连着 x），首先来说明几个性质：
性质 1：由于 x 的度数至少是 2（割点性质），所以其余每个点的度数也至少是 2，这意味着 V 中没有「叶子」。
性质 2：删除 x 会把 V 分成若干个连通块，每个连通块至少有 3 个点（根据性质 1）。
现在的问题是：操作后，每个连通块是否都有点与 x 相连？
反证法：假设有一个连通块 U 没有点与 x 相连，设这个连通块的大小为 |U|。
由于操作前 x 除了与这 |U| 个点相连外，还与其它点相连（注意我们假设 x 是割点），这说明 x 的度数至少是 |U|+1，但由于 x 的度数最小，所以 |U| 中必定有度数至少为 |U|+1 的点，但 U 中只有 |U| 个点，度数不可能超过 |U|，矛盾。所以操作后，每个连通块都至少有一个点与 x 相连。
