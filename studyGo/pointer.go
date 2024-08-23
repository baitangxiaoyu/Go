package main
import (
	"fmt"
)

func changeVal(p int){
	//p 只是值传递并不会改变 a 的值
	p = 10
}
func changValV2(p *int){
	//p 是一个指针变量，通过指针可以改变 a 的值
	*p = 10
}

func swap(a *int, b *int){
	temp := *a
	*a = *b
	*b = temp
}

func main(){
	a := 1
	b := 1
	c := 2
	d := 3
	changeVal(a)
	changValV2(&b)
	swap(&c,&d)
	// a => 1
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c,d,&c,&d)
}