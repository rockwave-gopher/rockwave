package main

import "fmt"

/*信道可以是 带缓冲的。将缓冲长度作为第二个参数提供给 make 来初始化一个带缓冲的信道：

ch := make(chan int, 100)
仅当信道的缓冲区填满后，向其发送数据时才会阻塞。当缓冲区为空时，接受方会阻塞。
deadlock! 当缓冲区小于放入的数量，会抛出错误，当前线程会deadlock！，需要特色处理
*/

func main() {
	ch := make(chan int,1)
	ch <- 1

	ch <- 2
	fmt.Println(<-ch)

	fmt.Println(<-ch)
}
