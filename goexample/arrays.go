package main
import "fmt"

func main() {
    var a [5]int    //初始化的一个数组
    fmt.Println("em:",a)    //0 0 0 0 0

    a[4] = 100              //基于下标访问
    fmt.Println("set:",a)
    fmt.Println("get:",a[4])

    fmt.Println("len:",len(a))  //len求长度
    b := [5]int{1,2,3,4,5}      //5个元素的数组
    fmt.Println("dc1:",b)       //打印

    var twoD [2][3]int          //二维数组打印
    for i := 0; i < 2; i++ {
        for j := 0; j < 3;j++ {
            twoD[i][j] = i + j
        }
    }
    fmt.Println("2d: ",twoD)
}
