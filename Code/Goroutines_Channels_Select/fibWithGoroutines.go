package main

import "fmt"

func fibonacci(c, quit chan int) {
	//base
	x, y := 0, 1
	for {
		select {
		//check if read on c is active
		case c <- x:
			x, y = y, x+y
		//check if read on quit is active
		case <-quit:
			fmt.Println("quit")
			//close channel
			close(c)
			return
		}
	}
}

func main() {
	//Init of channel c
	c := make(chan int)
	//Init of channel quit
	quit := make(chan int)
	go func() {
		//read 10 values from channel
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		//10 values readed, read from quit channel
		quit <- 0
	}()
	//run async method fibonacci
	fibonacci(c, quit)
}