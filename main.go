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
    switch x.(type) { // Get the type of x and check whether if it belongs to the supported types
    case int, float64, uint: // If true, print the number
        fmt.Println(x)
    default: // If false, print "x is not a number"
        fmt.Println("x is not a number")
    }
}

func main() {
    var x uint = 12
    printNumber(x) // outputs 12
    printNumber("Hello World") // outputs "x is not a number"
}
