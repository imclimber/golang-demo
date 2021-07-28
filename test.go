package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	ov := make(chan struct{})
	pool := NewGreedyPool(30, func() {
		close(ov)
	})

	go func() {
		for {
			time.Sleep(time.Second)
			fmt.Println(runtime.NumGoroutine())
		}
	}()

	time.Sleep(time.Second)
	for i := 0; i < 100; i++ {
		idx := i
		pool.Send(func() {
			idx += 1
			fmt.Println(idx)
			time.Sleep(time.Second)
		})

		fmt.Println("==========i========:", i)
	}

	pool.Over()
	fmt.Println("============= before <- over===========================")
	<-ov
	fmt.Println("============= after <- over===========================")
	fmt.Println("Over")
}

type PoolFunc func()

// Greedy Pool
type greedyPool struct {
	pool chan PoolFunc
	over chan struct{}

	close PoolFunc
}

// NewGreedyPool: return greed pool
func NewGreedyPool(size int, close PoolFunc) *greedyPool {
	pool := &greedyPool{
		pool:  make(chan PoolFunc, size),
		over:  make(chan struct{}),
		close: close,
	}

	go pool.scheduler(size)
	return pool
}

// Over: End of task issuance
func (g *greedyPool) Over() {
	close(g.over)

}

// Send: Issue a task
func (g *greedyPool) Send(fn PoolFunc) {
	g.pool <- fn
}

func (g *greedyPool) scheduler(size int) {
	var wg sync.WaitGroup
	for i := 0; i < size; i++ {
		wg.Add(1)
		go g.task(&wg)
	}

	fmt.Println("= =before wait===============")
	wg.Wait()
	fmt.Println("= ==after wait=======****************========")
	g.close()
	fmt.Println("========================================****************************")
}

func PrintStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	fmt.Printf("==> %s\n", string(buf[:n]))
}

func (g *greedyPool) task(wg *sync.WaitGroup) {
	defer func() {
		wg.Done()

		// recover
		if err := recover(); err != nil {
			PrintStack()
			fmt.Println("Recover Err: ", err)
		}
	}()

loop:
	for {
		select {
		case i, ex := <-g.pool:
			if !ex {
				break loop
			}
			fmt.Println("pool in ...........")
			i()
		case <-g.over:
			fmt.Println("over is closed ----------------------------")
			break loop
		}
	}
}
