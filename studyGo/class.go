package main

import "fmt"
//无论是方法名、常量、变量名还是结构体的名称，
//如果首字母大写，则可以被其他的包访问；如果首字母小写，则只能在本包中使用
//结构体首字母大写，代表其他包也可以使用
type Hero struct{
	// 属性名大写代表其他包可以使用了，像面向对象语言中的 public
	Name string
	Age int
	// 属性名小写代表只能结构体内部使用无法继承，像面向对象语言中的 private
	sex string
}

//方法名大写代表其他包可以使用
func (this *Hero) getName() string{
	fmt.Println("_________")
	return this.Name
}

func (this *Hero) setName(name string){
	this.Name = name
}

func (this *Hero) Show(){
	fmt.Println("name:",this.Name)
	fmt.Println("age:",this.Age)
}

//类的继承
type SuperHero struct{
	Hero //SuperHero类继承了Hero类的属性
	Lever int
}

func (this *SuperHero) getLever() int{
	return this.Lever
}

func (this *SuperHero) setLever(lever int){
  this.Lever = lever
}
//重写父类方法
// func (this *SuperHero) getName(){
// 	fmt.Println("superHero getName")
// }
//扩展自己的方法
func (this *SuperHero) Show(){
	fmt.Println("lever:",this)
}

func main(){
	hero := Hero{Name:"hero",Age:18}
	hero.Show()
	hero.setName("newHero")
	hero.Show()

	//继承第一张写法
	// superHero := SuperHero{Hero{"孙悟空",1000,"男"},999}
	//继承第二张写法
	var super SuperHero
	super.Name = "黑神话悟空"
	super.Age = 10000
	// super.Sex = "男"
	super.Lever = 9999
	super.setName("悟空")
	super.Show()
}