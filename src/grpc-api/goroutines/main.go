package main

import (
	"fmt"
	"time"
)

func main() {
	//runGoroutine()
	//runGoroutineChan()
	//runGoroutineSelect()
	runGoroutine2()
}

func runGoroutine() {
	go printMsg("goroutine 1")
	go printMsg("goroutine 2")

	// mainが終わったらgoroutineも終了するためsleep
	time.Sleep(2 * time.Second)

}

func runGoroutine2() {
	worker := 5
	dataChan := make(chan int, worker)
	//closeChan := make(chan struct{})
	result := []int{}

	for i := 0; i < worker; i++ {
		go func(num int) {
			dataChan <- num
		}(i)
	}

	count := 0
	for {
		select {
		case num := <-dataChan:
			result = append(result, num)
			count++
		}
		if count == worker {
			break
		}
	}
	close(dataChan)

	fmt.Println("result:", result)
}

func runGoroutineChan() {
	c := make(chan string)
	defer close(c)
	go func() {
		c <- "goroutine2"
	}()
	fmt.Println(<-c) // 受信されるまで待機
	fmt.Println("goroutine1")

}

func runGoroutineSelect() {

	c := make(chan string)
	defer close(c)

	go func() {
		time.Sleep(3 * time.Second)
		c <- "end process"
	}()

	// channel準備できるまで待機、受信または送信準備ができたら、対象分岐だけ実行して終了
	select {
	case v := <-c:
		fmt.Println("select response:", v)
	case <-time.After(5 * time.Second):
		fmt.Println("time out")
	}

	// defaultを追加するとchannel blockingなしで終了(各分岐でchannel準備ができていなかったら即終了)
	/*
		elect {
		case v := <-c:
			fmt.Println("select response:", v)
		case <-time.After(5 * time.Second):
			fmt.Println("time out")
		default:
			fmt.println("default")
		}
	*/

}

func printMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Println(msg, i)
		time.Sleep(200 * time.Millisecond)
	}
}
