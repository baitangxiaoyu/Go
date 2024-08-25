package main
import (
		"fmt" 
)

func handleArray(array [10]int){
	// array 为值传递
	fmt.Printf("array type= %T\n",array)
}

func handleMyArray(array []int){
		// array 为引用传递
		fmt.Printf("array type= %T\n",array)
		array[0] = 1
}
//声明slice四种方式
func creatSlice(){
	//声明一个slice,并且初始化长度为1，默认值为 1
	arr := []int{1}
	//声明一个slice，但是没有分配内存，此时直接赋值会报错
	var arr1 []int
	//给 arr1 分配长度为3的数组，默认值为 0，0，0
	arr1 = make([]int,3)
	//声明一个slice，同时分配长度为3的数组，默认值为 0，0，0
	var arr2 []int = make([]int,3)
	//声明一个slice,同时分配长度为3的数组，默认值为 0，0，0,通过 := 推导出 slice是一个切片
	arr3 := make([]int,3)
	//给切片分配值前先判断能否分配
	if arr1 == nil{
		fmt.Println("arr1没有分配内存空间")
	}else{
		fmt.Println("arr1分配了内存空间")
	}
}
//删除数组元素

func remove(array []int,index int)[]int{
	return  append(array[:index],array[index+1:]...)
}

func origintRemove(array []int,index int) []int {
	newArr := make([]int,len(array) - 1)
	k := 0
	for i := 0; i < (len(array) -1); {
		if i != index{
			newArr[i] = array[k]
			k++
		}else{
			k++
		}
		i++
	}
	return newArr
}
func main(){
	//固定长度为10的数组，数组的元素类型为int
	array :=[10]int{1,2,3,4}
	//动态数组 slice
	myArray := []int{0}
	for i := 0; i < len(array); i++ {
		fmt.Println(array[i])
	}
	handleArray(array)
	handleMyArray(myArray)
	fmt.Println(myArray,len(myArray))
}

