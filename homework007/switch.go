package homework007

import "fmt"

func Test_1() {
	fmt.Println("homework007_1")
	signal := "青"

	switch signal {
	case "赤":
		fmt.Println("止まってください")
	case "黄":
		fmt.Println("注意してください")
	case "青":
		fmt.Println("進んでください")
	}
}
