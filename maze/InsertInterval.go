package main

import (
	"fmt"
)


func insert(intervals [][]int, newInterval []int) [][]int {
	size := len(intervals)
	start := 0
	if size == 0{
		return intervals
	}
	mid := (size + start)/2
	for ;start<size;{
		if intervals[mid][1] <newInterval[0]{
			size = mid -1
			mid = (size+start) / 2
		}else if intervals[mid][0] > newInterval[1]{
			start = mid +1
			mid = (size+start) / 2
		}else{
			break
		}
	// intervals = append()

	}

	fmt.Println(size)
	return nil
}
func main()  {
	intervals := [][]int{{1,2},{3,5},{6,7},{8,10},{12,16}}
	newInterval := []int{4,8}
	//  intervals
	fmt.Println(intervals,newInterval)
	// insert(intervals,newInterval)
}
