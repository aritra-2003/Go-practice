package main

import (
	"fmt"
	"time"
)

func greet(phrase string,done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string,done chan bool) {
	time.Sleep(3 * time.Second) // simulate a slow, long-taking task
	fmt.Println("Hello!", phrase)
	done <- true
	close(done)
}

func main() {
	done :=make (chan bool)

	go greet("Nice to meet you!",done)
	// done[0] = make(chan bool)

	go greet("How are you?",done)
    // done[1] = make(chan bool)
	go slowGreet("How ... are ... you ...?",done)
	// done[2] = make(chan bool)
	// for _, ch := range done {
	// 	<-ch
	// }
// 	go greet("I hope you're liking the course!")
// 	time.Sleep(20* time.Second)
// 

 for  range done{
	
 }
}