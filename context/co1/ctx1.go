package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go monitor(ctx, "[监控1]")
	go monitor(ctx, "[监控2]")
	go monitor(ctx, "[监控3]")
	go monitor(ctx, "[监控4]")

	time.Sleep(10 * time.Second)
	fmt.Println("通知监控停止")
	cancel()
	time.Sleep(5 * time.Second)
}

func monitor(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println(name, "监控停止, 并退出")
			return
		default:
			fmt.Println(name, "goroutine监控中")
			time.Sleep(2 * time.Second)
		}
	}
}
