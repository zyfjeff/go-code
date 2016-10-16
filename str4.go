package main


//字符串是不可变的，要修改字符串将其转换为[]rune或[]byte，完成后再转换为string
//无论那种转换都会重新分配内存，并复制字节数组
func main(){
    s := "abcd"
    bs := []byte(s)
    bs[1] = 'B'
    println(string(bs))

    u := "电脑"
    us := []rune(u)
    us[1] = '话'
    println(string(us))
    //println(us)
}
