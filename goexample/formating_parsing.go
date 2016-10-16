package main
import "fmt"
import "time"

func main() {
    p := fmt.Println
    t := time.Now() //要打印的时间
    p(t.Format(time.RFC3339))   //按照RFC3339格式化
    t1,e := time.Parse(         //按照RFC3339格式化指定时间
        time.RFC3339,
        "2012-11-01T22:08:41+00:00")
    p(t1)                       //打印
    p(t.Format("3:04PM"))
    p(t.Format("Mon Jan _2 15:04:05 2006"))
    p(t.Format("2006-01-02T15:04:05.999999-07:00"))
    form := "3 04 PM"
    t2,e := time.Parse(form,"8 41 PM")
    p(12)


    fmt.Printf("%d-%02d-%02dT%02d:%02d-00:00\n",
        t.Year(),t.Month(),t.Day(),
        t.Hour(),t.Minute(),t.Second())

    ansic := "Mon Jan _2 15:04:05 2006"
    _,e = time.Parse(ansic,"8:41PM")
    p(e)
}
