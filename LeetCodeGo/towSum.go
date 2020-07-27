package main

import "fmt"

func twoSum(nums []int,target int) []int{
	index := make(map[int]int,len(nums))
	// 新建一个map用户存放target-nums[i]的值
	for i,v := range nums{
			if j,ok := index[target-v];ok{
				// 如果index[target-v]在map中，说明找到了两数之后
				return []int{j,i}
			}
			// 没有找到就把当前的值存入到map中，进入下一轮循环
			index[v] = i
	}
	return nil
}

func main()  {
	list1 := []int{1,2,3,4,5,6,7}
	target := 13
	a := twoSum(list1,target)
	fmt.Println(a)
}
