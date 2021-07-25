package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	// findTargetWithContext()  // 方法一
	findTarget_V1() // 方法二https://juejin.cn/post/6844904170760175630
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
