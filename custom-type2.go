package main

//类型分为两种，第一种是命名类型 bool int string
//第二种则是array slice map 和具体元素类型，长度有关属于未命名类型
func main(){
    type bigint int64
    var x bigint = 100
    println(x)
    z := 1234
    var b bigint = bigint(z) //必须显示转换,除非是常量
    var b2 int64 = int64(b)

    var s myslice = []int{1,2,3}    //未命名类型,隐式转换
    var s2 []int  = s
}
