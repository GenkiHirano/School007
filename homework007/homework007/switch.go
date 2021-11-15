package homework007

import "fmt"

func Test_1() {
	fmt.Println("homework007_1")
	var a string = "赤"
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
