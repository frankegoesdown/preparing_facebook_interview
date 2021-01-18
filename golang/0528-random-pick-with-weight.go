package main

import (
	"math/rand"
	"time"
)

func main() {

}

type Solution struct {
	sums []int
}

func Constructor(w []int) Solution {
	rand.Seed(time.Now().UnixNano())
	sums := make([]int, len(w))
	sums[0] = w[0]
	for i := 1; i < len(w); i++ {
		sums[i] = w[i] + sums[i-1]
	}
	return Solution{sums}
}

func (this *Solution) PickIndex() int {
	pick := rand.Intn(this.sums[len(this.sums)-1])
	l, r := 0, len(this.sums)-1
	for l <= r {
		mid := l + (r-l)/2
		if this.sums[mid] > pick {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return l
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(w);
 * param_1 := obj.PickIndex();
 */