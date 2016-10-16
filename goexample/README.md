## Introduce
这是来自于[go by example](https://gobyexample.com/)的例子，花了几天的时间写完了这些例子，感觉对我的帮助很大，对于
初学者来说，我的建议还是先找本go的书从头到尾看一下，然后再来看这些例子，每个例子都手敲一遍，对你的帮助还是很大的。
在敲这些例子的过程中，有一些疑问，也有一些知识的扩充，因此总结了本文。

## 你不知道的打印输出
在go中fmt包功能很强大，里面提供了Print，Println等打印方法，支持类似于C语言的格式化输出，最重要的是fmt包可以识别任意类型
功能很强大。下面是利用fmt来打印一个自定义的结构体。

```
package main
import "fmt"


type custom_struct struct{
    name string
    age int
}


func main() {
    man := custom_struct{"zhang",10}
    vars := []string{"first","second","last"}
    fmt.Println(man)
    fmt.Println(vars)
}
```
上面的代码可以直接打印custom_struct，还可以直接打印数组等，功能还是很强大的，但是却很少有人知道，go语言在内部也有一套
print，println方法用来打印你的输出，用的人很少，这些打印方法一般用于go内部debug使用，不支持打印自定义类型。

## 大容量的常量
在go语言中，数字可以直接写成指数形式，例如`3e20`，而且go中的常量可以接收任意精度的常量表达的值，例如下面这个例子。

```
package main
import "fmt"


func main() {
    const n = 5000
    const d = 3e100 / n
    fmt.Println(d)
}
```
常量d可以接收任意精度的值，除此之外常量还有另外一个特点就是无类型，可以通过显式的转换赋予其类型，或者通过上下文赋予
常量类型，例如赋值操作，函数调用，那么常量的类型就变成了赋值操作符左侧的类型，或者变成了函数调用的形参类型了。

## 无所不能的for
在go中，for循环可以是C中的while循环，for循环，一个关键字完成所有的循环操作，还支持类型python中的range的循环操作。

```
func main() {
    i := 1
    nums := []int{2,3,4}
	//while循环
    for i <= 3 {
        fmt.Println(i)
        i = i + 1
    }
	
	//普通的for循环
    for j := 7; j <= 9; j++ {
        fmt.Println(j)
    }

    //基于range的循环
    for i,num := range nums {
        if num == 3 {
            fmt.Println("index:",i)    
        }    
    }

	//死循环
    for {
        fmt.Println("loop")
        break
    }
}
```

## if还可以这样子
if语句大体上来看，和C中的if在使用上差别不大，只不过go语言中的if支持在条件的前面加上一个赋值表达式，刷新了我的三观，并且
在go中没有像C中的那个问号表达式了，下面是带赋值表达式的if

```
func main() {
    //这里是可以前置一个赋值表达式的
    if num := 9; num < 0 {
        fmt.Println(num,"is negative")
    } else if num < 10 {
        fmt.Println(num,"has 1 digit")
    } else {
        fmt.Println(num,"has multiple digits")
    }
}
```

## 不一样的switch
自从学习了golang，不断刷新我三观，这和我之前接触的C/C++，Python差别不是一般的大啊，在go中switch也同时兼备了正常和不正常的地方
正常使用的情况和C中的switch基本一样，但是go中的switch没有穿透的概念，意思就是说，匹配到指定的项后就结束停止了，不再向下。go
中的switch是不需要显示的break的。除此之外在和C中的switch别无差别，但是这只是go switch正常的一面，不正常的一面则是go的switch
支持多条件匹配,支持表达式的匹配。

```

    switch time.Now().Weekday() {
    case time.Saturday,time.Sunday:     //用逗号分割的多个条件，匹配任意一个都可
        fmt.Println("it's the weekend")
    default:
        fmt.Println("it's a weekday")
    }

    //类似与if/else的感觉
    t := time.Now()
    switch {
    case t.Hour() < 12:     //为true就执行下面的语句,为false就直接跳到default分支
        fmt.Println("it's before noon")
    case true:              //继续执行,如果这里是false，那么就跳转指向default分支
        fmt.Println("continue")
    default:
        fmt.Println("it's after noon")
    }
```

## arrays和slice不是一回事
go 中的arrays和C中的数组是一样的概念，而slice和python的列表是一个概念，只不过go中的slice必须是同一类型的而已。arrays是固定长度
而slice则不是固定长度，可以通过append方法向其追加元素，但是在go中却没有单独给slice配备一个关键字，这导致有的时候不知道怎么区分
arrays和slice。

```
arrays := [5]int{1,2,3,4,5} //这是一个数组，指明长度了
slices := []int{1,2,3}      //这是一个slice
slices2 := make([]int,3)    //这样也可以是一个slice
```
那么在go中slice是如何实现的呢?在讨论这个问题的时候，让我们来先看看go中的数组的一些实现细节，因为slice就是对array的抽象。
go中的数组和C中的数组类似，但还是有一些不一样的地方，比如go的数组变量代表的就是整个数组，而不是像像C那样是一个指针，指向
数组的首元素，因此在go中赋值和传递一个数组变量会导致对数组中的全部元素进行拷贝。数组在定义的时候需要指定长度，也可以让
编译器帮我们计算，通过如下代码实现:

```
b := [...]string{"test","dwqd"}
```

slice不需要指定长度，这是和数组最明显的区别，那么这个slice内部到底是什么样子的呢，是和c++中的vector一样吗? slice其实就是对
array的一个包装和抽象，每个slice内部都会对应一个数组，这个slice维护一个指针指向这个数组的首元素，维护这个数组的长度len,
还会记录数组的最大长度capacity。那么这个len和capacity是什么关系呢? 和C++的vector中的capacity一个意思吗?，恰恰不是。限于
篇幅这里我不再过多介绍slice的内部实现细节了，给大家推荐一篇博客[go-slice-and-internals](https://blog.golang.org/go-slices-usage-and-internals)


## 多返回值的奇妙体验
限于我之前一直写C，常常需要返回多个值的时候，通常都是借助与函数参数来实现，好在go提供了多返回值的能力。这样就极大了简化了
代码，在返回结果的同时，也返回出现的error，但是需要记住的是，只能返回两个值哦，千万不要贪多。

## 受限的指针
之所以选择学习go，主要是感觉go是一个类C的语言把，熟悉的struct，熟悉的指针，但是go中的指针是受限的，不支持算术运算因为需要
支持垃圾回收。通过指针作为函数参数，在函数内部修改会影响到原始的变量，默认情况下是值拷贝的，在函数内部修改不会影响原始变量。

## 像class的struct
go中的struct也是用来自定义数据类型的，需要需要以面向对象的方法来编程，就需要将方法和数据封装在一起，C中通过在struct中包含
函数指针来实现，go则不然，原生支持，直接给struct定义方法。

```
package main

import "fmt"

type rect struct {
    width,height int
}

//给rect自定义的数据类型，定义了两个方法
func (r *rect) area() int {
    return r.width * r.height
}

func (r rect) perim() int {
    r.width = 100;
    return 2 * r.width + 2*r.height
}


func main() {
    r := rect{width: 10,height: 5}
    fmt.Println("area: ",r.area()) //r 是一个值类型，可是area的参数类型是指针啊，这居然可以调用 
    fmt.Println("perlim: ",r.perim())
    fmt.Println(r)

    rp := &r
    fmt.Println("area: ",rp.area())
    fmt.Println("perlim: ",rp.perim())

}
```
在上面的例子中给rect定义了两个方法，一个是指针作为第一个参数，另一个则是值作为参数，但是在使用的时候居然可以不加区别的使用。
这其实都是go在背后帮我们做了转换，但是不影响最终的效果，指针作为参数的，在函数内部操作会直接影响到原有的变量。

## Interface神器
接口是什么?，说白了就是一堆方法的集合。C++没有接口，但是可以通过纯虚函数来实现一个接口，在java中原生就有接口的概念，但是我想说的是golang的
接口要比这两者都要强大。无论是C++还是java对于一个具体类需要实现对应接口的所有方法才能被这个接口抽象，而golang不需要，golang中你只需要实现
接口的至少一个方法即可。比如说对于经典企鹅是鸟但不会飞的这个问题，C++和java都没办法让企鹅实现一个鸟类的接口，但是golang确可以简化这个设计，
只需要实现鸟类接口中除了飞翔这个方法外的其他方法。在golang中还有一个空接口`interface{}`，这个可以代表任意类型，如果将空接口作为函数参数
那么这个函数可以接收任意类型的数据。但是空接口存在一个设计很ugly的地方，请看下面的这段代码:

```
func PrintAll(vals []interface{}) {
    for _,val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"test1","test2","test3"}
    PrintAll(names)
}
```
根据上面的描述，空接口应该是可以接受任意类型的数据的，上面的代码理论上可以正常运行，但是我不的不说我撒谎了，直接编译运行会有下面的错误:

```
cannot use names (type []string) as type []interface {} in argument to PrintAll
```
两者居然不能直接做转换，不是说空接口可以接收任意类型的数据的吗?，的确是的，不过上面的代码需要改一改。

```
func PrintAll(vals []interface{}) {
    for _,val := range vals {
        fmt.Println(val)
    }
}

func main() {
    names := []string{"test1","test2","test3"}
    vals := make([]interface{},len(names))
    for i,v := range names {
        vals[i] = v
    }
    PrintAll(vals)
}

```
现在可以正常工作了，偏要先赋值给空接口。这设计的确有点ugly。golang的接口还有另外一个不错的地方就是，你可以设计成接收指针的方法
也可以设计成接收value的方法，这都是ok的。但是这里面依然存在一些误区，请看下面这段代码:

```
type Animal interface {
    Speak() string
}

type Dog struct {

}

func (d Dog) Speak() string {
    return "Woof"
}

type Cat struct {

}

func (c Cat) Speak() string {
    return "Meow"
}

func main() {
    animals := []Animal{Dog{},Cat{}}
    for _,animal := range animals {
        fmt.Println(animal.Speak())
    }
}
```
上面的代码声明了Animal接口，然后Dog和Cat实现了这个接口。然后在main函数中通过Animal接口对Dog和Cat进行抽象，依次调用对应的Speak方法。
上面的代码工作正常，直到有一天，我将Cat的Speak方法由原来的`(c Cat)`改成了`(c *Cat)`接收一个指针。再次运行上面的代码，出现了错误:

```
cannot use Cat literal (type Cat) as type Animal in array element:
	Cat does not implement Animal (Speak method has pointer receiver)
```
传入的`Cat{}`是一个value，而Cat的Speak方法是接收一个指针。好吧，改动下代码把`Cat{}`改成`new(Cat)`，现在可以正常工作了，现在试着
将`(c *Cat)`改回`(c Cat)`然后依然传递一个Cat指针给这个方法。你会发现这居然可以。最后得出的结论就是，值类型不能传递给接收指针的方法
但是指针类型却可以传递给接收值的方法。是不是感觉设计上有点ugly。但这都是有原因的，更多的内容可以参考下面这两篇文章:

* [How to use interfaces in Go](http://jordanorelli.com/post/32665860244/how-to-use-interfaces-in-go)
* [Go Data Structures: Interfaces](http://research.swtch.com/interfaces)

## 专属的错误处理方式
在golang中错误处理都是通过返回值进行，golang内置一个error的接口，一般情况下你直接通过`errors.New`就可以产生一个带着指定错误消息的
error value，如果你想给自己的类型设计一个错误类型，只需要自定义一个错误类型，然后实现`Error`接口即可，代码如下:

```
package main

import "errors"
import "fmt"

//返回一个错误,通过errors.New创建一个错误对象
func f1(arg int) (int,error) {
    if arg == 42 {
        return -1,errors.New("Can't work with 42")
    }

    return arg + 3,nil
}

type argError struct {
    arg int
    prob string
}

//给结构体定义了一个错误
func (e *argError) Error() string {
    return fmt.Sprintf("%d - %s",e.arg,e.prob)
}

//只要实现了Error方法都可以通过error接受，error就是一个接口
func f2(arg int) (int,error) {
    if arg == 42 {
        return -1,&argError{arg,"can't work with it"}
    }
    return arg + 3,nil
}

func main() {
    //遍历slice，对每个值调用f1和f2的方法
    for _,i := range []int{7,42} {
        if r,e := f1(i); e != nil {
            fmt.Println("f1 failed:",e)
        } else {
            fmt.Println("f1 worked:",r)
        }
    }
    for _,i := range []int{7,42} {
        if r,e := f2(i); e!= nil {
            fmt.Println("f2 failed:",e)
        } else {
            fmt.Println("f2 workded:",r)
        }
    }
}
```
除了可以自定义error类型，还可以对error接口进行扩充，有关更多的关于error的故事可以参考下面这篇文章:

[Error handling and Go](https://blog.golang.org/error-handling-and-go)

## CSP模型和channel
golang中的Goroutine是典型的CSP模型，每一个协程就是一个work，各个work之间通过刚channel进行通信，channel就好比是一个通道，类似于
Unix的Pipe,和Actor模型很相似。golang中的channel默认是没有buffer的，也就是说，一次只能发送一条消息，并且只有等对方接收到了消息后
才可以再次发送，如果channel中没有消息那么会导致阻塞，知道有消息为止。

```
done := make(chan bool,1)   //创建了一个bool类型的channel，buffer大小是1
done <- true                //向channl发送消息
flag := <-done              //从chanlle中接收消息
```
默认情况下channel是全双工的，但是你可以设置channel只能发或者只能接收。

```
pings <-chan string         //只能发的channel
pongs chan<- string         //只能收的channel
```

## select IO多路服用
写过网络编程的都知道IO多路服用，也就是同一时间内可以监听多个套接字，golang在语言层面提供了select关键字，不过这不是用来监听套接字的，
而是用来同时监听多个channel。下面是一个select使用的例子:

```
package main

import "time"
import "fmt"

func main() {
    c1 := make(chan string)
    c2 := make(chan string)
    fmt.Println(time.Second)
    //两个协程，通过c1和c2通信
    go func() {
        time.Sleep(time.Second * 1)
        c1 <- "one"
    }()

    go func() {
        time.Sleep(time.Second * 2)
        c2 <- "two"
    }()

    //select监控多个描述符
    for i := 0;i < 2;i++ {
        select {
        case msg1 := <-c1:
            fmt.Println("received",msg1)
        case msg2 := <-c2:
            fmt.Println("received",msg2)
        }
    }
}
```
默认的channel读写是阻塞的，可以通过结合select变成非阻塞,通过在select中加入default，当上面所有的channel都没有消息的时候，会立即
跳转到default。这样就巧妙的实现了非阻塞。除此之外channel还可以被关闭，当channel被关闭了，说明没有消息要发送过了，此时如果去读channel
会立即返回非空的error，channel还可以使用for range来进行遍历，前提是需要将channel关闭，否则在遍历的时候会阻塞。


## time和channel
golang的time package带有定时器的功能，而定时器和channel完美融合，创建一个定时器会返回一个channel，在定时器到期之前读这个channel
是阻塞的，直到定时时间到达这个channel就会变成可读的。

```
func main() {
    //创建了一个定时器，2秒后会发送事件到timer1.C channel
    timer1 := time.NewTimer(time.Second * 2)

    //等待定时器到期
    data := <-timer1.C          //接收到的数据 2016-07-16 15:24:19.337701998 +0800 CST
    fmt.Println("Timer 1 expired")
    fmt.Println("Timer 1 expired",data)
    timer2 := time.NewTimer(time.Second)
    go func() {
        <- timer2.C
        fmt.Println("Timer 2 expired")
    }()

    //关闭定时器
    stop2 := timer2.Stop()
    if stop2 {
        fmt.Println("Timer 2 stopped")
    }
}
``` 

time package除了具有定时器的功能外，还有一个Ticker，Ticker同样也是和channel完美融合的一个功能，创建一个Ticker会返回channel
通过range这个channel。来表示每次interval的到来。再结合go的协程就很容易实现一个定时任务的功能。

```
func main() {
    //创建了定时器,每time.Millisecond * 500就产生事件，发送到chnanel
    ticker := time.NewTicker(time.Millisecond * 500)
    go func() {
        for t := range ticker.C {
            fmt.Println("Tick at",t)
        }
    }()
    //睡上一段事件，然后关闭
    time.Sleep(time.Millisecond * 1600)
    ticker.Stop()
    fmt.Println("Ticker stopped")
}
```

## goroutines和work pool
goroutines结合channel很容易就可以实现一个work pool，开上N个goroutines，然后这N个goroutines共同去读channel，读到channel
就去执行相应的工作，然后结果通过另外一个channel传出来即可，模型很简单。用go实现起来还是很容易的。

```
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
```

## rate limiting与channel
限速这是一个用于控制资源利用率和保证服务质量的一种机制，golang通过goroutines,channel还有tickers优雅的支持了这个机制。
比如处理web请求的限速，每接收一个请求就先读取一个tick，这个tick每隔固定时间才可读，这样就可以实现限速的功能。

```
func main() {
    requests := make(chan int,5)
    //发送五条消息
    for i := 1; i <= 5;i++ {
        requests <- i
    }

    close(requests)
    //创建了定时器，然后遍历channel，打印信息
    limiter := time.Tick(time.Millisecond * 200)
    for req := range requests {
        <-limiter   //每隔time.Millisecond * 200，起到了限速的作用
        fmt.Println("requests",req,time.Now())
    }

}
```

但是上面的限速存在一个问题，就是并发数只有1，如果可以在拥有固定的并发数的情况下限速呢?，这就需要借助channel的buffer功能了。
上面的time.Tick返回的channel是没有buffer的，所以一次只能处理一个请求，如果这个channel是有buffer的，比如这个buffer的大小是N
那么可以同时并发接收N个请求，想处理第N+1个请求就需要等待固定时间才可以。

```
func main() {
    //创建了另外一个time.Time类似的channel
    burstyLimiter := make(chan time.Time,3)
    //发送三个
    for i := 0; i < 3; i++ {
        burstyLimiter <- time.Now()
    }

    go func() {
        for t := range time.Tick(time.Millisecond * 200) {
            burstyLimiter <- t //每200 * time.Millisecond 就发送一个事件到burstyLimiter
        }
    }()


    burstyRequests := make(chan int,5)
    for i := 1; i <= 5; i++ {
        burstyRequests <- i //发送五个数据
    }
    close(burstyRequests)
    for req := range burstyRequests { //现在开始限速读取
        <-burstyLimiter//在读前三个的时候是不会阻塞的，直到读取第四个的 时候才开始通过Limiter限速
        fmt.Println("request",req,time.Now())
    }
}
```

## 自定义sort和Interface
golang的sort package自带排序的功能，但是如果要对用户自己定义的数据结构进行排序这就不好半了，在C++中要求用户对关系运算符重载即可
在golang中则需要和interface完美融合，只要用户实现Len,Less,Swap三个接口即可,就是这么简单。

```
package main

import "fmt"
import "sort"

//string slice的别名，给这个别名struct 添加方法
type ByLength []string

func (s ByLength) Len() int {
    return len(s)
}

func (s ByLength) Swap(i,j int) {
    s[i],s[j] = s[j],s[i]
}

func (s ByLength) Less(i,j int) bool {
    return len(s[i]) < len(s[j])
}

//sort接口需要实现 Swap Less和len即可
func main() {
    fruits := []string{"peach","banana","kiwi"}
    sort.Sort(ByLength(fruits))
    fmt.Println(fruits)
}
```

## signal和channel
golang再一次将unix上的signals和channel结合了起来，unix上通过给信号注册处理函数来完成信号处理，在golang中，通过把信号
和channel关联起来，当有信号到来channel就可读了。返回的结果就是signal的号码。

```
import "fmt"
import "os"
import "os/signal"
import "syscall"

func main() {
    //os.Signal类型的chn
    sigs := make(chan os.Signal,1)
    done := make(chan bool,1)
    //通过Norify来注册信号，
    signal.Notify(sigs,syscall.SIGINT,syscall.SIGTERM)  //将SIGINT和SIGTERM和sigs channel结合起来

    //协成来收集信号，然后发送done chan来表示完成
    go func() {
        sig :=  <-sigs
        fmt.Println()
        fmt.Println(sig)
        done <- true
    }()

    fmt.Println("awaiting signal")
    <-done
    fmt.Println("exiting")
}
```
