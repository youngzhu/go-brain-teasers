package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := append(s1[:1], 10)
	bigger := []int{5, 6, 7, 8}
	s3 := append(s1[:1], bigger...)

	// 如果容量（cap）够用，则共用底层数组
	// 如果不够，则新建一个底层数组，并把相应数据复制过去

	printSlice(s1) //len=3, cap=3, [1 10 3]
	printSlice(s2) //len=2, cap=3, [1 10]
	printSlice(s3) //len=5, cap=6, [1 5 6 7 8]

}

func printSlice(s []int) {
	fmt.Printf("len=%d, cap=%d, %v\n", len(s), cap(s), s)
}
