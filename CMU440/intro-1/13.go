package main

import (
	"fmt"
	"time"
)

type myMsg struct {
	seqNum  int
	message string
}

func main() {
	fmt.Println("Go channels starting")
	ch1 := make(chan *myMsg)
	go chanSender(ch1, "cs1")

	ch2 := make(chan *myMsg)
	go chanSender(ch1, "cs2")

	for {
		select {
		case msg, ok := <-ch1:
			if ok {
				fmt.Println("CH1: ", msg.seqNum, ":", msg.message)
			}
		case msg, ok := <-ch2:
			if ok {
				fmt.Println("CH2: ", msg.seqNum, ":", msg.message)
			} else {
				fmt.Println("Ch2 closed")
			}

		}
	}
}

func chanSender(out chan<- *myMsg, prefix string) {
	seqNum := 0
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Second)
		out <- &myMsg{seqNum, fmt.Sprintf("%s: %s", prefix, "moo")}
		seqNum++
	}
	close(out)
}
