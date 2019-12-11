package main

import (
	"fmt"
)




func init()  {
	fmt.Println("初始化方法！可以有多个")
}

//包函数参数可以只定义不使用，函数中参数定义必须使用
var a ,b,c = 1,true,"str"
var x string   //定义x为字符串类型，值为0值""
var y = ""  //定义y，根据值自动判断类型

func main() {
	a,b,c := 1, 1,1  //只能用于函数中，并且必须存在新声明的值
	//a,b,c := 2, 2,2 报错
	//a,b,e := 2,2,2 可以，因为有新声明的值e

	//测试if  else if
	fmt.Println(a+b+c)
	aa := test1(100,20)
	fmt.Println(aa)
	aa = test1(200,200)
	fmt.Println(aa)
	aa = test1(1000,2000)
	fmt.Println(aa)

	/////////////////////测试switch 判断变量类型
	test2(a)
	//////////////////////测试  fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
	//比较x范围  sitch 后不能跟x  ，如果case是固定值，可以写成switch x
	test3(3)

	///////////////////测试for三种形式
	test4()
	fmt.Println()

	//值传递
	x , y := 100 ,200
	fmt.Printf("转换前x：%d，y：%d ", x ,y)
	test5(x,y)
	fmt.Printf("转换后x：%d，y：%d ", x ,y)
	fmt.Println()

	///////////////////////////引用传递
	fmt.Printf("转换前x：%d，y：%d ", x ,y)
	test6(&x, &y)
	fmt.Printf("转换后x：%d，y：%d ", x ,y)

	fmt.Println()
	/////////////////////////函数作为参数
	fmt.Println(test7("mao", func(str string) string {
		return str + "-形参内定义函数行为"
	}))
	//已定义了函数
	fmt.Println(test7("mao", mao1))

	////////////////////////////闭包是匿名函数
	num := test8()(1)
	fmt.Println("闭包匿名函数：")
	fmt.Println(num)
	num = test8()(num)
	fmt.Println(num)
	funcc := test8()
	fmt.Println(funcc(num))

	//////////////////////////////////测试 “方法”
	var hh hanshu = "strrrr"
	fangfa := hh.getStringAppend("append")
	fmt.Println("调用函数的方法拼接后：",fangfa)
	//声明正方形结构体
	var sq square
	//给正方形结构体边长赋值
	sq.length = 3
	//调用方法求周长
	fmt.Printf("正方形边长：%d,周长：%d",sq.length,sq.getPerimeter())

	/////////////////////////////测试数组的三种方法
	test9()
	//输入两个数组，将两个数组中数求和,其中参数2只能是长度为3的数组，所以需要显示声明数组长度为3
	arr1 := []int{1,2,3,4,5,6,7}
	arr2 := [3]int{1,2,3}
	fmt.Printf("数组1：%v,数组2：%v中值求和得：%v \n",arr1,arr2,test10(arr1,arr2))


	///////////////////////////////////测试指针，返回内存地址
	str := "asdfasdfasdf"
	add := test11(str)
	fmt.Printf("字符串值：%v,地址为：%v",str,add)
}

func mao1(str string) string{
	return str + "-现有函数"
}


func  test1(a , b int) string{
	if a == 100{
		return "a=100"
	}else if b ==200 {
		return "b=200"
	}else{
		return "a!=100  && b!=200"
	}
}

//switch 语句还可以被用于 type-switch 来判断某个 interface 变量中实际存储的变量类型。
func test2(x interface{}) {

	switch i := x.(type) {
	case nil:
		fmt.Println(" x 的类型 :%T",i)
	case int:
		fmt.Println("x 是 int 型")
	case float64:
		fmt.Println("x 是 float64 型")
	case func(int) float64:
		fmt.Println("x 是 func(int) 型")
	case bool, string:
		fmt.Println("x 是 bool 或 string 型" )
	default:
		fmt.Println("未知型")
	}
}


//使用 fallthrough 会强制执行后面的 case 语句，fallthrough 不会判断下一条 case 的表达式结果是否为 true。
func test3(x int){
	//比较x范围  sitch 后不能跟x  ，如果case是固定值，可以写成switch x
	switch {
	case x > 1:
		fmt.Println("x大于1")
		fallthrough
	case x > 2:
		fmt.Println("x大于2")
	case x == 3:
		fmt.Println("x大于3")
	}

	switch x{
	case 1:
		fmt.Println("x==1")
		fallthrough
	case 2:
		fmt.Println("x==2")
	case 3:
		fmt.Println("x==3")
	}
}


