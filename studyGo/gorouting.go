package main
import (
	"time"
	"fmt"
	"runtime"
)

//从协程
func Task(){
	i := 0
	for {
		i++
		fmt.Printf("task = %v\n",i)
		time.Sleep(time.Second)
	}
}

// 主协程 主协程执行完后 从协程也会自动结束
func main(){
	// 启动一个协程
	go Task()

	//c匿名函数
	go func(){
		defer fmt.Println("defer A")
		func(){
			defer fmt.Println("defer B")
			// 退出多层函数栈
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
		// return 只能退出当前函数栈
		return 
	}()
	// 此时匿名函数和主函数是并行执行的 无法获取到返回值 通过管道获取另一个线程的返回值
	go func(a int,b int)boolean{
		return a + b >= 10
	}(1,2)

	i := 0
	for{
		i++
		fmt.Printf("main = %v\n",i)
		time.Sleep(time.Second)
	}

}