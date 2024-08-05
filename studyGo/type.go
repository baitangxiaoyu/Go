package main 
import (
	"fmt"
	"math/rand"
)

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
	deposit()
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