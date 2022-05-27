package main

import "fmt"

// 翻转所有字符串
func reverse(s []byte) {
	left :=0
	right := len(s)-1
	for left < right {
		temp := s[left]
		s[left] = s[right]
		s[right] = temp
		left ++
		right --
	}
}

// 每2k个字符串翻转前k个
func reverser2K(s []byte, k int) {
	for i:=0; i<len(s); i += 2*k {
		if (i + k) < len(s) {
			reverse(s[i: i+k])
			continue
		} else {
			reverse(s[i: ])
		}
	}
}

// 空格替换成%20
// 使用go的话需要考虑数组新建的问题
// 方法一：遍历添加，空间复杂度O（N），新建一个空的byte数组，从前往后依次添加要添加的字符串即可
// 方法二：原地添加，空间复杂度O（1）, 如下
func replaceSpace(s string) string {
	ss := []byte(s)
	length := len(ss)
	spaceCount := 0
	for _, i := range ss {
		if i == ' ' {
			spaceCount += 2
		}
	}
	// 扩展原有数组切片
	tmp	:= make([]byte, spaceCount)
	b := append(ss, tmp...)
	i := length - 1 // 原数组指针
	j := len(b) - 1 // 新数组指针
	for i >= 0 {
		if ss[i] != ' ' {  // 此处如果是双引号，判断则不通过，不知道为啥
			b[j] = ss[i]
			j--  // 不是空格的情况，扩展后的数组每次往前走一步
		} else {
			b[j] = '0'
			b[j-1] = '2'
			b[j-2] = '%'
			j-=3 // 是空格的情况，扩展后的数组每次往前走三步
		}
		i--  // 原数组每次往左走一步
	}
	return string(b)
}

func main() {
	//s := []byte("a b")
	//reverser2K(s, 2)
	s := " a fsd dfasdf "
	fmt.Println(replaceSpace(s))
}
