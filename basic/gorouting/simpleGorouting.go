package main


/**
main goroutine将计算菲波那契数列的第45个元素值。由于计算函数使用低效的
递归，所以会运行相当长时间，在此期间我们想让用户看到一个可见的标识来表明程序依然
在正常运行，所以来做一个动画的小图标：
主函数返回时，所有的goroutine都会被直接打断，程序退出。
 */
import (
	"fmt"
	"time"
)

func main() {
	//start a spinner go routing
	go spinner(100 * time.Millisecond)

	const n = 45
	fibN := fib(n) // slow
	fmt.Printf("\rFibonacci(%d) = %d\n", n, fibN)
}
func spinner(delay time.Duration) {
	for {
		for _, r := range `-\|/` {
			fmt.Printf("\r%c", r)
			time.Sleep(delay)
		}
	}
}
func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}