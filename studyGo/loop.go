package main
import (
	"fmt"
	"strings"
	"math/rand"
	"time"
)
var era = "AD"
func main(){
	var command = "walk outside"
	var exit = strings.Contains(command, "outside")
	fmt.Println(exit)
	// if 判断
	var number = rand.Intn(100)
	if number < 60 {
		fmt.Println("没及格")
	}else if number > 60 && number < 80{
		fmt.Println("良好")
	}else{
		fmt.Println("优秀")
	}
	var year = 2000
	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)
	//switch fallthrough 可以继续执行下一个case
	switch leap {
	case true:
		fmt.Println("这是闰年")
	case false:
		fmt.Println("这是平年")
	default:
		fmt.Println("判断失败")
	}
	//for 循环
	var count = 5
	for count > 0{
		fmt.Println(count)
		time.Sleep(time.Second)
		count--
	}
	guessingGame()
	randomDate()
}

// 猜数字游戏
func guessingGame(){
	var num = rand.Intn(100)+1
		for {
			var randomNum = rand.Intn(100)+1
			if(randomNum == num){
				fmt.Printf("恭喜你，猜对了%v\n", num)
				break
			}
			fmt.Printf("猜错了%v\n", randomNum)
		}
}

//随机日期
func randomDate(){
	var count = 10
	for count > 0{
		year := rand.Intn(2024)
		month := rand.Intn(12) + 1
		daysInMonth := 31
		switch month{
		case 2:
			if year%400 == 0 || (year%100 !=0 && year%4 == 0){
				daysInMonth = 29
			}else{
				daysInMonth = 28
			}
		case 4,6,9,11:
			daysInMonth = 30
		}
		day := rand.Intn(daysInMonth) + 1
		fmt.Println(era,year,month,day)
		count--
	}
}