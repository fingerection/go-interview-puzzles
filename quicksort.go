package main

import(
    "fmt"
    "math/rand"
)

const (
    SIZE = 1000
)

func main() {
    fmt.Println("Start")
    l := make([]int, SIZE)
    for i:=0; i<SIZE; i++ {
        l[i]= rand.Intn(100000)
    }
    fmt.Println(l)
    quicksort(l,0,len(l)-1)
    fmt.Println(l)
}


func quicksort(l []int, start int, end int) {
    //fmt.Println("qs:", l, start, end)
    if start >= end {
        return
    }
    j := start
    n := l[start]
    for i:=start+1; i<=end; i++ {
        if l[i] <= n {
            l[j], l[i] = l[i], l[j]
            j++
        }
    }
    l[j] = n
    //fmt.Println("result:", l, start, j, end)
    quicksort(l, start, j-1)
    quicksort(l, j+1, end)
}
