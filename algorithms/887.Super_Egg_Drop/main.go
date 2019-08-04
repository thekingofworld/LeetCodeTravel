package main

import (
	"fmt"
	"math"
)

func main() {
	x := superEggDrop(2, 100)
	fmt.Println(x)
}

//算法一 复杂度：O(K*N*N)
func superEggDrop1(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	fPre := make([]int, N+1)
	fCur := make([]int, N+1)
	for i := range fPre {
		fPre[i] = i
		fCur[i] = i
	}
	for i := 2; i <= K; i++ {
		for e := 1; e <= N; e++ {
			fCur[e] = math.MaxInt64
			for j := 1; j <= e; j++ {
				fCur[e] = min(fCur[e], max(fPre[j-1], fCur[e-j]) + 1)
			}
		}
		for e := 1; e <= N; e++ {
			fPre[e] = fCur[e]
		}
	}
	return fCur[N]
}

//算法二 复杂度：O(N*N*log2(N))
func superEggDrop2(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	t := int(math.Ceil(math.Log2(float64(N+1))))
	if K >= t {
		return t
	}
	fPre := make([]int, N+1)
	fCur := make([]int, N+1)
	for i := range fPre {
		fPre[i] = i
		fCur[i] = i
	}
	for i := 2; i <= K; i++ {
		for e := 1; e <= N; e++ {
			t := int(math.Ceil(math.Log2(float64(e+1))))
			if i >= t {
				fCur[e] = t
				continue
			}
			fCur[e] = math.MaxInt64
			for j := 1; j <= e; j++ {
				fCur[e] = min(fCur[e], max(fPre[j-1], fCur[e-j]) + 1)
			}
		}
		for e := 1; e <= N; e++ {
			fPre[e] = fCur[e]
		}
	}
	return fCur[N]
}

//算法三 复杂度：O(N*log2(N)*log2(N))
func superEggDrop3(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	t := int(math.Ceil(math.Log2(float64(N+1))))
	if K >= t {
		return t
	}
	fPre := make([]int, N+1)
	fCur := make([]int, N+1)
	for i := range fPre {
		fPre[i] = i
		fCur[i] = i
	}
	for i := 2; i <= K; i++ {
		for e := 1; e <= N; e++ {
			t := int(math.Ceil(math.Log2(float64(e+1))))
			if i >= t {
				fCur[e] = t
				continue
			}
			fCur[e] = math.MaxInt64
			start := 1
			end := e
			for start <= end {
				mid := (start + end)/2
				if fPre[mid-1] < fCur[e-mid] {
					fCur[e] = min(fCur[e], fCur[e-mid] + 1)
					start = mid + 1
				} else if fPre[mid-1] > fCur[e-mid] {
					fCur[e] = min(fCur[e], fPre[mid-1] + 1)
					end = mid - 1
				} else {
					fCur[e] = min(fCur[e], fPre[mid-1] + 1)
					break
				}
			}
		}
		for e := 1; e <= N; e++ {
			fPre[e] = fCur[e]
		}
	}
	return fCur[N]
}

//算法四 复杂度：O(N*log2(N))
func superEggDrop(K int, N int) int {
	if K < 1 || N < 1 {
		return 0
	}
	t := int(math.Ceil(math.Log2(float64(N+1))))
	if K >= t {
		return t
	}
	fPre := make([]int, N+1)
	fCur := make([]int, N+1)
	for i := range fPre {
		fPre[i] = i
		fCur[i] = i
	}
	for i := 2; i <= K; i++ {
		p := 0
		for e := 2; e <= N; e++ {
			t := int(math.Ceil(math.Log2(float64(e+1))))
			if i >= t {
				fCur[e] = t
				continue
			}
			if fCur[p] >= fPre[e-p-1] {
				fCur[e] = fCur[e-1]
			} else {
				fCur[e] = fCur[e-1] + 1
				p = e - 1
			}
		}
		for e := 1; e <= N; e++ {
			fPre[e] = fCur[e]
		}
	}
	return fCur[N]
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}