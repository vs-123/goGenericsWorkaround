package main

import (
    "fmt"
)

func contains(x []string, y string) bool {
    for _, i := range x {
        if i == y {
            return true
        }
    }
    return false
}

func printNumber(x interface{}) {
    // Get the type of x and check whether if it's in the string slice that contains all supported types for this function
    switch x.(type) {
    case int, float64, uint:
        fmt.Println(x)
    default:
        fmt.Println("x is not a number")
    }
}

func main() {
    var x uint = 12
    printNumber(x) // outputs 12
    printNumber("Hello World") // outputs "x is not a number"
}
