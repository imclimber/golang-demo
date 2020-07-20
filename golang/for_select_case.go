package main

import (
	"fmt"
	"log"
	"time"
)

// 向channel写数据，写的时间给定为 一秒，超过一秒。则立即返回false。否则写成功返回true
func SendToChan(c chan<- int, f int) bool {
	select {
	case c <- f:
		fmt.Println("send, here c<-f")
		return true
	case <-time.After(time.Second * 10):
		log.Print("send, here time.after")
		return false
	}
}

// 每隔二秒，读channel里的数据
func readChan(c <-chan int) {
	for {
		time.Sleep(time.Second * 20)
		select {
		case m, ok := <-c:
			if !ok {
				log.Print("[r-no] channel is closed")
				return
			}
			log.Print("[r-yes] 读channel成功!", m)
		}
	}
}
func main() {
	pool := 3
	c := make(chan int, pool) // 申请缓冲为3int chan
	go readChan(c)            // 新开协程 每隔二秒，读channel里的数据
	for i := 0; i < 6; i++ {
		ok := SendToChan(c, i)
		if ok {
			log.Print("[w-yes] 写channel成功")
		} else {
			log.Print("[w-no] 写channel失败")
		}
	}
	close(c) // 关闭 channel
	log.Println("finish")
	time.Sleep(time.Second * 86400)
}

// 参考
//- [初识 golang time.After()](https://www.gwalker.cn/article-747088fgc346bc04bb94f6973d291cab.html)
