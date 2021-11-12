package main
import "fmt"

type integer int

func (i *integer) chage() {
	*i = *i + 1
}

type Student struct{
	Name string
	Age int
}

func (stu *Student) string() string {
	str := fmt.Sprintf("Name=[%v], Age=[%v]", stu.Name, stu.Age)
	return str
}
func main()  {
	var i integer
	i = 10
	i.chage()
	fmt.Println(i)

	stu := Student{
		Name: "tom",
		Age: 20,
	}
	fmt.Println(&stu)
}
