package main
import "fmt"
import "sort"

func main() {
    //待排序的序列
    strs := []string{"c","a","b"}
    //sort package，针对strig,int的排序
    sort.Strings(strs)
    fmt.Println("strings",strs)

    ints := []int{7,2,4}
    sort.Ints(ints)
    fmt.Println("Ints:  ",ints)
    //判断int类型的序列是否有序
    s := sort.IntsAreSorted(ints)
    fmt.Println("Sorted: ",s)
}
