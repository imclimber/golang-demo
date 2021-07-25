package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	findTarget()
}

func findTarget() {
	ctx, canFunc := context.WithCancel(context.Background())

	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	var wg sync.WaitGroup

	for i := 0; i < 9; i += 4 {
		wg.Add(1)
		go findTargetInner(ctx, canFunc, &wg, inputs, i, i+3, 5)
	}

	wg.Wait()
}

func findTargetInner(ctx context.Context, canf context.CancelFunc, wg *sync.WaitGroup, in []int, start, end int, target int) {
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
				} else {
					fmt.Println("cannot find,i:", i)
					// time.Sleep(time.Second * 1)
				}
			}
		}
	}
}

func findTargetWithChannel() {
	ctx, canFunc := context.WithCancel(context.Background())

	inputs := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}

	var wg sync.WaitGroup

	for i := 0; i < 9; i += 4 {
		wg.Add(1)
		go func(ctx context.Context, in []int, start, end int, target int) {
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
							canFunc()
						} else {
							fmt.Println("cannot find,i:", i)
							// time.Sleep(time.Second * 1)
						}
					}
				}
			}
		}(ctx, inputs, i, i+3, 5)
	}

	wg.Wait()
}
