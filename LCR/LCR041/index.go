package LCR041

type MovingAverage struct {
	cache []int
	size  int
}

/** Initialize your data structure here. */
func Constructor(size int) MovingAverage {
	return MovingAverage{
		cache: make([]int, 0),
		size:  size,
	}
}

func (this *MovingAverage) Next(val int) float64 {
	if len(this.cache) < this.size {
		this.cache = append(this.cache, val)
		return float64(sum(this.cache)) / float64(len(this.cache))
	}

	this.cache = append(this.cache[1:], val)
	return float64(sum(this.cache)) / float64(this.size)
}

func sum(nums []int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result
}

/**
 * Your MovingAverage object will be instantiated and called as such:
 * obj := Constructor(size);
 * param_1 := obj.Next(val);
 */
