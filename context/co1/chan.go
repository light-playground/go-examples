//goroutine 启动后无法对其进行控制
//大部分是等待自己结束,但存在会一直运行的goroutine
//chan + select
package main

import (
	"fmt"
	"time"
)

func main() {
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("退出监控")
				return
			default:
				fmt.Println("持续监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	fmt.Println("监控可以停止了")
	stop <- true

	time.Sleep(5 * time.Second)
}
