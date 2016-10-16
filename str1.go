package main



//默认值是空字符串
//用索引号访问某字节 如s[i]
//不能用序号获取字节元素指针 &s[i] 是非法的
//不可变类型,无法修改字节数组
//字节数组尾部不包含NULL


func main(){
    s := "abc"
    z := `a
b\r\n\x00
c`
    //定义不转义的字符
    println(z)

    println(s[0] == '\x61',s[1] == 'b',s[2] == 0x63)
}
