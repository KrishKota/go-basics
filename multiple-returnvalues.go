package main

import (
"fmt"
"strconv"
)

func vals() (int, int) {
    return 3, 7
}

func intToString(a int)string{
return strconv(a)
}

func main() {

    // Here we use the 2 different return values from the
    // call with _multiple assignment_.
    a, b := vals()
    fmt.Println(a)
    fmt.Println(b)

    // If you only want a subset of the returned values,
    // use the blank identifier `_`.
    _, c := vals()
    fmt.Println(c)
    
    stringValue := intToString(c)
    fmt.Println(stringValue)
}
