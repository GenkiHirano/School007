package homework007

import (
	"fmt"
	"time"
)

func task1() {
	fmt.Println("田中さんゴール！")
}

func task2() {
	time.Sleep(5 * time.Second)
	fmt.Println("佐藤さんゴール！")
}

func Test_3() {
	fmt.Println("homework007_3")
	go task1()
	go task2()
	time.Sleep(5 * time.Second)
}
