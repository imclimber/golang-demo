package main

/*
- [并发编程的数据竞争问题以及解决之道](https://juejin.cn/post/6844904170760175630)
假设有一个超长的切片，切片的元素类型为 int，切片中的元素为乱序排列。
限时 5 秒，使用多个 goroutine 查找切片中是否存在给定值，在找到目标值或者超时后立刻结束所有 goroutine 的执行。

比如切片为：[23, 32, 78, 43, 76, 65, 345, 762, …… 915, 86]，查找的目标值为 345，如果切片中存在目标值程序输出:”Found it!” 并且立即取消仍在执行查找任务的 goroutine。
如果在超时时间为找到目标值程序输出:”Timeout! Not Found”，同时立即取消仍在执行查找任务的 goroutine。
*/

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// findTargetWithContext()  // 方法一
	findTargetWithChannel() // 方法二
}

func findTargetWithContext() {
	ctx, canFunc := context.WithCancel(context.Background())

	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	var wg sync.WaitGroup

	for i := 0; i < 9; i += 4 {
		wg.Add(1)
		go findTargetWithContextInner(ctx, canFunc, &wg, inputs, i, i+3, 5)
	}

	wg.Wait()
}

func findTargetWithContextInner(ctx context.Context, canf context.CancelFunc, wg *sync.WaitGroup, in []int, start, end int, target int) {
	defer wg.Done()

	fmt.Println("in,start,end,target", in, start, end, target)
	timeTrigger := time.After(time.Second * 3)
	for {
		select {
		case <-ctx.Done():
			fmt.Println("canceld")
			return
		case <-timeTrigger:
			fmt.Println("time is up")
			return
		default:
			for i := start; i <= end; i++ {
				if in[i] == target {
					fmt.Println("found:", i)
					// time.Sleep(time.Second * 5)
					canf()

					// 这里是否 return 程序不会出现问题，context 的好处就体现出来了。
					return
				} else {
					fmt.Println("cannot find,i:", i)
					// time.Sleep(time.Second * 1)
				}
			}
		}
	}
}

func findTargetWithChannel() {
	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	ch := make(chan int)
	foundCh := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < 9; i += 4 {
		wg.Add(1)
		go findTargetWithChannelInner(ch, foundCh, &wg, inputs, i, i+3, 5)
	}

	a, ok := <-foundCh
	if ok {
		fmt.Println("a:", a)
		close(ch)
	}

	wg.Wait()
}

func findTargetWithChannelInner(stopCh chan int, foundCh chan int, wg *sync.WaitGroup, in []int, start, end int, target int) {
	defer wg.Done()

	fmt.Println("in,start,end,target", in, start, end, target)
	timeTrigger := time.After(time.Second * 3)
	for {
		select {
		case <-stopCh:
			fmt.Println("channel canceld")
			return
		case <-timeTrigger:
			fmt.Println("time is up")
			return
		default:
			for i := start; i <= end; i++ {
				if in[i] == target {
					fmt.Println("found:", i)
					foundCh <- i
					// time.Sleep(time.Second * 5)

					// 此处一定要返回，否则第二次进入此循环会导致死锁，原因是第二次向 foundCh写入，但是main外层只读取了一次，会造成阻塞，main又在等待，因此死锁。
					return
				} else {
					// fmt.Println("cannot find,i:", i)
					// time.Sleep(time.Second * 1)
				}
			}
		}
	}
}
