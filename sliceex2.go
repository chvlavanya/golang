/* Creating a slice from an array */

package main

import "fmt"

func main() {

    myArray := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9} // we need an array

    mySlice1 := myArray[1:5] // slice from the 1st up to 5th

    mySlice2 := mySlice1[0:2] // slicing a slice

    fmt.Println(mySlice1)

    fmt.Println(mySlice2)
    
    mySlice2[1] = 22          // updating mySlice2 will also update mySlice1 and myArray
    
    fmt.Println(mySlice1)     // this now reads [1 22 3 4]
    fmt.Println(mySlice2)     // this now reads [1 22 3 4]
    fmt.Println(myArray)      // this now reads [0 1 22 3 4 5 6 7 8 9]
                              //      updating the slice updates the array

}

