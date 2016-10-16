package bubblesort

func BubbleSort(values []int) {
    flag := true
    //len(values) - 1趟排序，
    for i := 0; i < len(values) - 1; i++ {
        flag = true
        //每次拿到一个数和后面的比较，如果大于后者就交换，这样最大的数就被替换到最后面了
        for j := 0; j < len(values) - i -1;j++ {
            if values[j] > values[j + 1] {
                values[j], values[j + 1] = values[j + 1],values[j]
                flag = false
            } //end of if
        }   //end of for j=
        if flag == true {
            break
        }
    }   //end of for i=
}
