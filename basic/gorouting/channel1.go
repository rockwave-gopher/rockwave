package main


/*信道 信道是带有类型的管道，你可以通过它用信道操作符 <- 来发送或者接收值。
ch <- v    // 将 v 发送至信道 ch。
v := <-ch  // 从 ch 接收值并赋予 v。
和映射与切片一样，信道在使用前必须创建:ch := make(chan int)
*/

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // 将和送入 c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)

	go sum(s[:len(s)/2], c)

	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从 c 中接收

	fmt.Println(x, y, x+y)
}
