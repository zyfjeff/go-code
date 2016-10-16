package main
var x int
var f float32 = 1.6
var s = "abc"

//一次性定义多个变量
var w,e,r int
//q是abc,n是123
var q,n = "abc",123
var(
    a int
    b float32
)

// 这种声明方式只能在函数内部 z := "liu"
//未使用的局部变量是个错误

func test()(int,string) {
    return 1,"abc";
}

func main(){
    _,_ = test(); //使用特殊变量_忽略
    a,b := 1,23.4; //多变量赋值
    //多变量赋值，先计算所有相关值，然后再从左到右依次赋值
    z,f := [3]int{0,1,2},0.2
    println(n);
    println(a);
    println(b);
    println(z[0]);
    s  := "wang" //容易造成异或，:=可以在函数内部定义变量，也可以是修改全局变量
    println(s);
    f = 1.2
    println(f)
    println("Hello","World");
}


