package main

import (
	"fmt"
	"strings"
)

func shout(ping <-chan string, pong chan<- string) {
	for {
		str := <-ping
		str = strings.ToUpper(str) + "!!!"
		pong <- str
	}
}

func main() {
	ping := make(chan string)
	pong := make(chan string)

	go shout(ping, pong)

	var str string
	for {
		fmt.Print("Enter the string: ")
		fmt.Scan(&str)
		str = strings.ToLower(str)
		if str == "q" {
			break
		} else {
			ping <- str
			fmt.Printf("Shout the string: %s", <-pong)
			fmt.Println()
		}
	}

	fmt.Println("Done with the channels")
	close(ping)
	close(pong)
}
