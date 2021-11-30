/* Create a new struct, of type nasaMission. To create this type, include people mapped to []astro, number mapped to int, and message mapped to string. Once you've created this, create a nasaMission struct named s. Initialize people as p, init as 3, and message as success. Display s (with and without its key fields) to end your program. */

package main

import "fmt"

type astro struct {
    name     string
    age      int
    mission  string
    isneeded bool
}

// this is our new struct
type nasaMission struct {
        people []astro
        number int
        message string
}

func main() {

    ast1 := astro{"Megan McArthur", 35, "ISS", false}
    ast2 := astro{"Barry Wilmore", 41, "Boeing Crew Flight Test (CFT)", true}
    ast3 := astro{"Raja Chari", 39, "SpaceX Crew-3", true}

    fmt.Println(ast1)
    fmt.Println(ast2)
    fmt.Println(ast3)

    //   name   slice of astro
        p := []astro{ast1, ast2, ast3}

    fmt.Println(p)

    // select data from the struct
    fmt.Println(p[1].mission)

    // initialize a nasaMission struct, "s"
    s := nasaMission{p, 3, "success"}

    // display "s" without fields
    fmt.Println(s)

    // display "s" with fields
    fmt.Printf("%+v", s)

}
