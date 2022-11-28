package main

import "fmt"

func send_msg(ch chan string) {

	for i := 0; i < 5; i++ {
		var msg string
		fmt.Scanln(&msg)
		ch <- msg
	}

}

func main() {

	my_ch := make(chan string)
	go send_msg(my_ch)
	for v := range my_ch {
		fmt.Println(v)
		for true {
			fmt.Println(<-my_ch)
		}

	}
}
