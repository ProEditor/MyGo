package main

import (
	"fmt"
	"math/rand"
)

//const用来表示常量,常量不能使用:=语法
const ABC = 123456
const (
	Big   = 1 << 100
	Small = 1 >> 99
)

func main() {
	//按照惯例,如果方法的首字母为大写,则表示这个方法为public,否则为private
	fmt.Println("这是输出的随机数", rand.Intn(10))
	// fmt.Printf("hello, world\n")
	//Reverse("hello, world\n")
	fmt.Println(swap("Hello", "OGE"))
	//使用var语句可以一次性定义多个变量,var语句可以在package和function级别中使用
	//使用var还可以初始化这一串变量的值
	var a, b, c, d int = 1, 2, 3, 4
	fmt.Println(a, b, c, d)
	fmt.Println(nameReturn("I Have a name", "hahahha"))
	fmt.Println(ABC)
	// needTypeCheck()
	testFor()
	// foreverLoop()
	// goIf()
	goSwitch()

}

func Reverse(s string) string {
	//在方法里面,:=语句可以实现var的功能,go会自动监测类型
	r := []rune(s)
	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-i {
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

//go支持一个函数返回多个结果
func swap(x, y string) (string, string) {
	return y, x
}

//go还支持返回带名称的值
func nameReturn(x, y string) (a, b string) {
	a = y
	b = x
	return
}
func needInt(x int) int {
	return x*10 + 1
}
func needFloat(x float64) float64 {
	return x * 0.1
}

//如果在定义常量的时候,没有指定一个类型,那么会根据需要自动进行推导
func needTypeCheck() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

//for循环不需要括号,初始语句和post语句可以省略,如果连分号都不要了,那for就是一个while循环了
func testFor() {
	sum := 1
	for sum < 100 {
		sum += sum
		fmt.Println(sum)
	}
}

//如果在for中啥都不写,那么这就是一个死循环了
func foreverLoop() {
	for {
		fmt.Println("这是一个死循环")
	}

}

//go中的if语句,go中的if语句可以有一个前置语句,在if判断之前执行它
//下面的if中定义了v变量,这个变量在else语句中也是可以访问的
//需要注意if后面的}括号和else必须在一行,否则编译的时候会报错
func goIf() {
	if v := 100; v < 100 {
		fmt.Println("90是小于100的")
	} else {

		fmt.Print(v)
	}
}

//go中的switch语法基本和if是一致的
func goSwitch() {
	switch os := "a"; os {
	case "a":
		fmt.Println("case a")
	case "b":
		fmt.Println("case b")
	case "c":
		fmt.Println("case c")
	}
}
