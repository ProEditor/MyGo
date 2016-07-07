package main

import (
	"fmt"
	"math/rand"
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
