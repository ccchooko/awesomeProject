package main

import "fmt"

func main() {
	fmt.Println("hello world")
	var nums = []int{1, 4, 8, 9, 10, 20}
	res := binary2(nums, 21)
	fmt.Println(res)
}

// 区间左闭右闭算法, 即[left, right]
func binary1(nums []int, target int) int {
	start := 0
	end := len(nums) - 1 //  没有-1就会导致数组无对应target退出不了
	res := -1
	for mid := (start + end) / 2; start <= end; mid = (start + end) / 2 { // end-1之后，正常start=end的情况在循环内也可能发生，如果此时不加，直接跳出循环return-1
		if nums[mid] > target {
			end = mid - 1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			res = mid
			return res
		}
	}
	return res
}

// 区间左闭右开算法, 即[left, right)
func binary2(nums []int, target int) int {
	start := 0
	end := len(nums) // 对比左右都闭区间，区别
	res := -1
	for start < end { // 左开右闭时，循环内不会出现start=end的是正常判断的情况，因为相等的话就属于找到循环外的情况不存在这个值该退出了
		mid := (start + end) / 2
		if nums[mid] > target {
			end = mid // 注意和上面方发的区别在这里，闭区间时右边界值不会取到mid，因为最后以为元素下标是len-1,因此不用再减1
		} else if nums[mid] < target {
			start = mid + 1
		} else {
			res = mid
			return res
		}
	}
	return res
}
