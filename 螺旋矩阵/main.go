package main

import "fmt"

func execute(n int) [][]int {
	arr := make([][]int, n)
	for i := 0; i < n; i++ {
		arr[i] = make([]int, n)
	}
	var (
		startx = 0
		starty = 0
		loop   = n / 2
		mid    = n / 2
		count  = 1
		i, j   int
		offset = 1
	)
	for loop >= 0 {
		i = startx
		j = starty
		for j = starty; j < starty+n-offset; j++ {
			arr[startx][j] = count
			count += 1
		}
		for i = startx; i < startx+n-offset; i++ {
			arr[i][j] = count
			count += 1
		}
		for ; j > starty; j-- {
			arr[i][j] = count
			count += 1
		}
		for ; i > startx; i-- {
			arr[i][j] = count
			count += 1
		}

		startx += 1
		starty += 1

		offset += 2
		loop -= 1
	}
	if n%2 != 0 {
		arr[mid][mid] = count
	}
	return arr
}

func main() {
	fmt.Println(execute(4))
}
