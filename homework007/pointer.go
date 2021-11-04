package homework007

import "fmt"

func Test_2() {
	fmt.Println("homework007_2")
	a := 50

	var pa *int
	pa = &a

	fmt.Println(pa)
	fmt.Println(*pa)
}
