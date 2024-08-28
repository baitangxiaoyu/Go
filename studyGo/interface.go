package main
import "fmt"

//通用万能类型 interface{}空接口相当于 TS中的any
// 变量可以分为pair( type + value)
// type 可以分为 static type(int、string...)  +  concrete type(interface)

func myType(arg interface{}){

	value, ok := arg.(string)

	if ok {
		fmt.Println("string: ", value)
		fmt.Printf("%T\n",value)
	} else {
		fmt.Printf("not string %v\n",arg)
	}

}

type Reader interface{
	ReadBook()
}

type Writer interface{
	WriterBook()
}
type Book struct{
	title string 
}

func (this *Book) ReadBook(){
	fmt.Println("Reading book: ", this.title)
}

func (this *Book) WriterBook(){
	fmt.Println("Writing book: ", this.title)
}

func main(){
	book := Book{"Go语言"}
	myType(book)
	myType("Hello World")
	myType(100)
	myType(true)
	myType(3.14)
	myType(nil)

	//pair
	// var a string
	// pair<statictype:string value:abc>
	// a = "abc"
	//pair<type:string,value:abc>
	// var allType interface{}
	// allType = a
	// pair<type:Book,value:pair结构体>
	b := &Book{title: "pair结构体"}
	// pair<type:,value> 没有具体类型
	var r Reader
	r  = b 
	r.ReadBook()
	var w Writer
	//动态类型
	w = r.(Writer)
	w.WriterBook()
}