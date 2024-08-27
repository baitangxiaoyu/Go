package main
import "fmt"

//定义一个自定义类型
type myIni int
//定义一个结构体
type book struct{
	title string
	content string
}

func changeBook(Book book){
	// 传递的是Book的副本，而不是指针
	Book.title = "Go"
}

func changeBookV2(Book *book){
	// 使用指针传递，修改的是结构体
   Book.title = "Go"
}

func main(){
	var number myIni = 10
	fmt.Println(number)
	var myBook book
	myBook.title = "Go语言入门"
	myBook.content = "Go语言入门是Golang语言的一个入门教程，通过这个教程可以掌握Golang语言的基础知识。"
	changeBookV2(&myBook)
	fmt.Printf("%v\n", myBook)
}