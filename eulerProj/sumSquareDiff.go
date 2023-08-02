package main

import "fmt"


func sumSquareDiff(num int) (int, int, int) {
    
    s := 0
    sq := 0

    for i := 0; i <= num; i++ {
        s += i*i
        sq += i
    }

    sq = sq*sq
    diff := sq - s

    return s, sq, diff 
}


func main() {

    s, sq, diff := sumSquareDiff(100)
    fmt.Printf("The sum of squares is %d \n the square of sums is %d, and the difference is %d", s, sq, diff)
}
