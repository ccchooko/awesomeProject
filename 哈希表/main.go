package main

import "fmt"

func isAnagram(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	exists := make(map[byte]int, 0)
	for i:=0;i<len(s);i++ {
		if v,ok:=exists[s[i]];v>=0&&ok{
			exists[s[i]]=v+1
		}else{
			exists[s[i]]=1
		}
	}
	for i:=0;i<len(t);i++{
		if v,ok:=exists[t[i]];v>=1&&ok{
			exists[t[i]]=v-1
		}else{
			return false
		}
	}
	return true
}

func isAnagram2(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	records := [26]int32{}
	for _, r := range s {
		records[r-[]rune("a")[0]] += 1
	}
	for _, r := range t {
		records[r-[]rune("a")[0]] -= 1
	}
	for _, item := range records {
		if item != 0 {
			return false
		}
	}
	return true
}

func twosum(arr []int, target int) []int {
	valM := map[int]int{}
	for idx, num := range arr{
		if valIdx, ok := valM[target-num]; ok {
			return []int{valIdx, idx}
		} else {
			valM[num] = idx
		}
	}
	return []int{}
}

func main() {
	//fmt.Println(isAnagram2("anagram", "nagaram"))
	fmt.Println(twosum([]int{1,2,3,4}, 5))
}