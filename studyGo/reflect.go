package main
import (
	"fmt"
 "reflect"
 "encoding/json"	
)

// reflect.Typeof() 获取当前结构体的类型
// reflect.ValueOf() 获取当前结构体的值
// `` 结构体标签 可以给结构体添加说明 
// struct => json 结构体名称要大写
type User struct {
	Id   int			`json:"id"`
	Name string   `json:"name"`
	Sex string    `json:"sex:`
	Age int       `json:"age"`
	Hobby []string `json:"hobby"`
}

func (this *User) call(){
	fmt.Println("User call...")
}

func structToJson(arg interface{}){

	jsonBytes , err := json.Marshal(arg)
	if err != nil{
		fmt.Println("json marshal failed")
		return 
	}
	fmt.Println(string(jsonBytes))
	structUser := User{}
	e := json.Unmarshal(jsonBytes, &structUser)
	if e != nil{
		fmt.Println("json marshal failed")
		return 
	}
	fmt.Printf("%v\n",structUser) 

}



func main(){
	var double = 3.14

	fmt.Println(reflect.TypeOf(double))
	fmt.Println(reflect.ValueOf(double))

	u := User{Id: 1,Name: "zhangsan",Sex: "男", Age: 25,Hobby: []string{"篮球","足球"}}
	u.call()
	structToJson(u)
	fmt.Println("--------------------------------")
	//获取结构体类型
	userType := reflect.TypeOf(u)
	//获取结构体的值
	userValue := reflect.ValueOf(u)
	//获取结构题中包含的属性
	for i:=0; i<userType.NumField(); i++{
		field := userType.Field(i)
		value := userValue.Field(i)
		fmt.Println(field.Name, field.Type,value)
	}
	//获取结构体中所有的方法
	for i:=0; i< userType.NumMethod(); i++{
		 m := userType.Method(i)
		 fmt.Println(m,userType.Method(i))
	}
	//获取结构体中标签
	el := reflect.TypeOf(User{})
	for i:=0; i<el.NumField();i++{
		tag := el.Field(i).Tag
		fmt.Println(tag)
	}
}