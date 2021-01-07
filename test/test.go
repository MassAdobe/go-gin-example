/**
 * @Time : 2020/12/18 1:38 下午
 * @Author : MassAdobe
 * @Description: test
**/
package main

import (
	"log"
	"time"
)

func main() {
	cb := make(chan bool, 1)

	go func() {
		timer2 := time.NewTimer(time.Second * 1)
		tick := time.Tick(time.Second * 5)
		select {
		case <-timer2.C:
			log.Println("退出协程")
			return
		case <-tick: // 模拟超时任务
			log.Println(222) // 处理业务代码
			cb <- true
			return
		}
	}()

	log.Println("回到主线程")

	timer := time.NewTimer(time.Second * 2)
	select {
	case <-timer.C:
		log.Println("time out")
		cb <- false
	case tmp := <-cb:
		log.Println(tmp)
	}

	time.Sleep(time.Second * 10)
	log.Println(123)
}
