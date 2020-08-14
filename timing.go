package main

import (
	"fmt"
	"time"
)

/*
func demo() {
	for {
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("1s timer")

		case <-time.After(time.Second * 2):
			fmt.Println("2s timer")
		}
	}
}

func main() {

	go demo()
	select {}
}*/
/*
看到2s的定时器一直没有执行，原因就是 select 每次执行都会重新执行 case 条件语句，
并重新注册到 select ,每次都会新构造一个 Timer 对象，所以2秒的定时器永远不会执行。
而且会造成内存的泄露。这样的设计倒是也有好处，定时器可用于另一个chan的超时处理*/

func demo(t1 interface{}, t2 interface{}) {
	for {
		select {
		case <-t1.(*time.Ticker).C:
			fmt.Println("1s timer")

		case <-t2.(*time.Ticker).C:
			fmt.Println("2s timer")
		}
	}
}

func main() {

	t1 := time.NewTicker(time.Second * 1)
	t2 := time.NewTicker(time.Second * 2)
	go demo(t1, t2)
	select {}
}

/*
//定时创建数据库
func TimeToCreatDb() {
    for {
        now := time.Now()  //获取当前时间，放到now里面，要给next用
        next := now.Add(time.Hour * 24) //通过now偏移24小时
        next = time.Date(next.Year(), next.Month(), next.Day(), 0, 0, 0, 0, next.Location()) //获取下一个凌晨的日期
        t := time.NewTimer(next.Sub(now))//计算当前时间到凌晨的时间间隔，设置一个定时器
        <-t.C
        Printf("凌晨创建一个文件: %v\n",time.Now())
        //以下为定时执行的操作
        creatdb()
    }
}
*/
