package homework007

import (
	"fmt"
)

func task4() {
	number := 2
	switch number {
	case 1:
		fmt.Printf("田中くんの点数は")
	case 2:
		fmt.Printf("佐藤さんの点数は")
	case 3:
		fmt.Printf("鈴木くんの点数は")
	}

func task5(c chan int) {
		sum := 80
		c <- sum
}

func Test_5() {
	fmt.Println("homework007_5")

	c := make(chan int)
	task4()
	go task5(c)
	p := <-c
	fmt.Println(p)
}
