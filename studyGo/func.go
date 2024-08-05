package main
import (
	"fmt"
)

//函数多返回值
func main(){
 c := foo("main",1)
 d,e := foo1("main",2)
 f,g := foo2("main",3)
 h,i := foo3("main",4)
 fmt.Println(c)
 fmt.Println(d,e)
 fmt.Println(f,g)
 fmt.Println(h,i)
}
//匿名返回值
func foo(a string,b int)int{
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	c := 100
	return c
}

func foo1(a string,b int)(int,int){
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	return 100,200
}
//返回值有行参
func foo2(a string,b int)(r1 int,r2 int){
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	r1 = 100
	r2 = 200
	return r1,r2
}
func foo3(a string,b int)(r1,r2 int){
	fmt.Println("a = ",a)
	fmt.Println("b = ",b)
	r1 = 100
	r2 = 200
	return r1,r2
}