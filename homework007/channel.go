package homework007

import (
	"fmt"
)

func task3(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func Test_4() {
	fmt.Println("homework007_4")
	s := []int{1, 2, 3, 4}
	// 第7回スクール 宿題4 チャネルを実装してください。
	c := make(chan int)
	go task3(s, c)
	p := <-c
	fmt.Println(p)
}
