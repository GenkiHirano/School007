package homework007

import (
	"fmt"
)

func Test_1() {
	fmt.Println("homework007_1")
	// 第7回スクール 宿題1 スイッチを実装してください。
	signal := "青"
	switch signal {
	case "赤":
		fmt.Println("止まってください")
	case "黄":
		fmt.Println("注意してください")
	case "青":
		fmt.Println("進んでください")
	}

	var a string
	fmt.Printf("赤、黄、青から好きな色を入力してください。")
	fmt.Scan(&a)
	switch a {
	case "赤":
		fmt.Println("止まってください")
	case "黄":
		fmt.Println("注意してください")
	case "青":
		fmt.Println("進んでください")
	}

}
