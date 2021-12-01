/* Alta3 Research | RZFeeser
   Flags - Strings, Integers, Booleans
   and alternatives techniques to parsing the CLI    */
   
package main

import (
    "flag"              // basic CLI parsing
    "fmt"
)

func main() {

    // these all return pointers
    wordPtr := flag.String("word", "Don't panic!", "a string")
    numbPtr := flag.Int("numb", 42, "an int")
    boolPtr := flag.Bool("truthy", false, "a bool")

    // example of passing a pointer to one of the "flag.*Var()" functions
    var spiderman string    // define the string variable "spiderman"
    flag.StringVar(&spiderman, "spiderman", "Friendly neighborhood Spider-Man", "Catchphrase for Spider-Man")

    flag.Parse()

    // display the values captured
    fmt.Println("word:", *wordPtr)
    fmt.Println("numb:", *numbPtr)
    fmt.Println("truthy:", *boolPtr)
    fmt.Println("spiderman:", spiderman)
    fmt.Println("tail:", flag.Args())
}
//go build ~/flags03.go
// go run ~/flags03.go
//~/flags03 --help(or) ~/flags03 -word=carrot --truthy=true -numb 8675309 (or) ~/flags03 --spiderman="Giving up crime fighting to program GoLang" toothbrush mapleleaf toaster
