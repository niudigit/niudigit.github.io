package main

import "fmt"

type MethodUtils struct {
	a int
	b int
}

func (me *MethodUtils) sum(len float64, width float64) float64 {
	return len * width
}

func (me *MethodUtils) test() {
	for i := 0; i <me.a; i++{
		for j := 0; j < me.b; j++{
			fmt.Printf("*\t")
		}
		fmt.Println()

	}
}

func (me *MethodUtils) test01(m int, n int) {

}

func main()  {
	var me MethodUtils
	me.a = 8
	me.b = 10
	me.test()

	qae:=me.sum(10.3, 30.9)
	fmt.Println("面积为",qae)
}