package main

import "bufio"
import "flag"
import "fmt"
import "io"
import "os"
import "strconv"
import "time"
import "algorithms/bubblesort"
import "algorithms/qsort"

//flag 设置命令行参数的包，第一个参数是option选项，第二个参数是默认值，第三个参数是参数描述
var infile *string = flag.String("i","unsorted.dat","File contains values for sorting")
var outfile *string = flag.String("o","sorted.dat","File to receive sorted values")
var algorithm *string = flag.String("a","qsort","Sort algorithm")


func readValues(infile string)(values []int,err error) {
    file,err := os.Open(infile)
    if err != nil {
        fmt.Println("Failed to open the input file ",infile)
        return
    }
    defer file.Close()
    br := bufio.NewReader(file)
    values = make([]int,0)
    for {
        //ReadLine返回一行，不包含行结束符，如果一行过长，会设置isPrefix
        //line是一个slice字节数组
        line,isPrefix,err1 := br.ReadLine()
        if err1 != nil {
            if err1 != io.EOF {
                err = err1
            }
            break;
        }

        //提示行过长，
        if isPrefix {
            fmt.Println("A too long line,seems unexpected.")
            return
        }
        //把字节数组转换为string
        str := string(line)
        //字符串转换为数字
        value,err1 := strconv.Atoi(str)
        if err1 != nil {
            err = err1
            return
        }
        values = append(values,value)
    }
    return
}

func writeValues(values []int,outfile string) error {
    file, err := os.Create(outfile)
    if err != nil {
        fmt.Println("Failed to create the output file ",outfile)
        return err
    }
    defer file.Close()
    for _,value := range values{
        str := strconv.Itoa(value)
        file.WriteString(str + "\n")
    }
    return nil
}

func main() {
    flag.Parse()
    if infile != nil {
        fmt.Println("infile =",*infile,"outfile =",*outfile,"algorithm =",*algorithm)
    }
    values, err := readValues(*infile)
    if err == nil {
        t1 := time.Now()
        switch *algorithm {
            case "qsort":
                qsort.QuickSort(values)
            case "bubblesort":
                bubblesort.BubbleSort(values)
            default:
                fmt.Println("Sorting algorithm",*algorithm,"is either unknown or unsupported.")
        }
        t2 := time.Now()
        fmt.Println("The sorting process costs",t2.Sub(t1),"to complete")
        writeValues(values,*outfile)

    } else {
        fmt.Println(err)
    }
}
