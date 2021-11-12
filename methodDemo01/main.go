package main
import "fmt"

type Persion struct{
	Name string
}
func (persion Persion) test() {
	persion.Name = "jeck"
	fmt.Println("test(),name=",persion.Name)
}

func (persion Persion) speak() {
	fmt.Println("ll是好人")
}

func (persion Persion) jisuan() {
	count := 0
	for i := 0; i <=1000; i++ {
		count += i
	}
	fmt.Println("jisuan 0---1000的和为", count)
}

func (persion Persion) jisuan2(n int) {
	count := 0
	if n >= 1 {
		for i := 1; i <= n; i++ {
			count += i
		}
		fmt.Println("jisuan2的结果是 ", count)
	}else {
		fmt.Println("传入的参数有误")
	}
}

func (persion Persion) getSum(a int, b int) int {
	return a + b
}

func main() {
	var p Persion
	p.Name = "tom"
	p.test()
	s := p.getSum(10,20)
	fmt.Println(s)
	p.jisuan2(-1)
}