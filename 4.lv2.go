// 这里好像真的搞不太会
// 应该只有第三个能实现。。。
package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(3)
	go time1()
	go time2()
	go Ticker()
	wg.Wait()
}
func time1() {
	defer wg.Done()
	for {
		now := time.Now()                                                                    //获取当前时间，放到now里面，要给next用
		next := now.Add(time.Hour * 24)                                                      //通过now偏移24小时
		next = time.Date(next.Year(), next.Month(), next.Day(), 2, 0, 0, 0, next.Location()) //获取下一个凌晨2点的日期
		t := time.NewTimer(next.Sub(now))                                                    //计算当前时间到凌晨2点的时间间隔，设置一个定时器
		<-t.C
		fmt.Printf(" %v:我还能再战4小时！\n", time.Now())
	}
}

func time2() {
	defer wg.Done()
	for {
		now := time.Now()
		next := now.Add(time.Hour * 24)
		next = time.Date(next.Year(), next.Month(), next.Day(), 6, 0, 0, 0, next.Location())
		t := time.NewTimer(next.Sub(now))
		<-t.C
		fmt.Printf(" %v:我要去图书馆开卷！\n", time.Now())
	}
}
func Ticker() {
	defer wg.Done()
	ticker := time.NewTicker(2 * time.Second)
	for {
		select {
		case <-ticker.C:
			fmt.Printf("芜湖！起飞！\n")
		}
	}
}

/*至少编写3个计时器，完成以下功能：

●  在每天凌晨两点，输出"我还能再战4小时！"
●  在每天早上六点，输出"我要去图书馆开卷！"
●  自程序运行时起，每过半分钟输出"芜湖！起飞！"
●  其他你想完成的功能

文案可自行修改，测试时可以把时间间隔缩短。

要求：使用并发。

LV2EX

在上述功能的基础上添加额外的功能：

●  用户可以自己增加定时器（一次性的&重复提醒的）
●  用户可以指定删除某个已存在的定时器
●  用户可以在不删除定时器（重复提醒的）的情况下，取消定时器的下一次提醒
*/
