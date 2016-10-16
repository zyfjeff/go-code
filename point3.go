package main

func test() *int {
    x := 100
    return &x   //返回局部变量指针是安全的，编译器根据需要将其分配在GC Heap，或者直接分配在目标栈上　
}

func main(){

}
