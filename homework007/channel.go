package homework007

import (
	"fmt"
)

func task3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	//チャネルにデータを送信
	c <- sum
}

func Test_4() {
	fmt.Println("homework007_4")
	s := []int{1, 2, 3, 4}
	c := make(chan int)
	go task3(s, c)
	p := <-c
	fmt.Println(p)
}
