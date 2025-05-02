package main

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {
	//runGoroutine()
	//runGoroutineChan()
	//runGoroutineSelect()
	//runGoroutine2()
	//runGoroutineWaitGroup()
	//runGoroutineWaitGroup2()
	runGoroutineSample()
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

func runGoroutineWaitGroup() {
	var wg sync.WaitGroup

	for i := 0; i < 3; i++ {
		wg.Add(1) // goroutine count +1

		go func(n int) {
			defer wg.Done() // goroutine count -1
			fmt.Println("goroutine test :", n)
		}(i)
	}

	wg.Wait() // goroutineが終了(Done)まで待機
	fmt.Println("end")
}

func runGoroutineWaitGroup2() {
	var (
		worker   = 5
		dataChan = make(chan int)
		result   = []int{}
		wg       sync.WaitGroup
	)

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			dataChan <- n
		}(i)
	}

	go func() {
		wg.Wait()
		close(dataChan)
	}()

	for num := range dataChan {
		result = append(result, num)
	}

	fmt.Println("result:", result)
}

const maxWorkers = 5

func runGoroutineSample() {
	msgQueue := make(chan string)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	//defer close(msgQueue)
	var wg sync.WaitGroup
	sem := make(chan struct{}, maxWorkers)

	go func() {
		for i := 0; i < 20; i++ {
			msg := fmt.Sprintf("message-%d", i)
			msgQueue <- msg
			time.Sleep(100 * time.Millisecond)
		}
		close(msgQueue)
	}()

	go func() {
		wg.Wait()

	}()

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled")
			return
		case msg, ok := <-msgQueue:
			if !ok {
				log.Println("msgQueue closed")
				return
			}

			sem <- struct{}{}
			wg.Add(1)

			go func(m string) {
				defer func() {
					<-sem
					wg.Done()
				}()

				if err := handleMessage(m); err != nil {
					log.Printf("error: %v\n", err)
				}
			}(msg)
		}
	}

}

func handleMessage(msg string) error {
	log.Println("start message:", msg)

	if msg == "message-13" {
		return fmt.Errorf("error message: %s", msg)
	}
	log.Println("message:", msg)
	time.Sleep(500 * time.Millisecond)

	return nil
}
