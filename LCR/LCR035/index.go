package LCR035

import "strconv"

// 给定一个 24 小时制（小时:分钟 "HH:MM"）的时间列表，找出列表中任意两个时间的最小时间差并以分钟数表示。
/*func findMinDifference(timePoints []string) int {

	// 将时间转化为分钟
	minutes := make([]int, len(timePoints))
	for i := 0; i < len(timePoints); i++ {
		v := timePoints[i]
		// 将 HH:MM 转化为分钟
		hour := int(v[0]-'0')*10 + int(v[1]-'0')
		minute := int(v[3]-'0')*10 + int(v[4]-'0')

		minutes[i] = hour*60 + minute
		// TODO: 这里可以使用插入排序
	}

	result := 65536
	for i := 0; i < len(minutes); i++ {
		for j := i + 1; j < len(minutes); j++ {
			abs1 := minutes[i]
			abs2 := minutes[j]
			if abs1 > abs2 {
				abs1, abs2 = abs2, abs1
			}
			sub := abs2 - abs1

			subR := (1440 - abs2) + abs1

			if subR < sub {
				sub = subR
			}

			if sub < result {
				result = sub
			}
		}
	}

	return result
}*/

func findMinDifference(timePoints []string) int {
	marked := make([]bool, 1440)
	for _, timePoint := range timePoints {
		hour, _ := strconv.Atoi(timePoint[:2])
		minute, _ := strconv.Atoi(timePoint[3:])
		index := hour*60 + minute
		if marked[index] { // 如果有时间点出现两次，最小时间差就是0
			return 0
		}
		marked[index] = true
	}

	minVal, maxVal, prev, minDiff := 1440, 0, -1, 1440
	for i := 0; i < 1440; i++ {
		if marked[i] {
			if prev != -1 && i-prev < minDiff {
				minDiff = i - prev
			}
			minVal = minInt(minVal, i)
			maxVal = i
			prev = i
		}
	}

	// 考虑跨越午夜的情况
	return minInt(minDiff, minVal+1440-maxVal)
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}
