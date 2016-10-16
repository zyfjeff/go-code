package qsort

func quickSort(values []int,left,right int) {
    temp := values[left] //确定主元
    p := left
    i, j := left,right
    for i <= j {
        //从右边开始，找到第一个小于主元的位置
        for j >= p && values[j] >= temp {
            j--
        }
        //交换到最左边
        if j >= p {
            values[p] = values[j]
            p = j
        }
        //从左边开始，找到第一个
        if values[i] <= temp && i <= p {
            i++
        }

        if i <= p {
            values[p] = values[i]
            p = i
        }
    }
    values[p] = temp
    if p - left > 1 {
        quickSort(values,left,p - 1)
    }

    if right - p > 1 {
        quickSort(values,p + 1,right)
    }
}

func QuickSort(values []int) {
    quickSort(values,0,len(values) - 1)
}
