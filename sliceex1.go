/* Array Management */

package main

import "fmt"

func main() {
    myArray1 := [2]string{"a", "b"}
    fmt.Printf("myArray1 Before: %v\n", myArray1)
    myArray2 := myArray1 // this creates a new array
    myArray2[1] = "c"
    fmt.Printf("myArray1 After assignment: %v\n", myArray1)
    fmt.Printf("myArray2: %v\n", myArray2)
    test(myArray1) // this creates a new array
    fmt.Printf("myArray1 After Test Function Call: %v\n", myArray1)
}
func test(myArray [2]string) {
    myArray[0] = "d"
    fmt.Printf("myArray in Test function: %v\n", myArray)
}

