package main

func main(){
/*
    var a struct {x int} = 100 //syntax error
    var b []int = {1,2,3}
    c := struct {x int;y string}
    {
    }
*/
//初始化复合对象，必须使用类型标签,且左大括号必须在类型的尾部
//多个成员可以换行，但是最后一个必须接","号，或者"}"号
    var a = struct{x int}{100}
    var b = []int{1,2,3}

}
