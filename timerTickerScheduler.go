package main

import ("fmt"
 		"time"
		"os"
)

func main () {
    fmt.Println("start")
    time.Sleep(time.Second * 4)
    fmt.Println("after 4 seconds")

	//scheduler with sleep
	fmt.Println("\nscheduler with sleep\n")
	for true {
		fmt.Println("Hello !!")
		time.Sleep(1 * time.Second)
	}

	///timer
	var timer = time.NewTimer(4 * time.Second)
	fmt.Println("start")
	<-timer.C
	fmt.Println("finish")

	//after function

	fmt.Println("\nafter function\n")
	var ch = make(chan bool)

	time.AfterFunc(4*time.Second, func() {
		fmt.Println("expired")
		ch <- true
	})

	fmt.Println("start")
	<-ch
	fmt.Println("finish")

	//time after
	fmt.Println("\ntime after\n")
	<-time.After(4 * time.Second)
	fmt.Println("expired")

	//ticker
	fmt.Println("\nticker\n")
	done := make(chan bool)
    ticker := time.NewTicker(time.Second)

    go func() {
        time.Sleep(10 * time.Second) // wait for 10 seconds
        done <- true
    }()

    for {
        select {
        case <-done:
            ticker.Stop()
            return
        case t := <-ticker.C:
            fmt.Println("Hello !!", t)
        }
    }

	//timerXGoroutine

	fmt.Println("\ntimerXgoroutine\n")
	var timeout = 5
    var ch1 = make(chan bool)

    go timer1(timeout, ch1)
    go watcher(timeout, ch1)

    var input string
    fmt.Print("what is 725/25 ? ")
    fmt.Scan(&input)

    if input == "29" {
        fmt.Println("the answer is right!")
    } else {
        fmt.Println("the answer is wrong!")
    }
}

func timer1(timeout int, ch chan<- bool) {
    time.AfterFunc(time.Duration(timeout)*time.Second, func() {
        ch <- true
    })
}

func watcher(timeout int, ch <-chan bool) {
    <-ch
    fmt.Println("\ntime out! no answer more than", timeout, "seconds")
    os.Exit(0)
}