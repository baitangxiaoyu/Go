package main
import (
	"fmt"
	"math/rand"
)
//常量、变量、格式化输出
func main(){
	const name = "Go"
	var age = 15
	number := rand.Intn(10)
	fmt.Println("hello Go")
	// %4v 向左填充空格 %-4v 向右填充空格
	fmt.Printf("我的名字%v,年龄%v\n",name,age)
	fmt.Printf("随机数%v",number)
}
 