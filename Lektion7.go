// Lektion 7 - Förkunskaper till Uppgift 1 - 3
// Gorutines & channels
// https://yourbasic.org/golang/goroutines-explained/
// https://yourbasic.org/golang/channels-explained/

package main

/*
	func Publish(text string, delay time.Duration) {
		go func() {
			time.Sleep(delay)
			fmt.Println("BREAKING NEWS:", text)
		}()
	}
*/
func main() {
}

// thread - körs parallelt
// gorutine - lightweight thread
// go fmt.Println("Hello from another goroutine")
// fmt.Println("Hello from main goroutine")
// time.Sleep(time.Second)
/*Publish("A goroutine starts a new thread.", 5*time.Second)
fmt.Println("Let’s hope the news will published before I leave.")

// Wait for the news to be published.
time.Sleep(10 * time.Second)

fmt.Println("Ten seconds later: I’m leaving now.")*/

// time.Now()

// int_ch := make(chan int)
// string_ch := make(chan string, 10)

// Send 3 on the channel. ic <- 3
// Receive a string from the channel. n := <-sc

/*ch := make(chan string)
go func() {
	ch <- "Hello!"
	close(ch)
}()

fmt.Println(<-ch) // Print "Hello!".
fmt.Println(<-ch) // Print the zero value "" without blocking.
fmt.Println(<-ch) // Once again print "".
v, ok := <-ch     // v is "", ok is false.

// Receive values from ch until closed.
for v := range ch {
	fmt.Println(v) // Will not be executed.
}*/

/*If the capacity of a channel is zero or absent,
  the channel is unbuffered and the sender blocks until the receiver has received the value.
If the channel has a buffer, the sender blocks only until the value has been copied to the buffer;
if the buffer is full, this means waiting until some receiver has retrieved a value.
Receivers always block until there is data to receive.*/
