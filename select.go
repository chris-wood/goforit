package main

import "time"
import "fmt"
import "strconv"
import "math/rand"

func main() {

    c1 := make(chan bool)
    c2 := make(chan int)
    c3 := make(chan string)

    var input string // uninitialized
    fmt.Scanln(&input)
    waitStep, _ := strconv.Atoi(input)
    wait := int32(waitStep)

    rand.Seed(42)

    // start each in a separate goroutine
    go func() {
        time.Sleep(time.Second * time.Duration(rand.Int31n(wait)))
        c1 <- true
    }()
        go func() {
        time.Sleep(time.Second * time.Duration(rand.Int31n(wait)))
        c2 <- 2
    }()
    go func() {
        time.Sleep(time.Second * time.Duration(rand.Int31n(wait)))
        c3 <- "three"
    }()

    // see which one gets here first
    select {
    case msg1 := <-c1:
        fmt.Printf("received: %t\n", msg1)
    case msg2 := <-c2:
        fmt.Printf("received: %d\n", msg2)
    case msg3 := <-c3:
        fmt.Printf("received: %s\n", msg3)
    }
}

