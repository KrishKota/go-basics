package main

import "fmt"

func addString(nums ...string) {
    fmt.Print(nums, " ")
    total := []string{}
    for _, num := range nums {
        total = append(total,num)
    }
    fmt.Println(total)
}

func main() {

    // Variadic functions can be called in the usual way
    // with individual arguments.
    addString("hello", "world")
    addString("1", "2", "3")

    nums := []string{"1", "2", "3", "4"}
    addString(nums...)
}
