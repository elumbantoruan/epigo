package ch14sorting

import (
	"math"
	"sort"
)

type Interval struct {
	Start int
	End   int
}

// MergingInterval merges existing sorted intervals with a new one
//
// Existing intervals: {-4,-1}{0,2}{3,6}{7,9}{11,12}{14,17}
// New interval: {1,8}
// Result: {-4,-1}{0,9}{11,12}{14,17}
//
func MergingInterval(existingIntervals []Interval, newInterval Interval) []*Interval {
	combined := []Interval{}
	//combinedMap := make(map[int]Interval)
	//i := 0
	//for i < len(existingIntervals) {
	//	if newInterval.Start < existingIntervals[i].Start {
	//		_, ok := combinedMap[newInterval.Start]
	//		if !ok {
	//			combinedMap[newInterval.Start] = newInterval
	//		} else {
	//			combinedMap[existingIntervals[i].Start] = existingIntervals[i]
	//			i++
	//		}
	//	} else {
	//		combinedMap[existingIntervals[i].Start] = existingIntervals[i]
	//		i++
	//	}
	//}
	//
	//for _, value := range combinedMap {
	//	combined = append(combined, value)
	//}
	//

	combined = existingIntervals
	combined = append(combined, newInterval)
	sort.Slice(combined, func(i, j int) bool {
		return combined[i].Start < combined[j].Start
	})

	merged := []*Interval{}
	merged = append(merged, &combined[0])
	for i := 1; i < len(combined); i++ {
		last := merged[len(merged)-1]
		if last.End >= combined[i].Start {
			last.Start = int(math.Min(float64(last.Start), float64(combined[i].Start)))
			last.End = int(math.Max(float64(last.End), float64(combined[i].End)))
		} else {
			merged = append(merged, &combined[i])
		}
	}
	return merged
}
