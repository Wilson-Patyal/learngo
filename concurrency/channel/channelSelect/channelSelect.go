package main

import (
	"fmt"
	"time"
)

// In select statement, if there are multiple cases which match the same condition
// it chooses one at random
// between case 1 & 2  and between case 3 & 4, it will pick randomly.
// In each iteration, always only one select statement will be selected.

func server1(ch chan string) {
	for {
		time.Sleep(3 * time.Second)
		ch <- "This is from server 1"
	}
}

func server2(ch chan string) {
	for {
		time.Sleep(6 * time.Second)
		ch <- "This is from server 2"
	}
}

func main() {
	fmt.Println("Select with channels")
	fmt.Println("--------------------")

	channel1 := make(chan string)
	channel2 := make(chan string)

	go server1(channel1)
	go server2(channel2)

	i := 1
	for {
		i++
		select {
		case s1 := <-channel1:
			fmt.Println("Case one:", s1)
		case s2 := <-channel1:
			fmt.Println("Case two:", s2)
		case s3 := <-channel2:
			fmt.Println("Case one:", s3)
		case s4 := <-channel2:
			fmt.Println("Case two:", s4)
		}
		if i > 10 {
			break
		}
	}
	fmt.Println("Finished with the Channels Select")
	close(channel1)
	close(channel2)
}
