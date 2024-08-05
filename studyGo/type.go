package main 
import (
	"fmt"
	"math/rand"
	"math/big"
	"reflect"
)

/* 
	int64、uint64 => 零值 0
	float64 => 零值 0.00000
	string =>  零值 ""
	bool =>  零值 false
	结构体 => 零值  {0 false}
	数组（切片） => []
	指针 => 零值 nil
*/

func main(){
	// 除法会有精度问题，建议先做乘法
	third := 1.0 / 3
	fmt.Println(third)
	fmt.Printf("%v\n",third)
	fmt.Printf("%f\n",third)
	//小数点后3位
	fmt.Printf("%.3f\n",third)
	//总共4位，小数点后2位
	fmt.Printf("%4.2f\n",third)
	var bigInt = big.NewInt(86400)
	//Printf判断类型
	fmt.Printf("%T",bigInt)
	fmt.Println(reflect.TypeOf(bigInt))

	handleString()
	// deposit()
}

//多语言文本
func handleString(){
	//字符串字面值/原始字符串字面值
	fmt.Println("peace be upon you \n upon you be peace")
	// `` 获取 \n 的原始字符串字面量
	fmt.Println(`strings can span mulitiple lines with the \n escape sequences`)
	//类型别名
	type byte = uint8
	type rune = int32
}

// 存钱游戏
func deposit(){
	var money float64 = 0.0 
	list := [3]float64{0.05,0.10,0.25}
	for{
		if money > 20{
			break
		}
		count := rand.Intn(3)
		val := list[count]
		money += val
		fmt.Printf("%4.2f\n",money)
	 }
}