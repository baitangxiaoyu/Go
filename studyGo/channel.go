package main
import (
	"time"
	"fmt"
)


func CacheChannel(){
	c := make(chan int,3)

	go func(){
		defer fmt.Println("CacheChannel go run 结束")
		for i := 0; i < 4; i++ {
			c <- i
			fmt.Println("当前传输管道的值：",i)
			fmt.Println("当前管道的长度：",len(c),cap(c))
		}
		close(c)
	}()

	time.Sleep(time.Second)
	
	for i := 0; i < 4; i++ {
		num := <- c
		fmt.Println(num)
	}
}

func fibonacci(c,n chan int){
	x,y := 1,1

	for{
		select{
		case c <- x:
			x = y
			y = x + y
			fmt.Println(y)
		case <-n:
			fmt.Println("fibonacci 结束")
			return
		}
	}
}

func main(){
	//定义一个 channel
	c := make(chan int)
	CacheChannel()
	go func(){
		defer fmt.Println("go run 结束")
		fmt.Println("go run 开始")
		//将100发送给 channel
		c <- 100
		close(c)
	}()
	//从channel中接收数据 并赋值给 val
	// for {
	// 	if data,ok := <-c; ok{
	// 		fmt.Println(data)
	// 	} else {
	// 		break
	// 	}
	// }

	for data := range c {
		fmt.Println(data)
	}

	channel := make(chan int)
	quit := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			<- channel
		}
		quit <- 0
	}()
	fibonacci(channel,quit)
	fmt.Println("main 结束")
}