package main

import (
	"fmt"
	"go_code/day12_面向对象/封装test/model"
)
func main()  {
	var p model.Account

	p.SetUser("tom")
	p.SetPwd("666666")
	p.SetYue(1000.00)
	fmt.Printf(p.GetPwd(),p.GetUser(),p.GetYue())
}