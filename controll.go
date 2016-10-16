package main




func main(){

    //1.可以省略条件表达式括号
    //2.支持初始化语句，可定义代码块局部变量
    //3.代码块左大括号必须在条件表达式尾部

    x := 0
    if n := "abc"; x > 0 {
       println(n[2])
    } else if x < 0 {
        println(n[1])
    } else {
        println(n[0])
    }

    //不支持三元操作

    s := "abc"
    //常见的for语句，支持初始化语句
    for i,z := 0,len(s);i < z;i++ {
        println(s[i])
    }

    q := len(s)
    for q > 0 {             //替代 while(n > 0) {}
        println(s[q-1])       //替代 for(;n > 0;) {}
        q--
    }

    for {               //while(true) {}
        println(s)      //替代for{;;} {}
    }

}
