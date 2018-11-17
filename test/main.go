package main

import (
	"fmt"
	"sync"
	"time"
)

/*
func compute(value int){
	for i:=0 ; i < value; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main (){
	fmt.Println("Concurrency with GoRoutines")
	go compute(5)
	go compute(5)

	fmt.Scanln()
}
*/

func fiveMultiplier (c chan int, val int) {
	c <- val * 5
}


func main (){
	fmt.Println("Concurrency with GoRoutines and Channel")

	fooVal := make(chan int)

	/*
	go fiveMultiplier(fooval, 5)
	go fiveMultiplier(fooval, 3)

	v1 := <- fooval
	v2 := <- fooval

	fmt.Println(v1 , v2)
	*/

	for i := 0; i < 10; i++ {
		go fiveMultiplier(fooVal, i)
	}

	close(fooVal)

	for item := range fooVal {
		fmt.Println(item)
	}

	time.Sleep(time.Second * 2)
}
