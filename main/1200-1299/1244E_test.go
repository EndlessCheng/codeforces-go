package main

import (
	"github.com/EndlessCheng/codeforces-go/main/testutil"
	"testing"
)

func TestSol1244E(t *testing.T) {
	// just copy from website
	rawText := `
70 3956
246 495 357 259 209 422 399 443 252 537 524 299 538 234 247 558 527 529 153 366 453 415 476 410 144 472 346 125 299 321 363 334 297 316 346 309 497 281 163 396 482 254 447 318 316 444 308 332 508 505 328 287 450 557 265 199 298 240 258 232 424 229 292 196 150 281 321 234 443 282
outputCopy
92
inputCopy
4 5
3 1 7 5
outputCopy
2
inputCopy
3 10
100 100 100
outputCopy
0
inputCopy
10 9
4 5 5 7 5 4 5 2 4 3
outputCopy
1`
	testutil.AssertEqualCase(t, rawText, -1, Sol1244E)
}
