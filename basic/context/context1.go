package main

import (
	"fmt"
	"time"
	"context"
)

/**
context.Background() 返回一个空的Context，这个空的Context一般用于整个Context树的根节点。
然后我们使用context.WithCancel(parent)函数，创建一个可取消的子Context，
然后当作参数传给goroutine使用，这样就可以使用这个子Context跟踪这个goroutine。
 */
func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("监控退出，停止了...")
				return
			default:
				fmt.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	fmt.Println("可以了，通知监控停止")

	cancel()

	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

}
