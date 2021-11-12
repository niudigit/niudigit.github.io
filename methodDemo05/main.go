package main

import "fmt"

type Method struct {

}


func (me *Method) judge(n int) {
	if n % 2 == 0{
		fmt.Println(n,"是偶数")
	}else {
		fmt.Println(n,"是奇数")
	}
}

func (me *Method) print(n int, m int, key string) {
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf(key)

		}
		fmt.Println()
	}
}

type Calcuator struct {

}
func (ca Calcuator) test001(n int, m int, fh string) int {
	count := 0
	if fh == "+" {
		count = n + m
	} else if fh == "-" {
		count = n - m
	} else if fh == "*" {
		count = n * m
	} else if fh == "/" {
		count = n / m
	}
	return count
}
func main()  {
	var me Method
	me.judge(12)
	me.print(6,8,"$")

	var ca Calcuator
	qwe := ca.test001(10,20,"+")
	fmt.Println(qwe)

}


