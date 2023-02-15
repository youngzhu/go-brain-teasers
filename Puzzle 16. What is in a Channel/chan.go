package main

import "fmt"

func main() {
	ch := make(chan int, 3)
	ch <- 1
	ch <- 2
	ch <- 3
	<-ch
	close(ch)
	a := <-ch
	b := <-ch
	c := <-ch
	// 从一个关闭的channel中取值
	// 如果buffer中有值，则返回其中的值
	// 否则，返回该类型的零值
	fmt.Println(a, b, c) //2 3 0

}
