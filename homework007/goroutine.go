package homework007

import (
	"fmt"
	"time"
)

func task1() {
	time.Sleep(time.Second * 2)
	fmt.Println("佐藤さんゴール！")

}

func task2() {
	fmt.Println("田中さんゴール！")
}

func Test_3() {
	fmt.Println("homework007_3")
	go task1()
	go task2()
	time.Sleep(time.Second * 3)
}
