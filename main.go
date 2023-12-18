package main

import (
    "os"
    "fmt"
)
import "strconv"

type LineReader = func(string)

func main(){
    args := os.Args[1:]
    if len(args) == 0 {
        fmt.Println("You must pass the dayNumber argument")
        return
    }
        
    dayNumber, err := strconv.Atoi(args[0])
    if err != nil {
        fmt.Println("You need to pass a valid day number")
        return
    }

    switch dayNumber {
        case 1:
            RunDay1()
        case 2:
            RunDay2()
        case 3:
            RunDay3()
        case 4:
            RunDay4()
        case 5:
            RunDay5()
        default:
            fmt.Println("This day is not yet implemented")
    }
}

