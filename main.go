package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string)
	people := [5]string{"walter", "sly", "happy", "toby", "cabe"}
	for _, person := range people {
		go isSexy(person, c)
	}

	for i := 0; i < len(people); i++ {
		_ = <-c
	}
	// fmt.Println(result)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
	// time.Sleep(time.Second * 5)
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
}

func isSexy(person string, c chan string) {
	time.Sleep(time.Second * 2)
	fmt.Println(person + " is Sexy")
	c <- person + " is Sexy"
}
