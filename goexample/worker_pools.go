package main
import "fmt"
import "time"

//一个worker,接收jobs,通过result返回结果
// <-chan 接收
// chan<- 发送
func worker(id int,jobs <-chan int,result chan<- int) {
    for j := range jobs {
        fmt.Println("worker",id,"processing job",j)
        time.Sleep(time.Second)
        result <- j * 2
    }
}

func main() {
    jobs := make(chan int,100)
    results := make(chan int,100)
    //启动三个worker
    for w := 1; w <= 3; w++ {
        go worker(w,jobs,results)
    }
    //循环9次，发送任务
    for j := 1; j <= 9;j++ {
        jobs <- j
    }
    close(jobs)
    //循环得到结果
    for a := 1; a <= 9; a++ {
        <-results
    }
}
