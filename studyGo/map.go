package main
import (
	"fmt"
)

func printMap(hashMap map[string]string){
	//hashMap 是引用传递，所以可以直接修改map的值
		for key, val := range hashMap {
				fmt.Println("key: ", key)
				fmt.Println("val: ", val)
		}
}

func main(){
	//map第一种声明方式, 声明map, key为string value为string
	var myMap map[string]string
	if myMap == nil{
		fmt.Println("map is nil")
		myMap = make(map[string]string,10)
	}
	//使用map前，需要给map分配内存空间，否则会报错
	myMap["one"] = "Golang"
	myMap["two"] = "JS"

	//map第二种声明方式
	myMap1 := make(map[string]string)
	myMap1["one"] = "Python"
	myMap1["two"] = "Rust"
	fmt.Println(myMap1)

	//map第三种声明方式
	langMap := map[string]string{
		"en": "English",
		"cn":"Chinese",
		"jp":"Japanese",
	}

	//增加map元素
	langMap["fr"] = "French"

	printMap(langMap)
	fmt.Println("---------")
	//删除map元素
	delete(langMap,"fr")

	printMap(langMap)
	//修改map元素
	langMap["jp"] = "Chinese"

	printMap(langMap)
}