package main

import (
	"fmt"
	"math/rand"
	"time"
)

//const用来表示常量,常量不能使用:=语法
const ABC = 123456
const (
	Big   = 1 << 100
	Small = 1 >> 99
)

type Vertex struct {
	X int
	Y int
}

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
	// testFor()
	// foreverLoop()
	// goIf()
	// goSwitch()
	// noCoditionSwitch()
	// deferExec()
	// goPoint()
	// goStruct()
	// goArray()
	// goAppendElement()
	goArrayRange()

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

//Go也支持if then else 模式,就是简单的在switch中不设置条件
func noCoditionSwitch() {
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Good Morning")
	case t.Hour() < 17:
		fmt.Println("Good Afteroon")
	default:
		fmt.Println("Good Evening")
	}
}

//go中通过defer来延迟语句的执行,使用defer的语句在方法返回的时候才执行
func deferExec() {
	defer fmt.Println("这句话在前面,但是等到方法返回后才执行")
	fmt.Println("我在后面,但是我先执行")
}

//go中的指针
func goPoint() {
	a, b := 1, 2
	p := &a
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(*p)
}

//go中的结构体
//指针可以直接的引用结构体,而不需要解引用;
//初始化结构体的时候,可以只初始化一部分
func goStruct() {
	fmt.Println(Vertex{1, 2})
	fmt.Println("测试结构体和指针的引用")
	v := Vertex{X: 3, Y: 5}
	p := &v
	p.X = 1223
	p.Y = 2323
	fmt.Println(p)
}

//go的声明方式是 []T,其中T为数组的类型,你可以任意截取数组,截取数组并不会创建新的数组,所以当你改变截取后的数组后,原始数组也会跟着改变
func goArray() {
	names := [4]string{
		"Hello",
		"World",
		"Good",
		"Morning",
	}

	//下面两句的含义是一样的
	fmt.Println(names[0:2])
	fmt.Println(names[:2])
	fmt.Println(names)
	var nilArray []int
	fmt.Println(nilArray)
	if nilArray == nil {
		fmt.Println("nilArray")
	}

}

func goAppendElement() {
	names := [2]string{"Hello", "World"}
	var nmaes1 []string
	nmaes1 = append(nmaes1, "再说一次Hello World")
	fmt.Println(names)
	fmt.Println(nmaes1)
}
func goArrayRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	//可以使用如下方式遍历数组,i表示索引,v表示值
	for i, v := range pow {
		fmt.Printf("索引是%d,元素是%d\n", i, v)
	}
	//如果不需要索引或值,可以使用_代替
	for _, v := range pow {
		fmt.Printf("元素是%d\n", v)
	}

}
