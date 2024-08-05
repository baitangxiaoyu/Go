package main
import (
	"fmt"
	"math/rand"
)

//常量、变量、格式化输出
func main(){
	//常量
	const name = "Go"
	//const() 添加一个关键字iota，iota 每行的iota都会累加1，第一行的iota的默认值是0
	const (
		BEIJING = iota
		SHANGHAI
		SHENZHEN
		GUANGZHOU
	)
	//变量
	var age = 15
	// count := 100
	// var age int = 15
	number := rand.Intn(10)
	fmt.Println("hello Go")
	// %4v 向左填充空格 %-4v 向右填充空格
	fmt.Printf("我的名字%v,年龄%v\n",name,age)
	fmt.Printf("随机数%v",number)
}
 