var mapp = make(map[int] int)

//for 三种表达形式
func test4(){

	///////////////////////////第一种格式
	fmt.Print("for循环第一种形式：")
	for i:=0 ; i<10 ; i++{
		fmt.Print(i)
		if i< 9{
			fmt.Print(",")
		}
		//将值放入map中
		mapp[i] = i
	}
	fmt.Println()

	a:=0
	for a<10{
		a++
		fmt.Printf("a等于%d", a)
	}
	fmt.Println()
	b:=0
	for ;b<10 ; b++{
		fmt.Printf("b等于%d", b)
	}
	fmt.Println()

	strings := []string{"google", "runoob"}
	for i, s := range strings {
		fmt.Println(i, s)
	}


	numbers := [6]int{1, 2, 3, 5}
	for i,x:= range numbers {
		fmt.Printf("第 %d 位 x 的值 = %d\n", i,x)
	}
	fmt.Println()
	/////////////////////////////////////第二种格式
	//循环map中的值打印
	for key, value := range mapp{
		//fmt.Print(key)
		//fmt.Print(value)
		//fmt.Print("," )
		fmt.Printf("%d-%d", key, value)
		fmt.Print(" ")
	}
	fmt.Println()

	///////////////////////////////////////第三种格式，无限循环
	i:=0
	for{
		if(i <= 100){
			fmt.Printf("%d ,",i )
			i++
		}else{
			return
		}

	}

}

//值传递-复制参数值副本，改变函数中不改变原值
func test5(x,y int){

	x , y = y , x
	//以下是把参数副本的地址对应值进行了交换
	//aa := &x
	//bb := &y
	//*aa ,*bb = *bb, *aa
	//fmt.Printf("%d -----%d", x,y)


}

//引用传递-传递指针，改变函数中改变原值
func test6(x,y *int){

	*x , *y = *y , *x

}

//传递方法参数
type mao func(str string) string

func test7(str string, mao mao) string{
	return mao(str)
}


//Go 语言支持匿名函数，可作为闭包。匿名函数是一个"内联"语句或表达式。匿名函数的优越性在于可以直接使用函数内的变量，不必申明
//闭包是匿名函数
func test8() func (int) int{

	return func(i int) int{
		ii := 2
		i+=ii
		return i
	}

}

//方法： 函数和方法不一样，方法包含了接收者的函数
	//自定义一个string类型的type
	type hanshu string

	//给接受者hanshu定义一个将传入字符串拼接的方法
	func (h hanshu) getStringAppend(s string) string{
		 return string(h) + s
	}
	/////////////////////////////////////////////////
	//定义一个正方形结构体
	type square struct {
		length int
	}
	//给结构体square 定义一个求正方形周长变成的方法
	func (square square) getPerimeter() int{
		return square.length << 2
	}


	//数组 --定义数组三种方式
	func test9(){
		var arr1 = []int{1,2,3,4}  //根据初始化值判断长度
		fmt.Println("数组arr1：",arr1 )
		var arr2 = [3]int{5,6,7}  //定义长度为3的数组，并且初始化
		fmt.Println("数组arr2：",arr2 )
		var arr3 []int  //这样定义数组 是nil
		fmt.Println("数组arr3：",arr3 )
		var arr4 [4]int //这样才会定义
		boo := arr3 == nil
		fmt.Println("arr3 是否等于 nil：",boo)
		for i:=0 ; i<4 ; i++{
			arr4[i] = i
		}
		fmt.Println("数组arr4：",arr4)
	}
	//输入传入参数可以定义长度，可以不定义长度
	func test10 (arr1 []int , arr2[3] int) int{
		len1 := len(arr1)
		len2 := len(arr2)
		fmt.Printf("数组arr1长度：%v ,数组arr2长度: %v", len1, len2)
		sum := 0
		for i:=0;i<len1 ;i++  {
			sum +=arr1[i]
		}
		for i:=0;i<len2 ;i++  {
			sum +=arr2[i]
		}
		return sum
	}

	//指针 获取变量内存地址
	func test11(str string) *string{
		return &str
	}